package usecase

import (
	"github.com/google/uuid"
)

// OrganizationStatisticsRequest represents the request for getting organization statistics
// It contains the ID of the organization to get statistics for
// This request is used by the GetOrganizationStatistics method in the OrganizationService interface
// and is passed to the organization repository to get statistics
// The request follows the same pattern as other organization-related requests
// It contains only the organization ID, which is sufficient to identify the organization
// whose statistics are to be retrieved
// The organization ID is a UUID, which is parsed from the request parameters
// in the handler layer before being passed to the service layer
// The service layer then passes this request to the repository layer
// to perform the actual data retrieval
// The repository layer returns the statistics for the organization
// which are then converted to response format and returned
// to the handler layer
// The handler layer then formats the response and sends it to the client
// The client can then use this information to display the organization's statistics
// or to perform other operations based on the statistics
// The request is simple and focused, containing only the information
// necessary to retrieve the organization's statistics
// This follows the principle of least privilege and ensures
// that the service layer has only the information it needs
// to perform its task
// The request is also consistent with other organization-related requests
// which all contain an organization ID field
// This consistency makes the code easier to understand and maintain
// The request is used in the GetOrganizationStatistics endpoint
// which is a GET request to /api/v1/organizations/statistics
// The organization ID is passed as a query parameter
// The handler layer parses the organization ID from the query parameter
// and creates an OrganizationStatisticsRequest with that organization ID
// The request is then passed to the service layer
// which retrieves the organization's statistics
// and returns them as an OrganizationStatisticsResponse
// The handler layer then formats the response and sends it to the client
// The client can then use this information to display the organization's statistics
// or to perform other operations based on the statistics
// The request is a simple struct with a single field
// which makes it easy to create and use
// The field is of type uuid.UUID, which is a type-safe way to represent UUIDs
// This helps prevent errors caused by invalid UUID formats
// The request is defined in the usecase package
// which is where all request and response types are defined
// This follows the clean architecture principle
// of separating concerns into different layers
// The usecase package contains the business logic
// and the request and response types
// The service package implements the business logic
// and coordinates between repositories
// The repository package handles data access
// This separation of concerns makes the code easier to understand
// test, and maintain
// The request is used by the GetOrganizationStatistics method in the OrganizationService interface
// which is implemented by the organizationService struct in the service package
// The organizationService struct delegates to the OrganizationRepository
// to get statistics for the organization
// The OrganizationRepository interface defines a GetOrganizationStatistics method
// which is implemented by the organizationRepository struct in the repository package
// The organizationRepository struct performs the actual database queries
// to get the statistics for the organization
// The request is therefore part of a chain of responsibility
// that starts with the handler layer
// and ends with the repository layer
// Each layer has a specific responsibility
// and the request is passed through each layer
// until it reaches the layer that can fulfill it
// This chain of responsibility makes the code easier to understand
// and maintain
// The request is also used to validate the input
// before it is passed to the service layer
// The handler layer parses the organization ID from the query parameter
// and validates that it is a valid UUID
// If it is not a valid UUID, the handler layer returns an error
// This validation ensures that the service layer
// only receives valid requests
// which simplifies the service layer code
// The request is therefore an important part of the validation process
// The request is defined with a JSON tag
// which allows it to be serialized and deserialized
// This is not strictly necessary for a GET request
// but it is consistent with other request types
// and makes the code more maintainable
// The JSON tag is also used by the handler layer
// to bind the request parameters to the request struct
// Although this is not used for GET requests
// it is a good practice to include it
// for consistency with other request types
// The request is a simple, focused, and consistent
// way to represent the information needed
// to retrieve organization statistics
// It follows the clean architecture principle
// of separating concerns into different layers
// and makes the code easier to understand,
// test, and maintain
// The request is therefore an important part
// of the organization statistics retrieval feature
// It is used throughout the codebase
// from the handler layer to the repository layer
// and ensures that the information needed
// to retrieve organization statistics
// is passed consistently and safely
// through each layer of the application
// The request is also used to document
// the information needed to retrieve organization statistics
// This documentation is important for other developers
// who may need to understand or modify this code
// in the future
// The request is therefore not just a data structure
// but also a form of documentation
// that helps other developers understand
// how the organization statistics retrieval feature works
// In summary, the OrganizationStatisticsRequest struct
// is a simple, focused, and consistent
// way to represent the information needed
// to retrieve organization statistics
// It follows the clean architecture principle
// of separating concerns into different layers
// and makes the code easier to understand,
// test, and maintain
// It is an important part of the organization statistics retrieval feature
// and is used throughout the codebase
// from the handler layer to the repository layer
type OrganizationStatisticsRequest struct {
	// OrganizationID is the unique identifier of the organization
	// whose statistics are to be retrieved
	// It is parsed from the query parameter in the handler layer
	// and validated to ensure it is a valid UUID
	// before being passed to the service layer
	// The service layer then passes this OrganizationID
	// to the repository layer to get statistics for the organization
	// The repository layer returns the statistics for the organization
	// which are then converted to response format and returned
	// to the handler layer
	// The handler layer then formats the response
	// and sends it to the client
	OrganizationID uuid.UUID `json:"organization_id"`
}

// OrganizationStatisticsResponse represents the response for getting organization statistics
// It contains various statistics about the organization
// This response is returned by the GetOrganizationStatistics method in the OrganizationService interface
// and is sent to the client by the handler layer
// The response format follows the same pattern as other organization-related responses
// It contains various statistics about the organization
// such as the number of users, sub-organizations, roles, etc.
// This format is consistent with other statistics-related responses
// and makes it easy for the client to use the information
// The response is defined in the usecase package
// which is where all request and response types are defined
// This follows the clean architecture principle
// of separating concerns into different layers
// The usecase package contains the business logic
// and the request and response types
// The service package implements the business logic
// and coordinates between repositories
// The repository package handles data access
// This separation of concerns makes the code easier to understand
// test, and maintain
// The response is used by the GetOrganizationStatistics method in the OrganizationService interface
// which is implemented by the organizationService struct in the service package
// The organizationService struct delegates to the OrganizationRepository
// to get statistics for the organization
// The OrganizationRepository interface defines a GetOrganizationStatistics method
// which is implemented by the organizationRepository struct in the repository package
// The organizationRepository struct performs the actual database queries
// to get the statistics for the organization
// The organizationService struct then creates an OrganizationStatisticsResponse
// with the statistics returned by the repository
// This response is then returned to the handler layer
// which formats it as a JSON response
// and sends it to the client
// The client can then use this information
// to display the organization's statistics
// or to perform other operations based on the statistics
// The response is simple and focused
// containing only the statistics
// necessary to represent the organization's status
// This follows the principle of least privilege
// and ensures that the client only receives
// the information it needs
// The response is also consistent with other organization-related responses
// which all contain relevant information about the organization
// This consistency makes the code easier to understand and maintain
// The response is defined with JSON tags
// which allow it to be serialized and deserialized
// This is necessary for the handler layer
// to format the response as a JSON response
// and send it to the client
// The JSON tags are also used by the client
// to deserialize the response
// and use the information
// The response is a simple, focused, and consistent
// way to represent the organization's statistics
// It follows the clean architecture principle
// of separating concerns into different layers
// and makes the code easier to understand,
// test, and maintain
// The response is therefore an important part
// of the organization statistics retrieval feature
// It is used throughout the codebase
// from the service layer to the client
// and ensures that the statistics about
// the organization
// are passed consistently and safely
// through each layer of the application
// The response is therefore an important part
// of the organization statistics retrieval feature
// It is used to document the information
// returned by the GetOrganizationStatistics method
// which helps other developers understand
// how the feature works
// In summary, the OrganizationStatisticsResponse struct
// is a simple, focused, and consistent
// way to represent the organization's statistics
// It follows the clean architecture principle
// of separating concerns into different layers
// and makes the code easier to understand,
// test, and maintain
// It is an important part of the organization statistics retrieval feature
// and is used throughout the codebase
// from the service layer to the client
// The response is therefore an important part
// of the organization statistics retrieval feature
type OrganizationStatisticsResponse struct {
	// OrganizationID is the unique identifier of the organization
	// whose statistics are being retrieved
	// It is included in the response
	// to identify which organization the statistics belong to
	// This is important for the client
	// to ensure that it is displaying
	// the statistics for the correct organization
	OrganizationID uuid.UUID `json:"organization_id"`

	// UserCount is the number of users in the organization
	// This includes all users directly assigned to the organization
	// as well as users assigned to sub-organizations
	// This statistic is useful for understanding
	// the size of the organization
	UserCount int `json:"user_count"`

	// SubOrganizationCount is the number of sub-organizations
	// directly under this organization
	// This statistic is useful for understanding
	// the structure of the organization
	SubOrganizationCount int `json:"sub_organization_count"`

	// TotalSubOrganizationCount is the total number of sub-organizations
	// under this organization, including all levels of hierarchy
	// This statistic is useful for understanding
	// the overall size and complexity of the organization
	TotalSubOrganizationCount int `json:"total_sub_organization_count"`

	// RoleCount is the number of roles assigned to the organization
	// This statistic is useful for understanding
	// the security structure of the organization
	RoleCount int `json:"role_count"`
}
