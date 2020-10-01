// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run cmd/pdatagen/main.go".

package pdata

import (
	logsproto "go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/logs/v1"
)

// ResourceLogsSlice logically represents a slice of ResourceLogs.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceLogsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogsSlice struct {
	// orig points to the slice logsproto.ResourceLogs field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*logsproto.ResourceLogs
}

func newResourceLogsSlice(orig *[]*logsproto.ResourceLogs) ResourceLogsSlice {
	return ResourceLogsSlice{orig}
}

// NewResourceLogsSlice creates a ResourceLogsSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewResourceLogsSlice() ResourceLogsSlice {
	orig := []*logsproto.ResourceLogs(nil)
	return ResourceLogsSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewResourceLogsSlice()".
func (es ResourceLogsSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es ResourceLogsSlice) At(ix int) ResourceLogs {
	return newResourceLogs(&(*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ResourceLogsSlice) MoveAndAppendTo(dest ResourceLogsSlice) {
	if es.Len() == 0 {
		// Just to ensure that we always return a Slice with nil elements.
		*es.orig = nil
		return
	}
	if dest.Len() == 0 {
		*dest.orig = *es.orig
		*es.orig = nil
		return
	}
	*dest.orig = append(*dest.orig, *es.orig...)
	*es.orig = nil
	return
}

// CopyTo copies all elements from the current slice to the dest.
func (es ResourceLogsSlice) CopyTo(dest ResourceLogsSlice) {
	newLen := es.Len()
	if newLen == 0 {
		*dest.orig = []*logsproto.ResourceLogs(nil)
		return
	}
	oldLen := dest.Len()
	if newLen <= oldLen {
		(*dest.orig) = (*dest.orig)[:newLen]
		for i, el := range *es.orig {
			newResourceLogs(&el).CopyTo(newResourceLogs(&(*dest.orig)[i]))
		}
		return
	}
	origs := make([]logsproto.ResourceLogs, newLen)
	wrappers := make([]*logsproto.ResourceLogs, newLen)
	for i, el := range *es.orig {
		wrappers[i] = &origs[i]
		newResourceLogs(&el).CopyTo(newResourceLogs(&wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If newLen is 0 then the slice is replaced with a nil slice.
// 2. If the newLen <= len then equivalent with slice[0:newLen].
// 3. If the newLen > len then (newLen - len) empty elements will be appended to the slice.
//
// Here is how a new ResourceLogsSlice can be initialized:
// es := NewResourceLogsSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es ResourceLogsSlice) Resize(newLen int) {
	if newLen == 0 {
		(*es.orig) = []*logsproto.ResourceLogs(nil)
		return
	}
	oldLen := len(*es.orig)
	if newLen <= oldLen {
		(*es.orig) = (*es.orig)[:newLen]
		return
	}
	// TODO: Benchmark and optimize this logic.
	extraOrigs := make([]logsproto.ResourceLogs, newLen-oldLen)
	oldOrig := (*es.orig)
	for i := range extraOrigs {
		oldOrig = append(oldOrig, &extraOrigs[i])
	}
	(*es.orig) = oldOrig
}

// Append will increase the length of the ResourceLogsSlice by one and set the
// given ResourceLogs at that new position.  The original ResourceLogs
// could still be referenced so do not reuse it after passing it to this
// method.
func (es ResourceLogsSlice) Append(e *ResourceLogs) {
	(*es.orig) = append((*es.orig), *e.orig)
}

// ResourceLogs is a collection of logs from a Resource.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogs struct {
	// orig points to the pointer logsproto.ResourceLogs field contained somewhere else.
	// We use pointer-to-pointer to be able to modify it in InitEmpty func.
	orig **logsproto.ResourceLogs
}

func newResourceLogs(orig **logsproto.ResourceLogs) ResourceLogs {
	return ResourceLogs{orig}
}

// NewResourceLogs creates a new "nil" ResourceLogs.
// To initialize the struct call "InitEmpty".
//
// This must be used only in testing code since no "Set" method available.
func NewResourceLogs() ResourceLogs {
	orig := (*logsproto.ResourceLogs)(nil)
	return newResourceLogs(&orig)
}

// InitEmpty overwrites the current value with empty.
func (ms ResourceLogs) InitEmpty() {
	*ms.orig = &logsproto.ResourceLogs{}
}

// IsNil returns true if the underlying data are nil.
//
// Important: All other functions will cause a runtime error if this returns "true".
func (ms ResourceLogs) IsNil() bool {
	return *ms.orig == nil
}

// Resource returns the resource associated with this ResourceLogs.
// If no resource available, it creates an empty message and associates it with this ResourceLogs.
//
//  Empty initialized ResourceLogs will return "nil" Resource.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms ResourceLogs) Resource() Resource {
	return newResource(&(*ms.orig).Resource)
}

// Logs returns the Logs associated with this ResourceLogs.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms ResourceLogs) Logs() LogSlice {
	return newLogSlice(&(*ms.orig).Logs)
}

// CopyTo copies all properties from the current struct to the dest.
func (ms ResourceLogs) CopyTo(dest ResourceLogs) {
	if ms.IsNil() {
		*dest.orig = nil
		return
	}
	if dest.IsNil() {
		dest.InitEmpty()
	}
	ms.Resource().CopyTo(dest.Resource())
	ms.Logs().CopyTo(dest.Logs())
}

// LogSlice logically represents a slice of LogRecord.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewLogSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogSlice struct {
	// orig points to the slice logsproto.LogRecord field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*logsproto.LogRecord
}

func newLogSlice(orig *[]*logsproto.LogRecord) LogSlice {
	return LogSlice{orig}
}

// NewLogSlice creates a LogSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewLogSlice() LogSlice {
	orig := []*logsproto.LogRecord(nil)
	return LogSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewLogSlice()".
func (es LogSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es LogSlice) At(ix int) LogRecord {
	return newLogRecord(&(*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es LogSlice) MoveAndAppendTo(dest LogSlice) {
	if es.Len() == 0 {
		// Just to ensure that we always return a Slice with nil elements.
		*es.orig = nil
		return
	}
	if dest.Len() == 0 {
		*dest.orig = *es.orig
		*es.orig = nil
		return
	}
	*dest.orig = append(*dest.orig, *es.orig...)
	*es.orig = nil
	return
}

// CopyTo copies all elements from the current slice to the dest.
func (es LogSlice) CopyTo(dest LogSlice) {
	newLen := es.Len()
	if newLen == 0 {
		*dest.orig = []*logsproto.LogRecord(nil)
		return
	}
	oldLen := dest.Len()
	if newLen <= oldLen {
		(*dest.orig) = (*dest.orig)[:newLen]
		for i, el := range *es.orig {
			newLogRecord(&el).CopyTo(newLogRecord(&(*dest.orig)[i]))
		}
		return
	}
	origs := make([]logsproto.LogRecord, newLen)
	wrappers := make([]*logsproto.LogRecord, newLen)
	for i, el := range *es.orig {
		wrappers[i] = &origs[i]
		newLogRecord(&el).CopyTo(newLogRecord(&wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If newLen is 0 then the slice is replaced with a nil slice.
// 2. If the newLen <= len then equivalent with slice[0:newLen].
// 3. If the newLen > len then (newLen - len) empty elements will be appended to the slice.
//
// Here is how a new LogSlice can be initialized:
// es := NewLogSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es LogSlice) Resize(newLen int) {
	if newLen == 0 {
		(*es.orig) = []*logsproto.LogRecord(nil)
		return
	}
	oldLen := len(*es.orig)
	if newLen <= oldLen {
		(*es.orig) = (*es.orig)[:newLen]
		return
	}
	// TODO: Benchmark and optimize this logic.
	extraOrigs := make([]logsproto.LogRecord, newLen-oldLen)
	oldOrig := (*es.orig)
	for i := range extraOrigs {
		oldOrig = append(oldOrig, &extraOrigs[i])
	}
	(*es.orig) = oldOrig
}

// Append will increase the length of the LogSlice by one and set the
// given LogRecord at that new position.  The original LogRecord
// could still be referenced so do not reuse it after passing it to this
// method.
func (es LogSlice) Append(e *LogRecord) {
	(*es.orig) = append((*es.orig), *e.orig)
}

// LogRecord are experimental implementation of OpenTelemetry Log Data Model.

//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewLogRecord function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogRecord struct {
	// orig points to the pointer logsproto.LogRecord field contained somewhere else.
	// We use pointer-to-pointer to be able to modify it in InitEmpty func.
	orig **logsproto.LogRecord
}

func newLogRecord(orig **logsproto.LogRecord) LogRecord {
	return LogRecord{orig}
}

// NewLogRecord creates a new "nil" LogRecord.
// To initialize the struct call "InitEmpty".
//
// This must be used only in testing code since no "Set" method available.
func NewLogRecord() LogRecord {
	orig := (*logsproto.LogRecord)(nil)
	return newLogRecord(&orig)
}

// InitEmpty overwrites the current value with empty.
func (ms LogRecord) InitEmpty() {
	*ms.orig = &logsproto.LogRecord{}
}

// IsNil returns true if the underlying data are nil.
//
// Important: All other functions will cause a runtime error if this returns "true".
func (ms LogRecord) IsNil() bool {
	return *ms.orig == nil
}

// Timestamp returns the timestamp associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) Timestamp() TimestampUnixNano {
	return TimestampUnixNano((*ms.orig).TimestampUnixNano)
}

// SetTimestamp replaces the timestamp associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SetTimestamp(v TimestampUnixNano) {
	(*ms.orig).TimestampUnixNano = uint64(v)
}

// TraceID returns the traceid associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) TraceID() TraceID {
	return TraceID((*ms.orig).TraceId)
}

// SetTraceID replaces the traceid associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SetTraceID(v TraceID) {
	(*ms.orig).TraceId = []byte(v)
}

// SpanID returns the spanid associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SpanID() SpanID {
	return SpanID((*ms.orig).SpanId)
}

// SetSpanID replaces the spanid associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SetSpanID(v SpanID) {
	(*ms.orig).SpanId = []byte(v)
}

// Flags returns the flags associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) Flags() uint32 {
	return uint32((*ms.orig).Flags)
}

// SetFlags replaces the flags associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SetFlags(v uint32) {
	(*ms.orig).Flags = uint32(v)
}

// SeverityText returns the severitytext associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SeverityText() string {
	return (*ms.orig).SeverityText
}

// SetSeverityText replaces the severitytext associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SetSeverityText(v string) {
	(*ms.orig).SeverityText = v
}

// SeverityNumber returns the severitynumber associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SeverityNumber() logsproto.SeverityNumber {
	return logsproto.SeverityNumber((*ms.orig).SeverityNumber)
}

// SetSeverityNumber replaces the severitynumber associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SetSeverityNumber(v logsproto.SeverityNumber) {
	(*ms.orig).SeverityNumber = logsproto.SeverityNumber(v)
}

// ShortName returns the shortname associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) ShortName() string {
	return (*ms.orig).ShortName
}

// SetShortName replaces the shortname associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SetShortName(v string) {
	(*ms.orig).ShortName = v
}

// Body returns the body associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) Body() string {
	return (*ms.orig).Body
}

// SetBody replaces the body associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SetBody(v string) {
	(*ms.orig).Body = v
}

// Attributes returns the Attributes associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) Attributes() AttributeMap {
	return newAttributeMap(&(*ms.orig).Attributes)
}

// DroppedAttributesCount returns the droppedattributescount associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) DroppedAttributesCount() uint32 {
	return (*ms.orig).DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this LogRecord.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms LogRecord) SetDroppedAttributesCount(v uint32) {
	(*ms.orig).DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct to the dest.
func (ms LogRecord) CopyTo(dest LogRecord) {
	if ms.IsNil() {
		*dest.orig = nil
		return
	}
	if dest.IsNil() {
		dest.InitEmpty()
	}
	dest.SetTimestamp(ms.Timestamp())
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
	dest.SetFlags(ms.Flags())
	dest.SetSeverityText(ms.SeverityText())
	dest.SetSeverityNumber(ms.SeverityNumber())
	dest.SetShortName(ms.ShortName())
	dest.SetBody(ms.Body())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}