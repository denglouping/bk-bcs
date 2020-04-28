// Code generated by protoc-gen-go.
// source: gse-controller.proto
// DO NOT EDIT!

/*
Package gsecontroller is a generated protocol buffer package.

It is generated from these files:
	gse-controller.proto

It has these top-level messages:
	PublishReleasePreReq
	PublishReleasePreResp
	PublishReleaseReq
	PublishReleaseResp
	RollbackReleaseReq
	RollbackReleaseResp
	ReportReq
	ReportResp
	PullReleaseReq
	PullReleaseResp
	PullConfigSetListReq
	PullConfigSetListResp
*/
package gsecontroller

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bk-bscp/internal/protocol/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PublishReleasePreReq struct {
	Seq       uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Bid       string `protobuf:"bytes,2,opt,name=bid" json:"bid,omitempty"`
	Releaseid string `protobuf:"bytes,3,opt,name=releaseid" json:"releaseid,omitempty"`
	Operator  string `protobuf:"bytes,4,opt,name=operator" json:"operator,omitempty"`
}

func (m *PublishReleasePreReq) Reset()                    { *m = PublishReleasePreReq{} }
func (m *PublishReleasePreReq) String() string            { return proto.CompactTextString(m) }
func (*PublishReleasePreReq) ProtoMessage()               {}
func (*PublishReleasePreReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PublishReleasePreReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PublishReleasePreReq) GetBid() string {
	if m != nil {
		return m.Bid
	}
	return ""
}

func (m *PublishReleasePreReq) GetReleaseid() string {
	if m != nil {
		return m.Releaseid
	}
	return ""
}

func (m *PublishReleasePreReq) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type PublishReleasePreResp struct {
	Seq     uint64         `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode common.ErrCode `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg  string         `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *PublishReleasePreResp) Reset()                    { *m = PublishReleasePreResp{} }
func (m *PublishReleasePreResp) String() string            { return proto.CompactTextString(m) }
func (*PublishReleasePreResp) ProtoMessage()               {}
func (*PublishReleasePreResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PublishReleasePreResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PublishReleasePreResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *PublishReleasePreResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type PublishReleaseReq struct {
	Seq       uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Bid       string `protobuf:"bytes,2,opt,name=bid" json:"bid,omitempty"`
	Releaseid string `protobuf:"bytes,3,opt,name=releaseid" json:"releaseid,omitempty"`
	Operator  string `protobuf:"bytes,4,opt,name=operator" json:"operator,omitempty"`
}

func (m *PublishReleaseReq) Reset()                    { *m = PublishReleaseReq{} }
func (m *PublishReleaseReq) String() string            { return proto.CompactTextString(m) }
func (*PublishReleaseReq) ProtoMessage()               {}
func (*PublishReleaseReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PublishReleaseReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PublishReleaseReq) GetBid() string {
	if m != nil {
		return m.Bid
	}
	return ""
}

func (m *PublishReleaseReq) GetReleaseid() string {
	if m != nil {
		return m.Releaseid
	}
	return ""
}

func (m *PublishReleaseReq) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type PublishReleaseResp struct {
	Seq     uint64         `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode common.ErrCode `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg  string         `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *PublishReleaseResp) Reset()                    { *m = PublishReleaseResp{} }
func (m *PublishReleaseResp) String() string            { return proto.CompactTextString(m) }
func (*PublishReleaseResp) ProtoMessage()               {}
func (*PublishReleaseResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PublishReleaseResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PublishReleaseResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *PublishReleaseResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type RollbackReleaseReq struct {
	Seq       uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Bid       string `protobuf:"bytes,2,opt,name=bid" json:"bid,omitempty"`
	Releaseid string `protobuf:"bytes,3,opt,name=releaseid" json:"releaseid,omitempty"`
	Operator  string `protobuf:"bytes,4,opt,name=operator" json:"operator,omitempty"`
}

func (m *RollbackReleaseReq) Reset()                    { *m = RollbackReleaseReq{} }
func (m *RollbackReleaseReq) String() string            { return proto.CompactTextString(m) }
func (*RollbackReleaseReq) ProtoMessage()               {}
func (*RollbackReleaseReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RollbackReleaseReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *RollbackReleaseReq) GetBid() string {
	if m != nil {
		return m.Bid
	}
	return ""
}

func (m *RollbackReleaseReq) GetReleaseid() string {
	if m != nil {
		return m.Releaseid
	}
	return ""
}

func (m *RollbackReleaseReq) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RollbackReleaseResp struct {
	Seq     uint64         `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode common.ErrCode `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg  string         `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *RollbackReleaseResp) Reset()                    { *m = RollbackReleaseResp{} }
func (m *RollbackReleaseResp) String() string            { return proto.CompactTextString(m) }
func (*RollbackReleaseResp) ProtoMessage()               {}
func (*RollbackReleaseResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RollbackReleaseResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *RollbackReleaseResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *RollbackReleaseResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type ReportReq struct {
	Seq       uint64               `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Bid       string               `protobuf:"bytes,2,opt,name=bid" json:"bid,omitempty"`
	Appid     string               `protobuf:"bytes,3,opt,name=appid" json:"appid,omitempty"`
	Clusterid string               `protobuf:"bytes,4,opt,name=clusterid" json:"clusterid,omitempty"`
	Zoneid    string               `protobuf:"bytes,5,opt,name=zoneid" json:"zoneid,omitempty"`
	Dc        string               `protobuf:"bytes,6,opt,name=dc" json:"dc,omitempty"`
	IP        string               `protobuf:"bytes,7,opt,name=IP" json:"IP,omitempty"`
	Labels    string               `protobuf:"bytes,8,opt,name=labels" json:"labels,omitempty"`
	Infos     []*common.ReportInfo `protobuf:"bytes,9,rep,name=infos" json:"infos,omitempty"`
}

func (m *ReportReq) Reset()                    { *m = ReportReq{} }
func (m *ReportReq) String() string            { return proto.CompactTextString(m) }
func (*ReportReq) ProtoMessage()               {}
func (*ReportReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReportReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ReportReq) GetBid() string {
	if m != nil {
		return m.Bid
	}
	return ""
}

func (m *ReportReq) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *ReportReq) GetClusterid() string {
	if m != nil {
		return m.Clusterid
	}
	return ""
}

func (m *ReportReq) GetZoneid() string {
	if m != nil {
		return m.Zoneid
	}
	return ""
}

func (m *ReportReq) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

func (m *ReportReq) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *ReportReq) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *ReportReq) GetInfos() []*common.ReportInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

type ReportResp struct {
	Seq     uint64         `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode common.ErrCode `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg  string         `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *ReportResp) Reset()                    { *m = ReportResp{} }
func (m *ReportResp) String() string            { return proto.CompactTextString(m) }
func (*ReportResp) ProtoMessage()               {}
func (*ReportResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReportResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ReportResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *ReportResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type PullReleaseReq struct {
	Seq            uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Bid            string `protobuf:"bytes,2,opt,name=bid" json:"bid,omitempty"`
	Appid          string `protobuf:"bytes,3,opt,name=appid" json:"appid,omitempty"`
	Clusterid      string `protobuf:"bytes,4,opt,name=clusterid" json:"clusterid,omitempty"`
	Zoneid         string `protobuf:"bytes,5,opt,name=zoneid" json:"zoneid,omitempty"`
	Dc             string `protobuf:"bytes,6,opt,name=dc" json:"dc,omitempty"`
	IP             string `protobuf:"bytes,7,opt,name=IP" json:"IP,omitempty"`
	Labels         string `protobuf:"bytes,8,opt,name=labels" json:"labels,omitempty"`
	Cfgsetid       string `protobuf:"bytes,9,opt,name=cfgsetid" json:"cfgsetid,omitempty"`
	LocalReleaseid string `protobuf:"bytes,10,opt,name=localReleaseid" json:"localReleaseid,omitempty"`
	Releaseid      string `protobuf:"bytes,11,opt,name=releaseid" json:"releaseid,omitempty"`
}

func (m *PullReleaseReq) Reset()                    { *m = PullReleaseReq{} }
func (m *PullReleaseReq) String() string            { return proto.CompactTextString(m) }
func (*PullReleaseReq) ProtoMessage()               {}
func (*PullReleaseReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PullReleaseReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PullReleaseReq) GetBid() string {
	if m != nil {
		return m.Bid
	}
	return ""
}

func (m *PullReleaseReq) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *PullReleaseReq) GetClusterid() string {
	if m != nil {
		return m.Clusterid
	}
	return ""
}

func (m *PullReleaseReq) GetZoneid() string {
	if m != nil {
		return m.Zoneid
	}
	return ""
}

func (m *PullReleaseReq) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

func (m *PullReleaseReq) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *PullReleaseReq) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *PullReleaseReq) GetCfgsetid() string {
	if m != nil {
		return m.Cfgsetid
	}
	return ""
}

func (m *PullReleaseReq) GetLocalReleaseid() string {
	if m != nil {
		return m.LocalReleaseid
	}
	return ""
}

func (m *PullReleaseReq) GetReleaseid() string {
	if m != nil {
		return m.Releaseid
	}
	return ""
}

type PullReleaseResp struct {
	Seq     uint64          `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode common.ErrCode  `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg  string          `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
	Release *common.Release `protobuf:"bytes,4,opt,name=release" json:"release,omitempty"`
}

func (m *PullReleaseResp) Reset()                    { *m = PullReleaseResp{} }
func (m *PullReleaseResp) String() string            { return proto.CompactTextString(m) }
func (*PullReleaseResp) ProtoMessage()               {}
func (*PullReleaseResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PullReleaseResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PullReleaseResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *PullReleaseResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *PullReleaseResp) GetRelease() *common.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

type PullConfigSetListReq struct {
	Seq   uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Bid   string `protobuf:"bytes,2,opt,name=bid" json:"bid,omitempty"`
	Appid string `protobuf:"bytes,3,opt,name=appid" json:"appid,omitempty"`
}

func (m *PullConfigSetListReq) Reset()                    { *m = PullConfigSetListReq{} }
func (m *PullConfigSetListReq) String() string            { return proto.CompactTextString(m) }
func (*PullConfigSetListReq) ProtoMessage()               {}
func (*PullConfigSetListReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PullConfigSetListReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PullConfigSetListReq) GetBid() string {
	if m != nil {
		return m.Bid
	}
	return ""
}

func (m *PullConfigSetListReq) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

type PullConfigSetListResp struct {
	Seq        uint64              `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode    common.ErrCode      `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg     string              `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
	ConfigSets []*common.ConfigSet `protobuf:"bytes,4,rep,name=configSets" json:"configSets,omitempty"`
}

func (m *PullConfigSetListResp) Reset()                    { *m = PullConfigSetListResp{} }
func (m *PullConfigSetListResp) String() string            { return proto.CompactTextString(m) }
func (*PullConfigSetListResp) ProtoMessage()               {}
func (*PullConfigSetListResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PullConfigSetListResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PullConfigSetListResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *PullConfigSetListResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *PullConfigSetListResp) GetConfigSets() []*common.ConfigSet {
	if m != nil {
		return m.ConfigSets
	}
	return nil
}

func init() {
	proto.RegisterType((*PublishReleasePreReq)(nil), "gsecontroller.PublishReleasePreReq")
	proto.RegisterType((*PublishReleasePreResp)(nil), "gsecontroller.PublishReleasePreResp")
	proto.RegisterType((*PublishReleaseReq)(nil), "gsecontroller.PublishReleaseReq")
	proto.RegisterType((*PublishReleaseResp)(nil), "gsecontroller.PublishReleaseResp")
	proto.RegisterType((*RollbackReleaseReq)(nil), "gsecontroller.RollbackReleaseReq")
	proto.RegisterType((*RollbackReleaseResp)(nil), "gsecontroller.RollbackReleaseResp")
	proto.RegisterType((*ReportReq)(nil), "gsecontroller.ReportReq")
	proto.RegisterType((*ReportResp)(nil), "gsecontroller.ReportResp")
	proto.RegisterType((*PullReleaseReq)(nil), "gsecontroller.PullReleaseReq")
	proto.RegisterType((*PullReleaseResp)(nil), "gsecontroller.PullReleaseResp")
	proto.RegisterType((*PullConfigSetListReq)(nil), "gsecontroller.PullConfigSetListReq")
	proto.RegisterType((*PullConfigSetListResp)(nil), "gsecontroller.PullConfigSetListResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GSEController service

type GSEControllerClient interface {
	PublishReleasePre(ctx context.Context, in *PublishReleasePreReq, opts ...grpc.CallOption) (*PublishReleasePreResp, error)
	PublishRelease(ctx context.Context, in *PublishReleaseReq, opts ...grpc.CallOption) (*PublishReleaseResp, error)
	RollbackRelease(ctx context.Context, in *RollbackReleaseReq, opts ...grpc.CallOption) (*RollbackReleaseResp, error)
	Report(ctx context.Context, in *ReportReq, opts ...grpc.CallOption) (*ReportResp, error)
	PullRelease(ctx context.Context, in *PullReleaseReq, opts ...grpc.CallOption) (*PullReleaseResp, error)
	PullConfigSetList(ctx context.Context, in *PullConfigSetListReq, opts ...grpc.CallOption) (*PullConfigSetListResp, error)
}

type gSEControllerClient struct {
	cc *grpc.ClientConn
}

func NewGSEControllerClient(cc *grpc.ClientConn) GSEControllerClient {
	return &gSEControllerClient{cc}
}

func (c *gSEControllerClient) PublishReleasePre(ctx context.Context, in *PublishReleasePreReq, opts ...grpc.CallOption) (*PublishReleasePreResp, error) {
	out := new(PublishReleasePreResp)
	err := grpc.Invoke(ctx, "/gsecontroller.GSEController/PublishReleasePre", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSEControllerClient) PublishRelease(ctx context.Context, in *PublishReleaseReq, opts ...grpc.CallOption) (*PublishReleaseResp, error) {
	out := new(PublishReleaseResp)
	err := grpc.Invoke(ctx, "/gsecontroller.GSEController/PublishRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSEControllerClient) RollbackRelease(ctx context.Context, in *RollbackReleaseReq, opts ...grpc.CallOption) (*RollbackReleaseResp, error) {
	out := new(RollbackReleaseResp)
	err := grpc.Invoke(ctx, "/gsecontroller.GSEController/RollbackRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSEControllerClient) Report(ctx context.Context, in *ReportReq, opts ...grpc.CallOption) (*ReportResp, error) {
	out := new(ReportResp)
	err := grpc.Invoke(ctx, "/gsecontroller.GSEController/Report", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSEControllerClient) PullRelease(ctx context.Context, in *PullReleaseReq, opts ...grpc.CallOption) (*PullReleaseResp, error) {
	out := new(PullReleaseResp)
	err := grpc.Invoke(ctx, "/gsecontroller.GSEController/PullRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSEControllerClient) PullConfigSetList(ctx context.Context, in *PullConfigSetListReq, opts ...grpc.CallOption) (*PullConfigSetListResp, error) {
	out := new(PullConfigSetListResp)
	err := grpc.Invoke(ctx, "/gsecontroller.GSEController/PullConfigSetList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GSEController service

type GSEControllerServer interface {
	PublishReleasePre(context.Context, *PublishReleasePreReq) (*PublishReleasePreResp, error)
	PublishRelease(context.Context, *PublishReleaseReq) (*PublishReleaseResp, error)
	RollbackRelease(context.Context, *RollbackReleaseReq) (*RollbackReleaseResp, error)
	Report(context.Context, *ReportReq) (*ReportResp, error)
	PullRelease(context.Context, *PullReleaseReq) (*PullReleaseResp, error)
	PullConfigSetList(context.Context, *PullConfigSetListReq) (*PullConfigSetListResp, error)
}

func RegisterGSEControllerServer(s *grpc.Server, srv GSEControllerServer) {
	s.RegisterService(&_GSEController_serviceDesc, srv)
}

func _GSEController_PublishReleasePre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishReleasePreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSEControllerServer).PublishReleasePre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsecontroller.GSEController/PublishReleasePre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSEControllerServer).PublishReleasePre(ctx, req.(*PublishReleasePreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSEController_PublishRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSEControllerServer).PublishRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsecontroller.GSEController/PublishRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSEControllerServer).PublishRelease(ctx, req.(*PublishReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSEController_RollbackRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSEControllerServer).RollbackRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsecontroller.GSEController/RollbackRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSEControllerServer).RollbackRelease(ctx, req.(*RollbackReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSEController_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSEControllerServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsecontroller.GSEController/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSEControllerServer).Report(ctx, req.(*ReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSEController_PullRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSEControllerServer).PullRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsecontroller.GSEController/PullRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSEControllerServer).PullRelease(ctx, req.(*PullReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSEController_PullConfigSetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullConfigSetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSEControllerServer).PullConfigSetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsecontroller.GSEController/PullConfigSetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSEControllerServer).PullConfigSetList(ctx, req.(*PullConfigSetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GSEController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gsecontroller.GSEController",
	HandlerType: (*GSEControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishReleasePre",
			Handler:    _GSEController_PublishReleasePre_Handler,
		},
		{
			MethodName: "PublishRelease",
			Handler:    _GSEController_PublishRelease_Handler,
		},
		{
			MethodName: "RollbackRelease",
			Handler:    _GSEController_RollbackRelease_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _GSEController_Report_Handler,
		},
		{
			MethodName: "PullRelease",
			Handler:    _GSEController_PullRelease_Handler,
		},
		{
			MethodName: "PullConfigSetList",
			Handler:    _GSEController_PullConfigSetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gse-controller.proto",
}

func init() { proto.RegisterFile("gse-controller.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x9d, 0x74, 0x3a, 0x9d, 0xf6, 0x56, 0xd3, 0x32, 0xa6, 0x20, 0x13, 0x01, 0x2a, 0x01, 0xa1,
	0xb2, 0x68, 0x2b, 0xca, 0x07, 0xb0, 0xa8, 0x46, 0xa8, 0x12, 0xa0, 0x28, 0xb3, 0x40, 0x62, 0x45,
	0x1e, 0x6e, 0x31, 0xe3, 0x89, 0x53, 0x3b, 0xdd, 0xf0, 0x17, 0x2c, 0x61, 0xc3, 0x97, 0xf1, 0x2f,
	0xc8, 0x76, 0xd2, 0x47, 0x1a, 0xa6, 0x20, 0x14, 0xc4, 0xaa, 0xb9, 0xcf, 0x93, 0x73, 0x7d, 0x72,
	0x5d, 0xe8, 0x2d, 0x24, 0x19, 0x86, 0x3c, 0x4e, 0x05, 0x67, 0x8c, 0x88, 0x51, 0x22, 0x78, 0xca,
	0xd1, 0xd9, 0x42, 0x92, 0x8d, 0xd3, 0x1e, 0x06, 0x57, 0xc3, 0x40, 0x86, 0xc9, 0x98, 0xc6, 0x29,
	0x11, 0xb1, 0xcf, 0xc6, 0x3a, 0x2d, 0xe4, 0x6c, 0x1c, 0xf2, 0xeb, 0x6b, 0x1e, 0x67, 0x3f, 0xa6,
	0xda, 0x49, 0xa1, 0xe7, 0xae, 0x02, 0x46, 0xe5, 0x47, 0x8f, 0x30, 0xe2, 0x4b, 0xe2, 0x0a, 0xe2,
	0x91, 0x25, 0xba, 0x05, 0xc7, 0x92, 0x2c, 0xb1, 0xd5, 0xb7, 0x06, 0x75, 0x4f, 0x3d, 0x2a, 0x4f,
	0x40, 0x23, 0x5c, 0xeb, 0x5b, 0x83, 0x96, 0xa7, 0x1e, 0xd1, 0x7d, 0x68, 0x09, 0x53, 0x44, 0x23,
	0x7c, 0xac, 0xfd, 0x1b, 0x07, 0xb2, 0xa1, 0xc9, 0x13, 0x22, 0xfc, 0x94, 0x0b, 0x5c, 0xd7, 0xc1,
	0xb5, 0xed, 0x30, 0xb8, 0x53, 0x82, 0x2a, 0x93, 0x12, 0xd8, 0x67, 0x70, 0x4a, 0x84, 0x98, 0xf2,
	0x88, 0x68, 0xe8, 0xce, 0xa4, 0x3b, 0xca, 0x08, 0x5c, 0x18, 0xb7, 0x97, 0xc7, 0xd1, 0x5d, 0x68,
	0x10, 0x21, 0xde, 0xc8, 0x45, 0xf6, 0x32, 0x99, 0xe5, 0x2c, 0xe1, 0x7c, 0x17, 0xad, 0x7a, 0x82,
	0x14, 0x50, 0x11, 0xb2, 0x2a, 0x76, 0x02, 0x90, 0xc7, 0x19, 0x0b, 0xfc, 0xf0, 0xea, 0x9f, 0xd1,
	0xfb, 0x04, 0xb7, 0xf7, 0x30, 0xab, 0xe2, 0xf7, 0xc3, 0x82, 0x96, 0x47, 0x12, 0x2e, 0xd2, 0xdf,
	0xe5, 0xd5, 0x83, 0x13, 0x3f, 0x49, 0xd6, 0x9c, 0x8c, 0xa1, 0xd8, 0x86, 0x6c, 0x25, 0x53, 0x22,
	0x68, 0x94, 0x11, 0xda, 0x38, 0x14, 0xfa, 0x67, 0x1e, 0xab, 0x41, 0x9c, 0x18, 0x74, 0x63, 0xa1,
	0x0e, 0xd4, 0xa2, 0x10, 0x37, 0xb4, 0xaf, 0x16, 0x85, 0xca, 0x9e, 0xb9, 0xf8, 0xd4, 0xd8, 0x33,
	0x57, 0xd5, 0x31, 0x3f, 0x20, 0x4c, 0xe2, 0xa6, 0xa9, 0x33, 0x16, 0x1a, 0xc0, 0x09, 0x8d, 0xe7,
	0x5c, 0xe2, 0x56, 0xff, 0x78, 0xd0, 0x9e, 0xa0, 0x9c, 0xb6, 0x61, 0x32, 0x8b, 0xe7, 0xdc, 0x33,
	0x09, 0x8e, 0x0f, 0x90, 0xd3, 0xab, 0x6a, 0x84, 0xdf, 0x6a, 0xd0, 0x71, 0x57, 0x8c, 0xfd, 0xa1,
	0x3e, 0xfe, 0x87, 0x39, 0xda, 0xd0, 0x0c, 0xe7, 0x0b, 0x49, 0x52, 0x1a, 0xe1, 0x96, 0x51, 0x61,
	0x6e, 0xa3, 0xa7, 0xd0, 0x61, 0x3c, 0xf4, 0x73, 0x5a, 0x34, 0xc2, 0xa0, 0x33, 0x0a, 0xde, 0x5d,
	0x9d, 0xb7, 0x0b, 0x3a, 0x77, 0xbe, 0x58, 0xd0, 0xdd, 0x19, 0x4e, 0x45, 0xa7, 0xa0, 0x5a, 0x64,
	0xa8, 0x7a, 0x6c, 0xed, 0x4d, 0x8b, 0x1c, 0x3a, 0x8f, 0x3b, 0xae, 0xda, 0xca, 0x8c, 0x4d, 0x79,
	0x3c, 0xa7, 0x8b, 0x4b, 0x92, 0xbe, 0xa6, 0xf2, 0xef, 0xd4, 0xef, 0x7c, 0xb7, 0xd4, 0xca, 0xdd,
	0x6b, 0x59, 0x15, 0xd7, 0xe7, 0x00, 0x61, 0x8e, 0x24, 0x71, 0x5d, 0x7f, 0x03, 0xe7, 0x79, 0x97,
	0xf5, 0x3b, 0x78, 0x5b, 0x49, 0x93, 0xaf, 0x75, 0x38, 0x7b, 0x75, 0x79, 0x31, 0x5d, 0x5f, 0x65,
	0xe8, 0x43, 0x71, 0x6f, 0xbb, 0x82, 0xa0, 0xc7, 0xa3, 0x9d, 0xfb, 0x6e, 0x54, 0x76, 0x7b, 0xd9,
	0x4f, 0x0e, 0x27, 0xc9, 0xc4, 0x39, 0x42, 0xef, 0xd4, 0x77, 0xb1, 0x1d, 0x42, 0xfd, 0x1b, 0x2b,
	0x55, 0xef, 0x47, 0x07, 0x32, 0x74, 0xe3, 0xf7, 0xd0, 0x2d, 0x2c, 0x48, 0x54, 0xac, 0xdb, 0x5f,
	0xda, 0xb6, 0x73, 0x28, 0x45, 0xf7, 0x7e, 0x09, 0x0d, 0xb3, 0x30, 0x10, 0x2e, 0xe6, 0xe7, 0x6b,
	0xd2, 0xbe, 0xf7, 0x8b, 0x88, 0x6e, 0xf0, 0x16, 0xda, 0x5b, 0x82, 0x47, 0x0f, 0xf6, 0x08, 0x6d,
	0x6f, 0x0a, 0xfb, 0xe1, 0x4d, 0x61, 0xdd, 0x4f, 0x9f, 0x53, 0x41, 0x5a, 0x25, 0xe7, 0xb4, 0xaf,
	0xe7, 0x92, 0x73, 0x2a, 0x51, 0xa8, 0x73, 0x14, 0x34, 0xf4, 0x9f, 0x95, 0x17, 0x3f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xa7, 0x21, 0xdd, 0x1d, 0x02, 0x09, 0x00, 0x00,
}
