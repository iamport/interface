<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v2/payments/card/card.proto

namespace GPBMetadata\V2\Payments\Card;

class Card
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Api\Annotations::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
v2/payments/card/card.protocard_v2"�
CardRegisterRequest
card_uid (	
card_number (	
expiry (	
birth (	

pwd_2digit (	
customer_name (	
customer_tel (	
customer_email (	
customer_addr	 (	
customer_postcode
 (	

pg (	"�
CardRegisterResultData
card_uid (	
	card_name (	
	card_code (	
bin_numbrer (	
	card_type (
customer_name (	
customer_tel (	
customer_email (	
customer_addr	 (	
customer_postcode
 (	
inserted (
updated (
pg_provider (	
pg_id (	"h
CardRegisterResponse
code (
message (	1
response (2.card_v2.CardRegisterResultData"&
GetCardInfoRequest
card_uid (	"�
GetCardInfoData
card_uid (	
card_provider (	
	card_code (	

bin_number (	
	card_type (
customer_name (	
customer_tel (	
customer_email (	
customer_addr	 (	
customer_postcode
 (	
inserted (
updated ("`
GetCardInfoResponse
code (
message (	*
response (2.card_v2.GetCardInfoData")
DeleteCardInfoRequest
card_uid (	"7
DeleteCardInfoResponse
code (
message (	"�
PayByRegisteredCardRequest
card_uid (	
merchant_uid (	
amount (	
duty_free_amount (	

order_name (	
card_installment (!
interest_free_by_merchant (
custom_data (	
callback_url	 (	

buyer_addr
 (	
buyer_email (	

buyer_name (	
buyer_postcode (	
	buyer_tel (	"�
PayByRegisteredCardData
amount (	
	apply_num (	

buyer_addr (	
buyer_email (	

buyer_name (	
buyer_postcode (	
	buyer_tel (	
cancel_amount (	.
cancel_history	 (2.card_v2.CancelHistory
cancel_reason
 (	
cancel_receipt_urls (	
cancelled_at (
	card_code (	
	card_name (	

bin_number (	
card_installment (
	card_type (
cash_receipt_issued (
channel (	
currency (	
custom_data (	
card_uid (	
customer_uid_usage (	
fail_reason (	
	failed_at (
imp_uid (	
merchant_uid (	

order_name (	
paid_at (

pay_method (	
pg_id (	
pg_provider  (	
pg_tid! (	
receipt_url" (	

started_at# (
status$ (	

user_agent% (	"p
PayByRegisteredCardResponse
code (
message (	2
response (2 .card_v2.PayByRegisteredCardData"j
CancelHistory
pg_tid (	
amount (	
cancelled_at (
reason (	
receipt_url (	"{
CancelCardPaymentRequest
imp_uid (	
amount (	
merchant_uid (	
duty_free_amount (	
reason (	"o
CancelCardPaymenttResponse
code (
message (	2
response (2 .card_v2.PayByRegisteredCardData2�
PaymentCardServicew
CardRegisterV2RPC.card_v2.CardRegisterRequest.card_v2.CardRegisterResponse"%���"/api/payments/v2/card/info:*|
GetCardInfoV2RPC.card_v2.GetCardInfoRequest.card_v2.GetCardInfoResponse"-���\'%/api/payments/v2/card/info/{card_uid}�
DeleteCardInfoV2RPC.card_v2.DeleteCardInfoRequest.card_v2.DeleteCardInfoResponse"-���\'*%/api/payments/v2/card/info/{card_uid}�
PayByRegiseteredCardV2RPC#.card_v2.PayByRegisteredCardRequest$.card_v2.PayByRegisteredCardResponse",���&"$/api/payments/v2/card/pay/registered�
CancelRegiseteredCardV2RPC!.card_v2.CancelCardPaymentRequest#.card_v2.CancelCardPaymenttResponse"9���3"1/api/payments/v2/card/cancel/registered/{imp_uid}BMZ8github.com/iamport/interface/gen_src/go/v2/payments/card�V2.Payments.Cardbproto3'
        , true);

        static::$is_initialized = true;
    }
}

