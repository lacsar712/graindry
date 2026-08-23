package app

import (
	"context"

	"github.com/lacsar712/graindry/internal/clock"
)

// SegmentPlan describes inter-segment vent staging for drying batches.

type SegmentPlan struct {
	VentSteps int
}

func (a *App) ExecutePlan(ctx context.Context, plan SegmentPlan) error {
	if a.scheduler == nil {
		return nil
	}
	// Propagate the caller's context (typically the batch scope ctx) so an
	// operator revoke that cancels the batch also aborts the schedule queue.
	// Using context.Background() here would leave the old drying steps being
	// appended even after the batch cancellation — the cancel would stop at
	// the console layer while the backend keeps installing steps.
	return a.scheduler.InstallVentPlanCtx(ctx, clock.VentPlan{VentSteps: plan.VentSteps}, "segment-plan")
}

func (a *App) SegmentVentStepsDone() int {
	if a.scheduler == nil {
		return 0
	}
	return a.scheduler.VentStepsDone()
}
