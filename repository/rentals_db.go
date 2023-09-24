package repository

import (
	"fmt"
	"time"
	"context"
	"log/slog"
	"database/sql"
	
	_ "github.com/lib/pq"

	
	"simple-rentals-api/entity"
)

var RentalsDB *sql.DB
var err error

// TODO: add some unit tests

func InitRentalsDbConnection() {
		
	slog.Info("Trying to connect to the database...")
	
	// TODO: not to use hardcoded variables. Move them to a env file or use Vault
	host := "postgres"
	port := 5432
	user := "root"
	password := "root"
	dbname := "testingwithrentals"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname) //  sslmode=disable not use it on PROD
	RentalsDB, err = sql.Open("postgres", dsn)

	if err != nil {
		slog.Error("Couldn't open connection to DB", slog.String("Error", err.Error()))
		return
	}

	RentalsDB.SetMaxIdleConns(10)
	RentalsDB.SetMaxOpenConns(100)
	RentalsDB.SetConnMaxLifetime(5 * time.Minute)

	if err := PingDb(); err != nil {
		slog.Error("An error occured trying to connect to DB", slog.String("Error", err.Error()))
		CloseConnection()
		return
	}

	slog.Info("Successfully connected to the database")


	// TODO: Override the default logger to JSON format
}

func CloseConnection() {
	if RentalsDB == nil {
		slog.Error("Db connection is nil")
		return
	}

	if err := RentalsDB.Close(); err != nil {
		slog.Error("Couldn't close DB connection", slog.String("Error", err.Error()))
	}
}

func PingDb() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if RentalsDB == nil {
		return fmt.Errorf("Db connection is nil")
	}

	if err := RentalsDB.PingContext(ctx); err != nil {
		return fmt.Errorf("An error occurred trying to ping DB: %v", err)
	}

	return nil
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

func GenerateSQLQuery(params entity.QueryParams) string {
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

func ExecuteDatabaseQuery(query string) ([]entity.Rental, error) {

	// if err := PingDb(); err != nil {
	// 	CloseConnection()
	// 	// InitRentalsDbConnection()
	// 	return nil, err
	// }

	rows, err := RentalsDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rentals := []entity.Rental{}

	for rows.Next() {
		rental := entity.Rental{}
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
