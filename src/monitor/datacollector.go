package main

import ()

type DataCollector interface {
	StartGetData()

	Stop()
}
