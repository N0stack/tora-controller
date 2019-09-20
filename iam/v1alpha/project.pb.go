// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iam/v1alpha/project.proto

package piam

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProjectMembership int32

const (
	ProjectMembership_PROJECT_MEMBERSHIP_UNSPECIFIED ProjectMembership = 0
	// Owners have all of permissions.
	ProjectMembership_OWNER ProjectMembership = 1
	// Members have only assined permissions by Roles.
	ProjectMembership_MEMBER ProjectMembership = 2
)

var ProjectMembership_name = map[int32]string{
	0: "PROJECT_MEMBERSHIP_UNSPECIFIED",
	1: "OWNER",
	2: "MEMBER",
}

var ProjectMembership_value = map[string]int32{
	"PROJECT_MEMBERSHIP_UNSPECIFIED": 0,
	"OWNER":                          1,
	"MEMBER":                         2,
}

func (x ProjectMembership) String() string {
	return proto.EnumName(ProjectMembership_name, int32(x))
}

func (ProjectMembership) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{0}
}

type Project struct {
	// Name is a unique field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Annotations can store metadata used by the system for control.
	// In particular, implementation-dependent fields that can not be set as protobuf fields are targeted.
	// The control specified by n0stack may delete metadata specified by the user.
	Annotations map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Labels stores user-defined metadata.
	// The n0stack system must not rewrite this value.
	Labels               map[string]string            `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DisplayName          string                       `protobuf:"bytes,9,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Membership           map[string]ProjectMembership `protobuf:"bytes,32,rep,name=membership,proto3" json:"membership,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=n0stack.iam.v1alpha.ProjectMembership"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Project) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Project) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Project) GetMembership() map[string]ProjectMembership {
	if m != nil {
		return m.Membership
	}
	return nil
}

type ListProjectsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProjectsRequest) Reset()         { *m = ListProjectsRequest{} }
func (m *ListProjectsRequest) String() string { return proto.CompactTextString(m) }
func (*ListProjectsRequest) ProtoMessage()    {}
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{1}
}

func (m *ListProjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsRequest.Unmarshal(m, b)
}
func (m *ListProjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsRequest.Marshal(b, m, deterministic)
}
func (m *ListProjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsRequest.Merge(m, src)
}
func (m *ListProjectsRequest) XXX_Size() int {
	return xxx_messageInfo_ListProjectsRequest.Size(m)
}
func (m *ListProjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsRequest proto.InternalMessageInfo

type ListProjectsResponse struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListProjectsResponse) Reset()         { *m = ListProjectsResponse{} }
func (m *ListProjectsResponse) String() string { return proto.CompactTextString(m) }
func (*ListProjectsResponse) ProtoMessage()    {}
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{2}
}

func (m *ListProjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsResponse.Unmarshal(m, b)
}
func (m *ListProjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsResponse.Marshal(b, m, deterministic)
}
func (m *ListProjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsResponse.Merge(m, src)
}
func (m *ListProjectsResponse) XXX_Size() int {
	return xxx_messageInfo_ListProjectsResponse.Size(m)
}
func (m *ListProjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsResponse proto.InternalMessageInfo

func (m *ListProjectsResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type GetProjectRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProjectRequest) Reset()         { *m = GetProjectRequest{} }
func (m *GetProjectRequest) String() string { return proto.CompactTextString(m) }
func (*GetProjectRequest) ProtoMessage()    {}
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{3}
}

func (m *GetProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProjectRequest.Unmarshal(m, b)
}
func (m *GetProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProjectRequest.Marshal(b, m, deterministic)
}
func (m *GetProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProjectRequest.Merge(m, src)
}
func (m *GetProjectRequest) XXX_Size() int {
	return xxx_messageInfo_GetProjectRequest.Size(m)
}
func (m *GetProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProjectRequest proto.InternalMessageInfo

func (m *GetProjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateProjectRequest struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateProjectRequest) Reset()         { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()    {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{4}
}

func (m *CreateProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectRequest.Unmarshal(m, b)
}
func (m *CreateProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectRequest.Marshal(b, m, deterministic)
}
func (m *CreateProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectRequest.Merge(m, src)
}
func (m *CreateProjectRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProjectRequest.Size(m)
}
func (m *CreateProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectRequest proto.InternalMessageInfo

func (m *CreateProjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateProjectRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *CreateProjectRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type UpdateProjectRequest struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateProjectRequest) Reset()         { *m = UpdateProjectRequest{} }
func (m *UpdateProjectRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProjectRequest) ProtoMessage()    {}
func (*UpdateProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{5}
}

func (m *UpdateProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProjectRequest.Unmarshal(m, b)
}
func (m *UpdateProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProjectRequest.Marshal(b, m, deterministic)
}
func (m *UpdateProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProjectRequest.Merge(m, src)
}
func (m *UpdateProjectRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProjectRequest.Size(m)
}
func (m *UpdateProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProjectRequest proto.InternalMessageInfo

func (m *UpdateProjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateProjectRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *UpdateProjectRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type DeleteProjectRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProjectRequest) Reset()         { *m = DeleteProjectRequest{} }
func (m *DeleteProjectRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectRequest) ProtoMessage()    {}
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{6}
}

func (m *DeleteProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectRequest.Unmarshal(m, b)
}
func (m *DeleteProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectRequest.Merge(m, src)
}
func (m *DeleteProjectRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectRequest.Size(m)
}
func (m *DeleteProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectRequest proto.InternalMessageInfo

func (m *DeleteProjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddProjectMembershipRequest struct {
	ProjectName          string            `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	UserName             string            `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Membership           ProjectMembership `protobuf:"varint,3,opt,name=membership,proto3,enum=n0stack.iam.v1alpha.ProjectMembership" json:"membership,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AddProjectMembershipRequest) Reset()         { *m = AddProjectMembershipRequest{} }
func (m *AddProjectMembershipRequest) String() string { return proto.CompactTextString(m) }
func (*AddProjectMembershipRequest) ProtoMessage()    {}
func (*AddProjectMembershipRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{7}
}

func (m *AddProjectMembershipRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProjectMembershipRequest.Unmarshal(m, b)
}
func (m *AddProjectMembershipRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProjectMembershipRequest.Marshal(b, m, deterministic)
}
func (m *AddProjectMembershipRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProjectMembershipRequest.Merge(m, src)
}
func (m *AddProjectMembershipRequest) XXX_Size() int {
	return xxx_messageInfo_AddProjectMembershipRequest.Size(m)
}
func (m *AddProjectMembershipRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProjectMembershipRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddProjectMembershipRequest proto.InternalMessageInfo

func (m *AddProjectMembershipRequest) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *AddProjectMembershipRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AddProjectMembershipRequest) GetMembership() ProjectMembership {
	if m != nil {
		return m.Membership
	}
	return ProjectMembership_PROJECT_MEMBERSHIP_UNSPECIFIED
}

type DeleteProjectMembershipRequest struct {
	ProjectName          string   `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProjectMembershipRequest) Reset()         { *m = DeleteProjectMembershipRequest{} }
func (m *DeleteProjectMembershipRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectMembershipRequest) ProtoMessage()    {}
func (*DeleteProjectMembershipRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a8e632cf9e6b2, []int{8}
}

func (m *DeleteProjectMembershipRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectMembershipRequest.Unmarshal(m, b)
}
func (m *DeleteProjectMembershipRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectMembershipRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProjectMembershipRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectMembershipRequest.Merge(m, src)
}
func (m *DeleteProjectMembershipRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectMembershipRequest.Size(m)
}
func (m *DeleteProjectMembershipRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectMembershipRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectMembershipRequest proto.InternalMessageInfo

func (m *DeleteProjectMembershipRequest) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *DeleteProjectMembershipRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterEnum("n0stack.iam.v1alpha.ProjectMembership", ProjectMembership_name, ProjectMembership_value)
	proto.RegisterType((*Project)(nil), "n0stack.iam.v1alpha.Project")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v1alpha.Project.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v1alpha.Project.LabelsEntry")
	proto.RegisterMapType((map[string]ProjectMembership)(nil), "n0stack.iam.v1alpha.Project.MembershipEntry")
	proto.RegisterType((*ListProjectsRequest)(nil), "n0stack.iam.v1alpha.ListProjectsRequest")
	proto.RegisterType((*ListProjectsResponse)(nil), "n0stack.iam.v1alpha.ListProjectsResponse")
	proto.RegisterType((*GetProjectRequest)(nil), "n0stack.iam.v1alpha.GetProjectRequest")
	proto.RegisterType((*CreateProjectRequest)(nil), "n0stack.iam.v1alpha.CreateProjectRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v1alpha.CreateProjectRequest.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v1alpha.CreateProjectRequest.LabelsEntry")
	proto.RegisterType((*UpdateProjectRequest)(nil), "n0stack.iam.v1alpha.UpdateProjectRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v1alpha.UpdateProjectRequest.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v1alpha.UpdateProjectRequest.LabelsEntry")
	proto.RegisterType((*DeleteProjectRequest)(nil), "n0stack.iam.v1alpha.DeleteProjectRequest")
	proto.RegisterType((*AddProjectMembershipRequest)(nil), "n0stack.iam.v1alpha.AddProjectMembershipRequest")
	proto.RegisterType((*DeleteProjectMembershipRequest)(nil), "n0stack.iam.v1alpha.DeleteProjectMembershipRequest")
}

func init() { proto.RegisterFile("iam/v1alpha/project.proto", fileDescriptor_4f0a8e632cf9e6b2) }

var fileDescriptor_4f0a8e632cf9e6b2 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x6e, 0x3a, 0x45,
	0x14, 0x77, 0xa1, 0x5f, 0x1c, 0xda, 0x4a, 0xa7, 0x58, 0x71, 0xa9, 0x0d, 0x5d, 0x93, 0xda, 0x92,
	0xb2, 0x5b, 0x69, 0x4c, 0x2a, 0x36, 0x6a, 0x3f, 0xb6, 0x5a, 0x53, 0x28, 0xa1, 0x36, 0x26, 0xc6,
	0x04, 0x07, 0x18, 0x61, 0xdb, 0xfd, 0x72, 0x77, 0x69, 0x43, 0x9a, 0x26, 0x7e, 0x5c, 0x78, 0xaf,
	0xcf, 0xe0, 0xa5, 0x89, 0xcf, 0xa2, 0xaf, 0xe0, 0xbd, 0xaf, 0x60, 0x98, 0x1d, 0x60, 0x81, 0xe9,
	0x42, 0x63, 0x13, 0xff, 0x57, 0xcc, 0x9e, 0x39, 0xe7, 0x77, 0x7e, 0xf3, 0xe3, 0x9c, 0x33, 0x03,
	0x6f, 0x69, 0xd8, 0x50, 0xee, 0xde, 0xc3, 0xba, 0xdd, 0xc2, 0x8a, 0xed, 0x58, 0x37, 0xa4, 0xee,
	0xc9, 0xb6, 0x63, 0x79, 0x16, 0x5a, 0x35, 0xf7, 0x5c, 0x0f, 0xd7, 0x6f, 0x65, 0x0d, 0x1b, 0x32,
	0x73, 0x11, 0xd7, 0x9b, 0x96, 0xd5, 0xd4, 0x89, 0x82, 0x6d, 0x4d, 0xc1, 0xa6, 0x69, 0x79, 0xd8,
	0xd3, 0x2c, 0xd3, 0xf5, 0x43, 0xc4, 0x34, 0xdb, 0xa5, 0x5f, 0xb5, 0xf6, 0xb7, 0x0a, 0x31, 0x6c,
	0xaf, 0xc3, 0x36, 0x77, 0xe9, 0x4f, 0x3d, 0xd7, 0x24, 0x66, 0xce, 0xbd, 0xc7, 0xcd, 0x26, 0x71,
	0x14, 0xcb, 0xa6, 0xe1, 0xe3, 0x50, 0xd2, 0x4f, 0x33, 0x30, 0x5f, 0xf6, 0xf9, 0x20, 0x04, 0x33,
	0x26, 0x36, 0x48, 0x4a, 0xc8, 0x08, 0xdb, 0xb1, 0x0a, 0x5d, 0xa3, 0x4b, 0x88, 0x07, 0x82, 0x52,
	0xd1, 0x4c, 0x74, 0x3b, 0x9e, 0xcf, 0xc9, 0x1c, 0xce, 0x32, 0x83, 0x91, 0x8f, 0x06, 0xfe, 0xaa,
	0xe9, 0x39, 0x9d, 0x4a, 0x10, 0x01, 0x7d, 0x02, 0x73, 0x3a, 0xae, 0x11, 0xdd, 0x4d, 0xcd, 0x50,
	0xac, 0xed, 0x50, 0xac, 0x0b, 0xea, 0xea, 0xc3, 0xb0, 0x38, 0xb4, 0x09, 0x8b, 0x0d, 0xcd, 0xb5,
	0x75, 0xdc, 0xa9, 0x52, 0xba, 0x31, 0x4a, 0x37, 0xce, 0x6c, 0xa5, 0x2e, 0xeb, 0x0b, 0x00, 0x83,
	0x18, 0x35, 0xe2, 0xb8, 0x2d, 0xcd, 0x4e, 0x65, 0x68, 0xa2, 0xdd, 0xd0, 0x44, 0xc5, 0xbe, 0xbb,
	0x9f, 0x2c, 0x10, 0x2f, 0x7e, 0x04, 0x89, 0xd1, 0x33, 0xa1, 0x04, 0x44, 0x6f, 0x49, 0x87, 0x49,
	0xd5, 0x5d, 0xa2, 0x24, 0xcc, 0xde, 0x61, 0xbd, 0x4d, 0x52, 0x11, 0x6a, 0xf3, 0x3f, 0x0a, 0x91,
	0x03, 0x41, 0xfc, 0x00, 0xe2, 0x81, 0x73, 0x3c, 0x2b, 0x94, 0xc0, 0xeb, 0x23, 0xcc, 0x38, 0xe1,
	0x87, 0xc1, 0xf0, 0xe5, 0xfc, 0x56, 0xd8, 0x41, 0x07, 0x68, 0x81, 0x34, 0xd2, 0x1b, 0xb0, 0x7a,
	0xa1, 0xb9, 0x1e, 0xf3, 0x71, 0x2b, 0xe4, 0xbb, 0x36, 0x71, 0x3d, 0xa9, 0x0c, 0xc9, 0x61, 0xb3,
	0x6b, 0x5b, 0xa6, 0x4b, 0xd0, 0x01, 0x2c, 0xb0, 0x1a, 0x76, 0x53, 0x02, 0x15, 0x77, 0x3d, 0x2c,
	0x67, 0xa5, 0xef, 0x2d, 0xbd, 0x0b, 0x2b, 0x9f, 0x92, 0x1e, 0x20, 0x4b, 0xc3, 0xab, 0x3b, 0xe9,
	0xcf, 0x08, 0x24, 0x4f, 0x1c, 0x82, 0x3d, 0x32, 0xd9, 0x19, 0x7d, 0xcd, 0x2b, 0xd2, 0x02, 0x97,
	0x12, 0x0f, 0x73, 0x42, 0xc5, 0x16, 0x47, 0x2a, 0xf6, 0xfd, 0xe9, 0x81, 0x39, 0xe5, 0xfb, 0x3f,
	0x56, 0x13, 0x15, 0xf5, 0xda, 0x6e, 0xbc, 0xb8, 0xa8, 0x3c, 0xcc, 0x17, 0x11, 0x95, 0x0b, 0xfc,
	0x8a, 0x89, 0x9a, 0x85, 0xe4, 0x29, 0xd1, 0xc9, 0x34, 0x9a, 0x4a, 0xbf, 0x09, 0x90, 0x3e, 0x6a,
	0x34, 0xc6, 0x7b, 0x91, 0xc5, 0x6c, 0xc2, 0x22, 0x6b, 0x95, 0x6a, 0x20, 0x36, 0xce, 0x6c, 0x74,
	0xb4, 0xa5, 0x21, 0xd6, 0x76, 0x89, 0xe3, 0xef, 0xfb, 0x64, 0x16, 0xba, 0x06, 0xba, 0x79, 0x36,
	0x34, 0xf7, 0xa2, 0xcf, 0x1a, 0x07, 0x81, 0x48, 0xe9, 0x1b, 0xd8, 0x18, 0x3a, 0xd3, 0x8b, 0x33,
	0xcd, 0x96, 0x60, 0x65, 0x0c, 0x1b, 0x49, 0xb0, 0x51, 0xae, 0x5c, 0x7e, 0xae, 0x9e, 0x7c, 0x51,
	0x2d, 0xaa, 0xc5, 0x63, 0xb5, 0x72, 0xf5, 0xd9, 0x79, 0xb9, 0x7a, 0x5d, 0xba, 0x2a, 0xab, 0x27,
	0xe7, 0x67, 0xe7, 0xea, 0x69, 0xe2, 0x35, 0x14, 0x83, 0xd9, 0xcb, 0x2f, 0x4b, 0x6a, 0x25, 0x21,
	0x20, 0x80, 0x39, 0xdf, 0x2d, 0x11, 0xc9, 0xff, 0x33, 0x0f, 0xcb, 0x0c, 0xf0, 0x8a, 0x38, 0x77,
	0x5a, 0x9d, 0xa0, 0x1f, 0x04, 0x58, 0x0c, 0x8e, 0x2f, 0xc4, 0xbf, 0x6a, 0x38, 0x83, 0x4f, 0xdc,
	0x99, 0xc2, 0xd3, 0x9f, 0x85, 0xd2, 0xfa, 0x8f, 0x7f, 0xfd, 0xfd, 0x6b, 0x64, 0x0d, 0x25, 0xe9,
	0x5d, 0x3d, 0x72, 0xc5, 0xa3, 0x7b, 0x80, 0xc1, 0xbc, 0x43, 0xfc, 0xbf, 0x62, 0x6c, 0x20, 0x8a,
	0xa1, 0xd3, 0x54, 0x7a, 0x87, 0x66, 0x7c, 0x1b, 0xa5, 0x79, 0x19, 0x95, 0x87, 0xae, 0xdc, 0x8f,
	0xe8, 0x7b, 0x01, 0x96, 0x86, 0x46, 0x12, 0xda, 0x99, 0x7a, 0x6c, 0x4d, 0xc8, 0xbf, 0x45, 0xf3,
	0x67, 0xa4, 0xb0, 0xfc, 0x05, 0x21, 0x4b, 0x29, 0x0c, 0x35, 0xf0, 0x13, 0x14, 0x78, 0x4d, 0x3e,
	0x1d, 0x05, 0x71, 0x12, 0x85, 0x7b, 0x58, 0x1a, 0xaa, 0xe3, 0x27, 0x18, 0xf0, 0xfa, 0x57, 0x5c,
	0x93, 0xfd, 0x57, 0x96, 0xdc, 0x7b, 0x65, 0xc9, 0x6a, 0xf7, 0x95, 0xd5, 0x93, 0x3f, 0x1b, 0x2a,
	0xff, 0xef, 0x02, 0x24, 0x79, 0x8d, 0x8e, 0xf6, 0xb8, 0x04, 0x42, 0x66, 0xc2, 0x04, 0x25, 0x54,
	0xca, 0xe6, 0x63, 0xa9, 0xc0, 0x67, 0x13, 0xec, 0xd1, 0x47, 0x65, 0xd0, 0xe1, 0xca, 0x43, 0xbf,
	0x35, 0xa9, 0x50, 0x7f, 0x08, 0xf0, 0xe6, 0x13, 0x1d, 0x8f, 0xf6, 0x27, 0x6b, 0xf6, 0x5c, 0xd6,
	0xc7, 0x94, 0xf5, 0x61, 0xf6, 0x3f, 0xb0, 0x3e, 0xfe, 0x59, 0xf8, 0xe5, 0xa8, 0x86, 0x0e, 0x60,
	0x9e, 0x25, 0x92, 0x72, 0xfd, 0x25, 0x92, 0x5a, 0x9e, 0x67, 0xbb, 0x05, 0x45, 0x69, 0x6a, 0x5e,
	0xab, 0x5d, 0x93, 0xeb, 0x96, 0xa1, 0xb0, 0xbd, 0xde, 0x6f, 0x36, 0x22, 0x44, 0xf2, 0x09, 0x6c,
	0xdb, 0xba, 0x56, 0xa7, 0xb7, 0x86, 0x72, 0xe3, 0x5a, 0x66, 0x61, 0xcc, 0xf2, 0x55, 0xa6, 0x1b,
	0x20, 0xe3, 0x7a, 0x1f, 0x20, 0xf0, 0x94, 0xff, 0xd0, 0xd6, 0xb0, 0x51, 0x9b, 0xa3, 0x15, 0xb2,
	0xff, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0xae, 0xd4, 0xca, 0xe5, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectServiceClient interface {
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	// memberじゃない場合 notfound
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error)
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*Project, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddProjectMembership(ctx context.Context, in *AddProjectMembershipRequest, opts ...grpc.CallOption) (*Project, error)
	DeleteProjectMembership(ctx context.Context, in *DeleteProjectMembershipRequest, opts ...grpc.CallOption) (*Project, error)
}

type projectServiceClient struct {
	cc *grpc.ClientConn
}

func NewProjectServiceClient(cc *grpc.ClientConn) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v1alpha.ProjectService/ListProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v1alpha.ProjectService/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v1alpha.ProjectService/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v1alpha.ProjectService/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v1alpha.ProjectService/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) AddProjectMembership(ctx context.Context, in *AddProjectMembershipRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v1alpha.ProjectService/AddProjectMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProjectMembership(ctx context.Context, in *DeleteProjectMembershipRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v1alpha.ProjectService/DeleteProjectMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
type ProjectServiceServer interface {
	ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	// memberじゃない場合 notfound
	GetProject(context.Context, *GetProjectRequest) (*Project, error)
	CreateProject(context.Context, *CreateProjectRequest) (*Project, error)
	UpdateProject(context.Context, *UpdateProjectRequest) (*Project, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*empty.Empty, error)
	AddProjectMembership(context.Context, *AddProjectMembershipRequest) (*Project, error)
	DeleteProjectMembership(context.Context, *DeleteProjectMembershipRequest) (*Project, error)
}

// UnimplementedProjectServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (*UnimplementedProjectServiceServer) ListProjects(ctx context.Context, req *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (*UnimplementedProjectServiceServer) GetProject(ctx context.Context, req *GetProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (*UnimplementedProjectServiceServer) CreateProject(ctx context.Context, req *CreateProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (*UnimplementedProjectServiceServer) UpdateProject(ctx context.Context, req *UpdateProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (*UnimplementedProjectServiceServer) DeleteProject(ctx context.Context, req *DeleteProjectRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (*UnimplementedProjectServiceServer) AddProjectMembership(ctx context.Context, req *AddProjectMembershipRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProjectMembership not implemented")
}
func (*UnimplementedProjectServiceServer) DeleteProjectMembership(ctx context.Context, req *DeleteProjectMembershipRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProjectMembership not implemented")
}

func RegisterProjectServiceServer(s *grpc.Server, srv ProjectServiceServer) {
	s.RegisterService(&_ProjectService_serviceDesc, srv)
}

func _ProjectService_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v1alpha.ProjectService/ListProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ListProjects(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v1alpha.ProjectService/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v1alpha.ProjectService/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v1alpha.ProjectService/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProject(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v1alpha.ProjectService/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_AddProjectMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProjectMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).AddProjectMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v1alpha.ProjectService/AddProjectMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).AddProjectMembership(ctx, req.(*AddProjectMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProjectMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProjectMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v1alpha.ProjectService/DeleteProjectMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProjectMembership(ctx, req.(*DeleteProjectMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.iam.v1alpha.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProjects",
			Handler:    _ProjectService_ListProjects_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _ProjectService_GetProject_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _ProjectService_CreateProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _ProjectService_UpdateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _ProjectService_DeleteProject_Handler,
		},
		{
			MethodName: "AddProjectMembership",
			Handler:    _ProjectService_AddProjectMembership_Handler,
		},
		{
			MethodName: "DeleteProjectMembership",
			Handler:    _ProjectService_DeleteProjectMembership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam/v1alpha/project.proto",
}