package apollo

// User definition for the JSON unmarshaler
type User struct {
	Name string `json:"name"`
	Sudo bool   `json:"sudo"`
}

// Service definition for the JSON unmarshaler
type Service struct {
	Name     string `json:"name"`
	Critical bool   `json:"critical"`
}
