package handler

import (
	"cruiseapp/database"
	"cruiseapp/dto"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

const StatsStmt = `
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
		year = strconv.Itoa(time.Now().Year())
	}
	rows, err := db.Query(StatsStmt, year)
	if err != nil {
		HandleError(err, w)
		return
	}

	var stats []dto.Statistics
	defer rows.Close()
	for rows.Next() {
		var s dto.Statistics
		if err := rows.Scan(
			&s.Year,
			&s.Month,
			&s.Count,
			&s.AvgHours,
		); err != nil {
			HandleError(err, w)
			return
		}
		stats = append(stats, s)
	}
	result := dto.StatisticsResponse{
		Data: stats,
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		HandleError(err, w)
		return
	}
}
