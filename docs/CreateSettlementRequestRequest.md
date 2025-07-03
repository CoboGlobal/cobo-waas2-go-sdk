# CreateSettlementRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a settlement request. The request ID is provided by you and must be unique. | 
**AcquiringType** | Pointer to [**AcquiringType**](AcquiringType.md) |  | [optional] 
**PayoutChannel** | Pointer to [**PayoutChannel**](PayoutChannel.md) |  | [optional] 
**SettlementType** | Pointer to [**SettlementType**](SettlementType.md) |  | [optional] 
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


### GetAcquiringType

`func (o *CreateSettlementRequestRequest) GetAcquiringType() AcquiringType`

GetAcquiringType returns the AcquiringType field if non-nil, zero value otherwise.

### GetAcquiringTypeOk

`func (o *CreateSettlementRequestRequest) GetAcquiringTypeOk() (*AcquiringType, bool)`

GetAcquiringTypeOk returns a tuple with the AcquiringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiringType

`func (o *CreateSettlementRequestRequest) SetAcquiringType(v AcquiringType)`

SetAcquiringType sets AcquiringType field to given value.

### HasAcquiringType

`func (o *CreateSettlementRequestRequest) HasAcquiringType() bool`

HasAcquiringType returns a boolean if a field has been set.

### GetPayoutChannel

`func (o *CreateSettlementRequestRequest) GetPayoutChannel() PayoutChannel`

GetPayoutChannel returns the PayoutChannel field if non-nil, zero value otherwise.

### GetPayoutChannelOk

`func (o *CreateSettlementRequestRequest) GetPayoutChannelOk() (*PayoutChannel, bool)`

GetPayoutChannelOk returns a tuple with the PayoutChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutChannel

`func (o *CreateSettlementRequestRequest) SetPayoutChannel(v PayoutChannel)`

SetPayoutChannel sets PayoutChannel field to given value.

### HasPayoutChannel

`func (o *CreateSettlementRequestRequest) HasPayoutChannel() bool`

HasPayoutChannel returns a boolean if a field has been set.

### GetSettlementType

`func (o *CreateSettlementRequestRequest) GetSettlementType() SettlementType`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *CreateSettlementRequestRequest) GetSettlementTypeOk() (*SettlementType, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *CreateSettlementRequestRequest) SetSettlementType(v SettlementType)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *CreateSettlementRequestRequest) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.

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


