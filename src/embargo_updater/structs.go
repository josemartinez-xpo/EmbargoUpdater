package main

type APIResponse struct {
	Code string
	Data EmbargoData
}

type EmbargoData struct {
	EmbargoLocation []EmbargoItem
}

type EmbargoItem struct {
	DestSic string
	DestZip string
	EmbargoType string
	EmbargoId string
	StartDate int64
	EndDate int64
	InclZoneInd bool
	InclSatelliteInd bool
}