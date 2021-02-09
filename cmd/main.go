package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/tobibot/gpxEdit/cmd/gpxStruct"
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

	for i, trk := range data.Trks {
		for j, trkseg := range trk.Trksegs {
			for k, trkpt := range trkseg.Waypoints {
				originalLat, err := strconv.ParseFloat(trkpt.Lat, 64)
				if err != nil {
					return &gpxStruct.GpxStruct{}, err
				}
				newLat := originalLat + latAdjustment
				if newLat < -90 { // -100
					newLat = 180 + newLat
				} else if newLat > 90 {
					newLat = -180 + newLat
				}
				trkpt.Lat = strconv.FormatFloat(newLat, 'f', 8, 64)

				originalLon, err := strconv.ParseFloat(trkpt.Lon, 64)
				if err != nil {
					return &gpxStruct.GpxStruct{}, err
				}
				newLon := originalLon + lonAdjustment
				if newLon < -180 { // -190
					newLon = 360 + newLon
				} else if newLon > 180 {
					newLon = -360 + newLon
				}
				trkpt.Lon = strconv.FormatFloat(newLon, 'f', 8, 64)
				data.Trks[i].Trksegs[j].Waypoints[k] = trkpt
			}
		}
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
