package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
)

type Exspert interface {
	GetByName(name vo.StationName) ([]ExspertDTO, error)
	GetRoutingURL(from, to StationCode) (string, error)
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

type exspertBindObjects struct {
	ResultSet struct {
		Point []ExspertDTO `json:"Point"`
	} `json:"ResultSet"`
}

type exspertBindObject struct {
	ResultSet struct {
		Point ExspertDTO `json:"Point"`
	} `json:"ResultSet"`
}

type StationCode string

func (s StationCode) Value() string {
	return string(s)
}

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
	u.Path = path.Join(u.Path, "station")
	q := u.Query()
	q.Add("key", e.apiKey)
	q.Add("name", name.Value())
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
	bindObjects := &exspertBindObjects{}
	if err = json.Unmarshal(body, bindObjects); err != nil {
		bindObject := &exspertBindObject{}
		if err = json.Unmarshal(body, bindObject); err != nil {
			return nil, fmt.Errorf("Unmarshal(): %w", err)
		}
		return []ExspertDTO{bindObject.ResultSet.Point}, nil
	}
	if bindObjects.ResultSet.Point == nil {
		return nil, vo.ErrNotFound
	}
	return bindObjects.ResultSet.Point, nil
}

type routingURLBindObject struct {
	ResultSet struct {
		ResourceURI string `json:"ResourceURI"`
	} `json:"ResultSet"`
}

func (e exspert) GetRoutingURL(from, to StationCode) (string, error) {
	u, err := url.Parse(e.apiPath)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, "search")
	u.Path = path.Join(u.Path, "course")
	u.Path = path.Join(u.Path, "light")
	q := u.Query()
	q.Add("key", e.apiKey)
	q.Add("from", from.Value())
	q.Add("to", to.Value())
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return "", fmt.Errorf("NewRequest(): %w", err)
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Client.Do(): %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("通信失敗")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ReadAll(): %w", err)
	}
	bindObject := &routingURLBindObject{}
	if err = json.Unmarshal(body, bindObject); err != nil {
		return "", fmt.Errorf("Unmarshal(): %w", err)

	}
	return bindObject.ResultSet.ResourceURI, nil
}
