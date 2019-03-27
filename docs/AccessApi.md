# AccessApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------ | -------------| -----------
[**createAccessToken**](AccessApi.md#createAccessToken) | **POST** /access/token | 
[**createAccessTokenFromTicket**](AccessApi.md#createAccessTokenFromTicket) | **POST** /access/kerberos | 
[**createDownloadToken**](AccessApi.md#createDownloadToken) | **POST** /access/download-token | 
[**createUiExtensionToken**](AccessApi.md#createUiExtensionToken) | **POST** /access/ui-extension-token | 
[**getAccessStatus**](AccessApi.md#getAccessStatus) | **GET** /access | 
[**getLoginConfig**](AccessApi.md#getLoginConfig) | **GET** /access/config | 

<a name="createAccessToken"></a>
# **createAccessToken**

Creates a token for accessing the REST API via username/password

The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format &#39;Authorization: Bearer &lt;token&gt;&#39;.

### Example

```go

```


<a name="createAccessTokenFromTicket"></a>
# **createAccessTokenFromTicket**

### Example

```go

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **String**|  | [optional]
 **password** | **String**|  | [optional]

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/plain
 
 <a name="createAccessTokenFromTicket"></a>
 # **createAccessTokenFromTicket**
 
 Creates a token for accessing the REST API via Kerberos ticket exchange / SPNEGO negotiation
 
 The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format &#39;Authorization: Bearer &lt;token&gt;&#39;.
 
 ### Example
 ```go

 ```
 
 ### Parameters
 This endpoint does not need any parameter.
 
 ### Return type
 
 **String**
 
 ### Authorization
 
 No authorization required
 
 ### HTTP request headers
 
  - **Content-Type**: text/plain
  - **Accept**: text/plain
 
 <a name="createDownloadToken"></a>
 # **createDownloadToken**
 
 Creates a single use access token for downloading FlowFile content.
 
 The token returned is a base64 encoded string. It is valid for a single request up to five minutes from being issued. It is used as a query parameter name &#39;access_token&#39;.
 
 ### Example
 ```go
 
 ```
 
 ### Parameters
 This endpoint does not need any parameter.
 
 ### Return type
 
 **String**
 
 ### Authorization
 
 No authorization required
 
 ### HTTP request headers
 
  - **Content-Type**: application/x-www-form-urlencoded
  - **Accept**: text/plain
 
 <a name="createUiExtensionToken"></a>
 # **createUiExtensionToken**
 
 Creates a single use access token for accessing a NiFi UI extension.
 
 The token returned is a base64 encoded string. It is valid for a single request up to five minutes from being issued. It is used as a query parameter name &#39;access_token&#39;.
 
 ### Example
 ```go

 ```
 
 ### Parameters
 This endpoint does not need any parameter.
 
 ### Return type
 
 **String**
 
 ### Authorization
 
 No authorization required
 
 ### HTTP request headers
 
  - **Content-Type**: application/x-www-form-urlencoded
  - **Accept**: text/plain
 
 <a name="getAccessStatus"></a>
 # **getAccessStatus**
 
 Gets the status the client&#39;s access
 
 Note: This endpoint is subject to change as NiFi and it&#39;s REST API evolve.
 
 ### Example
 ```go

 ```
 
 ### Parameters
 This endpoint does not need any parameter.
 
 ### Return type
 
 [**AccessStatusEntity**](AccessStatusEntity.md)
 
 ### Authorization
 
 No authorization required
 
 ### HTTP request headers
 
  - **Content-Type**: *_/_*
  - **Accept**: application/json
 
 <a name="getLoginConfig"></a>
 # **getLoginConfig**
 
 Retrieves the access configuration for this NiFi
 
 
 
 ### Example
 ```go

 ```
 
 ### Parameters
 This endpoint does not need any parameter.
 
 ### Return type
 
 [**AccessConfigurationEntity**](AccessConfigurationEntity.md)
 
 ### Authorization
 
 No authorization required
 
 ### HTTP request headers
 
  - **Content-Type**: *_/_*
  - **Accept**: application/json
