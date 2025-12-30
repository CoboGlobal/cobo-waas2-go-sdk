# CreateReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **int32** | The start time of the report. Unix timestamp measured in seconds. | 
**EndTime** | **int32** | The end time of the report. Unix timestamp measured in seconds. | 
**ReportExportFormat** | [**ReportExportFormat**](ReportExportFormat.md) |  | 
**ReportTypes** | [**[]ReportType**](ReportType.md) |  | 
**TokenIds** | Pointer to **[]string** | Optional filter to include only items related to specified token IDs in the report. | [optional] 

## Methods

### NewCreateReportRequest

`func NewCreateReportRequest(startTime int32, endTime int32, reportExportFormat ReportExportFormat, reportTypes []ReportType, ) *CreateReportRequest`

NewCreateReportRequest instantiates a new CreateReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReportRequestWithDefaults

`func NewCreateReportRequestWithDefaults() *CreateReportRequest`

NewCreateReportRequestWithDefaults instantiates a new CreateReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *CreateReportRequest) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateReportRequest) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateReportRequest) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CreateReportRequest) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreateReportRequest) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreateReportRequest) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetReportExportFormat

`func (o *CreateReportRequest) GetReportExportFormat() ReportExportFormat`

GetReportExportFormat returns the ReportExportFormat field if non-nil, zero value otherwise.

### GetReportExportFormatOk

`func (o *CreateReportRequest) GetReportExportFormatOk() (*ReportExportFormat, bool)`

GetReportExportFormatOk returns a tuple with the ReportExportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportExportFormat

`func (o *CreateReportRequest) SetReportExportFormat(v ReportExportFormat)`

SetReportExportFormat sets ReportExportFormat field to given value.


### GetReportTypes

`func (o *CreateReportRequest) GetReportTypes() []ReportType`

GetReportTypes returns the ReportTypes field if non-nil, zero value otherwise.

### GetReportTypesOk

`func (o *CreateReportRequest) GetReportTypesOk() (*[]ReportType, bool)`

GetReportTypesOk returns a tuple with the ReportTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypes

`func (o *CreateReportRequest) SetReportTypes(v []ReportType)`

SetReportTypes sets ReportTypes field to given value.


### GetTokenIds

`func (o *CreateReportRequest) GetTokenIds() []string`

GetTokenIds returns the TokenIds field if non-nil, zero value otherwise.

### GetTokenIdsOk

`func (o *CreateReportRequest) GetTokenIdsOk() (*[]string, bool)`

GetTokenIdsOk returns a tuple with the TokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIds

`func (o *CreateReportRequest) SetTokenIds(v []string)`

SetTokenIds sets TokenIds field to given value.

### HasTokenIds

`func (o *CreateReportRequest) HasTokenIds() bool`

HasTokenIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


