package v1

import (
	"time"
)

type Parameter struct {
	MeasurementTime time.Time `binding:"required"        json:"messungsZeit"`
	Latitude        float64   `binding:"required"        json:"lat"`
	Longitude       float64   `binding:"required"        json:"long"`
	RoofSize        float64   `binding:"required"        json:"dachflaeche"`
	WaterLevel      float64   `binding:"required"        json:"gemessen"`
	RainForecast    float64   `binding:"required"        json:"regenvorhersage"`
	Draining        bool      `json:"entwaesserung"`
	DrainingTime    time.Time `json:"entwaesserungsZeit"`
	MacAddress      string    `binding:"required"        json:"macAdresse"`
}
