// Package main demonstrates basic usage of the Astrology API Go SDK.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/charts"
	"github.com/astro-api/astroapi-go/categories/data"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

func main() {
	apiKey := os.Getenv("ASTROLOGY_API_KEY")
	if apiKey == "" {
		log.Fatal("ASTROLOGY_API_KEY environment variable is required")
	}

	// Create the client.
	client := astroapi.NewClient(
		option.WithAPIKey(apiKey),
		option.WithMaxRetries(2),
	)

	ctx := context.Background()

	// Example 1: Get current moment data.
	fmt.Println("=== Current Moment ===")
	now, err := client.Data.GetNow(ctx)
	if err != nil {
		log.Printf("GetNow error: %v", err)
	} else {
		printJSON(now)
	}

	// Example 2: Get planetary positions for a birth chart.
	fmt.Println("\n=== Planetary Positions ===")
	subject := shared.Subject{
		Name: "Demo User",
		BirthData: shared.BirthData{
			Year:        1990,
			Month:       5,
			Day:         11,
			Hour:        14,
			Minute:      30,
			City:        "London",
			CountryCode: "GB",
		},
	}

	positions, err := client.Data.GetPositions(ctx, data.PositionsParams{
		Subject: subject,
	})
	if err != nil {
		log.Printf("GetPositions error: %v", err)
	} else {
		printJSON(positions)
	}

	// Example 3: Generate a natal chart.
	fmt.Println("\n=== Natal Chart ===")
	natalChart, err := client.Charts.GetNatal(ctx, charts.NatalChartParams{
		Subject: subject,
	})
	if err != nil {
		log.Printf("GetNatal error: %v", err)
	} else {
		printJSON(natalChart)
	}

	// Example 4: Daily horoscope for a sign.
	fmt.Println("\n=== Aries Daily Horoscope ===")
	_ = client.Horoscope // horoscope sub-client
	// Demonstrating field helpers.
	nameField := astroapi.String("Alice")
	fmt.Printf("Name field value: %v, present: %v\n", nameField.Value, nameField.IsPresent())
}

func printJSON(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("(error marshalling: %v)\n", err)
		return
	}
	fmt.Println(string(b))
}
