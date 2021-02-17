// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: src/TanksWar/protocol/requestMessage.proto

package protocol

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ServerProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId             string                              `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`          //组id
	AccountId           string                              `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`    //用户id
	ProtocolId          string                              `protobuf:"bytes,3,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"` //协议选择
	GamingProtocol      *ServerProtocol_GamingProtocol      `protobuf:"bytes,4,opt,name=gamingProtocol,proto3" json:"gamingProtocol,omitempty"`
	StartAndEndProtocol *ServerProtocol_StartAndEndProtocol `protobuf:"bytes,5,opt,name=startAndEndProtocol,proto3" json:"startAndEndProtocol,omitempty"`
	TanksInfo           *ServerProtocol_TanksInfo           `protobuf:"bytes,6,opt,name=tanksInfo,proto3" json:"tanksInfo,omitempty"`
}

func (x *ServerProtocol) Reset() {
	*x = ServerProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_TanksWar_protocol_requestMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerProtocol) ProtoMessage() {}

func (x *ServerProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_src_TanksWar_protocol_requestMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerProtocol.ProtoReflect.Descriptor instead.
func (*ServerProtocol) Descriptor() ([]byte, []int) {
	return file_src_TanksWar_protocol_requestMessage_proto_rawDescGZIP(), []int{0}
}

func (x *ServerProtocol) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ServerProtocol) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ServerProtocol) GetProtocolId() string {
	if x != nil {
		return x.ProtocolId
	}
	return ""
}

func (x *ServerProtocol) GetGamingProtocol() *ServerProtocol_GamingProtocol {
	if x != nil {
		return x.GamingProtocol
	}
	return nil
}

func (x *ServerProtocol) GetStartAndEndProtocol() *ServerProtocol_StartAndEndProtocol {
	if x != nil {
		return x.StartAndEndProtocol
	}
	return nil
}

func (x *ServerProtocol) GetTanksInfo() *ServerProtocol_TanksInfo {
	if x != nil {
		return x.TanksInfo
	}
	return nil
}

type ServerProtocol_GamingProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`                            //指令代码
	Steps     int32 `protobuf:"varint,2,opt,name=steps,proto3" json:"steps,omitempty"`                          //步数
	OldX      int32 `protobuf:"varint,3,opt,name=old_x,json=oldX,proto3" json:"old_x,omitempty"`                //旧x坐标
	OldY      int32 `protobuf:"varint,4,opt,name=old_y,json=oldY,proto3" json:"old_y,omitempty"`                //旧y坐标
	NewX      int32 `protobuf:"varint,5,opt,name=new_x,json=newX,proto3" json:"new_x,omitempty"`                //新x坐标
	NewY      int32 `protobuf:"varint,6,opt,name=new_y,json=newY,proto3" json:"new_y,omitempty"`                //新y坐标
	ScopeDir  int32 `protobuf:"varint,7,opt,name=scope_dir,json=scopeDir,proto3" json:"scope_dir,omitempty"`    //视镜方向
	WeaponDir int32 `protobuf:"varint,8,opt,name=weapon_dir,json=weaponDir,proto3" json:"weapon_dir,omitempty"` //炮筒方向
	TankDir   int32 `protobuf:"varint,9,opt,name=tank_dir,json=tankDir,proto3" json:"tank_dir,omitempty"`       //坦克方向
	Num       int32 `protobuf:"varint,10,opt,name=num,proto3" json:"num,omitempty"`                             //比赛人数
}

func (x *ServerProtocol_GamingProtocol) Reset() {
	*x = ServerProtocol_GamingProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_TanksWar_protocol_requestMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerProtocol_GamingProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerProtocol_GamingProtocol) ProtoMessage() {}

func (x *ServerProtocol_GamingProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_src_TanksWar_protocol_requestMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerProtocol_GamingProtocol.ProtoReflect.Descriptor instead.
func (*ServerProtocol_GamingProtocol) Descriptor() ([]byte, []int) {
	return file_src_TanksWar_protocol_requestMessage_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServerProtocol_GamingProtocol) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ServerProtocol_GamingProtocol) GetSteps() int32 {
	if x != nil {
		return x.Steps
	}
	return 0
}

func (x *ServerProtocol_GamingProtocol) GetOldX() int32 {
	if x != nil {
		return x.OldX
	}
	return 0
}

func (x *ServerProtocol_GamingProtocol) GetOldY() int32 {
	if x != nil {
		return x.OldY
	}
	return 0
}

func (x *ServerProtocol_GamingProtocol) GetNewX() int32 {
	if x != nil {
		return x.NewX
	}
	return 0
}

func (x *ServerProtocol_GamingProtocol) GetNewY() int32 {
	if x != nil {
		return x.NewY
	}
	return 0
}

func (x *ServerProtocol_GamingProtocol) GetScopeDir() int32 {
	if x != nil {
		return x.ScopeDir
	}
	return 0
}

func (x *ServerProtocol_GamingProtocol) GetWeaponDir() int32 {
	if x != nil {
		return x.WeaponDir
	}
	return 0
}

func (x *ServerProtocol_GamingProtocol) GetTankDir() int32 {
	if x != nil {
		return x.TankDir
	}
	return 0
}

func (x *ServerProtocol_GamingProtocol) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ServerProtocol_StartAndEndProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Live int32 `protobuf:"varint,1,opt,name=live,proto3" json:"live,omitempty"` //用户存亡
}

func (x *ServerProtocol_StartAndEndProtocol) Reset() {
	*x = ServerProtocol_StartAndEndProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_TanksWar_protocol_requestMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerProtocol_StartAndEndProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerProtocol_StartAndEndProtocol) ProtoMessage() {}

func (x *ServerProtocol_StartAndEndProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_src_TanksWar_protocol_requestMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerProtocol_StartAndEndProtocol.ProtoReflect.Descriptor instead.
func (*ServerProtocol_StartAndEndProtocol) Descriptor() ([]byte, []int) {
	return file_src_TanksWar_protocol_requestMessage_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ServerProtocol_StartAndEndProtocol) GetLive() int32 {
	if x != nil {
		return x.Live
	}
	return 0
}

type ServerProtocol_TanksInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Engine int32 `protobuf:"varint,1,opt,name=engine,proto3" json:"engine,omitempty"` //引擎
	Weapon int32 `protobuf:"varint,2,opt,name=weapon,proto3" json:"weapon,omitempty"` //炮筒
	Scope  int32 `protobuf:"varint,3,opt,name=scope,proto3" json:"scope,omitempty"`   //视镜
	Armor  int32 `protobuf:"varint,4,opt,name=armor,proto3" json:"armor,omitempty"`   //装甲
	Bullet int32 `protobuf:"varint,5,opt,name=bullet,proto3" json:"bullet,omitempty"` //子弹
}

func (x *ServerProtocol_TanksInfo) Reset() {
	*x = ServerProtocol_TanksInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_TanksWar_protocol_requestMessage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerProtocol_TanksInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerProtocol_TanksInfo) ProtoMessage() {}

func (x *ServerProtocol_TanksInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_TanksWar_protocol_requestMessage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerProtocol_TanksInfo.ProtoReflect.Descriptor instead.
func (*ServerProtocol_TanksInfo) Descriptor() ([]byte, []int) {
	return file_src_TanksWar_protocol_requestMessage_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ServerProtocol_TanksInfo) GetEngine() int32 {
	if x != nil {
		return x.Engine
	}
	return 0
}

func (x *ServerProtocol_TanksInfo) GetWeapon() int32 {
	if x != nil {
		return x.Weapon
	}
	return 0
}

func (x *ServerProtocol_TanksInfo) GetScope() int32 {
	if x != nil {
		return x.Scope
	}
	return 0
}

func (x *ServerProtocol_TanksInfo) GetArmor() int32 {
	if x != nil {
		return x.Armor
	}
	return 0
}

func (x *ServerProtocol_TanksInfo) GetBullet() int32 {
	if x != nil {
		return x.Bullet
	}
	return 0
}

var File_src_TanksWar_protocol_requestMessage_proto protoreflect.FileDescriptor

var file_src_TanksWar_protocol_requestMessage_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x72, 0x63, 0x2f, 0x54, 0x61, 0x6e, 0x6b, 0x73, 0x57, 0x61, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x84, 0x06, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x0e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x5e, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x6e,
	0x64, 0x45, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x52, 0x13, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x40, 0x0a, 0x09, 0x74, 0x61, 0x6e, 0x6b, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x54, 0x61, 0x6e, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x61,
	0x6e, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0xf7, 0x01, 0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x74, 0x65, 0x70, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x6f, 0x6c, 0x64, 0x5f, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6f, 0x6c, 0x64, 0x58, 0x12, 0x13, 0x0a, 0x05, 0x6f, 0x6c, 0x64,
	0x5f, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6f, 0x6c, 0x64, 0x59, 0x12, 0x13,
	0x0a, 0x05, 0x6e, 0x65, 0x77, 0x5f, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e,
	0x65, 0x77, 0x58, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x65, 0x77, 0x5f, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x59, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x44, 0x69, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f,
	0x64, 0x69, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x44, 0x69, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x6e, 0x6b, 0x5f, 0x64, 0x69, 0x72,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x61, 0x6e, 0x6b, 0x44, 0x69, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x1a, 0x29, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x76, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x1a, 0x7f, 0x0a, 0x09,
	0x54, 0x61, 0x6e, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x61, 0x72, 0x6d, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_TanksWar_protocol_requestMessage_proto_rawDescOnce sync.Once
	file_src_TanksWar_protocol_requestMessage_proto_rawDescData = file_src_TanksWar_protocol_requestMessage_proto_rawDesc
)

func file_src_TanksWar_protocol_requestMessage_proto_rawDescGZIP() []byte {
	file_src_TanksWar_protocol_requestMessage_proto_rawDescOnce.Do(func() {
		file_src_TanksWar_protocol_requestMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_TanksWar_protocol_requestMessage_proto_rawDescData)
	})
	return file_src_TanksWar_protocol_requestMessage_proto_rawDescData
}

var file_src_TanksWar_protocol_requestMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_src_TanksWar_protocol_requestMessage_proto_goTypes = []interface{}{
	(*ServerProtocol)(nil),                     // 0: protocol.ServerProtocol
	(*ServerProtocol_GamingProtocol)(nil),      // 1: protocol.ServerProtocol.GamingProtocol
	(*ServerProtocol_StartAndEndProtocol)(nil), // 2: protocol.ServerProtocol.StartAndEndProtocol
	(*ServerProtocol_TanksInfo)(nil),           // 3: protocol.ServerProtocol.TanksInfo
}
var file_src_TanksWar_protocol_requestMessage_proto_depIdxs = []int32{
	1, // 0: protocol.ServerProtocol.gamingProtocol:type_name -> protocol.ServerProtocol.GamingProtocol
	2, // 1: protocol.ServerProtocol.startAndEndProtocol:type_name -> protocol.ServerProtocol.StartAndEndProtocol
	3, // 2: protocol.ServerProtocol.tanksInfo:type_name -> protocol.ServerProtocol.TanksInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_src_TanksWar_protocol_requestMessage_proto_init() }
func file_src_TanksWar_protocol_requestMessage_proto_init() {
	if File_src_TanksWar_protocol_requestMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_TanksWar_protocol_requestMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerProtocol); i {
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
		file_src_TanksWar_protocol_requestMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerProtocol_GamingProtocol); i {
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
		file_src_TanksWar_protocol_requestMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerProtocol_StartAndEndProtocol); i {
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
		file_src_TanksWar_protocol_requestMessage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerProtocol_TanksInfo); i {
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
			RawDescriptor: file_src_TanksWar_protocol_requestMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_TanksWar_protocol_requestMessage_proto_goTypes,
		DependencyIndexes: file_src_TanksWar_protocol_requestMessage_proto_depIdxs,
		MessageInfos:      file_src_TanksWar_protocol_requestMessage_proto_msgTypes,
	}.Build()
	File_src_TanksWar_protocol_requestMessage_proto = out.File
	file_src_TanksWar_protocol_requestMessage_proto_rawDesc = nil
	file_src_TanksWar_protocol_requestMessage_proto_goTypes = nil
	file_src_TanksWar_protocol_requestMessage_proto_depIdxs = nil
}