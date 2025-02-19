# TravelRuleDepositNaturalEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedEntityType** | **string** | Specifies the type of entity associated with the transaction. - &#x60;LEGAL&#x60;: Legal entity. - &#x60;NATURAL&#x60;: Natural person.  | 
**FirstName** | **string** | The first name of the natural person. | 
**LastName** | **string** | The last name of the natural person. | 
**DateOfBirth** | Pointer to **string** | The date of birth of the natural person. This field is required when either of the following conditions is met: - &#x60;is_threshold_reached&#x60; is &#x60;true&#x60; in the response of the [Retrieve transaction limitations](https://www.cobo.com/developers/v2/api-references/travelrule/retrieve-transaction-limitations) operation. - &#x60;selected_entity_type&#x60; is &#x60;NATURAL&#x60;.  | [optional] 
**PlaceOfBirth** | Pointer to **string** | The place of birth of the natural person. This field is required when either of the following conditions is met: - &#x60;is_threshold_reached&#x60; is &#x60;true&#x60; in the response of the [Retrieve transaction limitations](https://www.cobo.com/developers/v2/api-references/travelrule/retrieve-transaction-limitations) operation. - &#x60;selected_entity_type&#x60; is &#x60;NATURAL&#x60;.  | [optional] 

## Methods

### NewTravelRuleDepositNaturalEntity

`func NewTravelRuleDepositNaturalEntity(selectedEntityType string, firstName string, lastName string, ) *TravelRuleDepositNaturalEntity`

NewTravelRuleDepositNaturalEntity instantiates a new TravelRuleDepositNaturalEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelRuleDepositNaturalEntityWithDefaults

`func NewTravelRuleDepositNaturalEntityWithDefaults() *TravelRuleDepositNaturalEntity`

NewTravelRuleDepositNaturalEntityWithDefaults instantiates a new TravelRuleDepositNaturalEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedEntityType

`func (o *TravelRuleDepositNaturalEntity) GetSelectedEntityType() string`

GetSelectedEntityType returns the SelectedEntityType field if non-nil, zero value otherwise.

### GetSelectedEntityTypeOk

`func (o *TravelRuleDepositNaturalEntity) GetSelectedEntityTypeOk() (*string, bool)`

GetSelectedEntityTypeOk returns a tuple with the SelectedEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedEntityType

`func (o *TravelRuleDepositNaturalEntity) SetSelectedEntityType(v string)`

SetSelectedEntityType sets SelectedEntityType field to given value.


### GetFirstName

`func (o *TravelRuleDepositNaturalEntity) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TravelRuleDepositNaturalEntity) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TravelRuleDepositNaturalEntity) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *TravelRuleDepositNaturalEntity) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TravelRuleDepositNaturalEntity) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TravelRuleDepositNaturalEntity) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDateOfBirth

`func (o *TravelRuleDepositNaturalEntity) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *TravelRuleDepositNaturalEntity) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *TravelRuleDepositNaturalEntity) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *TravelRuleDepositNaturalEntity) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *TravelRuleDepositNaturalEntity) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *TravelRuleDepositNaturalEntity) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *TravelRuleDepositNaturalEntity) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *TravelRuleDepositNaturalEntity) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


