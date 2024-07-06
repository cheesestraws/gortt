package gortt

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLineupExample(t *testing.T) {
	// This tests the example location lineup from the RTT API docs
	example := `
		{
		   "location": {
			   "name": "Bournemouth",
			   "crs": "BMH",
			   "tiploc": "BOMO"
		   },
		   "filter": null,
		   "services": [
			   {
				   "locationDetail": {
					   "realtimeActivated": true,
					   "tiploc": "BOMO",
					   "crs": "BMH",
					   "description": "Bournemouth",
					   "wttBookedArrival": "011630",
					   "wttBookedDeparture": "011830",
					   "gbttBookedArrival": "0117",
					   "gbttBookedDeparture": "0118",
					   "origin": [
						   {
							   "tiploc": "WATRLMN",
							   "description": "London Waterloo",
							   "workingTime": "230500",
							   "publicTime": "2305"
						   }
					   ],
					   "destination": [
						   {
							   "tiploc": "POOLE",
							   "description": "Poole",
							   "workingTime": "013000",
							   "publicTime": "0130"
						   }
					   ],
					   "isCall": true,
					   "isPublicCall": true,
					   "realtimeArrival": "0114",
					   "realtimeArrivalActual": false,
					   "realtimeDeparture": "0118",
					   "realtimeDepartureActual": false,
					   "platform": "3",
					   "platformConfirmed": false,
					   "platformChanged": false,
					   "displayAs": "CALL"
				   },
				   "serviceUid": "W90091",
				   "runDate": "2013-06-11",
				   "trainIdentity": "1B77",
				   "runningIdentity": "1B77",
				   "atocCode": "SW",
				   "atocName": "South West Trains",
				   "serviceType": "train",
				   "isPassenger": true
			   }
		   ]
		}`
	expected := RTTLocationLineup{
		Location: RTTLocationDetail{
			Name:   "Bournemouth",
			CRS:    "BMH",
			TIPLOC: "BOMO",
		},
		Services: []RTTLocationContainer{
			RTTLocationContainer{
				LocationDetail: RTTLocation{
					RealtimeActivated:   true,
					TIPLOC:              "BOMO",
					CRS:                 "BMH",
					Description:         "Bournemouth",
					WTTBookedArrival:    "011630",
					WTTBookedDeparture:  "011830",
					GBTTBookedArrival:   "0117",
					GBTTBookedDeparture: "0118",
					Origin: []RTTPair{
						RTTPair{
							TIPLOC:      "WATRLMN",
							Description: "London Waterloo",
							WorkingTime: "230500",
							PublicTime:  "2305",
						},
					},
					Destination: []RTTPair{
						RTTPair{
							TIPLOC:      "POOLE",
							Description: "Poole",
							WorkingTime: "013000",
							PublicTime:  "0130",
						},
					},
					IsCall:                  true,
					RealtimeArrival:         "0114",
					RealtimeArrivalActual:   false,
					RealtimeDeparture:       "0118",
					RealtimeDepartureActual: false,
					Platform:                "3",
					PlatformConfirmed:       false,
					PlatformChanged:         false,
					DisplayAs:               "CALL",
				},
				ServiceUID:      "W90091",
				RunDate:         "2013-06-11",
				TrainIdentity:   "1B77",
				RunningIdentity: "1B77",
				ATOCCode:        "SW",
				ATOCName:        "South West Trains",
				ServiceType:     "train",
				IsPassenger:     true,
			},
		},
	}

	var r RTTLocationLineup
	err := json.Unmarshal([]byte(example), &r)

	if err != nil {
		t.Errorf("failed to parse example: %v", err)
	}

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("unmarshalling gave wrong result")
	}
}
