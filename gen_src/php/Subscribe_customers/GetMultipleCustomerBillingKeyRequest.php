<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/subscribe_customers/subscribe_customers.proto

namespace Subscribe_customers;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>subscribe_customers.GetMultipleCustomerBillingKeyRequest</code>
 */
class GetMultipleCustomerBillingKeyRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated string customer_uid = 1;</code>
     */
    private $customer_uid;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $customer_uid
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\V1\SubscribeCustomers\SubscribeCustomers::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated string customer_uid = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCustomerUid()
    {
        return $this->customer_uid;
    }

    /**
     * Generated from protobuf field <code>repeated string customer_uid = 1;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCustomerUid($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->customer_uid = $arr;

        return $this;
    }

}

