# ListExchanges200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**SupportedTradingAccountTypes** | **[]string** | The supported trading account types of this exchange. | 

## Methods

### NewListExchanges200ResponseInner

`func NewListExchanges200ResponseInner(exchangeId ExchangeId, supportedTradingAccountTypes []string, ) *ListExchanges200ResponseInner`

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


### GetSupportedTradingAccountTypes

`func (o *ListExchanges200ResponseInner) GetSupportedTradingAccountTypes() []string`

GetSupportedTradingAccountTypes returns the SupportedTradingAccountTypes field if non-nil, zero value otherwise.

### GetSupportedTradingAccountTypesOk

`func (o *ListExchanges200ResponseInner) GetSupportedTradingAccountTypesOk() (*[]string, bool)`

GetSupportedTradingAccountTypesOk returns a tuple with the SupportedTradingAccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTradingAccountTypes

`func (o *ListExchanges200ResponseInner) SetSupportedTradingAccountTypes(v []string)`

SetSupportedTradingAccountTypes sets SupportedTradingAccountTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


