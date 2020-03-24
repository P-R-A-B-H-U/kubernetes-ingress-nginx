/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"path/filepath"
	"testing"

	"k8s.io/ingress-nginx/internal/nginx"
)

func TestCreateApiserverClient(t *testing.T) {
	_, _, err := createApiserverClient("", "", "")
	if err == nil {
		t.Fatal("Expected an error creating REST client without an API server URL or kubeconfig file.")
	}
}

func init() {
	// the default value of nginx.TemplatePath assumes the template exists in
	// the root filesystem and not in the rootfs directory
	path, err := filepath.Abs(filepath.Join("../../rootfs/", nginx.TemplatePath))
	if err == nil {
		nginx.TemplatePath = path
	}
}
