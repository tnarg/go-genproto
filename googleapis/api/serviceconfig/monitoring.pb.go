// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/monitoring.proto

package serviceconfig

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Monitoring configuration of the service.
//
// The example below shows how to configure monitored resources and metrics
// for monitoring. In the example, a monitored resource and two metrics are
// defined. The `library.googleapis.com/book/returned_count` metric is sent
// to both producer and consumer projects, whereas the
// `library.googleapis.com/book/overdue_count` metric is only sent to the
// consumer project.
//
//     monitored_resources:
//     - type: library.googleapis.com/branch
//       labels:
//       - key: /city
//         description: The city where the library branch is located in.
//       - key: /name
//         description: The name of the branch.
//     metrics:
//     - name: library.googleapis.com/book/returned_count
//       metric_kind: DELTA
//       value_type: INT64
//       labels:
//       - key: /customer_id
//     - name: library.googleapis.com/book/overdue_count
//       metric_kind: GAUGE
//       value_type: INT64
//       labels:
//       - key: /customer_id
//     monitoring:
//       producer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/returned_count
//       consumer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/returned_count
//         - library.googleapis.com/book/overdue_count
type Monitoring struct {
	// Monitoring configurations for sending metrics to the producer project.
	// There can be multiple producer destinations, each one must have a
	// different monitored resource type. A metric can be used in at most
	// one producer destination.
	ProducerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,1,rep,name=producer_destinations,json=producerDestinations" json:"producer_destinations,omitempty"`
	// Monitoring configurations for sending metrics to the consumer project.
	// There can be multiple consumer destinations, each one must have a
	// different monitored resource type. A metric can be used in at most
	// one consumer destination.
	ConsumerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,2,rep,name=consumer_destinations,json=consumerDestinations" json:"consumer_destinations,omitempty"`
}

func (m *Monitoring) Reset()                    { *m = Monitoring{} }
func (m *Monitoring) String() string            { return proto.CompactTextString(m) }
func (*Monitoring) ProtoMessage()               {}
func (*Monitoring) Descriptor() ([]byte, []int) { return fileDescriptorMonitoring, []int{0} }

func (m *Monitoring) GetProducerDestinations() []*Monitoring_MonitoringDestination {
	if m != nil {
		return m.ProducerDestinations
	}
	return nil
}

func (m *Monitoring) GetConsumerDestinations() []*Monitoring_MonitoringDestination {
	if m != nil {
		return m.ConsumerDestinations
	}
	return nil
}

// Configuration of a specific monitoring destination (the producer project
// or the consumer project).
type Monitoring_MonitoringDestination struct {
	// The monitored resource type. The type must be defined in
	// [Service.monitored_resources][google.api.Service.monitored_resources] section.
	MonitoredResource string `protobuf:"bytes,1,opt,name=monitored_resource,json=monitoredResource,proto3" json:"monitored_resource,omitempty"`
	// Names of the metrics to report to this monitoring destination.
	// Each name must be defined in [Service.metrics][google.api.Service.metrics] section.
	Metrics []string `protobuf:"bytes,2,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *Monitoring_MonitoringDestination) Reset()         { *m = Monitoring_MonitoringDestination{} }
func (m *Monitoring_MonitoringDestination) String() string { return proto.CompactTextString(m) }
func (*Monitoring_MonitoringDestination) ProtoMessage()    {}
func (*Monitoring_MonitoringDestination) Descriptor() ([]byte, []int) {
	return fileDescriptorMonitoring, []int{0, 0}
}

func (m *Monitoring_MonitoringDestination) GetMonitoredResource() string {
	if m != nil {
		return m.MonitoredResource
	}
	return ""
}

func (m *Monitoring_MonitoringDestination) GetMetrics() []string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*Monitoring)(nil), "google.api.Monitoring")
	proto.RegisterType((*Monitoring_MonitoringDestination)(nil), "google.api.Monitoring.MonitoringDestination")
}
func (m *Monitoring) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Monitoring) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProducerDestinations) > 0 {
		for _, msg := range m.ProducerDestinations {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMonitoring(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ConsumerDestinations) > 0 {
		for _, msg := range m.ConsumerDestinations {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMonitoring(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Monitoring_MonitoringDestination) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Monitoring_MonitoringDestination) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MonitoredResource) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMonitoring(dAtA, i, uint64(len(m.MonitoredResource)))
		i += copy(dAtA[i:], m.MonitoredResource)
	}
	if len(m.Metrics) > 0 {
		for _, s := range m.Metrics {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Monitoring(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Monitoring(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMonitoring(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Monitoring) Size() (n int) {
	var l int
	_ = l
	if len(m.ProducerDestinations) > 0 {
		for _, e := range m.ProducerDestinations {
			l = e.Size()
			n += 1 + l + sovMonitoring(uint64(l))
		}
	}
	if len(m.ConsumerDestinations) > 0 {
		for _, e := range m.ConsumerDestinations {
			l = e.Size()
			n += 1 + l + sovMonitoring(uint64(l))
		}
	}
	return n
}

func (m *Monitoring_MonitoringDestination) Size() (n int) {
	var l int
	_ = l
	l = len(m.MonitoredResource)
	if l > 0 {
		n += 1 + l + sovMonitoring(uint64(l))
	}
	if len(m.Metrics) > 0 {
		for _, s := range m.Metrics {
			l = len(s)
			n += 1 + l + sovMonitoring(uint64(l))
		}
	}
	return n
}

func sovMonitoring(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMonitoring(x uint64) (n int) {
	return sovMonitoring(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Monitoring) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoring
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
			return fmt.Errorf("proto: Monitoring: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Monitoring: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducerDestinations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
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
				return ErrInvalidLengthMonitoring
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProducerDestinations = append(m.ProducerDestinations, &Monitoring_MonitoringDestination{})
			if err := m.ProducerDestinations[len(m.ProducerDestinations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsumerDestinations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
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
				return ErrInvalidLengthMonitoring
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsumerDestinations = append(m.ConsumerDestinations, &Monitoring_MonitoringDestination{})
			if err := m.ConsumerDestinations[len(m.ConsumerDestinations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitoring
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
func (m *Monitoring_MonitoringDestination) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoring
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
			return fmt.Errorf("proto: MonitoringDestination: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoringDestination: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonitoredResource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
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
				return ErrInvalidLengthMonitoring
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MonitoredResource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
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
				return ErrInvalidLengthMonitoring
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metrics = append(m.Metrics, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitoring
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
func skipMonitoring(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMonitoring
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
					return 0, ErrIntOverflowMonitoring
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
					return 0, ErrIntOverflowMonitoring
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
				return 0, ErrInvalidLengthMonitoring
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMonitoring
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
				next, err := skipMonitoring(dAtA[start:])
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
	ErrInvalidLengthMonitoring = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMonitoring   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("google/api/monitoring.proto", fileDescriptorMonitoring) }

var fileDescriptorMonitoring = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0xc6, 0xd9, 0x28, 0xca, 0xad, 0xa0, 0x18, 0x3c, 0x08, 0xa7, 0x84, 0xc3, 0xea, 0x0a, 0x4d,
	0x40, 0x4b, 0x2b, 0x0f, 0x45, 0x2c, 0x84, 0x90, 0xd2, 0xe6, 0x5c, 0x37, 0xeb, 0x32, 0x70, 0x99,
	0x09, 0xbb, 0x1b, 0x3b, 0x5f, 0xc6, 0x67, 0xf0, 0x21, 0x2c, 0x7d, 0x04, 0xc9, 0x93, 0xc8, 0xe5,
	0xcf, 0x25, 0x8a, 0xd5, 0x75, 0x3b, 0xfb, 0x7d, 0xdf, 0xfc, 0x86, 0x19, 0x7e, 0xac, 0x89, 0xf4,
	0x52, 0xc5, 0xa2, 0x80, 0x38, 0x27, 0x04, 0x47, 0x06, 0x50, 0x47, 0x85, 0x21, 0x47, 0x3e, 0x6f,
	0xc4, 0x48, 0x14, 0x30, 0x39, 0x19, 0x18, 0x05, 0x22, 0x39, 0xe1, 0x80, 0xd0, 0x36, 0xce, 0xd3,
	0x0f, 0x8f, 0xf3, 0x87, 0x75, 0xdc, 0x17, 0x7c, 0x5c, 0x18, 0xca, 0x4a, 0xa9, 0xcc, 0x22, 0x53,
	0xd6, 0x01, 0x36, 0xee, 0x80, 0x4d, 0xb7, 0x66, 0x7b, 0x17, 0x67, 0x51, 0xdf, 0x38, 0xea, 0x63,
	0x83, 0xe7, 0x4d, 0x1f, 0x4a, 0x8f, 0xba, 0x56, 0x83, 0x4f, 0xbb, 0x42, 0x48, 0x42, 0x5b, 0xe6,
	0x7f, 0x11, 0xde, 0x26, 0x88, 0xae, 0xd5, 0x10, 0x31, 0x79, 0xe2, 0xe3, 0x7f, 0xed, 0xfe, 0x39,
	0xf7, 0xdb, 0x5d, 0xa9, 0x6c, 0x61, 0x94, 0xa5, 0xd2, 0x48, 0x15, 0xb0, 0x29, 0x9b, 0x8d, 0xd2,
	0xc3, 0xb5, 0x92, 0xb6, 0x82, 0x1f, 0xf0, 0xdd, 0x5c, 0x39, 0x03, 0xb2, 0x19, 0x6e, 0x94, 0x76,
	0xe5, 0xfc, 0xed, 0xb3, 0x0a, 0xd9, 0x57, 0x15, 0xb2, 0xef, 0x2a, 0x64, 0x7c, 0x5f, 0x52, 0x3e,
	0x18, 0x7b, 0x7e, 0xd0, 0xd3, 0x93, 0xd5, 0x96, 0x13, 0xf6, 0x78, 0xdb, 0xca, 0x9a, 0x96, 0x02,
	0x75, 0x44, 0x46, 0xc7, 0x5a, 0x61, 0x7d, 0x83, 0xb8, 0x91, 0x44, 0x01, 0xb6, 0x3e, 0x92, 0x55,
	0xe6, 0x15, 0xa4, 0x92, 0x84, 0x2f, 0xa0, 0xaf, 0x7e, 0x55, 0xef, 0xde, 0xf6, 0xdd, 0x75, 0x72,
	0xff, 0xbc, 0x53, 0x07, 0x2f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0x67, 0x17, 0x9d, 0x05,
	0x02, 0x00, 0x00,
}
