package fhirclient

type Resource interface {
	ResourceType() string
}
