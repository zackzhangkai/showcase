// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by thriftgo (0.1.7). DO NOT EDIT.

package meta

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type TMessageType int64

const (
	TMessageType_INVALID_MESSAGE_TYPE TMessageType = 0
	TMessageType_CALL                 TMessageType = 1
	TMessageType_REPLY                TMessageType = 2
	TMessageType_EXCEPTION            TMessageType = 3
	TMessageType_ONEWAY               TMessageType = 4
)

func (p TMessageType) String() string {
	switch p {
	case TMessageType_INVALID_MESSAGE_TYPE:
		return "INVALID_MESSAGE_TYPE"
	case TMessageType_CALL:
		return "CALL"
	case TMessageType_REPLY:
		return "REPLY"
	case TMessageType_EXCEPTION:
		return "EXCEPTION"
	case TMessageType_ONEWAY:
		return "ONEWAY"
	}
	return "<UNSET>"
}

func TMessageTypeFromString(s string) (TMessageType, error) {
	switch s {
	case "INVALID_MESSAGE_TYPE":
		return TMessageType_INVALID_MESSAGE_TYPE, nil
	case "CALL":
		return TMessageType_CALL, nil
	case "REPLY":
		return TMessageType_REPLY, nil
	case "EXCEPTION":
		return TMessageType_EXCEPTION, nil
	case "ONEWAY":
		return TMessageType_ONEWAY, nil
	}
	return TMessageType(0), fmt.Errorf("not a valid TMessageType string")
}

func TMessageTypePtr(v TMessageType) *TMessageType { return &v }
func (p *TMessageType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = TMessageType(result.Int64)
	return
}

func (p *TMessageType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type TTypeID int64

const (
	TTypeID_STOP   TTypeID = 0
	TTypeID_VOID   TTypeID = 1
	TTypeID_BOOL   TTypeID = 2
	TTypeID_BYTE   TTypeID = 3
	TTypeID_DOUBLE TTypeID = 4
	TTypeID_I16    TTypeID = 6
	TTypeID_I32    TTypeID = 8
	TTypeID_I64    TTypeID = 10
	TTypeID_STRING TTypeID = 11
	TTypeID_STRUCT TTypeID = 12
	TTypeID_MAP    TTypeID = 13
	TTypeID_SET    TTypeID = 14
	TTypeID_LIST   TTypeID = 15
	TTypeID_UTF8   TTypeID = 16
	TTypeID_UTF16  TTypeID = 17
)

func (p TTypeID) String() string {
	switch p {
	case TTypeID_STOP:
		return "STOP"
	case TTypeID_VOID:
		return "VOID"
	case TTypeID_BOOL:
		return "BOOL"
	case TTypeID_BYTE:
		return "BYTE"
	case TTypeID_DOUBLE:
		return "DOUBLE"
	case TTypeID_I16:
		return "I16"
	case TTypeID_I32:
		return "I32"
	case TTypeID_I64:
		return "I64"
	case TTypeID_STRING:
		return "STRING"
	case TTypeID_STRUCT:
		return "STRUCT"
	case TTypeID_MAP:
		return "MAP"
	case TTypeID_SET:
		return "SET"
	case TTypeID_LIST:
		return "LIST"
	case TTypeID_UTF8:
		return "UTF8"
	case TTypeID_UTF16:
		return "UTF16"
	}
	return "<UNSET>"
}

func TTypeIDFromString(s string) (TTypeID, error) {
	switch s {
	case "STOP":
		return TTypeID_STOP, nil
	case "VOID":
		return TTypeID_VOID, nil
	case "BOOL":
		return TTypeID_BOOL, nil
	case "BYTE":
		return TTypeID_BYTE, nil
	case "DOUBLE":
		return TTypeID_DOUBLE, nil
	case "I16":
		return TTypeID_I16, nil
	case "I32":
		return TTypeID_I32, nil
	case "I64":
		return TTypeID_I64, nil
	case "STRING":
		return TTypeID_STRING, nil
	case "STRUCT":
		return TTypeID_STRUCT, nil
	case "MAP":
		return TTypeID_MAP, nil
	case "SET":
		return TTypeID_SET, nil
	case "LIST":
		return TTypeID_LIST, nil
	case "UTF8":
		return TTypeID_UTF8, nil
	case "UTF16":
		return TTypeID_UTF16, nil
	}
	return TTypeID(0), fmt.Errorf("not a valid TTypeID string")
}

func TTypeIDPtr(v TTypeID) *TTypeID { return &v }
func (p *TTypeID) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = TTypeID(result.Int64)
	return
}

func (p *TTypeID) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type TRequiredness int64

const (
	TRequiredness_DEFAULT  TRequiredness = 0
	TRequiredness_REQUIRED TRequiredness = 1
	TRequiredness_OPTIONAL TRequiredness = 2
)

func (p TRequiredness) String() string {
	switch p {
	case TRequiredness_DEFAULT:
		return "DEFAULT"
	case TRequiredness_REQUIRED:
		return "REQUIRED"
	case TRequiredness_OPTIONAL:
		return "OPTIONAL"
	}
	return "<UNSET>"
}

func TRequirednessFromString(s string) (TRequiredness, error) {
	switch s {
	case "DEFAULT":
		return TRequiredness_DEFAULT, nil
	case "REQUIRED":
		return TRequiredness_REQUIRED, nil
	case "OPTIONAL":
		return TRequiredness_OPTIONAL, nil
	}
	return TRequiredness(0), fmt.Errorf("not a valid TRequiredness string")
}

func TRequirednessPtr(v TRequiredness) *TRequiredness { return &v }
func (p *TRequiredness) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = TRequiredness(result.Int64)
	return
}

func (p *TRequiredness) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type TypeMeta struct {
	TypeID    TTypeID   `thrift:"type_id,1,required" json:"type_id"`
	KeyType   *TypeMeta `thrift:"key_type,2,optional" json:"key_type,omitempty"`
	ValueType *TypeMeta `thrift:"value_type,3,optional" json:"value_type,omitempty"`
}

func NewTypeMeta() *TypeMeta {
	return &TypeMeta{}
}

func (p *TypeMeta) GetTypeID() (v TTypeID) {
	return p.TypeID
}

var TypeMeta_KeyType_DEFAULT *TypeMeta

func (p *TypeMeta) GetKeyType() (v *TypeMeta) {
	if !p.IsSetKeyType() {
		return TypeMeta_KeyType_DEFAULT
	}
	return p.KeyType
}

var TypeMeta_ValueType_DEFAULT *TypeMeta

func (p *TypeMeta) GetValueType() (v *TypeMeta) {
	if !p.IsSetValueType() {
		return TypeMeta_ValueType_DEFAULT
	}
	return p.ValueType
}

func (p *TypeMeta) IsSetKeyType() bool {
	return p.KeyType != nil
}

func (p *TypeMeta) IsSetValueType() bool {
	return p.ValueType != nil
}

func (p *TypeMeta) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TypeMeta(%+v)", *p)
}

type FieldMeta struct {
	FieldID      int16         `thrift:"field_id,1,required" json:"field_id"`
	Name         string        `thrift:"name,2,required" json:"name"`
	Requiredness TRequiredness `thrift:"requiredness,3,required" json:"requiredness"`
	FieldType    *TypeMeta     `thrift:"field_type,4,required" json:"field_type"`
}

func NewFieldMeta() *FieldMeta {
	return &FieldMeta{}
}

func (p *FieldMeta) GetFieldID() (v int16) {
	return p.FieldID
}

func (p *FieldMeta) GetName() (v string) {
	return p.Name
}

func (p *FieldMeta) GetRequiredness() (v TRequiredness) {
	return p.Requiredness
}

var FieldMeta_FieldType_DEFAULT *TypeMeta

func (p *FieldMeta) GetFieldType() (v *TypeMeta) {
	if !p.IsSetFieldType() {
		return FieldMeta_FieldType_DEFAULT
	}
	return p.FieldType
}

func (p *FieldMeta) IsSetFieldType() bool {
	return p.FieldType != nil
}

func (p *FieldMeta) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FieldMeta(%+v)", *p)
}

type StructMeta struct {
	Name     string       `thrift:"name,1,required" json:"name"`
	Category string       `thrift:"category,2,required" json:"category"`
	Fields   []*FieldMeta `thrift:"fields,3,required" json:"fields"`
}

func NewStructMeta() *StructMeta {
	return &StructMeta{}
}

func (p *StructMeta) GetName() (v string) {
	return p.Name
}

func (p *StructMeta) GetCategory() (v string) {
	return p.Category
}

func (p *StructMeta) GetFields() (v []*FieldMeta) {
	return p.Fields
}

func (p *StructMeta) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StructMeta(%+v)", *p)
}
