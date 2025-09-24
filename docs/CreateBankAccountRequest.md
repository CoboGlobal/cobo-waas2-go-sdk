# CreateBankAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtcBankAccountId** | **string** |  | 
**Info** | **map[string]interface{}** | JSON-formatted bank account details. The object should include the following fields: - beneficiary_name: Name of the account holder - beneficiary_address: Address of the account holder - account_number: Bank account number - bank_name: Name of the bank - bank_address: Address of the bank - iban: (Optional) International Bank Account Number - swift_or_bic: SWIFT or BIC code of the bank  | 

## Methods

### NewCreateBankAccountRequest

`func NewCreateBankAccountRequest(otcBankAccountId string, info map[string]interface{}, ) *CreateBankAccountRequest`

NewCreateBankAccountRequest instantiates a new CreateBankAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBankAccountRequestWithDefaults

`func NewCreateBankAccountRequestWithDefaults() *CreateBankAccountRequest`

NewCreateBankAccountRequestWithDefaults instantiates a new CreateBankAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtcBankAccountId

`func (o *CreateBankAccountRequest) GetOtcBankAccountId() string`

GetOtcBankAccountId returns the OtcBankAccountId field if non-nil, zero value otherwise.

### GetOtcBankAccountIdOk

`func (o *CreateBankAccountRequest) GetOtcBankAccountIdOk() (*string, bool)`

GetOtcBankAccountIdOk returns a tuple with the OtcBankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtcBankAccountId

`func (o *CreateBankAccountRequest) SetOtcBankAccountId(v string)`

SetOtcBankAccountId sets OtcBankAccountId field to given value.


### GetInfo

`func (o *CreateBankAccountRequest) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CreateBankAccountRequest) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CreateBankAccountRequest) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


