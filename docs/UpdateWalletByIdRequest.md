# UpdateWalletByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Apikey** | Pointer to **string** | The API Key for the exchange. | [optional] 
**Secret** | Pointer to **string** | The API Secret for the exchange. | [optional] 
**Passphrase** | Pointer to **string** | The passphrase for the exchange. | [optional] 
**GaCode** | Pointer to **string** | The ga_code for the exchange. | [optional] 
**SubAccountIds** | Pointer to **[]string** | The unique identifier associated with the exchange sub-account. It can be an email address, username, or a custom account ID. | [optional] 

## Methods

### NewUpdateWalletByIdRequest

`func NewUpdateWalletByIdRequest() *UpdateWalletByIdRequest`

NewUpdateWalletByIdRequest instantiates a new UpdateWalletByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWalletByIdRequestWithDefaults

`func NewUpdateWalletByIdRequestWithDefaults() *UpdateWalletByIdRequest`

NewUpdateWalletByIdRequestWithDefaults instantiates a new UpdateWalletByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWalletByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWalletByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWalletByIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWalletByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApikey

`func (o *UpdateWalletByIdRequest) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *UpdateWalletByIdRequest) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *UpdateWalletByIdRequest) SetApikey(v string)`

SetApikey sets Apikey field to given value.

### HasApikey

`func (o *UpdateWalletByIdRequest) HasApikey() bool`

HasApikey returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateWalletByIdRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateWalletByIdRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateWalletByIdRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateWalletByIdRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPassphrase

`func (o *UpdateWalletByIdRequest) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *UpdateWalletByIdRequest) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *UpdateWalletByIdRequest) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *UpdateWalletByIdRequest) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetGaCode

`func (o *UpdateWalletByIdRequest) GetGaCode() string`

GetGaCode returns the GaCode field if non-nil, zero value otherwise.

### GetGaCodeOk

`func (o *UpdateWalletByIdRequest) GetGaCodeOk() (*string, bool)`

GetGaCodeOk returns a tuple with the GaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaCode

`func (o *UpdateWalletByIdRequest) SetGaCode(v string)`

SetGaCode sets GaCode field to given value.

### HasGaCode

`func (o *UpdateWalletByIdRequest) HasGaCode() bool`

HasGaCode returns a boolean if a field has been set.

### GetSubAccountIds

`func (o *UpdateWalletByIdRequest) GetSubAccountIds() []string`

GetSubAccountIds returns the SubAccountIds field if non-nil, zero value otherwise.

### GetSubAccountIdsOk

`func (o *UpdateWalletByIdRequest) GetSubAccountIdsOk() (*[]string, bool)`

GetSubAccountIdsOk returns a tuple with the SubAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountIds

`func (o *UpdateWalletByIdRequest) SetSubAccountIds(v []string)`

SetSubAccountIds sets SubAccountIds field to given value.

### HasSubAccountIds

`func (o *UpdateWalletByIdRequest) HasSubAccountIds() bool`

HasSubAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


