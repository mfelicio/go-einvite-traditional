package web

type HttpConfig struct {
	Port       string
	SSL        string
	StaticPath string
}

type SessionConfig struct {
	Cookie      string
	HttpOnly    bool
	Secure      bool
	Ttl         int
	AutoRefresh bool
}

type SecurityKeys struct {
	SignKey       string
	EncryptionKey string

	RawSignKey       []byte
	RawEncryptionKey []byte
}

type WebConfig struct {
	Http     *HttpConfig
	Session  *SessionConfig
	Security *SecurityKeys
}
