# Controller Field Signals

Use this checklist to decide which fields are worth calling out in a high-level controller breakdown.

## Spec Signals (usually high-value)

- Lifecycle intent: `replicas`, `paused`, `suspend`, `deletionPolicy`, `cleanupPolicy`
- Rollout knobs: strategy type, surge/unavailable, partition/canary percentages
- Workload shape: image/version, resource requests/limits, topology/affinity
- Dependency wiring: references to secrets, configmaps, service accounts, classes, templates
- External integration: cloud/provider IDs, endpoint URLs, region/zone, credentials refs
- Ownership scope: selectors, namespace selectors, label matchers

## Status Signals (usually high-value)

- Conditions summarizing controller state (`Ready`, `Progressing`, `Degraded`, etc.)
- Observed generation synchronization (`observedGeneration`)
- Last applied/observed revision hashes or image versions
- Counters that indicate convergence (`readyReplicas`, `updatedReplicas`, `availableReplicas`)
- Failure surfaces (`lastError`, transient/permanent condition reasons)

## Metadata Signals (high-value when used by reconcile)

- Finalizers controlling delete semantics
- Controller-specific annotations that toggle behavior
- Labels used for selection, ownership, or policy targeting

## Usually Low-Value Boilerplate

- Embedded type/meta structs
- Generic timestamps unless used for logic
- Generated deepcopy registration details

## Quick Triage Questions

1. If this field changes, does reconcile behavior change?
2. Does this field explain why rollout converges or stalls?
3. Is this field essential for ownership/lifecycle safety?

If "yes" to any, include it in the relevant fields section.
