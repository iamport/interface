<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v2/payments/card/card.proto

namespace Card_v2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>card_v2.PayByRegisteredCardData</code>
 */
class PayByRegisteredCardData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string amount = 1;</code>
     */
    protected $amount = '';
    /**
     * Generated from protobuf field <code>string apply_num = 2;</code>
     */
    protected $apply_num = '';
    /**
     * Generated from protobuf field <code>string buyer_addr = 3;</code>
     */
    protected $buyer_addr = '';
    /**
     * Generated from protobuf field <code>string buyer_email = 4;</code>
     */
    protected $buyer_email = '';
    /**
     * Generated from protobuf field <code>string buyer_name = 5;</code>
     */
    protected $buyer_name = '';
    /**
     * Generated from protobuf field <code>string buyer_postcode = 6;</code>
     */
    protected $buyer_postcode = '';
    /**
     * Generated from protobuf field <code>string buyer_tel = 7;</code>
     */
    protected $buyer_tel = '';
    /**
     * Generated from protobuf field <code>string cancel_amount = 8;</code>
     */
    protected $cancel_amount = '';
    /**
     * Generated from protobuf field <code>repeated .card_v2.CancelHistory cancel_history = 9;</code>
     */
    private $cancel_history;
    /**
     * Generated from protobuf field <code>string cancel_reason = 10;</code>
     */
    protected $cancel_reason = '';
    /**
     * Generated from protobuf field <code>repeated string cancel_receipt_urls = 11;</code>
     */
    private $cancel_receipt_urls;
    /**
     * Generated from protobuf field <code>int32 cancelled_at = 12;</code>
     */
    protected $cancelled_at = 0;
    /**
     * Generated from protobuf field <code>string card_code = 13;</code>
     */
    protected $card_code = '';
    /**
     * Generated from protobuf field <code>string card_name = 14;</code>
     */
    protected $card_name = '';
    /**
     * Generated from protobuf field <code>string bin_number = 15;</code>
     */
    protected $bin_number = '';
    /**
     * Generated from protobuf field <code>int32 card_installment = 16;</code>
     */
    protected $card_installment = 0;
    /**
     * Generated from protobuf field <code>int32 card_type = 17;</code>
     */
    protected $card_type = 0;
    /**
     * Generated from protobuf field <code>bool cash_receipt_issued = 18;</code>
     */
    protected $cash_receipt_issued = false;
    /**
     * Generated from protobuf field <code>string channel = 19;</code>
     */
    protected $channel = '';
    /**
     * Generated from protobuf field <code>string currency = 20;</code>
     */
    protected $currency = '';
    /**
     * Generated from protobuf field <code>string custom_data = 21;</code>
     */
    protected $custom_data = '';
    /**
     * a.k.a.) customer_uid
     *
     * Generated from protobuf field <code>string card_uid = 22;</code>
     */
    protected $card_uid = '';
    /**
     * Generated from protobuf field <code>string customer_uid_usage = 23;</code>
     */
    protected $customer_uid_usage = '';
    /**
     * Generated from protobuf field <code>string fail_reason = 24;</code>
     */
    protected $fail_reason = '';
    /**
     * Generated from protobuf field <code>int32 failed_at = 25;</code>
     */
    protected $failed_at = 0;
    /**
     * Generated from protobuf field <code>string imp_uid = 26;</code>
     */
    protected $imp_uid = '';
    /**
     * Generated from protobuf field <code>string merchant_uid = 27;</code>
     */
    protected $merchant_uid = '';
    /**
     * Generated from protobuf field <code>string order_name = 28;</code>
     */
    protected $order_name = '';
    /**
     * Generated from protobuf field <code>int32 paid_at = 29;</code>
     */
    protected $paid_at = 0;
    /**
     * Generated from protobuf field <code>string pay_method = 30;</code>
     */
    protected $pay_method = '';
    /**
     * Generated from protobuf field <code>string pg_id = 31;</code>
     */
    protected $pg_id = '';
    /**
     * Generated from protobuf field <code>string pg_provider = 32;</code>
     */
    protected $pg_provider = '';
    /**
     * Generated from protobuf field <code>string pg_tid = 33;</code>
     */
    protected $pg_tid = '';
    /**
     * Generated from protobuf field <code>string receipt_url = 34;</code>
     */
    protected $receipt_url = '';
    /**
     * Generated from protobuf field <code>int32 started_at = 35;</code>
     */
    protected $started_at = 0;
    /**
     * Generated from protobuf field <code>string status = 36;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>string user_agent = 37;</code>
     */
    protected $user_agent = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $amount
     *     @type string $apply_num
     *     @type string $buyer_addr
     *     @type string $buyer_email
     *     @type string $buyer_name
     *     @type string $buyer_postcode
     *     @type string $buyer_tel
     *     @type string $cancel_amount
     *     @type \Card_v2\CancelHistory[]|\Google\Protobuf\Internal\RepeatedField $cancel_history
     *     @type string $cancel_reason
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $cancel_receipt_urls
     *     @type int $cancelled_at
     *     @type string $card_code
     *     @type string $card_name
     *     @type string $bin_number
     *     @type int $card_installment
     *     @type int $card_type
     *     @type bool $cash_receipt_issued
     *     @type string $channel
     *     @type string $currency
     *     @type string $custom_data
     *     @type string $card_uid
     *           a.k.a.) customer_uid
     *     @type string $customer_uid_usage
     *     @type string $fail_reason
     *     @type int $failed_at
     *     @type string $imp_uid
     *     @type string $merchant_uid
     *     @type string $order_name
     *     @type int $paid_at
     *     @type string $pay_method
     *     @type string $pg_id
     *     @type string $pg_provider
     *     @type string $pg_tid
     *     @type string $receipt_url
     *     @type int $started_at
     *     @type string $status
     *     @type string $user_agent
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\V2\Payments\Card\Card::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string amount = 1;</code>
     * @return string
     */
    public function getAmount()
    {
        return $this->amount;
    }

    /**
     * Generated from protobuf field <code>string amount = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setAmount($var)
    {
        GPBUtil::checkString($var, True);
        $this->amount = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string apply_num = 2;</code>
     * @return string
     */
    public function getApplyNum()
    {
        return $this->apply_num;
    }

    /**
     * Generated from protobuf field <code>string apply_num = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setApplyNum($var)
    {
        GPBUtil::checkString($var, True);
        $this->apply_num = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string buyer_addr = 3;</code>
     * @return string
     */
    public function getBuyerAddr()
    {
        return $this->buyer_addr;
    }

    /**
     * Generated from protobuf field <code>string buyer_addr = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setBuyerAddr($var)
    {
        GPBUtil::checkString($var, True);
        $this->buyer_addr = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string buyer_email = 4;</code>
     * @return string
     */
    public function getBuyerEmail()
    {
        return $this->buyer_email;
    }

    /**
     * Generated from protobuf field <code>string buyer_email = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setBuyerEmail($var)
    {
        GPBUtil::checkString($var, True);
        $this->buyer_email = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string buyer_name = 5;</code>
     * @return string
     */
    public function getBuyerName()
    {
        return $this->buyer_name;
    }

    /**
     * Generated from protobuf field <code>string buyer_name = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setBuyerName($var)
    {
        GPBUtil::checkString($var, True);
        $this->buyer_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string buyer_postcode = 6;</code>
     * @return string
     */
    public function getBuyerPostcode()
    {
        return $this->buyer_postcode;
    }

    /**
     * Generated from protobuf field <code>string buyer_postcode = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setBuyerPostcode($var)
    {
        GPBUtil::checkString($var, True);
        $this->buyer_postcode = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string buyer_tel = 7;</code>
     * @return string
     */
    public function getBuyerTel()
    {
        return $this->buyer_tel;
    }

    /**
     * Generated from protobuf field <code>string buyer_tel = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setBuyerTel($var)
    {
        GPBUtil::checkString($var, True);
        $this->buyer_tel = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string cancel_amount = 8;</code>
     * @return string
     */
    public function getCancelAmount()
    {
        return $this->cancel_amount;
    }

    /**
     * Generated from protobuf field <code>string cancel_amount = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setCancelAmount($var)
    {
        GPBUtil::checkString($var, True);
        $this->cancel_amount = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .card_v2.CancelHistory cancel_history = 9;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCancelHistory()
    {
        return $this->cancel_history;
    }

    /**
     * Generated from protobuf field <code>repeated .card_v2.CancelHistory cancel_history = 9;</code>
     * @param \Card_v2\CancelHistory[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCancelHistory($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Card_v2\CancelHistory::class);
        $this->cancel_history = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string cancel_reason = 10;</code>
     * @return string
     */
    public function getCancelReason()
    {
        return $this->cancel_reason;
    }

    /**
     * Generated from protobuf field <code>string cancel_reason = 10;</code>
     * @param string $var
     * @return $this
     */
    public function setCancelReason($var)
    {
        GPBUtil::checkString($var, True);
        $this->cancel_reason = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string cancel_receipt_urls = 11;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCancelReceiptUrls()
    {
        return $this->cancel_receipt_urls;
    }

    /**
     * Generated from protobuf field <code>repeated string cancel_receipt_urls = 11;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCancelReceiptUrls($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->cancel_receipt_urls = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 cancelled_at = 12;</code>
     * @return int
     */
    public function getCancelledAt()
    {
        return $this->cancelled_at;
    }

    /**
     * Generated from protobuf field <code>int32 cancelled_at = 12;</code>
     * @param int $var
     * @return $this
     */
    public function setCancelledAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->cancelled_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string card_code = 13;</code>
     * @return string
     */
    public function getCardCode()
    {
        return $this->card_code;
    }

    /**
     * Generated from protobuf field <code>string card_code = 13;</code>
     * @param string $var
     * @return $this
     */
    public function setCardCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->card_code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string card_name = 14;</code>
     * @return string
     */
    public function getCardName()
    {
        return $this->card_name;
    }

    /**
     * Generated from protobuf field <code>string card_name = 14;</code>
     * @param string $var
     * @return $this
     */
    public function setCardName($var)
    {
        GPBUtil::checkString($var, True);
        $this->card_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string bin_number = 15;</code>
     * @return string
     */
    public function getBinNumber()
    {
        return $this->bin_number;
    }

    /**
     * Generated from protobuf field <code>string bin_number = 15;</code>
     * @param string $var
     * @return $this
     */
    public function setBinNumber($var)
    {
        GPBUtil::checkString($var, True);
        $this->bin_number = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 card_installment = 16;</code>
     * @return int
     */
    public function getCardInstallment()
    {
        return $this->card_installment;
    }

    /**
     * Generated from protobuf field <code>int32 card_installment = 16;</code>
     * @param int $var
     * @return $this
     */
    public function setCardInstallment($var)
    {
        GPBUtil::checkInt32($var);
        $this->card_installment = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 card_type = 17;</code>
     * @return int
     */
    public function getCardType()
    {
        return $this->card_type;
    }

    /**
     * Generated from protobuf field <code>int32 card_type = 17;</code>
     * @param int $var
     * @return $this
     */
    public function setCardType($var)
    {
        GPBUtil::checkInt32($var);
        $this->card_type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool cash_receipt_issued = 18;</code>
     * @return bool
     */
    public function getCashReceiptIssued()
    {
        return $this->cash_receipt_issued;
    }

    /**
     * Generated from protobuf field <code>bool cash_receipt_issued = 18;</code>
     * @param bool $var
     * @return $this
     */
    public function setCashReceiptIssued($var)
    {
        GPBUtil::checkBool($var);
        $this->cash_receipt_issued = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string channel = 19;</code>
     * @return string
     */
    public function getChannel()
    {
        return $this->channel;
    }

    /**
     * Generated from protobuf field <code>string channel = 19;</code>
     * @param string $var
     * @return $this
     */
    public function setChannel($var)
    {
        GPBUtil::checkString($var, True);
        $this->channel = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string currency = 20;</code>
     * @return string
     */
    public function getCurrency()
    {
        return $this->currency;
    }

    /**
     * Generated from protobuf field <code>string currency = 20;</code>
     * @param string $var
     * @return $this
     */
    public function setCurrency($var)
    {
        GPBUtil::checkString($var, True);
        $this->currency = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string custom_data = 21;</code>
     * @return string
     */
    public function getCustomData()
    {
        return $this->custom_data;
    }

    /**
     * Generated from protobuf field <code>string custom_data = 21;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomData($var)
    {
        GPBUtil::checkString($var, True);
        $this->custom_data = $var;

        return $this;
    }

    /**
     * a.k.a.) customer_uid
     *
     * Generated from protobuf field <code>string card_uid = 22;</code>
     * @return string
     */
    public function getCardUid()
    {
        return $this->card_uid;
    }

    /**
     * a.k.a.) customer_uid
     *
     * Generated from protobuf field <code>string card_uid = 22;</code>
     * @param string $var
     * @return $this
     */
    public function setCardUid($var)
    {
        GPBUtil::checkString($var, True);
        $this->card_uid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string customer_uid_usage = 23;</code>
     * @return string
     */
    public function getCustomerUidUsage()
    {
        return $this->customer_uid_usage;
    }

    /**
     * Generated from protobuf field <code>string customer_uid_usage = 23;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomerUidUsage($var)
    {
        GPBUtil::checkString($var, True);
        $this->customer_uid_usage = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string fail_reason = 24;</code>
     * @return string
     */
    public function getFailReason()
    {
        return $this->fail_reason;
    }

    /**
     * Generated from protobuf field <code>string fail_reason = 24;</code>
     * @param string $var
     * @return $this
     */
    public function setFailReason($var)
    {
        GPBUtil::checkString($var, True);
        $this->fail_reason = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 failed_at = 25;</code>
     * @return int
     */
    public function getFailedAt()
    {
        return $this->failed_at;
    }

    /**
     * Generated from protobuf field <code>int32 failed_at = 25;</code>
     * @param int $var
     * @return $this
     */
    public function setFailedAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->failed_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string imp_uid = 26;</code>
     * @return string
     */
    public function getImpUid()
    {
        return $this->imp_uid;
    }

    /**
     * Generated from protobuf field <code>string imp_uid = 26;</code>
     * @param string $var
     * @return $this
     */
    public function setImpUid($var)
    {
        GPBUtil::checkString($var, True);
        $this->imp_uid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string merchant_uid = 27;</code>
     * @return string
     */
    public function getMerchantUid()
    {
        return $this->merchant_uid;
    }

    /**
     * Generated from protobuf field <code>string merchant_uid = 27;</code>
     * @param string $var
     * @return $this
     */
    public function setMerchantUid($var)
    {
        GPBUtil::checkString($var, True);
        $this->merchant_uid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string order_name = 28;</code>
     * @return string
     */
    public function getOrderName()
    {
        return $this->order_name;
    }

    /**
     * Generated from protobuf field <code>string order_name = 28;</code>
     * @param string $var
     * @return $this
     */
    public function setOrderName($var)
    {
        GPBUtil::checkString($var, True);
        $this->order_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 paid_at = 29;</code>
     * @return int
     */
    public function getPaidAt()
    {
        return $this->paid_at;
    }

    /**
     * Generated from protobuf field <code>int32 paid_at = 29;</code>
     * @param int $var
     * @return $this
     */
    public function setPaidAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->paid_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string pay_method = 30;</code>
     * @return string
     */
    public function getPayMethod()
    {
        return $this->pay_method;
    }

    /**
     * Generated from protobuf field <code>string pay_method = 30;</code>
     * @param string $var
     * @return $this
     */
    public function setPayMethod($var)
    {
        GPBUtil::checkString($var, True);
        $this->pay_method = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string pg_id = 31;</code>
     * @return string
     */
    public function getPgId()
    {
        return $this->pg_id;
    }

    /**
     * Generated from protobuf field <code>string pg_id = 31;</code>
     * @param string $var
     * @return $this
     */
    public function setPgId($var)
    {
        GPBUtil::checkString($var, True);
        $this->pg_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string pg_provider = 32;</code>
     * @return string
     */
    public function getPgProvider()
    {
        return $this->pg_provider;
    }

    /**
     * Generated from protobuf field <code>string pg_provider = 32;</code>
     * @param string $var
     * @return $this
     */
    public function setPgProvider($var)
    {
        GPBUtil::checkString($var, True);
        $this->pg_provider = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string pg_tid = 33;</code>
     * @return string
     */
    public function getPgTid()
    {
        return $this->pg_tid;
    }

    /**
     * Generated from protobuf field <code>string pg_tid = 33;</code>
     * @param string $var
     * @return $this
     */
    public function setPgTid($var)
    {
        GPBUtil::checkString($var, True);
        $this->pg_tid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string receipt_url = 34;</code>
     * @return string
     */
    public function getReceiptUrl()
    {
        return $this->receipt_url;
    }

    /**
     * Generated from protobuf field <code>string receipt_url = 34;</code>
     * @param string $var
     * @return $this
     */
    public function setReceiptUrl($var)
    {
        GPBUtil::checkString($var, True);
        $this->receipt_url = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 started_at = 35;</code>
     * @return int
     */
    public function getStartedAt()
    {
        return $this->started_at;
    }

    /**
     * Generated from protobuf field <code>int32 started_at = 35;</code>
     * @param int $var
     * @return $this
     */
    public function setStartedAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->started_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 36;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 36;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string user_agent = 37;</code>
     * @return string
     */
    public function getUserAgent()
    {
        return $this->user_agent;
    }

    /**
     * Generated from protobuf field <code>string user_agent = 37;</code>
     * @param string $var
     * @return $this
     */
    public function setUserAgent($var)
    {
        GPBUtil::checkString($var, True);
        $this->user_agent = $var;

        return $this;
    }

}

