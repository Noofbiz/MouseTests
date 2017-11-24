//+build mobilebind

package androidglue

import (
	"github.com/Noofbiz/mousetests"

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

func Touch(x, y, action int) {
	engo.TouchEvent(x, y, action)
}

func Stop() {
	running = false
	engo.MobileStop()
}
