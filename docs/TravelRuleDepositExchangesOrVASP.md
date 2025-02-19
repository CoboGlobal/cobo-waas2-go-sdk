# TravelRuleDepositExchangesOrVASP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationWalletType** | [**DestinationWalletType**](DestinationWalletType.md) |  | 
**VendorCode** | **string** | The vendor code of the VASP. | 
**VendorVaspId** | **string** | The unique identifier of the VASP. | 
**VendorVaspName** | Pointer to **string** | The vendor name. Use this field to specify the name of a vendor not listed. | [optional] 
**EntityInfo** | [**TravelRuleDepositExchangesOrVASPEntityInfo**](TravelRuleDepositExchangesOrVASPEntityInfo.md) |  | 

## Methods

### NewTravelRuleDepositExchangesOrVASP

`func NewTravelRuleDepositExchangesOrVASP(destinationWalletType DestinationWalletType, vendorCode string, vendorVaspId string, entityInfo TravelRuleDepositExchangesOrVASPEntityInfo, ) *TravelRuleDepositExchangesOrVASP`

NewTravelRuleDepositExchangesOrVASP instantiates a new TravelRuleDepositExchangesOrVASP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelRuleDepositExchangesOrVASPWithDefaults

`func NewTravelRuleDepositExchangesOrVASPWithDefaults() *TravelRuleDepositExchangesOrVASP`

NewTravelRuleDepositExchangesOrVASPWithDefaults instantiates a new TravelRuleDepositExchangesOrVASP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationWalletType

`func (o *TravelRuleDepositExchangesOrVASP) GetDestinationWalletType() DestinationWalletType`

GetDestinationWalletType returns the DestinationWalletType field if non-nil, zero value otherwise.

### GetDestinationWalletTypeOk

`func (o *TravelRuleDepositExchangesOrVASP) GetDestinationWalletTypeOk() (*DestinationWalletType, bool)`

GetDestinationWalletTypeOk returns a tuple with the DestinationWalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationWalletType

`func (o *TravelRuleDepositExchangesOrVASP) SetDestinationWalletType(v DestinationWalletType)`

SetDestinationWalletType sets DestinationWalletType field to given value.


### GetVendorCode

`func (o *TravelRuleDepositExchangesOrVASP) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *TravelRuleDepositExchangesOrVASP) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *TravelRuleDepositExchangesOrVASP) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.


### GetVendorVaspId

`func (o *TravelRuleDepositExchangesOrVASP) GetVendorVaspId() string`

GetVendorVaspId returns the VendorVaspId field if non-nil, zero value otherwise.

### GetVendorVaspIdOk

`func (o *TravelRuleDepositExchangesOrVASP) GetVendorVaspIdOk() (*string, bool)`

GetVendorVaspIdOk returns a tuple with the VendorVaspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorVaspId

`func (o *TravelRuleDepositExchangesOrVASP) SetVendorVaspId(v string)`

SetVendorVaspId sets VendorVaspId field to given value.


### GetVendorVaspName

`func (o *TravelRuleDepositExchangesOrVASP) GetVendorVaspName() string`

GetVendorVaspName returns the VendorVaspName field if non-nil, zero value otherwise.

### GetVendorVaspNameOk

`func (o *TravelRuleDepositExchangesOrVASP) GetVendorVaspNameOk() (*string, bool)`

GetVendorVaspNameOk returns a tuple with the VendorVaspName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorVaspName

`func (o *TravelRuleDepositExchangesOrVASP) SetVendorVaspName(v string)`

SetVendorVaspName sets VendorVaspName field to given value.

### HasVendorVaspName

`func (o *TravelRuleDepositExchangesOrVASP) HasVendorVaspName() bool`

HasVendorVaspName returns a boolean if a field has been set.

### GetEntityInfo

`func (o *TravelRuleDepositExchangesOrVASP) GetEntityInfo() TravelRuleDepositExchangesOrVASPEntityInfo`

GetEntityInfo returns the EntityInfo field if non-nil, zero value otherwise.

### GetEntityInfoOk

`func (o *TravelRuleDepositExchangesOrVASP) GetEntityInfoOk() (*TravelRuleDepositExchangesOrVASPEntityInfo, bool)`

GetEntityInfoOk returns a tuple with the EntityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityInfo

`func (o *TravelRuleDepositExchangesOrVASP) SetEntityInfo(v TravelRuleDepositExchangesOrVASPEntityInfo)`

SetEntityInfo sets EntityInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


