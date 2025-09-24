# TokenizationUpdatePermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Action** | [**TokenizationPermissionAction**](TokenizationPermissionAction.md) |  | 
**Address** | **string** | The address to manage permissions for. | 
**Permissions** | [**[]TokenizationTokenPermissionType**](TokenizationTokenPermissionType.md) | The list of permissions to operate on. | 
**AppInitiator** | Pointer to **string** | The initiator of the tokenization activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 
**RequestId** | Pointer to **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | [optional] 

## Methods

### NewTokenizationUpdatePermissionsRequest

`func NewTokenizationUpdatePermissionsRequest(source TokenizationTokenOperationSource, action TokenizationPermissionAction, address string, permissions []TokenizationTokenPermissionType, fee TransactionRequestFee, ) *TokenizationUpdatePermissionsRequest`

NewTokenizationUpdatePermissionsRequest instantiates a new TokenizationUpdatePermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUpdatePermissionsRequestWithDefaults

`func NewTokenizationUpdatePermissionsRequestWithDefaults() *TokenizationUpdatePermissionsRequest`

NewTokenizationUpdatePermissionsRequestWithDefaults instantiates a new TokenizationUpdatePermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationUpdatePermissionsRequest) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationUpdatePermissionsRequest) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationUpdatePermissionsRequest) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetAction

`func (o *TokenizationUpdatePermissionsRequest) GetAction() TokenizationPermissionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenizationUpdatePermissionsRequest) GetActionOk() (*TokenizationPermissionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenizationUpdatePermissionsRequest) SetAction(v TokenizationPermissionAction)`

SetAction sets Action field to given value.


### GetAddress

`func (o *TokenizationUpdatePermissionsRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationUpdatePermissionsRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationUpdatePermissionsRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPermissions

`func (o *TokenizationUpdatePermissionsRequest) GetPermissions() []TokenizationTokenPermissionType`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationUpdatePermissionsRequest) GetPermissionsOk() (*[]TokenizationTokenPermissionType, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationUpdatePermissionsRequest) SetPermissions(v []TokenizationTokenPermissionType)`

SetPermissions sets Permissions field to given value.


### GetAppInitiator

`func (o *TokenizationUpdatePermissionsRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *TokenizationUpdatePermissionsRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *TokenizationUpdatePermissionsRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *TokenizationUpdatePermissionsRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.

### GetFee

`func (o *TokenizationUpdatePermissionsRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TokenizationUpdatePermissionsRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TokenizationUpdatePermissionsRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.


### GetRequestId

`func (o *TokenizationUpdatePermissionsRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TokenizationUpdatePermissionsRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TokenizationUpdatePermissionsRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TokenizationUpdatePermissionsRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


