-- name: select-tank
SELECT * FROM brain_tank.tank;

-- name: update-tank
UPDATE brain_tank.tank
SET tank_location = ST_SetSRID(ST_MakePoint($3, $2), 4326), roof_size = $4 
WHERE mac_address = $1
RETURNING mac_address, tank_location, roof_size;

-- name: insert-tank
INSERT INTO
   brain_tank.tank(mac_address, tank_location, roof_size)
VALUES 
   ($1, ST_SetSRID(ST_MakePoint($3, $2), 4326), $4)
RETURNING mac_address, tank_location, roof_size; 

-- name: insert-measurement
INSERT INTO
   brain_tank.measurement(mac_address, measurement_time, waterlevel, draining, draining_time)
VALUES 
   ($1, $2, $3, $4, $5)
RETURNING mac_address, measurement_time, waterlevel, draining, draining_time; 