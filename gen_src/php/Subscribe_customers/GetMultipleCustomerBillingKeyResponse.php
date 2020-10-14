<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: subscribe_customers/subscribe_customers.proto

namespace Subscribe_customers;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>subscribe_customers.GetMultipleCustomerBillingKeyResponse</code>
 */
class GetMultipleCustomerBillingKeyResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 code = 1;</code>
     */
    protected $code = 0;
    /**
     * Generated from protobuf field <code>string message = 2;</code>
     */
    protected $message = '';
    /**
     * Generated from protobuf field <code>repeated .subscribe_customers.CustomerBillingKey response = 3;</code>
     */
    private $response;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $code
     *     @type string $message
     *     @type \Subscribe_customers\CustomerBillingKey[]|\Google\Protobuf\Internal\RepeatedField $response
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\SubscribeCustomers\SubscribeCustomers::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 code = 1;</code>
     * @return int
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * Generated from protobuf field <code>int32 code = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setCode($var)
    {
        GPBUtil::checkInt32($var);
        $this->code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string message = 2;</code>
     * @return string
     */
    public function getMessage()
    {
        return $this->message;
    }

    /**
     * Generated from protobuf field <code>string message = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMessage($var)
    {
        GPBUtil::checkString($var, True);
        $this->message = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .subscribe_customers.CustomerBillingKey response = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getResponse()
    {
        return $this->response;
    }

    /**
     * Generated from protobuf field <code>repeated .subscribe_customers.CustomerBillingKey response = 3;</code>
     * @param \Subscribe_customers\CustomerBillingKey[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setResponse($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Subscribe_customers\CustomerBillingKey::class);
        $this->response = $arr;

        return $this;
    }

}
