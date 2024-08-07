package model

// Response for "timezome area location" and "timezome area location region".
type TimeZone struct {
	TimeZone     string `json:"timezone"`
	DayOfWeek    int    `json:"day_of_week"`
	DayOfYear    int    `json:"day_of_year"`
	DateTime     string `json:"datetime"`
	UtcDatetime  string `json:"utc_datetime"`
	Unixtime     int    `json:"unixtime"`
	RawOffset    int    `json:"raw_offset"`
	WeekNumber   int    `json:"week_number"`
	Dst          bool   `json:"dst"`
	Abbreviation string `json:"abbreviation"`
	DstOffset    int    `json:"dst_offset"`
	DstFrom      string `json:"dst_from"`
	DstUntil     string `json:"dst_until"`
	ClientIp     string `json:"client_ip"`
}
