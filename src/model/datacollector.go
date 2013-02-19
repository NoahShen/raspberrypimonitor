package model

import ()

type DataCollector interface {
	StartGetData(dsCh chan<- *DataValue)
	Stop()
}

type DataValue struct {
	Id    string
	Value string
}
