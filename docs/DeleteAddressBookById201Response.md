# DeleteAddressBookById201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Submitted** | **bool** | Whether the request to delete the address book has been successfully submitted. - &#x60;true&#x60;: The request to delete the address book has been successfully submitted. - &#x60;false&#x60;: The request to delete the address book has not been submitted.  | 

## Methods

### NewDeleteAddressBookById201Response

`func NewDeleteAddressBookById201Response(submitted bool, ) *DeleteAddressBookById201Response`

NewDeleteAddressBookById201Response instantiates a new DeleteAddressBookById201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAddressBookById201ResponseWithDefaults

`func NewDeleteAddressBookById201ResponseWithDefaults() *DeleteAddressBookById201Response`

NewDeleteAddressBookById201ResponseWithDefaults instantiates a new DeleteAddressBookById201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubmitted

`func (o *DeleteAddressBookById201Response) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *DeleteAddressBookById201Response) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *DeleteAddressBookById201Response) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


