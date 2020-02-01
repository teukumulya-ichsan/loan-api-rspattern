package response

// Error should used to return business error message
type Error struct {
	Message string `json: "string"`
}
