# ReconDailyStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BizDate** | **string** | The business date (UTC), in YYYY-MM-DD format. | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. | 
**WalletId** | **string** | The wallet ID. | 
**OpeningBalance** | **string** | The opening balance at the start of the business date, expressed in the token&#39;s main unit. | 
**TotalDeposit** | **string** | The total deposit amount during the business date, expressed in the token&#39;s main unit. | 
**DepositCount** | **int32** | The number of deposits during the business date. | 
**TotalWithdrawal** | **string** | The total withdrawal amount during the business date, expressed in the token&#39;s main unit. | 
**WithdrawalCount** | **int32** | The number of withdrawals during the business date. | 
**ClosingBalance** | **string** | The closing balance at the end of the business date, expressed in the token&#39;s main unit. | 
**Status** | [**ReconStatementStatus**](ReconStatementStatus.md) |  | 

## Methods

### NewReconDailyStatement

`func NewReconDailyStatement(bizDate string, tokenId string, chainId string, walletId string, openingBalance string, totalDeposit string, depositCount int32, totalWithdrawal string, withdrawalCount int32, closingBalance string, status ReconStatementStatus, ) *ReconDailyStatement`

NewReconDailyStatement instantiates a new ReconDailyStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconDailyStatementWithDefaults

`func NewReconDailyStatementWithDefaults() *ReconDailyStatement`

NewReconDailyStatementWithDefaults instantiates a new ReconDailyStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBizDate

`func (o *ReconDailyStatement) GetBizDate() string`

GetBizDate returns the BizDate field if non-nil, zero value otherwise.

### GetBizDateOk

`func (o *ReconDailyStatement) GetBizDateOk() (*string, bool)`

GetBizDateOk returns a tuple with the BizDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizDate

`func (o *ReconDailyStatement) SetBizDate(v string)`

SetBizDate sets BizDate field to given value.


### GetTokenId

`func (o *ReconDailyStatement) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ReconDailyStatement) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ReconDailyStatement) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *ReconDailyStatement) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ReconDailyStatement) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ReconDailyStatement) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetWalletId

`func (o *ReconDailyStatement) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ReconDailyStatement) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ReconDailyStatement) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetOpeningBalance

`func (o *ReconDailyStatement) GetOpeningBalance() string`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *ReconDailyStatement) GetOpeningBalanceOk() (*string, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *ReconDailyStatement) SetOpeningBalance(v string)`

SetOpeningBalance sets OpeningBalance field to given value.


### GetTotalDeposit

`func (o *ReconDailyStatement) GetTotalDeposit() string`

GetTotalDeposit returns the TotalDeposit field if non-nil, zero value otherwise.

### GetTotalDepositOk

`func (o *ReconDailyStatement) GetTotalDepositOk() (*string, bool)`

GetTotalDepositOk returns a tuple with the TotalDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeposit

`func (o *ReconDailyStatement) SetTotalDeposit(v string)`

SetTotalDeposit sets TotalDeposit field to given value.


### GetDepositCount

`func (o *ReconDailyStatement) GetDepositCount() int32`

GetDepositCount returns the DepositCount field if non-nil, zero value otherwise.

### GetDepositCountOk

`func (o *ReconDailyStatement) GetDepositCountOk() (*int32, bool)`

GetDepositCountOk returns a tuple with the DepositCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositCount

`func (o *ReconDailyStatement) SetDepositCount(v int32)`

SetDepositCount sets DepositCount field to given value.


### GetTotalWithdrawal

`func (o *ReconDailyStatement) GetTotalWithdrawal() string`

GetTotalWithdrawal returns the TotalWithdrawal field if non-nil, zero value otherwise.

### GetTotalWithdrawalOk

`func (o *ReconDailyStatement) GetTotalWithdrawalOk() (*string, bool)`

GetTotalWithdrawalOk returns a tuple with the TotalWithdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithdrawal

`func (o *ReconDailyStatement) SetTotalWithdrawal(v string)`

SetTotalWithdrawal sets TotalWithdrawal field to given value.


### GetWithdrawalCount

`func (o *ReconDailyStatement) GetWithdrawalCount() int32`

GetWithdrawalCount returns the WithdrawalCount field if non-nil, zero value otherwise.

### GetWithdrawalCountOk

`func (o *ReconDailyStatement) GetWithdrawalCountOk() (*int32, bool)`

GetWithdrawalCountOk returns a tuple with the WithdrawalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalCount

`func (o *ReconDailyStatement) SetWithdrawalCount(v int32)`

SetWithdrawalCount sets WithdrawalCount field to given value.


### GetClosingBalance

`func (o *ReconDailyStatement) GetClosingBalance() string`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *ReconDailyStatement) GetClosingBalanceOk() (*string, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *ReconDailyStatement) SetClosingBalance(v string)`

SetClosingBalance sets ClosingBalance field to given value.


### GetStatus

`func (o *ReconDailyStatement) GetStatus() ReconStatementStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReconDailyStatement) GetStatusOk() (*ReconStatementStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReconDailyStatement) SetStatus(v ReconStatementStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


