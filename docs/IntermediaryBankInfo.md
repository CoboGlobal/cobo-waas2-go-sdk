# IntermediaryBankInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankName** | **string** | The name of the intermediary bank. | 
**BankAddress** | **string** | The address of the intermediary bank. | 
**BankSwiftCode** | **string** | The SWIFT or BIC code of the intermediary bank. | 

## Methods

### NewIntermediaryBankInfo

`func NewIntermediaryBankInfo(bankName string, bankAddress string, bankSwiftCode string, ) *IntermediaryBankInfo`

NewIntermediaryBankInfo instantiates a new IntermediaryBankInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntermediaryBankInfoWithDefaults

`func NewIntermediaryBankInfoWithDefaults() *IntermediaryBankInfo`

NewIntermediaryBankInfoWithDefaults instantiates a new IntermediaryBankInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankName

`func (o *IntermediaryBankInfo) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *IntermediaryBankInfo) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *IntermediaryBankInfo) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankAddress

`func (o *IntermediaryBankInfo) GetBankAddress() string`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *IntermediaryBankInfo) GetBankAddressOk() (*string, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *IntermediaryBankInfo) SetBankAddress(v string)`

SetBankAddress sets BankAddress field to given value.


### GetBankSwiftCode

`func (o *IntermediaryBankInfo) GetBankSwiftCode() string`

GetBankSwiftCode returns the BankSwiftCode field if non-nil, zero value otherwise.

### GetBankSwiftCodeOk

`func (o *IntermediaryBankInfo) GetBankSwiftCodeOk() (*string, bool)`

GetBankSwiftCodeOk returns a tuple with the BankSwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankSwiftCode

`func (o *IntermediaryBankInfo) SetBankSwiftCode(v string)`

SetBankSwiftCode sets BankSwiftCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


