package entity

import "time"

type SummariesRequest struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Project  string    `json:"project"`
	Timezone string    `json:"timezone"`
	Range    string    `json:"range"`
}

type SummariesResponse struct {
	Data            []Data          `json:"data"`
	Start           time.Time       `json:"start"`
	End             time.Time       `json:"end"`
	CumulativeTotal CumulativeTotal `json:"cumulative_total"`
	DailyAverage    DailyAverage    `json:"daily_average"`
}
type Data struct {
	Languages        []Languages        `json:"languages"`
	GrandTotal       GrandTotal         `json:"grand_total"`
	Editors          []Editors          `json:"editors"`
	OperatingSystems []OperatingSystems `json:"operating_systems"`
	Categories       []Categories       `json:"categories"`
	Dependencies     []Dependencies     `json:"dependencies"`
	Machines         []Machines         `json:"machines"`
	Projects         []Projects         `json:"projects"`
	Range            Range              `json:"range"`
}
type Languages struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Digital      string  `json:"digital"`
	Decimal      string  `json:"decimal"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
	Percent      float64 `json:"percent"`
}
type GrandTotal struct {
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	TotalSeconds float64 `json:"total_seconds"`
	Digital      string  `json:"digital"`
	Decimal      string  `json:"decimal"`
	Text         string  `json:"text"`
}
type Editors struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Digital      string  `json:"digital"`
	Decimal      string  `json:"decimal"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
	Percent      float64 `json:"percent"`
}
type OperatingSystems struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Digital      string  `json:"digital"`
	Decimal      string  `json:"decimal"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
	Percent      float64 `json:"percent"`
}
type Categories struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Digital      string  `json:"digital"`
	Decimal      string  `json:"decimal"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
	Percent      float64 `json:"percent"`
}
type Dependencies struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Digital      string  `json:"digital"`
	Decimal      string  `json:"decimal"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
	Percent      float64 `json:"percent"`
}
type Machines struct {
	Name          string  `json:"name"`
	TotalSeconds  float64 `json:"total_seconds"`
	MachineNameID string  `json:"machine_name_id"`
	Digital       string  `json:"digital"`
	Decimal       string  `json:"decimal"`
	Text          string  `json:"text"`
	Hours         int     `json:"hours"`
	Minutes       int     `json:"minutes"`
	Seconds       int     `json:"seconds"`
	Percent       float64 `json:"percent"`
}
type Projects struct {
	Name         string      `json:"name"`
	TotalSeconds float64     `json:"total_seconds"`
	Color        interface{} `json:"color"`
	Digital      string      `json:"digital"`
	Decimal      string      `json:"decimal"`
	Text         string      `json:"text"`
	Hours        int         `json:"hours"`
	Minutes      int         `json:"minutes"`
	Seconds      int         `json:"seconds"`
	Percent      float64     `json:"percent"`
}

type CumulativeTotal struct {
	Seconds float64 `json:"seconds"`
	Text    string  `json:"text"`
	Digital string  `json:"digital"`
	Decimal string  `json:"decimal"`
}
type DailyAverage struct {
	Holidays                      int    `json:"holidays"`
	DaysMinusHolidays             int    `json:"days_minus_holidays"`
	DaysIncludingHolidays         int    `json:"days_including_holidays"`
	Seconds                       int    `json:"seconds"`
	SecondsIncludingOtherLanguage int    `json:"seconds_including_other_language"`
	Text                          string `json:"text"`
	TextIncludingOtherLanguage    string `json:"text_including_other_language"`
}

type Range struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Date     string    `json:"date"`
	Text     string    `json:"text"`
	Timezone string    `json:"timezone"`
}
