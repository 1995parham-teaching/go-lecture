package request

type Name struct {
	Name  string `json:"name,omitempty"`
	Count *int   `json:"count,omitempty"`
}
