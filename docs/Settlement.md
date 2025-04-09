# Settlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettlementRequestId** | **string** | Internal system-generated identifier with a human-readable format. Used primarily for internal operations and logging. | 
**RequestId** | **string** | Client-provided external identifier. | 
**Status** | [**SettleRequestStatus**](SettleRequestStatus.md) |  | 
**Settlements** | [**[]SettlementDetail**](SettlementDetail.md) |  | 

## Methods

### NewSettlement

`func NewSettlement(settlementRequestId string, requestId string, status SettleRequestStatus, settlements []SettlementDetail, ) *Settlement`

NewSettlement instantiates a new Settlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementWithDefaults

`func NewSettlementWithDefaults() *Settlement`

NewSettlementWithDefaults instantiates a new Settlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettlementRequestId

`func (o *Settlement) GetSettlementRequestId() string`

GetSettlementRequestId returns the SettlementRequestId field if non-nil, zero value otherwise.

### GetSettlementRequestIdOk

`func (o *Settlement) GetSettlementRequestIdOk() (*string, bool)`

GetSettlementRequestIdOk returns a tuple with the SettlementRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementRequestId

`func (o *Settlement) SetSettlementRequestId(v string)`

SetSettlementRequestId sets SettlementRequestId field to given value.


### GetRequestId

`func (o *Settlement) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Settlement) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Settlement) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetStatus

`func (o *Settlement) GetStatus() SettleRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Settlement) GetStatusOk() (*SettleRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Settlement) SetStatus(v SettleRequestStatus)`

SetStatus sets Status field to given value.


### GetSettlements

`func (o *Settlement) GetSettlements() []SettlementDetail`

GetSettlements returns the Settlements field if non-nil, zero value otherwise.

### GetSettlementsOk

`func (o *Settlement) GetSettlementsOk() (*[]SettlementDetail, bool)`

GetSettlementsOk returns a tuple with the Settlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlements

`func (o *Settlement) SetSettlements(v []SettlementDetail)`

SetSettlements sets Settlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


