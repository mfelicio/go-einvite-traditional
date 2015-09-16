package auth

type authConfig struct {
	// ClientId is the OAuth client identifier used when communicating with
	// the configured OAuth provider.
	ClientId string

	// ClientSecret is the OAuth client secret used when communicating with
	// the configured OAuth provider.
	ClientSecret string

	// Scope identifies the level of access being requested. Multiple scope
	// values should be provided as a space-delimited string.
	Scope string

	// AuthURL is the URL the user will be directed to in order to grant
	// access.
	AuthURL string

	// TokenURL is the URL used to retrieve OAuth tokens.
	TokenURL string

	// RedirectURL is the URL to which the user will be returned after
	// granting (or denying) access.
	RedirectURL string

	// Optional, "online" (default) or "offline", no refresh token if "online"
	AccessType string

	// ApprovalPrompt indicates whether the user should be
	// re-prompted for consent. If set to "auto" (default) the
	// user will be prompted only if they haven't previously
	// granted consent and the code can only be exchanged for an
	// access token.
	// If set to "force" the user will always be prompted, and the
	// code can be exchanged for a refresh token.
	ApprovalPrompt string

	//This is the URL that Google has defined so that an authenticated application may obtain the user's info in json format
	ProfileInfoURL string
}
