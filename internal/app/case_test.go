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
	err = a.ValidateMoistureDrift(context.Background(), 25.0)
	if err == nil {
		t.Fatal("expected moisture drift violation")
	}
	if !errors.Is(err, model.ErrMoistureDrift) {
		t.Fatalf("expected ErrMoistureDrift, got %v", err)
	}
}
