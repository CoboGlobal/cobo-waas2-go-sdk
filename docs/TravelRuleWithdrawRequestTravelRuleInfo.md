# TravelRuleWithdrawRequestTravelRuleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationWalletType** | [**DestinationWalletType**](DestinationWalletType.md) |  | 
**SelfCustodyWalletChallenge** | **string** | The challenge obtained from a previous operation. | 
**SelfCustodyWalletAddress** | **string** | The address of the self-custodial wallet. | 
**SelfCustodyWalletSign** | **string** | The signed message from the self-custodial wallet. | 
**VendorCode** | **string** | The vendor code for exchanges or VASPs. | 
**VendorVaspId** | **string** | The unique identifier of the VASP. | 
**EntityInfo** | [**TravelRuleWithdrawExchangesOrVASPEntityInfo**](TravelRuleWithdrawExchangesOrVASPEntityInfo.md) |  | 

## Methods

### NewTravelRuleWithdrawRequestTravelRuleInfo

`func NewTravelRuleWithdrawRequestTravelRuleInfo(destinationWalletType DestinationWalletType, selfCustodyWalletChallenge string, selfCustodyWalletAddress string, selfCustodyWalletSign string, vendorCode string, vendorVaspId string, entityInfo TravelRuleWithdrawExchangesOrVASPEntityInfo, ) *TravelRuleWithdrawRequestTravelRuleInfo`

NewTravelRuleWithdrawRequestTravelRuleInfo instantiates a new TravelRuleWithdrawRequestTravelRuleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelRuleWithdrawRequestTravelRuleInfoWithDefaults

`func NewTravelRuleWithdrawRequestTravelRuleInfoWithDefaults() *TravelRuleWithdrawRequestTravelRuleInfo`

NewTravelRuleWithdrawRequestTravelRuleInfoWithDefaults instantiates a new TravelRuleWithdrawRequestTravelRuleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationWalletType

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetDestinationWalletType() DestinationWalletType`

GetDestinationWalletType returns the DestinationWalletType field if non-nil, zero value otherwise.

### GetDestinationWalletTypeOk

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetDestinationWalletTypeOk() (*DestinationWalletType, bool)`

GetDestinationWalletTypeOk returns a tuple with the DestinationWalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationWalletType

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) SetDestinationWalletType(v DestinationWalletType)`

SetDestinationWalletType sets DestinationWalletType field to given value.


### GetSelfCustodyWalletChallenge

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetSelfCustodyWalletChallenge() string`

GetSelfCustodyWalletChallenge returns the SelfCustodyWalletChallenge field if non-nil, zero value otherwise.

### GetSelfCustodyWalletChallengeOk

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetSelfCustodyWalletChallengeOk() (*string, bool)`

GetSelfCustodyWalletChallengeOk returns a tuple with the SelfCustodyWalletChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletChallenge

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) SetSelfCustodyWalletChallenge(v string)`

SetSelfCustodyWalletChallenge sets SelfCustodyWalletChallenge field to given value.


### GetSelfCustodyWalletAddress

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetSelfCustodyWalletAddress() string`

GetSelfCustodyWalletAddress returns the SelfCustodyWalletAddress field if non-nil, zero value otherwise.

### GetSelfCustodyWalletAddressOk

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetSelfCustodyWalletAddressOk() (*string, bool)`

GetSelfCustodyWalletAddressOk returns a tuple with the SelfCustodyWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletAddress

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) SetSelfCustodyWalletAddress(v string)`

SetSelfCustodyWalletAddress sets SelfCustodyWalletAddress field to given value.


### GetSelfCustodyWalletSign

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetSelfCustodyWalletSign() string`

GetSelfCustodyWalletSign returns the SelfCustodyWalletSign field if non-nil, zero value otherwise.

### GetSelfCustodyWalletSignOk

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetSelfCustodyWalletSignOk() (*string, bool)`

GetSelfCustodyWalletSignOk returns a tuple with the SelfCustodyWalletSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletSign

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) SetSelfCustodyWalletSign(v string)`

SetSelfCustodyWalletSign sets SelfCustodyWalletSign field to given value.


### GetVendorCode

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.


### GetVendorVaspId

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetVendorVaspId() string`

GetVendorVaspId returns the VendorVaspId field if non-nil, zero value otherwise.

### GetVendorVaspIdOk

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetVendorVaspIdOk() (*string, bool)`

GetVendorVaspIdOk returns a tuple with the VendorVaspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorVaspId

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) SetVendorVaspId(v string)`

SetVendorVaspId sets VendorVaspId field to given value.


### GetEntityInfo

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetEntityInfo() TravelRuleWithdrawExchangesOrVASPEntityInfo`

GetEntityInfo returns the EntityInfo field if non-nil, zero value otherwise.

### GetEntityInfoOk

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) GetEntityInfoOk() (*TravelRuleWithdrawExchangesOrVASPEntityInfo, bool)`

GetEntityInfoOk returns a tuple with the EntityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityInfo

`func (o *TravelRuleWithdrawRequestTravelRuleInfo) SetEntityInfo(v TravelRuleWithdrawExchangesOrVASPEntityInfo)`

SetEntityInfo sets EntityInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


