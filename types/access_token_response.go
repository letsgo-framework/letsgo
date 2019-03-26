package types

type TokenResponse struct {
	Access_token string `json:"access_token"`
	Expires_in   string `json:"expires_in"`
	Scope        string `json:"scope"`
	Token_type   string `json:"token_type"`
}
