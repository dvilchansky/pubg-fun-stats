package services

import (
	"pubg-fun-stats/parser"
	"pubg-fun-stats/parser/models/telemetry"
)

type TelemetryService interface {
	RequestMatchTelemetry(url string) (*telemetry.Telemetry, error)
}

type telemetryService struct {
	api *gopubg.API
}

func NewTelemetryService(api *gopubg.API) TelemetryService {
	return &telemetryService{
		api: api,
	}
}

func (ts *telemetryService) RequestMatchTelemetry(url string) (*telemetry.Telemetry, error) {
	t, err := ts.api.RequestTelemetry(url)
	if err != nil {
		return nil, err
	}
	return t, nil
}
