# CreateSettlementRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a settlement request. The request ID is provided by you and must be unique. | 
**Settlements** | [**[]CreateSettlement**](CreateSettlement.md) |  | 

## Methods

### NewCreateSettlementRequestRequest

`func NewCreateSettlementRequestRequest(requestId string, settlements []CreateSettlement, ) *CreateSettlementRequestRequest`

NewCreateSettlementRequestRequest instantiates a new CreateSettlementRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSettlementRequestRequestWithDefaults

`func NewCreateSettlementRequestRequestWithDefaults() *CreateSettlementRequestRequest`

NewCreateSettlementRequestRequestWithDefaults instantiates a new CreateSettlementRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateSettlementRequestRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateSettlementRequestRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateSettlementRequestRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetSettlements

`func (o *CreateSettlementRequestRequest) GetSettlements() []CreateSettlement`

GetSettlements returns the Settlements field if non-nil, zero value otherwise.

### GetSettlementsOk

`func (o *CreateSettlementRequestRequest) GetSettlementsOk() (*[]CreateSettlement, bool)`

GetSettlementsOk returns a tuple with the Settlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlements

`func (o *CreateSettlementRequestRequest) SetSettlements(v []CreateSettlement)`

SetSettlements sets Settlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


