package valuesets

type ValueSet interface {
	IsValid() bool
}

// required = iota
// other = string
