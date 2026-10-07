# greeter — PRD

## Problem Statement

Teams building and validating services on this platform need a minimal, known-good reference service to exercise end-to-end conventions (project layout, build, API contract, deployment) without the complexity of a real product. Today there is no small, canonical HTTP service that demonstrates these conventions end to end.

## Solution

Greeter is a small Go HTTP service that exposes a single endpoint, `GET /hello?name=X`, returning a JSON greeting for the given name. It exists to follow and demonstrate the conventions set out in `app-factory-kaj/e2e-reference`.

## Actors

- **API Consumer** — another service, script, or developer that calls the greeter API programmatically to obtain a JSON greeting. *assumed*

## User Stories

1. As an API Consumer, I want to call `GET /hello?name=X` and receive a JSON greeting that includes the given name, so that I can verify the service responds correctly.
2. As an API Consumer, I want to call `GET /hello` without a `name` parameter and still receive a generic JSON greeting, so that the endpoint behaves predictably even with incomplete input.

## Product Decisions

- Missing or empty `name` query parameter: the service returns a generic greeting (e.g. "Hello, World!") rather than an error. *assumed*
- Authentication: the endpoint is open and unauthenticated — no sign-in, API key, or token required, consistent with a minimal reference service. *assumed*
- The service follows the conventions in `app-factory-kaj/e2e-reference` for project structure and API design.

## Out of Scope

- Any persistence, storage, or state — the service is stateless.
- A user interface or frontend of any kind.
- Authentication, authorization, or rate limiting.
- Any endpoint other than `GET /hello`.

## Open Questions

*(none)*

## Further Notes

*(none)*