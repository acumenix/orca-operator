// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/tufin/orca-operator/pkg/apis/app/v1alpha1.Kite":       schema_pkg_apis_app_v1alpha1_Kite(ref),
		"github.com/tufin/orca-operator/pkg/apis/app/v1alpha1.KiteSpec":   schema_pkg_apis_app_v1alpha1_KiteSpec(ref),
		"github.com/tufin/orca-operator/pkg/apis/app/v1alpha1.KiteStatus": schema_pkg_apis_app_v1alpha1_KiteStatus(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_Kite(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Kite is the Schema for the kites API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tufin/orca-operator/pkg/apis/app/v1alpha1.KiteSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tufin/orca-operator/pkg/apis/app/v1alpha1.KiteStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tufin/orca-operator/pkg/apis/app/v1alpha1.KiteSpec", "github.com/tufin/orca-operator/pkg/apis/app/v1alpha1.KiteStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_KiteSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KiteSpec defines the desired state of Kite",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1alpha1_KiteStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KiteStatus defines the observed state of Kite",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
