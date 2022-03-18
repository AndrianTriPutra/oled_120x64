package display

import (
	"log"
	"time"

	Oled "github.com/mmalcek/nanohatoled"

	. "oled/time"
)

func Oled_Scroll(interval time.Duration) {
	var (
		side_oled uint8
	)

	oled, err := Oled.Open()
	if err != nil {
		log.Println("[ERROR] [Oled_Scroll]", "error Open")
	}
	defer oled.Close()

	ticker := time.NewTicker(interval * time.Second)
	side_oled = 1
	for {
		select {
		case <-ticker.C:
			var lines [6]string

			date_oled := Get_date("Asia/Jakarta")
			hms_oled := Get_hms("Asia/Jakarta")

			switch side_oled {
			case 1:
				//side A insert
				lines[0] = hms_oled
				lines[1] = "1 :"
				lines[2] = "2 :"
				lines[3] = "3 :"
				lines[4] = "4 :"
				lines[5] = date_oled

				oled.Clear()
				oled.New(180)
				oled.Text(25, 00, lines[0], true)
				oled.Text(0, 10, lines[1], true)
				oled.Text(0, 20, lines[2], true)
				oled.Text(0, 30, lines[3], true)
				oled.Text(0, 40, lines[4], true)
				oled.Text(15, 50, lines[5], true)
				oled.Send()

				side_oled++

			case 2:
				//side B insert
				lines[0] = hms_oled
				lines[1] = "A :"
				lines[2] = "B :"
				lines[3] = "C :"
				lines[4] = "D :"
				lines[5] = date_oled

				oled.Clear()
				oled.New(180)
				oled.Text(25, 00, lines[0], true)
				oled.Text(0, 10, lines[1], true)
				oled.Text(0, 20, lines[2], true)
				oled.Text(0, 30, lines[3], true)
				oled.Text(0, 40, lines[4], true)
				oled.Text(15, 50, lines[5], true)
				oled.Send()

				side_oled = 1

			}
		}
	}
}
