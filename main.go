package main

import (
	// "bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
	"log"
	"time"
)

func main() {

	defer func(start time.Time) {
		fmt.Printf("took us %v", time.Since(start))
	}(time.Now())

	f, err := os.Open("../measurements.txt")
	if err != nil {
		log.Fatal(err)
	}

	chunkSize := 500 * 1024 * 1024
	buf := make([]byte, chunkSize)
	for {
		n, err := f.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

	}

	// scanner := bufio.NewScanner(f)

	// stations := map[string][]float64{}
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	parts := strings.Split(line, ";")
	// 	station := parts[0]
	// 	temp, err := strconv.ParseFloat(parts[1], 64)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	//        _ = station
	//        _ = temp

	// if _, ok := stations[station]; !ok {
	// 	stations[station] = []float64{temp}
	// } else {
	// 	stations[station] = append(stations[station], temp)
	// }

	// }

}
