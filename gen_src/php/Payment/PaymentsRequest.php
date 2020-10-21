<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/payment/payment.proto

namespace Payment;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>payment.PaymentsRequest</code>
 */
class PaymentsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated string imp_uid = 1;</code>
     */
    private $imp_uid;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $imp_uid
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\V1\Payment\Payment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated string imp_uid = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getImpUid()
    {
        return $this->imp_uid;
    }

    /**
     * Generated from protobuf field <code>repeated string imp_uid = 1;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setImpUid($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->imp_uid = $arr;

        return $this;
    }

}

