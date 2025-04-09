# TransactionBabylonBusinessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraType** | [**TransactionExtraType**](TransactionExtraType.md) |  | 
**BabylonAddressInfo** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**BtcAddressInfo** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 

## Methods

### NewTransactionBabylonBusinessInfo

`func NewTransactionBabylonBusinessInfo(extraType TransactionExtraType, ) *TransactionBabylonBusinessInfo`

NewTransactionBabylonBusinessInfo instantiates a new TransactionBabylonBusinessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionBabylonBusinessInfoWithDefaults

`func NewTransactionBabylonBusinessInfoWithDefaults() *TransactionBabylonBusinessInfo`

NewTransactionBabylonBusinessInfoWithDefaults instantiates a new TransactionBabylonBusinessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraType

`func (o *TransactionBabylonBusinessInfo) GetExtraType() TransactionExtraType`

GetExtraType returns the ExtraType field if non-nil, zero value otherwise.

### GetExtraTypeOk

`func (o *TransactionBabylonBusinessInfo) GetExtraTypeOk() (*TransactionExtraType, bool)`

GetExtraTypeOk returns a tuple with the ExtraType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraType

`func (o *TransactionBabylonBusinessInfo) SetExtraType(v TransactionExtraType)`

SetExtraType sets ExtraType field to given value.


### GetBabylonAddressInfo

`func (o *TransactionBabylonBusinessInfo) GetBabylonAddressInfo() AddressInfo`

GetBabylonAddressInfo returns the BabylonAddressInfo field if non-nil, zero value otherwise.

### GetBabylonAddressInfoOk

`func (o *TransactionBabylonBusinessInfo) GetBabylonAddressInfoOk() (*AddressInfo, bool)`

GetBabylonAddressInfoOk returns a tuple with the BabylonAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBabylonAddressInfo

`func (o *TransactionBabylonBusinessInfo) SetBabylonAddressInfo(v AddressInfo)`

SetBabylonAddressInfo sets BabylonAddressInfo field to given value.

### HasBabylonAddressInfo

`func (o *TransactionBabylonBusinessInfo) HasBabylonAddressInfo() bool`

HasBabylonAddressInfo returns a boolean if a field has been set.

### GetBtcAddressInfo

`func (o *TransactionBabylonBusinessInfo) GetBtcAddressInfo() AddressInfo`

GetBtcAddressInfo returns the BtcAddressInfo field if non-nil, zero value otherwise.

### GetBtcAddressInfoOk

`func (o *TransactionBabylonBusinessInfo) GetBtcAddressInfoOk() (*AddressInfo, bool)`

GetBtcAddressInfoOk returns a tuple with the BtcAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtcAddressInfo

`func (o *TransactionBabylonBusinessInfo) SetBtcAddressInfo(v AddressInfo)`

SetBtcAddressInfo sets BtcAddressInfo field to given value.

### HasBtcAddressInfo

`func (o *TransactionBabylonBusinessInfo) HasBtcAddressInfo() bool`

HasBtcAddressInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


