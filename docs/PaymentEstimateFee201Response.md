# PaymentEstimateFee201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PaymentEstimatedFee**](PaymentEstimatedFee.md) | A list of estimated fees for the requested operations. | [optional] 

## Methods

### NewPaymentEstimateFee201Response

`func NewPaymentEstimateFee201Response() *PaymentEstimateFee201Response`

NewPaymentEstimateFee201Response instantiates a new PaymentEstimateFee201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentEstimateFee201ResponseWithDefaults

`func NewPaymentEstimateFee201ResponseWithDefaults() *PaymentEstimateFee201Response`

NewPaymentEstimateFee201ResponseWithDefaults instantiates a new PaymentEstimateFee201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentEstimateFee201Response) GetData() []PaymentEstimatedFee`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentEstimateFee201Response) GetDataOk() (*[]PaymentEstimatedFee, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentEstimateFee201Response) SetData(v []PaymentEstimatedFee)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentEstimateFee201Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


