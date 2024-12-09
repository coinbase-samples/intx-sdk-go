package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/credentials"
	"github.com/coinbase-samples/intx-sdk-go/portfolios"
)

func main() {
	credentials, err := credentials.ReadEnvCredentials("INTX_CREDENTIALS")
	if err != nil {
		panic(fmt.Sprintf("unable to read INTX credentials: %v", err))
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		panic(fmt.Sprintf("unable to load default http client: %v", err))
	}

	client := client.NewRestClient(credentials, httpClient)

	portfoliosSvc := portfolios.NewPortfoliosService(client)

	request := &portfolios.ListPortfoliosRequest{}

	response, err := portfoliosSvc.ListPortfolios(context.Background(), request)
	if err != nil {
		panic(fmt.Sprintf("unable to list portfolios: %v", err))
	}

	jsonResponse, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("error marshaling response to JSON: %v", err))
	}
	fmt.Println(string(jsonResponse))
}
