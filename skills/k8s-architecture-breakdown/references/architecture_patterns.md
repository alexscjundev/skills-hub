# Architecture Patterns

Use this file to keep outputs simple, scenario-oriented, and evidence-backed.

## Component Categories

- API Surface: CRDs, API types (`Spec`, `Status`), validation/defaulting hooks
- Control Loop: reconcilers, predicates, work queue behavior, finalizer handlers
- Domain Services: planner/executor/remediator modules called by reconcile
- Integrations: Kubernetes objects, cloud APIs, node/system interfaces, webhooks
- Signals: status conditions, events, metrics/logs that operators observe

## Separation of Concerns Checklist

For each component, answer:

1. What decision does it own?
2. What side effects can it perform?
3. What inputs does it trust?
4. What output or signal does it emit?

If the answer is unclear, mark `Unknown`.

## Flow Templates

### Failure / Remediation

1. Trigger detected (condition/event/health probe failure)
2. Classify failure and choose remediation plan
3. Execute remediation action(s)
4. Update status/condition and emit events
5. Requeue/retry behavior until convergence or terminal state

### Recovery / Healthy Again

1. Healthy signal detected
2. Clear or downgrade failure condition
3. Reconcile desired state and remove temporary remediation state
4. Publish healthy condition/status
5. Return to steady-state reconcile loop

## Diagram Rules

- Prefer one directional flow (`A --> B`) over dense graph layouts
- Keep to core control/data paths only
- Avoid edge labels unless needed for branch meaning
- Limit to <= 8 nodes, <= 12 edges unless user asks for more

## Evidence Tag Format

Use compact tags like:

- `(internal/foo/controller.go:88)`
- `(api/v1alpha1/foo_types.go:24)`

Attach at least one evidence tag for each major component and each flow step.
