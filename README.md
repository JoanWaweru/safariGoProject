# SafariGo ✈️🦁

A tiny Go app that turns simple trip info into a draft **itinerary** and **budget**. Perfect for learning Go by building something real.

## Why this exists

If you’re new to Go, big projects can feel scary. SafariGo is small on purpose: clean code, clear structure, and a couple of endpoints so you can practice the essentials without getting lost.

## What you’ll practice

* Building a **JSON HTTP API** with `net/http`
* **Structs**, validation, and simple business logic
* Writing and running **tests**
* Keeping code tidy with **formatting** and **vetting**

---

**Prereqs**

* Go **1.24+** installed (`go version`)

**Run**

```bash
go run .
```

The server starts on `:8080`.

**Test it quickly**

```bash
# Health
curl -i http://localhost:8080/healthz

# Create a plan
curl -s -X POST http://localhost:8080/plans \
  -H "Content-Type: application/json" \
  -d '{
    "origin": "Nairobi",
    "destinations": ["Diani", "Mombasa"],
    "start_date": "2025-09-01",
    "end_date": "2025-09-03",
    "budget_kes": 100000,
    "interests": ["beach", "food"]
  }' | jq
```

You’ll get a JSON plan with an **id**, **budget split**, and a per-day **itinerary**.

---

## Endpoints

### `GET /healthz`

Returns:

```json
{ "status": "ok" }
```

### `POST /plans`

Send:

```json
{
  "origin": "Nairobi",
  "destinations": ["Diani", "Mombasa"],
  "start_date": "YYYY-MM-DD",
  "end_date": "YYYY-MM-DD",
  "budget_kes": 100000,
  "interests": ["beach", "food"]
}
```

Returns (shortened):

```json
{
  "id": "20250901T120102.123456789",
  "summary": { "nights": 2 },
  "budget_split": {
    "accommodation": 40000,
    "transport": 25000,
    "food": 20000,
    "activities": 15000,
    "total": 100000
  },
  "itinerary": [
    { "date": "2025-09-01", "city": "Diani", "plan": ["Beach afternoon", "Snorkeling"], "est_cost_kes": 14000 },
    { "date": "2025-09-02", "city": "Mombasa", "plan": ["Old Town walk", "Fort Jesus visit"], "est_cost_kes": 14000 }
  ],
  "notes": ["This is a simple draft plan. Adjust activities and budget as you like."]
}
```

> **Budget logic (simple & tweakable):**
> Base split = 40% accommodation, 25% transport, 20% food, 15% activities.
> If interests include **beach** or **wildlife**, shift +5% to activities.
> If interests include **food**, shift +5% to food (from transport).

---

## How the code is laid out

```
safarigo/
  internal/
    models.go      # Request/response types and small structs
    planner.go     # Validation, budget split, itinerary builder
    planner_test.go# Your first table-driven test
  main.go          # HTTP server & routes
  go.mod
```

---

## Develop

**Run tests**

```bash
go test ./...
go test ./... -v -race
go test ./... -cover
```

**Keep it tidy**

```bash
go fmt ./...
go vet ./...
```

---

## Troubleshooting (quick)

* **“command not found: go”** → open a new terminal or reinstall Go.
* **Apple silicon weirdness** → ensure Terminal isn’t using Rosetta; `uname -m` should say `arm64`.
* **Port already in use** → change `addr := ":8081"` in `main.go`.

---

## Roadmap (what’s next)

* Better input validation + friendlier errors
* Concurrency with `context` (simulate parallel lookups)
* Persistence (SQLite) + `GET /plans/{id}`
* Optional CLI to export Markdown/CSV
* Dockerfile and a tiny README deploy guide
