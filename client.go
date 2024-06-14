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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var defaultV1ApiBaseUrl = "https://api.international.coinbase.com/api/v1"

type Client struct {
	httpClient  http.Client
	httpBaseUrl string
	Credentials *Credentials
}

func (c *Client) HttpBaseUrl() string {
	return c.httpBaseUrl
}

func (c *Client) HttpClient() *http.Client {
	return &c.httpClient
}

func (c *Client) SetBaseUrl(u string) *Client {
	c.httpBaseUrl = u
	return c
}

func NewClient(credentials *Credentials, httpClient http.Client) *Client {
	httpBaseUrl := os.Getenv("INTX_BASE_URL")
	if httpBaseUrl == "" {
		httpBaseUrl = defaultV1ApiBaseUrl
	}
	return &Client{
		httpBaseUrl: httpBaseUrl,
		Credentials: credentials,
		httpClient:  httpClient,
	}
}

func addIntxHeaders(req *http.Request, path string, body []byte, client core.Client, t time.Time) {
	c := client.(*Client)
	signature := sign(req.Method, path, t.Unix(), c.Credentials.SigningKey, body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("CB-ACCESS-KEY", c.Credentials.AccessKey)
	req.Header.Add("CB-ACCESS-PASSPHRASE", c.Credentials.Passphrase)
	req.Header.Add("CB-ACCESS-SIGN", signature)
	req.Header.Add("CB-ACCESS-TIMESTAMP", strconv.FormatInt(t.Unix(), 10))
}

func sign(method, path string, t int64, signingKey string, body []byte) string {
	key, err := base64.StdEncoding.DecodeString(signingKey)
	if err != nil {
		log.Fatalf("Error decoding signing key: %v", err)
	}

	message := fmt.Sprintf("%d%s%s", t, method, path)
	if len(body) > 0 {
		message += string(body)
	}

	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
