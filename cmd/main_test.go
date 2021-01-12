package main

import (
	"encoding/xml"
	"github.com/tobibot/gpxEdit/cmd/gpxStruct"
	"reflect"
	"testing"
)

func Test_readFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name       string
		args       args
		wantResult gpxStruct.GpxStruct
		wantErr    bool
	}{
		{name: "ReadFile", args: args{fileName: "../resources/test_twoPoints.gpx"}, wantErr: false, wantResult: GetTwoEntriesGpx()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := readFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("readFile() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func GetTwoEntriesGpx() gpxStruct.GpxStruct {
	return gpxStruct.GpxStruct{
		XMLName: xml.Name{
			Space: "http://www.topografix.com/GPX/1/1",
			Local: "gpx",
		},
		Text:           "\n    \n    \n",
		Xmlns:          "http://www.topografix.com/GPX/1/1",
		Gpxtpx:         "http://www.garmin.com/xmlschemas/TrackPointExtension/v1",
		Gpxx:           "http://www.garmin.com/xmlschemas/GpxExtensions/v3",
		Ns1:            "http://www.cluetrust.com/XML/GPXDATA/1/0",
		Xsi:            "http://www.w3.org/2001/XMLSchema-instance",
		Creator:        "Huami Amazfit Pace",
		Version:        "1.3",
		SchemaLocation: "http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd http://www.garmin.com/xmlschemas/GpxExtensions/v3 http://www.garmin.com/xmlschemas/GpxExtensionsv3.xsd http://www.garmin.com/xmlschemas/TrackPointExtension/v1 http://www.garmin.com/xmlschemas/TrackPointExtensionv1.xsd",
		Metadata:       GetMetadata(),

		Trk: GetTrk(),
	}
}

func GetTrk() gpxStruct.Trk {
	return gpxStruct.Trk{
		Text:   "\n        \n        \n        \n        \n    ",
		Name:   "Sport",
		Number: "2016",
		Type:   "run",
		Trkseg: GetTrkseg(),
	}
}

func GetTrkseg() gpxStruct.Trkseg {
	return gpxStruct.Trkseg{
		Text:  "\n            \n            \n        ",
		Trkpt: GetTrkpt(),
	}
}

func GetTrkpt() []gpxStruct.Trkpt {
	return []gpxStruct.Trkpt{
		{
			Text: "\n                \n                \n            ",
			Lat:  "51.00000000",
			Lon:  "7.00000000",
			Time: "2021-01-03T09:52:19Z",
			Extensions: gpxStruct.Extensions{
				Text: "\n                    \n                ",
				TrackPointExtension: gpxStruct.TrackPointExtension{
					Text: "\n                        \n                    ",
					Hr:   "102",
					Cad:  "",
				},
			},
			Ele: "",
		},
		{
			Text: "\n                \n                \n            ",
			Lat:  "50.95363072",
			Lon:  "7.13367296",
			Time: "2021-01-03T09:52:19Z",
			Extensions: gpxStruct.Extensions{
				Text: "\n                    \n                ",
				TrackPointExtension: gpxStruct.TrackPointExtension{
					Text: "\n                        \n                    ",
					Hr:   "102",
					Cad:  "",
				},
			},
			Ele: "",
		},
	}
}

func GetMetadata() gpxStruct.Metadata {
	return gpxStruct.Metadata{
		Text: "\n        \n        \n    ",
		Name: "Huami Amazfit Sports Watch",
		Time: "2021-01-03T09:52:17Z",
	}
}

func Test_adjustGpx(t *testing.T) {
	type args struct {
		data gpxStruct.GpxStruct
		lon  float64
		lat  float64
	}
	tests := []struct {
		name       string
		args       args
		wantResult gpxStruct.GpxStruct
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := adjustGpx(tt.args.data, tt.args.lon, tt.args.lat)
			if (err != nil) != tt.wantErr {
				t.Errorf("adjustGpx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("adjustGpx() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_writeFile(t *testing.T) {
	type args struct {
		fileName string
		data     gpxStruct.GpxStruct
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := writeFile(tt.args.fileName, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("writeFile() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
