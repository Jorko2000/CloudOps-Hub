package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GroupVersion = schema.GroupVersion{
	Group:   "cloudops.io",
	Version: "v1",
}
