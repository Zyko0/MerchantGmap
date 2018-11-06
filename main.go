package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("The company name is required as an argument")
		os.Exit(1)
	}
	companyName := os.Args[1]
	b, err := ioutil.ReadFile("api.key") // regular file containing your raw googlemaps api key
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	apiKey := string(b)

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.FindPlaceFromTextRequest{
		Input:     companyName,
		InputType: maps.FindPlaceFromTextInputTypeTextQuery,
	}
	resp, err := c.FindPlaceFromText(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	for _, result := range resp.Candidates {
		fmt.Println("-----------------------")
		fmt.Println("PlaceID: ", result.PlaceID)
		place, err := c.PlaceDetails(context.Background(), &maps.PlaceDetailsRequest{
			PlaceID: result.PlaceID,
		})
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}
		fmt.Println("Name: ", place.Name)
		fmt.Println("Icon: ", place.Icon)
		fmt.Println("Website: ", place.Website)
		fmt.Println("Location(lat,lng): ", place.Geometry.Location.Lat, place.Geometry.Location.Lng)
		fmt.Println("-----------------------")
	}
}
