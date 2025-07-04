package v1

import (
	"net"
	"time"
)

type Measurement struct {
	MacAddress      net.HardwareAddr `json:"mac_address"`
	MeasurementTime time.Time        `json:"measurement_time"`
	WaterLevel      float64          `json:"water_level"`
	Draining        float64          `json:"draining"`
	DrainingTime    time.Time        `json:"draining_time"`
}
