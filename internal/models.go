package internal

type PlanRequest struct {
	Origin       string   `json:"origin"`
	Destinations []string `json:"destinations"`
	StartDate    string   `json:"start_date"` // YYYY-MM-DD
	EndDate      string   `json:"end_date"`   // YYYY-MM-DD
	BudgetKES    int      `json:"budget_kes"`
	Interests    []string `json:"interests"`
}

type BudgetBreakdown struct {
	Accommodation int `json:"accommodation"`
	Transport     int `json:"transport"`
	Food          int `json:"food"`
	Activities    int `json:"activities"`
	Total         int `json:"total"`
}

type DayPlan struct {
	Date       string   `json:"date"`
	City       string   `json:"city"`
	Plan       []string `json:"plan"`
	EstCostKES int      `json:"est_cost_kes"`
}

type Plan struct {
	ID      string          `json:"id"`
	Summary PlanSummary     `json:"summary"`
	Days    []DayPlan       `json:"itinerary"`
	Notes   []string        `json:"notes"`
	Request PlanRequest     `json:"-"`
	Budget  BudgetBreakdown `json:"budget_split"`
}

type PlanSummary struct {
	Nights int `json:"nights"`
}
