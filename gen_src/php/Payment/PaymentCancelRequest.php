<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: payment/payment.proto

namespace Payment;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>payment.PaymentCancelRequest</code>
 */
class PaymentCancelRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string imp_uid = 1;</code>
     */
    protected $imp_uid = '';
    /**
     * Generated from protobuf field <code>string merchant_uid = 2;</code>
     */
    protected $merchant_uid = '';
    /**
     * Generated from protobuf field <code>double amount = 3;</code>
     */
    protected $amount = 0.0;
    /**
     * Generated from protobuf field <code>double tax_free = 4;</code>
     */
    protected $tax_free = 0.0;
    /**
     * Generated from protobuf field <code>double checksum = 5;</code>
     */
    protected $checksum = 0.0;
    /**
     * Generated from protobuf field <code>string reason = 6;</code>
     */
    protected $reason = '';
    /**
     * Generated from protobuf field <code>string refund_holder = 7;</code>
     */
    protected $refund_holder = '';
    /**
     * Generated from protobuf field <code>string refund_bank = 8;</code>
     */
    protected $refund_bank = '';
    /**
     * Generated from protobuf field <code>string refund_account = 9;</code>
     */
    protected $refund_account = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $imp_uid
     *     @type string $merchant_uid
     *     @type float $amount
     *     @type float $tax_free
     *     @type float $checksum
     *     @type string $reason
     *     @type string $refund_holder
     *     @type string $refund_bank
     *     @type string $refund_account
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Payment\Payment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string imp_uid = 1;</code>
     * @return string
     */
    public function getImpUid()
    {
        return $this->imp_uid;
    }

    /**
     * Generated from protobuf field <code>string imp_uid = 1;</code>
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
     * Generated from protobuf field <code>string merchant_uid = 2;</code>
     * @return string
     */
    public function getMerchantUid()
    {
        return $this->merchant_uid;
    }

    /**
     * Generated from protobuf field <code>string merchant_uid = 2;</code>
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
     * Generated from protobuf field <code>double amount = 3;</code>
     * @return float
     */
    public function getAmount()
    {
        return $this->amount;
    }

    /**
     * Generated from protobuf field <code>double amount = 3;</code>
     * @param float $var
     * @return $this
     */
    public function setAmount($var)
    {
        GPBUtil::checkDouble($var);
        $this->amount = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>double tax_free = 4;</code>
     * @return float
     */
    public function getTaxFree()
    {
        return $this->tax_free;
    }

    /**
     * Generated from protobuf field <code>double tax_free = 4;</code>
     * @param float $var
     * @return $this
     */
    public function setTaxFree($var)
    {
        GPBUtil::checkDouble($var);
        $this->tax_free = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>double checksum = 5;</code>
     * @return float
     */
    public function getChecksum()
    {
        return $this->checksum;
    }

    /**
     * Generated from protobuf field <code>double checksum = 5;</code>
     * @param float $var
     * @return $this
     */
    public function setChecksum($var)
    {
        GPBUtil::checkDouble($var);
        $this->checksum = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string reason = 6;</code>
     * @return string
     */
    public function getReason()
    {
        return $this->reason;
    }

    /**
     * Generated from protobuf field <code>string reason = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setReason($var)
    {
        GPBUtil::checkString($var, True);
        $this->reason = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string refund_holder = 7;</code>
     * @return string
     */
    public function getRefundHolder()
    {
        return $this->refund_holder;
    }

    /**
     * Generated from protobuf field <code>string refund_holder = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setRefundHolder($var)
    {
        GPBUtil::checkString($var, True);
        $this->refund_holder = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string refund_bank = 8;</code>
     * @return string
     */
    public function getRefundBank()
    {
        return $this->refund_bank;
    }

    /**
     * Generated from protobuf field <code>string refund_bank = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setRefundBank($var)
    {
        GPBUtil::checkString($var, True);
        $this->refund_bank = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string refund_account = 9;</code>
     * @return string
     */
    public function getRefundAccount()
    {
        return $this->refund_account;
    }

    /**
     * Generated from protobuf field <code>string refund_account = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setRefundAccount($var)
    {
        GPBUtil::checkString($var, True);
        $this->refund_account = $var;

        return $this;
    }

}
