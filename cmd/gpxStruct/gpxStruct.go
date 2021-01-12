package gpxStruct

import "encoding/xml"

type GpxStruct struct {
	XMLName        xml.Name `xml:"gpx"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Gpxtpx         string   `xml:"gpxtpx,attr"`
	Gpxx           string   `xml:"gpxx,attr"`
	Ns1            string   `xml:"ns1,attr"`
	Xsi            string   `xml:"xsi,attr"`
	Creator        string   `xml:"creator,attr"`
	Version        string   `xml:"version,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Metadata       Metadata `xml:"metadata"`
	Trk            Trk      `xml:"trk"`
}

type Metadata struct {
	Text string `xml:",chardata"`
	Name string `xml:"name"`
	Time string `xml:"time"`
}

type Trk struct {
	Text   string `xml:",chardata"`
	Name   string `xml:"name"`
	Number string `xml:"number"`
	Type   string `xml:"type"`
	Trkseg Trkseg `xml:"trkseg"`
}

type Trkseg struct {
	Text  string  `xml:",chardata"`
	Trkpt []Trkpt `xml:"trkpt"`
}

type Trkpt struct {
	Text       string     `xml:",chardata"`
	Lat        string     `xml:"lat,attr"`
	Lon        string     `xml:"lon,attr"`
	Time       string     `xml:"time"`
	Extensions Extensions `xml:"extensions"`
	Ele        string     `xml:"ele"`
}

type Extensions struct {
	Text                string              `xml:",chardata"`
	TrackPointExtension TrackPointExtension `xml:"TrackPointExtension"`
}

type TrackPointExtension struct {
	Text string `xml:",chardata"`
	Hr   string `xml:"hr"`
	Cad  string `xml:"cad"`
}
