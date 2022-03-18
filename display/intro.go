package display

import (
	"log"
	"time"

	Oled "github.com/mmalcek/nanohatoled"
)

func Oled_Intro() {

	oled, err := Oled.Open()
	if err != nil {
		log.Println("[ERROR] [Oled_Intro]", "error Open")
	}
	defer oled.Close()

	oled.Clear()
	oled.New(180)
	oled.Text(15, 30, "Welcome to", true)
	oled.Text(15, 50, "GOLANG DEV", true)
	oled.Send()
	time.Sleep(2 * time.Second)

}
