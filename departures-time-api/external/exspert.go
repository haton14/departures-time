package external

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
)

type Exspert interface {
	GetByName(name vo.StationName) ([]ExspertDTO, error)
}

type ExspertDTO struct {
	Station  ExspertStation  `json:"Station"`
	GeoPoint ExspertGeoPoint `json:"GeoPoint"`
}

type ExspertStation struct {
	Code        StationCode `json:"code"`
	StationName string      `json:"Name"`
}

type ExspertGeoPoint struct {
	Longitude string `json:"longi_d"`
	Latitude  string `json:"lati_d"`
}

type exspertBindObject struct {
	ResultSet struct {
		Point []ExspertDTO `json:"Point"`
	} `json:"ResultSet"`
}

type StationCode string

type exspert struct {
	apiPath string
	apiKey  string
}

func NewExspert(apiPath, apiKey string) Exspert {
	return exspert{
		apiPath: apiPath,
		apiKey:  apiKey,
	}
}

func (e exspert) GetByName(name vo.StationName) ([]ExspertDTO, error) {
	u, err := url.Parse(e.apiPath)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Add("key", e.apiKey)
	q.Add("name", name.Value())
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("通信失敗")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var bindObject exspertBindObject
	if err = json.Unmarshal(body, &bindObject); err != nil {
		return nil, err
	}
	return bindObject.ResultSet.Point, nil
}
