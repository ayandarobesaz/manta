// Code generated by protoc-gen-go.
// source: econ_shared_enums.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EGCEconBaseMsg int32

const (
	EGCEconBaseMsg_k_EMsgGCGenericResult EGCEconBaseMsg = 2579
)

var EGCEconBaseMsg_name = map[int32]string{
	2579: "k_EMsgGCGenericResult",
}
var EGCEconBaseMsg_value = map[string]int32{
	"k_EMsgGCGenericResult": 2579,
}

func (x EGCEconBaseMsg) Enum() *EGCEconBaseMsg {
	p := new(EGCEconBaseMsg)
	*p = x
	return p
}
func (x EGCEconBaseMsg) String() string {
	return proto.EnumName(EGCEconBaseMsg_name, int32(x))
}
func (x *EGCEconBaseMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCEconBaseMsg_value, data, "EGCEconBaseMsg")
	if err != nil {
		return err
	}
	*x = EGCEconBaseMsg(value)
	return nil
}
func (EGCEconBaseMsg) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

type EGCMsgResponse int32

const (
	EGCMsgResponse_k_EGCMsgResponseOK           EGCMsgResponse = 0
	EGCMsgResponse_k_EGCMsgResponseDenied       EGCMsgResponse = 1
	EGCMsgResponse_k_EGCMsgResponseServerError  EGCMsgResponse = 2
	EGCMsgResponse_k_EGCMsgResponseTimeout      EGCMsgResponse = 3
	EGCMsgResponse_k_EGCMsgResponseInvalid      EGCMsgResponse = 4
	EGCMsgResponse_k_EGCMsgResponseNoMatch      EGCMsgResponse = 5
	EGCMsgResponse_k_EGCMsgResponseUnknownError EGCMsgResponse = 6
	EGCMsgResponse_k_EGCMsgResponseNotLoggedOn  EGCMsgResponse = 7
	EGCMsgResponse_k_EGCMsgFailedToCreate       EGCMsgResponse = 8
)

var EGCMsgResponse_name = map[int32]string{
	0: "k_EGCMsgResponseOK",
	1: "k_EGCMsgResponseDenied",
	2: "k_EGCMsgResponseServerError",
	3: "k_EGCMsgResponseTimeout",
	4: "k_EGCMsgResponseInvalid",
	5: "k_EGCMsgResponseNoMatch",
	6: "k_EGCMsgResponseUnknownError",
	7: "k_EGCMsgResponseNotLoggedOn",
	8: "k_EGCMsgFailedToCreate",
}
var EGCMsgResponse_value = map[string]int32{
	"k_EGCMsgResponseOK":           0,
	"k_EGCMsgResponseDenied":       1,
	"k_EGCMsgResponseServerError":  2,
	"k_EGCMsgResponseTimeout":      3,
	"k_EGCMsgResponseInvalid":      4,
	"k_EGCMsgResponseNoMatch":      5,
	"k_EGCMsgResponseUnknownError": 6,
	"k_EGCMsgResponseNotLoggedOn":  7,
	"k_EGCMsgFailedToCreate":       8,
}

func (x EGCMsgResponse) Enum() *EGCMsgResponse {
	p := new(EGCMsgResponse)
	*p = x
	return p
}
func (x EGCMsgResponse) String() string {
	return proto.EnumName(EGCMsgResponse_name, int32(x))
}
func (x *EGCMsgResponse) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCMsgResponse_value, data, "EGCMsgResponse")
	if err != nil {
		return err
	}
	*x = EGCMsgResponse(value)
	return nil
}
func (EGCMsgResponse) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{1} }

type EGCPartnerRequestResponse int32

const (
	EGCPartnerRequestResponse_k_EPartnerRequestOK                     EGCPartnerRequestResponse = 1
	EGCPartnerRequestResponse_k_EPartnerRequestBadAccount             EGCPartnerRequestResponse = 2
	EGCPartnerRequestResponse_k_EPartnerRequestNotLinked              EGCPartnerRequestResponse = 3
	EGCPartnerRequestResponse_k_EPartnerRequestUnsupportedPartnerType EGCPartnerRequestResponse = 4
)

var EGCPartnerRequestResponse_name = map[int32]string{
	1: "k_EPartnerRequestOK",
	2: "k_EPartnerRequestBadAccount",
	3: "k_EPartnerRequestNotLinked",
	4: "k_EPartnerRequestUnsupportedPartnerType",
}
var EGCPartnerRequestResponse_value = map[string]int32{
	"k_EPartnerRequestOK":                     1,
	"k_EPartnerRequestBadAccount":             2,
	"k_EPartnerRequestNotLinked":              3,
	"k_EPartnerRequestUnsupportedPartnerType": 4,
}

func (x EGCPartnerRequestResponse) Enum() *EGCPartnerRequestResponse {
	p := new(EGCPartnerRequestResponse)
	*p = x
	return p
}
func (x EGCPartnerRequestResponse) String() string {
	return proto.EnumName(EGCPartnerRequestResponse_name, int32(x))
}
func (x *EGCPartnerRequestResponse) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCPartnerRequestResponse_value, data, "EGCPartnerRequestResponse")
	if err != nil {
		return err
	}
	*x = EGCPartnerRequestResponse(value)
	return nil
}
func (EGCPartnerRequestResponse) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{2} }

type EGCMsgUseItemResponse int32

const (
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed                    EGCMsgUseItemResponse = 0
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_GiftNoOtherPlayers          EGCMsgUseItemResponse = 1
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ServerError                 EGCMsgUseItemResponse = 2
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_MiniGameAlreadyStarted      EGCMsgUseItemResponse = 3
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted       EGCMsgUseItemResponse = 4
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted EGCMsgUseItemResponse = 5
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_NotInLowPriorityPool        EGCMsgUseItemResponse = 6
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_NotHighEnoughLevel          EGCMsgUseItemResponse = 7
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EventNotActive              EGCMsgUseItemResponse = 8
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted EGCMsgUseItemResponse = 9
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_MissingRequirement          EGCMsgUseItemResponse = 10
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew        EGCMsgUseItemResponse = 11
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EmoticonUnlock_Complete     EGCMsgUseItemResponse = 12
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_Compendium         EGCMsgUseItemResponse = 13
)

var EGCMsgUseItemResponse_name = map[int32]string{
	0:  "k_EGCMsgUseItemResponse_ItemUsed",
	1:  "k_EGCMsgUseItemResponse_GiftNoOtherPlayers",
	2:  "k_EGCMsgUseItemResponse_ServerError",
	3:  "k_EGCMsgUseItemResponse_MiniGameAlreadyStarted",
	4:  "k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted",
	5:  "k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted",
	6:  "k_EGCMsgUseItemResponse_NotInLowPriorityPool",
	7:  "k_EGCMsgUseItemResponse_NotHighEnoughLevel",
	8:  "k_EGCMsgUseItemResponse_EventNotActive",
	9:  "k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted",
	10: "k_EGCMsgUseItemResponse_MissingRequirement",
	11: "k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew",
	12: "k_EGCMsgUseItemResponse_EmoticonUnlock_Complete",
	13: "k_EGCMsgUseItemResponse_ItemUsed_Compendium",
}
var EGCMsgUseItemResponse_value = map[string]int32{
	"k_EGCMsgUseItemResponse_ItemUsed":                    0,
	"k_EGCMsgUseItemResponse_GiftNoOtherPlayers":          1,
	"k_EGCMsgUseItemResponse_ServerError":                 2,
	"k_EGCMsgUseItemResponse_MiniGameAlreadyStarted":      3,
	"k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted":       4,
	"k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted": 5,
	"k_EGCMsgUseItemResponse_NotInLowPriorityPool":        6,
	"k_EGCMsgUseItemResponse_NotHighEnoughLevel":          7,
	"k_EGCMsgUseItemResponse_EventNotActive":              8,
	"k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted": 9,
	"k_EGCMsgUseItemResponse_MissingRequirement":          10,
	"k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew":        11,
	"k_EGCMsgUseItemResponse_EmoticonUnlock_Complete":     12,
	"k_EGCMsgUseItemResponse_ItemUsed_Compendium":         13,
}

func (x EGCMsgUseItemResponse) Enum() *EGCMsgUseItemResponse {
	p := new(EGCMsgUseItemResponse)
	*p = x
	return p
}
func (x EGCMsgUseItemResponse) String() string {
	return proto.EnumName(EGCMsgUseItemResponse_name, int32(x))
}
func (x *EGCMsgUseItemResponse) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCMsgUseItemResponse_value, data, "EGCMsgUseItemResponse")
	if err != nil {
		return err
	}
	*x = EGCMsgUseItemResponse(value)
	return nil
}
func (EGCMsgUseItemResponse) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{3} }

type CMsgGenericResult struct {
	Eresult          *uint32 `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	DebugMessage     *string `protobuf:"bytes,2,opt,name=debug_message" json:"debug_message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgGenericResult) Reset()                    { *m = CMsgGenericResult{} }
func (m *CMsgGenericResult) String() string            { return proto.CompactTextString(m) }
func (*CMsgGenericResult) ProtoMessage()               {}
func (*CMsgGenericResult) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

const Default_CMsgGenericResult_Eresult uint32 = 2

func (m *CMsgGenericResult) GetEresult() uint32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgGenericResult_Eresult
}

func (m *CMsgGenericResult) GetDebugMessage() string {
	if m != nil && m.DebugMessage != nil {
		return *m.DebugMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*CMsgGenericResult)(nil), "dota.CMsgGenericResult")
	proto.RegisterEnum("dota.EGCEconBaseMsg", EGCEconBaseMsg_name, EGCEconBaseMsg_value)
	proto.RegisterEnum("dota.EGCMsgResponse", EGCMsgResponse_name, EGCMsgResponse_value)
	proto.RegisterEnum("dota.EGCPartnerRequestResponse", EGCPartnerRequestResponse_name, EGCPartnerRequestResponse_value)
	proto.RegisterEnum("dota.EGCMsgUseItemResponse", EGCMsgUseItemResponse_name, EGCMsgUseItemResponse_value)
}

var fileDescriptor26 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x4f, 0x13, 0x41,
	0x14, 0x65, 0xa1, 0x15, 0xb9, 0x8a, 0x19, 0x47, 0xf9, 0x10, 0x8c, 0x12, 0x35, 0x62, 0x0a, 0x82,
	0xc2, 0x83, 0x89, 0x0f, 0x26, 0xb4, 0xac, 0x85, 0x48, 0xa1, 0x01, 0xfa, 0xdc, 0x8c, 0xbb, 0xd7,
	0xed, 0xa4, 0xbb, 0x33, 0xeb, 0xcc, 0x6c, 0x49, 0xdf, 0xfc, 0x0f, 0x3e, 0xfa, 0xe6, 0x7f, 0xf1,
	0x7f, 0x79, 0xb7, 0x40, 0x63, 0x4b, 0x5b, 0x7c, 0x9b, 0xce, 0x39, 0xe7, 0xce, 0x3d, 0xe7, 0xde,
	0x2e, 0x2c, 0x61, 0xa0, 0x55, 0xd3, 0xb6, 0x84, 0xc1, 0xb0, 0x89, 0x2a, 0x4b, 0xec, 0x56, 0x6a,
	0xb4, 0xd3, 0xbc, 0x10, 0x6a, 0x27, 0x5e, 0x7c, 0x82, 0x87, 0x95, 0x9a, 0x8d, 0xaa, 0xa8, 0xd0,
	0xc8, 0xe0, 0x14, 0x6d, 0x16, 0x3b, 0xce, 0x61, 0x16, 0x4d, 0xef, 0xb8, 0xec, 0xad, 0x79, 0x6f,
	0xe6, 0x3f, 0x7a, 0x3b, 0x7c, 0x01, 0xe6, 0x43, 0xfc, 0x9a, 0x45, 0xcd, 0x04, 0xad, 0x15, 0x11,
	0x2e, 0x4f, 0x13, 0x32, 0x57, 0xda, 0x84, 0x07, 0x7e, 0xb5, 0xe2, 0xd3, 0x1b, 0x65, 0x61, 0x91,
	0x2a, 0xf1, 0x15, 0x58, 0x68, 0x37, 0xfd, 0xbc, 0x66, 0x65, 0xa0, 0x2a, 0xfb, 0xf9, 0xb8, 0xf4,
	0x6b, 0xba, 0x47, 0x27, 0x90, 0xee, 0x52, 0xad, 0x2c, 0xf2, 0x45, 0xe0, 0x44, 0x1f, 0xb8, 0x3b,
	0xf9, 0xc2, 0xa6, 0xa8, 0xcc, 0xe2, 0xf0, 0xfd, 0x3e, 0x2a, 0x89, 0x21, 0xf3, 0xf8, 0x73, 0x58,
	0x1d, 0xc6, 0xce, 0xd0, 0x74, 0xd0, 0xf8, 0xc6, 0x68, 0xc3, 0xa6, 0xf9, 0x2a, 0x2c, 0x0d, 0x13,
	0xce, 0x65, 0x82, 0x3a, 0x73, 0x6c, 0x66, 0x14, 0x78, 0xa8, 0x3a, 0x22, 0x96, 0x21, 0x2b, 0x8c,
	0x02, 0x8f, 0x75, 0x4d, 0xb8, 0xa0, 0xc5, 0x8a, 0x7c, 0x0d, 0x9e, 0x0e, 0x83, 0x0d, 0xd5, 0x56,
	0xfa, 0x42, 0x5d, 0x3e, 0x7c, 0x67, 0x54, 0x67, 0xc7, 0xda, 0x1d, 0xe9, 0x28, 0xc2, 0xf0, 0x44,
	0xb1, 0xd9, 0x7f, 0x6d, 0x7d, 0x16, 0x32, 0xc6, 0xf0, 0x5c, 0x57, 0x0c, 0x0a, 0x87, 0xec, 0x6e,
	0xe9, 0xb7, 0x07, 0x4f, 0x08, 0xaa, 0x0b, 0xe3, 0x28, 0xb7, 0x53, 0xfc, 0x9e, 0xa1, 0x75, 0xfd,
	0xa0, 0x96, 0xe0, 0x11, 0x29, 0x07, 0x41, 0x4a, 0xea, 0x3a, 0x8d, 0x41, 0xa0, 0x2c, 0xc2, 0xbd,
	0x20, 0xd0, 0x99, 0x72, 0x94, 0xc6, 0x33, 0x58, 0xb9, 0x41, 0xc8, 0xbb, 0x92, 0xaa, 0x4d, 0x71,
	0xce, 0xf0, 0x0d, 0x58, 0xbf, 0x81, 0x37, 0x94, 0xcd, 0xd2, 0x54, 0x1b, 0x87, 0xe1, 0x15, 0x70,
	0xde, 0x4d, 0x91, 0x15, 0x4a, 0x7f, 0x8a, 0xb0, 0x70, 0xd9, 0x7f, 0x83, 0x72, 0x73, 0x98, 0xf4,
	0x1b, 0x7c, 0x05, 0x6b, 0xd7, 0xd6, 0x86, 0xa0, 0x66, 0xfe, 0x83, 0xee, 0x42, 0x9a, 0xeb, 0x16,
	0x94, 0xc6, 0xb1, 0xaa, 0xf2, 0x1b, 0xf5, 0x75, 0xe2, 0x5a, 0x68, 0xea, 0xb1, 0xe8, 0xa2, 0xb1,
	0xe4, 0x6e, 0x1d, 0x5e, 0x8e, 0xe3, 0x0f, 0xce, 0x7c, 0x07, 0xb6, 0xc6, 0x11, 0x6b, 0x52, 0xc9,
	0xaa, 0x48, 0x70, 0x2f, 0xa6, 0xa8, 0xc3, 0xee, 0x99, 0x13, 0xb9, 0x2f, 0x72, 0xfe, 0x1e, 0xde,
	0xde, 0xd6, 0x72, 0xef, 0x60, 0xab, 0x46, 0xa8, 0x5c, 0x52, 0xe0, 0x1f, 0x60, 0x77, 0x9c, 0x64,
	0xdf, 0xe8, 0xf4, 0x94, 0x46, 0x59, 0xd6, 0x2a, 0xb3, 0x57, 0x6f, 0x5d, 0x0b, 0x8b, 0xfc, 0x1d,
	0x6c, 0x8e, 0x13, 0xd2, 0x30, 0x0e, 0xd5, 0x91, 0xbe, 0xa8, 0x1b, 0xa9, 0x8d, 0x74, 0xdd, 0xba,
	0xd6, 0x31, 0x2d, 0xd3, 0x84, 0xa8, 0x48, 0x71, 0x20, 0xa3, 0x96, 0xaf, 0x74, 0x16, 0xb5, 0x8e,
	0xb0, 0x83, 0x31, 0xed, 0x56, 0x09, 0x5e, 0x8f, 0xe3, 0xfb, 0x1d, 0x54, 0xf9, 0xcc, 0xf7, 0x02,
	0x27, 0x3b, 0xb4, 0x6b, 0x93, 0x6c, 0xf4, 0x9d, 0xf7, 0x44, 0x75, 0x2d, 0x95, 0xeb, 0xfb, 0x9f,
	0x9b, 0xd4, 0x54, 0x4d, 0x5a, 0x2b, 0x55, 0x94, 0x6f, 0x90, 0x34, 0x98, 0x90, 0x98, 0xc1, 0x24,
	0xdb, 0x7e, 0xa2, 0x9d, 0xa4, 0x2f, 0x47, 0x43, 0xc5, 0x3a, 0x68, 0x93, 0xa7, 0x63, 0xbc, 0x60,
	0xf7, 0xf8, 0x2e, 0x6c, 0xff, 0xa7, 0xa2, 0xa2, 0x93, 0x34, 0x46, 0xfa, 0xef, 0xdc, 0xe7, 0xdb,
	0xb0, 0x71, 0xab, 0x9f, 0x9c, 0x8e, 0x2a, 0x94, 0x59, 0xc2, 0xe6, 0xcb, 0xc5, 0x03, 0xef, 0x87,
	0x37, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x17, 0xef, 0x79, 0x01, 0x1f, 0x05, 0x00, 0x00,
}
