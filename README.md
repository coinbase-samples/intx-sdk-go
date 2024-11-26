# Coinbase International Exchange (INTX) Go SDK

[![GoDoc](https://godoc.org/github.com/coinbase-samples/intx-sdk-go?status.svg)](https://godoc.org/github.com/coinbase-samples/intx-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/coinbase-samples/intx-sdk-go)](https://goreportcard.com/report/coinbase-samples/intx-sdk-go)

## Overview

The *INTX Go SDK* is a sample library that demonstrates the structure of a [Coinbase International Exchange (INTX)](https://international.coinbase.com/) driver for
the [REST APIs](https://docs.cloud.coinbase.com/intx/reference).

## License

The *INTX Go SDK* sample library is free and open source and released under the [Apache License, Version 2.0](LICENSE).

The application and code are only available for demonstration purposes.

## Usage

To use the *INTX Go SDK*, initialize the [Credentials](credentials/credentials.go) struct and create a new client. The Credentials struct is JSON
enabled. Ensure that INTX API credentials are stored in a secure manner.

```
credentials := &credentials.Credentials{}
if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
    log.Fatalf("unable to load INTX_CREDENTIALS: %v", err)
}

httpClient, err := client.DefaultHttpClient()
if err != nil {
    log.Fatalf("unable to create the http client: %v", err)
}

restClient := client.NewRestClient(credentials, httpClient)
```

There are convenience functions to read the credentials as an environment variable (credentials.ReadEnvCredentials) and to deserialize the JSON structure (credentials.UnmarshalCredentials) if pulled from a different source. The JSON format expected by both is:

```
{
  "accessKey": "",
  "passphrase": "",
  "signingKey": "",
  "portfolioId": "",
}
```

Coinbase INTX API credentials can be created in the INTX web console under API on the left-hand panel. 

Once the client is initialized, make the desired call by creating the relevant service. For example, to [list portfolios](https://github.com/coinbase-samples/intx-sdk-go/blob/main/list_portfolios.go),
pass in the request object, check for an error, and if nil, process the response.


```
portfoliosSvc := portfolios.NewPortfoliosService(restClient)
response, err := portfoliosSvc.ListPortfolios(ctx, &portfolios.ListPortfoliosRequest{})
```

## Build

To build the sample library, ensure that [Go](https://go.dev/) 1.19+ is installed and then run:

```bash
go build ./...
```
