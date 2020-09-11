<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: payment/payment.proto

namespace Payment;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>payment.PaymentPage</code>
 */
class PaymentPage extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     */
    protected $total = 0;
    /**
     * Generated from protobuf field <code>int32 previous = 2;</code>
     */
    protected $previous = 0;
    /**
     * Generated from protobuf field <code>int32 next = 3;</code>
     */
    protected $next = 0;
    /**
     * Generated from protobuf field <code>repeated .payment.Payment list = 4;</code>
     */
    private $list;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $total
     *     @type int $previous
     *     @type int $next
     *     @type \Payment\Payment[]|\Google\Protobuf\Internal\RepeatedField $list
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Payment\Payment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     * @return int
     */
    public function getTotal()
    {
        return $this->total;
    }

    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setTotal($var)
    {
        GPBUtil::checkInt32($var);
        $this->total = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 previous = 2;</code>
     * @return int
     */
    public function getPrevious()
    {
        return $this->previous;
    }

    /**
     * Generated from protobuf field <code>int32 previous = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setPrevious($var)
    {
        GPBUtil::checkInt32($var);
        $this->previous = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 next = 3;</code>
     * @return int
     */
    public function getNext()
    {
        return $this->next;
    }

    /**
     * Generated from protobuf field <code>int32 next = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setNext($var)
    {
        GPBUtil::checkInt32($var);
        $this->next = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .payment.Payment list = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getList()
    {
        return $this->list;
    }

    /**
     * Generated from protobuf field <code>repeated .payment.Payment list = 4;</code>
     * @param \Payment\Payment[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setList($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Payment\Payment::class);
        $this->list = $arr;

        return $this;
    }

}

