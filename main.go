package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	interval := 15 * time.Second

	go updateJSONFile(interval)

	select {}
}

func updateJSONFile(interval time.Duration)  {
	for {
		water := generateRandomNumber(1, 100)
		wind := generateRandomNumber(1, 100)

		status := Status {
			Water: water,
			Wind: wind,
		}

		statusJSON, err := json.Marshal(status)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			continue
		}

		err = writeToFile("status.json", statusJSON)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}

		fmt.Printf("Updated status: Water - %d, Wind - %d\n", water, wind)

		time.Sleep(interval)
	}
}

func writeToFile(filename string, data []byte) error  {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
