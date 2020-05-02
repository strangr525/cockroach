// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/engine/enginepb/rocksdb.proto

package enginepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// SSTUserProperties contains the user-added properties of a single sstable.
type SSTUserProperties struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// ts_min is the minimum mvcc timestamp present in this sstable.
	TsMin *cockroach_util_hlc.Timestamp `protobuf:"bytes,2,opt,name=ts_min,json=tsMin" json:"ts_min,omitempty"`
	// ts_max is the maximum mvcc timestamp present in this sstable.
	TsMax *cockroach_util_hlc.Timestamp `protobuf:"bytes,3,opt,name=ts_max,json=tsMax" json:"ts_max,omitempty"`
}

func (m *SSTUserProperties) Reset()                    { *m = SSTUserProperties{} }
func (m *SSTUserProperties) String() string            { return proto.CompactTextString(m) }
func (*SSTUserProperties) ProtoMessage()               {}
func (*SSTUserProperties) Descriptor() ([]byte, []int) { return fileDescriptorRocksdb, []int{0} }

// SSTUserPropertiesCollection contains the user-added properties of every
// sstable in a RocksDB instance.
type SSTUserPropertiesCollection struct {
	Sst   []SSTUserProperties `protobuf:"bytes,1,rep,name=sst" json:"sst"`
	Error string              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *SSTUserPropertiesCollection) Reset()         { *m = SSTUserPropertiesCollection{} }
func (m *SSTUserPropertiesCollection) String() string { return proto.CompactTextString(m) }
func (*SSTUserPropertiesCollection) ProtoMessage()    {}
func (*SSTUserPropertiesCollection) Descriptor() ([]byte, []int) {
	return fileDescriptorRocksdb, []int{1}
}

func init() {
	proto.RegisterType((*SSTUserProperties)(nil), "cockroach.storage.engine.enginepb.SSTUserProperties")
	proto.RegisterType((*SSTUserPropertiesCollection)(nil), "cockroach.storage.engine.enginepb.SSTUserPropertiesCollection")
}
func (m *SSTUserProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSTUserProperties) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	if m.TsMin != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(m.TsMin.Size()))
		n1, err := m.TsMin.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.TsMax != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(m.TsMax.Size()))
		n2, err := m.TsMax.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *SSTUserPropertiesCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSTUserPropertiesCollection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sst) > 0 {
		for _, msg := range m.Sst {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRocksdb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	return i, nil
}

func encodeVarintRocksdb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SSTUserProperties) Size() (n int) {
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovRocksdb(uint64(l))
	}
	if m.TsMin != nil {
		l = m.TsMin.Size()
		n += 1 + l + sovRocksdb(uint64(l))
	}
	if m.TsMax != nil {
		l = m.TsMax.Size()
		n += 1 + l + sovRocksdb(uint64(l))
	}
	return n
}

func (m *SSTUserPropertiesCollection) Size() (n int) {
	var l int
	_ = l
	if len(m.Sst) > 0 {
		for _, e := range m.Sst {
			l = e.Size()
			n += 1 + l + sovRocksdb(uint64(l))
		}
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovRocksdb(uint64(l))
	}
	return n
}

func sovRocksdb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRocksdb(x uint64) (n int) {
	return sovRocksdb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SSTUserProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRocksdb
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
			return fmt.Errorf("proto: SSTUserProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSTUserProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TsMin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TsMin == nil {
				m.TsMin = &cockroach_util_hlc.Timestamp{}
			}
			if err := m.TsMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TsMax", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TsMax == nil {
				m.TsMax = &cockroach_util_hlc.Timestamp{}
			}
			if err := m.TsMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRocksdb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRocksdb
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
func (m *SSTUserPropertiesCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRocksdb
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
			return fmt.Errorf("proto: SSTUserPropertiesCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSTUserPropertiesCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sst", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sst = append(m.Sst, SSTUserProperties{})
			if err := m.Sst[len(m.Sst)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRocksdb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRocksdb
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
func skipRocksdb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRocksdb
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
					return 0, ErrIntOverflowRocksdb
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
					return 0, ErrIntOverflowRocksdb
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
				return 0, ErrInvalidLengthRocksdb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRocksdb
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
				next, err := skipRocksdb(dAtA[start:])
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
	ErrInvalidLengthRocksdb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRocksdb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage/engine/enginepb/rocksdb.proto", fileDescriptorRocksdb) }

var fileDescriptorRocksdb = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x17, 0xbb, 0x0d, 0x97, 0x9d, 0x0c, 0x3b, 0x94, 0x89, 0xb1, 0x16, 0x84, 0x9e, 0x52,
	0x98, 0x7b, 0x82, 0x79, 0x55, 0x90, 0x6e, 0x5e, 0xbc, 0x48, 0x1a, 0x43, 0x1b, 0xd6, 0xf6, 0x5f,
	0x92, 0x08, 0xbb, 0xfa, 0x06, 0x82, 0x2f, 0xd5, 0xa3, 0x47, 0x4f, 0xa2, 0xf5, 0x45, 0xa4, 0xad,
	0xc5, 0xc3, 0x0e, 0x7a, 0xca, 0x97, 0x90, 0xdf, 0xf7, 0xff, 0xfe, 0x1f, 0x3e, 0x37, 0x16, 0x34,
	0x4f, 0x64, 0x28, 0x8b, 0x44, 0x15, 0xfd, 0x51, 0xc6, 0xa1, 0x06, 0xb1, 0x35, 0x0f, 0x31, 0x2b,
	0x35, 0x58, 0x20, 0x67, 0x02, 0xc4, 0x56, 0x03, 0x17, 0x29, 0xfb, 0x01, 0x58, 0xf7, 0x93, 0xf5,
	0xc0, 0xdc, 0x7d, 0xb4, 0x2a, 0x0b, 0xd3, 0x4c, 0x84, 0x56, 0xe5, 0xd2, 0x58, 0x9e, 0x97, 0x1d,
	0x3c, 0x9f, 0x25, 0x90, 0x40, 0x2b, 0xc3, 0x46, 0x75, 0xaf, 0xfe, 0x0b, 0xc2, 0x47, 0xeb, 0xf5,
	0xe6, 0xd6, 0x48, 0x7d, 0xa3, 0xa1, 0x94, 0xda, 0x2a, 0x69, 0x08, 0xc1, 0xc3, 0x92, 0xdb, 0xd4,
	0x45, 0x1e, 0x0a, 0x26, 0x51, 0xab, 0xc9, 0x12, 0x8f, 0xad, 0xb9, 0xcf, 0x55, 0xe1, 0x1e, 0x78,
	0x28, 0x98, 0x2e, 0x4e, 0xd8, 0x6f, 0x9a, 0x66, 0x28, 0x4b, 0x33, 0xc1, 0x36, 0xfd, 0xd0, 0x68,
	0x64, 0xcd, 0xb5, 0x2a, 0x7a, 0x8a, 0xef, 0x5c, 0xe7, 0xbf, 0x14, 0xdf, 0xf9, 0x4f, 0x08, 0x1f,
	0xef, 0xa5, 0xba, 0x84, 0x2c, 0x93, 0xc2, 0x2a, 0x28, 0xc8, 0x15, 0x76, 0x8c, 0xb1, 0x2e, 0xf2,
	0x9c, 0x60, 0xba, 0x58, 0xb2, 0x3f, 0x6b, 0x61, 0x7b, 0x66, 0xab, 0x61, 0xf5, 0x7e, 0x3a, 0x88,
	0x1a, 0x1b, 0x32, 0xc3, 0x23, 0xa9, 0x35, 0xe8, 0x76, 0xb1, 0x49, 0xd4, 0x5d, 0x56, 0x7e, 0xf5,
	0x49, 0x07, 0x55, 0x4d, 0xd1, 0x6b, 0x4d, 0xd1, 0x5b, 0x4d, 0xd1, 0x47, 0x4d, 0xd1, 0xf3, 0x17,
	0x1d, 0xdc, 0x1d, 0xf6, 0xb6, 0xf1, 0xb8, 0x2d, 0xf1, 0xe2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x3a,
	0xf0, 0xb4, 0xfc, 0xc0, 0x01, 0x00, 0x00,
}