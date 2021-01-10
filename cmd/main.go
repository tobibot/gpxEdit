package main

import (
	"fmt"
	"log"
)

func main() {
	printUsage()

	inputFileName := "file.gpx"
	outputFileName := "outFile.gpx"
	latitudeAdjustment := 0.00001
	longitudeAdjustemnt := 0.000005

	gpx, err := readFile(inputFileName)

	if err != nil {
		log.Fatal(err)
	}

	gpxnew, err := adjustGpx(gpx, latitudeAdjustment, longitudeAdjustemnt)

	written, err := writeFile(outputFileName, gpxnew)

	if err != nil {
		log.Fatal(err)
	}

	if written {
		fmt.Printf("New file written to %s\n", outputFileName)
	} else {
		fmt.Println("file wasn't written")
	}

}

func readFile(fileName string) (result interface{}, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func adjustGpx(data interface{}, lon, lat float64) (result interface{}, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func writeFile(fileName string, data interface{}) (result bool, err error) {
	return false, fmt.Errorf("Not implemented")
}

func printUsage() {

}
