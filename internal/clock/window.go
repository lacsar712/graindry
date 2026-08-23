package clock

import (
	"time"

	"github.com/lacsar712/graindry/internal/model"
)

type AvgTempWindow struct {
	clk      Clock
	duration time.Duration
}

func NewAvgTempWindow(clk Clock, duration time.Duration) *AvgTempWindow {
	if duration <= 0 {
		duration = 2 * time.Minute
	}
	return &AvgTempWindow{clk: clk, duration: duration}
}

func (w *AvgTempWindow) Active(anchor time.Time) bool {
	return WindowElapsed(w.clk, anchor, w.duration)
}

func (w *AvgTempWindow) Require(anchor time.Time) error {
	if w.Active(anchor) {
		return nil
	}
	return model.ErrGradientHold
}
