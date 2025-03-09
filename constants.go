package goft

const (
	BaseURL = "https://api.intra.42.fr"

	Version          = "v2"
	BaseUrlVersioned = BaseURL + "/" + Version

	AuthURL  = BaseURL + "/oauth/authorize"
	TokenURL = BaseURL + "/oauth/token"
)
