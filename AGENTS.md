# AGENTS.md

Execution rules for agents working in this repository.

## Goal
Build a learning-focused TCP Reverse Proxy in Go, incrementally.
- One entrypoint that distributes to multiple backends
- Focus on low-to-mid level concerns (connection management, timeouts, backpressure, etc.)
- Keep design extensible for observability (logs/metrics/tracing)

## Ownership Split (Most Important)
- **Theme area (implemented from scratch by the user)**
  - Core proxy data plane
    - Accept / Dial
    - Bidirectional forwarding
    - Load-balancing strategy
    - Control behaviors (timeouts / retries / circuit breaker, etc.)
    - Connection management and congestion/backpressure behavior
  - As a rule, the user writes core logic from zero.

- **Non-theme area (primarily implemented by the agent)**
  - Project plumbing: CLI wiring, config loader, sample config, Makefile, Docker, CI skeleton
  - Observability scaffolding: logging foundation, metrics endpoint skeleton
  - Test infrastructure: test helpers, fixtures, benchmark harness, fault-injection scripts
  - Documentation: README, operation notes, validation steps

## Implementation Policy
1. **Do not overstep into the theme area**
   - Prefer interfaces, TODOs, stubs, and test viewpoints over directly implementing core logic.
2. **Ship in small working units**
   - Keep each PR-sized change focused on one objective.
3. **Prioritize measurability**
   - Each change should include at least one verification method (logs, metrics, or benchmark).
4. **Avoid destructive changes**
   - Large renames/restructures require explicit user agreement.

## Code Style
- Follow standard Go tools (`gofmt`, `go test`).
- Keep dependencies minimal; explain why when adding one.
- Write error messages that help identify root cause.

## Task List Operations
- Do not use schedules. `ROADMAP.md` must stay a **task list**.
- Use these task states:
  - `[ ]` Not started
  - `[-]` In progress
  - `[x]` Done

## Motivation Visibility
- Leave visible outcomes for progress:
  - Benchmark diffs
  - Screenshots (metrics/logs)
  - "Break and fix" records
