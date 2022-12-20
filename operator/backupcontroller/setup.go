package backupcontroller

import (
	"github.com/go-logr/logr"
	"github.com/k8up-io/k8up/v2/api/v1"
	"sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// +kubebuilder:rbac:groups=k8up.io,resources=backups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=k8up.io,resources=backups/status;backups/finalizers,verbs=get;update;patch
// +kubebuilder:rbac:groups=k8up.io,resources=prebackuppods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=k8up.io,resources=prebackuppods/status;prebackuppods/finalizers,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs="*"
// +kubebuilder:rbac:groups=core,resources=pods/exec,verbs="*"
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;create;delete
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles;rolebindings,verbs=get;list;create;delete

// SetupWithManager configures the reconciler.
func (r *BackupReconciler) SetupWithManager(mgr controllerruntime.Manager, _ logr.Logger) error {
	name := "backup.k8up.io"
	r.Kube = mgr.GetClient()
	return controllerruntime.NewControllerManagedBy(mgr).
		Named(name).
		For(&v1.Backup{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}