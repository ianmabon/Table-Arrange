package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type Person struct {
	Firstname string
	Lastname  string
	Table     string
}

func main() {
	//Elliot's code modified
	//empty tables
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	usedTables := []int{}

	var person []string

	csvFile, _ := os.Open("seating.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		//random seed
		rand.Seed(time.Now().UnixNano())
		//random table number
		var table = rand.Intn(33)

		var name = line[0] + " " + line[1]
		person = append(person, name)

		//if the table is full, next table
		if tableFill[table] < 8 {
			fmt.Println(name, table)
			tableFill[table]++
		} else {
			//else, add the integer assigned to the table to the usedTables slice.
			tableFill[table]++
			fmt.Println(line[0], line[1], table)
			usedTables = append(usedTables, table)
		}
	}
}

//two print statements?????
//random tables, but no kitchen crew or waiters
