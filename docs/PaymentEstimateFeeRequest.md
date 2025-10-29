# PaymentEstimateFeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | Pointer to [**PaymentFeeType**](PaymentFeeType.md) |  | [optional] 
**EstimateFees** | [**[]PaymentEstimateFee**](PaymentEstimateFee.md) | A list of token IDs and amounts for which fees will be calculated. | 

## Methods

### NewPaymentEstimateFeeRequest

`func NewPaymentEstimateFeeRequest(estimateFees []PaymentEstimateFee, ) *PaymentEstimateFeeRequest`

NewPaymentEstimateFeeRequest instantiates a new PaymentEstimateFeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentEstimateFeeRequestWithDefaults

`func NewPaymentEstimateFeeRequestWithDefaults() *PaymentEstimateFeeRequest`

NewPaymentEstimateFeeRequestWithDefaults instantiates a new PaymentEstimateFeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *PaymentEstimateFeeRequest) GetFeeType() PaymentFeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *PaymentEstimateFeeRequest) GetFeeTypeOk() (*PaymentFeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *PaymentEstimateFeeRequest) SetFeeType(v PaymentFeeType)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *PaymentEstimateFeeRequest) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### GetEstimateFees

`func (o *PaymentEstimateFeeRequest) GetEstimateFees() []PaymentEstimateFee`

GetEstimateFees returns the EstimateFees field if non-nil, zero value otherwise.

### GetEstimateFeesOk

`func (o *PaymentEstimateFeeRequest) GetEstimateFeesOk() (*[]PaymentEstimateFee, bool)`

GetEstimateFeesOk returns a tuple with the EstimateFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateFees

`func (o *PaymentEstimateFeeRequest) SetEstimateFees(v []PaymentEstimateFee)`

SetEstimateFees sets EstimateFees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


