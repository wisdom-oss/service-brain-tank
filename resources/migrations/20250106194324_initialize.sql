-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS brain_tank;

CREATE TABLE IF NOT EXISTS 
   brain_tank.tank (
       mac_address macaddr PRIMARY KEY,
       tank_location geometry(Point, 4326),
       roof_size double precision
   );

CREATE TABLE IF NOT EXISTS 
   brain_tank.measurement (
       mac_address macaddr references brain_tank.tank(mac_address),
       measurement_time timestamp without time zone,
       water_level double precision,
       draining double precision,
       draining_time timestamp without time zone,
       PRIMARY KEY(mac_address, measurement_time)
   );

SELECT create_hypertable('brain_tank.measurement', by_range('measurement_time'));
-- +goose StatementEnd
-- +goose Down