package gomesure

import (
	"testing"
	"time"
)

func TestMesure(t *testing.T) {
	var timing Mesure
	timing.Starttime()
	time.Sleep(1 * time.Second)
	timing.Elapsedtime()
	if timing.GetSecondes() < 1. {
	}
}
