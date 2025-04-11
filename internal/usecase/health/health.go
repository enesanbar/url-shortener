package health

import (
	"context"

	"go.uber.org/fx"

	"github.com/enesanbar/go-service/healthchecker"
)

type Checker struct {
}

type Params struct {
	fx.In
}

func NewChecker(p Params) *Checker {
	return &Checker{}
}

// Name returns the name of the health checker.
func (c *Checker) Name() string {
	return "my-health-checker"
}

// Check performs the health check and returns the result.
func (c *Checker) Check(ctx context.Context) *healthchecker.HealthCheckerProbeResult {
	return healthchecker.NewHealthCheckerProbeResult(true, "database ping success")
}
