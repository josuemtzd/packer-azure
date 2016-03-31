package resources

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

// ProviderOperationDetailsClient is the client for the
// ProviderOperationDetails methods of the Resources service.
type ProviderOperationDetailsClient struct {
	ManagementClient
}

// NewProviderOperationDetailsClient creates an instance of the
// ProviderOperationDetailsClient client.
func NewProviderOperationDetailsClient(subscriptionID string) ProviderOperationDetailsClient {
	return NewProviderOperationDetailsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderOperationDetailsClientWithBaseURI creates an instance of the
// ProviderOperationDetailsClient client.
func NewProviderOperationDetailsClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationDetailsClient {
	return ProviderOperationDetailsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets a list of resource providers.
//
// resourceProviderNamespace is resource identity.
func (client ProviderOperationDetailsClient) List(resourceProviderNamespace string, apiVersion string) (result ResourceProviderOperationDetailListResult, err error) {
	req, err := client.ListPreparer(resourceProviderNamespace, apiVersion)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ProviderOperationDetailsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ProviderOperationDetailsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/ProviderOperationDetailsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ProviderOperationDetailsClient) ListPreparer(resourceProviderNamespace string, apiVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/{resourceProviderNamespace}/operations"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderOperationDetailsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProviderOperationDetailsClient) ListResponder(resp *http.Response) (result ResourceProviderOperationDetailListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ProviderOperationDetailsClient) ListNextResults(lastResults ResourceProviderOperationDetailListResult) (result ResourceProviderOperationDetailListResult, err error) {
	req, err := lastResults.ResourceProviderOperationDetailListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ProviderOperationDetailsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ProviderOperationDetailsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/ProviderOperationDetailsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}
