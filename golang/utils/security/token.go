package security

// Token -
type Token struct {
	ID         string `json:"id"`
	Authorized bool   `json:"authorized"`
}
