package gcp

import "github.com/openshift/installer/pkg/destroy/providers"

func init() {
	providers.Registry["gcp"] = New
}
