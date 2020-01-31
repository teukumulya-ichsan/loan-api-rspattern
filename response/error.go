package response

// ServiceError should used to return business error message
type ResponseError struct {
	Message string `json: "string"`
}
