// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/genomics/v1/readgroup.proto

package genomics

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf3 "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A read group is all the data that's processed the same way by the sequencer.
type ReadGroup struct {
	// The server-generated read group ID, unique for all read groups.
	// Note: This is different than the @RG ID field in the SAM spec. For that
	// value, see [name][google.genomics.v1.ReadGroup.name].
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The dataset to which this read group belongs.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The read group name. This corresponds to the @RG ID field in the SAM spec.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// A free-form text description of this read group.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// A client-supplied sample identifier for the reads in this read group.
	SampleId string `protobuf:"bytes,5,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	// The experiment used to generate this read group.
	Experiment *ReadGroup_Experiment `protobuf:"bytes,6,opt,name=experiment" json:"experiment,omitempty"`
	// The predicted insert size of this read group. The insert size is the length
	// the sequenced DNA fragment from end-to-end, not including the adapters.
	PredictedInsertSize int32 `protobuf:"varint,7,opt,name=predicted_insert_size,json=predictedInsertSize,proto3" json:"predicted_insert_size,omitempty"`
	// The programs used to generate this read group. Programs are always
	// identical for all read groups within a read group set. For this reason,
	// only the first read group in a returned set will have this field
	// populated.
	Programs []*ReadGroup_Program `protobuf:"bytes,10,rep,name=programs" json:"programs,omitempty"`
	// The reference set the reads in this read group are aligned to.
	ReferenceSetId string `protobuf:"bytes,11,opt,name=reference_set_id,json=referenceSetId,proto3" json:"reference_set_id,omitempty"`
	// A map of additional read group information. This must be of the form
	// map<string, string[]> (string key mapping to a list of string values).
	Info map[string]*google_protobuf3.ListValue `protobuf:"bytes,12,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ReadGroup) Reset()                    { *m = ReadGroup{} }
func (m *ReadGroup) String() string            { return proto.CompactTextString(m) }
func (*ReadGroup) ProtoMessage()               {}
func (*ReadGroup) Descriptor() ([]byte, []int) { return fileDescriptorReadgroup, []int{0} }

func (m *ReadGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadGroup) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *ReadGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ReadGroup) GetSampleId() string {
	if m != nil {
		return m.SampleId
	}
	return ""
}

func (m *ReadGroup) GetExperiment() *ReadGroup_Experiment {
	if m != nil {
		return m.Experiment
	}
	return nil
}

func (m *ReadGroup) GetPredictedInsertSize() int32 {
	if m != nil {
		return m.PredictedInsertSize
	}
	return 0
}

func (m *ReadGroup) GetPrograms() []*ReadGroup_Program {
	if m != nil {
		return m.Programs
	}
	return nil
}

func (m *ReadGroup) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

func (m *ReadGroup) GetInfo() map[string]*google_protobuf3.ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReadGroup_Experiment struct {
	// A client-supplied library identifier; a library is a collection of DNA
	// fragments which have been prepared for sequencing from a sample. This
	// field is important for quality control as error or bias can be introduced
	// during sample preparation.
	LibraryId string `protobuf:"bytes,1,opt,name=library_id,json=libraryId,proto3" json:"library_id,omitempty"`
	// The platform unit used as part of this experiment, for example
	// flowcell-barcode.lane for Illumina or slide for SOLiD. Corresponds to the
	// @RG PU field in the SAM spec.
	PlatformUnit string `protobuf:"bytes,2,opt,name=platform_unit,json=platformUnit,proto3" json:"platform_unit,omitempty"`
	// The sequencing center used as part of this experiment.
	SequencingCenter string `protobuf:"bytes,3,opt,name=sequencing_center,json=sequencingCenter,proto3" json:"sequencing_center,omitempty"`
	// The instrument model used as part of this experiment. This maps to
	// sequencing technology in the SAM spec.
	InstrumentModel string `protobuf:"bytes,4,opt,name=instrument_model,json=instrumentModel,proto3" json:"instrument_model,omitempty"`
}

func (m *ReadGroup_Experiment) Reset()                    { *m = ReadGroup_Experiment{} }
func (m *ReadGroup_Experiment) String() string            { return proto.CompactTextString(m) }
func (*ReadGroup_Experiment) ProtoMessage()               {}
func (*ReadGroup_Experiment) Descriptor() ([]byte, []int) { return fileDescriptorReadgroup, []int{0, 0} }

func (m *ReadGroup_Experiment) GetLibraryId() string {
	if m != nil {
		return m.LibraryId
	}
	return ""
}

func (m *ReadGroup_Experiment) GetPlatformUnit() string {
	if m != nil {
		return m.PlatformUnit
	}
	return ""
}

func (m *ReadGroup_Experiment) GetSequencingCenter() string {
	if m != nil {
		return m.SequencingCenter
	}
	return ""
}

func (m *ReadGroup_Experiment) GetInstrumentModel() string {
	if m != nil {
		return m.InstrumentModel
	}
	return ""
}

type ReadGroup_Program struct {
	// The command line used to run this program.
	CommandLine string `protobuf:"bytes,1,opt,name=command_line,json=commandLine,proto3" json:"command_line,omitempty"`
	// The user specified locally unique ID of the program. Used along with
	// `prevProgramId` to define an ordering between programs.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The display name of the program. This is typically the colloquial name of
	// the tool used, for example 'bwa' or 'picard'.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The ID of the program run before this one.
	PrevProgramId string `protobuf:"bytes,4,opt,name=prev_program_id,json=prevProgramId,proto3" json:"prev_program_id,omitempty"`
	// The version of the program run.
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ReadGroup_Program) Reset()                    { *m = ReadGroup_Program{} }
func (m *ReadGroup_Program) String() string            { return proto.CompactTextString(m) }
func (*ReadGroup_Program) ProtoMessage()               {}
func (*ReadGroup_Program) Descriptor() ([]byte, []int) { return fileDescriptorReadgroup, []int{0, 1} }

func (m *ReadGroup_Program) GetCommandLine() string {
	if m != nil {
		return m.CommandLine
	}
	return ""
}

func (m *ReadGroup_Program) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadGroup_Program) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadGroup_Program) GetPrevProgramId() string {
	if m != nil {
		return m.PrevProgramId
	}
	return ""
}

func (m *ReadGroup_Program) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*ReadGroup)(nil), "google.genomics.v1.ReadGroup")
	proto.RegisterType((*ReadGroup_Experiment)(nil), "google.genomics.v1.ReadGroup.Experiment")
	proto.RegisterType((*ReadGroup_Program)(nil), "google.genomics.v1.ReadGroup.Program")
}
func (m *ReadGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadGroup) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.DatasetId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.DatasetId)))
		i += copy(dAtA[i:], m.DatasetId)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.SampleId) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.SampleId)))
		i += copy(dAtA[i:], m.SampleId)
	}
	if m.Experiment != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(m.Experiment.Size()))
		n1, err := m.Experiment.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.PredictedInsertSize != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(m.PredictedInsertSize))
	}
	if len(m.Programs) > 0 {
		for _, msg := range m.Programs {
			dAtA[i] = 0x52
			i++
			i = encodeVarintReadgroup(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ReferenceSetId) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.ReferenceSetId)))
		i += copy(dAtA[i:], m.ReferenceSetId)
	}
	if len(m.Info) > 0 {
		for k, _ := range m.Info {
			dAtA[i] = 0x62
			i++
			v := m.Info[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovReadgroup(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovReadgroup(uint64(len(k))) + msgSize
			i = encodeVarintReadgroup(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintReadgroup(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintReadgroup(dAtA, i, uint64(v.Size()))
				n2, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n2
			}
		}
	}
	return i, nil
}

func (m *ReadGroup_Experiment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadGroup_Experiment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LibraryId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.LibraryId)))
		i += copy(dAtA[i:], m.LibraryId)
	}
	if len(m.PlatformUnit) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.PlatformUnit)))
		i += copy(dAtA[i:], m.PlatformUnit)
	}
	if len(m.SequencingCenter) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.SequencingCenter)))
		i += copy(dAtA[i:], m.SequencingCenter)
	}
	if len(m.InstrumentModel) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.InstrumentModel)))
		i += copy(dAtA[i:], m.InstrumentModel)
	}
	return i, nil
}

func (m *ReadGroup_Program) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadGroup_Program) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CommandLine) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.CommandLine)))
		i += copy(dAtA[i:], m.CommandLine)
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.PrevProgramId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.PrevProgramId)))
		i += copy(dAtA[i:], m.PrevProgramId)
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintReadgroup(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	return i, nil
}

func encodeFixed64Readgroup(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Readgroup(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintReadgroup(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ReadGroup) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.DatasetId)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.SampleId)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	if m.Experiment != nil {
		l = m.Experiment.Size()
		n += 1 + l + sovReadgroup(uint64(l))
	}
	if m.PredictedInsertSize != 0 {
		n += 1 + sovReadgroup(uint64(m.PredictedInsertSize))
	}
	if len(m.Programs) > 0 {
		for _, e := range m.Programs {
			l = e.Size()
			n += 1 + l + sovReadgroup(uint64(l))
		}
	}
	l = len(m.ReferenceSetId)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	if len(m.Info) > 0 {
		for k, v := range m.Info {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovReadgroup(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovReadgroup(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovReadgroup(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ReadGroup_Experiment) Size() (n int) {
	var l int
	_ = l
	l = len(m.LibraryId)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.PlatformUnit)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.SequencingCenter)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.InstrumentModel)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	return n
}

func (m *ReadGroup_Program) Size() (n int) {
	var l int
	_ = l
	l = len(m.CommandLine)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.PrevProgramId)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovReadgroup(uint64(l))
	}
	return n
}

func sovReadgroup(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReadgroup(x uint64) (n int) {
	return sovReadgroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReadGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReadgroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatasetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DatasetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SampleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SampleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Experiment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Experiment == nil {
				m.Experiment = &ReadGroup_Experiment{}
			}
			if err := m.Experiment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PredictedInsertSize", wireType)
			}
			m.PredictedInsertSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PredictedInsertSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Programs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Programs = append(m.Programs, &ReadGroup_Program{})
			if err := m.Programs[len(m.Programs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferenceSetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReferenceSetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = make(map[string]*google_protobuf3.ListValue)
			}
			var mapkey string
			var mapvalue *google_protobuf3.ListValue
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowReadgroup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowReadgroup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthReadgroup
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowReadgroup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthReadgroup
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthReadgroup
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &google_protobuf3.ListValue{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipReadgroup(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthReadgroup
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Info[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReadgroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReadgroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadGroup_Experiment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReadgroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Experiment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Experiment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LibraryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LibraryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlatformUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlatformUnit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequencingCenter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SequencingCenter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstrumentModel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InstrumentModel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReadgroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReadgroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadGroup_Program) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReadgroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Program: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Program: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandLine", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommandLine = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevProgramId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevProgramId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReadgroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReadgroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReadgroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipReadgroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReadgroup
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReadgroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthReadgroup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowReadgroup
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipReadgroup(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthReadgroup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReadgroup   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("google/genomics/v1/readgroup.proto", fileDescriptorReadgroup) }

var fileDescriptorReadgroup = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0x95, 0xa7, 0xbf, 0x73, 0xa7, 0x3f, 0xf3, 0xf9, 0x13, 0x28, 0x1a, 0xa0, 0x1a, 0x8a, 0x80,
	0x41, 0x48, 0x19, 0x3a, 0x6c, 0x50, 0xbb, 0xa2, 0xa8, 0x82, 0x91, 0x8a, 0x54, 0x4d, 0x05, 0x0b,
	0x36, 0x91, 0x1b, 0xdf, 0x89, 0x2c, 0x12, 0x3b, 0xd8, 0xce, 0x88, 0xf6, 0x31, 0x78, 0x0a, 0x1e,
	0x05, 0xb1, 0xe2, 0x11, 0x50, 0x9f, 0x82, 0x25, 0xb2, 0xe3, 0xa4, 0x95, 0xa8, 0xba, 0x73, 0xce,
	0x39, 0x37, 0xe7, 0xfa, 0x9e, 0x9b, 0xc0, 0x6e, 0xa6, 0x54, 0x96, 0xe3, 0x38, 0x43, 0xa9, 0x0a,
	0x91, 0x9a, 0xf1, 0x62, 0x6f, 0xac, 0x91, 0xf1, 0x4c, 0xab, 0xaa, 0x8c, 0x4b, 0xad, 0xac, 0xa2,
	0xb4, 0xd6, 0xc4, 0x8d, 0x26, 0x5e, 0xec, 0x0d, 0xee, 0x87, 0x3a, 0x56, 0x8a, 0x31, 0x93, 0x52,
	0x59, 0x66, 0x85, 0x92, 0xa6, 0xae, 0x68, 0x59, 0xff, 0x74, 0x56, 0xcd, 0xc7, 0xc6, 0xea, 0x2a,
	0xb5, 0x35, 0xbb, 0xfb, 0x73, 0x15, 0xba, 0x33, 0x64, 0xfc, 0xad, 0xf3, 0xa0, 0x5b, 0xd0, 0x11,
	0x3c, 0x22, 0x43, 0x32, 0xea, 0xce, 0x3a, 0x82, 0xd3, 0x07, 0x00, 0x9c, 0x59, 0x66, 0xd0, 0x26,
	0x82, 0x47, 0x1d, 0x8f, 0x77, 0x03, 0x32, 0xe5, 0x94, 0xc2, 0xb2, 0x64, 0x05, 0x46, 0x4b, 0x9e,
	0xf0, 0x67, 0x3a, 0x84, 0x1e, 0x47, 0x93, 0x6a, 0x51, 0xba, 0x26, 0xa2, 0x65, 0x4f, 0x5d, 0x87,
	0xe8, 0x3d, 0xe8, 0x1a, 0x56, 0x94, 0x39, 0xba, 0x77, 0xae, 0x78, 0x7e, 0xbd, 0x06, 0xa6, 0x9c,
	0xbe, 0x03, 0xc0, 0xaf, 0x25, 0x6a, 0x51, 0xa0, 0xb4, 0xd1, 0xea, 0x90, 0x8c, 0x7a, 0x93, 0x51,
	0xfc, 0xef, 0xa5, 0xe3, 0xb6, 0xe9, 0xf8, 0xa8, 0xd5, 0xcf, 0xae, 0xd5, 0xd2, 0x09, 0xdc, 0x29,
	0x35, 0x72, 0x91, 0x5a, 0xe4, 0x89, 0x90, 0x06, 0xb5, 0x4d, 0x8c, 0xb8, 0xc0, 0x68, 0x6d, 0x48,
	0x46, 0x2b, 0xb3, 0xff, 0x5b, 0x72, 0xea, 0xb9, 0x53, 0x71, 0x81, 0xf4, 0x35, 0xac, 0x97, 0x5a,
	0x65, 0x9a, 0x15, 0x26, 0x82, 0xe1, 0xd2, 0xa8, 0x37, 0x79, 0x7c, 0xbb, 0xf7, 0x49, 0xad, 0x9e,
	0xb5, 0x65, 0x74, 0x04, 0x7d, 0x8d, 0x73, 0xd4, 0x28, 0x53, 0x4c, 0xc2, 0xe0, 0x7a, 0xfe, 0x92,
	0x5b, 0x2d, 0x7e, 0xea, 0xa7, 0x77, 0x00, 0xcb, 0x42, 0xce, 0x55, 0xb4, 0xe1, 0x8d, 0x9e, 0xde,
	0x6e, 0x34, 0x95, 0x73, 0x75, 0x24, 0xad, 0x3e, 0x9f, 0xf9, 0xa2, 0xc1, 0x77, 0x02, 0x70, 0x75,
	0x71, 0x17, 0x54, 0x2e, 0xce, 0x34, 0xd3, 0xe7, 0x49, 0x1b, 0x60, 0x37, 0x20, 0x53, 0x4e, 0x1f,
	0xc1, 0x66, 0x99, 0x33, 0x3b, 0x57, 0xba, 0x48, 0x2a, 0x29, 0x6c, 0x88, 0x72, 0xa3, 0x01, 0x3f,
	0x48, 0x61, 0xe9, 0x73, 0xf8, 0xcf, 0xe0, 0x97, 0x0a, 0x65, 0x2a, 0x64, 0x96, 0xa4, 0x28, 0x2d,
	0xea, 0x10, 0x6d, 0xff, 0x8a, 0x78, 0xe3, 0x71, 0xfa, 0x0c, 0xfa, 0x42, 0xba, 0x4d, 0x72, 0xf6,
	0x49, 0xa1, 0x38, 0xe6, 0x21, 0xeb, 0xed, 0x2b, 0xfc, 0xbd, 0x83, 0x07, 0xdf, 0x08, 0xac, 0x85,
	0x39, 0xd1, 0x87, 0xb0, 0x91, 0xaa, 0xa2, 0x60, 0x92, 0x27, 0xb9, 0x90, 0x18, 0x3a, 0xed, 0x05,
	0xec, 0x58, 0x48, 0x0c, 0x3b, 0xd8, 0x69, 0x77, 0xf0, 0xa6, 0x25, 0x7b, 0x02, 0xdb, 0xa5, 0xc6,
	0x45, 0x12, 0xa6, 0xee, 0xee, 0x5c, 0x9b, 0x6f, 0x3a, 0x38, 0x98, 0x4d, 0x39, 0x8d, 0x60, 0x6d,
	0x81, 0xda, 0xb8, 0x45, 0xac, 0x17, 0xad, 0x79, 0x1c, 0x9c, 0x42, 0xb7, 0x1d, 0x29, 0xed, 0xc3,
	0xd2, 0x67, 0x3c, 0x0f, 0xcd, 0xb8, 0x23, 0x7d, 0x01, 0x2b, 0x0b, 0x96, 0x57, 0xe8, 0xfb, 0xe8,
	0x4d, 0x06, 0x4d, 0x38, 0xcd, 0x47, 0x14, 0x1f, 0x0b, 0x63, 0x3f, 0x3a, 0xc5, 0xac, 0x16, 0xee,
	0x77, 0x5e, 0x91, 0xc3, 0xea, 0xc7, 0xe5, 0x0e, 0xf9, 0x75, 0xb9, 0x43, 0x7e, 0x5f, 0xee, 0x10,
	0xb8, 0x9b, 0xaa, 0xe2, 0x86, 0x50, 0x0f, 0xb7, 0xda, 0x54, 0x4f, 0xdc, 0xdb, 0x4e, 0xc8, 0xa7,
	0xfd, 0x46, 0xa5, 0x72, 0x26, 0xb3, 0x58, 0xe9, 0xcc, 0xfd, 0x03, 0xbc, 0xd7, 0xb8, 0xa6, 0x58,
	0x29, 0xcc, 0xf5, 0xff, 0xc2, 0x41, 0x73, 0xfe, 0x43, 0xc8, 0xd9, 0xaa, 0x57, 0xbe, 0xfc, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xe4, 0x07, 0x77, 0x57, 0x40, 0x04, 0x00, 0x00,
}
