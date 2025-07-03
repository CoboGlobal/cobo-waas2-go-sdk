# SOLBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseFee** | Pointer to **string** | The fundamental fee required for each transaction. It is charged to prevent spam transactions and network congestion, ensuring that only meaningful transactions consume network resources. | [optional] 
**RentAmount** | Pointer to **string** | The fee charged as rent for maintaining the state of accounts on the blockchain. This rent ensures accounts are stored on-chain over the long term and that there&#39;s sufficient balance to sustain the account state. | [optional] 

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


