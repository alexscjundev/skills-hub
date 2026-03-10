# Dirt Simple Kubernetes Controller

This repo contains a tiny custom controller that reconciles an `Echo` custom resource.

## What it does

- Watches `Echo` objects (`playground.aljun.dev/v1alpha1`).
- Copies `spec.message` into `status.observedMessage`.
- Stamps `status.lastReconciled` with an RFC3339 UTC timestamp.

## Quick start

Prereqs:

- A reachable Kubernetes cluster (`kubectl` context configured).
- Go 1.25+.

Run:

```bash
make install-crd
make run
```

In another shell:

```bash
kubectl apply -f config/samples/playground_v1alpha1_echo.yaml
kubectl get echo demo-echo -o yaml
```

Clean up:

```bash
make uninstall-crd
```
