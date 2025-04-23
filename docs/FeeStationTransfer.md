# FeeStationTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID of the transferred token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
**Destination** | Pointer to [**FeeStationDestination**](FeeStationDestination.md) |  | [optional] 

## Methods

### NewFeeStationTransfer

`func NewFeeStationTransfer(tokenId string, ) *FeeStationTransfer`

NewFeeStationTransfer instantiates a new FeeStationTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeStationTransferWithDefaults

`func NewFeeStationTransferWithDefaults() *FeeStationTransfer`

NewFeeStationTransferWithDefaults instantiates a new FeeStationTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *FeeStationTransfer) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *FeeStationTransfer) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *FeeStationTransfer) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDestination

`func (o *FeeStationTransfer) GetDestination() FeeStationDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FeeStationTransfer) GetDestinationOk() (*FeeStationDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FeeStationTransfer) SetDestination(v FeeStationDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *FeeStationTransfer) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


