package errors

// ServiceError should used to return business error message
type ServiceError struct {
	Message string `json: "string"`
}
