# \CheckApi

All URIs are relative to *http://localhost/getcoins*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDigitalCurrencyAddressGet**](CheckApi.md#CheckDigitalCurrencyAddressGet) | **Get** /check/digital-currency-address | Check if digital currency address is in CSL
[**CheckNameGet**](CheckApi.md#CheckNameGet) | **Get** /check/name | Check if user has watchlist hits.



## CheckDigitalCurrencyAddressGet

> CheckDigitalCurrencyAddressResponse CheckDigitalCurrencyAddressGet(ctx).Doa(doa).Execute()

Check if digital currency address is in CSL

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    doa := true // bool | Digital currency address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckApi.CheckDigitalCurrencyAddressGet(context.Background()).Doa(doa).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckApi.CheckDigitalCurrencyAddressGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckDigitalCurrencyAddressGet`: CheckDigitalCurrencyAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckApi.CheckDigitalCurrencyAddressGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckDigitalCurrencyAddressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **doa** | **bool** | Digital currency address | 

### Return type

[**CheckDigitalCurrencyAddressResponse**](CheckDigitalCurrencyAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckNameGet

> CheckNameResponse CheckNameGet(ctx).Minscore(minscore).FirstName(firstName).LastName(lastName).GetcoinsID(getcoinsID).MiddleName(middleName).Address(address).Country(country).Cached(cached).Execute()

Check if user has watchlist hits.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    minscore := int32(56) // int32 | Sets minimum score to define a match
    firstName := "firstName_example" // string | Name to search CSL list for match. First and last name are required.
    lastName := "lastName_example" // string | Name to search CSL list for match.
    getcoinsID := "getcoinsID_example" // string | Getcoins user id to relate to blockscore person id. This way, we don't need to ask blockscore for the ID and we can minimize calls to the service
    middleName := "middleName_example" // string | Not required but will help return more accurate results. (optional)
    address := "address_example" // string | Street address of person being searched (optional)
    country := "country_example" // string | Two letter country code to make search more specific (optional)
    cached := true // bool | Boolean to denote whether to bypass cache or not. May be omitted if unsure. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckApi.CheckNameGet(context.Background()).Minscore(minscore).FirstName(firstName).LastName(lastName).GetcoinsID(getcoinsID).MiddleName(middleName).Address(address).Country(country).Cached(cached).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckApi.CheckNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckNameGet`: CheckNameResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckApi.CheckNameGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minscore** | **int32** | Sets minimum score to define a match | 
 **firstName** | **string** | Name to search CSL list for match. First and last name are required. | 
 **lastName** | **string** | Name to search CSL list for match. | 
 **getcoinsID** | **string** | Getcoins user id to relate to blockscore person id. This way, we don&#39;t need to ask blockscore for the ID and we can minimize calls to the service | 
 **middleName** | **string** | Not required but will help return more accurate results. | 
 **address** | **string** | Street address of person being searched | 
 **country** | **string** | Two letter country code to make search more specific | 
 **cached** | **bool** | Boolean to denote whether to bypass cache or not. May be omitted if unsure. | [default to false]

### Return type

[**CheckNameResponse**](CheckNameResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

