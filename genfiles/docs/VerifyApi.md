# \VerifyApi

All URIs are relative to *http://localhost/getcoins*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyDocumentPost**](VerifyApi.md#VerifyDocumentPost) | **Post** /verify/document | Verify user identity with blockscore. Blockscore already checks OFAC and PEP so we might as well return those results every call to get our money&#39;s worth 



## VerifyDocumentPost

> VerifyUserResponse VerifyDocumentPost(ctx).NameFirst(nameFirst).NameLast(nameLast).DocumentType(documentType).DocumentValue(documentValue).Birthday(birthday).AddressStreet1(addressStreet1).AddressCity(addressCity).AddressSubdivision(addressSubdivision).AddressPostalCode(addressPostalCode).AddressCountryCode(addressCountryCode).Id(id).NameMiddle(nameMiddle).AddressStreet2(addressStreet2).PhoneNumber(phoneNumber).IpAddress(ipAddress).Note(note).Execute()

Verify user identity with blockscore. Blockscore already checks OFAC and PEP so we might as well return those results every call to get our money's worth 

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
    nameFirst := "nameFirst_example" // string | 
    nameLast := "nameLast_example" // string | 
    documentType := "documentType_example" // string | Must be either SSN or Drivers License
    documentValue := "documentValue_example" // string | 
    birthday := "birthday_example" // string | MM/DD/YYYY
    addressStreet1 := "addressStreet1_example" // string | 
    addressCity := "addressCity_example" // string | 
    addressSubdivision := "addressSubdivision_example" // string | 
    addressPostalCode := "addressPostalCode_example" // string | 
    addressCountryCode := "addressCountryCode_example" // string | 
    id := "id_example" // string |  (optional)
    nameMiddle := "nameMiddle_example" // string |  (optional)
    addressStreet2 := "addressStreet2_example" // string |  (optional)
    phoneNumber := "phoneNumber_example" // string |  (optional)
    ipAddress := "ipAddress_example" // string |  (optional)
    note := "note_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VerifyApi.VerifyDocumentPost(context.Background()).NameFirst(nameFirst).NameLast(nameLast).DocumentType(documentType).DocumentValue(documentValue).Birthday(birthday).AddressStreet1(addressStreet1).AddressCity(addressCity).AddressSubdivision(addressSubdivision).AddressPostalCode(addressPostalCode).AddressCountryCode(addressCountryCode).Id(id).NameMiddle(nameMiddle).AddressStreet2(addressStreet2).PhoneNumber(phoneNumber).IpAddress(ipAddress).Note(note).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifyApi.VerifyDocumentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyDocumentPost`: VerifyUserResponse
    fmt.Fprintf(os.Stdout, "Response from `VerifyApi.VerifyDocumentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyDocumentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFirst** | **string** |  | 
 **nameLast** | **string** |  | 
 **documentType** | **string** | Must be either SSN or Drivers License | 
 **documentValue** | **string** |  | 
 **birthday** | **string** | MM/DD/YYYY | 
 **addressStreet1** | **string** |  | 
 **addressCity** | **string** |  | 
 **addressSubdivision** | **string** |  | 
 **addressPostalCode** | **string** |  | 
 **addressCountryCode** | **string** |  | 
 **id** | **string** |  | 
 **nameMiddle** | **string** |  | 
 **addressStreet2** | **string** |  | 
 **phoneNumber** | **string** |  | 
 **ipAddress** | **string** |  | 
 **note** | **string** |  | 

### Return type

[**VerifyUserResponse**](VerifyUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

