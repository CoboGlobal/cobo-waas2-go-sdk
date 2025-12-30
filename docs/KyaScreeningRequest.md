# KyaScreeningRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. This ID is used for idempotency to prevent duplicate submissions and for troubleshooting purposes. | 
**Address** | **string** | The blockchain address to be screened. For chains requiring memo (XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, TON), append the memo to the address using \&quot;|\&quot; separator (e.g., \&quot;address|memo\&quot;).  | 
**ChainId** | **string** | The chain identifier (e.g., ETH, BTC, TRON). | 
**Note** | Pointer to **string** | Optional note for this address screening. | [optional] 

## Methods

### NewKyaScreeningRequest

`func NewKyaScreeningRequest(requestId string, address string, chainId string, ) *KyaScreeningRequest`

NewKyaScreeningRequest instantiates a new KyaScreeningRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKyaScreeningRequestWithDefaults

`func NewKyaScreeningRequestWithDefaults() *KyaScreeningRequest`

NewKyaScreeningRequestWithDefaults instantiates a new KyaScreeningRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *KyaScreeningRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *KyaScreeningRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *KyaScreeningRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetAddress

`func (o *KyaScreeningRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *KyaScreeningRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *KyaScreeningRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *KyaScreeningRequest) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *KyaScreeningRequest) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *KyaScreeningRequest) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetNote

`func (o *KyaScreeningRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *KyaScreeningRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *KyaScreeningRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *KyaScreeningRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


