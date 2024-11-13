# CheckAddressChainsValidity200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. | 
**Validity** | **bool** | Whether the address is valid for the specified chain.  - &#x60;true&#x60;: The address is valid.  - &#x60;false&#x60;: The address is invalid.  | 

## Methods

### NewCheckAddressChainsValidity200ResponseInner

`func NewCheckAddressChainsValidity200ResponseInner(chainId string, validity bool, ) *CheckAddressChainsValidity200ResponseInner`

NewCheckAddressChainsValidity200ResponseInner instantiates a new CheckAddressChainsValidity200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAddressChainsValidity200ResponseInnerWithDefaults

`func NewCheckAddressChainsValidity200ResponseInnerWithDefaults() *CheckAddressChainsValidity200ResponseInner`

NewCheckAddressChainsValidity200ResponseInnerWithDefaults instantiates a new CheckAddressChainsValidity200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *CheckAddressChainsValidity200ResponseInner) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CheckAddressChainsValidity200ResponseInner) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CheckAddressChainsValidity200ResponseInner) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetValidity

`func (o *CheckAddressChainsValidity200ResponseInner) GetValidity() bool`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *CheckAddressChainsValidity200ResponseInner) GetValidityOk() (*bool, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *CheckAddressChainsValidity200ResponseInner) SetValidity(v bool)`

SetValidity sets Validity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


