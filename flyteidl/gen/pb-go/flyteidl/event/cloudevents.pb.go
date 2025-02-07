// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/event/cloudevents.proto

package event

import (
	fmt "fmt"
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// This is the cloud event parallel to the raw WorkflowExecutionEvent message. It's filled in with additional
// information that downstream consumers may find useful.
type CloudEventWorkflowExecution struct {
	RawEvent        *WorkflowExecutionEvent `protobuf:"bytes,1,opt,name=raw_event,json=rawEvent,proto3" json:"raw_event,omitempty"`
	OutputInterface *core.TypedInterface    `protobuf:"bytes,2,opt,name=output_interface,json=outputInterface,proto3" json:"output_interface,omitempty"`
	// The following are ExecutionMetadata fields
	// We can't have the ExecutionMetadata object directly because of import cycle
	ArtifactIds        []*core.ArtifactID                `protobuf:"bytes,3,rep,name=artifact_ids,json=artifactIds,proto3" json:"artifact_ids,omitempty"`
	ReferenceExecution *core.WorkflowExecutionIdentifier `protobuf:"bytes,4,opt,name=reference_execution,json=referenceExecution,proto3" json:"reference_execution,omitempty"`
	Principal          string                            `protobuf:"bytes,5,opt,name=principal,proto3" json:"principal,omitempty"`
	// The ID of the LP that generated the execution that generated the Artifact.
	// Here for provenance information.
	// Launch plan IDs are easier to get than workflow IDs so we'll use these for now.
	LaunchPlanId         *core.Identifier `protobuf:"bytes,6,opt,name=launch_plan_id,json=launchPlanId,proto3" json:"launch_plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CloudEventWorkflowExecution) Reset()         { *m = CloudEventWorkflowExecution{} }
func (m *CloudEventWorkflowExecution) String() string { return proto.CompactTextString(m) }
func (*CloudEventWorkflowExecution) ProtoMessage()    {}
func (*CloudEventWorkflowExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8af3ecc827e5d5e, []int{0}
}

func (m *CloudEventWorkflowExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudEventWorkflowExecution.Unmarshal(m, b)
}
func (m *CloudEventWorkflowExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudEventWorkflowExecution.Marshal(b, m, deterministic)
}
func (m *CloudEventWorkflowExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudEventWorkflowExecution.Merge(m, src)
}
func (m *CloudEventWorkflowExecution) XXX_Size() int {
	return xxx_messageInfo_CloudEventWorkflowExecution.Size(m)
}
func (m *CloudEventWorkflowExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudEventWorkflowExecution.DiscardUnknown(m)
}

var xxx_messageInfo_CloudEventWorkflowExecution proto.InternalMessageInfo

func (m *CloudEventWorkflowExecution) GetRawEvent() *WorkflowExecutionEvent {
	if m != nil {
		return m.RawEvent
	}
	return nil
}

func (m *CloudEventWorkflowExecution) GetOutputInterface() *core.TypedInterface {
	if m != nil {
		return m.OutputInterface
	}
	return nil
}

func (m *CloudEventWorkflowExecution) GetArtifactIds() []*core.ArtifactID {
	if m != nil {
		return m.ArtifactIds
	}
	return nil
}

func (m *CloudEventWorkflowExecution) GetReferenceExecution() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ReferenceExecution
	}
	return nil
}

func (m *CloudEventWorkflowExecution) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func (m *CloudEventWorkflowExecution) GetLaunchPlanId() *core.Identifier {
	if m != nil {
		return m.LaunchPlanId
	}
	return nil
}

type CloudEventNodeExecution struct {
	RawEvent *NodeExecutionEvent `protobuf:"bytes,1,opt,name=raw_event,json=rawEvent,proto3" json:"raw_event,omitempty"`
	// The relevant task execution if applicable
	TaskExecId *core.TaskExecutionIdentifier `protobuf:"bytes,2,opt,name=task_exec_id,json=taskExecId,proto3" json:"task_exec_id,omitempty"`
	// The typed interface for the task that produced the event.
	OutputInterface *core.TypedInterface `protobuf:"bytes,3,opt,name=output_interface,json=outputInterface,proto3" json:"output_interface,omitempty"`
	// The following are ExecutionMetadata fields
	// We can't have the ExecutionMetadata object directly because of import cycle
	ArtifactIds []*core.ArtifactID `protobuf:"bytes,4,rep,name=artifact_ids,json=artifactIds,proto3" json:"artifact_ids,omitempty"`
	Principal   string             `protobuf:"bytes,5,opt,name=principal,proto3" json:"principal,omitempty"`
	// The ID of the LP that generated the execution that generated the Artifact.
	// Here for provenance information.
	// Launch plan IDs are easier to get than workflow IDs so we'll use these for now.
	LaunchPlanId         *core.Identifier `protobuf:"bytes,6,opt,name=launch_plan_id,json=launchPlanId,proto3" json:"launch_plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CloudEventNodeExecution) Reset()         { *m = CloudEventNodeExecution{} }
func (m *CloudEventNodeExecution) String() string { return proto.CompactTextString(m) }
func (*CloudEventNodeExecution) ProtoMessage()    {}
func (*CloudEventNodeExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8af3ecc827e5d5e, []int{1}
}

func (m *CloudEventNodeExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudEventNodeExecution.Unmarshal(m, b)
}
func (m *CloudEventNodeExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudEventNodeExecution.Marshal(b, m, deterministic)
}
func (m *CloudEventNodeExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudEventNodeExecution.Merge(m, src)
}
func (m *CloudEventNodeExecution) XXX_Size() int {
	return xxx_messageInfo_CloudEventNodeExecution.Size(m)
}
func (m *CloudEventNodeExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudEventNodeExecution.DiscardUnknown(m)
}

var xxx_messageInfo_CloudEventNodeExecution proto.InternalMessageInfo

func (m *CloudEventNodeExecution) GetRawEvent() *NodeExecutionEvent {
	if m != nil {
		return m.RawEvent
	}
	return nil
}

func (m *CloudEventNodeExecution) GetTaskExecId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.TaskExecId
	}
	return nil
}

func (m *CloudEventNodeExecution) GetOutputInterface() *core.TypedInterface {
	if m != nil {
		return m.OutputInterface
	}
	return nil
}

func (m *CloudEventNodeExecution) GetArtifactIds() []*core.ArtifactID {
	if m != nil {
		return m.ArtifactIds
	}
	return nil
}

func (m *CloudEventNodeExecution) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func (m *CloudEventNodeExecution) GetLaunchPlanId() *core.Identifier {
	if m != nil {
		return m.LaunchPlanId
	}
	return nil
}

type CloudEventTaskExecution struct {
	RawEvent             *TaskExecutionEvent `protobuf:"bytes,1,opt,name=raw_event,json=rawEvent,proto3" json:"raw_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CloudEventTaskExecution) Reset()         { *m = CloudEventTaskExecution{} }
func (m *CloudEventTaskExecution) String() string { return proto.CompactTextString(m) }
func (*CloudEventTaskExecution) ProtoMessage()    {}
func (*CloudEventTaskExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8af3ecc827e5d5e, []int{2}
}

func (m *CloudEventTaskExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudEventTaskExecution.Unmarshal(m, b)
}
func (m *CloudEventTaskExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudEventTaskExecution.Marshal(b, m, deterministic)
}
func (m *CloudEventTaskExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudEventTaskExecution.Merge(m, src)
}
func (m *CloudEventTaskExecution) XXX_Size() int {
	return xxx_messageInfo_CloudEventTaskExecution.Size(m)
}
func (m *CloudEventTaskExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudEventTaskExecution.DiscardUnknown(m)
}

var xxx_messageInfo_CloudEventTaskExecution proto.InternalMessageInfo

func (m *CloudEventTaskExecution) GetRawEvent() *TaskExecutionEvent {
	if m != nil {
		return m.RawEvent
	}
	return nil
}

// This event is to be sent by Admin after it creates an execution.
type CloudEventExecutionStart struct {
	// The execution created.
	ExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// The launch plan used.
	LaunchPlanId *core.Identifier `protobuf:"bytes,2,opt,name=launch_plan_id,json=launchPlanId,proto3" json:"launch_plan_id,omitempty"`
	WorkflowId   *core.Identifier `protobuf:"bytes,3,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// Artifact inputs to the workflow execution for which we have the full Artifact ID. These are likely the result of artifact queries that are run.
	ArtifactIds []*core.ArtifactID `protobuf:"bytes,4,rep,name=artifact_ids,json=artifactIds,proto3" json:"artifact_ids,omitempty"`
	// Artifact inputs to the workflow execution for which we only have the tracking bit that's installed into the Literal's metadata by the Artifact service.
	ArtifactTrackers     []string `protobuf:"bytes,5,rep,name=artifact_trackers,json=artifactTrackers,proto3" json:"artifact_trackers,omitempty"`
	Principal            string   `protobuf:"bytes,6,opt,name=principal,proto3" json:"principal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudEventExecutionStart) Reset()         { *m = CloudEventExecutionStart{} }
func (m *CloudEventExecutionStart) String() string { return proto.CompactTextString(m) }
func (*CloudEventExecutionStart) ProtoMessage()    {}
func (*CloudEventExecutionStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8af3ecc827e5d5e, []int{3}
}

func (m *CloudEventExecutionStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudEventExecutionStart.Unmarshal(m, b)
}
func (m *CloudEventExecutionStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudEventExecutionStart.Marshal(b, m, deterministic)
}
func (m *CloudEventExecutionStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudEventExecutionStart.Merge(m, src)
}
func (m *CloudEventExecutionStart) XXX_Size() int {
	return xxx_messageInfo_CloudEventExecutionStart.Size(m)
}
func (m *CloudEventExecutionStart) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudEventExecutionStart.DiscardUnknown(m)
}

var xxx_messageInfo_CloudEventExecutionStart proto.InternalMessageInfo

func (m *CloudEventExecutionStart) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

func (m *CloudEventExecutionStart) GetLaunchPlanId() *core.Identifier {
	if m != nil {
		return m.LaunchPlanId
	}
	return nil
}

func (m *CloudEventExecutionStart) GetWorkflowId() *core.Identifier {
	if m != nil {
		return m.WorkflowId
	}
	return nil
}

func (m *CloudEventExecutionStart) GetArtifactIds() []*core.ArtifactID {
	if m != nil {
		return m.ArtifactIds
	}
	return nil
}

func (m *CloudEventExecutionStart) GetArtifactTrackers() []string {
	if m != nil {
		return m.ArtifactTrackers
	}
	return nil
}

func (m *CloudEventExecutionStart) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func init() {
	proto.RegisterType((*CloudEventWorkflowExecution)(nil), "flyteidl.event.CloudEventWorkflowExecution")
	proto.RegisterType((*CloudEventNodeExecution)(nil), "flyteidl.event.CloudEventNodeExecution")
	proto.RegisterType((*CloudEventTaskExecution)(nil), "flyteidl.event.CloudEventTaskExecution")
	proto.RegisterType((*CloudEventExecutionStart)(nil), "flyteidl.event.CloudEventExecutionStart")
}

func init() { proto.RegisterFile("flyteidl/event/cloudevents.proto", fileDescriptor_f8af3ecc827e5d5e) }

var fileDescriptor_f8af3ecc827e5d5e = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xa6, 0xcd, 0x6e, 0xb1, 0xd3, 0xb2, 0xae, 0xf1, 0xc2, 0x58, 0x77, 0x35, 0xf4, 0x42, 0x8a,
	0x62, 0x02, 0xeb, 0x9d, 0x3f, 0x2c, 0xba, 0x2e, 0x98, 0x0b, 0x45, 0xe2, 0x82, 0xb0, 0x5e, 0x84,
	0x69, 0xe6, 0x24, 0x3b, 0x74, 0x3a, 0x13, 0x26, 0x13, 0xeb, 0x3e, 0x83, 0xef, 0xe1, 0xeb, 0xf9,
	0x0a, 0x92, 0xc9, 0x5f, 0x93, 0x16, 0xb4, 0xe8, 0xde, 0x84, 0xc9, 0x39, 0xdf, 0xf7, 0x9d, 0x9f,
	0x6f, 0x18, 0x64, 0x47, 0xec, 0x5a, 0x01, 0x25, 0xcc, 0x85, 0x6f, 0xc0, 0x95, 0x1b, 0x32, 0x91,
	0x11, 0x7d, 0x4c, 0x9d, 0x44, 0x0a, 0x25, 0xcc, 0x83, 0x0a, 0xe1, 0xe8, 0xf0, 0x64, 0xd2, 0x61,
	0xe8, 0x6f, 0x81, 0x9d, 0x1c, 0xd5, 0xb9, 0x50, 0x48, 0x70, 0x19, 0x55, 0x20, 0x31, 0x2b, 0x95,
	0x26, 0xc7, 0xed, 0x2c, 0xe5, 0x0a, 0x64, 0x84, 0x43, 0x28, 0xd3, 0x8f, 0xda, 0x69, 0x2c, 0x15,
	0x8d, 0x70, 0xa8, 0x02, 0x4a, 0x4a, 0xc0, 0xc3, 0x0e, 0x9f, 0x00, 0x57, 0x34, 0xa2, 0x20, 0x2b,
	0x81, 0x58, 0x88, 0x98, 0x81, 0xab, 0xff, 0xe6, 0x59, 0xe4, 0x2a, 0xba, 0x84, 0x54, 0xe1, 0x65,
	0x52, 0x00, 0xa6, 0x3f, 0x0d, 0xf4, 0xe0, 0x2c, 0x1f, 0xf0, 0x3c, 0xef, 0xf9, 0x8b, 0x90, 0x8b,
	0x88, 0x89, 0xd5, 0xf9, 0x77, 0x08, 0x33, 0x45, 0x05, 0x37, 0xcf, 0xd0, 0x50, 0xe2, 0x55, 0xa0,
	0x27, 0xb2, 0x7a, 0x76, 0x6f, 0x36, 0x3a, 0x79, 0xec, 0xb4, 0xc7, 0x77, 0x36, 0x58, 0x5a, 0xcb,
	0xbf, 0x25, 0xf1, 0x4a, 0x9f, 0xcc, 0xf7, 0xe8, 0x50, 0x64, 0x2a, 0xc9, 0x54, 0x50, 0x0f, 0x68,
	0xf5, 0xb5, 0xd6, 0x71, 0xa3, 0x95, 0x0f, 0xe0, 0x5c, 0x5c, 0x27, 0x40, 0xbc, 0x0a, 0xe4, 0xdf,
	0x2e, 0x68, 0x75, 0xc0, 0x7c, 0x85, 0xc6, 0x6b, 0x4b, 0x48, 0x2d, 0xc3, 0x36, 0x66, 0xa3, 0x93,
	0xfb, 0x1d, 0x95, 0x37, 0x25, 0xc4, 0x7b, 0xe7, 0x8f, 0x2a, 0xb8, 0x47, 0x52, 0xf3, 0x2b, 0xba,
	0x2b, 0x21, 0x02, 0x09, 0x3c, 0x84, 0x00, 0xaa, 0x6e, 0xad, 0x3d, 0xdd, 0xca, 0x93, 0x8e, 0xc8,
	0xc6, 0x54, 0x5e, 0xbd, 0x5c, 0xdf, 0xac, 0x65, 0x9a, 0x4d, 0x1d, 0xa1, 0x61, 0x22, 0x29, 0x0f,
	0x69, 0x82, 0x99, 0xb5, 0x6f, 0xf7, 0x66, 0x43, 0xbf, 0x09, 0x98, 0xa7, 0xe8, 0x80, 0xe1, 0x8c,
	0x87, 0x57, 0x41, 0xc2, 0x30, 0x0f, 0x28, 0xb1, 0x06, 0xba, 0x6a, 0xb7, 0xf5, 0xb5, 0x22, 0xe3,
	0x82, 0xf0, 0x89, 0x61, 0xee, 0x91, 0xe9, 0x0f, 0x03, 0xdd, 0x6b, 0x8c, 0xfa, 0x28, 0xc8, 0x5a,
	0xe9, 0xd3, 0x4d, 0x93, 0xa6, 0x5d, 0x93, 0x5a, 0x8c, 0x4d, 0x83, 0xc6, 0x0a, 0xa7, 0x0b, 0xbd,
	0x93, 0xbc, 0xb7, 0x7e, 0xd7, 0xe8, 0xc2, 0x1c, 0x9c, 0x2e, 0xb6, 0x6d, 0x03, 0xa9, 0x32, 0xe1,
	0x91, 0xad, 0x56, 0x1b, 0xff, 0xc5, 0xea, 0xbd, 0x9d, 0xac, 0xbe, 0x61, 0x37, 0x2e, 0xd7, 0xcd,
	0x68, 0xed, 0xe5, 0xaf, 0xcc, 0x68, 0x31, 0x3a, 0x66, 0x4c, 0x7f, 0xf5, 0x91, 0xd5, 0x88, 0xd7,
	0xb0, 0xcf, 0x0a, 0x4b, 0x65, 0x7e, 0x40, 0xe3, 0xfa, 0xe2, 0xe6, 0x7d, 0xf7, 0x76, 0xbe, 0xbb,
	0x23, 0x68, 0x82, 0x5b, 0x16, 0xd1, 0xdf, 0x69, 0x11, 0xe6, 0x0b, 0x34, 0x5a, 0x95, 0xc5, 0x72,
	0xb6, 0xf1, 0x27, 0x36, 0xaa, 0xd0, 0x1e, 0xf9, 0x47, 0x87, 0x9f, 0xa2, 0x3b, 0x35, 0x5b, 0x49,
	0x1c, 0x2e, 0x40, 0xa6, 0xd6, 0xbe, 0x6d, 0xcc, 0x86, 0xfe, 0x61, 0x95, 0xb8, 0x28, 0xe3, 0xed,
	0xeb, 0x30, 0xe8, 0x5c, 0x87, 0xb7, 0xaf, 0x2f, 0x5f, 0xc6, 0x54, 0x5d, 0x65, 0x73, 0x27, 0x14,
	0x4b, 0x57, 0x97, 0x17, 0x32, 0x2e, 0x0e, 0x6e, 0xfd, 0xc2, 0xc6, 0xc0, 0xdd, 0x64, 0xfe, 0x2c,
	0x16, 0x6e, 0xfb, 0xb9, 0x9f, 0x0f, 0xf4, 0x53, 0xfa, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xef, 0x43, 0x56, 0x39, 0x06, 0x00, 0x00,
}
