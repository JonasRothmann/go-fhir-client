// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/endpoint-environment
*/type EndpointEnvironment string

const (
	// Production environment and is expected to contain real data and should be protected appropriately
	EndpointEnvironmentProd EndpointEnvironment = "prod"
	// Staging environment typically used while preparing for a release to production
	EndpointEnvironmentStaging EndpointEnvironment = "staging"
	// Development environment used while building systems
	EndpointEnvironmentDev EndpointEnvironment = "dev"
	// Test environment, it is not intended for production usage.
	EndpointEnvironmentTest EndpointEnvironment = "test"
	// Training environment, it is not intended for production usage and typically contains data specifically prepared for training usage.
	EndpointEnvironmentTrain EndpointEnvironment = "train"
)
