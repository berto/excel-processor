package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func seedShips(csvPath string) error {
	reader, err := readCSV(csvPath)
	if err != nil {
		return err
	}
	ships, quit := processCSV(reader)
	for {
		select {
		case ship := <-ships:
			insertShip(ship)
		case <-quit:
			return nil
		}
	}
}

func processCSV(r *csv.Reader) (chan Ship, chan bool) {
	ships := make(chan Ship, 10)
	quit := make(chan bool, 1)
	if _, err := r.Read(); err != nil {
		quit <- true
		return ships, quit
	}

	go func() {
		defer close(ships)
		defer close(quit)
		for {
			record, err := r.Read()
			if err == io.EOF {
				quit <- true
				return
			}

			name := record[0]
			size, _ := strconv.Atoi(record[1])

			ships <- Ship{
				Name: name,
				Size: size,
			}
		}
	}()
	return ships, quit
}

func readCSV(path string) (*csv.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return csv.NewReader(bufio.NewReader(file)), nil
}
