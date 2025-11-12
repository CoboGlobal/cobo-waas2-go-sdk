# FeeStationSponsorTransactionDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalFeeAmount** | **string** | The total amount used to sponsor the gas fee required for executing the main transaction on behalf of the user. | 
**CommissionFee** | **string** | The commission fee charged for sponsoring the gas fee.  | 

## Methods

### NewFeeStationSponsorTransactionDescription

`func NewFeeStationSponsorTransactionDescription(totalFeeAmount string, commissionFee string, ) *FeeStationSponsorTransactionDescription`

NewFeeStationSponsorTransactionDescription instantiates a new FeeStationSponsorTransactionDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeStationSponsorTransactionDescriptionWithDefaults

`func NewFeeStationSponsorTransactionDescriptionWithDefaults() *FeeStationSponsorTransactionDescription`

NewFeeStationSponsorTransactionDescriptionWithDefaults instantiates a new FeeStationSponsorTransactionDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalFeeAmount

`func (o *FeeStationSponsorTransactionDescription) GetTotalFeeAmount() string`

GetTotalFeeAmount returns the TotalFeeAmount field if non-nil, zero value otherwise.

### GetTotalFeeAmountOk

`func (o *FeeStationSponsorTransactionDescription) GetTotalFeeAmountOk() (*string, bool)`

GetTotalFeeAmountOk returns a tuple with the TotalFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFeeAmount

`func (o *FeeStationSponsorTransactionDescription) SetTotalFeeAmount(v string)`

SetTotalFeeAmount sets TotalFeeAmount field to given value.


### GetCommissionFee

`func (o *FeeStationSponsorTransactionDescription) GetCommissionFee() string`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *FeeStationSponsorTransactionDescription) GetCommissionFeeOk() (*string, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *FeeStationSponsorTransactionDescription) SetCommissionFee(v string)`

SetCommissionFee sets CommissionFee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


