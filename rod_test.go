package math

import (
	"log"
	"testing"

	"github.com/go-rod/rod/lib/launcher"
)

func TestRodSetup(t *testing.T) {
	str := launcher.NewBrowser().MustGet()
	log.Println(str)
}
