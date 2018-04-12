package networkpolicy

import (
	"testing"

	"github.com/manifoldco/heighliner/pkg/api/v1alpha1"
)

func TestServiceForRelease(t *testing.T) {
	release := &v1alpha1.Release{
		SemVer: &v1alpha1.SemVerRelease{
			Name:    "test-application",
			Version: "1.2.3",
			Build:   "201804121722",
		},
	}

	t.Run("without ports", func(t *testing.T) {
		np := &v1alpha1.NetworkPolicy{
			Spec: v1alpha1.NetworkPolicySpec{},
		}

		obj, err := serviceForRelease(np, release)
		if obj != nil {
			t.Errorf("Expected object to be nil, got %#v", obj)
		}

		if err != nil {
			t.Errorf("Expected error to be nil, got %s", err)
		}
	})

	t.Run("with a set of ports", func(t *testing.T) {
		np := &v1alpha1.NetworkPolicy{
			Spec: v1alpha1.NetworkPolicySpec{
				Ports: []v1alpha1.NetworkPort{
					{
						Name:       "headless",
						TargetPort: 8080,
						Port:       80,
					},
				},
			},
		}

		obj, err := serviceForRelease(np, release)
		if obj == nil {
			t.Errorf("Expected object to be nil, got %#v", obj)
		}

		if err != nil {
			t.Errorf("Expected error to be nil, got %s", err)
		}
	})
}
