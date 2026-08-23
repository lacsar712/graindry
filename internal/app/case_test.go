package app

import (
	"context"
	"errors"
	"testing"

	"github.com/lacsar712/graindry/internal/config"
	"github.com/lacsar712/graindry/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	err = a.HandleHeatOvertemp(context.Background(), model.TowerID(a.cfg.TowerID), 95.0)
	if err == nil {
		t.Fatal("expected heat overtemperature error")
	}
	if !errors.Is(err, model.ErrHeatOvertemp) {
		t.Fatalf("expected ErrHeatOvertemp, got %v", err)
	}
}
