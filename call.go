/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package intx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const emptyQueryParams = ""

type apiRequest struct {
	path                   string
	query                  string
	httpMethod             string
	body                   []byte
	expectedHttpStatusCode int
	client                 Client
}

type apiResponse struct {
	request        *apiRequest
	body           []byte
	httpStatusCode int
	httpStatusMsg  string
	err            error
	errorMessage   *ErrorMessage
}

func post(
	ctx context.Context,
	client Client,
	path,
	query string,
	request,
	response interface{},
) error {
	return call(ctx, client, path, query, http.MethodPost, http.StatusOK, request, response)
}

func get(
	ctx context.Context,
	client Client,
	path,
	query string,
	request,
	response interface{},
) error {
	return call(ctx, client, path, query, http.MethodGet, http.StatusOK, request, response)
}

func put(
	ctx context.Context,
	client Client,
	path,
	query string,
	request,
	response interface{},
) error {
	return call(ctx, client, path, query, http.MethodPut, http.StatusOK, request, response)
}

func del(
	ctx context.Context,
	client Client,
	path,
	query string,
	request,
	response interface{},
) error {
	return call(ctx, client, path, query, http.MethodDelete, http.StatusOK, request, response)
}

func call(
	ctx context.Context,
	client Client,
	path,
	query,
	httpMethod string,
	expectedHttpStatusCode int,
	request,
	response interface{},
) error {

	if client.Credentials == nil {
		return errors.New("credentials not set")
	}

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	resp := makeCall(
		ctx,
		&apiRequest{
			path:                   path,
			query:                  query,
			httpMethod:             httpMethod,
			body:                   body,
			expectedHttpStatusCode: expectedHttpStatusCode,
			client:                 client,
		},
	)

	if resp.err != nil {
		return resp.err
	}

	if err := json.Unmarshal(resp.body, response); err != nil {
		return err
	}

	return nil
}

func makeCall(ctx context.Context, request *apiRequest) *apiResponse {

	response := &apiResponse{
		request: request,
	}

	callUrl := fmt.Sprintf("%s%s%s", request.client.HttpBaseUrl, request.path, request.query)

	parsedUrl, err := url.Parse(callUrl)
	if err != nil {
		response.err = fmt.Errorf("invalid URL: %s - %w", callUrl, err)
		return response
	}

	var requestBody []byte
	if request.httpMethod == http.MethodPost || request.httpMethod == http.MethodPut {
		requestBody = request.body
	}

	t := time.Now().Unix()

	req, err := http.NewRequestWithContext(ctx, request.httpMethod, callUrl, bytes.NewReader(requestBody))
	if err != nil {
		response.err = err
		return response
	}

	signature := sign(request.httpMethod, parsedUrl.Path, t, request.client.Credentials.SigningKey, requestBody)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("CB-ACCESS-KEY", request.client.Credentials.AccessKey)
	req.Header.Add("CB-ACCESS-PASSPHRASE", request.client.Credentials.Passphrase)
	req.Header.Add("CB-ACCESS-SIGN", signature)
	req.Header.Add("CB-ACCESS-TIMESTAMP", fmt.Sprintf("%d", t))

	res, err := request.client.HttpClient.Do(req)
	if err != nil {
		response.err = err
		return response
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		response.err = err
		return response
	}

	response.body = body
	response.httpStatusCode = res.StatusCode
	response.httpStatusMsg = res.Status

	if request.expectedHttpStatusCode > 0 && res.StatusCode != request.expectedHttpStatusCode {

		var errMsg ErrorMessage
		if strings.Contains(string(response.body), "message") {
			_ = json.Unmarshal(response.body, &errMsg)
			response.errorMessage = &errMsg
		}

		responseMsg := string(body)

		if response.errorMessage != nil && len(response.errorMessage.Value) > 0 {
			responseMsg = response.errorMessage.Value
		}

		response.err = fmt.Errorf(
			"expected status code: %d - received: %d - status msg: %s - url %s - msg: %s",
			request.expectedHttpStatusCode,
			res.StatusCode,
			res.Status,
			callUrl,
			responseMsg,
		)
	}

	return response
}
