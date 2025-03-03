package goft

// VERSION of goft, follows Semantic Versioning. (http://semver.org/)
const VERSION = "0.1.0"

const (
	BaseUrl = "https://api.intra.42.fr"

	Version          = "v2"
	BaseUrlVersioned = BaseUrl + "/" + Version

	AuthUrl  = BaseUrl + "/oauth/authorize"
	TokenUrl = BaseUrl + "/oauth/token"
)
