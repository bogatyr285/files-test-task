// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/files-api.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetFilesRequest) Reset() {
	*x = GetFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilesRequest) ProtoMessage() {}

func (x *GetFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilesRequest.ProtoReflect.Descriptor instead.
func (*GetFilesRequest) Descriptor() ([]byte, []int) {
	return file_proto_files_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetFilesRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	OwnerId   string `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerType string `protobuf:"bytes,4,opt,name=ownerType,proto3" json:"ownerType,omitempty"`
	Url       string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	ObjectKey string `protobuf:"bytes,6,opt,name=object_key,json=objectKey,proto3" json:"object_key,omitempty"`
	Size      int64  `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_proto_files_api_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *File) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *File) GetOwnerType() string {
	if x != nil {
		return x.OwnerType
	}
	return ""
}

func (x *File) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *File) GetObjectKey() string {
	if x != nil {
		return x.ObjectKey
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FilesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*File `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *FilesList) Reset() {
	*x = FilesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesList) ProtoMessage() {}

func (x *FilesList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesList.ProtoReflect.Descriptor instead.
func (*FilesList) Descriptor() ([]byte, []int) {
	return file_proto_files_api_proto_rawDescGZIP(), []int{2}
}

func (x *FilesList) GetRecords() []*File {
	if x != nil {
		return x.Records
	}
	return nil
}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UploadFileRequest_Info
	//	*UploadFileRequest_ChunkData
	Data isUploadFileRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_files_api_proto_rawDescGZIP(), []int{3}
}

func (m *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadFileRequest) GetInfo() *FileInfo {
	if x, ok := x.GetData().(*UploadFileRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadFileRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UploadFileRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_Info struct {
	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadFileRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadFileRequest_Info) isUploadFileRequest_Data() {}

func (*UploadFileRequest_ChunkData) isUploadFileRequest_Data() {}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId       string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerType     string `protobuf:"bytes,2,opt,name=owner_type,json=ownerType,proto3" json:"owner_type,omitempty"`
	FileExtension string `protobuf:"bytes,3,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty"`
	Filename      string `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_proto_files_api_proto_rawDescGZIP(), []int{4}
}

func (x *FileInfo) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *FileInfo) GetOwnerType() string {
	if x != nil {
		return x.OwnerType
	}
	return ""
}

func (x *FileInfo) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

func (x *FileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

// update statuses
type UpdateFilesStatusesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApproveIds []string `protobuf:"bytes,1,rep,name=approve_ids,json=approveIds,proto3" json:"approve_ids,omitempty"`
	DeleteIds  []string `protobuf:"bytes,2,rep,name=delete_ids,json=deleteIds,proto3" json:"delete_ids,omitempty"`
}

func (x *UpdateFilesStatusesRequest) Reset() {
	*x = UpdateFilesStatusesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFilesStatusesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFilesStatusesRequest) ProtoMessage() {}

func (x *UpdateFilesStatusesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFilesStatusesRequest.ProtoReflect.Descriptor instead.
func (*UpdateFilesStatusesRequest) Descriptor() ([]byte, []int) {
	return file_proto_files_api_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFilesStatusesRequest) GetApproveIds() []string {
	if x != nil {
		return x.ApproveIds
	}
	return nil
}

func (x *UpdateFilesStatusesRequest) GetDeleteIds() []string {
	if x != nil {
		return x.DeleteIds
	}
	return nil
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApprovedCount int64 `protobuf:"varint,1,opt,name=approved_count,json=approvedCount,proto3" json:"approved_count,omitempty"`
	DeletedCount  int64 `protobuf:"varint,2,opt,name=deleted_count,json=deletedCount,proto3" json:"deleted_count,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_files_api_proto_rawDescGZIP(), []int{6}
}

func (x *StatusResponse) GetApprovedCount() int64 {
	if x != nil {
		return x.ApprovedCount
	}
	return 0
}

func (x *StatusResponse) GetDeletedCount() int64 {
	if x != nil {
		return x.DeletedCount
	}
	return 0
}

var File_proto_files_api_proto protoreflect.FileDescriptor

var file_proto_files_api_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2d, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x23,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x42, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x73, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x87, 0x01, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x64, 0x73, 0x22, 0x5c, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0xae, 0x02, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x41, 0x50, 0x49, 0x12, 0x56,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x71, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_files_api_proto_rawDescOnce sync.Once
	file_proto_files_api_proto_rawDescData = file_proto_files_api_proto_rawDesc
)

func file_proto_files_api_proto_rawDescGZIP() []byte {
	file_proto_files_api_proto_rawDescOnce.Do(func() {
		file_proto_files_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_api_proto_rawDescData)
	})
	return file_proto_files_api_proto_rawDescData
}

var file_proto_files_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_files_api_proto_goTypes = []interface{}{
	(*GetFilesRequest)(nil),            // 0: files_api.grpc.schema.GetFilesRequest
	(*File)(nil),                       // 1: files_api.grpc.schema.File
	(*FilesList)(nil),                  // 2: files_api.grpc.schema.FilesList
	(*UploadFileRequest)(nil),          // 3: files_api.grpc.schema.UploadFileRequest
	(*FileInfo)(nil),                   // 4: files_api.grpc.schema.FileInfo
	(*UpdateFilesStatusesRequest)(nil), // 5: files_api.grpc.schema.UpdateFilesStatusesRequest
	(*StatusResponse)(nil),             // 6: files_api.grpc.schema.StatusResponse
}
var file_proto_files_api_proto_depIdxs = []int32{
	1, // 0: files_api.grpc.schema.FilesList.records:type_name -> files_api.grpc.schema.File
	4, // 1: files_api.grpc.schema.UploadFileRequest.info:type_name -> files_api.grpc.schema.FileInfo
	0, // 2: files_api.grpc.schema.FilesAPI.GetFiles:input_type -> files_api.grpc.schema.GetFilesRequest
	3, // 3: files_api.grpc.schema.FilesAPI.UploadFile:input_type -> files_api.grpc.schema.UploadFileRequest
	5, // 4: files_api.grpc.schema.FilesAPI.UpdateFilesStatuses:input_type -> files_api.grpc.schema.UpdateFilesStatusesRequest
	2, // 5: files_api.grpc.schema.FilesAPI.GetFiles:output_type -> files_api.grpc.schema.FilesList
	1, // 6: files_api.grpc.schema.FilesAPI.UploadFile:output_type -> files_api.grpc.schema.File
	6, // 7: files_api.grpc.schema.FilesAPI.UpdateFilesStatuses:output_type -> files_api.grpc.schema.StatusResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_files_api_proto_init() }
func file_proto_files_api_proto_init() {
	if File_proto_files_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_files_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilesRequest); i {
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
		file_proto_files_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_proto_files_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesList); i {
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
		file_proto_files_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_proto_files_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_proto_files_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFilesStatusesRequest); i {
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
		file_proto_files_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
	file_proto_files_api_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*UploadFileRequest_Info)(nil),
		(*UploadFileRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_files_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_files_api_proto_goTypes,
		DependencyIndexes: file_proto_files_api_proto_depIdxs,
		MessageInfos:      file_proto_files_api_proto_msgTypes,
	}.Build()
	File_proto_files_api_proto = out.File
	file_proto_files_api_proto_rawDesc = nil
	file_proto_files_api_proto_goTypes = nil
	file_proto_files_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FilesAPIClient is the client API for FilesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FilesAPIClient interface {
	GetFiles(ctx context.Context, in *GetFilesRequest, opts ...grpc.CallOption) (*FilesList, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FilesAPI_UploadFileClient, error)
	UpdateFilesStatuses(ctx context.Context, in *UpdateFilesStatusesRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type filesAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesAPIClient(cc grpc.ClientConnInterface) FilesAPIClient {
	return &filesAPIClient{cc}
}

func (c *filesAPIClient) GetFiles(ctx context.Context, in *GetFilesRequest, opts ...grpc.CallOption) (*FilesList, error) {
	out := new(FilesList)
	err := c.cc.Invoke(ctx, "/files_api.grpc.schema.FilesAPI/GetFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesAPIClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FilesAPI_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FilesAPI_serviceDesc.Streams[0], "/files_api.grpc.schema.FilesAPI/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &filesAPIUploadFileClient{stream}
	return x, nil
}

type FilesAPI_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*File, error)
	grpc.ClientStream
}

type filesAPIUploadFileClient struct {
	grpc.ClientStream
}

func (x *filesAPIUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *filesAPIUploadFileClient) CloseAndRecv() (*File, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(File)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *filesAPIClient) UpdateFilesStatuses(ctx context.Context, in *UpdateFilesStatusesRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/files_api.grpc.schema.FilesAPI/UpdateFilesStatuses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesAPIServer is the server API for FilesAPI service.
type FilesAPIServer interface {
	GetFiles(context.Context, *GetFilesRequest) (*FilesList, error)
	UploadFile(FilesAPI_UploadFileServer) error
	UpdateFilesStatuses(context.Context, *UpdateFilesStatusesRequest) (*StatusResponse, error)
}

// UnimplementedFilesAPIServer can be embedded to have forward compatible implementations.
type UnimplementedFilesAPIServer struct {
}

func (*UnimplementedFilesAPIServer) GetFiles(context.Context, *GetFilesRequest) (*FilesList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (*UnimplementedFilesAPIServer) UploadFile(FilesAPI_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedFilesAPIServer) UpdateFilesStatuses(context.Context, *UpdateFilesStatusesRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFilesStatuses not implemented")
}

func RegisterFilesAPIServer(s *grpc.Server, srv FilesAPIServer) {
	s.RegisterService(&_FilesAPI_serviceDesc, srv)
}

func _FilesAPI_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesAPIServer).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/files_api.grpc.schema.FilesAPI/GetFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesAPIServer).GetFiles(ctx, req.(*GetFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesAPI_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FilesAPIServer).UploadFile(&filesAPIUploadFileServer{stream})
}

type FilesAPI_UploadFileServer interface {
	SendAndClose(*File) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type filesAPIUploadFileServer struct {
	grpc.ServerStream
}

func (x *filesAPIUploadFileServer) SendAndClose(m *File) error {
	return x.ServerStream.SendMsg(m)
}

func (x *filesAPIUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FilesAPI_UpdateFilesStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFilesStatusesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesAPIServer).UpdateFilesStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/files_api.grpc.schema.FilesAPI/UpdateFilesStatuses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesAPIServer).UpdateFilesStatuses(ctx, req.(*UpdateFilesStatusesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FilesAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "files_api.grpc.schema.FilesAPI",
	HandlerType: (*FilesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFiles",
			Handler:    _FilesAPI_GetFiles_Handler,
		},
		{
			MethodName: "UpdateFilesStatuses",
			Handler:    _FilesAPI_UpdateFilesStatuses_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FilesAPI_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/files-api.proto",
}
