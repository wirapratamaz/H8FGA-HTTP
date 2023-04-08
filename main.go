package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Payload struct {
	Water       int    `json:"water"`
	WaterStatus string `json:"water_status"`
	Wind        int    `json:"wind"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	rand.Seed(time.Now().Unix())

	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		var waterStatus string
		if water < 5 {
			waterStatus = "aman"
		} else if water <= 8 {
			waterStatus = "siaga"
		} else {
			waterStatus = "bahaya"
		}

		var windStatus string
		if wind < 6 {
			windStatus = "aman"
		} else if wind <= 15 {
			windStatus = "siaga"
		} else {
			windStatus = "bahaya"
		}

		//create payload
		payload := Payload{
			Water:       water,
			WaterStatus: waterStatus,
			Wind:        wind,
			WindStatus:  windStatus,
		}

		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error ecoding payload", err)
			continue
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
		if err != nil {
			fmt.Println("Error sending POST request", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("{\n\"water\": %d,\n\"wind\": %d\n}\n", water, wind)
		fmt.Printf("status water: %s\n", waterStatus)
		fmt.Printf("status wind %s\n", windStatus)

		time.Sleep(15 * time.Second)
	}
}
