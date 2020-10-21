<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v2/payments/card/card.proto

namespace Card_v2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>card_v2.CancelCardPaymentRequest</code>
 */
class CancelCardPaymentRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string imp_uid = 1;</code>
     */
    protected $imp_uid = '';
    /**
     * Generated from protobuf field <code>string amount = 2;</code>
     */
    protected $amount = '';
    /**
     * Generated from protobuf field <code>string merchant_uid = 3;</code>
     */
    protected $merchant_uid = '';
    /**
     * Generated from protobuf field <code>string duty_free_amount = 4;</code>
     */
    protected $duty_free_amount = '';
    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     */
    protected $reason = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $imp_uid
     *     @type string $amount
     *     @type string $merchant_uid
     *     @type string $duty_free_amount
     *     @type string $reason
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\V2\Payments\Card\Card::initOnce();
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
     * Generated from protobuf field <code>string amount = 2;</code>
     * @return string
     */
    public function getAmount()
    {
        return $this->amount;
    }

    /**
     * Generated from protobuf field <code>string amount = 2;</code>
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
     * Generated from protobuf field <code>string merchant_uid = 3;</code>
     * @return string
     */
    public function getMerchantUid()
    {
        return $this->merchant_uid;
    }

    /**
     * Generated from protobuf field <code>string merchant_uid = 3;</code>
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
     * Generated from protobuf field <code>string duty_free_amount = 4;</code>
     * @return string
     */
    public function getDutyFreeAmount()
    {
        return $this->duty_free_amount;
    }

    /**
     * Generated from protobuf field <code>string duty_free_amount = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setDutyFreeAmount($var)
    {
        GPBUtil::checkString($var, True);
        $this->duty_free_amount = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     * @return string
     */
    public function getReason()
    {
        return $this->reason;
    }

    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setReason($var)
    {
        GPBUtil::checkString($var, True);
        $this->reason = $var;

        return $this;
    }

}

