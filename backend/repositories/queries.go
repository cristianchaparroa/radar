package repositories



const (
	// AroundLocationQuery is a query to search the profile ID's around a specific location.
	// The params required for use successful this query are:
	// - latitude of the center of search
	// - longitude of the center of search
	// - radius of the center of search to find profiles
	// - unit in which should performed the search.
	AroundLocationQuery = `SELECT profile_id,
		latitude,
		longitude,
		distance
		FROM (
		SELECT z.profile_id,
		z.latitude,
		z.longitude,
		p.radius,
		p.distance_unit
		* DEGREES(ACOS(LEAST(1.0, COS(RADIANS(p.latpoint))
		* COS(RADIANS(z.latitude))
		* COS(RADIANS(p.longpoint - z.longitude))
		+ SIN(RADIANS(p.latpoint))
		* SIN(RADIANS(z.latitude))))) AS distance
		FROM locations AS z
		JOIN (
		SELECT 
		%f   AS latpoint,
		%f  AS longpoint,
		%f    AS radius,
		%d AS distance_unit
		) AS p ON 1 = 1
		WHERE z.latitude
		BETWEEN p.latpoint - (p.radius / p.distance_unit)
		AND p.latpoint + (p.radius / p.distance_unit)
		AND z.longitude
		BETWEEN p.longpoint - (p.radius / (p.distance_unit * COS(RADIANS(p.latpoint))))
		AND p.longpoint + (p.radius / (p.distance_unit * COS(RADIANS(p.latpoint))))
		) AS d
		WHERE distance <= radius
		ORDER BY distance`
)