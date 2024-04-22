// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: client_event.proto

package pbce

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ClientEvent source resource reference: pkg/dal/table/client-event.go
type ClientEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec                *ClientEventSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment          *ClientEventAttachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	HeartbeatTime       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=heartbeat_time,json=heartbeatTime,proto3" json:"heartbeat_time,omitempty"`
	MessageType         string                 `protobuf:"bytes,5,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	OriginalReleaseName string                 `protobuf:"bytes,6,opt,name=original_release_name,json=originalReleaseName,proto3" json:"original_release_name,omitempty"`
	TargetReleaseName   string                 `protobuf:"bytes,7,opt,name=target_release_name,json=targetReleaseName,proto3" json:"target_release_name,omitempty"`
}

func (x *ClientEvent) Reset() {
	*x = ClientEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEvent) ProtoMessage() {}

func (x *ClientEvent) ProtoReflect() protoreflect.Message {
	mi := &file_client_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEvent.ProtoReflect.Descriptor instead.
func (*ClientEvent) Descriptor() ([]byte, []int) {
	return file_client_event_proto_rawDescGZIP(), []int{0}
}

func (x *ClientEvent) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ClientEvent) GetSpec() *ClientEventSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ClientEvent) GetAttachment() *ClientEventAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *ClientEvent) GetHeartbeatTime() *timestamppb.Timestamp {
	if x != nil {
		return x.HeartbeatTime
	}
	return nil
}

func (x *ClientEvent) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *ClientEvent) GetOriginalReleaseName() string {
	if x != nil {
		return x.OriginalReleaseName
	}
	return ""
}

func (x *ClientEvent) GetTargetReleaseName() string {
	if x != nil {
		return x.TargetReleaseName
	}
	return ""
}

// ClientEventSpec source resource reference: pkg/dal/table/client-event.go
type ClientEventSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalReleaseId         uint32                 `protobuf:"varint,1,opt,name=original_release_id,json=originalReleaseId,proto3" json:"original_release_id,omitempty"`
	TargetReleaseId           uint32                 `protobuf:"varint,2,opt,name=target_release_id,json=targetReleaseId,proto3" json:"target_release_id,omitempty"`
	StartTime                 *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime                   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	TotalSeconds              float64                `protobuf:"fixed64,5,opt,name=total_seconds,json=totalSeconds,proto3" json:"total_seconds,omitempty"`
	TotalFileSize             float64                `protobuf:"fixed64,6,opt,name=total_file_size,json=totalFileSize,proto3" json:"total_file_size,omitempty"`
	DownloadFileSize          float64                `protobuf:"fixed64,7,opt,name=download_file_size,json=downloadFileSize,proto3" json:"download_file_size,omitempty"`
	TotalFileNum              uint32                 `protobuf:"varint,8,opt,name=total_file_num,json=totalFileNum,proto3" json:"total_file_num,omitempty"`
	DownloadFileNum           uint32                 `protobuf:"varint,9,opt,name=download_file_num,json=downloadFileNum,proto3" json:"download_file_num,omitempty"`
	ReleaseChangeStatus       string                 `protobuf:"bytes,10,opt,name=release_change_status,json=releaseChangeStatus,proto3" json:"release_change_status,omitempty"`
	ReleaseChangeFailedReason string                 `protobuf:"bytes,11,opt,name=release_change_failed_reason,json=releaseChangeFailedReason,proto3" json:"release_change_failed_reason,omitempty"`
	FailedDetailReason        string                 `protobuf:"bytes,12,opt,name=failed_detail_reason,json=failedDetailReason,proto3" json:"failed_detail_reason,omitempty"`
	SpecificFailedReason      string                 `protobuf:"bytes,13,opt,name=specific_failed_reason,json=specificFailedReason,proto3" json:"specific_failed_reason,omitempty"`
}

func (x *ClientEventSpec) Reset() {
	*x = ClientEventSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventSpec) ProtoMessage() {}

func (x *ClientEventSpec) ProtoReflect() protoreflect.Message {
	mi := &file_client_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventSpec.ProtoReflect.Descriptor instead.
func (*ClientEventSpec) Descriptor() ([]byte, []int) {
	return file_client_event_proto_rawDescGZIP(), []int{1}
}

func (x *ClientEventSpec) GetOriginalReleaseId() uint32 {
	if x != nil {
		return x.OriginalReleaseId
	}
	return 0
}

func (x *ClientEventSpec) GetTargetReleaseId() uint32 {
	if x != nil {
		return x.TargetReleaseId
	}
	return 0
}

func (x *ClientEventSpec) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ClientEventSpec) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *ClientEventSpec) GetTotalSeconds() float64 {
	if x != nil {
		return x.TotalSeconds
	}
	return 0
}

func (x *ClientEventSpec) GetTotalFileSize() float64 {
	if x != nil {
		return x.TotalFileSize
	}
	return 0
}

func (x *ClientEventSpec) GetDownloadFileSize() float64 {
	if x != nil {
		return x.DownloadFileSize
	}
	return 0
}

func (x *ClientEventSpec) GetTotalFileNum() uint32 {
	if x != nil {
		return x.TotalFileNum
	}
	return 0
}

func (x *ClientEventSpec) GetDownloadFileNum() uint32 {
	if x != nil {
		return x.DownloadFileNum
	}
	return 0
}

func (x *ClientEventSpec) GetReleaseChangeStatus() string {
	if x != nil {
		return x.ReleaseChangeStatus
	}
	return ""
}

func (x *ClientEventSpec) GetReleaseChangeFailedReason() string {
	if x != nil {
		return x.ReleaseChangeFailedReason
	}
	return ""
}

func (x *ClientEventSpec) GetFailedDetailReason() string {
	if x != nil {
		return x.FailedDetailReason
	}
	return ""
}

func (x *ClientEventSpec) GetSpecificFailedReason() string {
	if x != nil {
		return x.SpecificFailedReason
	}
	return ""
}

// ClientEventAttachment source resource reference: pkg/dal/table/client-event.go
type ClientEventAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId   uint32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Uid        string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	BizId      uint32 `protobuf:"varint,3,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	AppId      uint32 `protobuf:"varint,4,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ClientMode string `protobuf:"bytes,5,opt,name=client_mode,json=clientMode,proto3" json:"client_mode,omitempty"`
	CursorId   string `protobuf:"bytes,6,opt,name=cursor_id,json=cursorId,proto3" json:"cursor_id,omitempty"`
}

func (x *ClientEventAttachment) Reset() {
	*x = ClientEventAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventAttachment) ProtoMessage() {}

func (x *ClientEventAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_client_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventAttachment.ProtoReflect.Descriptor instead.
func (*ClientEventAttachment) Descriptor() ([]byte, []int) {
	return file_client_event_proto_rawDescGZIP(), []int{2}
}

func (x *ClientEventAttachment) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ClientEventAttachment) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ClientEventAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *ClientEventAttachment) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ClientEventAttachment) GetClientMode() string {
	if x != nil {
		return x.ClientMode
	}
	return ""
}

func (x *ClientEventAttachment) GetCursorId() string {
	if x != nil {
		return x.CursorId
	}
	return ""
}

var File_client_event_proto protoreflect.FileDescriptor

var file_client_event_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02, 0x0a, 0x0b,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x63, 0x65,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x63,
	0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x89, 0x05,
	0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x2a, 0x0a, 0x11, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x32, 0x0a, 0x15,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3f, 0x0a, 0x1c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xb2, 0x01, 0x0a, 0x15, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x42, 0x5d,
	0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e,
	0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x75, 0x65, 0x4b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6b, 0x2d,
	0x62, 0x63, 0x73, 0x2f, 0x62, 0x63, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x62, 0x63, 0x73, 0x2d, 0x62, 0x73, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3b, 0x70, 0x62, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_event_proto_rawDescOnce sync.Once
	file_client_event_proto_rawDescData = file_client_event_proto_rawDesc
)

func file_client_event_proto_rawDescGZIP() []byte {
	file_client_event_proto_rawDescOnce.Do(func() {
		file_client_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_event_proto_rawDescData)
	})
	return file_client_event_proto_rawDescData
}

var file_client_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_client_event_proto_goTypes = []interface{}{
	(*ClientEvent)(nil),           // 0: pbce.ClientEvent
	(*ClientEventSpec)(nil),       // 1: pbce.ClientEventSpec
	(*ClientEventAttachment)(nil), // 2: pbce.ClientEventAttachment
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_client_event_proto_depIdxs = []int32{
	1, // 0: pbce.ClientEvent.spec:type_name -> pbce.ClientEventSpec
	2, // 1: pbce.ClientEvent.attachment:type_name -> pbce.ClientEventAttachment
	3, // 2: pbce.ClientEvent.heartbeat_time:type_name -> google.protobuf.Timestamp
	3, // 3: pbce.ClientEventSpec.start_time:type_name -> google.protobuf.Timestamp
	3, // 4: pbce.ClientEventSpec.end_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_client_event_proto_init() }
func file_client_event_proto_init() {
	if File_client_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEvent); i {
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
		file_client_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventSpec); i {
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
		file_client_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventAttachment); i {
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
			RawDescriptor: file_client_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_event_proto_goTypes,
		DependencyIndexes: file_client_event_proto_depIdxs,
		MessageInfos:      file_client_event_proto_msgTypes,
	}.Build()
	File_client_event_proto = out.File
	file_client_event_proto_rawDesc = nil
	file_client_event_proto_goTypes = nil
	file_client_event_proto_depIdxs = nil
}
