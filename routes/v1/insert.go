package v1

import (
	"encoding/json"
	"net/http"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"

	"microservice/internal/db"
	apiErrors "microservice/internal/errors"
	types "microservice/types/v1"
)

func Insert(c *gin.Context) {
	var parameters types.Parameter

	err := c.BindJSON(&parameters)
	if err != nil {
		c.Abort()
		res := apiErrors.ErrMissingParameter
		res.Errors = []error{err}
		res.Emit(c)
		return
	}

	err = db.LoadQueries()
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	query, err := db.Queries.Raw("select-tank")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	var tank types.Tank

	err = pgxscan.Get(c, db.Pool(), &tank, query, parameters.MacAddress)
	if err != nil {
		if pgxscan.NotFound(err) {
			query, err = db.Queries.Raw("insert-tank")
			if err != nil {
				c.Abort()
				_ = c.Error(err)
				return
			}

			err = pgxscan.Get(c, db.Pool(), &tank, query, parameters.MacAddress, parameters.Latitude,
				parameters.Longitude, parameters.RoofSize)
			if err != nil {
				c.Abort()
				_ = c.Error(err)
				return
			}
		} else {
			c.Abort()
			_ = c.Error(err)
			return
		}
	}

	if tank.TankLocation.X() != parameters.Longitude || tank.TankLocation.Y() != parameters.Latitude ||
		tank.RoofSize != parameters.RoofSize {
		query, err = db.Queries.Raw("update-tank")
		if err != nil {
			c.Abort()
			_ = c.Error(err)
			return
		}
		err = pgxscan.Get(c, db.Pool(), &tank, query, parameters.MacAddress,
			parameters.Latitude, parameters.Longitude, parameters.RoofSize)
		if err != nil {
			c.Abort()
			_ = c.Error(err)
			return
		}
	}

	query, err = db.Queries.Raw("insert-measurement")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	var measurement types.Measurement

	err = pgxscan.Get(c, db.Pool(), &measurement, query, parameters.MacAddress, parameters.MeasurementTime,
		parameters.WaterLevel, parameters.Draining, parameters.DrainingTime)
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	measurementJSON, err := json.Marshal(measurement)
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, json.RawMessage(measurementJSON))
}
