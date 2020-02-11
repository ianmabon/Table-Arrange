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
		Table2    []int
	}
	//Elliot's code modified
	//empty tables
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
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
		var table = rand.Intn(33)
		var table2 = rand.Intn(33)
		//table + 1
		table++
		table2++
		person[studentNum].Table = append(person[studentNum].Table, table)
		person[studentNum].Table2 = append(person[studentNum].Table2, table2)
		//never runs just used to declare student
		if table == 34 {
			fmt.Println(student)
		}
		if table == 32 {
			//if the group is full, next table
			if tableFill[table] < 7 {
				//if Table 1 = Table 2, then reshuffle
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)
					fmt.Println(person[studentNum], "KC")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum], "KC")
					tableFill[table]++
				}
			} else {
				//if Table 1 = Table 2, then reshuffle
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)
					fmt.Println(person[studentNum], "KC")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum], "KC")
					usedTables = append(usedTables, table)
				}
			}
		} else if table == 33 {
			//if the group is full, next table
			if tableFill[table] < 30 {
				//if Table 1 = Table 2, then reshuffle
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)
					fmt.Println(person[studentNum], "Waiter")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum], "Waiter")
					tableFill[table]++
				}
			} else {
				//else, add the integer assigned to the table to the usedTables slice.
				//if Table 1 = Table 2, then reshuffle
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)
					fmt.Println(person[studentNum], "Waiter")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum], "Waiter")
					usedTables = append(usedTables, table)
				}
			}
		} else {
			if tableFill[table] < 8 {
				//if the table is full, next table
				//if Table 1 = Table 2, then reshuffle
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)
					fmt.Println(person[studentNum])
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum])
					tableFill[table]++
				}
			} else {
				//else, add the integer assigned to the table to the usedTables slice.
				// tableFill[table]++
				//if Table 1 = Table 2, then reshuffle
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)
					fmt.Println(person[studentNum])
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum])
					usedTables = append(usedTables, table)
				}
			}
		}
	}
	// fmt.Println(person)
	// fmt.Println(usedTables)
}
