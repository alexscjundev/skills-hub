---
name: k8s-architecture-breakdown
description: Explain large Kubernetes controller codebases with direct architecture summaries: CRD purpose, controller responsibilities, and common failure/recovery flows. Always produce a no-fluff markdown architecture document with a minimal diagram, keyed by project name and short report scope.
---

# K8s Architecture Breakdown

## Overview

Generate architecture output for Kubernetes codebases with zero filler.
Use short factual statements only.

## Modes

Pick the smallest mode that satisfies the request:

- `map`: components, boundaries, and responsibilities.
- `flow`: one scenario path through the system (for example failure/remediation/recovery).
- `diagram`: no-fluff diagram only.
- `full`: `map` + one or two `flow`s + `diagram`.

If user does not specify mode, default to `full`.

## Output Is Always a Document

Never return architecture content only in chat.
Always write a markdown file and return the path plus a 1-2 line summary in chat.

Path rules:
- If user provides a path, use it.
- Otherwise derive:
  - `project_name`: basename of repo root, kebab-case.
  - `scope_slug`: short scope label from request (for example `controller-runtime`, `nvsentinel-remediation`, `overview`), kebab-case, max 24 chars.
  - target path: `docs/architecture/<project_name>-<scope_slug>.md`

## Workflow

1. Scope the architecture surface
- Find primary API groups/kinds, reconcilers, and shared services.
- State scope assumptions when codebase is too large.
- If user asks for specific components only, treat that as strict scope and keep other pieces as `External` only.

2. Build component map
- Identify top components (controllers, APIs/CRDs, services/helpers, external dependencies).
- Assign one responsibility sentence per component.
- Highlight boundaries and ownership handoffs.

3. Trace runtime flow(s)
- For `flow` or `full`, select scenario(s): failure path and/or recovery path.
- Walk step-by-step from trigger -> decision -> action -> status signal.
- Mark branch points and retries.

4. Generate no-fluff diagram
- Use Mermaid by default unless user requests ASCII.
- Keep diagram minimal (prefer <= 8 nodes, <= 12 edges).
- Include only components in the chosen scope.

5. Ground claims with evidence
- Attach compact evidence tags to each major claim, e.g. `(pkg/foo/controller.go:88)`.
- If unsure, mark `Unknown` instead of guessing.

## Output Contract

### Document Header
- Must appear at the top of the markdown file before any other section.
- Format:
  - `# <ProjectName> Architecture - <ScopeLabel>`
  - `In scope: <comma-separated components/features>`
  - `Out of scope: <what is intentionally excluded, or "none">`
- `ScopeLabel` should be human-readable and short (for example `Slinky Drainer + Node Drain`).

### TL;DR
- Always include first.
- 2-4 lines only.
- Format:
  - `System: <one-sentence purpose>`
  - `Core loop: <who reads what and updates what>`
  - `Key flows: <failure path> | <recovery path>`

### 1) Architecture Map
- Use a markdown table, not paragraphs.
- Required columns: `Actor | Action | Reason / Details | Evidence`.
- `Actor` format: `<Type> <Name>`.
- Prefer types: `CRD`, `Controller`, `Service`, `Integration`, `Status`, `External`.
- Include only in-scope components unless needed for handoff; mark those as `External`.
- Cell limits:
  - `Action`: <= 14 words.
  - `Reason / Details`: <= 18 words.
  - `Evidence`: 1-2 references max.

### 2) Common Flow(s)
- One subsection per scenario.
- Each scenario must be a markdown table.
- Required columns: `Step | Object | Object Context (field/path + tiny snippet) | Actor | Action | Reason / Details | Evidence`.
- `Object` must name the state artifact moving through the flow (record/CR/status/queue item/annotation).
- `Object Context` must include a concrete field/path or tiny code-like snippet (for example `status.conditions[type=Ready]`, `spec.nodeName`, `Exists(crName)`).
- `Action` format should stay direct: `<reads X> -> <does Y> -> <writes Z>`.
- Keep each scenario to 4-8 steps.
- Cell limits:
  - `Action`: <= 14 words.
  - `Reason / Details`: <= 14 words, plain language.
  - `Evidence`: 1-2 references max.

### 3) Diagram
- Provide one no-fluff Mermaid diagram.
- Do not add styling unless requested.

### 4) Unknowns / Confidence
- One line: `Confidence: high|medium|low`.
- List missing evidence or unresolved paths.

## No-Fluff Rules

- Do not write introductions, conclusions, or motivational text.
- Do not use abstract wording like "robust", "seamless", "comprehensive".
- Keep each statement to one concrete behavior.
- Prefer `reads`, `decides`, `creates`, `updates`, `requeues`, `marks`.
- If a line has no concrete verb and object, rewrite it.
- If a line does not help identify scope, responsibility, flow, or handoff, delete it.
- Do not use prose lists for Architecture Map or Flow steps; use tables only.
- If a table row exceeds the cell limits, split it into multiple rows.
- In flow tables, do not omit `Object Context`.

## Guardrails

- Prefer simplification over exhaustive listing.
- Never invent components or flows.
- Keep language operational and concrete.
- Avoid generic architecture jargon without code evidence.

## References

Load [Architecture Patterns](references/architecture_patterns.md) when choosing flow templates and diagram scope.

## Example Prompts

"Use $k8s-architecture-breakdown in full mode. Focus on remediation and recovery flows. Write the document."

"Use $k8s-architecture-breakdown in diagram mode only. Mermaid, max 8 nodes."

"Use $k8s-architecture-breakdown for only these components: slinky-drainer controller and node-drainer reconciler. Exclude everything else unless needed as External handoff."

"Use $k8s-architecture-breakdown and format Architecture Map + Flows as tables with columns Actor, Action, Reason/Details, Evidence."
