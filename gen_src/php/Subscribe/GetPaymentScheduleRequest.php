<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/subscribe/subscribe.proto

namespace Subscribe;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>subscribe.GetPaymentScheduleRequest</code>
 */
class GetPaymentScheduleRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string merchant_uid = 1;</code>
     */
    protected $merchant_uid = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $merchant_uid
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\V1\Subscribe\Subscribe::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string merchant_uid = 1;</code>
     * @return string
     */
    public function getMerchantUid()
    {
        return $this->merchant_uid;
    }

    /**
     * Generated from protobuf field <code>string merchant_uid = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setMerchantUid($var)
    {
        GPBUtil::checkString($var, True);
        $this->merchant_uid = $var;

        return $this;
    }

}

