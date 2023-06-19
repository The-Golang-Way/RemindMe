package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName = "hey_thats_me"
	markValue = "1"
)

func main(){
	if len(os.Args) < 3  {
		fmt.Printf("Usage:%s <time> <message>\n", os.Args[0])
		os.Exit(1)
	}

	now := time.Now()
	when := when.New(nil)

	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1], now)

	if t != nil {
		fmt.Println("Can't parse time!")
		panic(t)
	}

	if err != nil {
		panic(err)
	}

	if now.After(t.Time){
		fmt.Println("pick a time in the future!")
		os.Exit(2)
	}

	timeDiff := t.Time.Sub(now)
	if os.Getenv(markName) == markValue{
		time.Sleep(timeDiff)
		beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png" )
	}
}