# SelfCustodyWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationWalletType** | [**DestinationWalletType**](DestinationWalletType.md) |  | 
**SelfCustodyWalletChallenge** | **string** | The challenge obtained from a previous operation. | 
**SelfCustodyWalletAddress** | **string** | The address of the self-custodial wallet. | 
**SelfCustodyWalletSign** | **string** | The signed message from the self-custodial wallet. | 

## Methods

### NewSelfCustodyWallet

`func NewSelfCustodyWallet(destinationWalletType DestinationWalletType, selfCustodyWalletChallenge string, selfCustodyWalletAddress string, selfCustodyWalletSign string, ) *SelfCustodyWallet`

NewSelfCustodyWallet instantiates a new SelfCustodyWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfCustodyWalletWithDefaults

`func NewSelfCustodyWalletWithDefaults() *SelfCustodyWallet`

NewSelfCustodyWalletWithDefaults instantiates a new SelfCustodyWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationWalletType

`func (o *SelfCustodyWallet) GetDestinationWalletType() DestinationWalletType`

GetDestinationWalletType returns the DestinationWalletType field if non-nil, zero value otherwise.

### GetDestinationWalletTypeOk

`func (o *SelfCustodyWallet) GetDestinationWalletTypeOk() (*DestinationWalletType, bool)`

GetDestinationWalletTypeOk returns a tuple with the DestinationWalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationWalletType

`func (o *SelfCustodyWallet) SetDestinationWalletType(v DestinationWalletType)`

SetDestinationWalletType sets DestinationWalletType field to given value.


### GetSelfCustodyWalletChallenge

`func (o *SelfCustodyWallet) GetSelfCustodyWalletChallenge() string`

GetSelfCustodyWalletChallenge returns the SelfCustodyWalletChallenge field if non-nil, zero value otherwise.

### GetSelfCustodyWalletChallengeOk

`func (o *SelfCustodyWallet) GetSelfCustodyWalletChallengeOk() (*string, bool)`

GetSelfCustodyWalletChallengeOk returns a tuple with the SelfCustodyWalletChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletChallenge

`func (o *SelfCustodyWallet) SetSelfCustodyWalletChallenge(v string)`

SetSelfCustodyWalletChallenge sets SelfCustodyWalletChallenge field to given value.


### GetSelfCustodyWalletAddress

`func (o *SelfCustodyWallet) GetSelfCustodyWalletAddress() string`

GetSelfCustodyWalletAddress returns the SelfCustodyWalletAddress field if non-nil, zero value otherwise.

### GetSelfCustodyWalletAddressOk

`func (o *SelfCustodyWallet) GetSelfCustodyWalletAddressOk() (*string, bool)`

GetSelfCustodyWalletAddressOk returns a tuple with the SelfCustodyWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletAddress

`func (o *SelfCustodyWallet) SetSelfCustodyWalletAddress(v string)`

SetSelfCustodyWalletAddress sets SelfCustodyWalletAddress field to given value.


### GetSelfCustodyWalletSign

`func (o *SelfCustodyWallet) GetSelfCustodyWalletSign() string`

GetSelfCustodyWalletSign returns the SelfCustodyWalletSign field if non-nil, zero value otherwise.

### GetSelfCustodyWalletSignOk

`func (o *SelfCustodyWallet) GetSelfCustodyWalletSignOk() (*string, bool)`

GetSelfCustodyWalletSignOk returns a tuple with the SelfCustodyWalletSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletSign

`func (o *SelfCustodyWallet) SetSelfCustodyWalletSign(v string)`

SetSelfCustodyWalletSign sets SelfCustodyWalletSign field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


