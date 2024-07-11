package services

import (
	"fmt"
	"golang-testing/src/api/domain/locations"
	"golang-testing/src/api/providers/locations_provider"
	"golang-testing/src/api/utils/errors"
)

type locationsService struct{}

type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	// now, in this services package, we create a PUBLIC variable named "LocationsService" (interface Type)
	LocationsService locationsServiceInterface
)

func init() {
	LocationsService = &locationsService{}
}

func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	fmt.Println("inside service")
	return locations_provider.GetCountry(countryId)
}
