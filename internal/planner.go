package internal

import (
	"errors"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

var sampleActivities = map[string][]string{
	"Nairobi":   {"National Museum visit", "Local food tour", "Karura walk"},
	"Diani":     {"Beach afternoon", "Snorkeling", "Seafood dinner"},
	"Mombasa":   {"Old Town walk", "Fort Jesus visit", "Street food crawl"},
	"Maasai Mara": {"Game drive", "Sundowner", "Village visit"},
}

// Validate basic inputs and parse dates.
func validate(req PlanRequest) (time.Time, time.Time, error) {
	if req.BudgetKES <= 0 {
		return time.Time{}, time.Time{}, errors.New("budget_kes must be > 0")
	}
	if len(req.Destinations) == 0 {
		return time.Time{}, time.Time{}, errors.New("destinations cannot be empty")
	}
	s, err := time.Parse(dateLayout, req.StartDate)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("start_date must be YYYY-MM-DD")
	}
	e, err := time.Parse(dateLayout, req.EndDate)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("end_date must be YYYY-MM-DD")
	}
	if !e.After(s) && !e.Equal(s) {
		return time.Time{}, time.Time{}, errors.New("end_date must be >= start_date")
	}
	return s, e, nil
}

// AllocateBudget uses a base split and tweaks for interests.
func AllocateBudget(total int, interests []string) BudgetBreakdown {
	// Base split: 40/25/20/15
	accom := total * 40 / 100
	trans := total * 25 / 100
	food := total * 20 / 100
	act  := total - accom - trans - food

	// Adjust based on interests (very simple rules)
	interestSet := map[string]bool{}
	for _, v := range interests {
		interestSet[strings.ToLower(v)] = true
	}
	// Example: beach/wildlife → more activities; food → more food
	if interestSet["beach"] || interestSet["wildlife"] {
		move := total * 5 / 100
		if accom >= move {
			accom -= move
			act += move
		}
	}
	if interestSet["food"] {
		move := total * 5 / 100
		if trans >= move {
			trans -= move
			food += move
		}
	}
	return BudgetBreakdown{
		Accommodation: accom,
		Transport:     trans,
		Food:          food,
		Activities:    act,
		Total:         total,
	}
}

// BuildPlan generates a simple per-day itinerary from sampleActivities.
func BuildPlan(id string, req PlanRequest) (Plan, error) {
	start, end, err := validate(req)
	if err != nil {
		return Plan{}, err
	}

	// Days = inclusive date range
	daysCount := int(end.Sub(start).Hours()/24) + 1
	if daysCount < 1 {
		daysCount = 1
	}

	bud := AllocateBudget(req.BudgetKES, req.Interests)
	perDayFood := bud.Food / daysCount
	perDayAct := bud.Activities / daysCount

	days := make([]DayPlan, 0, daysCount)
	// Rotate through destinations across days
	for i := 0; i < daysCount; i++ {
		city := req.Destinations[i%len(req.Destinations)]
		acts := sampleActivities[city]
		if len(acts) == 0 {
			acts = []string{"Free exploration"}
		}
		plan := []string{}
		// pick up to 2 activities per day
		plan = append(plan, acts[i%len(acts)])
		if len(acts) > 1 {
			plan = append(plan, acts[(i+1)%len(acts)])
		}

		date := start.AddDate(0, 0, i).Format(dateLayout)
		days = append(days, DayPlan{
			Date:       date,
			City:       city,
			Plan:       plan,
			EstCostKES: perDayFood + perDayAct, // rough per-day spend (food+activities)
		})
	}

	nights := daysCount - 1
	if nights < 0 {
		nights = 0
	}
	return Plan{
		ID: id,
		Summary: PlanSummary{
			Nights: nights,
		},
		Days:   days,
		Notes:  []string{"This is a simple draft plan. Adjust activities and budget as you like."},
		Request: req,
		Budget:  bud,
	}, nil
}
