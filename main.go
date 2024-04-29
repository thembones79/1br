package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../measurements.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)


    stations := map[string][ ]float64{}
	for scanner.Scan() {
		line := scanner.Text()
        parts := strings.Split(line, ";")
        station := parts[0]
        temp, err := strconv.ParseFloat(parts[1], 64)
        if err != nil {
            log.Fatal(err)
        }


        if _,ok := stations[station]; !ok {
           stations[station] = []float64(temp)
        }else{
stations[station] = append(stations[station],temp)
        }












	}

}
