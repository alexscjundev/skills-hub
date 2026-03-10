GO ?= go

.PHONY: run test install-crd uninstall-crd

run:
	$(GO) run ./cmd/controller

test:
	$(GO) test ./...

install-crd:
	kubectl apply -f config/crd/bases/playground.aljun.dev_echoes.yaml

uninstall-crd:
	kubectl delete -f config/crd/bases/playground.aljun.dev_echoes.yaml --ignore-not-found
