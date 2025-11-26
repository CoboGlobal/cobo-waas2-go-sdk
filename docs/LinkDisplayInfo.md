# LinkDisplayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeveloperName** | Pointer to **string** | The brand name to display to end users on the page. This helps users identify who is providing the payment service.  | [optional] 
**Logo** | Pointer to **string** | The URL of the logo image to display to end users on the page.  | [optional] 

## Methods

### NewLinkDisplayInfo

`func NewLinkDisplayInfo() *LinkDisplayInfo`

NewLinkDisplayInfo instantiates a new LinkDisplayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkDisplayInfoWithDefaults

`func NewLinkDisplayInfoWithDefaults() *LinkDisplayInfo`

NewLinkDisplayInfoWithDefaults instantiates a new LinkDisplayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeveloperName

`func (o *LinkDisplayInfo) GetDeveloperName() string`

GetDeveloperName returns the DeveloperName field if non-nil, zero value otherwise.

### GetDeveloperNameOk

`func (o *LinkDisplayInfo) GetDeveloperNameOk() (*string, bool)`

GetDeveloperNameOk returns a tuple with the DeveloperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperName

`func (o *LinkDisplayInfo) SetDeveloperName(v string)`

SetDeveloperName sets DeveloperName field to given value.

### HasDeveloperName

`func (o *LinkDisplayInfo) HasDeveloperName() bool`

HasDeveloperName returns a boolean if a field has been set.

### GetLogo

`func (o *LinkDisplayInfo) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *LinkDisplayInfo) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *LinkDisplayInfo) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *LinkDisplayInfo) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


