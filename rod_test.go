package math

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func TestRodSetup(t *testing.T) {

	GetBrowserBinary("google")

	launch := GetBrowserLauncher()
	browserContext, _ := GetBrowserContext(launch)
	browserInstance := SetBrowserContextOptions(browserContext)

	launchedPage, _ := NewPageNavigate(browserInstance, "https://www.wikipedia.org/", time.Minute*1, 3)

	launchedPage.MustElement("#searchInput").MustInput("earth")
	launchedPage.MustElement("#search-form > fieldset > button").MustClick()

	time.Sleep(time.Second * 3)
	launchedPage.Close()
	browserContext.Close()
}

func GetBrowserBinary(host string) *launcher.Browser {
	browser := launcher.NewBrowser()
	if host == "google" {
		browser.Hosts = []launcher.Host{launcher.HostGoogle}
	} else if host == "npm" {
		browser.Hosts = []launcher.Host{launcher.HostNPM}
	} else if host == "azure" {
		browser.Hosts = []launcher.Host{launcher.HostPlaywright}
	}
	browser.Context = context.Background()
	browser.Revision = launcher.RevisionDefault
	browser.MustGet()
	return browser
}

func GetBrowserLauncher() *launcher.Launcher {
	launch := launcher.New().
		Devtools(false).
		Headless(true).
		Leakless(false).
		NoSandbox(true).
		StartURL("")
	return launch
}

func GetBrowserContext(launcher *launcher.Launcher) (*rod.Browser, error) {
	var browser *rod.Browser
	controlUrl, err := launcher.Launch()
	if err != nil {
		return nil, err
	}
	browser = rod.New().ControlURL(controlUrl).MustConnect().Context(context.Background())
	return browser, nil
}

func SetBrowserContextOptions(browserContext *rod.Browser) *rod.Browser {
	browserContext.SlowMotion(time.Second * 1)
	browserContext.Trace(true)
	return browserContext
}

func NewPageNavigate(browser *rod.Browser, desURL string, timeOut time.Duration, maxRetryTimes int) (*rod.Page, error) {

	page, err := NewPage(browser)
	if err != nil {
		return nil, err
	}
	page = page.Timeout(timeOut)
	err = rod.Try(func() {
		wait := page.MustWaitNavigation()
		pageInst := page.MustNavigate(desURL)
		log.Println(pageInst)
		if pageInst != nil {
			return
		} else {
			wait()
		}
	})
	if errors.Is(err, context.DeadlineExceeded) {
		return nil, err
	}
	return page, nil
}

func NewPage(browser *rod.Browser) (*rod.Page, error) {
	page, err := browser.Page(proto.TargetCreateTarget{URL: ""})
	if err != nil {
		return nil, err
	}
	return page, err
}
