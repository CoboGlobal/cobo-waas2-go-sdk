# ListExchanges200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**Subtitle** | Pointer to **string** | Introduction of the supported features of this exchange. | [optional] 
**SupportedSubWalletIds** | **[]string** | The trading accounts(sub_wallet_ids) supported for this exchange. | 

## Methods

### NewListExchanges200ResponseInner

`func NewListExchanges200ResponseInner(exchangeId ExchangeId, supportedSubWalletIds []string, ) *ListExchanges200ResponseInner`

NewListExchanges200ResponseInner instantiates a new ListExchanges200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListExchanges200ResponseInnerWithDefaults

`func NewListExchanges200ResponseInnerWithDefaults() *ListExchanges200ResponseInner`

NewListExchanges200ResponseInnerWithDefaults instantiates a new ListExchanges200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeId

`func (o *ListExchanges200ResponseInner) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *ListExchanges200ResponseInner) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *ListExchanges200ResponseInner) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetSubtitle

`func (o *ListExchanges200ResponseInner) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ListExchanges200ResponseInner) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ListExchanges200ResponseInner) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *ListExchanges200ResponseInner) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetSupportedSubWalletIds

`func (o *ListExchanges200ResponseInner) GetSupportedSubWalletIds() []string`

GetSupportedSubWalletIds returns the SupportedSubWalletIds field if non-nil, zero value otherwise.

### GetSupportedSubWalletIdsOk

`func (o *ListExchanges200ResponseInner) GetSupportedSubWalletIdsOk() (*[]string, bool)`

GetSupportedSubWalletIdsOk returns a tuple with the SupportedSubWalletIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSubWalletIds

`func (o *ListExchanges200ResponseInner) SetSupportedSubWalletIds(v []string)`

SetSupportedSubWalletIds sets SupportedSubWalletIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


