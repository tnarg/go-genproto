// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/appengine/v1/deploy.proto

package appengine

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Code and application artifacts used to deploy a version to App Engine.
type Deployment struct {
	// Manifest of the files stored in Google Cloud Storage that are included
	// as part of this version. All files must be readable using the
	// credentials supplied with this call.
	Files map[string]*FileInfo `protobuf:"bytes,1,rep,name=files" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// A Docker image that App Engine uses to run the version.
	// Only applicable for instances in App Engine flexible environment.
	Container *ContainerInfo `protobuf:"bytes,2,opt,name=container" json:"container,omitempty"`
	// The zip file for this deployment, if this is a zip deployment.
	Zip *ZipInfo `protobuf:"bytes,3,opt,name=zip" json:"zip,omitempty"`
}

func (m *Deployment) Reset()                    { *m = Deployment{} }
func (m *Deployment) String() string            { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()               {}
func (*Deployment) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{0} }

func (m *Deployment) GetFiles() map[string]*FileInfo {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Deployment) GetContainer() *ContainerInfo {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *Deployment) GetZip() *ZipInfo {
	if m != nil {
		return m.Zip
	}
	return nil
}

// Single source file that is part of the version to be deployed. Each source
// file that is deployed must be specified separately.
type FileInfo struct {
	// URL source to use to fetch this file. Must be a URL to a resource in
	// Google Cloud Storage in the form
	// 'http(s)://storage.googleapis.com/\<bucket\>/\<object\>'.
	SourceUrl string `protobuf:"bytes,1,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// The SHA1 hash of the file, in hex.
	Sha1Sum string `protobuf:"bytes,2,opt,name=sha1_sum,json=sha1Sum,proto3" json:"sha1_sum,omitempty"`
	// The MIME type of the file.
	//
	// Defaults to the value from Google Cloud Storage.
	MimeType string `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (m *FileInfo) Reset()                    { *m = FileInfo{} }
func (m *FileInfo) String() string            { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()               {}
func (*FileInfo) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{1} }

func (m *FileInfo) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

func (m *FileInfo) GetSha1Sum() string {
	if m != nil {
		return m.Sha1Sum
	}
	return ""
}

func (m *FileInfo) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

// Docker image that is used to start a VM container for the version you
// deploy.
type ContainerInfo struct {
	// URI to the hosted container image in a Docker repository. The URI must be
	// fully qualified and include a tag or digest.
	// Examples: "gcr.io/my-project/image:tag" or "gcr.io/my-project/image@digest"
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *ContainerInfo) Reset()                    { *m = ContainerInfo{} }
func (m *ContainerInfo) String() string            { return proto.CompactTextString(m) }
func (*ContainerInfo) ProtoMessage()               {}
func (*ContainerInfo) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{2} }

func (m *ContainerInfo) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type ZipInfo struct {
	// URL of the zip file to deploy from. Must be a URL to a resource in
	// Google Cloud Storage in the form
	// 'http(s)://storage.googleapis.com/\<bucket\>/\<object\>'.
	SourceUrl string `protobuf:"bytes,3,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// An estimate of the number of files in a zip for a zip deployment.
	// If set, must be greater than or equal to the actual number of files.
	// Used for optimizing performance; if not provided, deployment may be slow.
	FilesCount int32 `protobuf:"varint,4,opt,name=files_count,json=filesCount,proto3" json:"files_count,omitempty"`
}

func (m *ZipInfo) Reset()                    { *m = ZipInfo{} }
func (m *ZipInfo) String() string            { return proto.CompactTextString(m) }
func (*ZipInfo) ProtoMessage()               {}
func (*ZipInfo) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{3} }

func (m *ZipInfo) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

func (m *ZipInfo) GetFilesCount() int32 {
	if m != nil {
		return m.FilesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Deployment)(nil), "google.appengine.v1.Deployment")
	proto.RegisterType((*FileInfo)(nil), "google.appengine.v1.FileInfo")
	proto.RegisterType((*ContainerInfo)(nil), "google.appengine.v1.ContainerInfo")
	proto.RegisterType((*ZipInfo)(nil), "google.appengine.v1.ZipInfo")
}
func (m *Deployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deployment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Files) > 0 {
		for k, _ := range m.Files {
			dAtA[i] = 0xa
			i++
			v := m.Files[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovDeploy(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovDeploy(uint64(len(k))) + msgSize
			i = encodeVarintDeploy(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeploy(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintDeploy(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	if m.Container != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDeploy(dAtA, i, uint64(m.Container.Size()))
		n2, err := m.Container.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Zip != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDeploy(dAtA, i, uint64(m.Zip.Size()))
		n3, err := m.Zip.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *FileInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SourceUrl) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeploy(dAtA, i, uint64(len(m.SourceUrl)))
		i += copy(dAtA[i:], m.SourceUrl)
	}
	if len(m.Sha1Sum) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDeploy(dAtA, i, uint64(len(m.Sha1Sum)))
		i += copy(dAtA[i:], m.Sha1Sum)
	}
	if len(m.MimeType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDeploy(dAtA, i, uint64(len(m.MimeType)))
		i += copy(dAtA[i:], m.MimeType)
	}
	return i, nil
}

func (m *ContainerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Image) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeploy(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	return i, nil
}

func (m *ZipInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ZipInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SourceUrl) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDeploy(dAtA, i, uint64(len(m.SourceUrl)))
		i += copy(dAtA[i:], m.SourceUrl)
	}
	if m.FilesCount != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDeploy(dAtA, i, uint64(m.FilesCount))
	}
	return i, nil
}

func encodeFixed64Deploy(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Deploy(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDeploy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Deployment) Size() (n int) {
	var l int
	_ = l
	if len(m.Files) > 0 {
		for k, v := range m.Files {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovDeploy(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovDeploy(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovDeploy(uint64(mapEntrySize))
		}
	}
	if m.Container != nil {
		l = m.Container.Size()
		n += 1 + l + sovDeploy(uint64(l))
	}
	if m.Zip != nil {
		l = m.Zip.Size()
		n += 1 + l + sovDeploy(uint64(l))
	}
	return n
}

func (m *FileInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.SourceUrl)
	if l > 0 {
		n += 1 + l + sovDeploy(uint64(l))
	}
	l = len(m.Sha1Sum)
	if l > 0 {
		n += 1 + l + sovDeploy(uint64(l))
	}
	l = len(m.MimeType)
	if l > 0 {
		n += 1 + l + sovDeploy(uint64(l))
	}
	return n
}

func (m *ContainerInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovDeploy(uint64(l))
	}
	return n
}

func (m *ZipInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.SourceUrl)
	if l > 0 {
		n += 1 + l + sovDeploy(uint64(l))
	}
	if m.FilesCount != 0 {
		n += 1 + sovDeploy(uint64(m.FilesCount))
	}
	return n
}

func sovDeploy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeploy(x uint64) (n int) {
	return sovDeploy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Deployment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploy
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
			return fmt.Errorf("proto: Deployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Files", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploy
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
				return ErrInvalidLengthDeploy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Files == nil {
				m.Files = make(map[string]*FileInfo)
			}
			var mapkey string
			var mapvalue *FileInfo
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDeploy
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
							return ErrIntOverflowDeploy
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
						return ErrInvalidLengthDeploy
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
							return ErrIntOverflowDeploy
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
						return ErrInvalidLengthDeploy
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthDeploy
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &FileInfo{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDeploy(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthDeploy
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Files[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Container", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploy
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
				return ErrInvalidLengthDeploy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Container == nil {
				m.Container = &ContainerInfo{}
			}
			if err := m.Container.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zip", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploy
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
				return ErrInvalidLengthDeploy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Zip == nil {
				m.Zip = &ZipInfo{}
			}
			if err := m.Zip.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeploy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeploy
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
func (m *FileInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploy
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
			return fmt.Errorf("proto: FileInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploy
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
				return ErrInvalidLengthDeploy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha1Sum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploy
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
				return ErrInvalidLengthDeploy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha1Sum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MimeType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploy
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
				return ErrInvalidLengthDeploy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MimeType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeploy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeploy
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
func (m *ContainerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploy
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
			return fmt.Errorf("proto: ContainerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploy
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
				return ErrInvalidLengthDeploy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeploy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeploy
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
func (m *ZipInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeploy
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
			return fmt.Errorf("proto: ZipInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ZipInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploy
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
				return ErrInvalidLengthDeploy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilesCount", wireType)
			}
			m.FilesCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeploy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FilesCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeploy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeploy
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
func skipDeploy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeploy
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
					return 0, ErrIntOverflowDeploy
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
					return 0, ErrIntOverflowDeploy
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
				return 0, ErrInvalidLengthDeploy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeploy
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
				next, err := skipDeploy(dAtA[start:])
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
	ErrInvalidLengthDeploy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeploy   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("google/appengine/v1/deploy.proto", fileDescriptorDeploy) }

var fileDescriptorDeploy = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0x86, 0xc9, 0xd4, 0x71, 0xa7, 0x67, 0x10, 0x24, 0x0a, 0xd6, 0x75, 0x77, 0x2c, 0x05, 0x61,
	0xf0, 0x22, 0x65, 0x76, 0x6f, 0x44, 0xbd, 0x58, 0x76, 0x55, 0xd8, 0xbb, 0xa5, 0x2a, 0xc2, 0xde,
	0x94, 0x58, 0xb3, 0x31, 0xd8, 0x26, 0x21, 0x4d, 0x07, 0xea, 0x9b, 0xf8, 0x46, 0x5e, 0xfa, 0x08,
	0x32, 0x4f, 0xb2, 0x24, 0xe9, 0xce, 0xb0, 0x43, 0xef, 0x7a, 0xfe, 0x7e, 0xff, 0x9f, 0x93, 0x9c,
	0x03, 0x29, 0x57, 0x8a, 0xd7, 0x2c, 0xa7, 0x5a, 0x33, 0xc9, 0x85, 0x64, 0xf9, 0x7a, 0x95, 0xff,
	0x60, 0xba, 0x56, 0x3d, 0xd1, 0x46, 0x59, 0x85, 0x9f, 0x04, 0x82, 0x6c, 0x09, 0xb2, 0x5e, 0x1d,
	0x1e, 0x6d, 0x6d, 0x22, 0xa7, 0x52, 0x2a, 0x4b, 0xad, 0x50, 0xb2, 0x0d, 0x96, 0xec, 0xcf, 0x04,
	0xe0, 0x83, 0xcf, 0x68, 0x98, 0xb4, 0xf8, 0x0c, 0xa6, 0x37, 0xa2, 0x66, 0x6d, 0x82, 0xd2, 0x68,
	0x39, 0x3f, 0x79, 0x4d, 0x46, 0x12, 0xc9, 0x8e, 0x27, 0x9f, 0x1c, 0xfc, 0x51, 0x5a, 0xd3, 0x17,
	0xc1, 0x88, 0xcf, 0x20, 0xae, 0x94, 0xb4, 0x54, 0x48, 0x66, 0x92, 0x49, 0x8a, 0x96, 0xf3, 0x93,
	0x6c, 0x34, 0xe5, 0xe2, 0x8e, 0xba, 0x94, 0x37, 0xaa, 0xd8, 0x99, 0x30, 0x81, 0xe8, 0xb7, 0xd0,
	0x49, 0xe4, 0xbd, 0x47, 0xa3, 0xde, 0x6b, 0xa1, 0xbd, 0xcb, 0x81, 0x87, 0xdf, 0x00, 0x76, 0x6d,
	0xe0, 0xc7, 0x10, 0xfd, 0x62, 0x7d, 0x82, 0x52, 0xb4, 0x8c, 0x0b, 0xf7, 0x89, 0x4f, 0x61, 0xba,
	0xa6, 0x75, 0xc7, 0x86, 0x6e, 0x8e, 0x47, 0x13, 0x5d, 0x82, 0x8f, 0x0c, 0xec, 0xdb, 0xc9, 0x1b,
	0x94, 0x51, 0x98, 0xdd, 0xc9, 0xf8, 0x18, 0xa0, 0x55, 0x9d, 0xa9, 0x58, 0xd9, 0x99, 0x7a, 0x48,
	0x8f, 0x83, 0xf2, 0xd5, 0xd4, 0xf8, 0x39, 0xcc, 0xda, 0x9f, 0x74, 0x55, 0xb6, 0x5d, 0xe3, 0x8f,
	0x89, 0x8b, 0x03, 0x57, 0x7f, 0xee, 0x1a, 0xfc, 0x02, 0xe2, 0x46, 0x34, 0xac, 0xb4, 0xbd, 0x66,
	0xfe, 0x52, 0x71, 0x31, 0x73, 0xc2, 0x97, 0x5e, 0xb3, 0xec, 0x15, 0x3c, 0xba, 0xf7, 0x0e, 0xf8,
	0x29, 0x4c, 0x45, 0x43, 0x39, 0x1b, 0x8e, 0x08, 0x45, 0x76, 0x09, 0x07, 0xc3, 0x95, 0xf7, 0x1a,
	0x89, 0xf6, 0x1b, 0x79, 0x09, 0x73, 0x3f, 0x87, 0xb2, 0x52, 0x9d, 0xb4, 0xc9, 0x83, 0x14, 0x2d,
	0xa7, 0x05, 0x78, 0xe9, 0xc2, 0x29, 0xe7, 0xe6, 0xef, 0x66, 0x81, 0xfe, 0x6d, 0x16, 0xe8, 0xff,
	0x66, 0x81, 0xe0, 0x59, 0xa5, 0x9a, 0xb1, 0xf7, 0x38, 0x9f, 0x87, 0x21, 0x5f, 0xb9, 0x25, 0xb9,
	0x42, 0xd7, 0xef, 0x07, 0x86, 0xab, 0x9a, 0x4a, 0x4e, 0x94, 0xe1, 0x39, 0x67, 0xd2, 0xaf, 0x50,
	0x1e, 0x7e, 0x51, 0x2d, 0xda, 0x7b, 0xab, 0xf9, 0x6e, 0x5b, 0x7c, 0x7f, 0xe8, 0xc1, 0xd3, 0xdb,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x7f, 0x85, 0xed, 0xc2, 0x02, 0x00, 0x00,
}
