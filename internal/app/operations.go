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
	tol := a.cfg.MoistureTolerancePct
	if tol < 0 {
		tol = 0
	}
	lo := a.cfg.TargetMoistPct - tol
	hi := a.cfg.TargetMoistPct + tol
	if moistPct >= lo && moistPct <= hi {
		return nil
	}
	// 双向容差带外飘出：包成 ErrMoistureDrift，使调度端可经 errors.Is 走含水偏离处置分支。
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
