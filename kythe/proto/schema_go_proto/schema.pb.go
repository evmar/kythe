// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/schema.proto

package schema_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NodeKind int32

const (
	NodeKind_UNKNOWN_NODE_KIND NodeKind = 0
	NodeKind_ABS               NodeKind = 1
	NodeKind_ABSVAR            NodeKind = 2
	NodeKind_ANCHOR            NodeKind = 3
	NodeKind_CONSTANT          NodeKind = 4
	NodeKind_DIAGNOSTIC        NodeKind = 5
	NodeKind_DOC               NodeKind = 6
	NodeKind_FILE              NodeKind = 7
	NodeKind_INTERFACE         NodeKind = 8
	NodeKind_FUNCTION          NodeKind = 9
	NodeKind_LOOKUP            NodeKind = 10
	NodeKind_MACRO             NodeKind = 11
	NodeKind_META              NodeKind = 12
	NodeKind_NAME              NodeKind = 13
	NodeKind_PACKAGE           NodeKind = 14
	NodeKind_PROCESS           NodeKind = 15
	NodeKind_RECORD            NodeKind = 16
	NodeKind_SUM               NodeKind = 17
	NodeKind_SYMBOL            NodeKind = 18
	NodeKind_TALIAS            NodeKind = 19
	NodeKind_TAPP              NodeKind = 20
	NodeKind_TBUILTIN          NodeKind = 21
	NodeKind_TNOMINAL          NodeKind = 22
	NodeKind_TSIGMA            NodeKind = 23
	NodeKind_VARIABLE          NodeKind = 24
	NodeKind_VCS               NodeKind = 25
)

var NodeKind_name = map[int32]string{
	0:  "UNKNOWN_NODE_KIND",
	1:  "ABS",
	2:  "ABSVAR",
	3:  "ANCHOR",
	4:  "CONSTANT",
	5:  "DIAGNOSTIC",
	6:  "DOC",
	7:  "FILE",
	8:  "INTERFACE",
	9:  "FUNCTION",
	10: "LOOKUP",
	11: "MACRO",
	12: "META",
	13: "NAME",
	14: "PACKAGE",
	15: "PROCESS",
	16: "RECORD",
	17: "SUM",
	18: "SYMBOL",
	19: "TALIAS",
	20: "TAPP",
	21: "TBUILTIN",
	22: "TNOMINAL",
	23: "TSIGMA",
	24: "VARIABLE",
	25: "VCS",
}
var NodeKind_value = map[string]int32{
	"UNKNOWN_NODE_KIND": 0,
	"ABS":               1,
	"ABSVAR":            2,
	"ANCHOR":            3,
	"CONSTANT":          4,
	"DIAGNOSTIC":        5,
	"DOC":               6,
	"FILE":              7,
	"INTERFACE":         8,
	"FUNCTION":          9,
	"LOOKUP":            10,
	"MACRO":             11,
	"META":              12,
	"NAME":              13,
	"PACKAGE":           14,
	"PROCESS":           15,
	"RECORD":            16,
	"SUM":               17,
	"SYMBOL":            18,
	"TALIAS":            19,
	"TAPP":              20,
	"TBUILTIN":          21,
	"TNOMINAL":          22,
	"TSIGMA":            23,
	"VARIABLE":          24,
	"VCS":               25,
}

func (x NodeKind) String() string {
	return proto.EnumName(NodeKind_name, int32(x))
}
func (NodeKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schema_91faf3e8973c7842, []int{0}
}

type Subkind int32

const (
	Subkind_UNKNOWN_SUBKIND Subkind = 0
	Subkind_CATEGORY        Subkind = 1
	Subkind_CLASS           Subkind = 2
	Subkind_CONSTRUCTOR     Subkind = 3
	Subkind_DESTRUCTOR      Subkind = 4
	Subkind_ENUM            Subkind = 5
	Subkind_ENUM_CLASS      Subkind = 6
	Subkind_FIELD           Subkind = 7
	Subkind_IMPLICIT        Subkind = 8
	Subkind_IMPORT          Subkind = 9
	Subkind_INITIALIZER     Subkind = 10
	Subkind_LOCAL           Subkind = 11
	Subkind_LOCAL_PARAMETER Subkind = 12
	Subkind_METHOD          Subkind = 13
	Subkind_NAMESPACE       Subkind = 14
	Subkind_STRUCT          Subkind = 15
	Subkind_TYPE            Subkind = 16
	Subkind_UNION           Subkind = 17
)

var Subkind_name = map[int32]string{
	0:  "UNKNOWN_SUBKIND",
	1:  "CATEGORY",
	2:  "CLASS",
	3:  "CONSTRUCTOR",
	4:  "DESTRUCTOR",
	5:  "ENUM",
	6:  "ENUM_CLASS",
	7:  "FIELD",
	8:  "IMPLICIT",
	9:  "IMPORT",
	10: "INITIALIZER",
	11: "LOCAL",
	12: "LOCAL_PARAMETER",
	13: "METHOD",
	14: "NAMESPACE",
	15: "STRUCT",
	16: "TYPE",
	17: "UNION",
}
var Subkind_value = map[string]int32{
	"UNKNOWN_SUBKIND": 0,
	"CATEGORY":        1,
	"CLASS":           2,
	"CONSTRUCTOR":     3,
	"DESTRUCTOR":      4,
	"ENUM":            5,
	"ENUM_CLASS":      6,
	"FIELD":           7,
	"IMPLICIT":        8,
	"IMPORT":          9,
	"INITIALIZER":     10,
	"LOCAL":           11,
	"LOCAL_PARAMETER": 12,
	"METHOD":          13,
	"NAMESPACE":       14,
	"STRUCT":          15,
	"TYPE":            16,
	"UNION":           17,
}

func (x Subkind) String() string {
	return proto.EnumName(Subkind_name, int32(x))
}
func (Subkind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schema_91faf3e8973c7842, []int{1}
}

type FactName int32

const (
	FactName_UNKNOWN_FACT_NAME FactName = 0
	FactName_CODE              FactName = 1
	FactName_COMPLETE          FactName = 2
	FactName_CONTEXT_URL       FactName = 3
	FactName_DETAILS           FactName = 4
	FactName_DOC_URI           FactName = 5
	FactName_LABEL             FactName = 6
	FactName_LOC_END           FactName = 7
	FactName_LOC_START         FactName = 8
	FactName_MESSAGE           FactName = 9
	FactName_NODE_KIND         FactName = 10
	FactName_PARAM_DEFAULT     FactName = 11
	FactName_RULE_CLASS        FactName = 12
	FactName_SNIPPET_END       FactName = 13
	FactName_SNIPPET_START     FactName = 14
	FactName_SUBKIND           FactName = 15
	FactName_TEXT              FactName = 16
	FactName_TEXT_ENCODING     FactName = 17
	FactName_VISIBILITY        FactName = 18
)

var FactName_name = map[int32]string{
	0:  "UNKNOWN_FACT_NAME",
	1:  "CODE",
	2:  "COMPLETE",
	3:  "CONTEXT_URL",
	4:  "DETAILS",
	5:  "DOC_URI",
	6:  "LABEL",
	7:  "LOC_END",
	8:  "LOC_START",
	9:  "MESSAGE",
	10: "NODE_KIND",
	11: "PARAM_DEFAULT",
	12: "RULE_CLASS",
	13: "SNIPPET_END",
	14: "SNIPPET_START",
	15: "SUBKIND",
	16: "TEXT",
	17: "TEXT_ENCODING",
	18: "VISIBILITY",
}
var FactName_value = map[string]int32{
	"UNKNOWN_FACT_NAME": 0,
	"CODE":              1,
	"COMPLETE":          2,
	"CONTEXT_URL":       3,
	"DETAILS":           4,
	"DOC_URI":           5,
	"LABEL":             6,
	"LOC_END":           7,
	"LOC_START":         8,
	"MESSAGE":           9,
	"NODE_KIND":         10,
	"PARAM_DEFAULT":     11,
	"RULE_CLASS":        12,
	"SNIPPET_END":       13,
	"SNIPPET_START":     14,
	"SUBKIND":           15,
	"TEXT":              16,
	"TEXT_ENCODING":     17,
	"VISIBILITY":        18,
}

func (x FactName) String() string {
	return proto.EnumName(FactName_name, int32(x))
}
func (FactName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schema_91faf3e8973c7842, []int{2}
}

type EdgeKind int32

const (
	EdgeKind_UNKNOWN_EDGE_KIND        EdgeKind = 0
	EdgeKind_ALIASES                  EdgeKind = 1
	EdgeKind_ALIASES_ROOT             EdgeKind = 2
	EdgeKind_ANNOTATED_BY             EdgeKind = 3
	EdgeKind_BOUNDED_LOWER            EdgeKind = 4
	EdgeKind_BOUNDED_UPPER            EdgeKind = 5
	EdgeKind_CHILD_OF                 EdgeKind = 6
	EdgeKind_CHILD_OF_CONTEXT         EdgeKind = 7
	EdgeKind_COMPLETES                EdgeKind = 8
	EdgeKind_COMPLETES_UNIQUELY       EdgeKind = 9
	EdgeKind_DEFINES                  EdgeKind = 10
	EdgeKind_DEFINES_BINDING          EdgeKind = 11
	EdgeKind_DEPENDS                  EdgeKind = 12
	EdgeKind_DOCUMENTS                EdgeKind = 13
	EdgeKind_EXPORTS                  EdgeKind = 14
	EdgeKind_EXTENDS                  EdgeKind = 15
	EdgeKind_GENERATES                EdgeKind = 16
	EdgeKind_IMPUTES                  EdgeKind = 17
	EdgeKind_INSTANTIATES             EdgeKind = 18
	EdgeKind_INSTANTIATES_SPECULATIVE EdgeKind = 19
	EdgeKind_NAMED                    EdgeKind = 20
	EdgeKind_OVERRIDES                EdgeKind = 21
	EdgeKind_OVERRIDES_ROOT           EdgeKind = 22
	EdgeKind_OVERRIDES_TRANSITIVE     EdgeKind = 23
	EdgeKind_PARAM                    EdgeKind = 24
	EdgeKind_REF                      EdgeKind = 25
	EdgeKind_REF_CALL                 EdgeKind = 26
	EdgeKind_REF_CALL_IMPLICIT        EdgeKind = 27
	EdgeKind_REF_DOC                  EdgeKind = 28
	EdgeKind_REF_EXPANDS              EdgeKind = 29
	EdgeKind_REF_EXPANDS_TRANSITIVE   EdgeKind = 30
	EdgeKind_REF_FILE                 EdgeKind = 31
	EdgeKind_REF_IMPLICIT             EdgeKind = 32
	EdgeKind_REF_IMPORTS              EdgeKind = 33
	EdgeKind_REF_INCLUDES             EdgeKind = 34
	EdgeKind_REF_INIT                 EdgeKind = 35
	EdgeKind_REF_INIT_IMPLICIT        EdgeKind = 36
	EdgeKind_REF_QUERIES              EdgeKind = 37
	EdgeKind_SATISFIES                EdgeKind = 38
	EdgeKind_SPECIALIZES              EdgeKind = 39
	EdgeKind_SPECIALIZES_SPECULATIVE  EdgeKind = 40
	EdgeKind_TAGGED                   EdgeKind = 41
	EdgeKind_TYPED                    EdgeKind = 42
	EdgeKind_UNDEFINES                EdgeKind = 43
)

var EdgeKind_name = map[int32]string{
	0:  "UNKNOWN_EDGE_KIND",
	1:  "ALIASES",
	2:  "ALIASES_ROOT",
	3:  "ANNOTATED_BY",
	4:  "BOUNDED_LOWER",
	5:  "BOUNDED_UPPER",
	6:  "CHILD_OF",
	7:  "CHILD_OF_CONTEXT",
	8:  "COMPLETES",
	9:  "COMPLETES_UNIQUELY",
	10: "DEFINES",
	11: "DEFINES_BINDING",
	12: "DEPENDS",
	13: "DOCUMENTS",
	14: "EXPORTS",
	15: "EXTENDS",
	16: "GENERATES",
	17: "IMPUTES",
	18: "INSTANTIATES",
	19: "INSTANTIATES_SPECULATIVE",
	20: "NAMED",
	21: "OVERRIDES",
	22: "OVERRIDES_ROOT",
	23: "OVERRIDES_TRANSITIVE",
	24: "PARAM",
	25: "REF",
	26: "REF_CALL",
	27: "REF_CALL_IMPLICIT",
	28: "REF_DOC",
	29: "REF_EXPANDS",
	30: "REF_EXPANDS_TRANSITIVE",
	31: "REF_FILE",
	32: "REF_IMPLICIT",
	33: "REF_IMPORTS",
	34: "REF_INCLUDES",
	35: "REF_INIT",
	36: "REF_INIT_IMPLICIT",
	37: "REF_QUERIES",
	38: "SATISFIES",
	39: "SPECIALIZES",
	40: "SPECIALIZES_SPECULATIVE",
	41: "TAGGED",
	42: "TYPED",
	43: "UNDEFINES",
}
var EdgeKind_value = map[string]int32{
	"UNKNOWN_EDGE_KIND":        0,
	"ALIASES":                  1,
	"ALIASES_ROOT":             2,
	"ANNOTATED_BY":             3,
	"BOUNDED_LOWER":            4,
	"BOUNDED_UPPER":            5,
	"CHILD_OF":                 6,
	"CHILD_OF_CONTEXT":         7,
	"COMPLETES":                8,
	"COMPLETES_UNIQUELY":       9,
	"DEFINES":                  10,
	"DEFINES_BINDING":          11,
	"DEPENDS":                  12,
	"DOCUMENTS":                13,
	"EXPORTS":                  14,
	"EXTENDS":                  15,
	"GENERATES":                16,
	"IMPUTES":                  17,
	"INSTANTIATES":             18,
	"INSTANTIATES_SPECULATIVE": 19,
	"NAMED":                    20,
	"OVERRIDES":                21,
	"OVERRIDES_ROOT":           22,
	"OVERRIDES_TRANSITIVE":     23,
	"PARAM":                    24,
	"REF":                      25,
	"REF_CALL":                 26,
	"REF_CALL_IMPLICIT":        27,
	"REF_DOC":                  28,
	"REF_EXPANDS":              29,
	"REF_EXPANDS_TRANSITIVE":   30,
	"REF_FILE":                 31,
	"REF_IMPLICIT":             32,
	"REF_IMPORTS":              33,
	"REF_INCLUDES":             34,
	"REF_INIT":                 35,
	"REF_INIT_IMPLICIT":        36,
	"REF_QUERIES":              37,
	"SATISFIES":                38,
	"SPECIALIZES":              39,
	"SPECIALIZES_SPECULATIVE":  40,
	"TAGGED":                   41,
	"TYPED":                    42,
	"UNDEFINES":                43,
}

func (x EdgeKind) String() string {
	return proto.EnumName(EdgeKind_name, int32(x))
}
func (EdgeKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schema_91faf3e8973c7842, []int{3}
}

type Index struct {
	EdgeKinds            []*Index_EdgeKinds `protobuf:"bytes,1,rep,name=edge_kinds,json=edgeKinds" json:"edge_kinds,omitempty"`
	NodeKinds            []*Index_NodeKinds `protobuf:"bytes,2,rep,name=node_kinds,json=nodeKinds" json:"node_kinds,omitempty"`
	Subkinds             []*Index_Subkinds  `protobuf:"bytes,3,rep,name=subkinds" json:"subkinds,omitempty"`
	FactNames            []*Index_FactNames `protobuf:"bytes,4,rep,name=fact_names,json=factNames" json:"fact_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_91faf3e8973c7842, []int{0}
}
func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (dst *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(dst, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetEdgeKinds() []*Index_EdgeKinds {
	if m != nil {
		return m.EdgeKinds
	}
	return nil
}

func (m *Index) GetNodeKinds() []*Index_NodeKinds {
	if m != nil {
		return m.NodeKinds
	}
	return nil
}

func (m *Index) GetSubkinds() []*Index_Subkinds {
	if m != nil {
		return m.Subkinds
	}
	return nil
}

func (m *Index) GetFactNames() []*Index_FactNames {
	if m != nil {
		return m.FactNames
	}
	return nil
}

type Index_EdgeKinds struct {
	Prefix               string              `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	EdgeKind             map[string]EdgeKind `protobuf:"bytes,2,rep,name=edge_kind,json=edgeKind" json:"edge_kind,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=kythe.proto.schema.EdgeKind"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Index_EdgeKinds) Reset()         { *m = Index_EdgeKinds{} }
func (m *Index_EdgeKinds) String() string { return proto.CompactTextString(m) }
func (*Index_EdgeKinds) ProtoMessage()    {}
func (*Index_EdgeKinds) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_91faf3e8973c7842, []int{0, 0}
}
func (m *Index_EdgeKinds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index_EdgeKinds.Unmarshal(m, b)
}
func (m *Index_EdgeKinds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index_EdgeKinds.Marshal(b, m, deterministic)
}
func (dst *Index_EdgeKinds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index_EdgeKinds.Merge(dst, src)
}
func (m *Index_EdgeKinds) XXX_Size() int {
	return xxx_messageInfo_Index_EdgeKinds.Size(m)
}
func (m *Index_EdgeKinds) XXX_DiscardUnknown() {
	xxx_messageInfo_Index_EdgeKinds.DiscardUnknown(m)
}

var xxx_messageInfo_Index_EdgeKinds proto.InternalMessageInfo

func (m *Index_EdgeKinds) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Index_EdgeKinds) GetEdgeKind() map[string]EdgeKind {
	if m != nil {
		return m.EdgeKind
	}
	return nil
}

type Index_NodeKinds struct {
	Prefix               string              `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	NodeKind             map[string]NodeKind `protobuf:"bytes,2,rep,name=node_kind,json=nodeKind" json:"node_kind,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=kythe.proto.schema.NodeKind"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Index_NodeKinds) Reset()         { *m = Index_NodeKinds{} }
func (m *Index_NodeKinds) String() string { return proto.CompactTextString(m) }
func (*Index_NodeKinds) ProtoMessage()    {}
func (*Index_NodeKinds) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_91faf3e8973c7842, []int{0, 1}
}
func (m *Index_NodeKinds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index_NodeKinds.Unmarshal(m, b)
}
func (m *Index_NodeKinds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index_NodeKinds.Marshal(b, m, deterministic)
}
func (dst *Index_NodeKinds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index_NodeKinds.Merge(dst, src)
}
func (m *Index_NodeKinds) XXX_Size() int {
	return xxx_messageInfo_Index_NodeKinds.Size(m)
}
func (m *Index_NodeKinds) XXX_DiscardUnknown() {
	xxx_messageInfo_Index_NodeKinds.DiscardUnknown(m)
}

var xxx_messageInfo_Index_NodeKinds proto.InternalMessageInfo

func (m *Index_NodeKinds) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Index_NodeKinds) GetNodeKind() map[string]NodeKind {
	if m != nil {
		return m.NodeKind
	}
	return nil
}

type Index_Subkinds struct {
	Prefix               string             `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	Subkind              map[string]Subkind `protobuf:"bytes,2,rep,name=subkind" json:"subkind,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=kythe.proto.schema.Subkind"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Index_Subkinds) Reset()         { *m = Index_Subkinds{} }
func (m *Index_Subkinds) String() string { return proto.CompactTextString(m) }
func (*Index_Subkinds) ProtoMessage()    {}
func (*Index_Subkinds) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_91faf3e8973c7842, []int{0, 2}
}
func (m *Index_Subkinds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index_Subkinds.Unmarshal(m, b)
}
func (m *Index_Subkinds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index_Subkinds.Marshal(b, m, deterministic)
}
func (dst *Index_Subkinds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index_Subkinds.Merge(dst, src)
}
func (m *Index_Subkinds) XXX_Size() int {
	return xxx_messageInfo_Index_Subkinds.Size(m)
}
func (m *Index_Subkinds) XXX_DiscardUnknown() {
	xxx_messageInfo_Index_Subkinds.DiscardUnknown(m)
}

var xxx_messageInfo_Index_Subkinds proto.InternalMessageInfo

func (m *Index_Subkinds) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Index_Subkinds) GetSubkind() map[string]Subkind {
	if m != nil {
		return m.Subkind
	}
	return nil
}

type Index_FactNames struct {
	Prefix               string              `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	FactName             map[string]FactName `protobuf:"bytes,2,rep,name=fact_name,json=factName" json:"fact_name,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=kythe.proto.schema.FactName"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Index_FactNames) Reset()         { *m = Index_FactNames{} }
func (m *Index_FactNames) String() string { return proto.CompactTextString(m) }
func (*Index_FactNames) ProtoMessage()    {}
func (*Index_FactNames) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_91faf3e8973c7842, []int{0, 3}
}
func (m *Index_FactNames) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index_FactNames.Unmarshal(m, b)
}
func (m *Index_FactNames) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index_FactNames.Marshal(b, m, deterministic)
}
func (dst *Index_FactNames) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index_FactNames.Merge(dst, src)
}
func (m *Index_FactNames) XXX_Size() int {
	return xxx_messageInfo_Index_FactNames.Size(m)
}
func (m *Index_FactNames) XXX_DiscardUnknown() {
	xxx_messageInfo_Index_FactNames.DiscardUnknown(m)
}

var xxx_messageInfo_Index_FactNames proto.InternalMessageInfo

func (m *Index_FactNames) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Index_FactNames) GetFactName() map[string]FactName {
	if m != nil {
		return m.FactName
	}
	return nil
}

func init() {
	proto.RegisterType((*Index)(nil), "kythe.proto.schema.Index")
	proto.RegisterType((*Index_EdgeKinds)(nil), "kythe.proto.schema.Index.EdgeKinds")
	proto.RegisterMapType((map[string]EdgeKind)(nil), "kythe.proto.schema.Index.EdgeKinds.EdgeKindEntry")
	proto.RegisterType((*Index_NodeKinds)(nil), "kythe.proto.schema.Index.NodeKinds")
	proto.RegisterMapType((map[string]NodeKind)(nil), "kythe.proto.schema.Index.NodeKinds.NodeKindEntry")
	proto.RegisterType((*Index_Subkinds)(nil), "kythe.proto.schema.Index.Subkinds")
	proto.RegisterMapType((map[string]Subkind)(nil), "kythe.proto.schema.Index.Subkinds.SubkindEntry")
	proto.RegisterType((*Index_FactNames)(nil), "kythe.proto.schema.Index.FactNames")
	proto.RegisterMapType((map[string]FactName)(nil), "kythe.proto.schema.Index.FactNames.FactNameEntry")
	proto.RegisterEnum("kythe.proto.schema.NodeKind", NodeKind_name, NodeKind_value)
	proto.RegisterEnum("kythe.proto.schema.Subkind", Subkind_name, Subkind_value)
	proto.RegisterEnum("kythe.proto.schema.FactName", FactName_name, FactName_value)
	proto.RegisterEnum("kythe.proto.schema.EdgeKind", EdgeKind_name, EdgeKind_value)
}

func init() { proto.RegisterFile("kythe/proto/schema.proto", fileDescriptor_schema_91faf3e8973c7842) }

var fileDescriptor_schema_91faf3e8973c7842 = []byte{
	// 1273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x6f, 0x14, 0x37,
	0x17, 0x66, 0x37, 0xd9, 0xec, 0xae, 0x37, 0x1f, 0x27, 0x26, 0x84, 0x25, 0xf0, 0xbe, 0x6f, 0x5e,
	0xe8, 0x47, 0x9a, 0x4a, 0x41, 0xd0, 0x9b, 0xaa, 0x17, 0x95, 0xbc, 0x33, 0x67, 0x17, 0x2b, 0x1e,
	0x7b, 0xb0, 0x3d, 0x21, 0xe9, 0xcd, 0x28, 0x90, 0xe5, 0x43, 0x94, 0x04, 0x91, 0x50, 0xc1, 0x7d,
	0x7f, 0x4c, 0x7f, 0x49, 0x6f, 0x2a, 0xb5, 0xff, 0xa8, 0xaa, 0x8e, 0x67, 0x3c, 0x6c, 0x5a, 0x20,
	0x15, 0x77, 0xe7, 0xd8, 0x3e, 0x8f, 0xcf, 0xf3, 0x9c, 0x67, 0x3c, 0x6c, 0xf8, 0xfc, 0xed, 0xd9,
	0xd3, 0xe9, 0xed, 0x97, 0xaf, 0x4e, 0xce, 0x4e, 0x6e, 0x9f, 0x3e, 0x7a, 0x3a, 0x7d, 0x71, 0xb8,
	0x13, 0x12, 0xce, 0xc3, 0x4e, 0x95, 0xec, 0x54, 0x3b, 0x37, 0xff, 0xe8, 0xb1, 0x8e, 0x3c, 0x3e,
	0x9a, 0xbe, 0xe1, 0x23, 0xc6, 0xa6, 0x47, 0x4f, 0xa6, 0xe5, 0xf3, 0x67, 0xc7, 0x47, 0xa7, 0xc3,
	0xd6, 0xe6, 0xdc, 0xd6, 0xe0, 0xee, 0xad, 0x9d, 0x7f, 0x96, 0xec, 0x84, 0xe3, 0x3b, 0x78, 0xf4,
	0x64, 0xba, 0x4b, 0x47, 0x6d, 0x7f, 0x1a, 0x43, 0xc2, 0x38, 0x3e, 0x39, 0x8a, 0x18, 0xed, 0x8b,
	0x30, 0xf4, 0xc9, 0x51, 0xc4, 0x38, 0x8e, 0x21, 0xff, 0x9e, 0xf5, 0x4e, 0x5f, 0x3f, 0xac, 0x10,
	0xe6, 0x02, 0xc2, 0xcd, 0x0f, 0x23, 0xb8, 0xfa, 0xa4, 0x6d, 0x6a, 0xa8, 0x87, 0xc7, 0x87, 0x8f,
	0xce, 0xca, 0xe3, 0xc3, 0x17, 0xd3, 0xd3, 0xe1, 0xfc, 0x45, 0x3d, 0x8c, 0x0f, 0x1f, 0x9d, 0x69,
	0x3a, 0x6a, 0xfb, 0x8f, 0x63, 0xb8, 0xf1, 0x7b, 0x8b, 0xf5, 0x1b, 0x82, 0x7c, 0x9d, 0x2d, 0xbc,
	0x7c, 0x35, 0x7d, 0xfc, 0xec, 0xcd, 0xb0, 0xb5, 0xd9, 0xda, 0xea, 0xdb, 0x3a, 0xe3, 0x9a, 0xf5,
	0x1b, 0xc5, 0x6a, 0xb2, 0x77, 0xfe, 0x85, 0x60, 0x4d, 0x84, 0xc7, 0x67, 0xaf, 0xde, 0xda, 0x5e,
	0x94, 0x6f, 0xe3, 0x80, 0x2d, 0x9d, 0xdb, 0xe2, 0xc0, 0xe6, 0x9e, 0x4f, 0xdf, 0xd6, 0xb7, 0x52,
	0xc8, 0xef, 0xb2, 0xce, 0x4f, 0x87, 0x3f, 0xbe, 0x9e, 0x0e, 0xdb, 0x9b, 0xad, 0xad, 0xe5, 0xbb,
	0x37, 0xde, 0x77, 0x5d, 0xc4, 0xb0, 0xd5, 0xd1, 0xef, 0xda, 0xdf, 0xb6, 0x02, 0xa1, 0x46, 0xed,
	0x8f, 0x11, 0x6a, 0xc6, 0x77, 0x31, 0xa1, 0x06, 0xaf, 0x89, 0x6a, 0x42, 0x71, 0x96, 0x44, 0xe8,
	0xdc, 0xd6, 0x27, 0x12, 0x8a, 0x18, 0xb3, 0x84, 0x7e, 0x6d, 0xb1, 0x5e, 0x1c, 0xfe, 0x07, 0xf9,
	0x48, 0xd6, 0xad, 0x6d, 0x51, 0xb3, 0xb9, 0x7d, 0xb1, 0x93, 0x62, 0x50, 0x71, 0x89, 0xf5, 0x1b,
	0x0f, 0xd8, 0xe2, 0xec, 0xc6, 0x7b, 0x98, 0xdc, 0x39, 0xcf, 0xe4, 0xfa, 0xfb, 0xae, 0xaa, 0x21,
	0xfe, 0x3e, 0x99, 0xc6, 0x83, 0x1f, 0x9b, 0x4c, 0x63, 0xea, 0x8b, 0x27, 0xd3, 0xe0, 0x35, 0x51,
	0x3d, 0x99, 0xe8, 0x70, 0x9a, 0xcc, 0xb9, 0xad, 0x4f, 0x9c, 0x4c, 0xc4, 0x98, 0x21, 0xb4, 0xfd,
	0x5b, 0x9b, 0xf5, 0xe2, 0xc4, 0xf8, 0x15, 0xb6, 0x5a, 0xe8, 0x5d, 0x6d, 0x1e, 0xe8, 0x52, 0x9b,
	0x14, 0xcb, 0x5d, 0xa9, 0x53, 0xb8, 0xc4, 0xbb, 0x6c, 0x4e, 0x8c, 0x1c, 0xb4, 0x38, 0x63, 0x0b,
	0x62, 0xe4, 0xf6, 0x84, 0x85, 0x76, 0x88, 0x75, 0x72, 0xcf, 0x58, 0x98, 0xe3, 0x8b, 0xac, 0x97,
	0x18, 0xed, 0xbc, 0xd0, 0x1e, 0xe6, 0xf9, 0x32, 0x63, 0xa9, 0x14, 0x13, 0x6d, 0x9c, 0x97, 0x09,
	0x74, 0xa8, 0x3c, 0x35, 0x09, 0x2c, 0xf0, 0x1e, 0x9b, 0x1f, 0x4b, 0x85, 0xd0, 0xe5, 0x4b, 0xac,
	0x2f, 0xb5, 0x47, 0x3b, 0x16, 0x09, 0x42, 0x8f, 0xea, 0xc7, 0x85, 0x4e, 0xbc, 0x34, 0x1a, 0xfa,
	0x84, 0xac, 0x8c, 0xd9, 0x2d, 0x72, 0x60, 0xbc, 0xcf, 0x3a, 0x99, 0x48, 0xac, 0x81, 0x01, 0x55,
	0x67, 0xe8, 0x05, 0x2c, 0x52, 0xa4, 0x45, 0x86, 0xb0, 0xc4, 0x07, 0xac, 0x9b, 0x8b, 0x64, 0x57,
	0x4c, 0x10, 0x96, 0x43, 0x62, 0x4d, 0x82, 0xce, 0xc1, 0x0a, 0x81, 0x58, 0x4c, 0x8c, 0x4d, 0x01,
	0xa8, 0x01, 0x57, 0x64, 0xb0, 0x4a, 0x8b, 0xee, 0x20, 0x1b, 0x19, 0x05, 0x9c, 0x62, 0x2f, 0x94,
	0x14, 0x0e, 0x2e, 0x13, 0xa0, 0x17, 0x79, 0x0e, 0x6b, 0xd4, 0x89, 0x1f, 0x15, 0x52, 0x79, 0xa9,
	0xe1, 0x4a, 0xc8, 0xb4, 0xc9, 0xa4, 0x16, 0x0a, 0xd6, 0x43, 0x85, 0x93, 0x93, 0x4c, 0xc0, 0x55,
	0xda, 0xd9, 0x13, 0x56, 0x8a, 0x91, 0x42, 0x18, 0xd2, 0x05, 0x7b, 0x89, 0x83, 0x6b, 0xdb, 0x3f,
	0xb7, 0x59, 0xb7, 0x76, 0x0d, 0xbf, 0xcc, 0x56, 0xa2, 0x98, 0xae, 0x18, 0xd5, 0x52, 0x92, 0x52,
	0xc2, 0xe3, 0xc4, 0xd8, 0x03, 0x68, 0x11, 0xbb, 0x44, 0x09, 0xe7, 0xa0, 0xcd, 0x57, 0xd8, 0x20,
	0x48, 0x68, 0x8b, 0xc4, 0x07, 0x4d, 0x49, 0x45, 0x6c, 0xf2, 0x79, 0xea, 0x11, 0x75, 0x91, 0x41,
	0x87, 0x76, 0x28, 0x2a, 0xab, 0xd2, 0x05, 0x42, 0x19, 0x4b, 0x54, 0x29, 0x74, 0x09, 0x5e, 0x66,
	0xb9, 0x92, 0x89, 0xf4, 0xd0, 0xa3, 0x86, 0x65, 0x96, 0x1b, 0xeb, 0xa1, 0x4f, 0xf8, 0x52, 0x4b,
	0x2f, 0x85, 0x92, 0x3f, 0xa0, 0xad, 0x94, 0x55, 0x26, 0x11, 0x0a, 0x06, 0xd4, 0x69, 0x08, 0xcb,
	0x5c, 0x58, 0x91, 0xa1, 0x47, 0x0b, 0x8b, 0x54, 0x9c, 0xa1, 0xbf, 0x67, 0x52, 0x58, 0xa2, 0x71,
	0x91, 0xe0, 0x2e, 0xa7, 0x71, 0x2d, 0x07, 0x19, 0x43, 0x63, 0xb0, 0x12, 0xa4, 0x3b, 0xc8, 0x11,
	0x80, 0x00, 0x0b, 0x4d, 0x13, 0x5c, 0xdd, 0xfe, 0xa5, 0xcd, 0x7a, 0xd1, 0x6c, 0xb3, 0xa6, 0x1a,
	0x8b, 0xc4, 0x97, 0x61, 0x74, 0x97, 0xa8, 0x30, 0x31, 0x29, 0x42, 0xab, 0x72, 0x4f, 0x96, 0x2b,
	0xf4, 0xd8, 0x08, 0xe1, 0x71, 0xdf, 0x97, 0x85, 0x55, 0x30, 0x47, 0x63, 0x4d, 0xd1, 0x0b, 0xa9,
	0x1c, 0xcc, 0x87, 0xc4, 0x24, 0x65, 0x61, 0x25, 0x74, 0x02, 0x05, 0x31, 0x42, 0x05, 0x0b, 0xb4,
	0xae, 0x4c, 0x52, 0xa2, 0x4e, 0x2b, 0x77, 0x51, 0xe2, 0xbc, 0xb0, 0x24, 0xc3, 0x80, 0x75, 0x33,
	0x74, 0x8e, 0x4c, 0xd2, 0x0f, 0x54, 0x1a, 0x6b, 0x33, 0xbe, 0xca, 0x96, 0x02, 0xe9, 0x32, 0xc5,
	0xb1, 0x28, 0x94, 0x87, 0x01, 0xc9, 0x6b, 0x0b, 0x85, 0xb5, 0xbc, 0x8b, 0xd4, 0x90, 0xd3, 0x32,
	0xcf, 0xd1, 0x07, 0xf8, 0x25, 0xaa, 0x89, 0x0b, 0xd5, 0x15, 0xc1, 0x7a, 0x71, 0xc6, 0x95, 0x24,
	0xb8, 0xef, 0x01, 0xe8, 0x64, 0x20, 0x82, 0x3a, 0x31, 0xa9, 0xd4, 0x13, 0x58, 0x25, 0xf4, 0x3d,
	0xe9, 0xe4, 0x48, 0x2a, 0xe9, 0x0f, 0x80, 0x6f, 0xff, 0xd9, 0x61, 0xbd, 0xf8, 0x0b, 0x98, 0x95,
	0x0a, 0xd3, 0x49, 0xf3, 0xfd, 0x0d, 0x58, 0x37, 0x38, 0x15, 0xe9, 0x1b, 0x04, 0xb6, 0x58, 0x27,
	0xa5, 0x35, 0xc6, 0x43, 0x3b, 0xac, 0x68, 0x6d, 0xbc, 0xf0, 0x98, 0x96, 0xa3, 0x03, 0x98, 0xa3,
	0x7b, 0x47, 0xa6, 0xd0, 0x29, 0xa6, 0xa5, 0x32, 0x0f, 0x90, 0xec, 0x33, 0xb3, 0x54, 0xe4, 0x39,
	0x5a, 0xe8, 0x04, 0xdd, 0xef, 0x49, 0x95, 0x96, 0x66, 0x0c, 0x0b, 0x7c, 0x8d, 0x41, 0xcc, 0xca,
	0x7a, 0x00, 0x95, 0x94, 0x71, 0x36, 0x0e, 0x7a, 0x7c, 0x9d, 0xf1, 0x26, 0x2d, 0x0b, 0x2d, 0xef,
	0x17, 0xa8, 0x0e, 0xa0, 0x5f, 0xcd, 0x68, 0x2c, 0x35, 0x3a, 0x60, 0x64, 0xa7, 0x3a, 0x29, 0x47,
	0x52, 0x07, 0xde, 0x83, 0xea, 0x44, 0x8e, 0x3a, 0x25, 0x49, 0x97, 0x58, 0x3f, 0x35, 0x49, 0x91,
	0xa1, 0xf6, 0xae, 0xfa, 0x8a, 0x71, 0x9f, 0x7c, 0xea, 0x2a, 0x29, 0x71, 0xdf, 0x87, 0x83, 0x2b,
	0x74, 0x70, 0x82, 0x1a, 0xad, 0xa0, 0xeb, 0x81, 0xf6, 0x64, 0x96, 0x17, 0x94, 0xac, 0x12, 0x6d,
	0x59, 0xbd, 0x39, 0x32, 0x6c, 0x73, 0x7e, 0x83, 0x0d, 0x67, 0x57, 0x4a, 0x97, 0x63, 0x52, 0x28,
	0xe1, 0xe5, 0x1e, 0xc2, 0x65, 0x72, 0x0b, 0x59, 0x2f, 0x85, 0x35, 0x82, 0x35, 0x7b, 0x68, 0xad,
	0x4c, 0xd1, 0xc1, 0x15, 0xce, 0xd9, 0x72, 0x93, 0x56, 0xa2, 0xae, 0xf3, 0x21, 0x5b, 0x7b, 0xb7,
	0xe6, 0xad, 0xd0, 0x4e, 0x06, 0x9c, 0xab, 0x84, 0x13, 0x2c, 0x53, 0x7d, 0xf7, 0x16, 0xc7, 0x70,
	0x8d, 0xa4, 0xb4, 0x38, 0x2e, 0x13, 0xa1, 0x14, 0x6c, 0xd0, 0x18, 0x63, 0x56, 0x36, 0x9f, 0xe3,
	0x75, 0xea, 0x9e, 0x96, 0xe9, 0x2d, 0xbc, 0x41, 0xae, 0xa2, 0x04, 0xf7, 0x73, 0x41, 0x54, 0xff,
	0xc3, 0x37, 0xd8, 0xfa, 0xcc, 0xc2, 0xec, 0x95, 0xff, 0x8d, 0xf0, 0xe1, 0xf1, 0xfc, 0x1f, 0x11,
	0xa7, 0xac, 0x41, 0xde, 0x8c, 0x60, 0xd5, 0xc7, 0xee, 0xe0, 0xff, 0xcd, 0x11, 0x9d, 0xa8, 0x82,
	0x38, 0xde, 0x8c, 0x10, 0xf4, 0x06, 0xc0, 0xad, 0xd8, 0x21, 0x65, 0xef, 0x70, 0x3e, 0x8b, 0x38,
	0xf7, 0x0b, 0xb4, 0x12, 0x1d, 0x7c, 0x4e, 0x42, 0x39, 0xe1, 0xa5, 0x1b, 0x53, 0xfa, 0x45, 0xf8,
	0x14, 0x72, 0x4c, 0xaa, 0x47, 0xc4, 0xc1, 0x97, 0xfc, 0x3a, 0xbb, 0x3a, 0xb3, 0x70, 0x4e, 0xf0,
	0xad, 0xea, 0x85, 0x9d, 0x4c, 0x30, 0x85, 0xaf, 0x48, 0x34, 0x7a, 0x26, 0x52, 0xd8, 0x26, 0x4c,
	0xf2, 0x61, 0xe5, 0x96, 0xaf, 0x1f, 0x2e, 0x84, 0x5f, 0xd4, 0x37, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x2d, 0x13, 0xa4, 0x51, 0x09, 0x0b, 0x00, 0x00,
}
