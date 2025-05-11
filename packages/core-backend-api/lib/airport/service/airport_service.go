package airportservice

import (
	airportrepository "github.com/gbenga504/travel-assistant/lib/airport/repository"
)

type AirportService struct {
	respository *airportrepository.AirportRepository
}

func NewAirportService(repository *airportrepository.AirportRepository) *AirportService {
	return &AirportService{
		repository,
	}
}

func (s *AirportService) SearchAirports(searchTerm string) (result []airportrepository.AirportSchema) {
	return s.respository.SearchAirports(searchTerm)
}
