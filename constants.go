package goft

const (
	BaseUrl = "https://api.intra.42.fr"

	Version          = "v2"
	BaseUrlVersioned = BaseUrl + "/" + Version

	AuthUrl  = BaseUrl + "/oauth/authorize"
	TokenUrl = BaseUrl + "/oauth/token"
)
