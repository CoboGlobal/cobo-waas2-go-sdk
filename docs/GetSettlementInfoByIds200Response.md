# GetSettlementInfoByIds200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PspTokenBalances** | Pointer to [**[]SettlementInfo**](SettlementInfo.md) |  | [optional] 
**TokenBalances** | Pointer to [**[]SettlementInfo**](SettlementInfo.md) |  | [optional] 

## Methods

### NewGetSettlementInfoByIds200Response

`func NewGetSettlementInfoByIds200Response() *GetSettlementInfoByIds200Response`

NewGetSettlementInfoByIds200Response instantiates a new GetSettlementInfoByIds200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSettlementInfoByIds200ResponseWithDefaults

`func NewGetSettlementInfoByIds200ResponseWithDefaults() *GetSettlementInfoByIds200Response`

NewGetSettlementInfoByIds200ResponseWithDefaults instantiates a new GetSettlementInfoByIds200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPspTokenBalances

`func (o *GetSettlementInfoByIds200Response) GetPspTokenBalances() []SettlementInfo`

GetPspTokenBalances returns the PspTokenBalances field if non-nil, zero value otherwise.

### GetPspTokenBalancesOk

`func (o *GetSettlementInfoByIds200Response) GetPspTokenBalancesOk() (*[]SettlementInfo, bool)`

GetPspTokenBalancesOk returns a tuple with the PspTokenBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspTokenBalances

`func (o *GetSettlementInfoByIds200Response) SetPspTokenBalances(v []SettlementInfo)`

SetPspTokenBalances sets PspTokenBalances field to given value.

### HasPspTokenBalances

`func (o *GetSettlementInfoByIds200Response) HasPspTokenBalances() bool`

HasPspTokenBalances returns a boolean if a field has been set.

### GetTokenBalances

`func (o *GetSettlementInfoByIds200Response) GetTokenBalances() []SettlementInfo`

GetTokenBalances returns the TokenBalances field if non-nil, zero value otherwise.

### GetTokenBalancesOk

`func (o *GetSettlementInfoByIds200Response) GetTokenBalancesOk() (*[]SettlementInfo, bool)`

GetTokenBalancesOk returns a tuple with the TokenBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBalances

`func (o *GetSettlementInfoByIds200Response) SetTokenBalances(v []SettlementInfo)`

SetTokenBalances sets TokenBalances field to given value.

### HasTokenBalances

`func (o *GetSettlementInfoByIds200Response) HasTokenBalances() bool`

HasTokenBalances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


