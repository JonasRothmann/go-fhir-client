package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode"

	stderrors "errors"

	"github.com/dave/jennifer/jen"
	"github.com/jonasrothmann/go-fhir-client/custom"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/stoewer/go-strcase"
)

var httpClient = http.DefaultClient
var cache = NewCache()

func init() {
	httpClient.Transport = NewThrottledTransport(time.Second, 1200, http.DefaultTransport)
}

type enumOptions struct {
	Name    string
	Value   string
	Comment *string
}

type enum struct {
	Name    string
	Options []enumOptions
	Comment *string
}

func (e enum) generateStatement() *jen.Statement {
	enumName := e.Name
	statement := &jen.Statement{}
	if e.Comment != nil {
		statement.Comment(*e.Comment)
	}
	statement.Type().Id(enumName).String().Line().Line()

	if len(e.Options) == 0 {
		return statement
	}

	optionCodes := make([]jen.Code, len(e.Options))
	for i, option := range e.Options {
		name := fmt.Sprintf("%s%s", enumName, strcase.UpperCamelCase(option.Name))
		statement := ptr(jen.Statement([]jen.Code{}))
		if option.Comment != nil {
			statement.Comment(*option.Comment).Line()
		}
		optionCodes[i] = statement.Id(name).Id(e.Name).Op("=").Lit(option.Value)
	}
	statement.Const().Defs(optionCodes...)

	return statement
}

func (vs ValueSet) generateComment() *string {
	parts := []string{}

	if vs.Url != nil {
		parts = append(parts, string(*vs.Url))
	}

	if vs.Description != nil {
		parts = append(parts, string(*vs.Description))
	}

	if vs.Url != nil {
		parts = append(parts, string(*vs.Url))
	}

	lines := []string{}
	for _, inc := range vs.Compose.Include {
		if inc.System == nil {
			log.Debug().Interface("inc", inc).Msg("missed include")
			continue
		}

		if len(lines) == 0 {
			lines = append(lines, "Composed of following code systems:")
		}

		lines = append(lines, fmt.Sprintf("- %s", *inc.System))

		if len(inc.Concept) > 0 {
			for _, concept := range inc.Concept {
				lines = append(lines, fmt.Sprintf(" - %s", concept.Code))
			}
		}
	}

	if len(lines) > 0 {
		parts = append(parts, lines...)
	}

	if len(parts) == 0 {
		return nil
	}

	return ptr(strings.Join(lines, "\n"))
}

func (vs ValueSet) generateEnum(ctx context.Context) (e enum, err error) {
	codesystems, err := vs.getCodeSystems(ctx)
	if err != nil {
		return e, err
	}
	if len(codesystems) == 0 {
		return e, errors.New("No codesystems supplied")
	}

	e.Name = *vs.Name
	e.Comment = vs.generateComment()

	generatedNames := make(map[string]int)

	for _, cs := range codesystems {
		whitelistedIds := map[custom.Code]struct{}{}
		if cs.ValueSetConcepts != nil {
			for _, concept := range cs.ValueSetConcepts {
				whitelistedIds[concept.Code] = struct{}{}
			}
		}

		for _, concept := range filterConcepts(cs.Concept, cs.ValueSetFilters) {
			if len(whitelistedIds) > 0 {
				if _, ok := whitelistedIds[concept.Code]; !ok {
					continue
				}
			}

			opts := enumOptions{
				Value: string(concept.Code),
			}

			baseName := ""
			if concept.Display != nil {
				baseName = formatNameAndValue(*concept.Display, string(concept.Code))
			} else {
				baseName = strcase.UpperCamelCase(string(concept.Code))
			}

			name := baseName
			count := generatedNames[name]
			if count > 0 {
				name = fmt.Sprintf("%s%d", baseName, count+1)
			}
			generatedNames[baseName]++

			opts.Name = name

			if concept.Definition != nil {
				opts.Comment = definitionToComment(*concept.Definition)
			}

			e.Options = append(e.Options, opts)
		}
	}

	return e, err
}

type ValueSetCodeSystem struct {
	CodeSystem

	ValueSetConcepts []ValueSetComposeIncludeConcept
	ValueSetFilters  []ValueSetComposeIncludeFilter
}

func (vs ValueSet) getCodeSystems(ctx context.Context) ([]ValueSetCodeSystem, error) {
	codesystems := make([]ValueSetCodeSystem, len(vs.Compose.Include))
	for i, inc := range vs.Compose.Include {
		if inc.System == nil {
			log.Info().Interface("include", inc).Msgf("Skipping %s's #%d include", *vs.Name, i)
			continue
		}

		res, err := getCodeSystem(ctx, inc)
		if err != nil {
			log.Error().Err(err).Msgf("couldn't get codesystem for %s", *vs.Name)
			continue
		}

		codesystems[i] = ValueSetCodeSystem{
			CodeSystem:       res,
			ValueSetConcepts: inc.Concept,
			ValueSetFilters:  inc.Filter,
		}
	}
	return codesystems, nil
}

func getCodeSystem(ctx context.Context, inc ValueSetComposeInclude) (res CodeSystem, err error) {
	systemUrl := string(*inc.System)

	if result, err := cache.Get(systemUrl); err == nil {
		return result, nil
	} else {
		if !errors.Is(err, errCacheMiss) {
			return res, err
		}
	}

	urls, err := getUrl(*inc.System)
	if err != nil {
		return res, err
	}

	var errs []error
	var foundValidURL bool
	for _, url := range urls {
		req, reqErr := http.NewRequestWithContext(ctx, "GET", url, nil)
		if reqErr != nil {
			errs = append(errs, fmt.Errorf("error creating request for url %s: %w", url, reqErr))
			continue
		}

		resp, respErr := httpClient.Do(req)
		if respErr != nil {
			errs = append(errs, fmt.Errorf("error performing request for url %s: %w", url, respErr))
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusAccepted {
				log.Warn().Msgf("url %s returned status code %d, may not be an error but not ideal", url, resp.StatusCode)
			} else {
				errs = append(errs, fmt.Errorf("url %s returned status code %d", url, resp.StatusCode))
				continue
			}
		}

		var currentRes CodeSystem
		decodeErr := json.NewDecoder(resp.Body).Decode(&currentRes)
		if decodeErr != nil {
			errs = append(errs, fmt.Errorf("error decoding json from url %s: %w", url, decodeErr))
			continue
		}

		res = currentRes
		foundValidURL = true
		cacheErr := cache.Set(url, res)
		if cacheErr != nil {
			log.Warn().Err(cacheErr).Msgf("error setting cache for url %s", url)
		}
		break
	}

	if !foundValidURL {
		if len(errs) > 0 {
			return res, fmt.Errorf("no valid CodeSystem found, encountered errors: %w", stderrors.Join(errs...))
		} else {
			return res, errors.New("no valid CodeSystem found and no errors reported, possible issue with getUrl or no URLs returned")
		}
	}

	return res, err
}

func getUrl(uri custom.URI) ([]string, error) {
	uri = custom.URI(strings.Replace(string(uri), "/", "-", 0))

	urlCases := []struct {
		prefixes []string
		replace  []string
	}{{
		prefixes: []string{"http://terminology.hl7.org/CodeSystem/"},
		replace:  []string{"https://terminology.hl7.org/6.1.0/CodeSystem-"},
	}, {
		prefixes: []string{"http://hl7.org/fhir/CodeSystem/", "http://hl7.org/fhir/"},
		replace:  []string{"http://hl7.org/fhir/codesystem-", "https://build.fhir.org/codesystem-"},
	}}

	for _, urlCase := range urlCases {
		for _, prefix := range urlCase.prefixes {
			if result, found := strings.CutPrefix(string(uri), prefix); found {
				urls := make([]string, len(urlCase.replace))
				for i, replace := range urlCase.replace {
					urls[i] = replace + result + ".json"
				}
				return urls, nil
			}
		}
	}

	return nil, errors.Errorf("unknown url: %s", uri)
}

func definitionToComment(def string) *string {
	if strings.HasPrefix(def, "**") {
		strBlocks := strings.Split(def, "**")
		for i, strBlock := range strBlocks {
			switch strings.ToLower(strBlock) {
			case "description", "description:":
				content := strings.TrimSpace(strBlocks[i+1])
				return &content
			}
		}

		return nil
	}

	return &def
}

var (
	nonAlphaNum     = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	validIdentRegex = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)
)

func isAbbreviation(val string) bool {
	if val == "" {
		return false
	}
	for _, char := range val {
		if !isalnum(char) {
			continue
		}
		if !isUpper(char) {
			return false
		}
	}
	return true
}

func isalnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isUpper(r rune) bool {
	return unicode.IsUpper(r)
}

func formatNameAndValue(name, value string) string {
	if isAbbreviation(value) || !validIdentRegex.MatchString(value) {
		parts := nonAlphaNum.Split(name, -1)
		if len(parts) == 0 {
			return strcase.UpperCamelCase(name)
		}
		result := parts[0]
		for _, part := range parts[1:] {
			if part == "" {
				continue
			}
			result += strings.ToUpper(part[:1]) + part[1:]
		}
		return strcase.UpperCamelCase(result)
	}
	return strcase.UpperCamelCase(value)
}
