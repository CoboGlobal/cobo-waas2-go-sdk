# CreateKyaScreeningsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Screenings** | [**[]KyaScreeningRequest**](KyaScreeningRequest.md) | List of address screening requests. Maximum 50 addresses per call. | 

## Methods

### NewCreateKyaScreeningsBody

`func NewCreateKyaScreeningsBody(screenings []KyaScreeningRequest, ) *CreateKyaScreeningsBody`

NewCreateKyaScreeningsBody instantiates a new CreateKyaScreeningsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKyaScreeningsBodyWithDefaults

`func NewCreateKyaScreeningsBodyWithDefaults() *CreateKyaScreeningsBody`

NewCreateKyaScreeningsBodyWithDefaults instantiates a new CreateKyaScreeningsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScreenings

`func (o *CreateKyaScreeningsBody) GetScreenings() []KyaScreeningRequest`

GetScreenings returns the Screenings field if non-nil, zero value otherwise.

### GetScreeningsOk

`func (o *CreateKyaScreeningsBody) GetScreeningsOk() (*[]KyaScreeningRequest, bool)`

GetScreeningsOk returns a tuple with the Screenings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenings

`func (o *CreateKyaScreeningsBody) SetScreenings(v []KyaScreeningRequest)`

SetScreenings sets Screenings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


