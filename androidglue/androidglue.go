//+build mobilebind

package androidglue

import (
	"github.com/Noofbiz/MouseTests"

	"engo.io/engo"
)

var running bool

func Start(width, height int) {
	running = true
	mousetests.Start(width, height)
}

func Update() {
	engo.RunIteration()
}

func IsRunning() bool {
	return running
}

func Touch(x, y, id, action int) {
	engo.TouchEvent(x, y, id, action)
}

func Stop() {
	running = false
	engo.MobileStop()
}
