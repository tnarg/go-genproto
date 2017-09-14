// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/cloud/audit/audit_log.proto

/*
	Package audit is a generated protocol buffer package.

	It is generated from these files:
		google/cloud/audit/audit_log.proto

	It has these top-level messages:
		AuditLog
		AuthenticationInfo
		AuthorizationInfo
		RequestMetadata
*/
package audit

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

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

// Common audit log format for Google Cloud Platform API operations.
type AuditLog struct {
	// The name of the API service performing the operation. For example,
	// `"datastore.googleapis.com"`.
	ServiceName string `protobuf:"bytes,7,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The name of the service method or operation.
	// For API calls, this should be the name of the API method.
	// For example,
	//
	//     "google.datastore.v1.Datastore.RunQuery"
	//     "google.logging.v1.LoggingService.DeleteLog"
	MethodName string `protobuf:"bytes,8,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// The resource or collection that is the target of the operation.
	// The name is a scheme-less URI, not including the API service name.
	// For example:
	//
	//     "shelves/SHELF_ID/books"
	//     "shelves/SHELF_ID/books/BOOK_ID"
	ResourceName string `protobuf:"bytes,11,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The number of items returned from a List or Query API method,
	// if applicable.
	NumResponseItems int64 `protobuf:"varint,12,opt,name=num_response_items,json=numResponseItems,proto3" json:"num_response_items,omitempty"`
	// The status of the overall operation.
	Status *google_rpc.Status `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	// Authentication information.
	AuthenticationInfo *AuthenticationInfo `protobuf:"bytes,3,opt,name=authentication_info,json=authenticationInfo" json:"authentication_info,omitempty"`
	// Authorization information. If there are multiple
	// resources or permissions involved, then there is
	// one AuthorizationInfo element for each {resource, permission} tuple.
	AuthorizationInfo []*AuthorizationInfo `protobuf:"bytes,9,rep,name=authorization_info,json=authorizationInfo" json:"authorization_info,omitempty"`
	// Metadata about the operation.
	RequestMetadata *RequestMetadata `protobuf:"bytes,4,opt,name=request_metadata,json=requestMetadata" json:"request_metadata,omitempty"`
	// The operation request. This may not include all request parameters,
	// such as those that are too large, privacy-sensitive, or duplicated
	// elsewhere in the log record.
	// It should never include user-generated data, such as file contents.
	// When the JSON object represented here has a proto equivalent, the proto
	// name will be indicated in the `@type` property.
	Request *google_protobuf2.Struct `protobuf:"bytes,16,opt,name=request" json:"request,omitempty"`
	// The operation response. This may not include all response elements,
	// such as those that are too large, privacy-sensitive, or duplicated
	// elsewhere in the log record.
	// It should never include user-generated data, such as file contents.
	// When the JSON object represented here has a proto equivalent, the proto
	// name will be indicated in the `@type` property.
	Response *google_protobuf2.Struct `protobuf:"bytes,17,opt,name=response" json:"response,omitempty"`
	// Other service-specific data about the request, response, and other
	// activities.
	ServiceData *google_protobuf1.Any `protobuf:"bytes,15,opt,name=service_data,json=serviceData" json:"service_data,omitempty"`
}

func (m *AuditLog) Reset()                    { *m = AuditLog{} }
func (m *AuditLog) String() string            { return proto.CompactTextString(m) }
func (*AuditLog) ProtoMessage()               {}
func (*AuditLog) Descriptor() ([]byte, []int) { return fileDescriptorAuditLog, []int{0} }

func (m *AuditLog) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AuditLog) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *AuditLog) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AuditLog) GetNumResponseItems() int64 {
	if m != nil {
		return m.NumResponseItems
	}
	return 0
}

func (m *AuditLog) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AuditLog) GetAuthenticationInfo() *AuthenticationInfo {
	if m != nil {
		return m.AuthenticationInfo
	}
	return nil
}

func (m *AuditLog) GetAuthorizationInfo() []*AuthorizationInfo {
	if m != nil {
		return m.AuthorizationInfo
	}
	return nil
}

func (m *AuditLog) GetRequestMetadata() *RequestMetadata {
	if m != nil {
		return m.RequestMetadata
	}
	return nil
}

func (m *AuditLog) GetRequest() *google_protobuf2.Struct {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AuditLog) GetResponse() *google_protobuf2.Struct {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *AuditLog) GetServiceData() *google_protobuf1.Any {
	if m != nil {
		return m.ServiceData
	}
	return nil
}

// Authentication information for the operation.
type AuthenticationInfo struct {
	// The email address of the authenticated user making the request.
	PrincipalEmail string `protobuf:"bytes,1,opt,name=principal_email,json=principalEmail,proto3" json:"principal_email,omitempty"`
}

func (m *AuthenticationInfo) Reset()                    { *m = AuthenticationInfo{} }
func (m *AuthenticationInfo) String() string            { return proto.CompactTextString(m) }
func (*AuthenticationInfo) ProtoMessage()               {}
func (*AuthenticationInfo) Descriptor() ([]byte, []int) { return fileDescriptorAuditLog, []int{1} }

func (m *AuthenticationInfo) GetPrincipalEmail() string {
	if m != nil {
		return m.PrincipalEmail
	}
	return ""
}

// Authorization information for the operation.
type AuthorizationInfo struct {
	// The resource being accessed, as a REST-style string. For example:
	//
	//     bigquery.googlapis.com/projects/PROJECTID/datasets/DATASETID
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The required IAM permission.
	Permission string `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	// Whether or not authorization for `resource` and `permission`
	// was granted.
	Granted bool `protobuf:"varint,3,opt,name=granted,proto3" json:"granted,omitempty"`
}

func (m *AuthorizationInfo) Reset()                    { *m = AuthorizationInfo{} }
func (m *AuthorizationInfo) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationInfo) ProtoMessage()               {}
func (*AuthorizationInfo) Descriptor() ([]byte, []int) { return fileDescriptorAuditLog, []int{2} }

func (m *AuthorizationInfo) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *AuthorizationInfo) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

func (m *AuthorizationInfo) GetGranted() bool {
	if m != nil {
		return m.Granted
	}
	return false
}

// Metadata about the request.
type RequestMetadata struct {
	// The IP address of the caller.
	CallerIp string `protobuf:"bytes,1,opt,name=caller_ip,json=callerIp,proto3" json:"caller_ip,omitempty"`
	// The user agent of the caller.
	// This information is not authenticated and should be treated accordingly.
	// For example:
	//
	// +   `google-api-python-client/1.4.0`:
	//     The request was made by the Google API client for Python.
	// +   `Cloud SDK Command Line Tool apitools-client/1.0 gcloud/0.9.62`:
	//     The request was made by the Google Cloud SDK CLI (gcloud).
	// +   `AppEngine-Google; (+http://code.google.com/appengine; appid: s~my-project`:
	//     The request was made from the `my-project` App Engine app.
	CallerSuppliedUserAgent string `protobuf:"bytes,2,opt,name=caller_supplied_user_agent,json=callerSuppliedUserAgent,proto3" json:"caller_supplied_user_agent,omitempty"`
}

func (m *RequestMetadata) Reset()                    { *m = RequestMetadata{} }
func (m *RequestMetadata) String() string            { return proto.CompactTextString(m) }
func (*RequestMetadata) ProtoMessage()               {}
func (*RequestMetadata) Descriptor() ([]byte, []int) { return fileDescriptorAuditLog, []int{3} }

func (m *RequestMetadata) GetCallerIp() string {
	if m != nil {
		return m.CallerIp
	}
	return ""
}

func (m *RequestMetadata) GetCallerSuppliedUserAgent() string {
	if m != nil {
		return m.CallerSuppliedUserAgent
	}
	return ""
}

func init() {
	proto.RegisterType((*AuditLog)(nil), "google.cloud.audit.AuditLog")
	proto.RegisterType((*AuthenticationInfo)(nil), "google.cloud.audit.AuthenticationInfo")
	proto.RegisterType((*AuthorizationInfo)(nil), "google.cloud.audit.AuthorizationInfo")
	proto.RegisterType((*RequestMetadata)(nil), "google.cloud.audit.RequestMetadata")
}
func (m *AuditLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuditLog) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(m.Status.Size()))
		n1, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.AuthenticationInfo != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(m.AuthenticationInfo.Size()))
		n2, err := m.AuthenticationInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.RequestMetadata != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(m.RequestMetadata.Size()))
		n3, err := m.RequestMetadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.ServiceName) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(len(m.ServiceName)))
		i += copy(dAtA[i:], m.ServiceName)
	}
	if len(m.MethodName) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(len(m.MethodName)))
		i += copy(dAtA[i:], m.MethodName)
	}
	if len(m.AuthorizationInfo) > 0 {
		for _, msg := range m.AuthorizationInfo {
			dAtA[i] = 0x4a
			i++
			i = encodeVarintAuditLog(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ResourceName) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(len(m.ResourceName)))
		i += copy(dAtA[i:], m.ResourceName)
	}
	if m.NumResponseItems != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(m.NumResponseItems))
	}
	if m.ServiceData != nil {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(m.ServiceData.Size()))
		n4, err := m.ServiceData.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Request != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(m.Request.Size()))
		n5, err := m.Request.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Response != nil {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(m.Response.Size()))
		n6, err := m.Response.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *AuthenticationInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticationInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PrincipalEmail) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(len(m.PrincipalEmail)))
		i += copy(dAtA[i:], m.PrincipalEmail)
	}
	return i, nil
}

func (m *AuthorizationInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthorizationInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Resource) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(len(m.Resource)))
		i += copy(dAtA[i:], m.Resource)
	}
	if len(m.Permission) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(len(m.Permission)))
		i += copy(dAtA[i:], m.Permission)
	}
	if m.Granted {
		dAtA[i] = 0x18
		i++
		if m.Granted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *RequestMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CallerIp) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(len(m.CallerIp)))
		i += copy(dAtA[i:], m.CallerIp)
	}
	if len(m.CallerSuppliedUserAgent) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuditLog(dAtA, i, uint64(len(m.CallerSuppliedUserAgent)))
		i += copy(dAtA[i:], m.CallerSuppliedUserAgent)
	}
	return i, nil
}

func encodeFixed64AuditLog(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32AuditLog(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAuditLog(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AuditLog) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovAuditLog(uint64(l))
	}
	if m.AuthenticationInfo != nil {
		l = m.AuthenticationInfo.Size()
		n += 1 + l + sovAuditLog(uint64(l))
	}
	if m.RequestMetadata != nil {
		l = m.RequestMetadata.Size()
		n += 1 + l + sovAuditLog(uint64(l))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovAuditLog(uint64(l))
	}
	l = len(m.MethodName)
	if l > 0 {
		n += 1 + l + sovAuditLog(uint64(l))
	}
	if len(m.AuthorizationInfo) > 0 {
		for _, e := range m.AuthorizationInfo {
			l = e.Size()
			n += 1 + l + sovAuditLog(uint64(l))
		}
	}
	l = len(m.ResourceName)
	if l > 0 {
		n += 1 + l + sovAuditLog(uint64(l))
	}
	if m.NumResponseItems != 0 {
		n += 1 + sovAuditLog(uint64(m.NumResponseItems))
	}
	if m.ServiceData != nil {
		l = m.ServiceData.Size()
		n += 1 + l + sovAuditLog(uint64(l))
	}
	if m.Request != nil {
		l = m.Request.Size()
		n += 2 + l + sovAuditLog(uint64(l))
	}
	if m.Response != nil {
		l = m.Response.Size()
		n += 2 + l + sovAuditLog(uint64(l))
	}
	return n
}

func (m *AuthenticationInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.PrincipalEmail)
	if l > 0 {
		n += 1 + l + sovAuditLog(uint64(l))
	}
	return n
}

func (m *AuthorizationInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Resource)
	if l > 0 {
		n += 1 + l + sovAuditLog(uint64(l))
	}
	l = len(m.Permission)
	if l > 0 {
		n += 1 + l + sovAuditLog(uint64(l))
	}
	if m.Granted {
		n += 2
	}
	return n
}

func (m *RequestMetadata) Size() (n int) {
	var l int
	_ = l
	l = len(m.CallerIp)
	if l > 0 {
		n += 1 + l + sovAuditLog(uint64(l))
	}
	l = len(m.CallerSuppliedUserAgent)
	if l > 0 {
		n += 1 + l + sovAuditLog(uint64(l))
	}
	return n
}

func sovAuditLog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuditLog(x uint64) (n int) {
	return sovAuditLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuditLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuditLog
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
			return fmt.Errorf("proto: AuditLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuditLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &google_rpc.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticationInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthenticationInfo == nil {
				m.AuthenticationInfo = &AuthenticationInfo{}
			}
			if err := m.AuthenticationInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestMetadata == nil {
				m.RequestMetadata = &RequestMetadata{}
			}
			if err := m.RequestMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizationInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizationInfo = append(m.AuthorizationInfo, &AuthorizationInfo{})
			if err := m.AuthorizationInfo[len(m.AuthorizationInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumResponseItems", wireType)
			}
			m.NumResponseItems = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumResponseItems |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ServiceData == nil {
				m.ServiceData = &google_protobuf1.Any{}
			}
			if err := m.ServiceData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &google_protobuf2.Struct{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &google_protobuf2.Struct{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuditLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuditLog
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
func (m *AuthenticationInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuditLog
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
			return fmt.Errorf("proto: AuthenticationInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticationInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrincipalEmail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrincipalEmail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuditLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuditLog
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
func (m *AuthorizationInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuditLog
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
			return fmt.Errorf("proto: AuthorizationInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthorizationInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permission = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
			m.Granted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAuditLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuditLog
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
func (m *RequestMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuditLog
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
			return fmt.Errorf("proto: RequestMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallerIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallerIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallerSuppliedUserAgent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditLog
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
				return ErrInvalidLengthAuditLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallerSuppliedUserAgent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuditLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuditLog
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
func skipAuditLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuditLog
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
					return 0, ErrIntOverflowAuditLog
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
					return 0, ErrIntOverflowAuditLog
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
				return 0, ErrInvalidLengthAuditLog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuditLog
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
				next, err := skipAuditLog(dAtA[start:])
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
	ErrInvalidLengthAuditLog = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuditLog   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("google/cloud/audit/audit_log.proto", fileDescriptorAuditLog) }

var fileDescriptorAuditLog = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x15, 0x36, 0x6d, 0xad, 0xbb, 0xd1, 0xd6, 0x20, 0x1a, 0xca, 0x54, 0x4a, 0x27, 0xa0,
	0x42, 0x28, 0x11, 0xdb, 0x61, 0x87, 0x89, 0x43, 0x27, 0x38, 0x54, 0x82, 0x69, 0x4a, 0x41, 0x48,
	0x5c, 0x22, 0x2f, 0x71, 0x33, 0x8b, 0xc4, 0x36, 0xfe, 0x83, 0x34, 0x3e, 0x21, 0x47, 0x3e, 0x02,
	0xea, 0x9d, 0xef, 0x80, 0xf2, 0xc6, 0x29, 0x5d, 0x3b, 0xb8, 0x44, 0xf2, 0xf3, 0xfc, 0x5e, 0xfb,
	0x7d, 0x1f, 0xbd, 0x2d, 0x1a, 0x65, 0x42, 0x64, 0x39, 0x0d, 0x93, 0x5c, 0xd8, 0x34, 0x24, 0x36,
	0x65, 0xa6, 0xfa, 0xc6, 0xb9, 0xc8, 0x02, 0xa9, 0x84, 0x11, 0x18, 0x57, 0x4c, 0x00, 0x4c, 0x00,
	0x6e, 0xff, 0xc0, 0xd5, 0x11, 0xc9, 0x42, 0xc2, 0xb9, 0x30, 0xc4, 0x30, 0xc1, 0x75, 0x55, 0xd1,
	0x7f, 0xe8, 0x5c, 0x38, 0x5d, 0xda, 0x79, 0x48, 0xf8, 0xb5, 0xb3, 0x0e, 0xd6, 0x2d, 0x6d, 0x94,
	0x4d, 0x8c, 0x73, 0x7b, 0xce, 0x55, 0x32, 0x09, 0xb5, 0x21, 0xc6, 0xba, 0x1b, 0x47, 0xbf, 0xb7,
	0x51, 0x63, 0x52, 0xbe, 0xfc, 0x4e, 0x64, 0xf8, 0x05, 0xda, 0xa9, 0x4c, 0xff, 0xce, 0xd0, 0x1b,
	0xb7, 0x8e, 0x70, 0xe0, 0x3a, 0x54, 0x32, 0x09, 0x66, 0xe0, 0x44, 0x8e, 0xc0, 0x9f, 0xd0, 0x3d,
	0x62, 0xcd, 0x15, 0xe5, 0x86, 0x25, 0xd0, 0x63, 0xcc, 0xf8, 0x5c, 0xf8, 0x5b, 0x50, 0xf8, 0x2c,
	0xd8, 0x1c, 0x2d, 0x98, 0xdc, 0xc0, 0xa7, 0x7c, 0x2e, 0x22, 0x4c, 0x36, 0x34, 0x7c, 0x8e, 0x3a,
	0x8a, 0x7e, 0xb5, 0x54, 0x9b, 0xb8, 0xa0, 0x86, 0xa4, 0xc4, 0x10, 0x7f, 0x1b, 0x6e, 0x3d, 0xbc,
	0xed, 0xd6, 0xa8, 0x62, 0xdf, 0x3b, 0x34, 0x6a, 0xab, 0x9b, 0x02, 0x7e, 0x82, 0xf6, 0x34, 0x55,
	0xdf, 0x58, 0x42, 0x63, 0x4e, 0x0a, 0xea, 0xef, 0x0e, 0xbd, 0x71, 0x33, 0x6a, 0x39, 0xed, 0x9c,
	0x14, 0x14, 0x3f, 0x46, 0xad, 0x82, 0x9a, 0x2b, 0x91, 0x56, 0x44, 0x03, 0x08, 0x54, 0x49, 0x00,
	0x7c, 0x40, 0xd0, 0xa9, 0x50, 0xec, 0xfb, 0xca, 0xac, 0xcd, 0xe1, 0xd6, 0xb8, 0x75, 0xf4, 0xf4,
	0x5f, 0xb3, 0x2e, 0x69, 0x18, 0xb5, 0x4b, 0xd6, 0x25, 0x7c, 0x88, 0xf6, 0x15, 0xd5, 0xc2, 0xaa,
	0xba, 0xb5, 0x16, 0x3c, 0xbc, 0x57, 0x8b, 0xf0, 0xf4, 0x4b, 0x84, 0xb9, 0x2d, 0x62, 0x45, 0xb5,
	0x14, 0x5c, 0xd3, 0x98, 0x19, 0x5a, 0x68, 0x7f, 0x6f, 0xe8, 0x8d, 0xb7, 0xa2, 0x0e, 0xb7, 0x45,
	0xe4, 0x8c, 0x69, 0xa9, 0xe3, 0x93, 0xbf, 0xc3, 0x42, 0x70, 0x6d, 0x08, 0xee, 0x7e, 0xdd, 0x62,
	0xbd, 0x1c, 0xc1, 0x84, 0x5f, 0x2f, 0x23, 0x78, 0x53, 0xa6, 0xf4, 0x0a, 0xed, 0xba, 0xe0, 0xfc,
	0x0e, 0xd4, 0xf4, 0x36, 0x6a, 0x66, 0xb0, 0x50, 0x51, 0xcd, 0xe1, 0x63, 0xd4, 0xa8, 0xbb, 0xf2,
	0xbb, 0xff, 0xaf, 0x59, 0x82, 0xa3, 0xd7, 0x08, 0x6f, 0xee, 0x01, 0x7e, 0x8e, 0xda, 0x52, 0x31,
	0x9e, 0x30, 0x49, 0xf2, 0x98, 0x16, 0x84, 0xe5, 0xbe, 0x07, 0x59, 0xdc, 0x5d, 0xca, 0x6f, 0x4b,
	0x75, 0xc4, 0x50, 0x77, 0x23, 0x5a, 0xdc, 0x87, 0x46, 0x20, 0x32, 0x57, 0xb6, 0x3c, 0xe3, 0x01,
	0x42, 0x92, 0xaa, 0x82, 0x69, 0xcd, 0x04, 0x87, 0xb5, 0x6e, 0x46, 0x2b, 0x0a, 0xf6, 0xd1, 0x6e,
	0xa6, 0x08, 0x37, 0x34, 0x85, 0xd5, 0x6d, 0x44, 0xf5, 0x71, 0xf4, 0x05, 0xb5, 0xd7, 0x76, 0x0b,
	0x3f, 0x42, 0xcd, 0x84, 0xe4, 0x39, 0x55, 0x31, 0x93, 0xf5, 0x4b, 0x95, 0x30, 0x95, 0xf8, 0x14,
	0xf5, 0x9d, 0xa9, 0xad, 0x94, 0x39, 0xa3, 0x69, 0x6c, 0x35, 0x55, 0x31, 0xc9, 0x28, 0x37, 0xee,
	0xe5, 0x5e, 0x45, 0xcc, 0x1c, 0xf0, 0x51, 0x53, 0x35, 0x29, 0xed, 0x33, 0xfe, 0x63, 0x31, 0xf0,
	0x7e, 0x2e, 0x06, 0xde, 0xaf, 0xc5, 0xc0, 0x43, 0x0f, 0x12, 0x51, 0xdc, 0xb2, 0x55, 0x67, 0xfb,
	0xf5, 0x2f, 0xf5, 0xa2, 0xcc, 0xf7, 0xc2, 0xfb, 0x7c, 0xe2, 0xa0, 0x4c, 0xe4, 0x84, 0x67, 0x81,
	0x50, 0x59, 0x98, 0x51, 0x0e, 0xe9, 0x87, 0x95, 0x45, 0x24, 0xd3, 0xab, 0x7f, 0x42, 0xa7, 0xf0,
	0xbd, 0xdc, 0x01, 0xe6, 0xf8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x29, 0xde, 0xc2, 0xa7,
	0x04, 0x00, 0x00,
}
