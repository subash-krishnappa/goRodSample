package sample

import (
	"log"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func TestSampleRodSetup(t *testing.T) {

	path, _ := launcher.LookPath()
	log.Println(path)
	debugUrl := launcher.New().
		Bin(path).
		Devtools(false).
		Headless(true).
		Leakless(true).
		NoSandbox(true).
		RemoteDebuggingPort(9222).
		Proxy(""). //provide your org proxy address here
		MustLaunch()
	log.Println(debugUrl)

	browser := rod.New().ControlURL(debugUrl)
	browser.SlowMotion(time.Second * 1)
	browser.Trace(true)

	page := browser.MustConnect().MustIgnoreCertErrors(true).MustPage("https://www.wikipedia.org/")

	page.MustElement("#searchInput").MustInput("earth")
	page.MustElement("#search-form > fieldset > button").MustClick()

	time.Sleep(time.Second * 3)
	page.Close()
	browser.Close()
}
