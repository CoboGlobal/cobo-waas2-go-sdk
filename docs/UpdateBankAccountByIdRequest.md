# UpdateBankAccountByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | **map[string]interface{}** | JSON-formatted bank account details. The object should include the following fields: - beneficiary_name: Name of the account holder - beneficiary_address: Address of the account holder - account_number: Bank account number - bank_name: Name of the bank - bank_address: Address of the bank - iban: (Optional) International Bank Account Number - swift_or_bic: SWIFT or BIC code of the bank  | 

## Methods

### NewUpdateBankAccountByIdRequest

`func NewUpdateBankAccountByIdRequest(info map[string]interface{}, ) *UpdateBankAccountByIdRequest`

NewUpdateBankAccountByIdRequest instantiates a new UpdateBankAccountByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBankAccountByIdRequestWithDefaults

`func NewUpdateBankAccountByIdRequestWithDefaults() *UpdateBankAccountByIdRequest`

NewUpdateBankAccountByIdRequestWithDefaults instantiates a new UpdateBankAccountByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *UpdateBankAccountByIdRequest) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UpdateBankAccountByIdRequest) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UpdateBankAccountByIdRequest) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


