package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather(city string, apiKey string) (string, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=fr", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// try to read body for a helpful message
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error: %s (status %d)", strings.TrimSpace(string(b)), resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	var data WeatherData
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if len(data.Weather) == 0 {
		return "Ville introuvable", nil
	}

	return fmt.Sprintf("🌡 Température : %.1f°C\n💧 Humidité : %d%%\n🌥 Description : %s",
		data.Main.Temp, data.Main.Humidity, data.Weather[0].Description), nil
}

func main() {
	// Prefer API key from environment variable OPENWEATHER_API_KEY
	// If not set, fall back to the provided API key.
	apiKey := strings.TrimSpace(os.Getenv("OPENWEATHER_API_KEY"))
	if apiKey == "" {
		apiKey = "7f3387f63a6f529cd981bf83b929de31"
		fmt.Println("OPENWEATHER_API_KEY non défini — utilisation de la clé intégrée.")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Entrez une ville (ou 'q' pour quitter) : ")
		city, _ := reader.ReadString('\n')
		city = strings.TrimSpace(city)
		if city == "" {
			fmt.Println("Veuillez entrer une ville valide")
			continue
		}
		if city == "q" || city == "quit" || city == "exit" {
			fmt.Println("Au revoir !")
			return
		}

		weather, err := getWeather(city, apiKey)
		if err != nil {
			fmt.Printf("Erreur: %v\n", err)
			continue
		}
		fmt.Println(weather)
	}
}
