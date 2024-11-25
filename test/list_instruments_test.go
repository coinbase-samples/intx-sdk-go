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
	"context"
	"testing"
	"time"

	"github.com/coinbase-samples/intx-sdk-go/instruments"
)

func TestListInstruments(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := instruments.NewInstrumentsService(client)

	response, err := service.ListInstruments(ctx, &instruments.ListInstrumentsRequest{})
	if err != nil {
		t.Fatalf("ListInstruments returned an error: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response from ListInstruments")
	}

	if len(response.Instruments) == 0 {
		t.Fatal("expected at least one instrument in the response")
	}

	found := false
	for _, instrument := range response.Instruments {
		if instrument.Symbol == "ETH-PERP" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("expected to find ETH-PERP in the response")
	}
}
