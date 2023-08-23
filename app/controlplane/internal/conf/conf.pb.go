//
// Copyright 2023 The Chainloop Authors.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: conf.proto

package conf

import (
	v1 "github.com/chainloop-dev/chainloop/internal/credentials/api/credentials/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server             *Server                  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data               *Data                    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Auth               *Auth                    `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	Observability      *Bootstrap_Observability `protobuf:"bytes,4,opt,name=observability,proto3" json:"observability,omitempty"`
	CredentialsService *v1.Credentials          `protobuf:"bytes,5,opt,name=credentials_service,json=credentialsService,proto3" json:"credentials_service,omitempty"`
	// CAS Server endpoint
	CasServer *Bootstrap_CASServer `protobuf:"bytes,6,opt,name=cas_server,json=casServer,proto3" json:"cas_server,omitempty"`
	// Plugins directory
	// NOTE: plugins have the form of chainloop-plugin-<name>
	PluginsDir string `protobuf:"bytes,7,opt,name=plugins_dir,json=pluginsDir,proto3" json:"plugins_dir,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Bootstrap) GetObservability() *Bootstrap_Observability {
	if x != nil {
		return x.Observability
	}
	return nil
}

func (x *Bootstrap) GetCredentialsService() *v1.Credentials {
	if x != nil {
		return x.CredentialsService
	}
	return nil
}

func (x *Bootstrap) GetCasServer() *Bootstrap_CASServer {
	if x != nil {
		return x.CasServer
	}
	return nil
}

func (x *Bootstrap) GetPluginsDir() string {
	if x != nil {
		return x.PluginsDir
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc *Server_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	// HTTPMetrics defines the HTTP server that exposes prometheus metrics
	HttpMetrics *Server_HTTP `protobuf:"bytes,3,opt,name=http_metrics,json=httpMetrics,proto3" json:"http_metrics,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Server) GetHttpMetrics() *Server_HTTP {
	if x != nil {
		return x.HttpMetrics
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Authentication creates a JWT that uses this secret for signing
	GeneratedJwsHmacSecret        string     `protobuf:"bytes,2,opt,name=generated_jws_hmac_secret,json=generatedJwsHmacSecret,proto3" json:"generated_jws_hmac_secret,omitempty"`
	AllowList                     []string   `protobuf:"bytes,3,rep,name=allow_list,json=allowList,proto3" json:"allow_list,omitempty"`
	CasRobotAccountPrivateKeyPath string     `protobuf:"bytes,4,opt,name=cas_robot_account_private_key_path,json=casRobotAccountPrivateKeyPath,proto3" json:"cas_robot_account_private_key_path,omitempty"`
	Oidc                          *Auth_OIDC `protobuf:"bytes,6,opt,name=oidc,proto3" json:"oidc,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Auth) GetGeneratedJwsHmacSecret() string {
	if x != nil {
		return x.GeneratedJwsHmacSecret
	}
	return ""
}

func (x *Auth) GetAllowList() []string {
	if x != nil {
		return x.AllowList
	}
	return nil
}

func (x *Auth) GetCasRobotAccountPrivateKeyPath() string {
	if x != nil {
		return x.CasRobotAccountPrivateKeyPath
	}
	return ""
}

func (x *Auth) GetOidc() *Auth_OIDC {
	if x != nil {
		return x.Oidc
	}
	return nil
}

type Bootstrap_Observability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sentry *Bootstrap_Observability_Sentry `protobuf:"bytes,1,opt,name=sentry,proto3" json:"sentry,omitempty"`
}

func (x *Bootstrap_Observability) Reset() {
	*x = Bootstrap_Observability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap_Observability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap_Observability) ProtoMessage() {}

func (x *Bootstrap_Observability) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap_Observability.ProtoReflect.Descriptor instead.
func (*Bootstrap_Observability) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Bootstrap_Observability) GetSentry() *Bootstrap_Observability_Sentry {
	if x != nil {
		return x.Sentry
	}
	return nil
}

type Bootstrap_CASServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CAS server GRPC endpoint
	Grpc *Server_GRPC `protobuf:"bytes,1,opt,name=grpc,proto3" json:"grpc,omitempty"`
	// insecure is used to connect to the CAS server in development
	Insecure bool `protobuf:"varint,2,opt,name=insecure,proto3" json:"insecure,omitempty"`
	// CAS server external http(s) download endpoint
	// In the form of [scheme]://[host]/path i.e https://cas.chainloop.dev/download
	// https://github.com/chainloop-dev/chainloop/blob/126f47b6c0803eac844b8e3e1a21d582f00e4dc6/app/artifact-cas/internal/service/download.go#L34
	DownloadUrl string `protobuf:"bytes,3,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
}

func (x *Bootstrap_CASServer) Reset() {
	*x = Bootstrap_CASServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap_CASServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap_CASServer) ProtoMessage() {}

func (x *Bootstrap_CASServer) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap_CASServer.ProtoReflect.Descriptor instead.
func (*Bootstrap_CASServer) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Bootstrap_CASServer) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Bootstrap_CASServer) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *Bootstrap_CASServer) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

type Bootstrap_Observability_Sentry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dsn         string `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty"`
	Environment string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *Bootstrap_Observability_Sentry) Reset() {
	*x = Bootstrap_Observability_Sentry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap_Observability_Sentry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap_Observability_Sentry) ProtoMessage() {}

func (x *Bootstrap_Observability_Sentry) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap_Observability_Sentry.ProtoReflect.Descriptor instead.
func (*Bootstrap_Observability_Sentry) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Bootstrap_Observability_Sentry) GetDsn() string {
	if x != nil {
		return x.Dsn
	}
	return ""
}

func (x *Bootstrap_Observability_Sentry) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	// In the form of [scheme]://[host] i.e https://instance.chainloop.dev
	// Optional
	ExternalUrl string               `protobuf:"bytes,4,opt,name=external_url,json=externalUrl,proto3" json:"external_url,omitempty"`
	Timeout     *durationpb.Duration `protobuf:"bytes,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetExternalUrl() string {
	if x != nil {
		return x.ExternalUrl
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Auth_OIDC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain            string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	ClientId          string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret      string `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	RedirectUrlScheme string `protobuf:"bytes,4,opt,name=redirect_url_scheme,json=redirectUrlScheme,proto3" json:"redirect_url_scheme,omitempty"`
}

func (x *Auth_OIDC) Reset() {
	*x = Auth_OIDC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth_OIDC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_OIDC) ProtoMessage() {}

func (x *Auth_OIDC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_OIDC.ProtoReflect.Descriptor instead.
func (*Auth_OIDC) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Auth_OIDC) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Auth_OIDC) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Auth_OIDC) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Auth_OIDC) GetRedirectUrlScheme() string {
	if x != nil {
		return x.RedirectUrlScheme
	}
	return ""
}

var File_conf_proto protoreflect.FileDescriptor

var file_conf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc7, 0x04, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x12, 0x1f, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x0d, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0d, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x13, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x52, 0x12, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x61, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x43, 0x41, 0x53, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x09, 0x63, 0x61, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x44, 0x69, 0x72, 0x1a, 0x86, 0x01, 0x0a, 0x0d,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x37, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x3c, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x73, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x76, 0x0a, 0x09, 0x43, 0x41, 0x53, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x80, 0x03, 0x0a,
	0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48,
	0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x20, 0x0a, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x2f, 0x0a, 0x0c, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52,
	0x0b, 0x68, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x8c, 0x01, 0x0a,
	0x04, 0x48, 0x54, 0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x72, 0x0a, 0x04, 0x47,
	0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1b, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22,
	0x6e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0xde, 0x02, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x19, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6a, 0x77, 0x73, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4a, 0x77, 0x73, 0x48, 0x6d, 0x61, 0x63, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x49, 0x0a, 0x22, 0x63, 0x61, 0x73, 0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d,
	0x63, 0x61, 0x73, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a,
	0x04, 0x6f, 0x69, 0x64, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x52, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x1a, 0x90, 0x01,
	0x0a, 0x04, 0x4f, 0x49, 0x44, 0x43, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_conf_proto_rawDescOnce sync.Once
	file_conf_proto_rawDescData = file_conf_proto_rawDesc
)

func file_conf_proto_rawDescGZIP() []byte {
	file_conf_proto_rawDescOnce.Do(func() {
		file_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_proto_rawDescData)
	})
	return file_conf_proto_rawDescData
}

var file_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),                      // 0: Bootstrap
	(*Server)(nil),                         // 1: Server
	(*Data)(nil),                           // 2: Data
	(*Auth)(nil),                           // 3: Auth
	(*Bootstrap_Observability)(nil),        // 4: Bootstrap.Observability
	(*Bootstrap_CASServer)(nil),            // 5: Bootstrap.CASServer
	(*Bootstrap_Observability_Sentry)(nil), // 6: Bootstrap.Observability.Sentry
	(*Server_HTTP)(nil),                    // 7: Server.HTTP
	(*Server_GRPC)(nil),                    // 8: Server.GRPC
	(*Data_Database)(nil),                  // 9: Data.Database
	(*Auth_OIDC)(nil),                      // 10: Auth.OIDC
	(*v1.Credentials)(nil),                 // 11: credentials.v1.Credentials
	(*durationpb.Duration)(nil),            // 12: google.protobuf.Duration
}
var file_conf_proto_depIdxs = []int32{
	1,  // 0: Bootstrap.server:type_name -> Server
	2,  // 1: Bootstrap.data:type_name -> Data
	3,  // 2: Bootstrap.auth:type_name -> Auth
	4,  // 3: Bootstrap.observability:type_name -> Bootstrap.Observability
	11, // 4: Bootstrap.credentials_service:type_name -> credentials.v1.Credentials
	5,  // 5: Bootstrap.cas_server:type_name -> Bootstrap.CASServer
	7,  // 6: Server.http:type_name -> Server.HTTP
	8,  // 7: Server.grpc:type_name -> Server.GRPC
	7,  // 8: Server.http_metrics:type_name -> Server.HTTP
	9,  // 9: Data.database:type_name -> Data.Database
	10, // 10: Auth.oidc:type_name -> Auth.OIDC
	6,  // 11: Bootstrap.Observability.sentry:type_name -> Bootstrap.Observability.Sentry
	8,  // 12: Bootstrap.CASServer.grpc:type_name -> Server.GRPC
	12, // 13: Server.HTTP.timeout:type_name -> google.protobuf.Duration
	12, // 14: Server.GRPC.timeout:type_name -> google.protobuf.Duration
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_conf_proto_init() }
func file_conf_proto_init() {
	if File_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap_Observability); i {
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
		file_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap_CASServer); i {
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
		file_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap_Observability_Sentry); i {
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
		file_conf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
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
		file_conf_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_GRPC); i {
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
		file_conf_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Database); i {
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
		file_conf_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth_OIDC); i {
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
			RawDescriptor: file_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_proto_goTypes,
		DependencyIndexes: file_conf_proto_depIdxs,
		MessageInfos:      file_conf_proto_msgTypes,
	}.Build()
	File_conf_proto = out.File
	file_conf_proto_rawDesc = nil
	file_conf_proto_goTypes = nil
	file_conf_proto_depIdxs = nil
}
