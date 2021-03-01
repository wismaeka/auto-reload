package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/gin-gonic/gin"
)

// type SetStatus struct {
// 	Status string
// 	Water  int `json:"water"`
// 	Wind   int `json:"wind"`
// }
type SetStatus struct {
	Status string
}

type Level struct {
	Water int
	Wind  int
}

func statusLevel(c *gin.Context) {
	var level Level
	if level.Water < 5 && level.Wind < 6 {
		fmt.Println("Aman")
	} else if level.Water > 8 && level.Wind > 15 {
		fmt.Println("Bahaya")
	} else {
		fmt.Println("Siaga")
	}
}

func writeJson() {
	max := 100
	min := 1
	water := rand.Intn(max-min)+min
	wind := rand.Intn(max-min)+min
	jsonString := `level{water,wind}`
	jsonData := []byte(jsonString)
	var status SetStatus
	err := json.Unmarshal(jsonData, &status)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	
	
	r := gin.Default()
	r.GET("/", statusLevel)

	r.Run()
}
