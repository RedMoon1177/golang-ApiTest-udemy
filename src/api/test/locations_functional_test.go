package test

import (
	"encoding/json"
	"fmt"
	"golang-testing/src/api/utils/errors"
	"io"
	"net/http"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryNotFound(t *testing.T) {
	// /////// MOCKING THE RESPONSE /////////////////
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found","error": "not_found","status": 404,"cause": []}`,
	})
	///////////////////////////////////////////////////
	// fmt.Println("about to functional test get country not found")
	response, err := http.Get("http://localhost:8081/locations/countries/AR")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "Country not found", apiErr.Message)

}
