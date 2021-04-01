package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/shaural/GlobePersonalWebsite/server/pkg/db"
)

const (
	countryFile = "conf/countries.csv"
	stateFile   = "conf/states.csv"
)

func loadCountries(ldb db.Database) error {
	countryCsvFile, err := os.Open(countryFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file: %s, Error:", countryFile, err)
	}
	countryReader := csv.NewReader(countryCsvFile)
	for {
		country, err := countryReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lat, err := strconv.Atoi(country[2])
		if err != nil {
			log.Printf("Latitude cannot be parsed for Country: %v, Error: %v", country, err)
			lat = 0
		}
		lon, err := strconv.Atoi(country[3])
		if err != nil {
			log.Printf("Longitude cannot be parsed for Country: %v, Error: %v", country, err)
			lon = 0
		}
		newCountry := &db.Country{
			ID:        country[0],
			Name:      country[1],
			Latitude:  lat,
			Longitude: lon,
		}
		ldb.InsertCountry(newCountry)
		fmt.Printf("Insert Country: %v", newCountry)
	}
	return loadStates(ldb)
}

func loadStates(ldb db.Database) error {
	stateCsvFile, err := os.Open(stateFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file: %s, Error:", countryFile, err)
	}
	stateReader := csv.NewReader(stateCsvFile)
	for {
		state, err := stateReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		newState := &db.State{
			ID:        state[0],
			CountryID: state[1],
			Name:      state[2],
		}
		ldb.InsertState(newState)
		fmt.Printf("Insert State: %v", newState)
	}
}
