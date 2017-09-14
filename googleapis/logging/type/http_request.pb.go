// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/logging/type/http_request.proto

/*
	Package ltype is a generated protocol buffer package.

	It is generated from these files:
		google/logging/type/http_request.proto
		google/logging/type/log_severity.proto

	It has these top-level messages:
		HttpRequest
*/
package ltype

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A common proto for logging HTTP requests. Only contains semantics
// defined by the HTTP specification. Product-specific logging
// information MUST be defined in a separate message.
type HttpRequest struct {
	// The request method. Examples: `"GET"`, `"HEAD"`, `"PUT"`, `"POST"`.
	RequestMethod string `protobuf:"bytes,1,opt,name=request_method,json=requestMethod,proto3" json:"request_method,omitempty"`
	// The scheme (http, https), the host name, the path and the query
	// portion of the URL that was requested.
	// Example: `"http://example.com/some/info?color=red"`.
	RequestUrl string `protobuf:"bytes,2,opt,name=request_url,json=requestUrl,proto3" json:"request_url,omitempty"`
	// The size of the HTTP request message in bytes, including the request
	// headers and the request body.
	RequestSize int64 `protobuf:"varint,3,opt,name=request_size,json=requestSize,proto3" json:"request_size,omitempty"`
	// The response code indicating the status of response.
	// Examples: 200, 404.
	Status int32 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	// The size of the HTTP response message sent back to the client, in bytes,
	// including the response headers and the response body.
	ResponseSize int64 `protobuf:"varint,5,opt,name=response_size,json=responseSize,proto3" json:"response_size,omitempty"`
	// The user agent sent by the client. Example:
	// `"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Q312461; .NET CLR 1.0.3705)"`.
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The IP address (IPv4 or IPv6) of the client that issued the HTTP
	// request. Examples: `"192.168.1.1"`, `"FE80::0202:B3FF:FE1E:8329"`.
	RemoteIp string `protobuf:"bytes,7,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	// The IP address (IPv4 or IPv6) of the origin server that the request was
	// sent to.
	ServerIp string `protobuf:"bytes,13,opt,name=server_ip,json=serverIp,proto3" json:"server_ip,omitempty"`
	// The referer URL of the request, as defined in
	// [HTTP/1.1 Header Field Definitions](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	Referer string `protobuf:"bytes,8,opt,name=referer,proto3" json:"referer,omitempty"`
	// The request processing latency on the server, from the time the request was
	// received until the response was sent.
	Latency *google_protobuf1.Duration `protobuf:"bytes,14,opt,name=latency" json:"latency,omitempty"`
	// Whether or not a cache lookup was attempted.
	CacheLookup bool `protobuf:"varint,11,opt,name=cache_lookup,json=cacheLookup,proto3" json:"cache_lookup,omitempty"`
	// Whether or not an entity was served from cache
	// (with or without validation).
	CacheHit bool `protobuf:"varint,9,opt,name=cache_hit,json=cacheHit,proto3" json:"cache_hit,omitempty"`
	// Whether or not the response was validated with the origin server before
	// being served from cache. This field is only meaningful if `cache_hit` is
	// True.
	CacheValidatedWithOriginServer bool `protobuf:"varint,10,opt,name=cache_validated_with_origin_server,json=cacheValidatedWithOriginServer,proto3" json:"cache_validated_with_origin_server,omitempty"`
	// The number of HTTP response bytes inserted into cache. Set only when a
	// cache fill was attempted.
	CacheFillBytes int64 `protobuf:"varint,12,opt,name=cache_fill_bytes,json=cacheFillBytes,proto3" json:"cache_fill_bytes,omitempty"`
	// Protocol used for the request. Examples: "HTTP/1.1", "HTTP/2", "websocket"
	Protocol string `protobuf:"bytes,15,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (m *HttpRequest) Reset()                    { *m = HttpRequest{} }
func (m *HttpRequest) String() string            { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()               {}
func (*HttpRequest) Descriptor() ([]byte, []int) { return fileDescriptorHttpRequest, []int{0} }

func (m *HttpRequest) GetRequestMethod() string {
	if m != nil {
		return m.RequestMethod
	}
	return ""
}

func (m *HttpRequest) GetRequestUrl() string {
	if m != nil {
		return m.RequestUrl
	}
	return ""
}

func (m *HttpRequest) GetRequestSize() int64 {
	if m != nil {
		return m.RequestSize
	}
	return 0
}

func (m *HttpRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *HttpRequest) GetResponseSize() int64 {
	if m != nil {
		return m.ResponseSize
	}
	return 0
}

func (m *HttpRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *HttpRequest) GetRemoteIp() string {
	if m != nil {
		return m.RemoteIp
	}
	return ""
}

func (m *HttpRequest) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *HttpRequest) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

func (m *HttpRequest) GetLatency() *google_protobuf1.Duration {
	if m != nil {
		return m.Latency
	}
	return nil
}

func (m *HttpRequest) GetCacheLookup() bool {
	if m != nil {
		return m.CacheLookup
	}
	return false
}

func (m *HttpRequest) GetCacheHit() bool {
	if m != nil {
		return m.CacheHit
	}
	return false
}

func (m *HttpRequest) GetCacheValidatedWithOriginServer() bool {
	if m != nil {
		return m.CacheValidatedWithOriginServer
	}
	return false
}

func (m *HttpRequest) GetCacheFillBytes() int64 {
	if m != nil {
		return m.CacheFillBytes
	}
	return 0
}

func (m *HttpRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*HttpRequest)(nil), "google.logging.type.HttpRequest")
}
func (m *HttpRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RequestMethod) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(len(m.RequestMethod)))
		i += copy(dAtA[i:], m.RequestMethod)
	}
	if len(m.RequestUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(len(m.RequestUrl)))
		i += copy(dAtA[i:], m.RequestUrl)
	}
	if m.RequestSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(m.RequestSize))
	}
	if m.Status != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(m.Status))
	}
	if m.ResponseSize != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(m.ResponseSize))
	}
	if len(m.UserAgent) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(len(m.UserAgent)))
		i += copy(dAtA[i:], m.UserAgent)
	}
	if len(m.RemoteIp) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(len(m.RemoteIp)))
		i += copy(dAtA[i:], m.RemoteIp)
	}
	if len(m.Referer) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(len(m.Referer)))
		i += copy(dAtA[i:], m.Referer)
	}
	if m.CacheHit {
		dAtA[i] = 0x48
		i++
		if m.CacheHit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CacheValidatedWithOriginServer {
		dAtA[i] = 0x50
		i++
		if m.CacheValidatedWithOriginServer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CacheLookup {
		dAtA[i] = 0x58
		i++
		if m.CacheLookup {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CacheFillBytes != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(m.CacheFillBytes))
	}
	if len(m.ServerIp) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(len(m.ServerIp)))
		i += copy(dAtA[i:], m.ServerIp)
	}
	if m.Latency != nil {
		dAtA[i] = 0x72
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(m.Latency.Size()))
		n1, err := m.Latency.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Protocol) > 0 {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintHttpRequest(dAtA, i, uint64(len(m.Protocol)))
		i += copy(dAtA[i:], m.Protocol)
	}
	return i, nil
}

func encodeFixed64HttpRequest(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32HttpRequest(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHttpRequest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HttpRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.RequestMethod)
	if l > 0 {
		n += 1 + l + sovHttpRequest(uint64(l))
	}
	l = len(m.RequestUrl)
	if l > 0 {
		n += 1 + l + sovHttpRequest(uint64(l))
	}
	if m.RequestSize != 0 {
		n += 1 + sovHttpRequest(uint64(m.RequestSize))
	}
	if m.Status != 0 {
		n += 1 + sovHttpRequest(uint64(m.Status))
	}
	if m.ResponseSize != 0 {
		n += 1 + sovHttpRequest(uint64(m.ResponseSize))
	}
	l = len(m.UserAgent)
	if l > 0 {
		n += 1 + l + sovHttpRequest(uint64(l))
	}
	l = len(m.RemoteIp)
	if l > 0 {
		n += 1 + l + sovHttpRequest(uint64(l))
	}
	l = len(m.Referer)
	if l > 0 {
		n += 1 + l + sovHttpRequest(uint64(l))
	}
	if m.CacheHit {
		n += 2
	}
	if m.CacheValidatedWithOriginServer {
		n += 2
	}
	if m.CacheLookup {
		n += 2
	}
	if m.CacheFillBytes != 0 {
		n += 1 + sovHttpRequest(uint64(m.CacheFillBytes))
	}
	l = len(m.ServerIp)
	if l > 0 {
		n += 1 + l + sovHttpRequest(uint64(l))
	}
	if m.Latency != nil {
		l = m.Latency.Size()
		n += 1 + l + sovHttpRequest(uint64(l))
	}
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sovHttpRequest(uint64(l))
	}
	return n
}

func sovHttpRequest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHttpRequest(x uint64) (n int) {
	return sovHttpRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpRequest
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
			return fmt.Errorf("proto: HttpRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
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
				return ErrInvalidLengthHttpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
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
				return ErrInvalidLengthHttpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestSize", wireType)
			}
			m.RequestSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseSize", wireType)
			}
			m.ResponseSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResponseSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAgent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
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
				return ErrInvalidLengthHttpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserAgent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
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
				return ErrInvalidLengthHttpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Referer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
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
				return ErrInvalidLengthHttpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Referer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheHit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CacheHit = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheValidatedWithOriginServer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CacheValidatedWithOriginServer = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheLookup", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CacheLookup = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheFillBytes", wireType)
			}
			m.CacheFillBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CacheFillBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
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
				return ErrInvalidLengthHttpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latency", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
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
				return ErrInvalidLengthHttpRequest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Latency == nil {
				m.Latency = &google_protobuf1.Duration{}
			}
			if err := m.Latency.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpRequest
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
				return ErrInvalidLengthHttpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpRequest
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
func skipHttpRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpRequest
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
					return 0, ErrIntOverflowHttpRequest
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
					return 0, ErrIntOverflowHttpRequest
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
				return 0, ErrInvalidLengthHttpRequest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttpRequest
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
				next, err := skipHttpRequest(dAtA[start:])
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
	ErrInvalidLengthHttpRequest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpRequest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("google/logging/type/http_request.proto", fileDescriptorHttpRequest) }

var fileDescriptorHttpRequest = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcb, 0x8e, 0xd3, 0x3e,
	0x14, 0xc6, 0x95, 0xb9, 0xb5, 0x75, 0x2f, 0xff, 0x91, 0xff, 0x12, 0x78, 0x0a, 0x94, 0x32, 0x08,
	0x94, 0x55, 0x22, 0x31, 0x1b, 0x24, 0x56, 0x14, 0x04, 0x33, 0x68, 0x10, 0xa3, 0x0c, 0x17, 0x89,
	0x4d, 0x94, 0xb6, 0xa7, 0x8e, 0x85, 0x1b, 0x1b, 0xdb, 0x19, 0xd4, 0x79, 0x09, 0xde, 0x83, 0x0d,
	0xaf, 0xc1, 0x92, 0x47, 0x40, 0x7d, 0x12, 0x94, 0x63, 0x47, 0x02, 0x89, 0x4d, 0xa4, 0xf3, 0x7d,
	0xbf, 0xef, 0x9c, 0xd8, 0x3a, 0x26, 0x0f, 0xb9, 0x52, 0x5c, 0x42, 0x2a, 0x15, 0xe7, 0xa2, 0xe2,
	0xa9, 0xdb, 0x68, 0x48, 0x4b, 0xe7, 0x74, 0x6e, 0xe0, 0x73, 0x0d, 0xd6, 0x25, 0xda, 0x28, 0xa7,
	0xe8, 0xff, 0x9e, 0x4b, 0x02, 0x97, 0x34, 0xdc, 0xf8, 0x76, 0x08, 0x17, 0x5a, 0xa4, 0x45, 0x55,
	0x29, 0x57, 0x38, 0xa1, 0x2a, 0xeb, 0x23, 0xe3, 0x49, 0x70, 0xb1, 0x9a, 0xd7, 0xab, 0x74, 0x59,
	0x1b, 0x04, 0xbc, 0x7f, 0xfc, 0x7d, 0x8f, 0xf4, 0x4f, 0x9d, 0xd3, 0x99, 0x1f, 0x44, 0x1f, 0x90,
	0x51, 0x98, 0x99, 0xaf, 0xc1, 0x95, 0x6a, 0xc9, 0xa2, 0x69, 0x14, 0xf7, 0xb2, 0x61, 0x50, 0x5f,
	0xa3, 0x48, 0xef, 0x92, 0x7e, 0x8b, 0xd5, 0x46, 0xb2, 0x1d, 0x64, 0x48, 0x90, 0xde, 0x19, 0x49,
	0xef, 0x91, 0x41, 0x0b, 0x58, 0x71, 0x0d, 0x6c, 0x77, 0x1a, 0xc5, 0xbb, 0x59, 0x1b, 0xba, 0x14,
	0xd7, 0x40, 0x6f, 0x90, 0x03, 0xeb, 0x0a, 0x57, 0x5b, 0xb6, 0x37, 0x8d, 0xe2, 0xfd, 0x2c, 0x54,
	0xf4, 0x3e, 0x19, 0x1a, 0xb0, 0x5a, 0x55, 0x16, 0x7c, 0x76, 0x1f, 0xb3, 0x83, 0x56, 0xc4, 0xf0,
	0x1d, 0x42, 0x6a, 0x0b, 0x26, 0x2f, 0x38, 0x54, 0x8e, 0x1d, 0xe0, 0xfc, 0x5e, 0xa3, 0x3c, 0x6d,
	0x04, 0x7a, 0x8b, 0xf4, 0x0c, 0xac, 0x95, 0x83, 0x5c, 0x68, 0xd6, 0x41, 0xb7, 0xeb, 0x85, 0x33,
	0x4d, 0x19, 0xe9, 0x18, 0x58, 0x81, 0x01, 0xc3, 0xba, 0x68, 0xb5, 0x65, 0x13, 0x5b, 0x14, 0x8b,
	0x12, 0xf2, 0x52, 0x38, 0xd6, 0x9b, 0x46, 0x71, 0x37, 0xeb, 0xa2, 0x70, 0x2a, 0x1c, 0x7d, 0x45,
	0x8e, 0xbd, 0x79, 0x55, 0x48, 0xb1, 0x2c, 0x1c, 0x2c, 0xf3, 0x2f, 0xc2, 0x95, 0xb9, 0x32, 0x82,
	0x8b, 0x2a, 0xb7, 0x60, 0xae, 0xc0, 0x30, 0x82, 0xa9, 0x09, 0x92, 0xef, 0x5b, 0xf0, 0x83, 0x70,
	0xe5, 0x1b, 0xc4, 0x2e, 0x91, 0x6a, 0xae, 0xc7, 0xf7, 0x92, 0x4a, 0x7d, 0xaa, 0x35, 0xeb, 0x63,
	0xaa, 0x8f, 0xda, 0x39, 0x4a, 0x34, 0x26, 0x87, 0x1e, 0x59, 0x09, 0x29, 0xf3, 0xf9, 0xc6, 0x81,
	0x65, 0x03, 0xbc, 0x89, 0x11, 0xea, 0x2f, 0x84, 0x94, 0xb3, 0x46, 0x6d, 0xfe, 0xda, 0x0f, 0x6f,
	0x0e, 0x3b, 0xf4, 0x87, 0xf5, 0xc2, 0x99, 0xa6, 0x27, 0xa4, 0x23, 0x0b, 0x07, 0xd5, 0x62, 0xc3,
	0x46, 0xd3, 0x28, 0xee, 0x3f, 0x3a, 0x4a, 0xc2, 0x16, 0xb5, 0x2b, 0x91, 0x3c, 0x0f, 0x2b, 0x91,
	0xb5, 0x24, 0x1d, 0x93, 0x2e, 0xba, 0x0b, 0x25, 0xd9, 0x7f, 0xbe, 0x61, 0x5b, 0xcf, 0xbe, 0x46,
	0x3f, 0xb6, 0x93, 0xe8, 0xe7, 0x76, 0x12, 0xfd, 0xda, 0x4e, 0x22, 0x72, 0x73, 0xa1, 0xd6, 0xc9,
	0x3f, 0xf6, 0x72, 0x76, 0xf8, 0xc7, 0x5a, 0x5d, 0x34, 0xe1, 0x8b, 0xe8, 0xe3, 0xe3, 0x00, 0x72,
	0x25, 0x8b, 0x8a, 0x27, 0xca, 0xf0, 0x94, 0x43, 0x85, 0xad, 0x53, 0x6f, 0x15, 0x5a, 0xd8, 0xbf,
	0xde, 0xc1, 0x13, 0xd9, 0x7c, 0xbf, 0xed, 0x1c, 0xbd, 0xf4, 0xd1, 0x67, 0x52, 0xd5, 0xcb, 0xe4,
	0x3c, 0x4c, 0x7a, 0xbb, 0xd1, 0x30, 0x3f, 0xc0, 0x06, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x91, 0x3a, 0x8c, 0x5f, 0x47, 0x03, 0x00, 0x00,
}
