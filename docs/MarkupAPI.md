# \MarkupAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Preview**](MarkupAPI.md#Preview) | **Post** /api/latest/markup/preview | Preview markdown render



## Preview

> RestMarkup Preview(ctx).HtmlEscape(htmlEscape).UrlMode(urlMode).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Body(body).Execute()

Preview markdown render



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	htmlEscape := "htmlEscape_example" // string | (Optional) true if HTML should be escaped in the input markup, false otherwise. (optional)
	urlMode := "urlMode_example" // string | (Optional) The mode to use when building URLs. One of: ABSOLUTE, RELATIVE or, CONFIGURED. By default this is RELATIVE. (optional)
	includeHeadingId := "includeHeadingId_example" // string | (Optional) true if headers should contain an ID based on the heading content. (optional)
	hardwrap := "hardwrap_example" // string | (Optional) Whether the markup implementation should convert newlines to breaks. By default this is false which reflects the standard markdown specification. (optional)
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarkupAPI.Preview(context.Background()).HtmlEscape(htmlEscape).UrlMode(urlMode).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarkupAPI.Preview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Preview`: RestMarkup
	fmt.Fprintf(os.Stdout, "Response from `MarkupAPI.Preview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **htmlEscape** | **string** | (Optional) true if HTML should be escaped in the input markup, false otherwise. | 
 **urlMode** | **string** | (Optional) The mode to use when building URLs. One of: ABSOLUTE, RELATIVE or, CONFIGURED. By default this is RELATIVE. | 
 **includeHeadingId** | **string** | (Optional) true if headers should contain an ID based on the heading content. | 
 **hardwrap** | **string** | (Optional) Whether the markup implementation should convert newlines to breaks. By default this is false which reflects the standard markdown specification. | 
 **body** | **string** |  | 

### Return type

[**RestMarkup**](RestMarkup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

