package operation

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/yuhang2/mockdata/config"
)

var (
	db  *sql.DB
	err error
)

func Location(line int) {
	db, err = sql.Open("postgres", config.Config.RedShift.String())
	if err != nil {
		fmt.Printf("Err: %v", err)
	}
	defer db.Close()

	start := 0
	limit := 1000
	channelNum := int(math.Ceil(float64(line) / float64(limit)))
	execChannel := make(chan string, channelNum)

	type booking struct {
		pickUpLatitude   float64
		pickUpLongitude  float64
		dropOffLatitude  float64
		dropOffLongitude float64
		createTime       time.Time
	}
	var pickUpLatitude, pickUpLongitude, dropOffLatitude, dropOffLongitude float64
	var createTime time.Time
	bookingsChan := make(chan booking, line)

	csvFile, err := os.Create("bookings.csv")
	if err != nil {
		log.Println("create bookings.csv failed")
		return
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for line > 0 {
		if line < limit {
			limit = line
		}
		go func(start, limit int) {
			listQuery := "select pick_up_latitude, pick_up_longitude, drop_off_latitude, drop_off_longitude, created_at_local" +
				" from grab_road_bookings limit $2 offset $1"
			rows, err := db.Query(listQuery, start, limit)
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			defer rows.Close()

			for rows.Next() {
				err := rows.Scan(&pickUpLatitude, &pickUpLongitude, &dropOffLatitude, &dropOffLongitude, &createTime)
				if err != nil {
					log.Printf("scan grab_road_bookings error: %v\n", err)
					continue
				}
				bookingsChan <- booking{
					pickUpLatitude:   pickUpLatitude,
					pickUpLongitude:  pickUpLongitude,
					dropOffLatitude:  dropOffLatitude,
					dropOffLongitude: dropOffLongitude,
					createTime:       createTime,
				}
			}

			execChannel <- "ok"
		}(start, limit)
		start += limit
		line -= limit
	}

	for i := 0; i < channelNum; i++ {
		<-execChannel
	}
	records := []string{}
	for info := range bookingsChan {
		records = []string{
			strconv.FormatFloat(info.dropOffLatitude, 'f', 8, 64),
			strconv.FormatFloat(info.dropOffLatitude, 'f', 8, 64),
			strconv.FormatFloat(info.dropOffLatitude, 'f', 8, 64),
			strconv.FormatFloat(info.dropOffLatitude, 'f', 8, 64),
			info.createTime.Format("2006-01-02T15:04:05"),
		}
		err := writer.Write(records)
		if err != nil {
			fmt.Printf("save bookings info error, info: %v\n", info)
		}
	}
	writer.Flush()
}
