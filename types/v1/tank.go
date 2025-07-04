package v1

import (
	"net"

	"github.com/twpayne/go-geom"
)

type Tank struct {
	MacAddress   net.HardwareAddr `json:"mac_address"`
	TankLocation geom.Point       `json:"tank_location"`
	RoofSize     float64          `json:"roof_size"`
}
