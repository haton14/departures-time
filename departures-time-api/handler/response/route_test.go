package response_test

import (
	"testing"

	"github.com/haton14/departures-time/departures-time-api/handler/response"
	"github.com/stretchr/testify/assert"
)

func TestNewRouteGet(t *testing.T) {
	url := "https://roote.ekispert.net/result?arr=%E5%A4%A7%E6%A3%AE%E6%B5%B7%E5%B2%B8&arr_code=22567&connect=true&dep=%E5%A4%A7%E6%A3%AE(%E6%9D%B1%E4%BA%AC%E9%83%BD)&dep_code=22566&express=true&highway=true&hour&liner=true&local=true&minute&plane=true&shinkansen=true&ship=true&sleep=false&sort=time&surcharge=3&type=dep&via1=&via1_code=&via2=&via2_code="
	expected := response.RouteGet{
		URL: url,
	}
	actual := response.NewRouteGet(url)
	assert.Equal(t, expected, actual)
}
