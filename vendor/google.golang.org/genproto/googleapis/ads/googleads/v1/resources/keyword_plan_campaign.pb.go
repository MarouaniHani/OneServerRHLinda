// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/keyword_plan_campaign.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Keyword Plan campaign.
// Max number of keyword plan campaigns per plan allowed: 1.
type KeywordPlanCampaign struct {
	// The resource name of the Keyword Plan campaign.
	// KeywordPlanCampaign resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanCampaigns/{kp_campaign_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The keyword plan this campaign belongs to.
	KeywordPlan *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan,json=keywordPlan,proto3" json:"keyword_plan,omitempty"`
	// The ID of the Keyword Plan campaign.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Keyword Plan campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The languages targeted for the Keyword Plan campaign.
	// Max allowed: 1.
	LanguageConstants []*wrappers.StringValue `protobuf:"bytes,5,rep,name=language_constants,json=languageConstants,proto3" json:"language_constants,omitempty"`
	// Targeting network.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,6,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v1.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// A default max cpc bid in micros, and in the account currency, for all ad
	// groups under the campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,7,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// The geo targets.
	// Max number allowed: 20.
	GeoTargets           []*KeywordPlanGeoTarget `protobuf:"bytes,8,rep,name=geo_targets,json=geoTargets,proto3" json:"geo_targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordPlanCampaign) Reset()         { *m = KeywordPlanCampaign{} }
func (m *KeywordPlanCampaign) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCampaign) ProtoMessage()    {}
func (*KeywordPlanCampaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_campaign_7ac95313d7c76e43, []int{0}
}
func (m *KeywordPlanCampaign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCampaign.Unmarshal(m, b)
}
func (m *KeywordPlanCampaign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCampaign.Marshal(b, m, deterministic)
}
func (dst *KeywordPlanCampaign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCampaign.Merge(dst, src)
}
func (m *KeywordPlanCampaign) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCampaign.Size(m)
}
func (m *KeywordPlanCampaign) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCampaign.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCampaign proto.InternalMessageInfo

func (m *KeywordPlanCampaign) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanCampaign) GetKeywordPlan() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlan
	}
	return nil
}

func (m *KeywordPlanCampaign) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanCampaign) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlanCampaign) GetLanguageConstants() []*wrappers.StringValue {
	if m != nil {
		return m.LanguageConstants
	}
	return nil
}

func (m *KeywordPlanCampaign) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

func (m *KeywordPlanCampaign) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *KeywordPlanCampaign) GetGeoTargets() []*KeywordPlanGeoTarget {
	if m != nil {
		return m.GeoTargets
	}
	return nil
}

// A geo target.
// Next ID: 3
type KeywordPlanGeoTarget struct {
	// Required. The resource name of the geo target.
	GeoTargetConstant    *wrappers.StringValue `protobuf:"bytes,1,opt,name=geo_target_constant,json=geoTargetConstant,proto3" json:"geo_target_constant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KeywordPlanGeoTarget) Reset()         { *m = KeywordPlanGeoTarget{} }
func (m *KeywordPlanGeoTarget) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanGeoTarget) ProtoMessage()    {}
func (*KeywordPlanGeoTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_campaign_7ac95313d7c76e43, []int{1}
}
func (m *KeywordPlanGeoTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanGeoTarget.Unmarshal(m, b)
}
func (m *KeywordPlanGeoTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanGeoTarget.Marshal(b, m, deterministic)
}
func (dst *KeywordPlanGeoTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanGeoTarget.Merge(dst, src)
}
func (m *KeywordPlanGeoTarget) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanGeoTarget.Size(m)
}
func (m *KeywordPlanGeoTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanGeoTarget.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanGeoTarget proto.InternalMessageInfo

func (m *KeywordPlanGeoTarget) GetGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstant
	}
	return nil
}

func init() {
	proto.RegisterType((*KeywordPlanCampaign)(nil), "google.ads.googleads.v1.resources.KeywordPlanCampaign")
	proto.RegisterType((*KeywordPlanGeoTarget)(nil), "google.ads.googleads.v1.resources.KeywordPlanGeoTarget")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/keyword_plan_campaign.proto", fileDescriptor_keyword_plan_campaign_7ac95313d7c76e43)
}

var fileDescriptor_keyword_plan_campaign_7ac95313d7c76e43 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x8a, 0x13, 0x31,
	0x14, 0x86, 0xe9, 0xb4, 0xae, 0x9a, 0xd6, 0x05, 0xb3, 0x7b, 0x31, 0xac, 0x8b, 0x74, 0x57, 0x16,
	0x0a, 0x42, 0xc6, 0xae, 0xa2, 0x32, 0x22, 0x32, 0x5d, 0xa4, 0xea, 0xea, 0x52, 0xaa, 0x14, 0x91,
	0xc2, 0x90, 0x4e, 0x62, 0x08, 0xed, 0x24, 0x43, 0x92, 0xd9, 0xa2, 0xf7, 0xbe, 0x88, 0x97, 0x3e,
	0x85, 0xd7, 0x3e, 0x8a, 0x4f, 0x21, 0xcd, 0x4c, 0x66, 0x8b, 0xdb, 0x75, 0xbc, 0x3b, 0x93, 0xfc,
	0xdf, 0x39, 0xff, 0x9c, 0x73, 0x02, 0x9e, 0x33, 0x29, 0xd9, 0x82, 0x06, 0x98, 0xe8, 0xa0, 0x08,
	0x57, 0xd1, 0x79, 0x3f, 0x50, 0x54, 0xcb, 0x5c, 0x25, 0x54, 0x07, 0x73, 0xfa, 0x65, 0x29, 0x15,
	0x89, 0xb3, 0x05, 0x16, 0x71, 0x82, 0xd3, 0x0c, 0x73, 0x26, 0x50, 0xa6, 0xa4, 0x91, 0xf0, 0xa0,
	0x60, 0x10, 0x26, 0x1a, 0x55, 0x38, 0x3a, 0xef, 0xa3, 0x0a, 0xdf, 0x7b, 0x7a, 0x55, 0x05, 0x2a,
	0xf2, 0xf4, 0xaf, 0xec, 0x82, 0x9a, 0xa5, 0x54, 0xf3, 0x22, 0xf9, 0xde, 0xdd, 0x92, 0xb4, 0x5f,
	0xb3, 0xfc, 0x73, 0xb0, 0x54, 0x38, 0xcb, 0xa8, 0xd2, 0xe5, 0xfd, 0xbe, 0xcb, 0x9c, 0xf1, 0x00,
	0x0b, 0x21, 0x0d, 0x36, 0x5c, 0x8a, 0xf2, 0xf6, 0xf0, 0x67, 0x0b, 0xec, 0x9c, 0x16, 0xc9, 0x47,
	0x0b, 0x2c, 0x4e, 0x4a, 0xe3, 0xf0, 0x1e, 0xb8, 0xe5, 0xcc, 0xc5, 0x02, 0xa7, 0xd4, 0x6f, 0x74,
	0x1b, 0xbd, 0x9b, 0xe3, 0x8e, 0x3b, 0x3c, 0xc3, 0x29, 0x85, 0x2f, 0x40, 0x67, 0xdd, 0x98, 0xef,
	0x75, 0x1b, 0xbd, 0xf6, 0xf1, 0x7e, 0xf9, 0x8f, 0xc8, 0x39, 0x42, 0xef, 0x8d, 0xe2, 0x82, 0x4d,
	0xf0, 0x22, 0xa7, 0xe3, 0xf6, 0xfc, 0xa2, 0x1a, 0xbc, 0x0f, 0x3c, 0x4e, 0xfc, 0xa6, 0xc5, 0xee,
	0x5c, 0xc2, 0x5e, 0x0b, 0xf3, 0xf8, 0x51, 0x41, 0x79, 0x9c, 0xc0, 0x07, 0xa0, 0x65, 0x9d, 0xb4,
	0xfe, 0xa3, 0x8a, 0x55, 0xc2, 0x53, 0x00, 0x17, 0x58, 0xb0, 0x1c, 0x33, 0x1a, 0x27, 0x52, 0x68,
	0x83, 0x85, 0xd1, 0xfe, 0xb5, 0x6e, 0xb3, 0x96, 0xbf, 0xed, 0xb8, 0x13, 0x87, 0xc1, 0xaf, 0x60,
	0x77, 0xd3, 0x14, 0xfc, 0xad, 0x6e, 0xa3, 0xb7, 0x7d, 0xfc, 0x0a, 0x5d, 0x35, 0x63, 0x3b, 0x40,
	0xb4, 0xd6, 0xe3, 0xb3, 0x02, 0x7c, 0x29, 0xf2, 0x74, 0xc3, 0xf1, 0x18, 0xce, 0x2f, 0x9d, 0xc1,
	0x08, 0x6c, 0x27, 0x59, 0x12, 0xcf, 0x38, 0x89, 0x53, 0x9e, 0x28, 0xa9, 0xfd, 0xeb, 0xf5, 0x3d,
	0xeb, 0x24, 0x59, 0x32, 0xe0, 0xe4, 0x9d, 0x05, 0xe0, 0x47, 0xd0, 0x66, 0x54, 0xc6, 0x06, 0x2b,
	0x46, 0x8d, 0xf6, 0x6f, 0xd8, 0x26, 0x3c, 0x41, 0xb5, 0x9b, 0xb9, 0x6e, 0x71, 0x48, 0xe5, 0x07,
	0xcb, 0x8f, 0x01, 0x73, 0xa1, 0x3e, 0x24, 0x60, 0x77, 0x93, 0x06, 0xbe, 0x05, 0x3b, 0x17, 0x15,
	0xab, 0xfe, 0xdb, 0x45, 0xaa, 0x6d, 0x7f, 0x95, 0xde, 0xf5, 0x7f, 0xf0, 0xcd, 0x03, 0x47, 0x89,
	0x4c, 0xeb, 0x0d, 0x0f, 0xfc, 0x0d, 0xfb, 0x3c, 0x5a, 0x55, 0x19, 0x35, 0x3e, 0xbd, 0x29, 0x71,
	0x26, 0x57, 0x03, 0x46, 0x52, 0xb1, 0x80, 0x51, 0x61, 0x3d, 0xb8, 0x67, 0x97, 0x71, 0xfd, 0x8f,
	0x77, 0xfe, 0xac, 0x8a, 0xbe, 0x7b, 0xcd, 0x61, 0x14, 0xfd, 0xf0, 0x0e, 0x86, 0x45, 0xca, 0x88,
	0x68, 0x54, 0x84, 0xab, 0x68, 0xd2, 0x47, 0x63, 0xa7, 0xfc, 0xe5, 0x34, 0xd3, 0x88, 0xe8, 0x69,
	0xa5, 0x99, 0x4e, 0xfa, 0xd3, 0x4a, 0xf3, 0xdb, 0x3b, 0x2a, 0x2e, 0xc2, 0x30, 0x22, 0x3a, 0x0c,
	0x2b, 0x55, 0x18, 0x4e, 0xfa, 0x61, 0x58, 0xe9, 0x66, 0x5b, 0xd6, 0xec, 0xc3, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x0c, 0xaa, 0xfa, 0x04, 0x93, 0x04, 0x00, 0x00,
}
