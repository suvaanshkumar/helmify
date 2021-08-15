package rbac

import (
	"github.com/arttor/helmify/internal"
	"github.com/arttor/helmify/pkg/helmify"
	"github.com/stretchr/testify/assert"
	"testing"
)

const clusterRoleBindingYaml = `apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: my-operator-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: my-operator-manager-role
subjects:
- kind: ServiceAccount
  name: my-operator-controller-manager
  namespace: my-operator-system`

func Test_clusterRoleBinding_Process(t *testing.T) {
	var testInstance clusterRoleBinding

	t.Run("processed", func(t *testing.T) {
		obj := internal.GenerateObj(clusterRoleBindingYaml)
		processed, _, err := testInstance.Process(helmify.ChartInfo{}, obj)
		assert.NoError(t, err)
		assert.Equal(t, true, processed)
	})
	t.Run("skipped", func(t *testing.T) {
		obj := internal.TestNs
		processed, _, err := testInstance.Process(helmify.ChartInfo{}, obj)
		assert.NoError(t, err)
		assert.Equal(t, false, processed)
	})
}
