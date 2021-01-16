package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/tobibot/gpxEdit/cmd/gpxStruct"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
)

var (
	inFile              = flag.String("in", "", "Input file name")
	outFile             = flag.String("out", "", "Output file name")
	latitudeAdjustment  = flag.Float64("lat", 0.0, "latitude adjustment (optional, defaults to 0.0)")
	longitudeAdjustment = flag.Float64("lon", 0.0, "longitude adjustment (optional, defaults to 0.0)")
)

func main() {
	flag.Parse()

	if *inFile == "" {
		printUsage()
		os.Exit(-1)
	}

	if *outFile == "" {
		printUsage()
		os.Exit(-1)
	}

	fmt.Println("Params:")
	fmt.Printf("inFile: %s\n", *inFile)
	fmt.Printf("outFile: %s\n", *outFile)
	fmt.Printf("latitude adjustment: %f\n", *latitudeAdjustment)
	fmt.Printf("longitude adjustment: %f\n\n", *longitudeAdjustment)

	gpx, err := readFile(*inFile)
	checkError(err)

	gpxNew, err := adjustGpx(&gpx, *latitudeAdjustment, *longitudeAdjustment)
	checkError(err)

	written, err := writeFile(*outFile, gpxNew)
	checkError(err)

	if written {
		fmt.Printf("New file written to %s\n", *outFile)
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

	var execName string
	switch {
	case runtime.GOOS == "darwin":
		execName = "./gpxEdit"
	case runtime.GOOS == "linux":
		execName = "./gpxEdit"
	default:
		execName = "gpxEdit.exe"
	}

	msg := "This tool transforms your track by the input you give north or south and east or west.\n"
	msg += "Params:\n"
	msg += "-in PathToYourInputFile 	(gpx-Input)\n"
	msg += "-out PathToYourOutputFile 	(desired place to write the result)\n"
	msg += "-lat PathToYourOutputFile 	(longitude adjustment for each track point in your gpx-File. Default: 0.0. Decimal value.)\n"
	msg += "-lon PathToYourOutputFile 	(latitude adjustment for each track point in your gpx-File. Default: 0.0. Decimal value.)\n"
	msg += fmt.Sprintf("Sample: %s -in myLastRun.gpx -out myLastRunAdjusted.gpx -lat 0.13\n\n", execName)
	fmt.Println(msg)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
