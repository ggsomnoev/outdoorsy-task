package helpers

import (
    "fmt"
    "time"
	"log/slog"
	"database/sql"
	
	_ "github.com/lib/pq"
)

var RentalsDB *sql.DB
var err error

// TODO: add some unit tests

func InitRentalsDbConnection() error {
	// TODO: not to use hardcoded variables. Move them to a env file or use Vault
	host := "127.0.0.1"
	port := 5434
	user := "root"
	password := "root"
	dbname := "testingwithrentals"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname) //  sslmode=disable not use it on PROD
    RentalsDB, err = sql.Open("postgres", dsn)

    if err != nil {
        return err
    }
	
	slog.Info("Successfully connected to the database")

	RentalsDB.SetMaxIdleConns(10)
	RentalsDB.SetMaxOpenConns(100)
	RentalsDB.SetConnMaxLifetime(5 * time.Minute)

	if err := RentalsDB.Ping(); err != nil {
        return err
    }

	// TODO: Override the default logger to JSON format

    return nil
}


// TODO: Add the retry mechanism when the db connection fails
func RetryDBConnection(dsn string) error {	
	const maxRetries = 5
	
	var retryInterval = time.Second	

	for i := 0; i < maxRetries; i++ {				
		
		RentalsDB, err = sql.Open("postgres", dsn)
		
		if err == nil {
			slog.Info("Successfully connected to the database")
			return nil
		}

		slog.Error("Failed to connect to the database", slog.String("error", err.Error()))
		slog.Error("Retrying in...", slog.Duration("interval", retryInterval))
		
		time.Sleep(retryInterval)

		retryInterval *= 2
	}

	return fmt.Errorf("Max retries reached, unable to connect to the database")
}

var BaseQuery = `
	SELECT
		r.id AS id,
		r.name AS name,
		r.description AS description,
		r.type AS type,
		r.vehicle_make AS make,
		r.vehicle_model AS model,
		r.vehicle_year AS year,
		r.vehicle_length AS length,
		r.sleeps AS sleeps,
		r.price_per_day AS "price.day",
		r.home_city AS "location.city",
		r.home_state AS "location.state",
		r.home_zip AS "location.zip",
		r.home_country AS "location.country",
		r.lat AS "location.lat",
		r.lng AS "location.lng",
		r.primary_image_url AS primary_image_url,
		u.id AS "user.id",
		u.first_name AS "user.first_name",
		u.last_name AS "user.last_name"
	FROM
		rentals r
	JOIN
		users u ON r.user_id = u.id
	WHERE
	1=1
`

func GenerateSQLQuery(params QueryParams) string {
	query := BaseQuery

	if params.PriceMin > 0 {
		query += fmt.Sprintf(" AND price_per_day >= %.2f", params.PriceMin)
	}
	if params.PriceMax > 0 {
		query += fmt.Sprintf(" AND price_per_day <= %.2f", params.PriceMax)
	}
	if len(params.Ids) > 0 {
		query += fmt.Sprintf(" AND r.id IN (%s)", params.Ids)
	}
	// TODO: adds some logic for coordinates? Compare the ones from the db that are in range or calculate distance?
	// if len(params.NearF) == 2 {
	// }
	if params.Sort != "" {
		query += fmt.Sprintf(" ORDER BY %s", params.Sort)
	}
	if params.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", params.Limit)
	}
	if params.Offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", params.Offset)
	}


	return query
}

func ExecuteDatabaseQuery(query string) ([]Rental, error) {

	rows, err := RentalsDB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    rentals := []Rental{}
    
    for rows.Next() {
        rental := Rental{}
		if err = rows.Scan(
			&rental.ID,
			&rental.Name,
			&rental.Description,
			&rental.Type,
			&rental.Make,
			&rental.Model,
			&rental.Year,
			&rental.Length,
			&rental.Sleeps,
			&rental.Price.Day,
			&rental.Location.City,
			&rental.Location.State,
			&rental.Location.Zip,
			&rental.Location.Country,
			&rental.Location.Lat,
			&rental.Location.Lng,
			&rental.PrimaryImageURL,
			&rental.User.ID,
			&rental.User.FirstName,
			&rental.User.LastName,
		); err != nil {
            return nil, err
        }

        rentals = append(rentals, rental)
    }

    return rentals, nil
}
