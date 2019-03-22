package repositories

import "database/sql"

type TelemetryRepository interface {
	//Fetch(lim int64) ([]*telemetry.Telemetry, error)
	//Store(match *telemetry.Telemetry) error
}

type telemetrySQLRepository struct {
	Conn *sql.DB
}

func NewTelemetrySQLRepository(conn *sql.DB) TelemetryRepository {
	return &telemetrySQLRepository{conn}
}
