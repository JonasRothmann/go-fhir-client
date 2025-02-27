package main

import (
	"regexp"
	"slices"
	"strings"

	"github.com/jonasrothmann/go-fhir-client/custom"
	"github.com/rs/zerolog/log"
)

type HierarchicalCodeSystemConcept struct {
	ParentCode    *custom.Code
	Parent        *HierarchicalCodeSystemConcept
	ChildrenCodes []custom.Code
	Children      []*HierarchicalCodeSystemConcept
	Code          custom.Code
}

func toHierarchical(inputs []CodeSystemConcept) map[custom.Code]*HierarchicalCodeSystemConcept {
	childrenByParent := map[custom.Code][]custom.Code{}
	for _, concept := range inputs {
		if parent, ok := concept.GetProperty("subsumedBy"); ok && parent.ValueCode != nil {
			childrenByParent[*parent.ValueCode] = append(childrenByParent[*parent.ValueCode], concept.Code)
		}
	}
	concepts := make(map[custom.Code]*HierarchicalCodeSystemConcept, len(inputs))

	for _, concept := range inputs {
		h := &HierarchicalCodeSystemConcept{
			Code: concept.Code,
		}
		if parent, ok := concept.GetProperty("subsumedBy"); ok == true && parent.ValueCode != nil {
			h.ParentCode = parent.ValueCode
		}
		if children, ok := childrenByParent[concept.Code]; ok {
			h.ChildrenCodes = children
		}
		concepts[concept.Code] = h
	}

	for _, concept := range concepts {
		if concept.ParentCode != nil {
			if parentConcept, ok := concepts[*concept.ParentCode]; ok {
				concept.Parent = parentConcept
			}
		}
		if len(concept.ChildrenCodes) > 0 {
			concept.Children = make([]*HierarchicalCodeSystemConcept, 0, len(concept.ChildrenCodes))
			for _, childCode := range concept.ChildrenCodes {
				if childConcept, ok := concepts[childCode]; ok {
					concept.Children = append(concept.Children, childConcept)
				}
			}
		}
	}
	return concepts
}

func filterConcepts(inputs []CodeSystemConcept, filters []ValueSetComposeIncludeFilter) []CodeSystemConcept {
	hierarchicalConcepts := toHierarchical(inputs)

	return slices.DeleteFunc(inputs, func(concept CodeSystemConcept) bool {
		hierarchy, ok := hierarchicalConcepts[concept.Code]
		if !ok || hierarchy == nil {
			log.Fatal().Interface("concept", concept).Msg("unexpected missing hierarchy")
		}

		for _, filter := range filters {
			if filter.Property != "concept" {
				log.Warn().Interface("concept", concept).Interface("filter", filter).Msgf("Unhandled filter property: %s", filter.Property)
				continue
			}

			isMatch := filter.Value == string(concept.Code)

			switch filter.Op {
			case "=":
				return !isMatch
			case "is-a":
				curr := hierarchy
				for curr != nil {
					//log.Debug().Interface("filter", filter).Interface("curr", curr).Send()
					if string(curr.Code) == filter.Value {
						return false
					}
					curr = curr.Parent
				}
				return true
			case "descendent-of":
				curr := hierarchy.Parent
				for curr != nil {
					if string(curr.Code) == filter.Value {
						return false
					}
					curr = curr.Parent
				}
				return true
			case "is-not-a":
				curr := hierarchy
				isA := false
				for curr != nil {
					if string(curr.Code) == filter.Value {
						isA = false
						break
					}
					curr = curr.Parent
				}
				return !isA
			case "regex":
				matched, err := regexp.MatchString(filter.Value, string(concept.Code))
				if err != nil {
					log.Fatal().Msgf("Invalid regex: %s", filter.Value)
				}
				if !matched {
					log.Printf("Skipped due to 'regex' filter: %s does not match %s", concept.Code, filter.Value)
					return true
				}
			case "in":
				values := strings.Split(filter.Value, ",")
				found := false
				for _, val := range values {
					if val == string(concept.Code) {
						found = true
						break
					}
				}
				return !found
			case "not-in":
				values := strings.Split(filter.Value, ",")
				for _, val := range values {
					if val == string(concept.Code) {
						log.Printf("Skipped due to 'not-in' filter: %s is in %s", concept.Code, filter.Value)
						return true
					}
				}
			case "generalizes":
				queue := []*HierarchicalCodeSystemConcept{hierarchy}
				for len(queue) > 0 {
					curr := queue[0]
					queue = queue[1:]

					if string(curr.Code) == filter.Value {
						return false
					}
					queue = append(queue, curr.Children...)
				}
				return true
			case "child-of":
				if hierarchy.ParentCode != nil && string(*hierarchy.ParentCode) == filter.Value {
					return false
				}
				return true
			case "descendent-leaf":
				isDescendant := false
				curr := hierarchy.Parent
				for curr != nil {
					if string(curr.Code) == filter.Value {
						isDescendant = true
						break
					}
					curr = curr.Parent
				}

				if isDescendant && len(hierarchy.Children) == 0 {
					return false
				}
				return true
			case "exists":
				panic("exists not implemented")
			default:
				log.Fatal().Msgf("Unhandled filter op: %s", filter.Op)
			}
		}

		return false
	})
}
