package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	envVar = "hey_thats_me"
	envVal = "1"
)

func main(){
	var userInputTime, userInputMsg string
	fmt.Println("Hey there! Want to send a reminder to your future self?")
	fmt.Print("Drop the time <hh:mm> and message: ")
	fmt.Scan(&userInputTime, &userInputMsg)
	fmt.Println(userInputTime + userInputMsg)

	now := time.Now()
	
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(userInputTime, now)

	if t == nil {
		fmt.Println("Can't parse time!")
		os.Exit(2)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if now.After(t.Time){
		fmt.Println("Pick a time in the future!")
		os.Exit(3)
	}

	timeDiff := t.Time.Sub(now)
	if os.Getenv(envVar) == envVal{
		time.Sleep(timeDiff)
		err = beeep.Alert("Reminder", userInputMsg, "assets/information.png" )
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		cmd := exec.Command(os.Args[0], userInputTime, userInputMsg)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", envVar, envVal))
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("Reminder will be displayed", timeDiff.Round(time.Second))
		os.Exit(0)
	}
}