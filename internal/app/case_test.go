package app

import (
	"context"
	"testing"

	"github.com/lacsar712/graindry/internal/config"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := a.RunDryRamp(ctx, float64(config.Default().DryingRampMinutes)); err == nil {
		t.Fatal("expected cancel during dry ramp")
	}
}
