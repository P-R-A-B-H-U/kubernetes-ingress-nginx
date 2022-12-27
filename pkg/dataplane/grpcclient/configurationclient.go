/*
Copyright 2022 The Kubernetes Authors.

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

package grpcclient

import (
	"fmt"

	json "github.com/json-iterator/go"
	"google.golang.org/grpc"

	"k8s.io/ingress-nginx/internal/ingress/controller/config"
	"k8s.io/ingress-nginx/pkg/apis/ingress"
	"k8s.io/klog/v2"
)

func (c *Client) ConfigurationService() {
	stream, err := c.ConfigurationClient.WatchConfigurations(c.ctx, c.Backendname, grpc.WaitForReady(true))
	if err != nil {
		c.grpcErrCh <- fmt.Errorf("error creating configuration client: %w", err)
		return
	}

	for {
		cfg, err := stream.Recv()
		if err != nil {
			c.grpcErrCh <- fmt.Errorf("error getting configuration: %w", err)
			return
		}
		var configtemplate *config.TemplateConfig

		switch op := cfg.Op.(type) {
		case *ingress.Configurations_FullconfigOp:
			if err := json.Unmarshal(op.FullconfigOp.Configuration, &configtemplate); err != nil {
				klog.Errorf("error unmarshalling config: %s", err)
				continue
			}
		default:
			klog.Warningf("Operation not implemented: %+v", op)
			continue
		}
		c.ConfigCh <- configtemplate
	}
}
