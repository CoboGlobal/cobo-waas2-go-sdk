# TSSEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | The event ID. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The time when the event occurred, in Unix timestamp format, measured in milliseconds. | [optional] 
**NodeId** | Pointer to **string** | The event publisher&#39;s TSS Node ID. | [optional] 
**EventType** | [**TSSEventType**](TSSEventType.md) |  | 
**Data** | Pointer to [**TSSEventData**](TSSEventData.md) |  | [optional] 

## Methods

### NewTSSEvent

`func NewTSSEvent(eventType TSSEventType, ) *TSSEvent`

NewTSSEvent instantiates a new TSSEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSEventWithDefaults

`func NewTSSEventWithDefaults() *TSSEvent`

NewTSSEventWithDefaults instantiates a new TSSEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *TSSEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *TSSEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *TSSEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *TSSEvent) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *TSSEvent) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TSSEvent) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TSSEvent) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TSSEvent) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetNodeId

`func (o *TSSEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TSSEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TSSEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *TSSEvent) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetEventType

`func (o *TSSEvent) GetEventType() TSSEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TSSEvent) GetEventTypeOk() (*TSSEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TSSEvent) SetEventType(v TSSEventType)`

SetEventType sets EventType field to given value.


### GetData

`func (o *TSSEvent) GetData() TSSEventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TSSEvent) GetDataOk() (*TSSEventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TSSEvent) SetData(v TSSEventData)`

SetData sets Data field to given value.

### HasData

`func (o *TSSEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


