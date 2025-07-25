/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type MembercloudsyncAPI interface {
	/*
		List Retrieve membercloudsync objects

		Returns a list of membercloudsync objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return MembercloudsyncAPIListRequest
	*/
	List(ctx context.Context) MembercloudsyncAPIListRequest

	// ListExecute executes the request
	//  @return ListMembercloudsyncResponse
	ListExecute(r MembercloudsyncAPIListRequest) (*ListMembercloudsyncResponse, *http.Response, error)
	/*
		Read Get a specific membercloudsync object

		Returns a specific membercloudsync object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the membercloudsync object
		@return MembercloudsyncAPIReadRequest
	*/
	Read(ctx context.Context, reference string) MembercloudsyncAPIReadRequest

	// ReadExecute executes the request
	//  @return GetMembercloudsyncResponse
	ReadExecute(r MembercloudsyncAPIReadRequest) (*GetMembercloudsyncResponse, *http.Response, error)
	/*
		Update Update a membercloudsync object

		Updates a specific membercloudsync object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the membercloudsync object
		@return MembercloudsyncAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) MembercloudsyncAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateMembercloudsyncResponse
	UpdateExecute(r MembercloudsyncAPIUpdateRequest) (*UpdateMembercloudsyncResponse, *http.Response, error)
}

// MembercloudsyncAPIService MembercloudsyncAPI service
type MembercloudsyncAPIService internal.Service

type MembercloudsyncAPIListRequest struct {
	ctx              context.Context
	ApiService       MembercloudsyncAPI
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
func (r MembercloudsyncAPIListRequest) ReturnFields(returnFields string) MembercloudsyncAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MembercloudsyncAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) MembercloudsyncAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r MembercloudsyncAPIListRequest) MaxResults(maxResults int32) MembercloudsyncAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r MembercloudsyncAPIListRequest) ReturnAsObject(returnAsObject int32) MembercloudsyncAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r MembercloudsyncAPIListRequest) Paging(paging int32) MembercloudsyncAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r MembercloudsyncAPIListRequest) PageId(pageId string) MembercloudsyncAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r MembercloudsyncAPIListRequest) Filters(filters map[string]interface{}) MembercloudsyncAPIListRequest {
	r.filters = &filters
	return r
}

func (r MembercloudsyncAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) MembercloudsyncAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r MembercloudsyncAPIListRequest) Execute() (*ListMembercloudsyncResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve membercloudsync objects

Returns a list of membercloudsync objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MembercloudsyncAPIListRequest
*/
func (a *MembercloudsyncAPIService) List(ctx context.Context) MembercloudsyncAPIListRequest {
	return MembercloudsyncAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListMembercloudsyncResponse
func (a *MembercloudsyncAPIService) ListExecute(r MembercloudsyncAPIListRequest) (*ListMembercloudsyncResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListMembercloudsyncResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MembercloudsyncAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/membercloudsync"

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

type MembercloudsyncAPIReadRequest struct {
	ctx              context.Context
	ApiService       MembercloudsyncAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r MembercloudsyncAPIReadRequest) ReturnFields(returnFields string) MembercloudsyncAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MembercloudsyncAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) MembercloudsyncAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r MembercloudsyncAPIReadRequest) ReturnAsObject(returnAsObject int32) MembercloudsyncAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r MembercloudsyncAPIReadRequest) Execute() (*GetMembercloudsyncResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific membercloudsync object

Returns a specific membercloudsync object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the membercloudsync object
	@return MembercloudsyncAPIReadRequest
*/
func (a *MembercloudsyncAPIService) Read(ctx context.Context, reference string) MembercloudsyncAPIReadRequest {
	return MembercloudsyncAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetMembercloudsyncResponse
func (a *MembercloudsyncAPIService) ReadExecute(r MembercloudsyncAPIReadRequest) (*GetMembercloudsyncResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetMembercloudsyncResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MembercloudsyncAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/membercloudsync/{reference}"
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

type MembercloudsyncAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       MembercloudsyncAPI
	reference        string
	membercloudsync  *Membercloudsync
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r MembercloudsyncAPIUpdateRequest) Membercloudsync(membercloudsync Membercloudsync) MembercloudsyncAPIUpdateRequest {
	r.membercloudsync = &membercloudsync
	return r
}

// Enter the field names followed by comma
func (r MembercloudsyncAPIUpdateRequest) ReturnFields(returnFields string) MembercloudsyncAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MembercloudsyncAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) MembercloudsyncAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r MembercloudsyncAPIUpdateRequest) ReturnAsObject(returnAsObject int32) MembercloudsyncAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r MembercloudsyncAPIUpdateRequest) Execute() (*UpdateMembercloudsyncResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a membercloudsync object

Updates a specific membercloudsync object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the membercloudsync object
	@return MembercloudsyncAPIUpdateRequest
*/
func (a *MembercloudsyncAPIService) Update(ctx context.Context, reference string) MembercloudsyncAPIUpdateRequest {
	return MembercloudsyncAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateMembercloudsyncResponse
func (a *MembercloudsyncAPIService) UpdateExecute(r MembercloudsyncAPIUpdateRequest) (*UpdateMembercloudsyncResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateMembercloudsyncResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MembercloudsyncAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/membercloudsync/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.membercloudsync == nil {
		return localVarReturnValue, nil, internal.ReportError("membercloudsync is required and must be specified")
	}

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
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.membercloudsync
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
