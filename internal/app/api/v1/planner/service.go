package planner

import desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/planner"

// Implementation is a Service implementation
type Implementation struct {
	desc.UnimplementedPlannerServer

	plannerUC PlannerUC
}

// NewService return new instance of Implementation.
func NewService(plannerUC PlannerUC) *Implementation {
	return &Implementation{
		plannerUC: plannerUC,
	}
}
