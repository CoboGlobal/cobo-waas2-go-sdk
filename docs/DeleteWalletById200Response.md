# DeleteWalletById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Submitted** | **bool** | Whether the request to delete the wallet has been successfully submitted. - &#x60;true&#x60;: The request to delete the wallet has been successfully submitted. - &#x60;false&#x60;: The request to delete the wallet has not been submitted.  | 

## Methods

### NewDeleteWalletById200Response

`func NewDeleteWalletById200Response(submitted bool, ) *DeleteWalletById200Response`

NewDeleteWalletById200Response instantiates a new DeleteWalletById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteWalletById200ResponseWithDefaults

`func NewDeleteWalletById200ResponseWithDefaults() *DeleteWalletById200Response`

NewDeleteWalletById200ResponseWithDefaults instantiates a new DeleteWalletById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubmitted

`func (o *DeleteWalletById200Response) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *DeleteWalletById200Response) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *DeleteWalletById200Response) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


