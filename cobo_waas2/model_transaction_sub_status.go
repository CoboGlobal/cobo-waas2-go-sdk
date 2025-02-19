/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// TransactionSubStatus The transaction sub-status. Possible values include:    - `PendingDoubleCheck`: The transaction is pending a double check.    - `RejectedDoubleCheck`: The transaction is rejected because it failed a double check.   - `PendingSpenderCheck`: The transaction is pending a spender check.   - `RejectedSpenderAuth`: The transaction is rejected by the spender.   - `PendingTravelRuleCheck`: The transaction is undergoing a Travel Rule check.   - `PendingTravelRuleInfo`: The transaction is awaiting users to provide information related to the Travel Rule.   - `RejectedTravelRule`: The transaction is rejected because it failed to comply with the Travel Rule.   - `RejectedTravelRuleDueToCompliance`: The transaction is rejected because it failed the cross-check of the Travel Rule.    - `RejectedTravelRuleDueToUnsupportedToken`: The transaction is rejected because the token is not supported by the Travel Rule.   - `PendingRiskControlCheck`: The transaction is pending for a Risk Control check.   - `PendingApproverCheck`: The transaction is pending approval from the approver.   - `RejectedApproverAuth`: The transaction is rejected by the approver.   - `RejectedbyMobileCosigner`: The transaction is rejected by a mobile cosigner.   - `Built`: The transaction has been built but not signed yet.   - `RejectedCoboCheck`: The transaction is rejected because it failed the internal check by Cobo.   - `RejectedWhiteList`: The transaction is rejected because the sender or receiver is not included in a whitelist.   - `PendingWaitSigner`: The transaction is pending signature.   - `PendingApprovalStart`: The transaction approval is waiting to be started.           - For [MPC Wallets (User-Controlled Wallets)](https://manuals.cobo.com/en/portal/mpc-wallets/ucw/introduction), you need to use the Client App and call the UCW SDK to start the transaction approval process.     - For [MPC Wallets (Organization-Controlled Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/introduction)):       - If you are using the [server co-signer](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups), this status indicates that the TSS Node will soon request the callback server to start the [risk controls](https://manuals.cobo.com/en/portal/risk-controls/introduction) check. No further action is required from you at this stage.       - If you are using the [mobile co-signer](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups), key share holders need to use their [Cobo Guard](https://manuals.cobo.com/en/guard/introduction) to approve the transaction and participate in the signing process.   - `FailedBySigner`: The transaction failed during the signing process.   - `SignatureVerificationSuccess`: The transaction's signature has been successfully verified.   - `SignatureVerificationFailed`: The transaction's signature failed verification.   - `FailedBroadcasting`: The transaction failed to be broadcast to the blockchain network.   - `CanceledBySpender`: The transaction is canceled by a Spender.   - `CanceledByAPI`: The transaction is canceled by a [Cancel transaction](https://www.cobo.com/developers/v2/api-references/transactions/cancel-transaction) operation.   - `Queue`: The transaction is queued to be processed by Cobo Portal.   - `Reverting`: The transaction is being reverted due to failure on the blockchain.   - `OnchainRejection`: The transaction is rejected from being added to the blockchain.   - `FailedOnChain`: The transaction failed on the blockchain.   - `PendingBlockConfirmations`: The transaction is awaiting the required number of confirmations.   - `ReplacedByNewTransaction`: The transaction has been replaced by a new transaction. 
type TransactionSubStatus string

// List of TransactionSubStatus
const (
	TRANSACTIONSUBSTATUS_REJECTED_KYT TransactionSubStatus = "RejectedKYT"
	TRANSACTIONSUBSTATUS_PENDING_DOUBLE_CHECK TransactionSubStatus = "PendingDoubleCheck"
	TRANSACTIONSUBSTATUS_PENDING_SPENDER_CHECK TransactionSubStatus = "PendingSpenderCheck"
	TRANSACTIONSUBSTATUS_PENDING_RISK_CONTROL_CHECK TransactionSubStatus = "PendingRiskControlCheck"
	TRANSACTIONSUBSTATUS_PENDING_APPROVER_CHECK TransactionSubStatus = "PendingApproverCheck"
	TRANSACTIONSUBSTATUS_REJECTED_COBO_CHECK TransactionSubStatus = "RejectedCoboCheck"
	TRANSACTIONSUBSTATUS_REJECTED_WHITE_LIST TransactionSubStatus = "RejectedWhiteList"
	TRANSACTIONSUBSTATUS_REJECTED_DOUBLE_CHECK TransactionSubStatus = "RejectedDoubleCheck"
	TRANSACTIONSUBSTATUS_REJECTED_SPENDER_AUTH TransactionSubStatus = "RejectedSpenderAuth"
	TRANSACTIONSUBSTATUS_REJECTED_RISK_CONTROL_CHECK TransactionSubStatus = "RejectedRiskControlCheck"
	TRANSACTIONSUBSTATUS_REJECTED_APPROVER_AUTH TransactionSubStatus = "RejectedApproverAuth"
	TRANSACTIONSUBSTATUS_REJECTEDBY_MOBILE_COSIGNER TransactionSubStatus = "RejectedbyMobileCosigner"
	TRANSACTIONSUBSTATUS_BUILT TransactionSubStatus = "Built"
	TRANSACTIONSUBSTATUS_PENDING_WAIT_SIGNER TransactionSubStatus = "PendingWaitSigner"
	TRANSACTIONSUBSTATUS_PENDING_APPROVAL_START TransactionSubStatus = "PendingApprovalStart"
	TRANSACTIONSUBSTATUS_FAILED_BY_SIGNER TransactionSubStatus = "FailedBySigner"
	TRANSACTIONSUBSTATUS_FAILED_BROADCASTING TransactionSubStatus = "FailedBroadcasting"
	TRANSACTIONSUBSTATUS_FAILED_ON_CHAIN TransactionSubStatus = "FailedOnChain"
	TRANSACTIONSUBSTATUS_REVERTING TransactionSubStatus = "Reverting"
	TRANSACTIONSUBSTATUS_QUEUE TransactionSubStatus = "Queue"
	TRANSACTIONSUBSTATUS_PENDING_BLOCK_CONFIRMATIONS TransactionSubStatus = "PendingBlockConfirmations"
	TRANSACTIONSUBSTATUS_REPLACED_BY_NEW_TRANSACTION TransactionSubStatus = "ReplacedByNewTransaction"
	TRANSACTIONSUBSTATUS_CANCELED_BY_SPENDER TransactionSubStatus = "CanceledBySpender"
	TRANSACTIONSUBSTATUS_CANCELED_BY_API TransactionSubStatus = "CanceledByAPI"
	TRANSACTIONSUBSTATUS_ONCHAIN_REJECTION TransactionSubStatus = "OnchainRejection"
	TRANSACTIONSUBSTATUS_REJECTED_TRAVEL_RULE TransactionSubStatus = "RejectedTravelRule"
	TRANSACTIONSUBSTATUS_REJECTED_TRAVEL_RULE_DUE_TO_COMPLIANCE TransactionSubStatus = "RejectedTravelRuleDueToCompliance"
	TRANSACTIONSUBSTATUS_PENDING_TRAVEL_RULE_INFO TransactionSubStatus = "PendingTravelRuleInfo"
	TRANSACTIONSUBSTATUS_PENDING_TRAVEL_RULE_CHECK TransactionSubStatus = "PendingTravelRuleCheck"
	TRANSACTIONSUBSTATUS_REJECTED_TRAVEL_RULE_DUE_TO_UNSUPPORTED_TOKEN TransactionSubStatus = "RejectedTravelRuleDueToUnsupportedToken"
	TRANSACTIONSUBSTATUS_SIGNATURE_VERIFICATION_SUCCESS TransactionSubStatus = "SignatureVerificationSuccess"
	TRANSACTIONSUBSTATUS_SIGNATURE_VERIFICATION_FAILED TransactionSubStatus = "SignatureVerificationFailed"
)

// All allowed values of TransactionSubStatus enum
var AllowedTransactionSubStatusEnumValues = []TransactionSubStatus{
	"RejectedKYT",
	"PendingDoubleCheck",
	"PendingSpenderCheck",
	"PendingRiskControlCheck",
	"PendingApproverCheck",
	"RejectedCoboCheck",
	"RejectedWhiteList",
	"RejectedDoubleCheck",
	"RejectedSpenderAuth",
	"RejectedRiskControlCheck",
	"RejectedApproverAuth",
	"RejectedbyMobileCosigner",
	"Built",
	"PendingWaitSigner",
	"PendingApprovalStart",
	"FailedBySigner",
	"FailedBroadcasting",
	"FailedOnChain",
	"Reverting",
	"Queue",
	"PendingBlockConfirmations",
	"ReplacedByNewTransaction",
	"CanceledBySpender",
	"CanceledByAPI",
	"OnchainRejection",
	"RejectedTravelRule",
	"RejectedTravelRuleDueToCompliance",
	"PendingTravelRuleInfo",
	"PendingTravelRuleCheck",
	"RejectedTravelRuleDueToUnsupportedToken",
	"SignatureVerificationSuccess",
	"SignatureVerificationFailed",
}

func (v *TransactionSubStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionSubStatus(value)
	for _, existing := range AllowedTransactionSubStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = TransactionSubStatus("unknown")
    return nil
}

// NewTransactionSubStatusFromValue returns a pointer to a valid TransactionSubStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionSubStatusFromValue(v string) (*TransactionSubStatus, error) {
	ev := TransactionSubStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionSubStatus: valid values are %v", v, AllowedTransactionSubStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionSubStatus) IsValid() bool {
	for _, existing := range AllowedTransactionSubStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionSubStatus value
func (v TransactionSubStatus) Ptr() *TransactionSubStatus {
	return &v
}

type NullableTransactionSubStatus struct {
	value *TransactionSubStatus
	isSet bool
}

func (v NullableTransactionSubStatus) Get() *TransactionSubStatus {
	return v.value
}

func (v *NullableTransactionSubStatus) Set(val *TransactionSubStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionSubStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionSubStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionSubStatus(val *TransactionSubStatus) *NullableTransactionSubStatus {
	return &NullableTransactionSubStatus{value: val, isSet: true}
}

func (v NullableTransactionSubStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionSubStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

