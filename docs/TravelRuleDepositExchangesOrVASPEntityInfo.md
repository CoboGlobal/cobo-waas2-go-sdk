# TravelRuleDepositExchangesOrVASPEntityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedEntityType** | **string** | Specifies the type of entity associated with the transaction. | 
**LegalName** | **string** | The legal name of the entity. | 
**DateOfIncorporation** | Pointer to **string** | The incorporation date of the entity. This field is required when: - **Calling**: &#x60;travel_rule/transaction/limitation&#x60; API returns &#x60;is_threshold_reached &#x3D; true&#x60;. - **Entity Type**: LEGAL. Otherwise, this field can be omitted.  | [optional] 
**PlaceOfIncorporation** | Pointer to **string** | The place of incorporation of the entity. This field is required when: - **Calling**: &#x60;travel_rule/transaction/limitation&#x60; API returns &#x60;is_threshold_reached &#x3D; true&#x60;. - **Entity Type**: LEGAL. Otherwise, this field can be omitted.  | [optional] 
**FirstName** | **string** | The first name of the user. | 
**LastName** | **string** | The last name of the user. | 
**DateOfBirth** | Pointer to **string** | The date of birth of the user. This field is required when: - **Calling**: &#x60;travel_rule/transaction/limitation&#x60; API returns &#x60;is_threshold_reached &#x3D; true&#x60;. - **Entity Type**: NATURAL. Otherwise, this field can be omitted.  | [optional] 
**PlaceOfBirth** | Pointer to **string** | The place of birth of the user. This field is required when: - **Calling**: &#x60;travel_rule/transaction/limitation&#x60; API returns &#x60;is_threshold_reached &#x3D; true&#x60;. - **Entity Type**: NATURAL. Otherwise, this field can be omitted.  | [optional] 

## Methods

### NewTravelRuleDepositExchangesOrVASPEntityInfo

`func NewTravelRuleDepositExchangesOrVASPEntityInfo(selectedEntityType string, legalName string, firstName string, lastName string, ) *TravelRuleDepositExchangesOrVASPEntityInfo`

NewTravelRuleDepositExchangesOrVASPEntityInfo instantiates a new TravelRuleDepositExchangesOrVASPEntityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelRuleDepositExchangesOrVASPEntityInfoWithDefaults

`func NewTravelRuleDepositExchangesOrVASPEntityInfoWithDefaults() *TravelRuleDepositExchangesOrVASPEntityInfo`

NewTravelRuleDepositExchangesOrVASPEntityInfoWithDefaults instantiates a new TravelRuleDepositExchangesOrVASPEntityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedEntityType

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetSelectedEntityType() string`

GetSelectedEntityType returns the SelectedEntityType field if non-nil, zero value otherwise.

### GetSelectedEntityTypeOk

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetSelectedEntityTypeOk() (*string, bool)`

GetSelectedEntityTypeOk returns a tuple with the SelectedEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedEntityType

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) SetSelectedEntityType(v string)`

SetSelectedEntityType sets SelectedEntityType field to given value.


### GetLegalName

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetDateOfIncorporation

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetDateOfIncorporation() string`

GetDateOfIncorporation returns the DateOfIncorporation field if non-nil, zero value otherwise.

### GetDateOfIncorporationOk

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetDateOfIncorporationOk() (*string, bool)`

GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfIncorporation

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) SetDateOfIncorporation(v string)`

SetDateOfIncorporation sets DateOfIncorporation field to given value.

### HasDateOfIncorporation

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) HasDateOfIncorporation() bool`

HasDateOfIncorporation returns a boolean if a field has been set.

### GetPlaceOfIncorporation

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetPlaceOfIncorporation() string`

GetPlaceOfIncorporation returns the PlaceOfIncorporation field if non-nil, zero value otherwise.

### GetPlaceOfIncorporationOk

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetPlaceOfIncorporationOk() (*string, bool)`

GetPlaceOfIncorporationOk returns a tuple with the PlaceOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfIncorporation

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) SetPlaceOfIncorporation(v string)`

SetPlaceOfIncorporation sets PlaceOfIncorporation field to given value.

### HasPlaceOfIncorporation

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) HasPlaceOfIncorporation() bool`

HasPlaceOfIncorporation returns a boolean if a field has been set.

### GetFirstName

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDateOfBirth

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *TravelRuleDepositExchangesOrVASPEntityInfo) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


