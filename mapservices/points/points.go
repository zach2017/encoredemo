// Service url takes URLs, generates random short IDs, and stores the URLs in a database.
package points

import (
	"context"
)

type Location struct {
	Name string  `json:"name"`
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
}

type Locations struct {
	Sports []Location `json:"sports"`
	Music  []Location `json:"music"`
}

// Shorten shortens a URL.
//encore:api public method=GET path=/getPoints
func getPoints(ctx context.Context) (*Locations, error) {

	data := Locations{
        Sports: []Location{
            {
                Name: "Riverwalk Stadium",
                Lon:  -86.3078,
                Lat:  32.3802,
            },
            {
                Name: "Cramton Bowl",
                Lon:  -86.2988,
                Lat:  32.3835,
            },
        },
        Music: []Location{
            {
                Name: "Capitol Sounds",
                Lon:  -86.3021,
                Lat:  32.3773,
            },
            {
                Name: "Montgomery Performing Arts Centre",
                Lon:  -86.3080,
                Lat:  32.3748,
            },
        },

	}	
	return &data, nil
}
