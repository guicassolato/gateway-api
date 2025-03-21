/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1 "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/typed/apis/v1"
)

type FakeGatewayV1 struct {
	*testing.Fake
}

func (c *FakeGatewayV1) GRPCRoutes(namespace string) v1.GRPCRouteInterface {
	return newFakeGRPCRoutes(c, namespace)
}

func (c *FakeGatewayV1) Gateways(namespace string) v1.GatewayInterface {
	return newFakeGateways(c, namespace)
}

func (c *FakeGatewayV1) GatewayClasses() v1.GatewayClassInterface {
	return newFakeGatewayClasses(c)
}

func (c *FakeGatewayV1) HTTPRoutes(namespace string) v1.HTTPRouteInterface {
	return newFakeHTTPRoutes(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeGatewayV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
