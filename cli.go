package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/Songmu/prompter"
	"gopkg.in/headzoo/surf.v1"
)

const kinnosukeUrl string = "https://www.4628.jp/"
const ua string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36"
const clockingIdIn string = "1"
const clockingIdOut string = "2"

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

func clockIn(clockingOut bool) error {
	var clockingId string
	if clockingOut {
		clockingId = clockingIdOut
	} else {
		clockingId = clockingIdIn
	}

	browser := surf.NewBrowser()
	browser.SetUserAgent(ua)

	if err := browser.Open(kinnosukeUrl); err != nil {
		return err
	}

	loginForm, _ := browser.Form("[id='form1']")
	loginForm.Input("y_companycd", os.Getenv("KINNOSUKE_COMPANYCD"))
	loginForm.Input("y_logincd", os.Getenv("KINNOSUKE_LOGINCD"))
	loginForm.Input("password", os.Getenv("KINNOSUKE_PASSWORD"))
	if err := loginForm.Submit(); err != nil {
		return err
	}

	// check if operated from internal network
	mes := browser.Find(".txt_12_red").Text()
	if len(mes) > 0 {
		return errors.New(mes)
	}

	timeRecorderForm, _ := browser.Form("[id='tr_submit_form']")
	timeRecorderForm.Input("timerecorder_stamping_type", clockingId)
	if err := timeRecorderForm.Submit(); err != nil {
		return err
	}

	selection := browser.Find("#timerecorder_txt")
	reg := regexp.MustCompile(`\d\d:\d\d`)

	clockInTime := reg.FindString(selection.Eq(0).Text())
	clockOutTime := reg.FindString(selection.Eq(1).Text())

	if clockingOut {
		fmt.Printf("%s %s\n", clockInTime, clockOutTime)
	} else {
		fmt.Println(clockInTime)
	}

	return nil
}

func (cli *CLI) Run(args []string) int {
	var (
		yes bool
		out bool

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&yes, "yes", false, "Skip y/n prompt")
	flags.BoolVar(&yes, "y", false, "Skip y/n prompt (Short)")
	flags.BoolVar(&out, "out", false, "Clocking out")
	flags.BoolVar(&out, "o", false, "Clocking out (Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if !yes && !prompter.YN("OK?", true) {
		fmt.Println("Canceled")
		return ExitCodeError
	}

	err := clockIn(out)
	if err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	return ExitCodeOK
}
