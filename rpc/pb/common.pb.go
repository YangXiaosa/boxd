// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common.proto

package rpcpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pb "github.com/BOXFoundation/boxd/core/pb"

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

type Utxo struct {
	OutPoint    *pb.OutPoint `protobuf:"bytes,1,opt,name=out_point,json=outPoint" json:"out_point,omitempty"`
	TxOut       *pb.TxOut    `protobuf:"bytes,2,opt,name=tx_out,json=txOut" json:"tx_out,omitempty"`
	BlockHeight int32        `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	IsCoinbase  bool         `protobuf:"varint,4,opt,name=is_coinbase,json=isCoinbase,proto3" json:"is_coinbase,omitempty"`
	IsSpent     bool         `protobuf:"varint,5,opt,name=is_spent,json=isSpent,proto3" json:"is_spent,omitempty"`
}

func (m *Utxo) Reset()         { *m = Utxo{} }
func (m *Utxo) String() string { return proto.CompactTextString(m) }
func (*Utxo) ProtoMessage()    {}
func (*Utxo) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_41720900d7f939cd, []int{0}
}
func (m *Utxo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Utxo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Utxo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Utxo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Utxo.Merge(dst, src)
}
func (m *Utxo) XXX_Size() int {
	return m.Size()
}
func (m *Utxo) XXX_DiscardUnknown() {
	xxx_messageInfo_Utxo.DiscardUnknown(m)
}

var xxx_messageInfo_Utxo proto.InternalMessageInfo

func (m *Utxo) GetOutPoint() *pb.OutPoint {
	if m != nil {
		return m.OutPoint
	}
	return nil
}

func (m *Utxo) GetTxOut() *pb.TxOut {
	if m != nil {
		return m.TxOut
	}
	return nil
}

func (m *Utxo) GetBlockHeight() int32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Utxo) GetIsCoinbase() bool {
	if m != nil {
		return m.IsCoinbase
	}
	return false
}

func (m *Utxo) GetIsSpent() bool {
	if m != nil {
		return m.IsSpent
	}
	return false
}

type BaseResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_41720900d7f939cd, []int{1}
}
func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(dst, src)
}
func (m *BaseResponse) XXX_Size() int {
	return m.Size()
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BaseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Utxo)(nil), "rpcpb.Utxo")
	proto.RegisterType((*BaseResponse)(nil), "rpcpb.BaseResponse")
}
func (m *Utxo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Utxo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OutPoint != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.OutPoint.Size()))
		n1, err := m.OutPoint.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.TxOut != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.TxOut.Size()))
		n2, err := m.TxOut.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.BlockHeight != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.BlockHeight))
	}
	if m.IsCoinbase {
		dAtA[i] = 0x20
		i++
		if m.IsCoinbase {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.IsSpent {
		dAtA[i] = 0x28
		i++
		if m.IsSpent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *BaseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Utxo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OutPoint != nil {
		l = m.OutPoint.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.TxOut != nil {
		l = m.TxOut.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovCommon(uint64(m.BlockHeight))
	}
	if m.IsCoinbase {
		n += 2
	}
	if m.IsSpent {
		n += 2
	}
	return n
}

func (m *BaseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovCommon(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Utxo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Utxo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Utxo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutPoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OutPoint == nil {
				m.OutPoint = &pb.OutPoint{}
			}
			if err := m.OutPoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxOut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TxOut == nil {
				m.TxOut = &pb.TxOut{}
			}
			if err := m.TxOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsCoinbase", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
			m.IsCoinbase = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSpent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
			m.IsSpent = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *BaseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: BaseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommon
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
				next, err := skipCommon(dAtA[start:])
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
	ErrInvalidLengthCommon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_41720900d7f939cd) }

var fileDescriptor_common_41720900d7f939cd = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xbd, 0x4e, 0x23, 0x31,
	0x14, 0x85, 0xe3, 0xdd, 0x4c, 0x7e, 0x9c, 0xac, 0xb4, 0x72, 0xe5, 0xdd, 0x62, 0x08, 0x11, 0x45,
	0x1a, 0x66, 0x24, 0x68, 0x28, 0xa8, 0x82, 0x84, 0xe8, 0x02, 0x06, 0x24, 0xba, 0xd1, 0xd8, 0xb1,
	0x12, 0x2b, 0x19, 0x5f, 0x6b, 0x7c, 0x8d, 0xf2, 0x18, 0x3c, 0x0f, 0x4f, 0x40, 0x99, 0x92, 0x12,
	0x25, 0x2f, 0x82, 0xc6, 0x24, 0x95, 0xcf, 0xf9, 0x7c, 0xec, 0x7b, 0x74, 0xe9, 0x50, 0x41, 0x55,
	0x81, 0xcd, 0x5c, 0x0d, 0x08, 0x2c, 0xa9, 0x9d, 0x72, 0xf2, 0xff, 0xd5, 0xc2, 0xe0, 0x32, 0xc8,
	0x4c, 0x41, 0x95, 0x4f, 0x67, 0x2f, 0xb7, 0x10, 0xec, 0xbc, 0x44, 0x03, 0x36, 0x7f, 0x08, 0x46,
	0xad, 0xbc, 0x59, 0xbf, 0xea, 0x3a, 0x57, 0x50, 0xeb, 0xdc, 0xc9, 0x5c, 0xae, 0x41, 0xad, 0x7e,
	0x3e, 0x18, 0xbf, 0x13, 0xda, 0x7e, 0xc6, 0x0d, 0xb0, 0x73, 0xda, 0x87, 0x80, 0x85, 0x03, 0x63,
	0x91, 0x93, 0x11, 0x99, 0x0c, 0x2e, 0xfe, 0x66, 0xcd, 0x0b, 0x27, 0xb3, 0x59, 0xc0, 0xfb, 0x86,
	0x8b, 0x1e, 0x1c, 0x14, 0x3b, 0xa3, 0x1d, 0xdc, 0x14, 0x10, 0x90, 0xff, 0x8a, 0xd9, 0x3f, 0xc7,
	0xec, 0xd3, 0x66, 0x16, 0x50, 0x24, 0xd8, 0x1c, 0xec, 0x94, 0x0e, 0xe3, 0xb0, 0x62, 0xa9, 0xcd,
	0x62, 0x89, 0xfc, 0xf7, 0x88, 0x4c, 0x12, 0x31, 0x88, 0xec, 0x2e, 0x22, 0x76, 0x42, 0x07, 0xc6,
	0x17, 0x0a, 0x8c, 0x95, 0xa5, 0xd7, 0xbc, 0x3d, 0x22, 0x93, 0x9e, 0xa0, 0xc6, 0xdf, 0x1c, 0x08,
	0xfb, 0x47, 0x7b, 0xc6, 0x17, 0xde, 0x69, 0x8b, 0x3c, 0x89, 0xb7, 0x5d, 0xe3, 0x1f, 0x1b, 0x3b,
	0xbe, 0xa6, 0xc3, 0x69, 0xe9, 0xb5, 0xd0, 0xde, 0x81, 0xf5, 0x9a, 0x31, 0xda, 0x56, 0x30, 0xd7,
	0xb1, 0x7e, 0x22, 0xa2, 0x66, 0x9c, 0x76, 0x2b, 0xed, 0x7d, 0xb9, 0xd0, 0xb1, 0x69, 0x5f, 0x1c,
	0xed, 0x94, 0x7f, 0xec, 0x52, 0xb2, 0xdd, 0xa5, 0xe4, 0x6b, 0x97, 0x92, 0xb7, 0x7d, 0xda, 0xda,
	0xee, 0xd3, 0xd6, 0xe7, 0x3e, 0x6d, 0xc9, 0x4e, 0xdc, 0xcd, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x3b, 0xa2, 0x37, 0x64, 0x6c, 0x01, 0x00, 0x00,
}
