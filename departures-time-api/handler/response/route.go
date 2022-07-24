package response

type RouteGet struct {
	URL string `json:"url"`
}

func NewRouteGet(url string) RouteGet {
	return RouteGet{
		URL: url,
	}
}
