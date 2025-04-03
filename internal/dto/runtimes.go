package dto


type RuntimeResponse struct {
	Language string `json:"language"`
	Version string `json:"version"`
	Aliases []string `json:"aliases"`
}