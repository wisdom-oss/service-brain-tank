package v1

import (
	"time"
)

type Parameter struct {
	MeasurementTime time.Time `binding:"required" json:"messungsZeit"`
	Latitude        float64   `binding:"required" json:"lat"`
	Longitude       float64   `binding:"required" json:"long"`
	RoofSize        float64   `binding:"required" json:"dachflaeche"`
	WaterLevel      float64   `binding:"required" json:"gemessen"`
	Draining        float64   `binding:"required" json:"entwaesserung"`
	DrainingTime    time.Time `binding:"required" json:"entwaesserugsZeit"`
	MacAddress      string    `binding:"required" json:"macAdresse"`
}
