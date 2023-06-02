package test

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
)

// An example of how to test the Kubernetes resource config in examples/kubernetes-basic-example using Terratest.
func TestKubernetesBasicExample(t *testing.T) {
	t.Parallel()

	// website::tag::1::Path to the Kubernetes resource config we will test
	kubeResourcePath, err := filepath.Abs("../deployment/nginx-deployment.yaml")
	require.NoError(t, err)

	if err != nil {
		log.Fatal(err)
	}

	// To ensure we can reuse the resource config on the same cluster to test different scenarios, we setup a unique
	// namespace for the resources for this test.
	// Note that namespaces must be lowercase.
	namespaceName := fmt.Sprintf("kubernetes-basic-example-%s", strings.ToLower(random.UniqueId()))

	// - Random namespace
	options := k8s.NewKubectlOptions("", "", namespaceName)

	k8s.CreateNamespace(t, options, namespaceName)

	// website::tag::5::Make sure to delete the namespace at the end of the test
	defer k8s.DeleteNamespace(t, options, namespaceName)

	defer k8s.KubectlDelete(t, options, kubeResourcePath)

	// This will run `kubectl apply -f RESOURCE_CONFIG` and fail the test if there are any errors
	k8s.KubectlApply(t, options, kubeResourcePath)

	fmt.Println("\nKubectl apply done")

	time.Sleep(30 * time.Second)

	fmt.Println("\nAfter sleep")

	// This will get the service resource and verify that it exists and was retrieved successfully. This function will
	// fail the test if the there is an error retrieving the service resource from Kubernetes.
	service := k8s.GetService(t, options, "nginx-service-test")
	require.Equal(t, service.Name, "nginx-service-test")
}
