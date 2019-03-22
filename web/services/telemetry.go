package services

import (
	"pubg-fun-stats/parser"
	"pubg-fun-stats/parser/models/telemetry"
	"pubg-fun-stats/repositories"
)

type TelemetryService interface {
	//Fetch(num int64) ([]*telemetry.Telemetry, error)
	RequestMatchTelemetry(url string) (*telemetry.Telemetry, error)
	//Store(m *telemetry.Telemetry) error
}

type telemetryService struct {
	repo repositories.TelemetryRepository
	api  *gopubg.API
}

func NewTelemetryService(repo repositories.TelemetryRepository, api *gopubg.API) TelemetryService {
	return &telemetryService{
		repo: repo,
		api:  api,
	}
}

func (ts *telemetryService) RequestMatchTelemetry(url string) (*telemetry.Telemetry, error) {
	t, err := ts.api.RequestTelemetry(url)
	if err != nil {
		return nil, err
	}
	return t, nil
}
