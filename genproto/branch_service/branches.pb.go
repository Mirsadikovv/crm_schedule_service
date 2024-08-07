// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: branches.proto

package branch_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type BranchPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BranchPrimaryKey) Reset() {
	*x = BranchPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchPrimaryKey) ProtoMessage() {}

func (x *BranchPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchPrimaryKey.ProtoReflect.Descriptor instead.
func (*BranchPrimaryKey) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{0}
}

func (x *BranchPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchName     string `protobuf:"bytes,1,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	BranchLocation string `protobuf:"bytes,2,opt,name=branch_location,json=branchLocation,proto3" json:"branch_location,omitempty"`
	Phone          string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	OpenTime       string `protobuf:"bytes,4,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	CloseTime      string `protobuf:"bytes,5,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
}

func (x *CreateBranch) Reset() {
	*x = CreateBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBranch) ProtoMessage() {}

func (x *CreateBranch) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBranch.ProtoReflect.Descriptor instead.
func (*CreateBranch) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBranch) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *CreateBranch) GetBranchLocation() string {
	if x != nil {
		return x.BranchLocation
	}
	return ""
}

func (x *CreateBranch) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateBranch) GetOpenTime() string {
	if x != nil {
		return x.OpenTime
	}
	return ""
}

func (x *CreateBranch) GetCloseTime() string {
	if x != nil {
		return x.CloseTime
	}
	return ""
}

type GetBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchName     string `protobuf:"bytes,1,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	BranchLocation string `protobuf:"bytes,2,opt,name=branch_location,json=branchLocation,proto3" json:"branch_location,omitempty"`
	Phone          string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	OpenTime       string `protobuf:"bytes,4,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	CloseTime      string `protobuf:"bytes,5,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	CreatedAt      string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Id             string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBranch) Reset() {
	*x = GetBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranch) ProtoMessage() {}

func (x *GetBranch) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranch.ProtoReflect.Descriptor instead.
func (*GetBranch) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{2}
}

func (x *GetBranch) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *GetBranch) GetBranchLocation() string {
	if x != nil {
		return x.BranchLocation
	}
	return ""
}

func (x *GetBranch) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetBranch) GetOpenTime() string {
	if x != nil {
		return x.OpenTime
	}
	return ""
}

func (x *GetBranch) GetCloseTime() string {
	if x != nil {
		return x.CloseTime
	}
	return ""
}

func (x *GetBranch) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetBranch) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GetBranch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchName     string `protobuf:"bytes,1,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	BranchLocation string `protobuf:"bytes,2,opt,name=branch_location,json=branchLocation,proto3" json:"branch_location,omitempty"`
	Phone          string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	OpenTime       string `protobuf:"bytes,4,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	CloseTime      string `protobuf:"bytes,5,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	Id             string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateBranch) Reset() {
	*x = UpdateBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBranch) ProtoMessage() {}

func (x *UpdateBranch) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBranch.ProtoReflect.Descriptor instead.
func (*UpdateBranch) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBranch) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *UpdateBranch) GetBranchLocation() string {
	if x != nil {
		return x.BranchLocation
	}
	return ""
}

func (x *UpdateBranch) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateBranch) GetOpenTime() string {
	if x != nil {
		return x.OpenTime
	}
	return ""
}

func (x *UpdateBranch) GetCloseTime() string {
	if x != nil {
		return x.CloseTime
	}
	return ""
}

func (x *UpdateBranch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListBranchRequest) Reset() {
	*x = GetListBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBranchRequest) ProtoMessage() {}

func (x *GetListBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBranchRequest.ProtoReflect.Descriptor instead.
func (*GetListBranchRequest) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{4}
}

func (x *GetListBranchRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListBranchRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListBranchRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListBranchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   int64        `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Branchs []*GetBranch `protobuf:"bytes,2,rep,name=Branchs,proto3" json:"Branchs,omitempty"`
}

func (x *GetListBranchResponse) Reset() {
	*x = GetListBranchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBranchResponse) ProtoMessage() {}

func (x *GetListBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBranchResponse.ProtoReflect.Descriptor instead.
func (*GetListBranchResponse) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{5}
}

func (x *GetListBranchResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListBranchResponse) GetBranchs() []*GetBranch {
	if x != nil {
		return x.Branchs
	}
	return nil
}

var File_branches_proto protoreflect.FileDescriptor

var file_branches_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x67, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x22, 0x0a, 0x10, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0xf5, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x22, 0x65, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x52, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x73, 0x32, 0x9e, 0x03, 0x0a, 0x0d,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x1a, 0x1c, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1c, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1f, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x1a, 0x1c, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67,
	0x6f, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branches_proto_rawDescOnce sync.Once
	file_branches_proto_rawDescData = file_branches_proto_rawDesc
)

func file_branches_proto_rawDescGZIP() []byte {
	file_branches_proto_rawDescOnce.Do(func() {
		file_branches_proto_rawDescData = protoimpl.X.CompressGZIP(file_branches_proto_rawDescData)
	})
	return file_branches_proto_rawDescData
}

var file_branches_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_branches_proto_goTypes = []interface{}{
	(*BranchPrimaryKey)(nil),      // 0: branch_service_go.BranchPrimaryKey
	(*CreateBranch)(nil),          // 1: branch_service_go.CreateBranch
	(*GetBranch)(nil),             // 2: branch_service_go.GetBranch
	(*UpdateBranch)(nil),          // 3: branch_service_go.UpdateBranch
	(*GetListBranchRequest)(nil),  // 4: branch_service_go.GetListBranchRequest
	(*GetListBranchResponse)(nil), // 5: branch_service_go.GetListBranchResponse
	(*empty.Empty)(nil),           // 6: google.protobuf.Empty
}
var file_branches_proto_depIdxs = []int32{
	2, // 0: branch_service_go.GetListBranchResponse.Branchs:type_name -> branch_service_go.GetBranch
	1, // 1: branch_service_go.BranchService.Create:input_type -> branch_service_go.CreateBranch
	0, // 2: branch_service_go.BranchService.GetByID:input_type -> branch_service_go.BranchPrimaryKey
	4, // 3: branch_service_go.BranchService.GetList:input_type -> branch_service_go.GetListBranchRequest
	3, // 4: branch_service_go.BranchService.Update:input_type -> branch_service_go.UpdateBranch
	0, // 5: branch_service_go.BranchService.Delete:input_type -> branch_service_go.BranchPrimaryKey
	2, // 6: branch_service_go.BranchService.Create:output_type -> branch_service_go.GetBranch
	2, // 7: branch_service_go.BranchService.GetByID:output_type -> branch_service_go.GetBranch
	5, // 8: branch_service_go.BranchService.GetList:output_type -> branch_service_go.GetListBranchResponse
	2, // 9: branch_service_go.BranchService.Update:output_type -> branch_service_go.GetBranch
	6, // 10: branch_service_go.BranchService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_branches_proto_init() }
func file_branches_proto_init() {
	if File_branches_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_branches_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchPrimaryKey); i {
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
		file_branches_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBranch); i {
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
		file_branches_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBranch); i {
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
		file_branches_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBranch); i {
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
		file_branches_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListBranchRequest); i {
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
		file_branches_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListBranchResponse); i {
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
			RawDescriptor: file_branches_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_branches_proto_goTypes,
		DependencyIndexes: file_branches_proto_depIdxs,
		MessageInfos:      file_branches_proto_msgTypes,
	}.Build()
	File_branches_proto = out.File
	file_branches_proto_rawDesc = nil
	file_branches_proto_goTypes = nil
	file_branches_proto_depIdxs = nil
}
