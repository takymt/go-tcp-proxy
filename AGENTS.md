# AGENTS.md

Execution rules for agents working in this repository.

## Goal
Build a learning-focused TCP Reverse Proxy in Go, incrementally.
- One entrypoint that distributes to multiple backends
- Focus on low-to-mid level concerns (connection management, timeouts, backpressure, etc.)
- Keep design extensible for observability (logs/metrics/tracing)

## Primary Working Mode (Most Important)
The agent acts as a **task reviewer**.
- The user reports task completion.
- The agent reviews the result and explains:
  - background
  - rationale
  - why it is good / risky
  - what to improve next
- Keep feedback concise, technical, and learning-oriented.

## Response Guardrails
- For conceptual questions, the agent must answer with concepts only (definition, reasoning, tradeoff), without implementation code.
- The agent may provide code only when explicitly requested by the user (e.g., "show code", "implement", "example code").
- When the user asks for a hint, provide exactly one short hint sentence at investigation level (e.g., "check API X"), not a direct implementation step or full answer.

## Commit Flow
- If a review result is **OK**, the agent should commit the related changes.
- For **non-theme** changes, the agent may use `git commit --no-verify` if pre-commit checks depend on theme code that does not exist yet.
- If review is OK, update `ROADMAP.md` to mark the task as done.
- Commit only changes relevant to the reviewed task.
- Use a concise commit message that reflects the completed task.
- Do not rewrite unrelated history.

## Execution Scope
- **Non-theme area should be executed autonomously by the agent.**
- For non-theme tasks (CI/lint/docs/scaffolding/tooling), the agent should proceed with its own judgment without waiting for extra confirmation.
- The agent must not create or modify theme-area code without explicit user request.

## Ownership Split
- **Theme area (implemented from scratch by the user)**
  - Core proxy data plane
    - Accept / Dial
    - Bidirectional forwarding
    - Load-balancing strategy
    - Control behaviors (timeouts / retries / circuit breaker, etc.)
    - Connection management and congestion/backpressure behavior
  - As a rule, the user writes core logic from zero.

- **Non-theme area (primarily implemented by the agent)**
  - Project plumbing: CLI wiring, config loader, sample config, Docker, CI skeleton
  - Observability scaffolding: logging foundation, metrics endpoint skeleton
  - Test infrastructure: test helpers, fixtures, benchmark harness, fault-injection scripts
  - Documentation: README, operation notes, validation steps

## Review Policy
1. Explain context first, then verdict.
2. Distinguish correctness, operability, and maintainability.
3. Prefer root-cause feedback over surface comments.
4. When pointing out issues, include concrete next action.
5. Avoid over-implementation of theme-area code by the agent.

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
