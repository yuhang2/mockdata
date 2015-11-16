package operation

import (
	"fmt"

	"github.com/yuhang2/mockdata/config"
)

func Location(line int) {
	fmt.Println("Get Location", line)
	fmt.Println(config.Config)
}

/*
var BookingList = func(start time.Time, stop time.Time, fileName *string) error {
	db, err := sql.Open("postgres", string(config.Config))
	defer db.Close()
	if err != nil {
		return err
	}
	listQuery := "select booking_code, created_at_local from grab_road_bookings where created_at_local >= $1 and created_at_local <= $2 and picking_up_time is not NULL and dropping_off_time is not NULL"
	rows, err := db.Query(listQuery, start, stop)
	if err != nil {
		return err
	}
	defer rows.Close()
	var bookings [][]string
	var booking_code string
	var create_at_local time.Time
	for rows.Next() {
		err = rows.Scan(&booking_code, &create_at_local)
		if err != nil {
			log.Printf("scan grab_road_bookings error: %v", err)
			continue
		}
		bookings = append(bookings, []string{
			booking_code,
			create_at_local.Format("2006-01-02T15:04:05"),
		})
	}
	fileOut, err := os.Create(*fileName)
	defer fileOut.Close()
	if err != nil {
		log.Printf("create file %s error: %v", *fileName, err)
		return nil
	}
	w := csv.NewWriter(fileOut)
	for _, record := range bookings {
		if err := w.Write(record); err != nil {
			log.Printf("error writing record to csv, err: %v", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Printf("error: ", err)
	}
	return nil
}
*/
