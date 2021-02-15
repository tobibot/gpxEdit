package gpxStruct

import "encoding/xml"

type GpxStruct struct {
	XMLName        xml.Name  `xml:"gpx"`
	Text           string    `xml:",chardata"`
	Xmlns          string    `xml:"xmlns,attr"`
	Gpxtpx         string    `xml:"gpxtpx,attr"`
	Gpxx           string    `xml:"gpxx,attr"`
	Ns1            string    `xml:"ns1,attr,omitempty"`
	Xsi            string    `xml:"xsi,attr,omitempty"`
	Creator        string    `xml:"creator,attr"`
	Version        string    `xml:"version,attr"`
	SchemaLocation string    `xml:"schemaLocation,attr"`
	Metadata       Metadata  `xml:"metadata,omitempty"`
	Trks           []Trk     `xml:"trk"`
	Routes         []Rte     `xml:"rte,omitempty"`
	Waypoints      Waypoints `xml:"wpt,omitempty"`
}

type Metadata struct {
	Text string `xml:",chardata"`
	Name string `xml:"name"`
	Time string `xml:"time"`
}

type Trk struct {
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name"`
	Number  string   `xml:"number"`
	Type    string   `xml:"type"`
	Trksegs []Trkseg `xml:"trkseg"`
}

type Rte struct {
	XMLName   xml.Name  `xml:"rte"`
	Name      string    `xml:"name,omitempty"`
	Cmt       string    `xml:"cmt,omitempty"`
	Desc      string    `xml:"desc,omitempty"`
	Src       string    `xml:"src,omitempty"`
	Links     []Link    `xml:"link"`
	Number    int       `xml:"number,omitempty"`
	Type      string    `xml:"type,omitempty"`
	Waypoints Waypoints `xml:"wpt,omitempty"`
}

type Trkseg struct {
	Text      string `xml:",chardata"`
	Waypoints `xml:"trkpt"`
}

type Waypoints []Wpt

type Wpt struct {
	Text       string     `xml:",chardata"`
	Lat        string     `xml:"lat,attr"`
	Lon        string     `xml:"lon,attr"`
	Time       string     `xml:"time"`
	Extensions Extensions `xml:"extensions,omitempty"`
	Ele        string     `xml:"ele"`
}

type Link struct {
	XMLName xml.Name `xml:"link"`
	URL     string   `xml:"href,attr,omitempty"`
	Text    string   `xml:"text,omitempty"`
	Type    string   `xml:"type,omitempty"`
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
