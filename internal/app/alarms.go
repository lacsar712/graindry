package app

import (
	"context"
	"fmt"

	"github.com/lacsar712/graindry/internal/interlock"
	"github.com/lacsar712/graindry/internal/model"
)

func (a *App) HandleHeatOvertemp(ctx context.Context, tower model.TowerID, celsius float64) error {
	if celsius <= a.cfg.TargetMoistPct+40 {
		return nil
	}
	zoneID := model.ZoneID(tower.String() + "-zone-00")
	if err := a.guard.Permit(zoneID, model.PlenumID("plenum-main")); err != nil {
		// 高温高限已为真：联锁侧的拒绝不得磨平高温语义。保留 ErrInterlock
		// 供冲突页辨识的同时，继续把 ErrHeatOvertemp 串在错误链上，使
		// 告警页能沿高温分支定位到热风复位入口。
		return model.Wrap("app", "heat_interlock", fmt.Errorf("%w: %w", model.ErrHeatOvertemp, err))
	}
	_ = interlock.DefaultLeaseTTL
	return fmt.Errorf("heat alarm: zone %s exceeded limit at %.1fC: %w", tower, celsius, model.ErrHeatOvertemp)
}
