package cosm

import (
	"time"
)

const (
	Type_Success = 1
	Type_Failed  = 0
)

type Status string

const (
	Live   = Status("live")
	Frozen = Status("frozen")
)

type Exposure string

const (
	Indoors  = Exposure("indoors")
	Outdoors = Exposure("outdoors")
)

type Feed struct {
	Title       string       `json:"title,omitempty"`
	Id          string       `json:"id,omitempty"`
	UpdatedTime *time.Time   `json:"updated,omitempty"`
	CreatedTime *time.Time   `json:"created,omitempty"`
	Creator     string       `json:"creator,omitempty"`
	Url         string       `json:"feed,omitempty"`
	Status      Status       `json:"status,omitempty"`
	Description string       `json:"description,omitempty"`
	Website     string       `json:"website,omitempty"`
	Icon        string       `json:"icon,omitempty"`
	Tags        []string     `json:"tags,omitempty"`
	Location    *Location    `json:"location,omitempty"`
	Datastreams []Datastream `json:"datastreams,omitempty"`
	Private     bool         `json:"private,omitempty"`
}

type Location struct {
	Name        string   `json:"name"`
	Domain      string   `json:"domain,omitempty"`
	Exposure    Exposure `json:"exposure,omitempty"`
	Disposition string   `json:"disposition,omitempty"`
	Lat         float64  `json:"lat,omitempty"`
	lon         float64  `json:"lon,omitempty"`
	Elevation   float64  `json:"ele,omitempty"`
}

type Datastream struct {
	Id           string      `json:"id,omitempty"`
	UpdatedTime  *time.Time  `json:"at,omitempty"`
	Tags         []string    `json:"tags,omitempty"`
	Unit         *Unit       `json:"units,omitempty"`
	MinValue     string      `json:"min_value,omitempty"`
	MaxValue     string      `json:"max_value,omitempty"`
	CurrentValue string      `json:"current_value,omitempty"`
	Datapoints   []Datapoint `json:"datapoints,omitempty"`
}

type Unit struct {
	Type   string `json:"label,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}

type Datapoint struct {
	At    time.Time `json:"at,omitempty"`
	Value string    `json:"value"`
}
