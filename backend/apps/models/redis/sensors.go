package redis

type Device struct {
	DvcId 		int64		`json:"dvcId"`
	Attrbs 		[]Attrb		`json:"attrbs"`
	//Name            *string		`json:"name"`
}

type Attrb struct {
	AttrbCd 	int32		`json:"attrbCd"`
	Value 		string		`json:"value"`
	Used 		bool		`json:"used"`
}

type Sensors struct {
	Interval 	int32		`json:"interval"`
	Devices 	[]Device	`json:"devices"`
}