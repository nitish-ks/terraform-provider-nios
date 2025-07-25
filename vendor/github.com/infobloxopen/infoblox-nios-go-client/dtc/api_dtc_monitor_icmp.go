/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type DtcMonitorIcmpAPI interface {
	/*
		Create Create a dtc:monitor:icmp object

		Creates a new dtc:monitor:icmp object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DtcMonitorIcmpAPICreateRequest
	*/
	Create(ctx context.Context) DtcMonitorIcmpAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateDtcMonitorIcmpResponse
	CreateExecute(r DtcMonitorIcmpAPICreateRequest) (*CreateDtcMonitorIcmpResponse, *http.Response, error)
	/*
		Delete Delete a dtc:monitor:icmp object

		Deletes a specific dtc:monitor:icmp object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the dtc:monitor:icmp object
		@return DtcMonitorIcmpAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) DtcMonitorIcmpAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r DtcMonitorIcmpAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve dtc:monitor:icmp objects

		Returns a list of dtc:monitor:icmp objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DtcMonitorIcmpAPIListRequest
	*/
	List(ctx context.Context) DtcMonitorIcmpAPIListRequest

	// ListExecute executes the request
	//  @return ListDtcMonitorIcmpResponse
	ListExecute(r DtcMonitorIcmpAPIListRequest) (*ListDtcMonitorIcmpResponse, *http.Response, error)
	/*
		Read Get a specific dtc:monitor:icmp object

		Returns a specific dtc:monitor:icmp object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the dtc:monitor:icmp object
		@return DtcMonitorIcmpAPIReadRequest
	*/
	Read(ctx context.Context, reference string) DtcMonitorIcmpAPIReadRequest

	// ReadExecute executes the request
	//  @return GetDtcMonitorIcmpResponse
	ReadExecute(r DtcMonitorIcmpAPIReadRequest) (*GetDtcMonitorIcmpResponse, *http.Response, error)
	/*
		Update Update a dtc:monitor:icmp object

		Updates a specific dtc:monitor:icmp object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the dtc:monitor:icmp object
		@return DtcMonitorIcmpAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) DtcMonitorIcmpAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateDtcMonitorIcmpResponse
	UpdateExecute(r DtcMonitorIcmpAPIUpdateRequest) (*UpdateDtcMonitorIcmpResponse, *http.Response, error)
}

// DtcMonitorIcmpAPIService DtcMonitorIcmpAPI service
type DtcMonitorIcmpAPIService internal.Service

type DtcMonitorIcmpAPICreateRequest struct {
	ctx              context.Context
	ApiService       DtcMonitorIcmpAPI
	dtcMonitorIcmp   *DtcMonitorIcmp
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r DtcMonitorIcmpAPICreateRequest) DtcMonitorIcmp(dtcMonitorIcmp DtcMonitorIcmp) DtcMonitorIcmpAPICreateRequest {
	r.dtcMonitorIcmp = &dtcMonitorIcmp
	return r
}

// Enter the field names followed by comma
func (r DtcMonitorIcmpAPICreateRequest) ReturnFields(returnFields string) DtcMonitorIcmpAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcMonitorIcmpAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcMonitorIcmpAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DtcMonitorIcmpAPICreateRequest) ReturnAsObject(returnAsObject int32) DtcMonitorIcmpAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DtcMonitorIcmpAPICreateRequest) Execute() (*CreateDtcMonitorIcmpResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a dtc:monitor:icmp object

Creates a new dtc:monitor:icmp object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DtcMonitorIcmpAPICreateRequest
*/
func (a *DtcMonitorIcmpAPIService) Create(ctx context.Context) DtcMonitorIcmpAPICreateRequest {
	return DtcMonitorIcmpAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateDtcMonitorIcmpResponse
func (a *DtcMonitorIcmpAPIService) CreateExecute(r DtcMonitorIcmpAPICreateRequest) (*CreateDtcMonitorIcmpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateDtcMonitorIcmpResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorIcmpAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:icmp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dtcMonitorIcmp == nil {
		return localVarReturnValue, nil, internal.ReportError("dtcMonitorIcmp is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.dtcMonitorIcmp != nil {
		if r.dtcMonitorIcmp.ExtAttrs == nil {
			r.dtcMonitorIcmp.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.dtcMonitorIcmp.ExtAttrs)[k]; !ok {
				(*r.dtcMonitorIcmp.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.dtcMonitorIcmp
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

type DtcMonitorIcmpAPIDeleteRequest struct {
	ctx        context.Context
	ApiService DtcMonitorIcmpAPI
	reference  string
}

func (r DtcMonitorIcmpAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a dtc:monitor:icmp object

Deletes a specific dtc:monitor:icmp object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the dtc:monitor:icmp object
	@return DtcMonitorIcmpAPIDeleteRequest
*/
func (a *DtcMonitorIcmpAPIService) Delete(ctx context.Context, reference string) DtcMonitorIcmpAPIDeleteRequest {
	return DtcMonitorIcmpAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *DtcMonitorIcmpAPIService) DeleteExecute(r DtcMonitorIcmpAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorIcmpAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:icmp/{reference}"
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

type DtcMonitorIcmpAPIListRequest struct {
	ctx              context.Context
	ApiService       DtcMonitorIcmpAPI
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
func (r DtcMonitorIcmpAPIListRequest) ReturnFields(returnFields string) DtcMonitorIcmpAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcMonitorIcmpAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcMonitorIcmpAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r DtcMonitorIcmpAPIListRequest) MaxResults(maxResults int32) DtcMonitorIcmpAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r DtcMonitorIcmpAPIListRequest) ReturnAsObject(returnAsObject int32) DtcMonitorIcmpAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r DtcMonitorIcmpAPIListRequest) Paging(paging int32) DtcMonitorIcmpAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r DtcMonitorIcmpAPIListRequest) PageId(pageId string) DtcMonitorIcmpAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r DtcMonitorIcmpAPIListRequest) Filters(filters map[string]interface{}) DtcMonitorIcmpAPIListRequest {
	r.filters = &filters
	return r
}

func (r DtcMonitorIcmpAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) DtcMonitorIcmpAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r DtcMonitorIcmpAPIListRequest) Execute() (*ListDtcMonitorIcmpResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve dtc:monitor:icmp objects

Returns a list of dtc:monitor:icmp objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DtcMonitorIcmpAPIListRequest
*/
func (a *DtcMonitorIcmpAPIService) List(ctx context.Context) DtcMonitorIcmpAPIListRequest {
	return DtcMonitorIcmpAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListDtcMonitorIcmpResponse
func (a *DtcMonitorIcmpAPIService) ListExecute(r DtcMonitorIcmpAPIListRequest) (*ListDtcMonitorIcmpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListDtcMonitorIcmpResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorIcmpAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:icmp"

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

type DtcMonitorIcmpAPIReadRequest struct {
	ctx              context.Context
	ApiService       DtcMonitorIcmpAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r DtcMonitorIcmpAPIReadRequest) ReturnFields(returnFields string) DtcMonitorIcmpAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcMonitorIcmpAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcMonitorIcmpAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DtcMonitorIcmpAPIReadRequest) ReturnAsObject(returnAsObject int32) DtcMonitorIcmpAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DtcMonitorIcmpAPIReadRequest) Execute() (*GetDtcMonitorIcmpResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific dtc:monitor:icmp object

Returns a specific dtc:monitor:icmp object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the dtc:monitor:icmp object
	@return DtcMonitorIcmpAPIReadRequest
*/
func (a *DtcMonitorIcmpAPIService) Read(ctx context.Context, reference string) DtcMonitorIcmpAPIReadRequest {
	return DtcMonitorIcmpAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetDtcMonitorIcmpResponse
func (a *DtcMonitorIcmpAPIService) ReadExecute(r DtcMonitorIcmpAPIReadRequest) (*GetDtcMonitorIcmpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetDtcMonitorIcmpResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorIcmpAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:icmp/{reference}"
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

type DtcMonitorIcmpAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       DtcMonitorIcmpAPI
	reference        string
	dtcMonitorIcmp   *DtcMonitorIcmp
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r DtcMonitorIcmpAPIUpdateRequest) DtcMonitorIcmp(dtcMonitorIcmp DtcMonitorIcmp) DtcMonitorIcmpAPIUpdateRequest {
	r.dtcMonitorIcmp = &dtcMonitorIcmp
	return r
}

// Enter the field names followed by comma
func (r DtcMonitorIcmpAPIUpdateRequest) ReturnFields(returnFields string) DtcMonitorIcmpAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcMonitorIcmpAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcMonitorIcmpAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DtcMonitorIcmpAPIUpdateRequest) ReturnAsObject(returnAsObject int32) DtcMonitorIcmpAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DtcMonitorIcmpAPIUpdateRequest) Execute() (*UpdateDtcMonitorIcmpResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a dtc:monitor:icmp object

Updates a specific dtc:monitor:icmp object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the dtc:monitor:icmp object
	@return DtcMonitorIcmpAPIUpdateRequest
*/
func (a *DtcMonitorIcmpAPIService) Update(ctx context.Context, reference string) DtcMonitorIcmpAPIUpdateRequest {
	return DtcMonitorIcmpAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateDtcMonitorIcmpResponse
func (a *DtcMonitorIcmpAPIService) UpdateExecute(r DtcMonitorIcmpAPIUpdateRequest) (*UpdateDtcMonitorIcmpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateDtcMonitorIcmpResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorIcmpAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:icmp/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dtcMonitorIcmp == nil {
		return localVarReturnValue, nil, internal.ReportError("dtcMonitorIcmp is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.dtcMonitorIcmp != nil {
		if r.dtcMonitorIcmp.ExtAttrs == nil {
			r.dtcMonitorIcmp.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.dtcMonitorIcmp.ExtAttrs)[k]; !ok {
				(*r.dtcMonitorIcmp.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.dtcMonitorIcmp
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
