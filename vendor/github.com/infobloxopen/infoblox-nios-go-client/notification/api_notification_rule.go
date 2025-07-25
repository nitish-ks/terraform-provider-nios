/*
Infoblox NOTIFICATION API

OpenAPI specification for Infoblox NIOS WAPI NOTIFICATION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type NotificationRuleAPI interface {
	/*
		Create Create a notification:rule object

		Creates a new notification:rule object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return NotificationRuleAPICreateRequest
	*/
	Create(ctx context.Context) NotificationRuleAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateNotificationRuleResponse
	CreateExecute(r NotificationRuleAPICreateRequest) (*CreateNotificationRuleResponse, *http.Response, error)
	/*
		Delete Delete a notification:rule object

		Deletes a specific notification:rule object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the notification:rule object
		@return NotificationRuleAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) NotificationRuleAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r NotificationRuleAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve notification:rule objects

		Returns a list of notification:rule objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return NotificationRuleAPIListRequest
	*/
	List(ctx context.Context) NotificationRuleAPIListRequest

	// ListExecute executes the request
	//  @return ListNotificationRuleResponse
	ListExecute(r NotificationRuleAPIListRequest) (*ListNotificationRuleResponse, *http.Response, error)
	/*
		Read Get a specific notification:rule object

		Returns a specific notification:rule object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the notification:rule object
		@return NotificationRuleAPIReadRequest
	*/
	Read(ctx context.Context, reference string) NotificationRuleAPIReadRequest

	// ReadExecute executes the request
	//  @return GetNotificationRuleResponse
	ReadExecute(r NotificationRuleAPIReadRequest) (*GetNotificationRuleResponse, *http.Response, error)
	/*
		Update Update a notification:rule object

		Updates a specific notification:rule object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the notification:rule object
		@return NotificationRuleAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) NotificationRuleAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateNotificationRuleResponse
	UpdateExecute(r NotificationRuleAPIUpdateRequest) (*UpdateNotificationRuleResponse, *http.Response, error)
}

// NotificationRuleAPIService NotificationRuleAPI service
type NotificationRuleAPIService internal.Service

type NotificationRuleAPICreateRequest struct {
	ctx              context.Context
	ApiService       NotificationRuleAPI
	notificationRule *NotificationRule
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r NotificationRuleAPICreateRequest) NotificationRule(notificationRule NotificationRule) NotificationRuleAPICreateRequest {
	r.notificationRule = &notificationRule
	return r
}

// Enter the field names followed by comma
func (r NotificationRuleAPICreateRequest) ReturnFields(returnFields string) NotificationRuleAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r NotificationRuleAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) NotificationRuleAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r NotificationRuleAPICreateRequest) ReturnAsObject(returnAsObject int32) NotificationRuleAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r NotificationRuleAPICreateRequest) Execute() (*CreateNotificationRuleResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a notification:rule object

Creates a new notification:rule object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return NotificationRuleAPICreateRequest
*/
func (a *NotificationRuleAPIService) Create(ctx context.Context) NotificationRuleAPICreateRequest {
	return NotificationRuleAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateNotificationRuleResponse
func (a *NotificationRuleAPIService) CreateExecute(r NotificationRuleAPICreateRequest) (*CreateNotificationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateNotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "NotificationRuleAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/notification:rule"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.notificationRule == nil {
		return localVarReturnValue, nil, internal.ReportError("notificationRule is required and must be specified")
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
	localVarPostBody = r.notificationRule
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

type NotificationRuleAPIDeleteRequest struct {
	ctx        context.Context
	ApiService NotificationRuleAPI
	reference  string
}

func (r NotificationRuleAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a notification:rule object

Deletes a specific notification:rule object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the notification:rule object
	@return NotificationRuleAPIDeleteRequest
*/
func (a *NotificationRuleAPIService) Delete(ctx context.Context, reference string) NotificationRuleAPIDeleteRequest {
	return NotificationRuleAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *NotificationRuleAPIService) DeleteExecute(r NotificationRuleAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "NotificationRuleAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/notification:rule/{reference}"
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

type NotificationRuleAPIListRequest struct {
	ctx              context.Context
	ApiService       NotificationRuleAPI
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
func (r NotificationRuleAPIListRequest) ReturnFields(returnFields string) NotificationRuleAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r NotificationRuleAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) NotificationRuleAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r NotificationRuleAPIListRequest) MaxResults(maxResults int32) NotificationRuleAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r NotificationRuleAPIListRequest) ReturnAsObject(returnAsObject int32) NotificationRuleAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r NotificationRuleAPIListRequest) Paging(paging int32) NotificationRuleAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r NotificationRuleAPIListRequest) PageId(pageId string) NotificationRuleAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r NotificationRuleAPIListRequest) Filters(filters map[string]interface{}) NotificationRuleAPIListRequest {
	r.filters = &filters
	return r
}

func (r NotificationRuleAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) NotificationRuleAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r NotificationRuleAPIListRequest) Execute() (*ListNotificationRuleResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve notification:rule objects

Returns a list of notification:rule objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return NotificationRuleAPIListRequest
*/
func (a *NotificationRuleAPIService) List(ctx context.Context) NotificationRuleAPIListRequest {
	return NotificationRuleAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListNotificationRuleResponse
func (a *NotificationRuleAPIService) ListExecute(r NotificationRuleAPIListRequest) (*ListNotificationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListNotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "NotificationRuleAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/notification:rule"

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

type NotificationRuleAPIReadRequest struct {
	ctx              context.Context
	ApiService       NotificationRuleAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r NotificationRuleAPIReadRequest) ReturnFields(returnFields string) NotificationRuleAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r NotificationRuleAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) NotificationRuleAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r NotificationRuleAPIReadRequest) ReturnAsObject(returnAsObject int32) NotificationRuleAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r NotificationRuleAPIReadRequest) Execute() (*GetNotificationRuleResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific notification:rule object

Returns a specific notification:rule object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the notification:rule object
	@return NotificationRuleAPIReadRequest
*/
func (a *NotificationRuleAPIService) Read(ctx context.Context, reference string) NotificationRuleAPIReadRequest {
	return NotificationRuleAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetNotificationRuleResponse
func (a *NotificationRuleAPIService) ReadExecute(r NotificationRuleAPIReadRequest) (*GetNotificationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetNotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "NotificationRuleAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/notification:rule/{reference}"
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

type NotificationRuleAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       NotificationRuleAPI
	reference        string
	notificationRule *NotificationRule
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r NotificationRuleAPIUpdateRequest) NotificationRule(notificationRule NotificationRule) NotificationRuleAPIUpdateRequest {
	r.notificationRule = &notificationRule
	return r
}

// Enter the field names followed by comma
func (r NotificationRuleAPIUpdateRequest) ReturnFields(returnFields string) NotificationRuleAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r NotificationRuleAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) NotificationRuleAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r NotificationRuleAPIUpdateRequest) ReturnAsObject(returnAsObject int32) NotificationRuleAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r NotificationRuleAPIUpdateRequest) Execute() (*UpdateNotificationRuleResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a notification:rule object

Updates a specific notification:rule object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the notification:rule object
	@return NotificationRuleAPIUpdateRequest
*/
func (a *NotificationRuleAPIService) Update(ctx context.Context, reference string) NotificationRuleAPIUpdateRequest {
	return NotificationRuleAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateNotificationRuleResponse
func (a *NotificationRuleAPIService) UpdateExecute(r NotificationRuleAPIUpdateRequest) (*UpdateNotificationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateNotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "NotificationRuleAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/notification:rule/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.notificationRule == nil {
		return localVarReturnValue, nil, internal.ReportError("notificationRule is required and must be specified")
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
	localVarPostBody = r.notificationRule
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
