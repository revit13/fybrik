// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.1
// source: data_catalog_request.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CatalogDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CredentialPath string `protobuf:"bytes,1,opt,name=credential_path,json=credentialPath,proto3" json:"credential_path,omitempty"` // link to vault plugin for reading k8s secret with user credentials
	DatasetId      string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`                // identifier of asset - always needed. JSON expected. Interpreted by the Connector, can contain any additional information as part of JSON
}

func (x *CatalogDatasetRequest) Reset() {
	*x = CatalogDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_catalog_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatalogDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatalogDatasetRequest) ProtoMessage() {}

func (x *CatalogDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_catalog_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatalogDatasetRequest.ProtoReflect.Descriptor instead.
func (*CatalogDatasetRequest) Descriptor() ([]byte, []int) {
	return file_data_catalog_request_proto_rawDescGZIP(), []int{0}
}

func (x *CatalogDatasetRequest) GetCredentialPath() string {
	if x != nil {
		return x.CredentialPath
	}
	return ""
}

func (x *CatalogDatasetRequest) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

var File_data_catalog_request_proto protoreflect.FileDescriptor

var file_data_catalog_request_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x5f, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x42, 0x35, 0x0a, 0x09, 0x69, 0x6f, 0x2e,
	0x66, 0x79, 0x62, 0x72, 0x69, 0x6b, 0x5a, 0x28, 0x66, 0x79, 0x62, 0x72, 0x69, 0x6b, 0x2e, 0x69,
	0x6f, 0x2f, 0x66, 0x79, 0x62, 0x72, 0x69, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_catalog_request_proto_rawDescOnce sync.Once
	file_data_catalog_request_proto_rawDescData = file_data_catalog_request_proto_rawDesc
)

func file_data_catalog_request_proto_rawDescGZIP() []byte {
	file_data_catalog_request_proto_rawDescOnce.Do(func() {
		file_data_catalog_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_catalog_request_proto_rawDescData)
	})
	return file_data_catalog_request_proto_rawDescData
}

var file_data_catalog_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_data_catalog_request_proto_goTypes = []interface{}{
	(*CatalogDatasetRequest)(nil), // 0: connectors.CatalogDatasetRequest
}
var file_data_catalog_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_catalog_request_proto_init() }
func file_data_catalog_request_proto_init() {
	if File_data_catalog_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_catalog_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatalogDatasetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_catalog_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_catalog_request_proto_goTypes,
		DependencyIndexes: file_data_catalog_request_proto_depIdxs,
		MessageInfos:      file_data_catalog_request_proto_msgTypes,
	}.Build()
	File_data_catalog_request_proto = out.File
	file_data_catalog_request_proto_rawDesc = nil
	file_data_catalog_request_proto_goTypes = nil
	file_data_catalog_request_proto_depIdxs = nil
}
