# TravelRuleWithdrawExchangesOrVASPEntityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedEntityType** | [**SelectedEntityType**](SelectedEntityType.md) |  | 
**LegalName** | **string** | The legal name of the entity. | 
**DateOfIncorporation** | Pointer to **string** | The date of incorporation of the entity. | [optional] 
**PlaceOfIncorporation** | Pointer to **string** | The place of incorporation of the entity. | [optional] 
**FirstName** | **string** | The first name of the natural person. | 
**LastName** | **string** | The last name of the natural person. | 
**DateOfBirth** | Pointer to **string** | The date of birth of the natural person. | [optional] 
**PlaceOfBirth** | Pointer to **string** | The place of birth of the natural person. | [optional] 

## Methods

### NewTravelRuleWithdrawExchangesOrVASPEntityInfo

`func NewTravelRuleWithdrawExchangesOrVASPEntityInfo(selectedEntityType SelectedEntityType, legalName string, firstName string, lastName string, ) *TravelRuleWithdrawExchangesOrVASPEntityInfo`

NewTravelRuleWithdrawExchangesOrVASPEntityInfo instantiates a new TravelRuleWithdrawExchangesOrVASPEntityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelRuleWithdrawExchangesOrVASPEntityInfoWithDefaults

`func NewTravelRuleWithdrawExchangesOrVASPEntityInfoWithDefaults() *TravelRuleWithdrawExchangesOrVASPEntityInfo`

NewTravelRuleWithdrawExchangesOrVASPEntityInfoWithDefaults instantiates a new TravelRuleWithdrawExchangesOrVASPEntityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedEntityType

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetSelectedEntityType() SelectedEntityType`

GetSelectedEntityType returns the SelectedEntityType field if non-nil, zero value otherwise.

### GetSelectedEntityTypeOk

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetSelectedEntityTypeOk() (*SelectedEntityType, bool)`

GetSelectedEntityTypeOk returns a tuple with the SelectedEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedEntityType

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) SetSelectedEntityType(v SelectedEntityType)`

SetSelectedEntityType sets SelectedEntityType field to given value.


### GetLegalName

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetDateOfIncorporation

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetDateOfIncorporation() string`

GetDateOfIncorporation returns the DateOfIncorporation field if non-nil, zero value otherwise.

### GetDateOfIncorporationOk

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetDateOfIncorporationOk() (*string, bool)`

GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfIncorporation

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) SetDateOfIncorporation(v string)`

SetDateOfIncorporation sets DateOfIncorporation field to given value.

### HasDateOfIncorporation

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) HasDateOfIncorporation() bool`

HasDateOfIncorporation returns a boolean if a field has been set.

### GetPlaceOfIncorporation

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetPlaceOfIncorporation() string`

GetPlaceOfIncorporation returns the PlaceOfIncorporation field if non-nil, zero value otherwise.

### GetPlaceOfIncorporationOk

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetPlaceOfIncorporationOk() (*string, bool)`

GetPlaceOfIncorporationOk returns a tuple with the PlaceOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfIncorporation

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) SetPlaceOfIncorporation(v string)`

SetPlaceOfIncorporation sets PlaceOfIncorporation field to given value.

### HasPlaceOfIncorporation

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) HasPlaceOfIncorporation() bool`

HasPlaceOfIncorporation returns a boolean if a field has been set.

### GetFirstName

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDateOfBirth

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *TravelRuleWithdrawExchangesOrVASPEntityInfo) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


