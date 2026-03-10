---
name: controller-breakdown
description: Explain Kubernetes controller code at a high level by identifying what resource the controller represents and which spec/status/metadata fields matter most. Use when reading large Go controller codebases, reviewing reconciliation logic, onboarding to a new operator, or when a user asks for a concise mental model of controller behavior.
---

# Controller Breakdown

## Overview

Produce a concise, high-signal breakdown of a Kubernetes controller.
Focus on two outcomes: what resource concept the controller manages, and which fields are operationally important.
Default to concise output, and switch to ultra-brief when user asks for "simple", "TL;DR", "3 sentences", or "5 lines".

## Workflow

1. Locate the primary API type(s)
- Find `type <Kind>Spec` and `type <Kind>Status`.
- Find where the controller watches and reconciles (`For(&...{})`, `Owns`, `Watches`).
- If there are multiple controllers, choose the one matching the user's scope, or state assumptions.

2. Identify what the resource represents
- Infer the real-world intent from reconcile actions, not from naming alone.
- Describe the resource in one sentence: "This resource represents X and drives Y."

3. Trace the reconcile contract
- Summarize desired inputs (`spec`, labels/annotations, referenced objects).
- Summarize produced outputs (`status`, child resources, events, conditions, finalizers).
- Note idempotency strategy and retry/error behavior if visible.

4. Extract relevant fields
- Prioritize fields that change behavior, rollout, ownership, lifecycle, or external integrations.
- Skip boilerplate unless it affects behavior.
- Group fields by `Spec`, `Status`, and `Metadata` (labels/annotations/finalizers).

5. Produce a compact explanation
- Use the output format below.
- Keep output short enough for onboarding docs, PR discussion, or handoff notes.
- Explicitly mark uncertainty when codepaths are incomplete.

## Brevity Knob

Select one mode before writing:

- `ultra-brief`: Use when asked for "TL;DR", "simpler", "3 sentences", "5 lines", or equivalent.
  - Start with `TL;DR`.
  - Maximum 5 lines total.
  - Maximum 3 sentences total.
  - Mention only: resource meaning, reconcile action, and top 1-3 fields.
- `standard`: Use for normal breakdown requests.
  - Include full sections below.

## Output Format

Use this structure:

### TL;DR
Always include this section first.
- In `ultra-brief` mode, this is the entire answer.
- In `standard` mode, keep it to 2-3 sentences.

### 1) Resource Represents
One sentence describing the domain entity and responsibility boundary.

### 2) Controller Intent
2-4 bullets describing the reconciliation loop at a high level:
- desired state inputs
- actions taken
- observed state feedback

### 3) Relevant Fields
List only high-value fields.

`Spec`
- `<field>`: why it matters operationally

`Status`
- `<field/condition>`: what signal it gives operators

`Metadata`
- `<label/annotation/finalizer>`: effect on lifecycle or ownership

### 4) Mental Model
One short paragraph on how to reason about this controller during incident response or feature changes.

## Guardrails

- Do not claim behavior not evidenced in code.
- Distinguish observed behavior from inferred intent.
- Prefer file/identifier references when possible (reconciler type, API type, key methods).
- If the codebase is huge, first narrow scope by API group/version/kind.

## References

Load [Controller Field Signals](references/controller_field_signals.md) when deciding which fields are high-value vs boilerplate.

## Example Prompt

"Use $controller-breakdown. TL;DR mode: explain this controller in 3 sentences max and 5 lines max."
