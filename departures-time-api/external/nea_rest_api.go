package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
)

type NeaRestApi interface {
	GetNearbyStations(lo vo.Longitude, la vo.Latitude) ([]NeaRestApiDTO, error)
}

type NeaRestApiDTO struct {
	StationName string    `json:"station_name"`
	Location    []float64 `json:"location"`
	Distance    int       `json:"distance"`
}

type neaRestApiBindObject struct {
	Data []NeaRestApiDTO `json:"data"`
}

type neaRestApi struct {
	apiPath string
}

func NewNeaRestApi(apiPath string) NeaRestApi {
	return neaRestApi{apiPath: apiPath}
}

func (h neaRestApi) GetNearbyStations(lo vo.Longitude, la vo.Latitude) ([]NeaRestApiDTO, error) {
	u, _ := url.Parse(h.apiPath)
	q := u.Query()
	q.Add("method", "getStations")
	q.Add("lon", fmt.Sprint(lo.Value()))
	q.Add("lat", fmt.Sprint(la.Value()))
	q.Add("limit", "20")
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest(): %w", err)
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client.Do(): %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("通信失敗")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll(): %w", err)
	}
	var bindObject neaRestApiBindObject
	if err = json.Unmarshal(body, &bindObject); err != nil {
		return nil, fmt.Errorf("Unmarshal(): %w", err)
	}
	return bindObject.Data, nil
}
