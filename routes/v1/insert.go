package v1

import (
	"net/http"
	"strings"
	"time"

	"microservice/internal/db"
	apiErrors "microservice/internal/errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
	"github.com/twpayne/go-geom"
)

func Insert(c *gin.Context) {
	var parameters struct {
		MeasurementTime time.Time `json:"messungsZeit" binding:"required"`
		Latitude        float64   `json:"lat" binding:"required"`
		Longitude       float64   `json:"long" binding:"required"`
		RoofSize        float64   `json:"dachflaeche" binding:"required"`
		Waterlevel      float64   `json:"gemessen" binding:"required"`
		Draining        float64   `json:"entwaesserung" binding:"required"`
		DrainingTime    time.Time `json:"entwaesserugsZeit" binding:"required"`
		MacAddress      string    `json:"macAdresse" binding:"required"`
	}

	err := c.BindJSON(&parameters)
	if err != nil {
		c.Abort()
		res := apiErrors.ErrMissingParameter
		res.Errors = []error{err}
		res.Emit(c)
		return
	}

	db.LoadQueries()

	//Bug seems to be here!
	query, err := db.Queries.Raw("select-tank")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	type Tank struct {
		MacAddress   pgtype.Macaddr
		TankLocation geom.Point
		RoofSize     float64
	}

	var selected []Tank

	err = pgxscan.Select(c, db.Pool(), &selected, query)
	if err != nil {
		if !pgxscan.NotFound(err) {
			c.Abort()
			_ = c.Error(err)
			return
		}
	}

	var found int = -1

	var resultTank Tank
	for i := 0; i < len(selected); i++ {
		if strings.EqualFold(selected[i].MacAddress.Addr.String(), parameters.MacAddress) {
			found = i
		}
	}

	print(found)

	if found == -1 {
		query, err = db.Queries.Raw("insert-tank")
		if err != nil {
			c.Abort()
			_ = c.Error(err)
			return
		}

		err = pgxscan.Get(c, db.Pool(), &resultTank, query, parameters.MacAddress, parameters.Latitude, parameters.Longitude, parameters.RoofSize)
		if err != nil {
			c.Abort()
			_ = c.Error(err)
			return
		}
	} else {
		if selected[found].TankLocation.X() != parameters.Longitude || selected[found].TankLocation.Y() != parameters.Latitude || selected[found].RoofSize != parameters.RoofSize {
			query, err = db.Queries.Raw("update-tank")
			if err != nil {
				c.Abort()
				_ = c.Error(err)
				return
			}
			err = pgxscan.Get(c, db.Pool(), &resultTank, query, parameters.MacAddress, parameters.Latitude, parameters.Longitude, parameters.RoofSize)
			if err != nil {
				c.Abort()
				_ = c.Error(err)
				return
			}
		}
	}

	query, err = db.Queries.Raw("insert-measurement")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	var resultMeasurement struct {
		MacAddress      pgtype.Macaddr
		MeasurementTime time.Time
		Waterlevel      float64
		Draining        float64
		DrainingTime    time.Time
	}

	err = pgxscan.Get(c, db.Pool(), &resultMeasurement, query, parameters.MacAddress, parameters.MeasurementTime, parameters.Waterlevel, parameters.Draining, parameters.DrainingTime)
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	c.String(http.StatusCreated, "Measurement of controller "+resultMeasurement.MacAddress.Addr.String()+"at time point "+resultMeasurement.MeasurementTime.String()+" inserted successfully!")
}
