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

func main() {
	type Person struct {
		Firstname string
		Lastname  string
		Table     []int
	}
	//Elliot's code modified
	//empty tables
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//used table array
	usedTables := []int{}
	//defined person
	var person []Person

	csvFile, _ := os.Open("seating.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		person = append(person, Person{
			Firstname: line[0],
			Lastname:  line[1],
		})
	}

	//random seed
	rand.Seed(time.Now().UnixNano())

	for studentNum, student := range person {
		//random table number
		var table = rand.Intn(30)
		// table = append(table, Person.Table)
		//table + 1
		table++
		person[studentNum].Table = append(person[studentNum].Table, table)

		if table == 29 {
			//if the group is full, next table
			if tableFill[table] < 7 {
				fmt.Println(student, "KC")
				tableFill[table]++
			} else {
				//else, add the integer assigned to the table to the usedTables slice.
				fmt.Println(student, "KC")
				usedTables = append(usedTables, table)
			}
		} else if table == 30 {
			//if the group is full, next table
			if tableFill[table] < 30 {
				fmt.Println(student, "Waiter")
				tableFill[table]++
			} else {
				//else, add the integer assigned to the table to the usedTables slice.
				fmt.Println(student, "Waiter")
				usedTables = append(usedTables, table)
			}
		} else {
			if tableFill[table] < 8 {
				//if the table is full, next table
				fmt.Println(student, person[studentNum].Table)
				tableFill[table]++
			} else {
				//else, add the integer assigned to the table to the usedTables slice.
				// tableFill[table]++
				fmt.Println(student, person[studentNum].Table)
				usedTables = append(usedTables, table)
			}
		}
	}

	// fmt.Println(person)
	// fmt.Println(usedTables)
}
