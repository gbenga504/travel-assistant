package llmcontext

import (
	airportrepository "github.com/gbenga504/travel-assistant/lib/airport/repository"
)

type LLMContext struct {
	AirportRepository *airportrepository.AirportRepository
}
