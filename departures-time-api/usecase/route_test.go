package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/haton14/departures-time/departures-time-api/external"
	mock_external "github.com/haton14/departures-time/departures-time-api/external/mock"
	"github.com/haton14/departures-time/departures-time-api/usecase"
	"github.com/stretchr/testify/assert"
)

func TestRouteGetRouting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	exspert := mock_external.NewMockExspert(ctrl)
	mockData := "https://roote.ekispert.net/result?arr=%E5%A4%A7%E6%A3%AE%E6%B5%B7%E5%B2%B8&arr_code=22567&connect=true&dep=%E5%A4%A7%E6%A3%AE(%E6%9D%B1%E4%BA%AC%E9%83%BD)&dep_code=22566&express=true&highway=true&hour&liner=true&local=true&minute&plane=true&shinkansen=true&ship=true&sleep=false&sort=time&surcharge=3&type=dep&via1=&via1_code=&via2=&via2_code="
	t.Run("期待通りデータが取れること", func(t *testing.T) {
		exspert.
			EXPECT().
			GetRoutingURL(external.StationCode("22566"), external.StationCode("22567")).
			Return(mockData, nil)
		r := usecase.NewRoute(exspert)
		actual, err := r.GetRouting("22566", "22567")
		assert.NoError(t, err)
		assert.Equal(t, mockData, actual)
	})

	t.Run("Exspert.GetRoutingURLでエラー", func(t *testing.T) {
		exspert.
			EXPECT().
			GetRoutingURL(external.StationCode("22566"), external.StationCode("22567dummy")).
			Return("", errors.New("other error"))
		r := usecase.NewRoute(exspert)
		actual, err := r.GetRouting("22566", "22567dummy")
		assert.Error(t, err)
		assert.Equal(t, "", actual)
	})
}
