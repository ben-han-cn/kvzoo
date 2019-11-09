// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/operations.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Metadata used across all long running operations returned by AutoML API.
type OperationMetadata struct {
	// Ouptut only. Details of specific operation. Even if this field is empty,
	// the presence allows to distinguish different types of operations.
	//
	// Types that are valid to be assigned to Details:
	//	*OperationMetadata_DeleteDetails
	//	*OperationMetadata_DeployModelDetails
	//	*OperationMetadata_UndeployModelDetails
	//	*OperationMetadata_CreateModelDetails
	//	*OperationMetadata_CreateDatasetDetails
	//	*OperationMetadata_ImportDataDetails
	//	*OperationMetadata_BatchPredictDetails
	//	*OperationMetadata_ExportDataDetails
	//	*OperationMetadata_ExportModelDetails
	Details isOperationMetadata_Details `protobuf_oneof:"details"`
	// Output only. Progress of operation. Range: [0, 100].
	// Not used currently.
	ProgressPercent int32 `protobuf:"varint,13,opt,name=progress_percent,json=progressPercent,proto3" json:"progress_percent,omitempty"`
	// Output only. Partial failures encountered.
	// E.g. single files that couldn't be read.
	// This field should never exceed 20 entries.
	// Status details field will contain standard GCP error details.
	PartialFailures []*status.Status `protobuf:"bytes,2,rep,name=partial_failures,json=partialFailures,proto3" json:"partial_failures,omitempty"`
	// Output only. Time when the operation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Time when the operation was updated for the last time.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OperationMetadata) Reset()         { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()    {}
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{0}
}

func (m *OperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationMetadata.Unmarshal(m, b)
}
func (m *OperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationMetadata.Marshal(b, m, deterministic)
}
func (m *OperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationMetadata.Merge(m, src)
}
func (m *OperationMetadata) XXX_Size() int {
	return xxx_messageInfo_OperationMetadata.Size(m)
}
func (m *OperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_OperationMetadata proto.InternalMessageInfo

type isOperationMetadata_Details interface {
	isOperationMetadata_Details()
}

type OperationMetadata_DeleteDetails struct {
	DeleteDetails *DeleteOperationMetadata `protobuf:"bytes,8,opt,name=delete_details,json=deleteDetails,proto3,oneof"`
}

type OperationMetadata_DeployModelDetails struct {
	DeployModelDetails *DeployModelOperationMetadata `protobuf:"bytes,24,opt,name=deploy_model_details,json=deployModelDetails,proto3,oneof"`
}

type OperationMetadata_UndeployModelDetails struct {
	UndeployModelDetails *UndeployModelOperationMetadata `protobuf:"bytes,25,opt,name=undeploy_model_details,json=undeployModelDetails,proto3,oneof"`
}

type OperationMetadata_CreateModelDetails struct {
	CreateModelDetails *CreateModelOperationMetadata `protobuf:"bytes,10,opt,name=create_model_details,json=createModelDetails,proto3,oneof"`
}

type OperationMetadata_CreateDatasetDetails struct {
	CreateDatasetDetails *CreateDatasetOperationMetadata `protobuf:"bytes,30,opt,name=create_dataset_details,json=createDatasetDetails,proto3,oneof"`
}

type OperationMetadata_ImportDataDetails struct {
	ImportDataDetails *ImportDataOperationMetadata `protobuf:"bytes,15,opt,name=import_data_details,json=importDataDetails,proto3,oneof"`
}

type OperationMetadata_BatchPredictDetails struct {
	BatchPredictDetails *BatchPredictOperationMetadata `protobuf:"bytes,16,opt,name=batch_predict_details,json=batchPredictDetails,proto3,oneof"`
}

type OperationMetadata_ExportDataDetails struct {
	ExportDataDetails *ExportDataOperationMetadata `protobuf:"bytes,21,opt,name=export_data_details,json=exportDataDetails,proto3,oneof"`
}

type OperationMetadata_ExportModelDetails struct {
	ExportModelDetails *ExportModelOperationMetadata `protobuf:"bytes,22,opt,name=export_model_details,json=exportModelDetails,proto3,oneof"`
}

func (*OperationMetadata_DeleteDetails) isOperationMetadata_Details() {}

func (*OperationMetadata_DeployModelDetails) isOperationMetadata_Details() {}

func (*OperationMetadata_UndeployModelDetails) isOperationMetadata_Details() {}

func (*OperationMetadata_CreateModelDetails) isOperationMetadata_Details() {}

func (*OperationMetadata_CreateDatasetDetails) isOperationMetadata_Details() {}

func (*OperationMetadata_ImportDataDetails) isOperationMetadata_Details() {}

func (*OperationMetadata_BatchPredictDetails) isOperationMetadata_Details() {}

func (*OperationMetadata_ExportDataDetails) isOperationMetadata_Details() {}

func (*OperationMetadata_ExportModelDetails) isOperationMetadata_Details() {}

func (m *OperationMetadata) GetDetails() isOperationMetadata_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *OperationMetadata) GetDeleteDetails() *DeleteOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_DeleteDetails); ok {
		return x.DeleteDetails
	}
	return nil
}

func (m *OperationMetadata) GetDeployModelDetails() *DeployModelOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_DeployModelDetails); ok {
		return x.DeployModelDetails
	}
	return nil
}

func (m *OperationMetadata) GetUndeployModelDetails() *UndeployModelOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_UndeployModelDetails); ok {
		return x.UndeployModelDetails
	}
	return nil
}

func (m *OperationMetadata) GetCreateModelDetails() *CreateModelOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_CreateModelDetails); ok {
		return x.CreateModelDetails
	}
	return nil
}

func (m *OperationMetadata) GetCreateDatasetDetails() *CreateDatasetOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_CreateDatasetDetails); ok {
		return x.CreateDatasetDetails
	}
	return nil
}

func (m *OperationMetadata) GetImportDataDetails() *ImportDataOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_ImportDataDetails); ok {
		return x.ImportDataDetails
	}
	return nil
}

func (m *OperationMetadata) GetBatchPredictDetails() *BatchPredictOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_BatchPredictDetails); ok {
		return x.BatchPredictDetails
	}
	return nil
}

func (m *OperationMetadata) GetExportDataDetails() *ExportDataOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_ExportDataDetails); ok {
		return x.ExportDataDetails
	}
	return nil
}

func (m *OperationMetadata) GetExportModelDetails() *ExportModelOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_ExportModelDetails); ok {
		return x.ExportModelDetails
	}
	return nil
}

func (m *OperationMetadata) GetProgressPercent() int32 {
	if m != nil {
		return m.ProgressPercent
	}
	return 0
}

func (m *OperationMetadata) GetPartialFailures() []*status.Status {
	if m != nil {
		return m.PartialFailures
	}
	return nil
}

func (m *OperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *OperationMetadata) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OperationMetadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OperationMetadata_DeleteDetails)(nil),
		(*OperationMetadata_DeployModelDetails)(nil),
		(*OperationMetadata_UndeployModelDetails)(nil),
		(*OperationMetadata_CreateModelDetails)(nil),
		(*OperationMetadata_CreateDatasetDetails)(nil),
		(*OperationMetadata_ImportDataDetails)(nil),
		(*OperationMetadata_BatchPredictDetails)(nil),
		(*OperationMetadata_ExportDataDetails)(nil),
		(*OperationMetadata_ExportModelDetails)(nil),
	}
}

// Details of operations that perform deletes of any entities.
type DeleteOperationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOperationMetadata) Reset()         { *m = DeleteOperationMetadata{} }
func (m *DeleteOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*DeleteOperationMetadata) ProtoMessage()    {}
func (*DeleteOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{1}
}

func (m *DeleteOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOperationMetadata.Unmarshal(m, b)
}
func (m *DeleteOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOperationMetadata.Marshal(b, m, deterministic)
}
func (m *DeleteOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOperationMetadata.Merge(m, src)
}
func (m *DeleteOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_DeleteOperationMetadata.Size(m)
}
func (m *DeleteOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOperationMetadata proto.InternalMessageInfo

// Details of DeployModel operation.
type DeployModelOperationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployModelOperationMetadata) Reset()         { *m = DeployModelOperationMetadata{} }
func (m *DeployModelOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*DeployModelOperationMetadata) ProtoMessage()    {}
func (*DeployModelOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{2}
}

func (m *DeployModelOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployModelOperationMetadata.Unmarshal(m, b)
}
func (m *DeployModelOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployModelOperationMetadata.Marshal(b, m, deterministic)
}
func (m *DeployModelOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployModelOperationMetadata.Merge(m, src)
}
func (m *DeployModelOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_DeployModelOperationMetadata.Size(m)
}
func (m *DeployModelOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployModelOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeployModelOperationMetadata proto.InternalMessageInfo

// Details of UndeployModel operation.
type UndeployModelOperationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UndeployModelOperationMetadata) Reset()         { *m = UndeployModelOperationMetadata{} }
func (m *UndeployModelOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*UndeployModelOperationMetadata) ProtoMessage()    {}
func (*UndeployModelOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{3}
}

func (m *UndeployModelOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UndeployModelOperationMetadata.Unmarshal(m, b)
}
func (m *UndeployModelOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UndeployModelOperationMetadata.Marshal(b, m, deterministic)
}
func (m *UndeployModelOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UndeployModelOperationMetadata.Merge(m, src)
}
func (m *UndeployModelOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_UndeployModelOperationMetadata.Size(m)
}
func (m *UndeployModelOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UndeployModelOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UndeployModelOperationMetadata proto.InternalMessageInfo

// Details of CreateDataset operation.
type CreateDatasetOperationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDatasetOperationMetadata) Reset()         { *m = CreateDatasetOperationMetadata{} }
func (m *CreateDatasetOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateDatasetOperationMetadata) ProtoMessage()    {}
func (*CreateDatasetOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{4}
}

func (m *CreateDatasetOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDatasetOperationMetadata.Unmarshal(m, b)
}
func (m *CreateDatasetOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDatasetOperationMetadata.Marshal(b, m, deterministic)
}
func (m *CreateDatasetOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDatasetOperationMetadata.Merge(m, src)
}
func (m *CreateDatasetOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateDatasetOperationMetadata.Size(m)
}
func (m *CreateDatasetOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDatasetOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDatasetOperationMetadata proto.InternalMessageInfo

// Details of CreateModel operation.
type CreateModelOperationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateModelOperationMetadata) Reset()         { *m = CreateModelOperationMetadata{} }
func (m *CreateModelOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateModelOperationMetadata) ProtoMessage()    {}
func (*CreateModelOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{5}
}

func (m *CreateModelOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateModelOperationMetadata.Unmarshal(m, b)
}
func (m *CreateModelOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateModelOperationMetadata.Marshal(b, m, deterministic)
}
func (m *CreateModelOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateModelOperationMetadata.Merge(m, src)
}
func (m *CreateModelOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateModelOperationMetadata.Size(m)
}
func (m *CreateModelOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateModelOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateModelOperationMetadata proto.InternalMessageInfo

// Details of ImportData operation.
type ImportDataOperationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportDataOperationMetadata) Reset()         { *m = ImportDataOperationMetadata{} }
func (m *ImportDataOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*ImportDataOperationMetadata) ProtoMessage()    {}
func (*ImportDataOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{6}
}

func (m *ImportDataOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportDataOperationMetadata.Unmarshal(m, b)
}
func (m *ImportDataOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportDataOperationMetadata.Marshal(b, m, deterministic)
}
func (m *ImportDataOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportDataOperationMetadata.Merge(m, src)
}
func (m *ImportDataOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_ImportDataOperationMetadata.Size(m)
}
func (m *ImportDataOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportDataOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ImportDataOperationMetadata proto.InternalMessageInfo

// Details of ExportData operation.
type ExportDataOperationMetadata struct {
	// Output only. Information further describing this export data's output.
	OutputInfo           *ExportDataOperationMetadata_ExportDataOutputInfo `protobuf:"bytes,1,opt,name=output_info,json=outputInfo,proto3" json:"output_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *ExportDataOperationMetadata) Reset()         { *m = ExportDataOperationMetadata{} }
func (m *ExportDataOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*ExportDataOperationMetadata) ProtoMessage()    {}
func (*ExportDataOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{7}
}

func (m *ExportDataOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportDataOperationMetadata.Unmarshal(m, b)
}
func (m *ExportDataOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportDataOperationMetadata.Marshal(b, m, deterministic)
}
func (m *ExportDataOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportDataOperationMetadata.Merge(m, src)
}
func (m *ExportDataOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_ExportDataOperationMetadata.Size(m)
}
func (m *ExportDataOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportDataOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ExportDataOperationMetadata proto.InternalMessageInfo

func (m *ExportDataOperationMetadata) GetOutputInfo() *ExportDataOperationMetadata_ExportDataOutputInfo {
	if m != nil {
		return m.OutputInfo
	}
	return nil
}

// Further describes this export data's output.
// Supplements
// [OutputConfig][google.cloud.automl.v1.OutputConfig].
type ExportDataOperationMetadata_ExportDataOutputInfo struct {
	// The output location to which the exported data is written.
	//
	// Types that are valid to be assigned to OutputLocation:
	//	*ExportDataOperationMetadata_ExportDataOutputInfo_GcsOutputDirectory
	OutputLocation       isExportDataOperationMetadata_ExportDataOutputInfo_OutputLocation `protobuf_oneof:"output_location"`
	XXX_NoUnkeyedLiteral struct{}                                                          `json:"-"`
	XXX_unrecognized     []byte                                                            `json:"-"`
	XXX_sizecache        int32                                                             `json:"-"`
}

func (m *ExportDataOperationMetadata_ExportDataOutputInfo) Reset() {
	*m = ExportDataOperationMetadata_ExportDataOutputInfo{}
}
func (m *ExportDataOperationMetadata_ExportDataOutputInfo) String() string {
	return proto.CompactTextString(m)
}
func (*ExportDataOperationMetadata_ExportDataOutputInfo) ProtoMessage() {}
func (*ExportDataOperationMetadata_ExportDataOutputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{7, 0}
}

func (m *ExportDataOperationMetadata_ExportDataOutputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportDataOperationMetadata_ExportDataOutputInfo.Unmarshal(m, b)
}
func (m *ExportDataOperationMetadata_ExportDataOutputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportDataOperationMetadata_ExportDataOutputInfo.Marshal(b, m, deterministic)
}
func (m *ExportDataOperationMetadata_ExportDataOutputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportDataOperationMetadata_ExportDataOutputInfo.Merge(m, src)
}
func (m *ExportDataOperationMetadata_ExportDataOutputInfo) XXX_Size() int {
	return xxx_messageInfo_ExportDataOperationMetadata_ExportDataOutputInfo.Size(m)
}
func (m *ExportDataOperationMetadata_ExportDataOutputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportDataOperationMetadata_ExportDataOutputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExportDataOperationMetadata_ExportDataOutputInfo proto.InternalMessageInfo

type isExportDataOperationMetadata_ExportDataOutputInfo_OutputLocation interface {
	isExportDataOperationMetadata_ExportDataOutputInfo_OutputLocation()
}

type ExportDataOperationMetadata_ExportDataOutputInfo_GcsOutputDirectory struct {
	GcsOutputDirectory string `protobuf:"bytes,1,opt,name=gcs_output_directory,json=gcsOutputDirectory,proto3,oneof"`
}

func (*ExportDataOperationMetadata_ExportDataOutputInfo_GcsOutputDirectory) isExportDataOperationMetadata_ExportDataOutputInfo_OutputLocation() {
}

func (m *ExportDataOperationMetadata_ExportDataOutputInfo) GetOutputLocation() isExportDataOperationMetadata_ExportDataOutputInfo_OutputLocation {
	if m != nil {
		return m.OutputLocation
	}
	return nil
}

func (m *ExportDataOperationMetadata_ExportDataOutputInfo) GetGcsOutputDirectory() string {
	if x, ok := m.GetOutputLocation().(*ExportDataOperationMetadata_ExportDataOutputInfo_GcsOutputDirectory); ok {
		return x.GcsOutputDirectory
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExportDataOperationMetadata_ExportDataOutputInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExportDataOperationMetadata_ExportDataOutputInfo_GcsOutputDirectory)(nil),
	}
}

// Details of BatchPredict operation.
type BatchPredictOperationMetadata struct {
	// Output only. The input config that was given upon starting this
	// batch predict operation.
	InputConfig *BatchPredictInputConfig `protobuf:"bytes,1,opt,name=input_config,json=inputConfig,proto3" json:"input_config,omitempty"`
	// Output only. Information further describing this batch predict's output.
	OutputInfo           *BatchPredictOperationMetadata_BatchPredictOutputInfo `protobuf:"bytes,2,opt,name=output_info,json=outputInfo,proto3" json:"output_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                              `json:"-"`
	XXX_unrecognized     []byte                                                `json:"-"`
	XXX_sizecache        int32                                                 `json:"-"`
}

func (m *BatchPredictOperationMetadata) Reset()         { *m = BatchPredictOperationMetadata{} }
func (m *BatchPredictOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*BatchPredictOperationMetadata) ProtoMessage()    {}
func (*BatchPredictOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{8}
}

func (m *BatchPredictOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchPredictOperationMetadata.Unmarshal(m, b)
}
func (m *BatchPredictOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchPredictOperationMetadata.Marshal(b, m, deterministic)
}
func (m *BatchPredictOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchPredictOperationMetadata.Merge(m, src)
}
func (m *BatchPredictOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_BatchPredictOperationMetadata.Size(m)
}
func (m *BatchPredictOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchPredictOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_BatchPredictOperationMetadata proto.InternalMessageInfo

func (m *BatchPredictOperationMetadata) GetInputConfig() *BatchPredictInputConfig {
	if m != nil {
		return m.InputConfig
	}
	return nil
}

func (m *BatchPredictOperationMetadata) GetOutputInfo() *BatchPredictOperationMetadata_BatchPredictOutputInfo {
	if m != nil {
		return m.OutputInfo
	}
	return nil
}

// Further describes this batch predict's output.
// Supplements
//
// [BatchPredictOutputConfig][google.cloud.automl.v1.BatchPredictOutputConfig].
type BatchPredictOperationMetadata_BatchPredictOutputInfo struct {
	// The output location into which prediction output is written.
	//
	// Types that are valid to be assigned to OutputLocation:
	//	*BatchPredictOperationMetadata_BatchPredictOutputInfo_GcsOutputDirectory
	OutputLocation       isBatchPredictOperationMetadata_BatchPredictOutputInfo_OutputLocation `protobuf_oneof:"output_location"`
	XXX_NoUnkeyedLiteral struct{}                                                              `json:"-"`
	XXX_unrecognized     []byte                                                                `json:"-"`
	XXX_sizecache        int32                                                                 `json:"-"`
}

func (m *BatchPredictOperationMetadata_BatchPredictOutputInfo) Reset() {
	*m = BatchPredictOperationMetadata_BatchPredictOutputInfo{}
}
func (m *BatchPredictOperationMetadata_BatchPredictOutputInfo) String() string {
	return proto.CompactTextString(m)
}
func (*BatchPredictOperationMetadata_BatchPredictOutputInfo) ProtoMessage() {}
func (*BatchPredictOperationMetadata_BatchPredictOutputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{8, 0}
}

func (m *BatchPredictOperationMetadata_BatchPredictOutputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchPredictOperationMetadata_BatchPredictOutputInfo.Unmarshal(m, b)
}
func (m *BatchPredictOperationMetadata_BatchPredictOutputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchPredictOperationMetadata_BatchPredictOutputInfo.Marshal(b, m, deterministic)
}
func (m *BatchPredictOperationMetadata_BatchPredictOutputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchPredictOperationMetadata_BatchPredictOutputInfo.Merge(m, src)
}
func (m *BatchPredictOperationMetadata_BatchPredictOutputInfo) XXX_Size() int {
	return xxx_messageInfo_BatchPredictOperationMetadata_BatchPredictOutputInfo.Size(m)
}
func (m *BatchPredictOperationMetadata_BatchPredictOutputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchPredictOperationMetadata_BatchPredictOutputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BatchPredictOperationMetadata_BatchPredictOutputInfo proto.InternalMessageInfo

type isBatchPredictOperationMetadata_BatchPredictOutputInfo_OutputLocation interface {
	isBatchPredictOperationMetadata_BatchPredictOutputInfo_OutputLocation()
}

type BatchPredictOperationMetadata_BatchPredictOutputInfo_GcsOutputDirectory struct {
	GcsOutputDirectory string `protobuf:"bytes,1,opt,name=gcs_output_directory,json=gcsOutputDirectory,proto3,oneof"`
}

func (*BatchPredictOperationMetadata_BatchPredictOutputInfo_GcsOutputDirectory) isBatchPredictOperationMetadata_BatchPredictOutputInfo_OutputLocation() {
}

func (m *BatchPredictOperationMetadata_BatchPredictOutputInfo) GetOutputLocation() isBatchPredictOperationMetadata_BatchPredictOutputInfo_OutputLocation {
	if m != nil {
		return m.OutputLocation
	}
	return nil
}

func (m *BatchPredictOperationMetadata_BatchPredictOutputInfo) GetGcsOutputDirectory() string {
	if x, ok := m.GetOutputLocation().(*BatchPredictOperationMetadata_BatchPredictOutputInfo_GcsOutputDirectory); ok {
		return x.GcsOutputDirectory
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BatchPredictOperationMetadata_BatchPredictOutputInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BatchPredictOperationMetadata_BatchPredictOutputInfo_GcsOutputDirectory)(nil),
	}
}

// Details of ExportModel operation.
type ExportModelOperationMetadata struct {
	// Output only. Information further describing the output of this model
	// export.
	OutputInfo           *ExportModelOperationMetadata_ExportModelOutputInfo `protobuf:"bytes,2,opt,name=output_info,json=outputInfo,proto3" json:"output_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *ExportModelOperationMetadata) Reset()         { *m = ExportModelOperationMetadata{} }
func (m *ExportModelOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*ExportModelOperationMetadata) ProtoMessage()    {}
func (*ExportModelOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{9}
}

func (m *ExportModelOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportModelOperationMetadata.Unmarshal(m, b)
}
func (m *ExportModelOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportModelOperationMetadata.Marshal(b, m, deterministic)
}
func (m *ExportModelOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportModelOperationMetadata.Merge(m, src)
}
func (m *ExportModelOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_ExportModelOperationMetadata.Size(m)
}
func (m *ExportModelOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportModelOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ExportModelOperationMetadata proto.InternalMessageInfo

func (m *ExportModelOperationMetadata) GetOutputInfo() *ExportModelOperationMetadata_ExportModelOutputInfo {
	if m != nil {
		return m.OutputInfo
	}
	return nil
}

// Further describes the output of model export.
// Supplements
// [ModelExportOutputConfig][google.cloud.automl.v1.ModelExportOutputConfig].
type ExportModelOperationMetadata_ExportModelOutputInfo struct {
	// The full path of the Google Cloud Storage directory created, into which
	// the model will be exported.
	GcsOutputDirectory   string   `protobuf:"bytes,1,opt,name=gcs_output_directory,json=gcsOutputDirectory,proto3" json:"gcs_output_directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportModelOperationMetadata_ExportModelOutputInfo) Reset() {
	*m = ExportModelOperationMetadata_ExportModelOutputInfo{}
}
func (m *ExportModelOperationMetadata_ExportModelOutputInfo) String() string {
	return proto.CompactTextString(m)
}
func (*ExportModelOperationMetadata_ExportModelOutputInfo) ProtoMessage() {}
func (*ExportModelOperationMetadata_ExportModelOutputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f887f15250a8bfee, []int{9, 0}
}

func (m *ExportModelOperationMetadata_ExportModelOutputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportModelOperationMetadata_ExportModelOutputInfo.Unmarshal(m, b)
}
func (m *ExportModelOperationMetadata_ExportModelOutputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportModelOperationMetadata_ExportModelOutputInfo.Marshal(b, m, deterministic)
}
func (m *ExportModelOperationMetadata_ExportModelOutputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportModelOperationMetadata_ExportModelOutputInfo.Merge(m, src)
}
func (m *ExportModelOperationMetadata_ExportModelOutputInfo) XXX_Size() int {
	return xxx_messageInfo_ExportModelOperationMetadata_ExportModelOutputInfo.Size(m)
}
func (m *ExportModelOperationMetadata_ExportModelOutputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportModelOperationMetadata_ExportModelOutputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExportModelOperationMetadata_ExportModelOutputInfo proto.InternalMessageInfo

func (m *ExportModelOperationMetadata_ExportModelOutputInfo) GetGcsOutputDirectory() string {
	if m != nil {
		return m.GcsOutputDirectory
	}
	return ""
}

func init() {
	proto.RegisterType((*OperationMetadata)(nil), "google.cloud.automl.v1.OperationMetadata")
	proto.RegisterType((*DeleteOperationMetadata)(nil), "google.cloud.automl.v1.DeleteOperationMetadata")
	proto.RegisterType((*DeployModelOperationMetadata)(nil), "google.cloud.automl.v1.DeployModelOperationMetadata")
	proto.RegisterType((*UndeployModelOperationMetadata)(nil), "google.cloud.automl.v1.UndeployModelOperationMetadata")
	proto.RegisterType((*CreateDatasetOperationMetadata)(nil), "google.cloud.automl.v1.CreateDatasetOperationMetadata")
	proto.RegisterType((*CreateModelOperationMetadata)(nil), "google.cloud.automl.v1.CreateModelOperationMetadata")
	proto.RegisterType((*ImportDataOperationMetadata)(nil), "google.cloud.automl.v1.ImportDataOperationMetadata")
	proto.RegisterType((*ExportDataOperationMetadata)(nil), "google.cloud.automl.v1.ExportDataOperationMetadata")
	proto.RegisterType((*ExportDataOperationMetadata_ExportDataOutputInfo)(nil), "google.cloud.automl.v1.ExportDataOperationMetadata.ExportDataOutputInfo")
	proto.RegisterType((*BatchPredictOperationMetadata)(nil), "google.cloud.automl.v1.BatchPredictOperationMetadata")
	proto.RegisterType((*BatchPredictOperationMetadata_BatchPredictOutputInfo)(nil), "google.cloud.automl.v1.BatchPredictOperationMetadata.BatchPredictOutputInfo")
	proto.RegisterType((*ExportModelOperationMetadata)(nil), "google.cloud.automl.v1.ExportModelOperationMetadata")
	proto.RegisterType((*ExportModelOperationMetadata_ExportModelOutputInfo)(nil), "google.cloud.automl.v1.ExportModelOperationMetadata.ExportModelOutputInfo")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/operations.proto", fileDescriptor_f887f15250a8bfee)
}

var fileDescriptor_f887f15250a8bfee = []byte{
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xf3, 0x34,
	0x18, 0xa6, 0x1d, 0x7f, 0x73, 0x19, 0xdd, 0xb2, 0xae, 0xeb, 0xba, 0x1f, 0xaa, 0x08, 0x89, 0x71,
	0x40, 0x42, 0x37, 0xe0, 0xa0, 0x83, 0x03, 0xb6, 0x02, 0x2b, 0xda, 0xc4, 0x54, 0x60, 0x42, 0x68,
	0x28, 0x72, 0x1d, 0x37, 0xb5, 0x96, 0xc4, 0x96, 0xe3, 0x54, 0xdb, 0x2d, 0xed, 0x9c, 0x2b, 0xe0,
	0x8c, 0x0b, 0xe0, 0x16, 0x90, 0xb8, 0x0a, 0x14, 0xdb, 0xe9, 0xdf, 0x12, 0x7f, 0x9f, 0xa6, 0xef,
	0xac, 0xf5, 0xf3, 0x3c, 0xef, 0xe3, 0xe7, 0xed, 0xeb, 0x57, 0x05, 0x9f, 0x04, 0x94, 0x06, 0x21,
	0x76, 0x51, 0x48, 0x53, 0xdf, 0x85, 0xa9, 0xa0, 0x51, 0xe8, 0x4e, 0xbb, 0x2e, 0x65, 0x98, 0x43,
	0x41, 0x68, 0x9c, 0x38, 0x8c, 0x53, 0x41, 0xad, 0xa6, 0x22, 0x3a, 0x92, 0xe8, 0x28, 0xa2, 0x33,
	0xed, 0xb6, 0x3f, 0x2e, 0x29, 0xe0, 0x43, 0x01, 0x13, 0x2c, 0x94, 0xba, 0xfd, 0x51, 0x09, 0x8b,
	0x50, 0x4d, 0xb0, 0x4b, 0x08, 0x11, 0xf5, 0x71, 0xa8, 0x39, 0x9f, 0x99, 0x38, 0x1e, 0x9e, 0xc2,
	0x30, 0x95, 0x57, 0xd6, 0x74, 0xb7, 0x84, 0xce, 0x38, 0xf6, 0x09, 0xca, 0x88, 0x5e, 0x82, 0xf9,
	0x94, 0x20, 0xac, 0x05, 0x65, 0x51, 0x96, 0x59, 0xfb, 0x9a, 0x25, 0xbf, 0x8d, 0xd2, 0xb1, 0x8b,
	0x23, 0x26, 0x1e, 0x57, 0x72, 0xce, 0x40, 0x41, 0x22, 0x9c, 0x08, 0x18, 0x31, 0x4d, 0xd8, 0xd5,
	0x04, 0xce, 0x90, 0x9b, 0x08, 0x28, 0x52, 0xdd, 0xdf, 0xf6, 0x81, 0x06, 0x20, 0x23, 0x2e, 0x8c,
	0x63, 0x2a, 0x16, 0xbb, 0x6f, 0xff, 0xb9, 0x0e, 0xb6, 0x7e, 0xca, 0x7f, 0x92, 0x6b, 0x2c, 0x60,
	0xd6, 0x5e, 0xeb, 0x37, 0xf0, 0xa1, 0x8f, 0x43, 0x2c, 0xb0, 0xe7, 0x63, 0x01, 0x49, 0x98, 0xb4,
	0xde, 0xef, 0x54, 0x8e, 0x6b, 0x27, 0xae, 0x53, 0xfc, 0x63, 0x39, 0x7d, 0xc9, 0x7e, 0x56, 0xe8,
	0xf2, 0xad, 0xe1, 0x86, 0x2a, 0xd4, 0x57, 0x75, 0xac, 0x09, 0x68, 0xf8, 0x98, 0x85, 0xf4, 0xd1,
	0x53, 0xcd, 0xcd, 0xeb, 0xb7, 0x64, 0xfd, 0x2f, 0xca, 0xeb, 0x67, 0x9a, 0xeb, 0x4c, 0x52, 0x64,
	0x62, 0xf9, 0x73, 0x3c, 0x77, 0x8a, 0x41, 0x33, 0x8d, 0x0b, 0xbd, 0xf6, 0xa4, 0xd7, 0x57, 0x65,
	0x5e, 0xbf, 0xc6, 0xbe, 0xd9, 0xad, 0x91, 0xc6, 0x05, 0x7e, 0x13, 0xd0, 0x40, 0x1c, 0x43, 0x81,
	0x57, 0xdc, 0x80, 0x39, 0xd9, 0x85, 0xd4, 0x94, 0x27, 0x43, 0x73, 0x7c, 0x21, 0x99, 0x76, 0xd2,
	0x6f, 0x61, 0xe6, 0x75, 0x64, 0x4e, 0xa6, 0xbc, 0xfa, 0x4a, 0x54, 0x98, 0x0c, 0x2d, 0x32, 0x72,
	0x3f, 0x0c, 0xb6, 0x49, 0xc4, 0x28, 0x17, 0xd2, 0x6f, 0x66, 0x56, 0x97, 0x66, 0xa7, 0x65, 0x66,
	0x03, 0x29, 0xc9, 0x4a, 0x15, 0x39, 0x6d, 0x91, 0x19, 0x9c, 0xdb, 0xdc, 0x83, 0x9d, 0x11, 0x14,
	0x68, 0xe2, 0xe9, 0x77, 0x34, 0x33, 0xda, 0x94, 0x46, 0x5f, 0x96, 0x19, 0x9d, 0x67, 0xa2, 0x1b,
	0xa5, 0x29, 0xb2, 0xda, 0x1e, 0x2d, 0x10, 0x16, 0x32, 0xe1, 0x87, 0xe7, 0x99, 0x76, 0xcc, 0x99,
	0xbe, 0x7b, 0x30, 0x66, 0xc2, 0x0f, 0xab, 0x99, 0x26, 0xa0, 0xa1, 0x6d, 0x96, 0x87, 0xa2, 0x69,
	0x1e, 0x0a, 0xe5, 0x53, 0x3e, 0x14, 0x78, 0x8e, 0xe7, 0x4e, 0x9f, 0x82, 0x4d, 0xc6, 0x69, 0xc0,
	0x71, 0x92, 0x78, 0x0c, 0x73, 0x84, 0x63, 0xd1, 0xda, 0xe8, 0x54, 0x8e, 0xdf, 0x19, 0xd6, 0xf3,
	0xf3, 0x1b, 0x75, 0x6c, 0x7d, 0x03, 0x36, 0x19, 0xe4, 0x82, 0xc0, 0xd0, 0x1b, 0x43, 0x12, 0xa6,
	0x1c, 0x27, 0xad, 0x6a, 0x67, 0xed, 0xb8, 0x76, 0x62, 0xe5, 0x17, 0xe2, 0x0c, 0x39, 0x3f, 0xcb,
	0x2d, 0x32, 0xac, 0x6b, 0xee, 0xf7, 0x9a, 0x6a, 0x9d, 0x81, 0x9a, 0x1e, 0xbf, 0x6c, 0x07, 0xb5,
	0xd6, 0x64, 0x94, 0x76, 0xae, 0xcc, 0x17, 0x94, 0xf3, 0x4b, 0xbe, 0xa0, 0x86, 0x40, 0xd1, 0xb3,
	0x83, 0x4c, 0x9c, 0x32, 0x7f, 0x26, 0x7e, 0xfb, 0xd5, 0x62, 0x45, 0xcf, 0x0e, 0xce, 0xd7, 0xc1,
	0x7b, 0xba, 0x81, 0xf6, 0x1e, 0xd8, 0x2d, 0xd9, 0x39, 0xf6, 0x11, 0x38, 0x30, 0xad, 0x0b, 0xbb,
	0x03, 0x8e, 0xcc, 0x4f, 0x3c, 0x63, 0x98, 0x9f, 0x4a, 0xe6, 0x61, 0x7a, 0xb8, 0xf6, 0x21, 0xd8,
	0x37, 0xcc, 0xbf, 0xfd, 0x6f, 0x05, 0xec, 0x1b, 0x66, 0xc9, 0x22, 0xa0, 0x46, 0x53, 0xc1, 0x52,
	0xe1, 0x91, 0x78, 0x4c, 0x5b, 0x15, 0xd9, 0xa5, 0xcb, 0x17, 0x4c, 0xe5, 0x22, 0x26, 0x0b, 0x0e,
	0xe2, 0x31, 0x1d, 0x02, 0x3a, 0xfb, 0xdc, 0xfe, 0x03, 0x34, 0x8a, 0x38, 0xd6, 0x09, 0x68, 0x04,
	0x28, 0xf1, 0xf4, 0x35, 0x7c, 0xc2, 0x31, 0x12, 0x94, 0x3f, 0xca, 0xbb, 0xac, 0x67, 0x33, 0x18,
	0xa0, 0x44, 0xd1, 0xfb, 0x39, 0x76, 0xbe, 0x05, 0xea, 0x9a, 0x1f, 0x52, 0x24, 0xef, 0x61, 0xff,
	0x55, 0x05, 0x87, 0xc6, 0x07, 0x6a, 0x0d, 0xc1, 0x07, 0x24, 0xce, 0x34, 0x88, 0xc6, 0x63, 0x12,
	0xe8, 0xb0, 0xee, 0xeb, 0xbc, 0xf6, 0x41, 0xa6, 0xbb, 0x90, 0xb2, 0x61, 0x8d, 0xcc, 0xbf, 0x58,
	0xd1, 0x72, 0xff, 0xaa, 0xb2, 0xe4, 0xd5, 0x8b, 0x16, 0xc8, 0x32, 0x5a, 0xdc, 0x43, 0x0f, 0x34,
	0x8b, 0x59, 0x6f, 0xaa, 0x8b, 0xff, 0x54, 0xc0, 0x81, 0x69, 0x27, 0x58, 0xf7, 0x45, 0x81, 0x7f,
	0x7c, 0xc9, 0x7a, 0x59, 0x02, 0x8b, 0xe3, 0x0e, 0xc0, 0x4e, 0x21, 0xc9, 0xfa, 0xdc, 0x94, 0xb6,
	0x30, 0xeb, 0x53, 0x05, 0xb4, 0x11, 0x8d, 0x4a, 0x2e, 0x7a, 0x53, 0xf9, 0xfd, 0x6b, 0x8d, 0x04,
	0x34, 0x84, 0x71, 0xe0, 0x50, 0x1e, 0xb8, 0x01, 0x8e, 0xe5, 0x9e, 0xd0, 0x7f, 0xc3, 0x20, 0x23,
	0xc9, 0xea, 0x3f, 0xab, 0x33, 0xf5, 0xe9, 0xa9, 0xda, 0xfc, 0x41, 0xc9, 0x2f, 0x64, 0xe1, 0x6f,
	0x53, 0x41, 0xaf, 0xaf, 0x9c, 0xdb, 0xee, 0xdf, 0x39, 0x70, 0x27, 0x81, 0x3b, 0x05, 0xdc, 0xdd,
	0x76, 0xff, 0xab, 0xee, 0x29, 0xa0, 0xd7, 0x93, 0x48, 0xaf, 0xa7, 0xa0, 0x5e, 0xef, 0xb6, 0x3b,
	0x7a, 0x57, 0xda, 0x9e, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x96, 0x02, 0x20, 0x48, 0xdb, 0x0a,
	0x00, 0x00,
}
