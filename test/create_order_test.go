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

	"github.com/coinbase-samples/intx-sdk-go/orders"
	"github.com/google/uuid"
)

func TestCreateOrder(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	clientOrderId, err := uuid.NewUUID()
	if err != nil {
		t.Fatal("could not generate UUID:", err)
	}

	service := orders.NewOrdersService(client)

	response, err := service.CreateOrder(
		ctx,
		&orders.CreateOrderRequest{
			PortfolioId:   client.Credentials().PortfolioId,
			ClientOrderId: clientOrderId.String(),
			Side:          "BUY",
			InstrumentId:  "b3469e0b-222c-4f8a-9f68-1f9e44d7e5e0",
			Type:          "LIMIT",
			Tif:           "GTC",
			Size:          "0.1",
			Price:         "43000",
		})
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected a not nil response")
	}

	if len(response.Order.OrderId) == 0 {
		t.Fatal("expected an activity id")
	}

}
