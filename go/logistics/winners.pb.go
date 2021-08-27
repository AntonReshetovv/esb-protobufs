// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.1
// source: proto/logistics/winners.proto

package logistics

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WinnersParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryIsoCode         string `protobuf:"bytes,1,opt,name=country_iso_code,json=countryIsoCode,proto3" json:"country_iso_code,omitempty"`
	GeoId                  string `protobuf:"bytes,2,opt,name=geo_id,json=geoId,proto3" json:"geo_id,omitempty"`
	OrderPrice             int32  `protobuf:"varint,3,opt,name=order_price,json=orderPrice,proto3" json:"order_price,omitempty"`
	OrderPositionsQuantity int32  `protobuf:"varint,4,opt,name=order_positions_quantity,json=orderPositionsQuantity,proto3" json:"order_positions_quantity,omitempty"`
}

func (x *WinnersParams) Reset() {
	*x = WinnersParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_winners_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinnersParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinnersParams) ProtoMessage() {}

func (x *WinnersParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_winners_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinnersParams.ProtoReflect.Descriptor instead.
func (*WinnersParams) Descriptor() ([]byte, []int) {
	return file_proto_logistics_winners_proto_rawDescGZIP(), []int{0}
}

func (x *WinnersParams) GetCountryIsoCode() string {
	if x != nil {
		return x.CountryIsoCode
	}
	return ""
}

func (x *WinnersParams) GetGeoId() string {
	if x != nil {
		return x.GeoId
	}
	return ""
}

func (x *WinnersParams) GetOrderPrice() int32 {
	if x != nil {
		return x.OrderPrice
	}
	return 0
}

func (x *WinnersParams) GetOrderPositionsQuantity() int32 {
	if x != nil {
		return x.OrderPositionsQuantity
	}
	return 0
}

type WinnersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Winner `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *WinnersResponse) Reset() {
	*x = WinnersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_winners_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinnersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinnersResponse) ProtoMessage() {}

func (x *WinnersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_winners_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinnersResponse.ProtoReflect.Descriptor instead.
func (*WinnersResponse) Descriptor() ([]byte, []int) {
	return file_proto_logistics_winners_proto_rawDescGZIP(), []int{1}
}

func (x *WinnersResponse) GetResults() []*Winner {
	if x != nil {
		return x.Results
	}
	return nil
}

type Winner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryMethodId   int32                  `protobuf:"varint,1,opt,name=delivery_method_id,json=deliveryMethodId,proto3" json:"delivery_method_id,omitempty"`
	DeliveryMethodName string                 `protobuf:"bytes,2,opt,name=delivery_method_name,json=deliveryMethodName,proto3" json:"delivery_method_name,omitempty"`
	WinnerCompany      *WinnerDeliveryCompany `protobuf:"bytes,3,opt,name=winner_company,json=winnerCompany,proto3" json:"winner_company,omitempty"`
	PaymentMethods     []*WinnerPaymentMethod `protobuf:"bytes,4,rep,name=payment_methods,json=paymentMethods,proto3" json:"payment_methods,omitempty"`
	DeliveryDate       []*WinnerDeliveryDate  `protobuf:"bytes,5,rep,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	Message            *WinnerMessage         `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Winner) Reset() {
	*x = Winner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_winners_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Winner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Winner) ProtoMessage() {}

func (x *Winner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_winners_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Winner.ProtoReflect.Descriptor instead.
func (*Winner) Descriptor() ([]byte, []int) {
	return file_proto_logistics_winners_proto_rawDescGZIP(), []int{2}
}

func (x *Winner) GetDeliveryMethodId() int32 {
	if x != nil {
		return x.DeliveryMethodId
	}
	return 0
}

func (x *Winner) GetDeliveryMethodName() string {
	if x != nil {
		return x.DeliveryMethodName
	}
	return ""
}

func (x *Winner) GetWinnerCompany() *WinnerDeliveryCompany {
	if x != nil {
		return x.WinnerCompany
	}
	return nil
}

func (x *Winner) GetPaymentMethods() []*WinnerPaymentMethod {
	if x != nil {
		return x.PaymentMethods
	}
	return nil
}

func (x *Winner) GetDeliveryDate() []*WinnerDeliveryDate {
	if x != nil {
		return x.DeliveryDate
	}
	return nil
}

func (x *Winner) GetMessage() *WinnerMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type WinnerDeliveryCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WinnerDeliveryCompany) Reset() {
	*x = WinnerDeliveryCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_winners_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinnerDeliveryCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinnerDeliveryCompany) ProtoMessage() {}

func (x *WinnerDeliveryCompany) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_winners_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinnerDeliveryCompany.ProtoReflect.Descriptor instead.
func (*WinnerDeliveryCompany) Descriptor() ([]byte, []int) {
	return file_proto_logistics_winners_proto_rawDescGZIP(), []int{3}
}

func (x *WinnerDeliveryCompany) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *WinnerDeliveryCompany) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WinnerPaymentMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DeliveryPrice int32          `protobuf:"varint,3,opt,name=delivery_price,json=deliveryPrice,proto3" json:"delivery_price,omitempty"`
	Message       *WinnerMessage `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WinnerPaymentMethod) Reset() {
	*x = WinnerPaymentMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_winners_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinnerPaymentMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinnerPaymentMethod) ProtoMessage() {}

func (x *WinnerPaymentMethod) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_winners_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinnerPaymentMethod.ProtoReflect.Descriptor instead.
func (*WinnerPaymentMethod) Descriptor() ([]byte, []int) {
	return file_proto_logistics_winners_proto_rawDescGZIP(), []int{4}
}

func (x *WinnerPaymentMethod) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WinnerPaymentMethod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WinnerPaymentMethod) GetDeliveryPrice() int32 {
	if x != nil {
		return x.DeliveryPrice
	}
	return 0
}

func (x *WinnerPaymentMethod) GetMessage() *WinnerMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type WinnerDeliveryDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date      string                    `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Intervals []*WinnerDeliveryInterval `protobuf:"bytes,2,rep,name=intervals,proto3" json:"intervals,omitempty"`
}

func (x *WinnerDeliveryDate) Reset() {
	*x = WinnerDeliveryDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_winners_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinnerDeliveryDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinnerDeliveryDate) ProtoMessage() {}

func (x *WinnerDeliveryDate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_winners_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinnerDeliveryDate.ProtoReflect.Descriptor instead.
func (*WinnerDeliveryDate) Descriptor() ([]byte, []int) {
	return file_proto_logistics_winners_proto_rawDescGZIP(), []int{5}
}

func (x *WinnerDeliveryDate) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *WinnerDeliveryDate) GetIntervals() []*WinnerDeliveryInterval {
	if x != nil {
		return x.Intervals
	}
	return nil
}

type WinnerDeliveryInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeFrom      string `protobuf:"bytes,1,opt,name=time_from,json=timeFrom,proto3" json:"time_from,omitempty"`
	TimeTo        string `protobuf:"bytes,2,opt,name=time_to,json=timeTo,proto3" json:"time_to,omitempty"`
	DayOfWeekType string `protobuf:"bytes,3,opt,name=day_of_week_type,json=dayOfWeekType,proto3" json:"day_of_week_type,omitempty"`
}

func (x *WinnerDeliveryInterval) Reset() {
	*x = WinnerDeliveryInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_winners_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinnerDeliveryInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinnerDeliveryInterval) ProtoMessage() {}

func (x *WinnerDeliveryInterval) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_winners_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinnerDeliveryInterval.ProtoReflect.Descriptor instead.
func (*WinnerDeliveryInterval) Descriptor() ([]byte, []int) {
	return file_proto_logistics_winners_proto_rawDescGZIP(), []int{6}
}

func (x *WinnerDeliveryInterval) GetTimeFrom() string {
	if x != nil {
		return x.TimeFrom
	}
	return ""
}

func (x *WinnerDeliveryInterval) GetTimeTo() string {
	if x != nil {
		return x.TimeTo
	}
	return ""
}

func (x *WinnerDeliveryInterval) GetDayOfWeekType() string {
	if x != nil {
		return x.DayOfWeekType
	}
	return ""
}

type WinnerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PriceBorder           int32  `protobuf:"varint,1,opt,name=price_border,json=priceBorder,proto3" json:"price_border,omitempty"`
	ProductQuantityBorder int32  `protobuf:"varint,2,opt,name=product_quantity_border,json=productQuantityBorder,proto3" json:"product_quantity_border,omitempty"`
	Text                  string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	ShortText             string `protobuf:"bytes,4,opt,name=short_text,json=shortText,proto3" json:"short_text,omitempty"`
	Header                string `protobuf:"bytes,5,opt,name=header,proto3" json:"header,omitempty"`
	ButtonText            string `protobuf:"bytes,6,opt,name=button_text,json=buttonText,proto3" json:"button_text,omitempty"`
}

func (x *WinnerMessage) Reset() {
	*x = WinnerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_winners_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinnerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinnerMessage) ProtoMessage() {}

func (x *WinnerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_winners_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinnerMessage.ProtoReflect.Descriptor instead.
func (*WinnerMessage) Descriptor() ([]byte, []int) {
	return file_proto_logistics_winners_proto_rawDescGZIP(), []int{7}
}

func (x *WinnerMessage) GetPriceBorder() int32 {
	if x != nil {
		return x.PriceBorder
	}
	return 0
}

func (x *WinnerMessage) GetProductQuantityBorder() int32 {
	if x != nil {
		return x.ProductQuantityBorder
	}
	return 0
}

func (x *WinnerMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *WinnerMessage) GetShortText() string {
	if x != nil {
		return x.ShortText
	}
	return ""
}

func (x *WinnerMessage) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *WinnerMessage) GetButtonText() string {
	if x != nil {
		return x.ButtonText
	}
	return ""
}

var File_proto_logistics_winners_proto protoreflect.FileDescriptor

var file_proto_logistics_winners_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x57, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x73, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x73, 0x6f,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x67, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x18,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3e, 0x0a, 0x0f, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xf2, 0x02, 0x0a, 0x06, 0x57, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x47, 0x0a, 0x0e, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x0d, 0x77, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x47, 0x0a, 0x0f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x12, 0x42, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x15, 0x57,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a,
	0x13, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x32, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x57, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x12, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a,
	0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x57, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x77,
	0x0a, 0x16, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x12, 0x27,
	0x0a, 0x10, 0x64, 0x61, 0x79, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x57,
	0x65, 0x65, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x0d, 0x57, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x17,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74,
	0x32, 0x75, 0x0a, 0x07, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x6a, 0x0a, 0x10, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x57, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2d, 0x77, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x6f, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_logistics_winners_proto_rawDescOnce sync.Once
	file_proto_logistics_winners_proto_rawDescData = file_proto_logistics_winners_proto_rawDesc
)

func file_proto_logistics_winners_proto_rawDescGZIP() []byte {
	file_proto_logistics_winners_proto_rawDescOnce.Do(func() {
		file_proto_logistics_winners_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_logistics_winners_proto_rawDescData)
	})
	return file_proto_logistics_winners_proto_rawDescData
}

var file_proto_logistics_winners_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_logistics_winners_proto_goTypes = []interface{}{
	(*WinnersParams)(nil),          // 0: logistics.WinnersParams
	(*WinnersResponse)(nil),        // 1: logistics.WinnersResponse
	(*Winner)(nil),                 // 2: logistics.Winner
	(*WinnerDeliveryCompany)(nil),  // 3: logistics.WinnerDeliveryCompany
	(*WinnerPaymentMethod)(nil),    // 4: logistics.WinnerPaymentMethod
	(*WinnerDeliveryDate)(nil),     // 5: logistics.WinnerDeliveryDate
	(*WinnerDeliveryInterval)(nil), // 6: logistics.WinnerDeliveryInterval
	(*WinnerMessage)(nil),          // 7: logistics.WinnerMessage
}
var file_proto_logistics_winners_proto_depIdxs = []int32{
	2, // 0: logistics.WinnersResponse.results:type_name -> logistics.Winner
	3, // 1: logistics.Winner.winner_company:type_name -> logistics.WinnerDeliveryCompany
	4, // 2: logistics.Winner.payment_methods:type_name -> logistics.WinnerPaymentMethod
	5, // 3: logistics.Winner.delivery_date:type_name -> logistics.WinnerDeliveryDate
	7, // 4: logistics.Winner.message:type_name -> logistics.WinnerMessage
	7, // 5: logistics.WinnerPaymentMethod.message:type_name -> logistics.WinnerMessage
	6, // 6: logistics.WinnerDeliveryDate.intervals:type_name -> logistics.WinnerDeliveryInterval
	0, // 7: logistics.Winners.CalculateWinners:input_type -> logistics.WinnersParams
	1, // 8: logistics.Winners.CalculateWinners:output_type -> logistics.WinnersResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_logistics_winners_proto_init() }
func file_proto_logistics_winners_proto_init() {
	if File_proto_logistics_winners_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_logistics_winners_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinnersParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_winners_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinnersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_winners_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Winner); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_winners_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinnerDeliveryCompany); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_winners_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinnerPaymentMethod); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_winners_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinnerDeliveryDate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_winners_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinnerDeliveryInterval); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_winners_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinnerMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_logistics_winners_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_logistics_winners_proto_goTypes,
		DependencyIndexes: file_proto_logistics_winners_proto_depIdxs,
		MessageInfos:      file_proto_logistics_winners_proto_msgTypes,
	}.Build()
	File_proto_logistics_winners_proto = out.File
	file_proto_logistics_winners_proto_rawDesc = nil
	file_proto_logistics_winners_proto_goTypes = nil
	file_proto_logistics_winners_proto_depIdxs = nil
}
