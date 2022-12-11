// Copyright 2015-present Basho Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go.
// source: riak_yokozuna.proto
// DO NOT EDIT!

/*
Package riak_yokozuna is a generated protocol buffer package.

It is generated from these files:
	riak_yokozuna.proto

It has these top-level messages:
	RpbYokozunaIndex
	RpbYokozunaIndexGetReq
	RpbYokozunaIndexGetResp
	RpbYokozunaIndexPutReq
	RpbYokozunaIndexDeleteReq
	RpbYokozunaSchema
	RpbYokozunaSchemaPutReq
	RpbYokozunaSchemaGetReq
	RpbYokozunaSchemaGetResp
*/
package riak_yokozuna

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RpbYokozunaIndex struct {
	Name             []byte  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Schema           []byte  `protobuf:"bytes,2,opt,name=schema" json:"schema,omitempty"`
	NVal             *uint32 `protobuf:"varint,3,opt,name=n_val" json:"n_val,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbYokozunaIndex) Reset()                    { *m = RpbYokozunaIndex{} }
func (m *RpbYokozunaIndex) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndex) ProtoMessage()               {}
func (*RpbYokozunaIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RpbYokozunaIndex) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RpbYokozunaIndex) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *RpbYokozunaIndex) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

// GET request - If a name is given, return matching index, else return all
type RpbYokozunaIndexGetReq struct {
	Name             []byte `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaIndexGetReq) Reset()                    { *m = RpbYokozunaIndexGetReq{} }
func (m *RpbYokozunaIndexGetReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexGetReq) ProtoMessage()               {}
func (*RpbYokozunaIndexGetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RpbYokozunaIndexGetReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaIndexGetResp struct {
	Index            []*RpbYokozunaIndex `protobuf:"bytes,1,rep,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *RpbYokozunaIndexGetResp) Reset()                    { *m = RpbYokozunaIndexGetResp{} }
func (m *RpbYokozunaIndexGetResp) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexGetResp) ProtoMessage()               {}
func (*RpbYokozunaIndexGetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RpbYokozunaIndexGetResp) GetIndex() []*RpbYokozunaIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

// PUT request - Create a new index
type RpbYokozunaIndexPutReq struct {
	Index            *RpbYokozunaIndex `protobuf:"bytes,1,req,name=index" json:"index,omitempty"`
	Timeout          *uint32           `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *RpbYokozunaIndexPutReq) Reset()                    { *m = RpbYokozunaIndexPutReq{} }
func (m *RpbYokozunaIndexPutReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexPutReq) ProtoMessage()               {}
func (*RpbYokozunaIndexPutReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RpbYokozunaIndexPutReq) GetIndex() *RpbYokozunaIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *RpbYokozunaIndexPutReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

// DELETE request - Remove an index
type RpbYokozunaIndexDeleteReq struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaIndexDeleteReq) Reset()                    { *m = RpbYokozunaIndexDeleteReq{} }
func (m *RpbYokozunaIndexDeleteReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexDeleteReq) ProtoMessage()               {}
func (*RpbYokozunaIndexDeleteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RpbYokozunaIndexDeleteReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaSchema struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Content          []byte `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaSchema) Reset()                    { *m = RpbYokozunaSchema{} }
func (m *RpbYokozunaSchema) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaSchema) ProtoMessage()               {}
func (*RpbYokozunaSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RpbYokozunaSchema) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RpbYokozunaSchema) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// PUT request - create or potentially update a new schema
type RpbYokozunaSchemaPutReq struct {
	Schema           *RpbYokozunaSchema `protobuf:"bytes,1,req,name=schema" json:"schema,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RpbYokozunaSchemaPutReq) Reset()                    { *m = RpbYokozunaSchemaPutReq{} }
func (m *RpbYokozunaSchemaPutReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaPutReq) ProtoMessage()               {}
func (*RpbYokozunaSchemaPutReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RpbYokozunaSchemaPutReq) GetSchema() *RpbYokozunaSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

// GET request - Return matching schema by name
type RpbYokozunaSchemaGetReq struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaSchemaGetReq) Reset()                    { *m = RpbYokozunaSchemaGetReq{} }
func (m *RpbYokozunaSchemaGetReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaGetReq) ProtoMessage()               {}
func (*RpbYokozunaSchemaGetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RpbYokozunaSchemaGetReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaSchemaGetResp struct {
	Schema           *RpbYokozunaSchema `protobuf:"bytes,1,req,name=schema" json:"schema,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RpbYokozunaSchemaGetResp) Reset()                    { *m = RpbYokozunaSchemaGetResp{} }
func (m *RpbYokozunaSchemaGetResp) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaGetResp) ProtoMessage()               {}
func (*RpbYokozunaSchemaGetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RpbYokozunaSchemaGetResp) GetSchema() *RpbYokozunaSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func init() {
	proto.RegisterType((*RpbYokozunaIndex)(nil), "RpbYokozunaIndex")
	proto.RegisterType((*RpbYokozunaIndexGetReq)(nil), "RpbYokozunaIndexGetReq")
	proto.RegisterType((*RpbYokozunaIndexGetResp)(nil), "RpbYokozunaIndexGetResp")
	proto.RegisterType((*RpbYokozunaIndexPutReq)(nil), "RpbYokozunaIndexPutReq")
	proto.RegisterType((*RpbYokozunaIndexDeleteReq)(nil), "RpbYokozunaIndexDeleteReq")
	proto.RegisterType((*RpbYokozunaSchema)(nil), "RpbYokozunaSchema")
	proto.RegisterType((*RpbYokozunaSchemaPutReq)(nil), "RpbYokozunaSchemaPutReq")
	proto.RegisterType((*RpbYokozunaSchemaGetReq)(nil), "RpbYokozunaSchemaGetReq")
	proto.RegisterType((*RpbYokozunaSchemaGetResp)(nil), "RpbYokozunaSchemaGetResp")
}

func init() { proto.RegisterFile("riak_yokozuna.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x69, 0xe7, 0x1c, 0x3c, 0xb3, 0xcd, 0x45, 0x70, 0xf5, 0x16, 0x72, 0xd0, 0xec, 0x92,
	0xc3, 0xae, 0xa2, 0x42, 0x11, 0x44, 0xbc, 0x8c, 0x7a, 0xf2, 0x34, 0xd2, 0xfa, 0x64, 0xa5, 0x6b,
	0x52, 0xd7, 0x54, 0xd4, 0xbf, 0x5e, 0xda, 0x4e, 0xe9, 0xda, 0x1e, 0x3c, 0x26, 0xfc, 0xbe, 0xf7,
	0x7d, 0xef, 0xe3, 0xc1, 0xd9, 0x2e, 0x56, 0xc9, 0xfa, 0xcb, 0x24, 0xe6, 0xbb, 0xd0, 0x4a, 0x66,
	0x3b, 0x63, 0x0d, 0xbf, 0x83, 0xd3, 0x20, 0x0b, 0x5f, 0xf6, 0x9f, 0x8f, 0xfa, 0x15, 0x3f, 0x29,
	0x81, 0x23, 0xad, 0x52, 0xf4, 0x1c, 0xe6, 0x0a, 0x42, 0x27, 0x70, 0x9c, 0x47, 0x1b, 0x4c, 0x95,
	0xe7, 0x32, 0x47, 0x10, 0x3a, 0x86, 0xa1, 0x5e, 0x7f, 0xa8, 0xad, 0x37, 0x60, 0x8e, 0x18, 0xf3,
	0x4b, 0x38, 0x6f, 0x0f, 0x78, 0x40, 0x1b, 0xe0, 0x7b, 0x63, 0x8c, 0x23, 0x08, 0xbf, 0x86, 0x79,
	0x2f, 0x97, 0x67, 0x94, 0xc1, 0x30, 0x2e, 0xdf, 0x9e, 0xc3, 0x06, 0xe2, 0x64, 0x39, 0x93, 0x6d,
	0x90, 0x3f, 0x75, 0x4d, 0x56, 0x45, 0x65, 0xd2, 0xd0, 0xba, 0xbd, 0x5a, 0x3a, 0x85, 0x91, 0x8d,
	0x53, 0x34, 0x85, 0xad, 0x16, 0x18, 0xf3, 0x05, 0x5c, 0xb4, 0xa1, 0x7b, 0xdc, 0xa2, 0xc5, 0xc3,
	0xd0, 0xae, 0x20, 0x7c, 0x09, 0xb3, 0x06, 0xfa, 0x5c, 0xd5, 0xd0, 0xaa, 0x67, 0x0a, 0xa3, 0xc8,
	0x68, 0x8b, 0xba, 0x1e, 0x4f, 0xf8, 0xcd, 0xc1, 0xa2, 0xb5, 0x66, 0x1f, 0x96, 0xff, 0x55, 0x59,
	0xa7, 0xa5, 0xb2, 0x43, 0xf2, 0xab, 0x1e, 0x79, 0xa7, 0xd0, 0x32, 0xdb, 0x2d, 0x78, 0xfd, 0x60,
	0x9e, 0xfd, 0xc7, 0xc8, 0x5f, 0xc0, 0x3c, 0x32, 0xa9, 0x0c, 0x55, 0xbe, 0x31, 0xb2, 0x3c, 0x8d,
	0xfa, 0x22, 0xc2, 0xe2, 0xcd, 0x9f, 0x04, 0xb1, 0x4a, 0x7e, 0xf1, 0x95, 0xff, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xf5, 0x13, 0xa5, 0x19, 0x3a, 0x02, 0x00, 0x00,
}
