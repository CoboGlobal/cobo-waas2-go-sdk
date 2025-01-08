# TravelRuleDepositRequestTravelRuleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationWalletType** | [**DestinationWalletType**](DestinationWalletType.md) |  | 
**VendorCode** | **string** | The vendor code for exchanges or VASPs. | 
**VendorVaspId** | **string** | The unique identifier of the VASP. | 
**VendorVaspName** | Pointer to **string** | The vendor name to be provided when selecting \&quot;Others\&quot; as the VASP case. This field allows customers to specify the name of a vendor not listed. | [optional] 
**EntityInfo** | [**TravelRuleDepositExchangesOrVASPEntityInfo**](TravelRuleDepositExchangesOrVASPEntityInfo.md) |  | 
**SelfCustodyWalletChallenge** | **string** | The challenge obtained from a previous operation. | 
**SelfCustodyWalletAddress** | **string** | The address of the self-custodial wallet. | 
**SelfCustodyWalletSign** | **string** | The signed message from the self-custodial wallet. | 

## Methods

### NewTravelRuleDepositRequestTravelRuleInfo

`func NewTravelRuleDepositRequestTravelRuleInfo(destinationWalletType DestinationWalletType, vendorCode string, vendorVaspId string, entityInfo TravelRuleDepositExchangesOrVASPEntityInfo, selfCustodyWalletChallenge string, selfCustodyWalletAddress string, selfCustodyWalletSign string, ) *TravelRuleDepositRequestTravelRuleInfo`

NewTravelRuleDepositRequestTravelRuleInfo instantiates a new TravelRuleDepositRequestTravelRuleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelRuleDepositRequestTravelRuleInfoWithDefaults

`func NewTravelRuleDepositRequestTravelRuleInfoWithDefaults() *TravelRuleDepositRequestTravelRuleInfo`

NewTravelRuleDepositRequestTravelRuleInfoWithDefaults instantiates a new TravelRuleDepositRequestTravelRuleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationWalletType

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetDestinationWalletType() DestinationWalletType`

GetDestinationWalletType returns the DestinationWalletType field if non-nil, zero value otherwise.

### GetDestinationWalletTypeOk

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetDestinationWalletTypeOk() (*DestinationWalletType, bool)`

GetDestinationWalletTypeOk returns a tuple with the DestinationWalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationWalletType

`func (o *TravelRuleDepositRequestTravelRuleInfo) SetDestinationWalletType(v DestinationWalletType)`

SetDestinationWalletType sets DestinationWalletType field to given value.


### GetVendorCode

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *TravelRuleDepositRequestTravelRuleInfo) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.


### GetVendorVaspId

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetVendorVaspId() string`

GetVendorVaspId returns the VendorVaspId field if non-nil, zero value otherwise.

### GetVendorVaspIdOk

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetVendorVaspIdOk() (*string, bool)`

GetVendorVaspIdOk returns a tuple with the VendorVaspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorVaspId

`func (o *TravelRuleDepositRequestTravelRuleInfo) SetVendorVaspId(v string)`

SetVendorVaspId sets VendorVaspId field to given value.


### GetVendorVaspName

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetVendorVaspName() string`

GetVendorVaspName returns the VendorVaspName field if non-nil, zero value otherwise.

### GetVendorVaspNameOk

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetVendorVaspNameOk() (*string, bool)`

GetVendorVaspNameOk returns a tuple with the VendorVaspName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorVaspName

`func (o *TravelRuleDepositRequestTravelRuleInfo) SetVendorVaspName(v string)`

SetVendorVaspName sets VendorVaspName field to given value.

### HasVendorVaspName

`func (o *TravelRuleDepositRequestTravelRuleInfo) HasVendorVaspName() bool`

HasVendorVaspName returns a boolean if a field has been set.

### GetEntityInfo

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetEntityInfo() TravelRuleDepositExchangesOrVASPEntityInfo`

GetEntityInfo returns the EntityInfo field if non-nil, zero value otherwise.

### GetEntityInfoOk

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetEntityInfoOk() (*TravelRuleDepositExchangesOrVASPEntityInfo, bool)`

GetEntityInfoOk returns a tuple with the EntityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityInfo

`func (o *TravelRuleDepositRequestTravelRuleInfo) SetEntityInfo(v TravelRuleDepositExchangesOrVASPEntityInfo)`

SetEntityInfo sets EntityInfo field to given value.


### GetSelfCustodyWalletChallenge

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetSelfCustodyWalletChallenge() string`

GetSelfCustodyWalletChallenge returns the SelfCustodyWalletChallenge field if non-nil, zero value otherwise.

### GetSelfCustodyWalletChallengeOk

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetSelfCustodyWalletChallengeOk() (*string, bool)`

GetSelfCustodyWalletChallengeOk returns a tuple with the SelfCustodyWalletChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletChallenge

`func (o *TravelRuleDepositRequestTravelRuleInfo) SetSelfCustodyWalletChallenge(v string)`

SetSelfCustodyWalletChallenge sets SelfCustodyWalletChallenge field to given value.


### GetSelfCustodyWalletAddress

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetSelfCustodyWalletAddress() string`

GetSelfCustodyWalletAddress returns the SelfCustodyWalletAddress field if non-nil, zero value otherwise.

### GetSelfCustodyWalletAddressOk

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetSelfCustodyWalletAddressOk() (*string, bool)`

GetSelfCustodyWalletAddressOk returns a tuple with the SelfCustodyWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletAddress

`func (o *TravelRuleDepositRequestTravelRuleInfo) SetSelfCustodyWalletAddress(v string)`

SetSelfCustodyWalletAddress sets SelfCustodyWalletAddress field to given value.


### GetSelfCustodyWalletSign

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetSelfCustodyWalletSign() string`

GetSelfCustodyWalletSign returns the SelfCustodyWalletSign field if non-nil, zero value otherwise.

### GetSelfCustodyWalletSignOk

`func (o *TravelRuleDepositRequestTravelRuleInfo) GetSelfCustodyWalletSignOk() (*string, bool)`

GetSelfCustodyWalletSignOk returns a tuple with the SelfCustodyWalletSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletSign

`func (o *TravelRuleDepositRequestTravelRuleInfo) SetSelfCustodyWalletSign(v string)`

SetSelfCustodyWalletSign sets SelfCustodyWalletSign field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


