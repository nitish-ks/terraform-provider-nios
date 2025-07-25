/*
Infoblox RPZ API

OpenAPI specification for Infoblox NIOS WAPI RPZ objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rpz

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type RecordRpzAAPI interface {
	/*
		Create Create a record:rpz:a object

		Creates a new record:rpz:a object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RecordRpzAAPICreateRequest
	*/
	Create(ctx context.Context) RecordRpzAAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateRecordRpzAResponse
	CreateExecute(r RecordRpzAAPICreateRequest) (*CreateRecordRpzAResponse, *http.Response, error)
	/*
		Delete Delete a record:rpz:a object

		Deletes a specific record:rpz:a object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:rpz:a object
		@return RecordRpzAAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) RecordRpzAAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r RecordRpzAAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve record:rpz:a objects

		Returns a list of record:rpz:a objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RecordRpzAAPIListRequest
	*/
	List(ctx context.Context) RecordRpzAAPIListRequest

	// ListExecute executes the request
	//  @return ListRecordRpzAResponse
	ListExecute(r RecordRpzAAPIListRequest) (*ListRecordRpzAResponse, *http.Response, error)
	/*
		Read Get a specific record:rpz:a object

		Returns a specific record:rpz:a object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:rpz:a object
		@return RecordRpzAAPIReadRequest
	*/
	Read(ctx context.Context, reference string) RecordRpzAAPIReadRequest

	// ReadExecute executes the request
	//  @return GetRecordRpzAResponse
	ReadExecute(r RecordRpzAAPIReadRequest) (*GetRecordRpzAResponse, *http.Response, error)
	/*
		Update Update a record:rpz:a object

		Updates a specific record:rpz:a object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:rpz:a object
		@return RecordRpzAAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) RecordRpzAAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateRecordRpzAResponse
	UpdateExecute(r RecordRpzAAPIUpdateRequest) (*UpdateRecordRpzAResponse, *http.Response, error)
}

// RecordRpzAAPIService RecordRpzAAPI service
type RecordRpzAAPIService internal.Service

type RecordRpzAAPICreateRequest struct {
	ctx              context.Context
	ApiService       RecordRpzAAPI
	recordRpzA       *RecordRpzA
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r RecordRpzAAPICreateRequest) RecordRpzA(recordRpzA RecordRpzA) RecordRpzAAPICreateRequest {
	r.recordRpzA = &recordRpzA
	return r
}

// Enter the field names followed by comma
func (r RecordRpzAAPICreateRequest) ReturnFields(returnFields string) RecordRpzAAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordRpzAAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordRpzAAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordRpzAAPICreateRequest) ReturnAsObject(returnAsObject int32) RecordRpzAAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordRpzAAPICreateRequest) Execute() (*CreateRecordRpzAResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a record:rpz:a object

Creates a new record:rpz:a object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordRpzAAPICreateRequest
*/
func (a *RecordRpzAAPIService) Create(ctx context.Context) RecordRpzAAPICreateRequest {
	return RecordRpzAAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateRecordRpzAResponse
func (a *RecordRpzAAPIService) CreateExecute(r RecordRpzAAPICreateRequest) (*CreateRecordRpzAResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateRecordRpzAResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordRpzAAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:rpz:a"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordRpzA == nil {
		return localVarReturnValue, nil, internal.ReportError("recordRpzA is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.recordRpzA != nil {
		if r.recordRpzA.ExtAttrs == nil {
			r.recordRpzA.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.recordRpzA.ExtAttrs)[k]; !ok {
				(*r.recordRpzA.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.recordRpzA
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

type RecordRpzAAPIDeleteRequest struct {
	ctx        context.Context
	ApiService RecordRpzAAPI
	reference  string
}

func (r RecordRpzAAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a record:rpz:a object

Deletes a specific record:rpz:a object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:rpz:a object
	@return RecordRpzAAPIDeleteRequest
*/
func (a *RecordRpzAAPIService) Delete(ctx context.Context, reference string) RecordRpzAAPIDeleteRequest {
	return RecordRpzAAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *RecordRpzAAPIService) DeleteExecute(r RecordRpzAAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordRpzAAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:rpz:a/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type RecordRpzAAPIListRequest struct {
	ctx              context.Context
	ApiService       RecordRpzAAPI
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
func (r RecordRpzAAPIListRequest) ReturnFields(returnFields string) RecordRpzAAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordRpzAAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordRpzAAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r RecordRpzAAPIListRequest) MaxResults(maxResults int32) RecordRpzAAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r RecordRpzAAPIListRequest) ReturnAsObject(returnAsObject int32) RecordRpzAAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r RecordRpzAAPIListRequest) Paging(paging int32) RecordRpzAAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r RecordRpzAAPIListRequest) PageId(pageId string) RecordRpzAAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r RecordRpzAAPIListRequest) Filters(filters map[string]interface{}) RecordRpzAAPIListRequest {
	r.filters = &filters
	return r
}

func (r RecordRpzAAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) RecordRpzAAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r RecordRpzAAPIListRequest) Execute() (*ListRecordRpzAResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve record:rpz:a objects

Returns a list of record:rpz:a objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordRpzAAPIListRequest
*/
func (a *RecordRpzAAPIService) List(ctx context.Context) RecordRpzAAPIListRequest {
	return RecordRpzAAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListRecordRpzAResponse
func (a *RecordRpzAAPIService) ListExecute(r RecordRpzAAPIListRequest) (*ListRecordRpzAResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListRecordRpzAResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordRpzAAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:rpz:a"

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

type RecordRpzAAPIReadRequest struct {
	ctx              context.Context
	ApiService       RecordRpzAAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r RecordRpzAAPIReadRequest) ReturnFields(returnFields string) RecordRpzAAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordRpzAAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordRpzAAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordRpzAAPIReadRequest) ReturnAsObject(returnAsObject int32) RecordRpzAAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordRpzAAPIReadRequest) Execute() (*GetRecordRpzAResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific record:rpz:a object

Returns a specific record:rpz:a object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:rpz:a object
	@return RecordRpzAAPIReadRequest
*/
func (a *RecordRpzAAPIService) Read(ctx context.Context, reference string) RecordRpzAAPIReadRequest {
	return RecordRpzAAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetRecordRpzAResponse
func (a *RecordRpzAAPIService) ReadExecute(r RecordRpzAAPIReadRequest) (*GetRecordRpzAResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetRecordRpzAResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordRpzAAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:rpz:a/{reference}"
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

type RecordRpzAAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       RecordRpzAAPI
	reference        string
	recordRpzA       *RecordRpzA
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r RecordRpzAAPIUpdateRequest) RecordRpzA(recordRpzA RecordRpzA) RecordRpzAAPIUpdateRequest {
	r.recordRpzA = &recordRpzA
	return r
}

// Enter the field names followed by comma
func (r RecordRpzAAPIUpdateRequest) ReturnFields(returnFields string) RecordRpzAAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordRpzAAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordRpzAAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordRpzAAPIUpdateRequest) ReturnAsObject(returnAsObject int32) RecordRpzAAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordRpzAAPIUpdateRequest) Execute() (*UpdateRecordRpzAResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a record:rpz:a object

Updates a specific record:rpz:a object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:rpz:a object
	@return RecordRpzAAPIUpdateRequest
*/
func (a *RecordRpzAAPIService) Update(ctx context.Context, reference string) RecordRpzAAPIUpdateRequest {
	return RecordRpzAAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateRecordRpzAResponse
func (a *RecordRpzAAPIService) UpdateExecute(r RecordRpzAAPIUpdateRequest) (*UpdateRecordRpzAResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateRecordRpzAResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordRpzAAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:rpz:a/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordRpzA == nil {
		return localVarReturnValue, nil, internal.ReportError("recordRpzA is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.recordRpzA != nil {
		if r.recordRpzA.ExtAttrs == nil {
			r.recordRpzA.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.recordRpzA.ExtAttrs)[k]; !ok {
				(*r.recordRpzA.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.recordRpzA
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
