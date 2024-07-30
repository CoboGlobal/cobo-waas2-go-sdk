# DeleteWalletById201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Submitted** | **bool** | Whether the request to delete the wallet has been successfully submitted. - &#x60;true&#x60;: The request to delete the wallet has been successfully submitted. - &#x60;false&#x60;: The request to delete the wallet has not been submitted.  | 

## Methods

### NewDeleteWalletById201Response

`func NewDeleteWalletById201Response(submitted bool, ) *DeleteWalletById201Response`

NewDeleteWalletById201Response instantiates a new DeleteWalletById201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteWalletById201ResponseWithDefaults

`func NewDeleteWalletById201ResponseWithDefaults() *DeleteWalletById201Response`

NewDeleteWalletById201ResponseWithDefaults instantiates a new DeleteWalletById201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubmitted

`func (o *DeleteWalletById201Response) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *DeleteWalletById201Response) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *DeleteWalletById201Response) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


