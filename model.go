package email

// MailConfig contain all configs for main services
type MailConfig struct {
	URL       string `json:"url,omitempty"`
	Port      string `json:"port,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
}
