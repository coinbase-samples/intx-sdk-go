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

package test

import (
	"encoding/json"
	"fmt"
	"github.com/coinbase-samples/intx-sdk-go"
	"net/http"
	"os"
)

func newLiveTestClient() (*intx.Client, error) {

	credentials, err := loadCredentialsFromEnv()
	if err != nil {
		return nil, err
	}

	client := intx.NewClient(credentials, http.Client{})
	return client, nil

}

func loadCredentialsFromEnv() (*intx.Credentials, error) {

	credentials := &intx.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
		return nil, fmt.Errorf("unable to deserialize intx credentials JSON: %w", err)
	}

	return credentials, nil
}
