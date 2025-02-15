package main

func (vs ValueSet) generateEnum() {

}

func (vs ValueSet) getCodeSystems() ([]CodeSystem, error) {
	/*codesystems := make([]CodeSystem, len(vs.Compose.Include))
	for _, inc := range vs.Compose.Include {
		codesystems = append(codesystems, inc.System)
	}
	url := fmt.Sprintf("https://terminology.hl7.org/6.1.0/CodeSystem-%s.json", vs)
	*/
	panic("not implemented")
}
