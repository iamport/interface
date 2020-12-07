// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: v2/basis/basis.proto

package basis

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CancelHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PgTid       string `protobuf:"bytes,1,opt,name=pg_tid,json=pgTid,proto3" json:"pg_tid,omitempty"`
	Amount      string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	CancelledAt int32  `protobuf:"varint,3,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
	Reason      string `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	ReceiptUrl  string `protobuf:"bytes,5,opt,name=receipt_url,json=receiptUrl,proto3" json:"receipt_url,omitempty"`
}

func (x *CancelHistory) Reset() {
	*x = CancelHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_basis_basis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelHistory) ProtoMessage() {}

func (x *CancelHistory) ProtoReflect() protoreflect.Message {
	mi := &file_v2_basis_basis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelHistory.ProtoReflect.Descriptor instead.
func (*CancelHistory) Descriptor() ([]byte, []int) {
	return file_v2_basis_basis_proto_rawDescGZIP(), []int{0}
}

func (x *CancelHistory) GetPgTid() string {
	if x != nil {
		return x.PgTid
	}
	return ""
}

func (x *CancelHistory) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *CancelHistory) GetCancelledAt() int32 {
	if x != nil {
		return x.CancelledAt
	}
	return 0
}

func (x *CancelHistory) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *CancelHistory) GetReceiptUrl() string {
	if x != nil {
		return x.ReceiptUrl
	}
	return ""
}

type UnitTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount            int32            `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	ApplyNum          string           `protobuf:"bytes,2,opt,name=apply_num,json=applyNum,proto3" json:"apply_num,omitempty"`
	BankCode          int32            `protobuf:"varint,3,opt,name=bank_code,json=bankCode,proto3" json:"bank_code,omitempty"`
	BankName          string           `protobuf:"bytes,4,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	BuyerAddr         string           `protobuf:"bytes,5,opt,name=buyer_addr,json=buyerAddr,proto3" json:"buyer_addr,omitempty"`
	BuyerEmail        string           `protobuf:"bytes,6,opt,name=buyer_email,json=buyerEmail,proto3" json:"buyer_email,omitempty"`
	BuyerName         string           `protobuf:"bytes,7,opt,name=buyer_name,json=buyerName,proto3" json:"buyer_name,omitempty"`
	BuyerPostcode     string           `protobuf:"bytes,8,opt,name=buyer_postcode,json=buyerPostcode,proto3" json:"buyer_postcode,omitempty"`
	BuyerTel          string           `protobuf:"bytes,9,opt,name=buyer_tel,json=buyerTel,proto3" json:"buyer_tel,omitempty"`
	CancelAmount      int32            `protobuf:"varint,10,opt,name=cancel_amount,json=cancelAmount,proto3" json:"cancel_amount,omitempty"`
	CancelHistory     []*CancelHistory `protobuf:"bytes,11,rep,name=cancel_history,json=cancelHistory,proto3" json:"cancel_history,omitempty"`
	CancelReason      string           `protobuf:"bytes,12,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason,omitempty"`
	CancelReceiptUrls []string         `protobuf:"bytes,13,rep,name=cancel_receipt_urls,json=cancelReceiptUrls,proto3" json:"cancel_receipt_urls,omitempty"`
	CancelledAt       int32            `protobuf:"varint,14,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
	CardCode          string           `protobuf:"bytes,15,opt,name=card_code,json=cardCode,proto3" json:"card_code,omitempty"`
	CardName          string           `protobuf:"bytes,16,opt,name=card_name,json=cardName,proto3" json:"card_name,omitempty"`
	BinNumber         string           `protobuf:"bytes,17,opt,name=bin_number,json=binNumber,proto3" json:"bin_number,omitempty"`
	CardInstallment   int32            `protobuf:"varint,18,opt,name=card_installment,json=cardInstallment,proto3" json:"card_installment,omitempty"`
	CardType          int32            `protobuf:"varint,19,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	CashReceiptIssued bool             `protobuf:"varint,20,opt,name=cash_receipt_issued,json=cashReceiptIssued,proto3" json:"cash_receipt_issued,omitempty"`
	Channel           string           `protobuf:"bytes,21,opt,name=channel,proto3" json:"channel,omitempty"`
	Currency          string           `protobuf:"bytes,22,opt,name=currency,proto3" json:"currency,omitempty"`
	CustomData        string           `protobuf:"bytes,23,opt,name=custom_data,json=customData,proto3" json:"custom_data,omitempty"`
	CardUid           string           `protobuf:"bytes,24,opt,name=card_uid,json=cardUid,proto3" json:"card_uid,omitempty"`
	CustomerUidUsage  string           `protobuf:"bytes,25,opt,name=customer_uid_usage,json=customerUidUsage,proto3" json:"customer_uid_usage,omitempty"`
	Escrow            bool             `protobuf:"varint,26,opt,name=escrow,proto3" json:"escrow,omitempty"`
	FailReason        string           `protobuf:"bytes,27,opt,name=fail_reason,json=failReason,proto3" json:"fail_reason,omitempty"`
	FailedAt          int32            `protobuf:"varint,28,opt,name=failed_at,json=failedAt,proto3" json:"failed_at,omitempty"`
	ImpUid            string           `protobuf:"bytes,29,opt,name=imp_uid,json=impUid,proto3" json:"imp_uid,omitempty"`
	MerchantUid       string           `protobuf:"bytes,30,opt,name=merchant_uid,json=merchantUid,proto3" json:"merchant_uid,omitempty"`
	OrderName         string           `protobuf:"bytes,31,opt,name=order_name,json=orderName,proto3" json:"order_name,omitempty"`
	PaidAt            int32            `protobuf:"varint,32,opt,name=paid_at,json=paidAt,proto3" json:"paid_at,omitempty"`
	PayMethod         string           `protobuf:"bytes,33,opt,name=pay_method,json=payMethod,proto3" json:"pay_method,omitempty"`
	PgId              string           `protobuf:"bytes,34,opt,name=pg_id,json=pgId,proto3" json:"pg_id,omitempty"`
	PgProvider        string           `protobuf:"bytes,35,opt,name=pg_provider,json=pgProvider,proto3" json:"pg_provider,omitempty"`
	PgTid             string           `protobuf:"bytes,36,opt,name=pg_tid,json=pgTid,proto3" json:"pg_tid,omitempty"`
	ReceiptUrl        string           `protobuf:"bytes,37,opt,name=receipt_url,json=receiptUrl,proto3" json:"receipt_url,omitempty"`
	StartedAt         int32            `protobuf:"varint,38,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	Status            string           `protobuf:"bytes,39,opt,name=status,proto3" json:"status,omitempty"`
	UserAgent         string           `protobuf:"bytes,40,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	VbankCode         string           `protobuf:"bytes,41,opt,name=vbank_code,json=vbankCode,proto3" json:"vbank_code,omitempty"`
	VbankDate         int32            `protobuf:"varint,42,opt,name=vbank_date,json=vbankDate,proto3" json:"vbank_date,omitempty"`
	VbankHolder       string           `protobuf:"bytes,43,opt,name=vbank_holder,json=vbankHolder,proto3" json:"vbank_holder,omitempty"`
	VbankIssuedAt     int32            `protobuf:"varint,44,opt,name=vbank_issued_at,json=vbankIssuedAt,proto3" json:"vbank_issued_at,omitempty"`
	VbankName         string           `protobuf:"bytes,45,opt,name=vbank_name,json=vbankName,proto3" json:"vbank_name,omitempty"`
	VbankNum          string           `protobuf:"bytes,46,opt,name=vbank_num,json=vbankNum,proto3" json:"vbank_num,omitempty"`
	CustomerEmail     string           `protobuf:"bytes,47,opt,name=customer_email,json=customerEmail,proto3" json:"customer_email,omitempty"`
}

func (x *UnitTx) Reset() {
	*x = UnitTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_basis_basis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitTx) ProtoMessage() {}

func (x *UnitTx) ProtoReflect() protoreflect.Message {
	mi := &file_v2_basis_basis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitTx.ProtoReflect.Descriptor instead.
func (*UnitTx) Descriptor() ([]byte, []int) {
	return file_v2_basis_basis_proto_rawDescGZIP(), []int{1}
}

func (x *UnitTx) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UnitTx) GetApplyNum() string {
	if x != nil {
		return x.ApplyNum
	}
	return ""
}

func (x *UnitTx) GetBankCode() int32 {
	if x != nil {
		return x.BankCode
	}
	return 0
}

func (x *UnitTx) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *UnitTx) GetBuyerAddr() string {
	if x != nil {
		return x.BuyerAddr
	}
	return ""
}

func (x *UnitTx) GetBuyerEmail() string {
	if x != nil {
		return x.BuyerEmail
	}
	return ""
}

func (x *UnitTx) GetBuyerName() string {
	if x != nil {
		return x.BuyerName
	}
	return ""
}

func (x *UnitTx) GetBuyerPostcode() string {
	if x != nil {
		return x.BuyerPostcode
	}
	return ""
}

func (x *UnitTx) GetBuyerTel() string {
	if x != nil {
		return x.BuyerTel
	}
	return ""
}

func (x *UnitTx) GetCancelAmount() int32 {
	if x != nil {
		return x.CancelAmount
	}
	return 0
}

func (x *UnitTx) GetCancelHistory() []*CancelHistory {
	if x != nil {
		return x.CancelHistory
	}
	return nil
}

func (x *UnitTx) GetCancelReason() string {
	if x != nil {
		return x.CancelReason
	}
	return ""
}

func (x *UnitTx) GetCancelReceiptUrls() []string {
	if x != nil {
		return x.CancelReceiptUrls
	}
	return nil
}

func (x *UnitTx) GetCancelledAt() int32 {
	if x != nil {
		return x.CancelledAt
	}
	return 0
}

func (x *UnitTx) GetCardCode() string {
	if x != nil {
		return x.CardCode
	}
	return ""
}

func (x *UnitTx) GetCardName() string {
	if x != nil {
		return x.CardName
	}
	return ""
}

func (x *UnitTx) GetBinNumber() string {
	if x != nil {
		return x.BinNumber
	}
	return ""
}

func (x *UnitTx) GetCardInstallment() int32 {
	if x != nil {
		return x.CardInstallment
	}
	return 0
}

func (x *UnitTx) GetCardType() int32 {
	if x != nil {
		return x.CardType
	}
	return 0
}

func (x *UnitTx) GetCashReceiptIssued() bool {
	if x != nil {
		return x.CashReceiptIssued
	}
	return false
}

func (x *UnitTx) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *UnitTx) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *UnitTx) GetCustomData() string {
	if x != nil {
		return x.CustomData
	}
	return ""
}

func (x *UnitTx) GetCardUid() string {
	if x != nil {
		return x.CardUid
	}
	return ""
}

func (x *UnitTx) GetCustomerUidUsage() string {
	if x != nil {
		return x.CustomerUidUsage
	}
	return ""
}

func (x *UnitTx) GetEscrow() bool {
	if x != nil {
		return x.Escrow
	}
	return false
}

func (x *UnitTx) GetFailReason() string {
	if x != nil {
		return x.FailReason
	}
	return ""
}

func (x *UnitTx) GetFailedAt() int32 {
	if x != nil {
		return x.FailedAt
	}
	return 0
}

func (x *UnitTx) GetImpUid() string {
	if x != nil {
		return x.ImpUid
	}
	return ""
}

func (x *UnitTx) GetMerchantUid() string {
	if x != nil {
		return x.MerchantUid
	}
	return ""
}

func (x *UnitTx) GetOrderName() string {
	if x != nil {
		return x.OrderName
	}
	return ""
}

func (x *UnitTx) GetPaidAt() int32 {
	if x != nil {
		return x.PaidAt
	}
	return 0
}

func (x *UnitTx) GetPayMethod() string {
	if x != nil {
		return x.PayMethod
	}
	return ""
}

func (x *UnitTx) GetPgId() string {
	if x != nil {
		return x.PgId
	}
	return ""
}

func (x *UnitTx) GetPgProvider() string {
	if x != nil {
		return x.PgProvider
	}
	return ""
}

func (x *UnitTx) GetPgTid() string {
	if x != nil {
		return x.PgTid
	}
	return ""
}

func (x *UnitTx) GetReceiptUrl() string {
	if x != nil {
		return x.ReceiptUrl
	}
	return ""
}

func (x *UnitTx) GetStartedAt() int32 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *UnitTx) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UnitTx) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *UnitTx) GetVbankCode() string {
	if x != nil {
		return x.VbankCode
	}
	return ""
}

func (x *UnitTx) GetVbankDate() int32 {
	if x != nil {
		return x.VbankDate
	}
	return 0
}

func (x *UnitTx) GetVbankHolder() string {
	if x != nil {
		return x.VbankHolder
	}
	return ""
}

func (x *UnitTx) GetVbankIssuedAt() int32 {
	if x != nil {
		return x.VbankIssuedAt
	}
	return 0
}

func (x *UnitTx) GetVbankName() string {
	if x != nil {
		return x.VbankName
	}
	return ""
}

func (x *UnitTx) GetVbankNum() string {
	if x != nil {
		return x.VbankNum
	}
	return ""
}

func (x *UnitTx) GetCustomerEmail() string {
	if x != nil {
		return x.CustomerEmail
	}
	return ""
}

var File_v2_basis_basis_proto protoreflect.FileDescriptor

var file_v2_basis_basis_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x32, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x61, 0x73, 0x69, 0x73, 0x5f, 0x76, 0x32,
	0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x67, 0x5f, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x67, 0x54, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x81, 0x0c,
	0x0a, 0x06, 0x55, 0x6e, 0x69, 0x74, 0x54, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61,
	0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x79,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x79,
	0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x79,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x75, 0x79, 0x65, 0x72, 0x54, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3e, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73, 0x5f,
	0x76, 0x32, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x0d, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x11, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x55, 0x72, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x64,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x63, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x61, 0x73, 0x68,
	0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x61, 0x73, 0x68, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x64, 0x55, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x55, 0x69, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x63, 0x72,
	0x6f, 0x77, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x6d, 0x70, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6d, 0x70, 0x55, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x69,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x69, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x67, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x67, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x67, 0x5f, 0x74, 0x69,
	0x64, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x67, 0x54, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x25, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x26, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x62, 0x61, 0x6e, 0x6b,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x62, 0x61, 0x6e, 0x6b,
	0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x5f,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x2d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x76, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x2f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x3d, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x5f, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f,
	0x62, 0x61, 0x73, 0x69, 0x73, 0xaa, 0x02, 0x08, 0x56, 0x32, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_basis_basis_proto_rawDescOnce sync.Once
	file_v2_basis_basis_proto_rawDescData = file_v2_basis_basis_proto_rawDesc
)

func file_v2_basis_basis_proto_rawDescGZIP() []byte {
	file_v2_basis_basis_proto_rawDescOnce.Do(func() {
		file_v2_basis_basis_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_basis_basis_proto_rawDescData)
	})
	return file_v2_basis_basis_proto_rawDescData
}

var file_v2_basis_basis_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v2_basis_basis_proto_goTypes = []interface{}{
	(*CancelHistory)(nil), // 0: basis_v2.CancelHistory
	(*UnitTx)(nil),        // 1: basis_v2.UnitTx
}
var file_v2_basis_basis_proto_depIdxs = []int32{
	0, // 0: basis_v2.UnitTx.cancel_history:type_name -> basis_v2.CancelHistory
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_basis_basis_proto_init() }
func file_v2_basis_basis_proto_init() {
	if File_v2_basis_basis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_basis_basis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelHistory); i {
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
		file_v2_basis_basis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitTx); i {
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
			RawDescriptor: file_v2_basis_basis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_basis_basis_proto_goTypes,
		DependencyIndexes: file_v2_basis_basis_proto_depIdxs,
		MessageInfos:      file_v2_basis_basis_proto_msgTypes,
	}.Build()
	File_v2_basis_basis_proto = out.File
	file_v2_basis_basis_proto_rawDesc = nil
	file_v2_basis_basis_proto_goTypes = nil
	file_v2_basis_basis_proto_depIdxs = nil
}