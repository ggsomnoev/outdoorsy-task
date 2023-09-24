package entity

type Rental struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Type            string `json:"type"`
	Make            string `json:"make"`
	Model           string `json:"model"`
	Year            string `json:"year"`
	Length          string `json:"length"`
	Sleeps          string `json:"sleeps"`
	PrimaryImageURL string `json:"primary_image_url"`
	Price           struct {
		Day string `json:"day"`
	} `json:"price"`
	Location struct {
		City    string `json:"city"`
		State   string `json:"state"`
		Zip     string `json:"zip"`
		Country string `json:"country"`
		Lat     string `json:"lat"`
		Lng     string `json:"lng"`
	} `json:"location"`
	User struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"user"`
}

type QueryParams struct {
	PriceMin float64   	`form:"price_min"`
	PriceMax float64   	`form:"price_max"`
	Limit    int       	`form:"limit"`
	Offset   int       	`form:"offset"`
	Sort     string    	`form:"sort"`
	Ids      string    	`form:"ids"`
	Near     string 	`form:"near"`
	NearF    []float64 
}
