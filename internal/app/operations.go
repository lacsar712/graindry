package app

import (
	"context"
	"fmt"
	"time"

	"github.com/lacsar712/graindry/internal/model"
)

func (a *App) ValidateMoistureDrift(ctx context.Context, moistPct float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	limit := a.cfg.TargetMoistPct + a.cfg.MaxGradientDeltaPct
	if moistPct <= limit {
		return nil
	}
	return fmt.Errorf("moisture: %w", model.ErrMoistureDrift)
}

func (a *App) ConfirmGradientHold(ctx context.Context, anchor time.Time) error {
	if a.avgWindow == nil {
		return model.Wrap("app", "window", model.ErrGradientHold)
	}
	if err := a.avgWindow.Require(anchor); err != nil {
		return fmt.Errorf("gradient hold: %w", err)
	}
	return nil
}
