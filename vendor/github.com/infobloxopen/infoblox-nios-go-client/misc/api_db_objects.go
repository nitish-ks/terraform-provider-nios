/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type DbObjectsAPI interface {
	/*
		List Retrieve db_objects objects

		Returns a list of db_objects objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DbObjectsAPIListRequest
	*/
	List(ctx context.Context) DbObjectsAPIListRequest

	// ListExecute executes the request
	//  @return ListDbObjectsResponse
	ListExecute(r DbObjectsAPIListRequest) (*ListDbObjectsResponse, *http.Response, error)
	/*
		Read Get a specific db_objects object

		Returns a specific db_objects object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the db_objects object
		@return DbObjectsAPIReadRequest
	*/
	Read(ctx context.Context, reference string) DbObjectsAPIReadRequest

	// ReadExecute executes the request
	//  @return GetDbObjectsResponse
	ReadExecute(r DbObjectsAPIReadRequest) (*GetDbObjectsResponse, *http.Response, error)
}

// DbObjectsAPIService DbObjectsAPI service
type DbObjectsAPIService internal.Service

type DbObjectsAPIListRequest struct {
	ctx              context.Context
	ApiService       DbObjectsAPI
	returnFields     *string
	returnFieldsPlus *string
	maxResults       *int32
	returnAsObject   *int32
	paging           *int32
	pageId           *string
	filters          *map[string]interface{}
	extattrfilter    *map[string]interface{}
}

// Enter the field names followed by comma
func (r DbObjectsAPIListRequest) ReturnFields(returnFields string) DbObjectsAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DbObjectsAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) DbObjectsAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r DbObjectsAPIListRequest) MaxResults(maxResults int32) DbObjectsAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r DbObjectsAPIListRequest) ReturnAsObject(returnAsObject int32) DbObjectsAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r DbObjectsAPIListRequest) Paging(paging int32) DbObjectsAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r DbObjectsAPIListRequest) PageId(pageId string) DbObjectsAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r DbObjectsAPIListRequest) Filters(filters map[string]interface{}) DbObjectsAPIListRequest {
	r.filters = &filters
	return r
}

func (r DbObjectsAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) DbObjectsAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r DbObjectsAPIListRequest) Execute() (*ListDbObjectsResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve db_objects objects

Returns a list of db_objects objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DbObjectsAPIListRequest
*/
func (a *DbObjectsAPIService) List(ctx context.Context) DbObjectsAPIListRequest {
	return DbObjectsAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListDbObjectsResponse
func (a *DbObjectsAPIService) ListExecute(r DbObjectsAPIListRequest) (*ListDbObjectsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListDbObjectsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DbObjectsAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/db_objects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.maxResults != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_max_results", r.maxResults, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	if r.paging != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_paging", r.paging, "form", "")
	}
	if r.pageId != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_id", r.pageId, "form", "")
	}
	if r.filters != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "filters", r.filters, "form", "")
	}
	if r.extattrfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "extattrfilter", r.extattrfilter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type DbObjectsAPIReadRequest struct {
	ctx              context.Context
	ApiService       DbObjectsAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r DbObjectsAPIReadRequest) ReturnFields(returnFields string) DbObjectsAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DbObjectsAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) DbObjectsAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DbObjectsAPIReadRequest) ReturnAsObject(returnAsObject int32) DbObjectsAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DbObjectsAPIReadRequest) Execute() (*GetDbObjectsResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific db_objects object

Returns a specific db_objects object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the db_objects object
	@return DbObjectsAPIReadRequest
*/
func (a *DbObjectsAPIService) Read(ctx context.Context, reference string) DbObjectsAPIReadRequest {
	return DbObjectsAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetDbObjectsResponse
func (a *DbObjectsAPIService) ReadExecute(r DbObjectsAPIReadRequest) (*GetDbObjectsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetDbObjectsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DbObjectsAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/db_objects/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
