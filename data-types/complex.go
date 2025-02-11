package datatypesold

/*
type Identifier struct {
	//Use      IdentifierUse
	Type     CodeableConcept
	System   URI
	Value    string
	Period   Period
	Assigner Reference[Organisation] // [Reference]([Organisation]) or just text
}

type Extension struct {
	URL URI
}

type Period struct {
	Element

	Start time.Time
	End   time.Time
}

type Coding struct {
	Element

	System       URI
	Version      string
	Code         Code
	Display      string
	UserSelected bool
}

type CodeableConcept struct {
	Element

	Coding Coding
	Text   string
}

var complexDataTypes = []string{"CodeableConcept", "Coding", "Period", "Extension", "Identifier", "Reference", "ReferenceElement", "Organisation", "Endpoint"}

func IsComplexDataType(name string) bool {
	return slices.Contains(complexDataTypes, name)
}*/
