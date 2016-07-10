// Code generated by protoc-gen-go.
// source: pb/library.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/library.proto

It has these top-level messages:
	Document
	CreateDocumentRequest
	CreateDocumentReply
	FindDocumentsRequest
	FindDocumentsReply
	FindDocumentsByIdRequest
	FindDocumentsByIdReply
	ReplaceDocumentByIdRequest
	ReplaceDocumentByIdReply
	DeleteDocumentsByIdRequest
	DeleteDocumentsByIdReply
*/
package pb

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

type Document struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *Document) Reset()                    { *m = Document{} }
func (m *Document) String() string            { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()               {}
func (*Document) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateDocumentRequest struct {
	UserId   string    `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Document *Document `protobuf:"bytes,2,opt,name=document" json:"document,omitempty"`
}

func (m *CreateDocumentRequest) Reset()                    { *m = CreateDocumentRequest{} }
func (m *CreateDocumentRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDocumentRequest) ProtoMessage()               {}
func (*CreateDocumentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateDocumentRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type CreateDocumentReply struct {
	Document *Document `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	Err      string    `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *CreateDocumentReply) Reset()                    { *m = CreateDocumentReply{} }
func (m *CreateDocumentReply) String() string            { return proto.CompactTextString(m) }
func (*CreateDocumentReply) ProtoMessage()               {}
func (*CreateDocumentReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateDocumentReply) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type FindDocumentsRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *FindDocumentsRequest) Reset()                    { *m = FindDocumentsRequest{} }
func (m *FindDocumentsRequest) String() string            { return proto.CompactTextString(m) }
func (*FindDocumentsRequest) ProtoMessage()               {}
func (*FindDocumentsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type FindDocumentsReply struct {
	Documents []*Document `protobuf:"bytes,1,rep,name=documents" json:"documents,omitempty"`
	Err       string      `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *FindDocumentsReply) Reset()                    { *m = FindDocumentsReply{} }
func (m *FindDocumentsReply) String() string            { return proto.CompactTextString(m) }
func (*FindDocumentsReply) ProtoMessage()               {}
func (*FindDocumentsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FindDocumentsReply) GetDocuments() []*Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

type FindDocumentsByIdRequest struct {
	UserId string   `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Ids    []string `protobuf:"bytes,2,rep,name=ids" json:"ids,omitempty"`
}

func (m *FindDocumentsByIdRequest) Reset()                    { *m = FindDocumentsByIdRequest{} }
func (m *FindDocumentsByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FindDocumentsByIdRequest) ProtoMessage()               {}
func (*FindDocumentsByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type FindDocumentsByIdReply struct {
	Documents []*Document `protobuf:"bytes,1,rep,name=documents" json:"documents,omitempty"`
	Err       string      `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *FindDocumentsByIdReply) Reset()                    { *m = FindDocumentsByIdReply{} }
func (m *FindDocumentsByIdReply) String() string            { return proto.CompactTextString(m) }
func (*FindDocumentsByIdReply) ProtoMessage()               {}
func (*FindDocumentsByIdReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FindDocumentsByIdReply) GetDocuments() []*Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

type ReplaceDocumentByIdRequest struct {
	UserId   string    `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Id       string    `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Document *Document `protobuf:"bytes,3,opt,name=document" json:"document,omitempty"`
}

func (m *ReplaceDocumentByIdRequest) Reset()                    { *m = ReplaceDocumentByIdRequest{} }
func (m *ReplaceDocumentByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*ReplaceDocumentByIdRequest) ProtoMessage()               {}
func (*ReplaceDocumentByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReplaceDocumentByIdRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type ReplaceDocumentByIdReply struct {
	Document *Document `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	Err      string    `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *ReplaceDocumentByIdReply) Reset()                    { *m = ReplaceDocumentByIdReply{} }
func (m *ReplaceDocumentByIdReply) String() string            { return proto.CompactTextString(m) }
func (*ReplaceDocumentByIdReply) ProtoMessage()               {}
func (*ReplaceDocumentByIdReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReplaceDocumentByIdReply) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type DeleteDocumentsByIdRequest struct {
	UserId string   `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Ids    []string `protobuf:"bytes,2,rep,name=ids" json:"ids,omitempty"`
}

func (m *DeleteDocumentsByIdRequest) Reset()                    { *m = DeleteDocumentsByIdRequest{} }
func (m *DeleteDocumentsByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDocumentsByIdRequest) ProtoMessage()               {}
func (*DeleteDocumentsByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type DeleteDocumentsByIdReply struct {
	Documents []*Document `protobuf:"bytes,1,rep,name=documents" json:"documents,omitempty"`
	Err       string      `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *DeleteDocumentsByIdReply) Reset()                    { *m = DeleteDocumentsByIdReply{} }
func (m *DeleteDocumentsByIdReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteDocumentsByIdReply) ProtoMessage()               {}
func (*DeleteDocumentsByIdReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteDocumentsByIdReply) GetDocuments() []*Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

func init() {
	proto.RegisterType((*Document)(nil), "pb.Document")
	proto.RegisterType((*CreateDocumentRequest)(nil), "pb.CreateDocumentRequest")
	proto.RegisterType((*CreateDocumentReply)(nil), "pb.CreateDocumentReply")
	proto.RegisterType((*FindDocumentsRequest)(nil), "pb.FindDocumentsRequest")
	proto.RegisterType((*FindDocumentsReply)(nil), "pb.FindDocumentsReply")
	proto.RegisterType((*FindDocumentsByIdRequest)(nil), "pb.FindDocumentsByIdRequest")
	proto.RegisterType((*FindDocumentsByIdReply)(nil), "pb.FindDocumentsByIdReply")
	proto.RegisterType((*ReplaceDocumentByIdRequest)(nil), "pb.ReplaceDocumentByIdRequest")
	proto.RegisterType((*ReplaceDocumentByIdReply)(nil), "pb.ReplaceDocumentByIdReply")
	proto.RegisterType((*DeleteDocumentsByIdRequest)(nil), "pb.DeleteDocumentsByIdRequest")
	proto.RegisterType((*DeleteDocumentsByIdReply)(nil), "pb.DeleteDocumentsByIdReply")
}

func init() { proto.RegisterFile("pb/library.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x94, 0x5f, 0x4f, 0xf2, 0x30,
	0x14, 0xc6, 0xb3, 0x35, 0xe1, 0xcf, 0x79, 0x5f, 0x09, 0x16, 0x85, 0xda, 0x10, 0x63, 0x7a, 0x45,
	0xbc, 0x80, 0x04, 0x3f, 0x81, 0x8a, 0x1a, 0x12, 0xb9, 0x10, 0x13, 0x62, 0xbc, 0x31, 0x8c, 0xf6,
	0x62, 0x09, 0xb2, 0xd9, 0x8d, 0x0b, 0x3e, 0x83, 0x5f, 0xda, 0x76, 0x63, 0x83, 0x41, 0x07, 0xc6,
	0xdd, 0xb5, 0x3d, 0xe7, 0xfc, 0x7a, 0x76, 0x9e, 0xa7, 0x83, 0xba, 0xef, 0xf4, 0xe6, 0xae, 0x23,
	0xa7, 0x72, 0xd5, 0xf5, 0xa5, 0x17, 0x7a, 0xd8, 0xf6, 0x1d, 0x36, 0x82, 0xca, 0xc0, 0x9b, 0x2d,
	0x3f, 0xc5, 0x22, 0xc4, 0x35, 0xb0, 0x5d, 0x4e, 0xac, 0x2b, 0xab, 0x53, 0x1d, 0xab, 0x15, 0x6e,
	0x41, 0x79, 0x19, 0x08, 0xf9, 0xa1, 0x0e, 0xed, 0xe8, 0xb0, 0xa4, 0xb7, 0x43, 0x8e, 0x09, 0x94,
	0x67, 0xde, 0x22, 0x54, 0x35, 0x04, 0x45, 0x81, 0x64, 0xcb, 0xde, 0xe1, 0xfc, 0x5e, 0x8a, 0x69,
	0x28, 0x12, 0xe8, 0x58, 0x7c, 0x2d, 0x45, 0x10, 0x6e, 0xb3, 0xac, 0x0c, 0xab, 0x03, 0x15, 0xbe,
	0xce, 0x8d, 0x6e, 0xf9, 0xd7, 0xff, 0xdf, 0xf5, 0x9d, 0x6e, 0x5a, 0x9f, 0x46, 0xd9, 0x0b, 0x34,
	0x76, 0xd9, 0xfe, 0x7c, 0x95, 0x01, 0x58, 0x87, 0x00, 0xb8, 0x0e, 0x48, 0x48, 0xb9, 0xfe, 0x16,
	0xbd, 0x64, 0x3d, 0x38, 0x7b, 0x74, 0x17, 0x3c, 0xc9, 0x0d, 0x8e, 0x75, 0xcb, 0xc6, 0x80, 0x77,
	0x0a, 0x74, 0x0b, 0xd7, 0x50, 0x4d, 0x2e, 0x09, 0x54, 0x01, 0xda, 0xeb, 0x61, 0x13, 0x36, 0x34,
	0xf1, 0x00, 0x24, 0xc3, 0xbc, 0x5b, 0x0d, 0xf9, 0xd1, 0xb1, 0x29, 0x8c, 0xcb, 0x03, 0x85, 0x41,
	0x1a, 0xa3, 0x96, 0x6c, 0x02, 0x4d, 0x03, 0xa6, 0x78, 0x7b, 0x1e, 0x50, 0x8d, 0x99, 0xce, 0xd2,
	0xb9, 0xff, 0xaa, 0xc1, 0xd8, 0x4c, 0x76, 0x6a, 0xa6, 0x6d, 0x99, 0xd0, 0x41, 0x9d, 0x27, 0x40,
	0x8c, 0x17, 0x16, 0x15, 0xfb, 0x09, 0xe8, 0x40, 0xcc, 0xc5, 0xc6, 0x3f, 0x7f, 0x9d, 0xf4, 0x1b,
	0x10, 0x23, 0xa8, 0xf0, 0xac, 0xfb, 0xdf, 0x08, 0xca, 0xcf, 0xf1, 0x1b, 0xc5, 0x03, 0xa8, 0x65,
	0xed, 0x8e, 0x2f, 0x34, 0xc8, 0xf8, 0xbc, 0x68, 0xcb, 0x14, 0xd2, 0xfd, 0xdc, 0xc2, 0x49, 0xc6,
	0x15, 0x98, 0xe8, 0x4c, 0x93, 0xe9, 0x69, 0xd3, 0x10, 0xd1, 0x88, 0x11, 0x9c, 0xee, 0x19, 0x0b,
	0xb7, 0xf7, 0x92, 0xb7, 0x86, 0x49, 0x69, 0x4e, 0x54, 0xe3, 0x5e, 0xa1, 0x61, 0x90, 0x17, 0x5f,
	0xea, 0x92, 0x7c, 0xa3, 0xd1, 0x76, 0x6e, 0x7c, 0x0d, 0x35, 0x48, 0x12, 0x43, 0xf3, 0x45, 0x8f,
	0xa1, 0x79, 0x5a, 0x3a, 0xa5, 0xe8, 0x37, 0x79, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x27,
	0xa9, 0xda, 0x3a, 0x05, 0x00, 0x00,
}