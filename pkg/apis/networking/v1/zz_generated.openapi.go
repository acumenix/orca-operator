// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/tufin/orca-operator/pkg/apis/networking/v1.Policy":       schema_pkg_apis_networking_v1_Policy(ref),
		"github.com/tufin/orca-operator/pkg/apis/networking/v1.PolicySpec":   schema_pkg_apis_networking_v1_PolicySpec(ref),
		"github.com/tufin/orca-operator/pkg/apis/networking/v1.PolicyStatus": schema_pkg_apis_networking_v1_PolicyStatus(ref),
	}
}

func schema_pkg_apis_networking_v1_Policy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Policy is the Schema for the policies API",
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
							Ref: ref("github.com/tufin/orca-operator/pkg/apis/networking/v1.PolicySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tufin/orca-operator/pkg/apis/networking/v1.PolicyStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tufin/orca-operator/pkg/apis/networking/v1.PolicySpec", "github.com/tufin/orca-operator/pkg/apis/networking/v1.PolicyStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_networking_v1_PolicySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PolicySpec defines the desired state of Policy",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_networking_v1_PolicyStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PolicyStatus defines the observed state of Policy",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}