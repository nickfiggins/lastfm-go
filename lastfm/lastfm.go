package lastfm

import "net/http"

var (
	UriApiSecBase  = "https://ws.audioscrobbler.com/2.0/"
	UriApiBase     = "http://ws.audioscrobbler.com/2.0/"
	UriBrowserBase = "https://www.last.fm/api/auth/"
)

type P map[string]interface{}

type Api struct {
	params  *apiParams
	Album   *albumApi
	Artist  *artistApi
	Chart   *chartApi
	Geo     *geoApi
	Library *libraryApi
	Tag     *tagApi
	Track   *trackApi
	User    *userApi
}

type apiParams struct {
	apikey    string
	secret    string
	sk        string
	useragent string
}

type ClientOption func(*Config)

type Config struct {
	client *http.Client
	url    string
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Config) {
		c.client = client
	}
}

func WithURL(url string) ClientOption {
	return func(c *Config) {
		c.url = url
	}
}

var httpClient = http.DefaultClient

func New(key, secret string, opts ...ClientOption) (api *Api) {
	var cfg Config
	for _, o := range opts {
		o(&cfg)
	}

	httpClient = cfg.client
	UriApiSecBase = cfg.url

	params := apiParams{key, secret, "", ""}
	api = &Api{
		params:  &params,
		Album:   &albumApi{&params},
		Artist:  &artistApi{&params},
		Chart:   &chartApi{&params},
		Geo:     &geoApi{&params},
		Library: &libraryApi{&params},
		Tag:     &tagApi{&params},
		Track:   &trackApi{&params},
		User:    &userApi{&params},
	}
	return
}

func (api *Api) SetSession(sessionkey string) {
	api.params.sk = sessionkey
}

func (api Api) GetSessionKey() (sk string) {
	sk = api.params.sk
	return
}

func (api *Api) SetUserAgent(useragent string) {
	api.params.useragent = useragent
}
