package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

var statRes status

type status struct {
	Status level `json:"Status"`
}

type level struct {
	Water     string `json:"Water"`
	Wind      string `json:"Wind"`
	Indicator string `json:"Indicator"`
}

func statusLevel(water int, wind int) (indicator string) {

	if water > 8 || wind > 15 {
		return "Bahaya"
	} else if (water >= 6 && water <= 8) || (wind >= 7 && wind <= 15) {
		return "Siaga"
	} else {
		return "Aman"
	}

}

func generateVal() (water, wind int) {

	rand.Seed(time.Now().UTC().UnixNano())
	water = rand.Intn(100)
	wind = rand.Intn(100)

	return water, wind

}

func writeJson(wt, wd int) (dstatus status) {
	water := strconv.Itoa(wt)
	wind := strconv.Itoa(wd)
	dataStatus := status{level{water + " meter", wind + " m/s", statusLevel(wt, wd)}}

	b, err := json.Marshal(dataStatus)

	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := os.OpenFile("test.json",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(b) + "\n"); err != nil {
		log.Println(err)
	}

	return dataStatus
}

func structToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

func main() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		water, wind := generateVal()
		statRes = writeJson(water, wind)
		mapdt, err1 := structToMap(statRes)
		if err1 != nil {
			panic(err1.Error())
		}
		err = tmpl.ExecuteTemplate(w, "index", mapdt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
