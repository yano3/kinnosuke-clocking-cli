package main

import (
	"flag"
	"fmt"
	"gopkg.in/headzoo/surf.v1"
	"os"
)

const kinnosukeUrl = "https://www.4628.jp/"
const ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36"
const clockingIdIn = "1"
const clockingIdOut = "2"

func main() {

	clockingOut := flag.Bool("out", false, "Clocking out")
	flag.Parse()

	var clockingId string
	if *clockingOut {
		clockingId = clockingIdOut
		fmt.Println("Clocking out...")
	} else {
		clockingId = clockingIdIn
		fmt.Println("Clocking in...")
	}

	browser := surf.NewBrowser()
	browser.SetUserAgent(ua)

	err := browser.Open(kinnosukeUrl)
	if err != nil {
		panic(err)
	}

	loginForm, _ := browser.Form("[id='form1']")
	loginForm.Input("y_companycd", os.Getenv("KINNOSUKE_COMPANYCD"))
	loginForm.Input("y_logincd", os.Getenv("KINNOSUKE_LOGINCD"))
	loginForm.Input("password", os.Getenv("KINNOSUKE_PASSWORD"))
	if loginForm.Submit() != nil {
		panic(err)
	}

	timeRecorderForm, _ := browser.Form("[id='tr_submit_form']")
	timeRecorderForm.Input("timerecorder_stamping_type", clockingId)
	if timeRecorderForm.Submit() != nil {
		panic(err)
	}

	fmt.Println("Success")
}
