package gortt

type RTTLocationLineup struct {
	Location RTTLocationDetail      `json:"location"`
	Services []RTTLocationContainer `json:"services"`
}

type RTTLocationDetail struct {
	// This type is	referred to but not defined in the docs
	// unless I am missing something.  I have inferred its
	// structure from data returned from the API.

	Name    string `json:"name"`
	CRS     string `json:"CRS"`
	TIPLOC  string `json:"TIPLOC"`
	Country string `json:"country"`
	System  string `json:"system"`
}

type RTTLocationContainer struct {
	LocationDetail   RTTLocation `json:"locationDetail"`
	ServiceUID       string      `json:"serviceUid"`
	RunDate          string      `json:"runDate"`
	TrainIdentity    string      `json:"trainIdentity"`
	RunningIdentity  string      `json:"runningIdentity"`
	ATOCCode         string      `json:"atocCode"`
	ATOCName         string      `json:"atocName"`
	ServiceType      string      `json:"serviceType"`
	IsPassenger      bool        `json:"isPassenger"`
	PlannedCancel    bool        `json:"plannedCancel"`
	Origin           []RTTPair   `json:"origin"`
	Destination      []RTTPair   `json:"destination"`
	CountdownMinutes int         `json:"countdownMinutes"`
}

type RTTPair struct {
	TIPLOC      string `json:"tiploc"`
	Description string `json:"description"`
	WorkingTime string `json:"workingTime"`
	PublicTime  string `json:"publicTime"`
}

type RTTLocation struct {
	RealtimeActivated   bool   `json:"realtimeActivated"`
	TIPLOC              string `json:"tiploc"`
	CRS                 string `json:"crs"`
	Description         string `json:"description"`
	WTTBookedArrival    string `json:"wttBookedArrival"`
	WTTBookedDeparture  string `json:"wttBookedDeparture"`
	GBTTBookedArrival   string `json:"gbttBookedArrival"`
	GBTTBookedDeparture string `json:"gbttBookedDeparture"`

	Origin       []RTTPair `json:"origin"`
	Destination  []RTTPair `json:"destination"`
	IsCall       bool      `json:"isCall"`
	IsCallPublic bool      `json:"isCallPublic"` // example calls this isPublicCall - to raise upstream?

	RealtimeArrival             string `json:"realtimeArrival"`
	RealtimeArrivalActual       bool   `json:"realtimeArrivalActual"`
	RealtimeArrivalNoReport     bool   `json:"realtimeArrivalNoReport"`
	RealtimeWTTArrivalLateness  int    `json:"realtimeWttArrivalLateness"`
	RealtimeGBTTArrivalLateness int    `json:"realtineGbttArrivalLateness"`

	RealtimeDeparture             string `json:"realtimeDeparture"`
	RealtimeDepartureActual       bool   `json:"realtimeDepartureActual"`
	RealtimeDepartureNoReport     bool   `json:"realtimeDepartureNoReport"`
	RealtimeWTTDepartureLateness  int    `json:"realtimeWttDepartureLateness"`
	RealtimeGBTTDepartureLateness int    `json:"realtimeGbttDepartureLateness"`

	RealtimePass         string `json:"realtimePass"`
	RealtimePassActual   bool   `json:"realtimePassActual"`
	RealtimePassNoReport bool   `json:"realtimePassNoReport"`

	/* At this point in the API documentation, RealtimeWTTArrivalLateness
	   is defined again - probably worth raising this upstream.  A field
	   may be missing here that the duplicate has replaced... */

	Platform          string `json:"platform"`
	PlatformConfirmed bool   `json:"platformConfirmed"`
	PlatformChanged   bool   `json:"platformChanged"`
	Line              string `json:"line"`
	LineConfirmed     bool   `json:"lineConfirmed"`
	Path              string `json:"path"`
	PathConfirmed     bool   `json:"pathConfirmed"`

	CancelReasonCode      string `json:"cancelReasonCode"`
	CancelReasonShortText string `json:"cancelReasonShortText"`
	CancelReasonLongText  string `json:"cancelReasonLongText"`

	DisplayAs       string `json:"displayAs"`
	ServiceLocation string `json:"ServiceLocation"`
}
