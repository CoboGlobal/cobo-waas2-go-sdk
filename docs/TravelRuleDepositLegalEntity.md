# TravelRuleDepositLegalEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedEntityType** | **string** | The entity type. Possible values include: - &#x60;LEGAL&#x60;: Legal entity. - &#x60;NATURAL&#x60;: Natural person.  | 
**LegalName** | **string** | The legal name of the entity. | 
**DateOfIncorporation** | Pointer to **string** | The date of incorporation of the entity. This field is required when either of the following conditions is met: - &#x60;is_threshold_reached&#x60; is &#x60;true&#x60; in the response of the [Retrieve transaction limitations](https://www.cobo.com/developers/v2/api-references/travelrule/retrieve-transaction-limitations) operation. - &#x60;selected_entity_type&#x60; is &#x60;LEGAL&#x60;.  | [optional] 
**PlaceOfIncorporation** | Pointer to **string** | The place of incorporation of the entity. This field is required when either of the following conditions is met: - &#x60;is_threshold_reached&#x60; is &#x60;true&#x60; in the response of the [Retrieve transaction limitations](https://www.cobo.com/developers/v2/api-references/travelrule/retrieve-transaction-limitations) operation. - &#x60;selected_entity_type&#x60; is &#x60;LEGAL&#x60;.  | [optional] 

## Methods

### NewTravelRuleDepositLegalEntity

`func NewTravelRuleDepositLegalEntity(selectedEntityType string, legalName string, ) *TravelRuleDepositLegalEntity`

NewTravelRuleDepositLegalEntity instantiates a new TravelRuleDepositLegalEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelRuleDepositLegalEntityWithDefaults

`func NewTravelRuleDepositLegalEntityWithDefaults() *TravelRuleDepositLegalEntity`

NewTravelRuleDepositLegalEntityWithDefaults instantiates a new TravelRuleDepositLegalEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedEntityType

`func (o *TravelRuleDepositLegalEntity) GetSelectedEntityType() string`

GetSelectedEntityType returns the SelectedEntityType field if non-nil, zero value otherwise.

### GetSelectedEntityTypeOk

`func (o *TravelRuleDepositLegalEntity) GetSelectedEntityTypeOk() (*string, bool)`

GetSelectedEntityTypeOk returns a tuple with the SelectedEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedEntityType

`func (o *TravelRuleDepositLegalEntity) SetSelectedEntityType(v string)`

SetSelectedEntityType sets SelectedEntityType field to given value.


### GetLegalName

`func (o *TravelRuleDepositLegalEntity) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *TravelRuleDepositLegalEntity) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *TravelRuleDepositLegalEntity) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetDateOfIncorporation

`func (o *TravelRuleDepositLegalEntity) GetDateOfIncorporation() string`

GetDateOfIncorporation returns the DateOfIncorporation field if non-nil, zero value otherwise.

### GetDateOfIncorporationOk

`func (o *TravelRuleDepositLegalEntity) GetDateOfIncorporationOk() (*string, bool)`

GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfIncorporation

`func (o *TravelRuleDepositLegalEntity) SetDateOfIncorporation(v string)`

SetDateOfIncorporation sets DateOfIncorporation field to given value.

### HasDateOfIncorporation

`func (o *TravelRuleDepositLegalEntity) HasDateOfIncorporation() bool`

HasDateOfIncorporation returns a boolean if a field has been set.

### GetPlaceOfIncorporation

`func (o *TravelRuleDepositLegalEntity) GetPlaceOfIncorporation() string`

GetPlaceOfIncorporation returns the PlaceOfIncorporation field if non-nil, zero value otherwise.

### GetPlaceOfIncorporationOk

`func (o *TravelRuleDepositLegalEntity) GetPlaceOfIncorporationOk() (*string, bool)`

GetPlaceOfIncorporationOk returns a tuple with the PlaceOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfIncorporation

`func (o *TravelRuleDepositLegalEntity) SetPlaceOfIncorporation(v string)`

SetPlaceOfIncorporation sets PlaceOfIncorporation field to given value.

### HasPlaceOfIncorporation

`func (o *TravelRuleDepositLegalEntity) HasPlaceOfIncorporation() bool`

HasPlaceOfIncorporation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


