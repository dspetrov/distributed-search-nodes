// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package model

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Task struct {
	SearchTerms          []string `protobuf:"bytes,1,rep,name=searchTerms,proto3" json:"searchTerms,omitempty"`
	Documents            []string `protobuf:"bytes,2,rep,name=documents,proto3" json:"documents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetSearchTerms() []string {
	if m != nil {
		return m.SearchTerms
	}
	return nil
}

func (m *Task) GetDocuments() []string {
	if m != nil {
		return m.Documents
	}
	return nil
}

type DocumentData struct {
	TermToFrequency      map[string]float64 `protobuf:"bytes,1,rep,name=termToFrequency,proto3" json:"termToFrequency,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DocumentData) Reset()         { *m = DocumentData{} }
func (m *DocumentData) String() string { return proto.CompactTextString(m) }
func (*DocumentData) ProtoMessage()    {}
func (*DocumentData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}
func (m *DocumentData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentData.Unmarshal(m, b)
}
func (m *DocumentData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentData.Marshal(b, m, deterministic)
}
func (m *DocumentData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentData.Merge(m, src)
}
func (m *DocumentData) XXX_Size() int {
	return xxx_messageInfo_DocumentData.Size(m)
}
func (m *DocumentData) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentData.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentData proto.InternalMessageInfo

func (m *DocumentData) GetTermToFrequency() map[string]float64 {
	if m != nil {
		return m.TermToFrequency
	}
	return nil
}

type Result struct {
	DocumentToDocumentData map[string]*DocumentData `protobuf:"bytes,1,rep,name=documentToDocumentData,proto3" json:"documentToDocumentData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral   struct{}                 `json:"-"`
	XXX_unrecognized       []byte                   `json:"-"`
	XXX_sizecache          int32                    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetDocumentToDocumentData() map[string]*DocumentData {
	if m != nil {
		return m.DocumentToDocumentData
	}
	return nil
}

type Request struct {
	SearchQuery          string   `protobuf:"bytes,1,opt,name=search_query,json=searchQuery,proto3" json:"search_query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSearchQuery() string {
	if m != nil {
		return m.SearchQuery
	}
	return ""
}

type Response struct {
	RelevantDocuments    []*Response_DocumentStats `protobuf:"bytes,1,rep,name=relevant_documents,json=relevantDocuments,proto3" json:"relevant_documents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetRelevantDocuments() []*Response_DocumentStats {
	if m != nil {
		return m.RelevantDocuments
	}
	return nil
}

type Response_DocumentStats struct {
	DocumentName         string   `protobuf:"bytes,1,opt,name=document_name,json=documentName,proto3" json:"document_name,omitempty"`
	Score                float64  `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
	DocumentSize         int64    `protobuf:"varint,3,opt,name=document_size,json=documentSize,proto3" json:"document_size,omitempty"`
	Author               string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_DocumentStats) Reset()         { *m = Response_DocumentStats{} }
func (m *Response_DocumentStats) String() string { return proto.CompactTextString(m) }
func (*Response_DocumentStats) ProtoMessage()    {}
func (*Response_DocumentStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4, 0}
}
func (m *Response_DocumentStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_DocumentStats.Unmarshal(m, b)
}
func (m *Response_DocumentStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_DocumentStats.Marshal(b, m, deterministic)
}
func (m *Response_DocumentStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_DocumentStats.Merge(m, src)
}
func (m *Response_DocumentStats) XXX_Size() int {
	return xxx_messageInfo_Response_DocumentStats.Size(m)
}
func (m *Response_DocumentStats) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_DocumentStats.DiscardUnknown(m)
}

var xxx_messageInfo_Response_DocumentStats proto.InternalMessageInfo

func (m *Response_DocumentStats) GetDocumentName() string {
	if m != nil {
		return m.DocumentName
	}
	return ""
}

func (m *Response_DocumentStats) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Response_DocumentStats) GetDocumentSize() int64 {
	if m != nil {
		return m.DocumentSize
	}
	return 0
}

func (m *Response_DocumentStats) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "model.Task")
	proto.RegisterType((*DocumentData)(nil), "model.DocumentData")
	proto.RegisterMapType((map[string]float64)(nil), "model.DocumentData.TermToFrequencyEntry")
	proto.RegisterType((*Result)(nil), "model.Result")
	proto.RegisterMapType((map[string]*DocumentData)(nil), "model.Result.DocumentToDocumentDataEntry")
	proto.RegisterType((*Request)(nil), "model.Request")
	proto.RegisterType((*Response)(nil), "model.Response")
	proto.RegisterType((*Response_DocumentStats)(nil), "model.Response.DocumentStats")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xda, 0x40,
	0x18, 0xd4, 0x62, 0xa0, 0xe5, 0x33, 0xa8, 0xed, 0x16, 0x21, 0x8b, 0xb6, 0x92, 0xeb, 0x5e, 0x8c,
	0x54, 0x71, 0x20, 0x97, 0x28, 0xc7, 0x88, 0x70, 0x8a, 0x22, 0x65, 0xf1, 0x39, 0x68, 0x63, 0x3e,
	0x09, 0x84, 0xed, 0x85, 0xdd, 0x35, 0x12, 0xbc, 0x40, 0xde, 0x24, 0x2f, 0x93, 0xb7, 0xc8, 0x93,
	0x44, 0xfe, 0xc3, 0x76, 0x44, 0x72, 0xf3, 0x8c, 0xe7, 0x1b, 0xcf, 0xf7, 0x8d, 0xc1, 0x0c, 0xc5,
	0x12, 0x83, 0xf1, 0x56, 0x0a, 0x2d, 0x68, 0x2b, 0x05, 0xce, 0x0c, 0x9a, 0x1e, 0x57, 0x1b, 0x6a,
	0x83, 0xa9, 0x90, 0x4b, 0x7f, 0xe5, 0xa1, 0x0c, 0x95, 0x45, 0x6c, 0xc3, 0xed, 0xb0, 0x2a, 0x45,
	0x7f, 0x43, 0x67, 0x29, 0xfc, 0x38, 0xc4, 0x48, 0x2b, 0xab, 0x91, 0xbe, 0x2f, 0x09, 0xe7, 0x99,
	0x40, 0x77, 0x9a, 0xa3, 0x29, 0xd7, 0x9c, 0x32, 0xf8, 0xa6, 0x51, 0x86, 0x9e, 0x98, 0x49, 0xdc,
	0xc5, 0x18, 0xf9, 0x87, 0xd4, 0xd4, 0x9c, 0xb8, 0xe3, 0x2c, 0x46, 0x55, 0x3d, 0xf6, 0xea, 0xd2,
	0x9b, 0x48, 0xcb, 0x03, 0x7b, 0x6f, 0x30, 0xbc, 0x86, 0xfe, 0x39, 0x21, 0xfd, 0x0e, 0xc6, 0x06,
	0x13, 0x7f, 0xe2, 0x76, 0x58, 0xf2, 0x48, 0xfb, 0xd0, 0xda, 0xf3, 0x20, 0x46, 0xab, 0x61, 0x13,
	0x97, 0xb0, 0x0c, 0x5c, 0x35, 0x2e, 0x89, 0xf3, 0x42, 0xa0, 0xcd, 0x50, 0xc5, 0x81, 0xa6, 0x1c,
	0x06, 0xc5, 0x02, 0x9e, 0xa8, 0xc6, 0xc9, 0x93, 0x8e, 0xf2, 0xa4, 0x99, 0xfc, 0x14, 0xb8, 0xae,
	0xcd, 0xa2, 0x7e, 0x60, 0x34, 0x7c, 0x80, 0x5f, 0x9f, 0x8c, 0x9d, 0x09, 0x3e, 0xaa, 0x06, 0x37,
	0x27, 0x3f, 0xcf, 0x1c, 0xab, 0xba, 0xcd, 0x7f, 0xf8, 0xc2, 0x92, 0x53, 0x28, 0x4d, 0xff, 0x42,
	0x37, 0xab, 0x6b, 0xb1, 0x8b, 0x51, 0x16, 0xa6, 0x79, 0x85, 0xf7, 0x09, 0xe5, 0xbc, 0x12, 0xf8,
	0xca, 0x50, 0x6d, 0x45, 0xa4, 0x90, 0xde, 0x02, 0x95, 0x18, 0xe0, 0x9e, 0x47, 0x7a, 0x51, 0x16,
	0x9b, 0x6d, 0xfe, 0xa7, 0xdc, 0x3c, 0x15, 0x9f, 0xbe, 0x3f, 0xd7, 0x5c, 0x2b, 0xf6, 0xa3, 0x18,
	0x2c, 0x68, 0x35, 0x7c, 0x22, 0xd0, 0xab, 0x89, 0xe8, 0x3f, 0xe8, 0x15, 0xb6, 0x8b, 0x88, 0x87,
	0x98, 0x07, 0xea, 0x16, 0xe4, 0x1d, 0x0f, 0x31, 0xe9, 0x49, 0xf9, 0x42, 0x9e, 0x7a, 0x4a, 0x41,
	0x6d, 0x54, 0xad, 0x8f, 0x68, 0x19, 0x36, 0x71, 0x8d, 0x72, 0x74, 0xbe, 0x3e, 0x22, 0x1d, 0x40,
	0x9b, 0xc7, 0x7a, 0x25, 0xa4, 0xd5, 0x4c, 0x8d, 0x73, 0xf4, 0xd8, 0x4e, 0xff, 0xef, 0x8b, 0xb7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0xcf, 0x6d, 0xda, 0xee, 0x02, 0x00, 0x00,
}
