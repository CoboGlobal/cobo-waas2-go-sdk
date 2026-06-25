# SOLBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseFee** | Pointer to **string** | A fixed fee charged per signature. The default is 5,000 lamports per signature. | [optional] 
**RentAmount** | Pointer to **string** | The one-time rent required to create and initialize a Solana token Associated Token Account (ATA) — a token sub-address that must be activated before the token can be received or used. This rent is paid by the main (source) address. It is populated only when an ATA must be activated for the transaction; otherwise it is null.  | [optional] 

## Methods

### NewSOLBase

`func NewSOLBase() *SOLBase`

NewSOLBase instantiates a new SOLBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSOLBaseWithDefaults

`func NewSOLBaseWithDefaults() *SOLBase`

NewSOLBaseWithDefaults instantiates a new SOLBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseFee

`func (o *SOLBase) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *SOLBase) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *SOLBase) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.

### HasBaseFee

`func (o *SOLBase) HasBaseFee() bool`

HasBaseFee returns a boolean if a field has been set.

### GetRentAmount

`func (o *SOLBase) GetRentAmount() string`

GetRentAmount returns the RentAmount field if non-nil, zero value otherwise.

### GetRentAmountOk

`func (o *SOLBase) GetRentAmountOk() (*string, bool)`

GetRentAmountOk returns a tuple with the RentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentAmount

`func (o *SOLBase) SetRentAmount(v string)`

SetRentAmount sets RentAmount field to given value.

### HasRentAmount

`func (o *SOLBase) HasRentAmount() bool`

HasRentAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


