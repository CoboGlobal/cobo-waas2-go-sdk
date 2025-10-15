# OtcFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeRate** | **string** | The otc fee rate. | 
**TokenId** | Pointer to **string** | The token id in otc. | [optional] 

## Methods

### NewOtcFee

`func NewOtcFee(feeRate string, ) *OtcFee`

NewOtcFee instantiates a new OtcFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtcFeeWithDefaults

`func NewOtcFeeWithDefaults() *OtcFee`

NewOtcFeeWithDefaults instantiates a new OtcFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeRate

`func (o *OtcFee) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *OtcFee) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *OtcFee) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetTokenId

`func (o *OtcFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *OtcFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *OtcFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *OtcFee) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


