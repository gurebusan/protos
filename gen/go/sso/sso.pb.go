// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: sso/sso.proto

package ssov1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_sso_sso_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_sso_sso_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type LoginRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Identifier:
	//
	//	*LoginRequest_Email
	//	*LoginRequest_Username
	Identifier    isLoginRequest_Identifier `protobuf_oneof:"identifier"`
	Password      string                    `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	AppId         int64                     `protobuf:"varint,4,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_sso_sso_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetIdentifier() isLoginRequest_Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		if x, ok := x.Identifier.(*LoginRequest_Email); ok {
			return x.Email
		}
	}
	return ""
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		if x, ok := x.Identifier.(*LoginRequest_Username); ok {
			return x.Username
		}
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type isLoginRequest_Identifier interface {
	isLoginRequest_Identifier()
}

type LoginRequest_Email struct {
	Email string `protobuf:"bytes,1,opt,name=email,proto3,oneof"`
}

type LoginRequest_Username struct {
	Username string `protobuf:"bytes,2,opt,name=username,proto3,oneof"`
}

func (*LoginRequest_Email) isLoginRequest_Identifier() {}

func (*LoginRequest_Username) isLoginRequest_Identifier() {}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_sso_sso_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type IsAdminRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RequstingUserId int64                  `protobuf:"varint,1,opt,name=requsting_user_id,json=requstingUserId,proto3" json:"requsting_user_id,omitempty"`
	TargetUserId    int64                  `protobuf:"varint,2,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *IsAdminRequest) Reset() {
	*x = IsAdminRequest{}
	mi := &file_sso_sso_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminRequest) ProtoMessage() {}

func (x *IsAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminRequest.ProtoReflect.Descriptor instead.
func (*IsAdminRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{4}
}

func (x *IsAdminRequest) GetRequstingUserId() int64 {
	if x != nil {
		return x.RequstingUserId
	}
	return 0
}

func (x *IsAdminRequest) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

type IsAdminResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsAdmin       bool                   `protobuf:"varint,1,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsAdminResponse) Reset() {
	*x = IsAdminResponse{}
	mi := &file_sso_sso_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminResponse) ProtoMessage() {}

func (x *IsAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminResponse.ProtoReflect.Descriptor instead.
func (*IsAdminResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{5}
}

func (x *IsAdminResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type GrantAdminRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RequstingUserId int64                  `protobuf:"varint,1,opt,name=requsting_user_id,json=requstingUserId,proto3" json:"requsting_user_id,omitempty"`
	TargetUserId    int64                  `protobuf:"varint,2,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GrantAdminRequest) Reset() {
	*x = GrantAdminRequest{}
	mi := &file_sso_sso_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantAdminRequest) ProtoMessage() {}

func (x *GrantAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantAdminRequest.ProtoReflect.Descriptor instead.
func (*GrantAdminRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{6}
}

func (x *GrantAdminRequest) GetRequstingUserId() int64 {
	if x != nil {
		return x.RequstingUserId
	}
	return 0
}

func (x *GrantAdminRequest) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

type GrantAdminResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrantAdminResponse) Reset() {
	*x = GrantAdminResponse{}
	mi := &file_sso_sso_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantAdminResponse) ProtoMessage() {}

func (x *GrantAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantAdminResponse.ProtoReflect.Descriptor instead.
func (*GrantAdminResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{7}
}

func (x *GrantAdminResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RevokeAdminRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RequstingUserId int64                  `protobuf:"varint,1,opt,name=requsting_user_id,json=requstingUserId,proto3" json:"requsting_user_id,omitempty"`
	TargetUserId    int64                  `protobuf:"varint,2,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RevokeAdminRequest) Reset() {
	*x = RevokeAdminRequest{}
	mi := &file_sso_sso_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeAdminRequest) ProtoMessage() {}

func (x *RevokeAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeAdminRequest.ProtoReflect.Descriptor instead.
func (*RevokeAdminRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{8}
}

func (x *RevokeAdminRequest) GetRequstingUserId() int64 {
	if x != nil {
		return x.RequstingUserId
	}
	return 0
}

func (x *RevokeAdminRequest) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

type RevokeAdminResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeAdminResponse) Reset() {
	*x = RevokeAdminResponse{}
	mi := &file_sso_sso_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeAdminResponse) ProtoMessage() {}

func (x *RevokeAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeAdminResponse.ProtoReflect.Descriptor instead.
func (*RevokeAdminResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{9}
}

func (x *RevokeAdminResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_sso_sso_proto protoreflect.FileDescriptor

const file_sso_sso_proto_rawDesc = "" +
	"\n" +
	"\rsso/sso.proto\x12\x04auth\x1a\x17validate/validate.proto\"\xa5\x01\n" +
	"\x0fRegisterRequest\x12\x1d\n" +
	"\x05email\x18\x01 \x01(\tB\a\xfaB\x04r\x02`\x01R\x05email\x126\n" +
	"\busername\x18\x02 \x01(\tB\x1a\xfaB\x17r\x15\x10\x03\x18\x142\x0f^[a-zA-Z0-9_]+$R\busername\x12;\n" +
	"\bpassword\x18\x03 \x01(\tB\x1f\xfaB\x1cr\x1a\x10\b2\x16^[a-zA-Z0-9!@#$%^&*]+$R\bpassword\"+\n" +
	"\x10RegisterResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"\xd4\x01\n" +
	"\fLoginRequest\x12\x1f\n" +
	"\x05email\x18\x01 \x01(\tB\a\xfaB\x04r\x02`\x01H\x00R\x05email\x128\n" +
	"\busername\x18\x02 \x01(\tB\x1a\xfaB\x17r\x15\x10\x03\x18\x142\x0f^[a-zA-Z0-9_]+$H\x00R\busername\x12;\n" +
	"\bpassword\x18\x03 \x01(\tB\x1f\xfaB\x1cr\x1a\x10\b2\x16^[a-zA-Z0-9!@#$%^&*]+$R\bpassword\x12\x1e\n" +
	"\x06app_id\x18\x04 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x05appIdB\f\n" +
	"\n" +
	"identifier\"%\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"t\n" +
	"\x0eIsAdminRequest\x123\n" +
	"\x11requsting_user_id\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x0frequstingUserId\x12-\n" +
	"\x0etarget_user_id\x18\x02 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\ftargetUserId\",\n" +
	"\x0fIsAdminResponse\x12\x19\n" +
	"\bis_admin\x18\x01 \x01(\bR\aisAdmin\"w\n" +
	"\x11GrantAdminRequest\x123\n" +
	"\x11requsting_user_id\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x0frequstingUserId\x12-\n" +
	"\x0etarget_user_id\x18\x02 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\ftargetUserId\".\n" +
	"\x12GrantAdminResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"x\n" +
	"\x12RevokeAdminRequest\x123\n" +
	"\x11requsting_user_id\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x0frequstingUserId\x12-\n" +
	"\x0etarget_user_id\x18\x02 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\ftargetUserId\"/\n" +
	"\x13RevokeAdminResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xb0\x02\n" +
	"\x04Auth\x129\n" +
	"\bRegister\x12\x15.auth.RegisterRequest\x1a\x16.auth.RegisterResponse\x120\n" +
	"\x05Login\x12\x12.auth.LoginRequest\x1a\x13.auth.LoginResponse\x126\n" +
	"\aIsAdmin\x12\x14.auth.IsAdminRequest\x1a\x15.auth.IsAdminResponse\x12?\n" +
	"\n" +
	"GrantAdmin\x12\x17.auth.GrantAdminRequest\x1a\x18.auth.GrantAdminResponse\x12B\n" +
	"\vRevokeAdmin\x12\x18.auth.RevokeAdminRequest\x1a\x19.auth.RevokeAdminResponseB\x18Z\x16gurebusan.sso.v1;ssov1b\x06proto3"

var (
	file_sso_sso_proto_rawDescOnce sync.Once
	file_sso_sso_proto_rawDescData []byte
)

func file_sso_sso_proto_rawDescGZIP() []byte {
	file_sso_sso_proto_rawDescOnce.Do(func() {
		file_sso_sso_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sso_sso_proto_rawDesc), len(file_sso_sso_proto_rawDesc)))
	})
	return file_sso_sso_proto_rawDescData
}

var file_sso_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_sso_sso_proto_goTypes = []any{
	(*RegisterRequest)(nil),     // 0: auth.RegisterRequest
	(*RegisterResponse)(nil),    // 1: auth.RegisterResponse
	(*LoginRequest)(nil),        // 2: auth.LoginRequest
	(*LoginResponse)(nil),       // 3: auth.LoginResponse
	(*IsAdminRequest)(nil),      // 4: auth.IsAdminRequest
	(*IsAdminResponse)(nil),     // 5: auth.IsAdminResponse
	(*GrantAdminRequest)(nil),   // 6: auth.GrantAdminRequest
	(*GrantAdminResponse)(nil),  // 7: auth.GrantAdminResponse
	(*RevokeAdminRequest)(nil),  // 8: auth.RevokeAdminRequest
	(*RevokeAdminResponse)(nil), // 9: auth.RevokeAdminResponse
}
var file_sso_sso_proto_depIdxs = []int32{
	0, // 0: auth.Auth.Register:input_type -> auth.RegisterRequest
	2, // 1: auth.Auth.Login:input_type -> auth.LoginRequest
	4, // 2: auth.Auth.IsAdmin:input_type -> auth.IsAdminRequest
	6, // 3: auth.Auth.GrantAdmin:input_type -> auth.GrantAdminRequest
	8, // 4: auth.Auth.RevokeAdmin:input_type -> auth.RevokeAdminRequest
	1, // 5: auth.Auth.Register:output_type -> auth.RegisterResponse
	3, // 6: auth.Auth.Login:output_type -> auth.LoginResponse
	5, // 7: auth.Auth.IsAdmin:output_type -> auth.IsAdminResponse
	7, // 8: auth.Auth.GrantAdmin:output_type -> auth.GrantAdminResponse
	9, // 9: auth.Auth.RevokeAdmin:output_type -> auth.RevokeAdminResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sso_sso_proto_init() }
func file_sso_sso_proto_init() {
	if File_sso_sso_proto != nil {
		return
	}
	file_sso_sso_proto_msgTypes[2].OneofWrappers = []any{
		(*LoginRequest_Email)(nil),
		(*LoginRequest_Username)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sso_sso_proto_rawDesc), len(file_sso_sso_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_sso_proto_goTypes,
		DependencyIndexes: file_sso_sso_proto_depIdxs,
		MessageInfos:      file_sso_sso_proto_msgTypes,
	}.Build()
	File_sso_sso_proto = out.File
	file_sso_sso_proto_goTypes = nil
	file_sso_sso_proto_depIdxs = nil
}
