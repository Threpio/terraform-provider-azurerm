package outputs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TestOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Test ...
func (c OutputsClient) Test(ctx context.Context, id OutputId, input Output) (result TestOperationResponse, err error) {
	req, err := c.preparerForTest(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outputs.OutputsClient", "Test", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForTest(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outputs.OutputsClient", "Test", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// TestThenPoll performs Test then polls until it's completed
func (c OutputsClient) TestThenPoll(ctx context.Context, id OutputId, input Output) error {
	result, err := c.Test(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Test: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Test: %+v", err)
	}

	return nil
}

// preparerForTest prepares the Test request.
func (c OutputsClient) preparerForTest(ctx context.Context, id OutputId, input Output) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/test", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForTest sends the Test request. The method will close the
// http.Response Body if it receives an error.
func (c OutputsClient) senderForTest(ctx context.Context, req *http.Request) (future TestOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
