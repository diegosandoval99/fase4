package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type EnergyData struct {
	Timestamp      string  `json:"timestamp"`
	EnergyConsumed float64 `json:"energy_consumed"`
}

var energyData = []EnergyData{}

func generateEnergyData() float64 {
	return 0.5 + rand.Float64()*(5.0-0.5) // Consumo entre 0.5 y 5 kWh
}

func getEnergyData(w http.ResponseWriter, r *http.Request) {
	// Genera un nuevo dato de consumo
	data := EnergyData{
		Timestamp:      time.Now().Format(time.RFC3339),
		EnergyConsumed: generateEnergyData(),
	}

	// Agregar a la base de datos temporal
	energyData = append(energyData, data)

	// Responder con los datos generados
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getReport(w http.ResponseWriter, r *http.Request) {
	var totalConsumption float64
	for _, data := range energyData {
		totalConsumption += data.EnergyConsumed
	}

	averageConsumption := totalConsumption / float64(len(energyData))

	report := map[string]interface{}{
		"total_consumption":   totalConsumption,
		"average_consumption": averageConsumption,
	}

	// Responder con el reporte
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func getRecommendations(w http.ResponseWriter, r *http.Request) {
	var totalConsumption float64
	for _, data := range energyData {
		totalConsumption += data.EnergyConsumed
	}

	averageConsumption := totalConsumption / float64(len(energyData))

	var recommendation string
	if averageConsumption > 3 {
		recommendation = "Reduce tu consumo apagando los equipos que no uses e implementa electrodomésticos energéticamente eficientes."
	} else {
		recommendation = "¡Buen trabajo! sigue usando las prácticas de eficiencia energética."
	}

	recommendationResponse := map[string]string{
		"recommendation": recommendation,
	}

	// Responder con la recomendación
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendationResponse)
}

func main() {
	// Configurar las rutas
	http.HandleFunc("/api/energy", getEnergyData)
	http.HandleFunc("/api/report", getReport)
	http.HandleFunc("/api/recommendations", getRecommendations)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
