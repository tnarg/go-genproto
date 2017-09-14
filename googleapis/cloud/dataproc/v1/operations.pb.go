// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/cloud/dataproc/v1/operations.proto

package dataproc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/longrunning"
import _ "github.com/gogo/protobuf/types"
import google_protobuf3 "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The operation state.
type ClusterOperationStatus_State int32

const (
	// Unused.
	ClusterOperationStatus_UNKNOWN ClusterOperationStatus_State = 0
	// The operation has been created.
	ClusterOperationStatus_PENDING ClusterOperationStatus_State = 1
	// The operation is running.
	ClusterOperationStatus_RUNNING ClusterOperationStatus_State = 2
	// The operation is done; either cancelled or completed.
	ClusterOperationStatus_DONE ClusterOperationStatus_State = 3
)

var ClusterOperationStatus_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "PENDING",
	2: "RUNNING",
	3: "DONE",
}
var ClusterOperationStatus_State_value = map[string]int32{
	"UNKNOWN": 0,
	"PENDING": 1,
	"RUNNING": 2,
	"DONE":    3,
}

func (x ClusterOperationStatus_State) String() string {
	return proto.EnumName(ClusterOperationStatus_State_name, int32(x))
}
func (ClusterOperationStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorOperations, []int{0, 0}
}

// The status of the operation.
type ClusterOperationStatus struct {
	// [Output-only] A message containing the operation state.
	State ClusterOperationStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.dataproc.v1.ClusterOperationStatus_State" json:"state,omitempty"`
	// [Output-only] A message containing the detailed operation state.
	InnerState string `protobuf:"bytes,2,opt,name=inner_state,json=innerState,proto3" json:"inner_state,omitempty"`
	// [Output-only]A message containing any operation metadata details.
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	// [Output-only] The time this state was entered.
	StateStartTime *google_protobuf3.Timestamp `protobuf:"bytes,4,opt,name=state_start_time,json=stateStartTime" json:"state_start_time,omitempty"`
}

func (m *ClusterOperationStatus) Reset()                    { *m = ClusterOperationStatus{} }
func (m *ClusterOperationStatus) String() string            { return proto.CompactTextString(m) }
func (*ClusterOperationStatus) ProtoMessage()               {}
func (*ClusterOperationStatus) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{0} }

func (m *ClusterOperationStatus) GetState() ClusterOperationStatus_State {
	if m != nil {
		return m.State
	}
	return ClusterOperationStatus_UNKNOWN
}

func (m *ClusterOperationStatus) GetInnerState() string {
	if m != nil {
		return m.InnerState
	}
	return ""
}

func (m *ClusterOperationStatus) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *ClusterOperationStatus) GetStateStartTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.StateStartTime
	}
	return nil
}

// Metadata describing the operation.
type ClusterOperationMetadata struct {
	// [Output-only] Name of the cluster for the operation.
	ClusterName string `protobuf:"bytes,7,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// [Output-only] Cluster UUID for the operation.
	ClusterUuid string `protobuf:"bytes,8,opt,name=cluster_uuid,json=clusterUuid,proto3" json:"cluster_uuid,omitempty"`
	// [Output-only] Current operation status.
	Status *ClusterOperationStatus `protobuf:"bytes,9,opt,name=status" json:"status,omitempty"`
	// [Output-only] The previous operation status.
	StatusHistory []*ClusterOperationStatus `protobuf:"bytes,10,rep,name=status_history,json=statusHistory" json:"status_history,omitempty"`
	// [Output-only] The operation type.
	OperationType string `protobuf:"bytes,11,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// [Output-only] Short description of operation.
	Description string `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *ClusterOperationMetadata) Reset()         { *m = ClusterOperationMetadata{} }
func (m *ClusterOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*ClusterOperationMetadata) ProtoMessage()    {}
func (*ClusterOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorOperations, []int{1}
}

func (m *ClusterOperationMetadata) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterOperationMetadata) GetClusterUuid() string {
	if m != nil {
		return m.ClusterUuid
	}
	return ""
}

func (m *ClusterOperationMetadata) GetStatus() *ClusterOperationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterOperationMetadata) GetStatusHistory() []*ClusterOperationStatus {
	if m != nil {
		return m.StatusHistory
	}
	return nil
}

func (m *ClusterOperationMetadata) GetOperationType() string {
	if m != nil {
		return m.OperationType
	}
	return ""
}

func (m *ClusterOperationMetadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ClusterOperationStatus)(nil), "google.cloud.dataproc.v1.ClusterOperationStatus")
	proto.RegisterType((*ClusterOperationMetadata)(nil), "google.cloud.dataproc.v1.ClusterOperationMetadata")
	proto.RegisterEnum("google.cloud.dataproc.v1.ClusterOperationStatus_State", ClusterOperationStatus_State_name, ClusterOperationStatus_State_value)
}
func (m *ClusterOperationStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterOperationStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.State))
	}
	if len(m.InnerState) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOperations(dAtA, i, uint64(len(m.InnerState)))
		i += copy(dAtA[i:], m.InnerState)
	}
	if len(m.Details) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOperations(dAtA, i, uint64(len(m.Details)))
		i += copy(dAtA[i:], m.Details)
	}
	if m.StateStartTime != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.StateStartTime.Size()))
		n1, err := m.StateStartTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ClusterOperationMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterOperationMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClusterName) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintOperations(dAtA, i, uint64(len(m.ClusterName)))
		i += copy(dAtA[i:], m.ClusterName)
	}
	if len(m.ClusterUuid) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintOperations(dAtA, i, uint64(len(m.ClusterUuid)))
		i += copy(dAtA[i:], m.ClusterUuid)
	}
	if m.Status != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.Status.Size()))
		n2, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.StatusHistory) > 0 {
		for _, msg := range m.StatusHistory {
			dAtA[i] = 0x52
			i++
			i = encodeVarintOperations(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.OperationType) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintOperations(dAtA, i, uint64(len(m.OperationType)))
		i += copy(dAtA[i:], m.OperationType)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintOperations(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	return i, nil
}

func encodeFixed64Operations(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Operations(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintOperations(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ClusterOperationStatus) Size() (n int) {
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovOperations(uint64(m.State))
	}
	l = len(m.InnerState)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.StateStartTime != nil {
		l = m.StateStartTime.Size()
		n += 1 + l + sovOperations(uint64(l))
	}
	return n
}

func (m *ClusterOperationMetadata) Size() (n int) {
	var l int
	_ = l
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	l = len(m.ClusterUuid)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovOperations(uint64(l))
	}
	if len(m.StatusHistory) > 0 {
		for _, e := range m.StatusHistory {
			l = e.Size()
			n += 1 + l + sovOperations(uint64(l))
		}
	}
	l = len(m.OperationType)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	return n
}

func sovOperations(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOperations(x uint64) (n int) {
	return sovOperations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClusterOperationStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: ClusterOperationStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterOperationStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (ClusterOperationStatus_State(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InnerState", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InnerState = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StateStartTime == nil {
				m.StateStartTime = &google_protobuf3.Timestamp{}
			}
			if err := m.StateStartTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperations
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
func (m *ClusterOperationMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: ClusterOperationMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterOperationMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterUuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterUuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &ClusterOperationStatus{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusHistory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusHistory = append(m.StatusHistory, &ClusterOperationStatus{})
			if err := m.StatusHistory[len(m.StatusHistory)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperations
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
func skipOperations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperations
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
					return 0, ErrIntOverflowOperations
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
					return 0, ErrIntOverflowOperations
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
				return 0, ErrInvalidLengthOperations
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOperations
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
				next, err := skipOperations(dAtA[start:])
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
	ErrInvalidLengthOperations = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperations   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("google/cloud/dataproc/v1/operations.proto", fileDescriptorOperations) }

var fileDescriptorOperations = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x27, 0xdd, 0x9f, 0x6e, 0xce, 0x56, 0x2a, 0x1f, 0x90, 0x55, 0xa6, 0x2e, 0x14, 0x21, 0x95,
	0x4b, 0xc2, 0x86, 0x84, 0x90, 0xb8, 0xa0, 0xd1, 0x89, 0x21, 0x20, 0xad, 0xb2, 0x55, 0x93, 0xb8,
	0x44, 0x5e, 0x62, 0x82, 0xa5, 0xc4, 0xb6, 0x6c, 0x67, 0x52, 0x4f, 0x3c, 0x0b, 0x6f, 0xc3, 0x81,
	0x03, 0x8f, 0x80, 0xfa, 0x24, 0xc8, 0x76, 0x52, 0x95, 0x8d, 0x1d, 0xe0, 0x14, 0x7f, 0xdf, 0xef,
	0x4f, 0xbe, 0xdf, 0x17, 0x07, 0x3c, 0x2d, 0x38, 0x2f, 0x4a, 0x12, 0x65, 0x25, 0xaf, 0xf3, 0x28,
	0xc7, 0x1a, 0x0b, 0xc9, 0xb3, 0xe8, 0xfa, 0x28, 0xe2, 0x82, 0x48, 0xac, 0x29, 0x67, 0x2a, 0x14,
	0x92, 0x6b, 0x0e, 0x91, 0xa3, 0x86, 0x96, 0x1a, 0xb6, 0xd4, 0xf0, 0xfa, 0x68, 0x70, 0xd0, 0x98,
	0x60, 0x41, 0x23, 0xcc, 0x18, 0xd7, 0xeb, 0xba, 0xc1, 0xe3, 0x06, 0x2d, 0x39, 0x2b, 0x64, 0xcd,
	0x18, 0x65, 0xc5, 0x2d, 0xf3, 0xc1, 0xc3, 0x86, 0x64, 0xab, 0xab, 0xfa, 0x73, 0x44, 0x2a, 0xa1,
	0x17, 0x0d, 0x78, 0x78, 0x13, 0xd4, 0xb4, 0x22, 0x4a, 0xe3, 0x4a, 0x38, 0xc2, 0xe8, 0x5b, 0x07,
	0x3c, 0x78, 0x53, 0xd6, 0x4a, 0x13, 0x39, 0x6d, 0x9d, 0xcf, 0x35, 0xd6, 0xb5, 0x82, 0x1f, 0xc0,
	0x96, 0xd2, 0x58, 0x13, 0xe4, 0x05, 0xde, 0xb8, 0x77, 0xfc, 0x22, 0xbc, 0x2b, 0x45, 0xf8, 0x77,
	0x83, 0xd0, 0x3c, 0x48, 0xe2, 0x4c, 0xe0, 0x21, 0xf0, 0x29, 0x63, 0x44, 0xa6, 0xce, 0xb3, 0x13,
	0x78, 0xe3, 0xdd, 0x04, 0xd8, 0x96, 0xe5, 0x41, 0x04, 0xba, 0x39, 0xd1, 0x98, 0x96, 0x0a, 0x6d,
	0x58, 0xb0, 0x2d, 0xe1, 0x04, 0xf4, 0xad, 0xc8, 0x48, 0xa5, 0x4e, 0x4d, 0x04, 0xb4, 0x19, 0x78,
	0x63, 0xff, 0x78, 0xd0, 0xce, 0xd4, 0xe6, 0x0b, 0x2f, 0xda, 0x7c, 0x49, 0xcf, 0x6a, 0xce, 0x8d,
	0xc4, 0x34, 0x47, 0x2f, 0xc1, 0x96, 0x7b, 0x91, 0x0f, 0xba, 0xf3, 0xf8, 0x7d, 0x3c, 0xbd, 0x8c,
	0xfb, 0xf7, 0x4c, 0x31, 0x3b, 0x8d, 0x27, 0xef, 0xe2, 0xb7, 0x7d, 0xcf, 0x14, 0xc9, 0x3c, 0x8e,
	0x4d, 0xd1, 0x81, 0x3b, 0x60, 0x73, 0x32, 0x8d, 0x4f, 0xfb, 0x1b, 0xa3, 0x1f, 0x1d, 0x80, 0x6e,
	0x46, 0xfc, 0x48, 0x34, 0x36, 0x2b, 0x80, 0x8f, 0xc0, 0x5e, 0xe6, 0xb0, 0x94, 0xe1, 0x8a, 0xa0,
	0xae, 0x9d, 0xdd, 0x6f, 0x7a, 0x31, 0xae, 0xc8, 0x3a, 0xa5, 0xae, 0x69, 0x8e, 0x76, 0xfe, 0xa0,
	0xcc, 0x6b, 0x9a, 0xc3, 0x33, 0xb0, 0xad, 0xec, 0xd2, 0xd0, 0xae, 0x0d, 0xf6, 0xec, 0x5f, 0x97,
	0x9d, 0x34, 0x7a, 0x78, 0x09, 0x7a, 0xee, 0x94, 0x7e, 0xa1, 0x4a, 0x73, 0xb9, 0x40, 0x20, 0xd8,
	0xf8, 0x2f, 0xc7, 0x7d, 0xe7, 0x73, 0xe6, 0x6c, 0xe0, 0x13, 0xd0, 0x5b, 0xdd, 0xbd, 0x54, 0x2f,
	0x04, 0x41, 0xbe, 0xcd, 0xb1, 0xbf, 0xea, 0x5e, 0x2c, 0x04, 0x81, 0x01, 0xf0, 0x73, 0xa2, 0x32,
	0x49, 0x85, 0x69, 0xa1, 0x3d, 0x97, 0x75, 0xad, 0x75, 0xf2, 0xf5, 0xfb, 0x72, 0xe8, 0xfd, 0x5c,
	0x0e, 0xbd, 0x5f, 0xcb, 0xa1, 0x07, 0x0e, 0x32, 0x5e, 0xdd, 0x39, 0xda, 0xc9, 0xfd, 0xd5, 0x50,
	0x6a, 0x66, 0x3e, 0xf1, 0xcc, 0xfb, 0xf4, 0xba, 0x21, 0x17, 0xbc, 0xc4, 0xac, 0x08, 0xb9, 0x2c,
	0xa2, 0x82, 0x30, 0x7b, 0x01, 0x22, 0x07, 0x61, 0x41, 0xd5, 0xed, 0xdf, 0xf2, 0x55, 0x7b, 0xbe,
	0xda, 0xb6, 0xe4, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x28, 0x71, 0xe0, 0xc2, 0x03,
	0x00, 0x00,
}
