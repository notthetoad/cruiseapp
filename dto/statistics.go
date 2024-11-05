package dto

type Statistics struct {
	Year     int
	Month    int
	Count    int
	AvgHours float64
}

type StatisticsResponse struct {
	Data []Statistics
}
