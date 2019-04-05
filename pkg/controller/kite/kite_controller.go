package kite

import (
	"context"

	appv1alpha1 "github.com/tufin/orca-operator/pkg/apis/app/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_kite")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Kite Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileKite{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("kite-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Kite
	err = c.Watch(&source.Kind{Type: &appv1alpha1.Kite{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Kite
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &appv1alpha1.Kite{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileKite{}

// ReconcileKite reconciles a Kite object
type ReconcileKite struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Kite object and makes changes based on the state read
// and what is in the Kite.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileKite) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Kite")

	// Fetch the Kite instance
	instance := &appv1alpha1.Kite{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	reqLogger.Info("TUFIN_GURU_URL",
		"Domain", instance.Spec.Domain,
		"Project", instance.Spec.Project,
		"KiteImage", instance.Spec.KiteImage,
		"EndPoints", instance.Spec.EndPoints,
		"IngnoredConfigMaps", instance.Spec.IngnoredConfigMaps,
		"Components", instance.Spec.Components,
		"KubePlatform", instance.Spec.KubePlatform,
	)
	deployment := getKiteDeployment(instance)
	daemonset := getConntrackDaemonset(instance)

	reconcileResult, err := r.createKiteDeployment(instance, deployment, request)
	if err != nil {
		return reconcileResult, err
	}
	reconcileResult, err = r.createConntrackDaemonset(instance, daemonset, request)
	return reconcileResult, err
}

func (r *ReconcileKite) createKiteDeployment(instance *appv1alpha1.Kite, deployment *appsv1.Deployment, request reconcile.Request) (reconcile.Result, error) {

	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Deploying Kite")

	// Set Kite instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, deployment, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if the Deployment already exists
	found := &appsv1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new deployment", "deployment.Namespace", deployment.Namespace, "deployment.Name", deployment.Name)
		err = r.client.Create(context.TODO(), deployment)
		if err != nil {
			return reconcile.Result{}, err
		}

		// deployment created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// deployment already exists - don't requeue
	reqLogger.Info("Skip reconcile: deployment already exists", "deployment.Namespace", found.Namespace, "deployment.Name", found.Name)
	return reconcile.Result{}, nil

}

func (r *ReconcileKite) createConntrackDaemonset(instance *appv1alpha1.Kite, daemonset *appsv1.DaemonSet, request reconcile.Request) (reconcile.Result, error) {

	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Deploying Conntrack")

	// Set Kite instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, daemonset, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if the Daemonset already exists
	found := &appsv1.DaemonSet{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: daemonset.Name, Namespace: daemonset.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new daemonset", "daemonset.Namespace", daemonset.Namespace, "daemonset.Name", daemonset.Name)
		err = r.client.Create(context.TODO(), daemonset)
		if err != nil {
			return reconcile.Result{}, err
		}

		// daemonset created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// daemonset already exists - don't requeue
	reqLogger.Info("Skip reconcile: daemonset already exists", "daemonset.Namespace", found.Namespace, "deployment.Name", found.Name)
	return reconcile.Result{}, nil

}
