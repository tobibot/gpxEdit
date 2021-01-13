package main

import (
	"encoding/xml"
	"fmt"
	"github.com/tobibot/gpxEdit/cmd/gpxStruct"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	printUsage()

	// ToDo: package Flags for input
	inputFileName := "./resources/testInput.gpx"
	outputFileName := "./resources/testOutput.gpx"
	latitudeAdjustment := 0.0041   // east-west
	longitudeAdjustment := -0.0001 // north-south

	gpx, err := readFile(inputFileName)
	checkError(err)

	gpxNew, err := adjustGpx(&gpx, latitudeAdjustment, longitudeAdjustment)
	checkError(err)

	written, err := writeFile(outputFileName, gpxNew)
	checkError(err)

	if written {
		fmt.Printf("New file written to %s\n", outputFileName)
	} else {
		fmt.Println("file wasn't written")
	}

}

func readFile(fileName string) (result gpxStruct.GpxStruct, err error) {

	xmlFile, err := os.Open(fileName)

	if err != nil {
		return gpxStruct.GpxStruct{}, err
	}

	fmt.Printf("Successfully Opened %s\n", fileName)
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		return gpxStruct.GpxStruct{}, err
	}

	var data gpxStruct.GpxStruct
	err = xml.Unmarshal(byteValue, &data)

	if err != nil {
		return gpxStruct.GpxStruct{}, err
	}

	return data, nil
}

func adjustGpx(data *gpxStruct.GpxStruct, lonAdjustment, latAdjustment float64) (dataOut *gpxStruct.GpxStruct, err error) {

	for i, trkpt := range data.Trk.Trkseg.Trkpt {
		originalLat, err := strconv.ParseFloat(trkpt.Lat, 64)
		trkpt.Lat = strconv.FormatFloat((originalLat + latAdjustment), 'f', 8, 64)

		if err != nil {
			return &gpxStruct.GpxStruct{}, err
		}

		originalLon, err := strconv.ParseFloat(trkpt.Lon, 64)
		trkpt.Lon = strconv.FormatFloat((originalLon + lonAdjustment), 'f', 8, 64)

		if err != nil {
			return &gpxStruct.GpxStruct{}, err
		}

		data.Trk.Trkseg.Trkpt[i] = trkpt
	}

	return data, nil
}

func writeFile(fileName string, data *gpxStruct.GpxStruct) (result bool, err error) {
	file, err := xml.MarshalIndent(data, "", " ")
	if err != nil {
		return false, err
	}

	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}

func printUsage() {

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
