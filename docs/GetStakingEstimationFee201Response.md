# GetStakingEstimationFee201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | Pointer to [**FeeType**](FeeType.md) |  | [optional] [default to FEETYPE_EVM_EIP_1559]
**FeeAmount** | Pointer to **string** | The amount of the estimated fee. | [optional] 
**TokenId** | Pointer to **string** | The token ID of the staking fee. | [optional] 

## Methods

### NewGetStakingEstimationFee201Response

`func NewGetStakingEstimationFee201Response() *GetStakingEstimationFee201Response`

NewGetStakingEstimationFee201Response instantiates a new GetStakingEstimationFee201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStakingEstimationFee201ResponseWithDefaults

`func NewGetStakingEstimationFee201ResponseWithDefaults() *GetStakingEstimationFee201Response`

NewGetStakingEstimationFee201ResponseWithDefaults instantiates a new GetStakingEstimationFee201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *GetStakingEstimationFee201Response) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *GetStakingEstimationFee201Response) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *GetStakingEstimationFee201Response) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *GetStakingEstimationFee201Response) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### GetFeeAmount

`func (o *GetStakingEstimationFee201Response) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *GetStakingEstimationFee201Response) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *GetStakingEstimationFee201Response) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *GetStakingEstimationFee201Response) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetTokenId

`func (o *GetStakingEstimationFee201Response) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetStakingEstimationFee201Response) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetStakingEstimationFee201Response) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *GetStakingEstimationFee201Response) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


