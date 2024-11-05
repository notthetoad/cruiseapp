package handler

import (
	"cruiseapp/database"
	"encoding/json"
	"net/http"
)

type statistics struct {
	Year     int
	Month    int
	Count    int
	AvgHours float64
}

const StatStmt = `
SELECT extract(year from c.start_date) as year,
       extract(month from c.start_date) as month,
       count(*),
       round(extract(epoch from avg(c.end_date - c.start_date))/3600, 2) as avg_hours
  FROM cruise c
  JOIN cruise_person p
    ON c.id = p.cruise_id
 WHERE extract(year from c.start_date) = $1
 GROUP by year, month
 ORDER by year, month;`

func StatisticsHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDb(r)
	query := r.URL.Query()
	year := query.Get("year")
	if year == "" {
		year = "2023"
	}
	res, err := db.Query(StatStmt, year)
	if err != nil {
		HandleError(err, w)
		return
	}
	var results []statistics
	for res.Next() {
		var s statistics
		if err := res.Scan(
			&s.Year,
			&s.Month,
			&s.Count,
			&s.AvgHours,
		); err != nil {
			HandleError(err, w)
			return
		}
		results = append(results, s)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&results)
	if err != nil {
		HandleError(err, w)
		return
	}
}
