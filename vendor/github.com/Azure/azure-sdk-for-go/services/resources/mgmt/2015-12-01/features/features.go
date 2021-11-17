package features

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// Client is the azure Feature Exposure Control (AFEC) provides a mechanism for the resource providers to control
// feature exposure to users. Resource providers typically use this mechanism to provide public/private preview for new
// features prior to making them generally available. Users need to explicitly register for AFEC features to get access
// to such functionality.
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient(subscriptionID string) Client {
	return NewClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return Client{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the preview feature with the specified name.
// Parameters:
// resourceProviderNamespace - the resource provider namespace for the feature.
// featureName - the name of the feature to get.
func (client Client) Get(ctx context.Context, resourceProviderNamespace string, featureName string) (result Result, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceProviderNamespace, featureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "features.Client", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client Client) GetPreparer(ctx context.Context, resourceProviderNamespace string, featureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName":               autorest.Encode("path", featureName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result Result, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the preview features in a provider namespace that are available through AFEC for the subscription.
// Parameters:
// resourceProviderNamespace - the namespace of the resource provider for getting features.
func (client Client) List(ctx context.Context, resourceProviderNamespace string) (result OperationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.olr.Response.Response != nil {
				sc = result.olr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceProviderNamespace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.olr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "features.Client", "List", resp, "Failure sending request")
		return
	}

	result.olr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "List", resp, "Failure responding to request")
		return
	}
	if result.olr.hasNextLink() && result.olr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(ctx context.Context, resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result OperationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client Client) listNextResults(ctx context.Context, lastResults OperationsListResult) (result OperationsListResult, err error) {
	req, err := lastResults.operationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.Client", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.Client", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListComplete(ctx context.Context, resourceProviderNamespace string) (result OperationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceProviderNamespace)
	return
}

// ListAll gets all the preview features that are available through AFEC for the subscription.
func (client Client) ListAll(ctx context.Context) (result OperationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListAll")
		defer func() {
			sc := -1
			if result.olr.Response.Response != nil {
				sc = result.olr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.olr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "features.Client", "ListAll", resp, "Failure sending request")
		return
	}

	result.olr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "ListAll", resp, "Failure responding to request")
		return
	}
	if result.olr.hasNextLink() && result.olr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client Client) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/features", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client Client) ListAllResponder(resp *http.Response) (result OperationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client Client) listAllNextResults(ctx context.Context, lastResults OperationsListResult) (result OperationsListResult, err error) {
	req, err := lastResults.operationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.Client", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.Client", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListAllComplete(ctx context.Context) (result OperationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// Register registers the preview feature for the subscription.
// Parameters:
// resourceProviderNamespace - the namespace of the resource provider.
// featureName - the name of the feature to register.
func (client Client) Register(ctx context.Context, resourceProviderNamespace string, featureName string) (result Result, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Register")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegisterPreparer(ctx, resourceProviderNamespace, featureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "Register", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegisterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "features.Client", "Register", resp, "Failure sending request")
		return
	}

	result, err = client.RegisterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "Register", resp, "Failure responding to request")
		return
	}

	return
}

// RegisterPreparer prepares the Register request.
func (client Client) RegisterPreparer(ctx context.Context, resourceProviderNamespace string, featureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName":               autorest.Encode("path", featureName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}/register", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegisterSender sends the Register request. The method will close the
// http.Response Body if it receives an error.
func (client Client) RegisterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegisterResponder handles the response to the Register request. The method always
// closes the http.Response Body.
func (client Client) RegisterResponder(resp *http.Response) (result Result, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Unregister unregisters the preview feature for the subscription.
// Parameters:
// resourceProviderNamespace - the namespace of the resource provider.
// featureName - the name of the feature to unregister.
func (client Client) Unregister(ctx context.Context, resourceProviderNamespace string, featureName string) (result Result, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Unregister")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UnregisterPreparer(ctx, resourceProviderNamespace, featureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "Unregister", nil, "Failure preparing request")
		return
	}

	resp, err := client.UnregisterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "features.Client", "Unregister", resp, "Failure sending request")
		return
	}

	result, err = client.UnregisterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "Unregister", resp, "Failure responding to request")
		return
	}

	return
}

// UnregisterPreparer prepares the Unregister request.
func (client Client) UnregisterPreparer(ctx context.Context, resourceProviderNamespace string, featureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName":               autorest.Encode("path", featureName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}/unregister", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UnregisterSender sends the Unregister request. The method will close the
// http.Response Body if it receives an error.
func (client Client) UnregisterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UnregisterResponder handles the response to the Unregister request. The method always
// closes the http.Response Body.
func (client Client) UnregisterResponder(resp *http.Response) (result Result, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
