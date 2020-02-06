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
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
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
		var table = rand.Intn(30)
		//table + 1
		table++

		var name = line[0] + " " + line[1]
		person = append(person, name)

		if table == 29 {
			//if the group is full, next table
			if tableFill[table] < 31 {
				fmt.Println(name, "KC")
				tableFill[table]++
			} else {
				//else, add the integer assigned to the table to the usedTables slice.
				fmt.Println(name, "KC")
				usedTables = append(usedTables, table)
			}
		} else if table == 30 {
			//if the group is full, next table
			if tableFill[table] < 30 {
				fmt.Println(name, "Waiter")
				tableFill[table]++
			} else {
				//else, add the integer assigned to the table to the usedTables slice.
				fmt.Println(name, "Waiter")
				usedTables = append(usedTables, table)
			}
		} else {
			if tableFill[table] < 8 {
				fmt.Println(name, table)
				tableFill[table]++
			} else {
				//else, add the integer assigned to the table to the usedTables slice.
				// tableFill[table]++
				fmt.Println(name, table)
				usedTables = append(usedTables, table)
			}
		}
		// fmt.Println(tableFill)
		// fmt.Println(usedTables)
	}
}
