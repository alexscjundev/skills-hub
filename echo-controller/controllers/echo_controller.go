package controllers

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	playgroundv1alpha1 "github.com/aljun/skills-playground/api/v1alpha1"
)

type EchoReconciler struct {
	client.Client
}

func (r *EchoReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var echo playgroundv1alpha1.Echo
	if err := r.Get(ctx, req.NamespacedName, &echo); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	desired := echo.Spec.Message
	if desired == "" {
		desired = "hello from controller"
	}

	if echo.Status.ObservedMessage == desired {
		return ctrl.Result{}, nil
	}

	echo.Status.ObservedMessage = desired
	echo.Status.LastReconciled = time.Now().UTC().Format(time.RFC3339)

	if err := r.Status().Update(ctx, &echo); err != nil {
		return ctrl.Result{}, err
	}

	logger.Info("reconciled Echo", "name", req.Name, "namespace", req.Namespace, "message", desired)
	return ctrl.Result{}, nil
}

func (r *EchoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&playgroundv1alpha1.Echo{}).
		Complete(r)
}
