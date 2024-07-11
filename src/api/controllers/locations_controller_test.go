package controllers

import (
	"encoding/json"
	"golang-testing/src/api/domain/locations"
	"golang-testing/src/api/services"
	"golang-testing/src/api/utils/errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

type locationsServiceMock struct{}

func (*locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

func TestGetCountryNotFound(t *testing.T) {

	// mock locationsService methods
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: "Country not found",
		}
	}

	services.LocationsService = &locationsServiceMock{}
	// create a test context to manage the test request and response
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	// call GetCountry() from the controllers
	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiErr errors.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}
