package dynamichelper

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"strings"

	"github.com/ghodss/yaml"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type UnstructuredObj struct {
	obj unstructured.Unstructured
}

func NewUnstructuredObj(obj unstructured.Unstructured) *UnstructuredObj {
	return &UnstructuredObj{obj}
}

func (o UnstructuredObj) GetObjectKind() schema.ObjectKind {
	return o.obj.GetObjectKind()
}

func (o UnstructuredObj) DeepCopyObject() kruntime.Object {
	if un := o.obj.DeepCopy(); un != nil {
		out := NewUnstructuredObj(*un)
		return out
	}
	return nil
}

func (o UnstructuredObj) GroupVersionKind() schema.GroupVersionKind {
	return o.obj.GroupVersionKind()
}

func (o UnstructuredObj) GetNamespace() string {
	return o.obj.GetNamespace()
}
func (o UnstructuredObj) GetName() string {
	return o.obj.GetName()
}

func (o *UnstructuredObj) DecodeUnstructured(data []byte) error {
	json, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	err = o.obj.UnmarshalJSON(json)
	if err != nil {
		return err
	}
	return nil
}

func isKindUnstructured(groupKind string) bool {
	return strings.HasSuffix(groupKind, ".constraints.gatekeeper.sh")
}
