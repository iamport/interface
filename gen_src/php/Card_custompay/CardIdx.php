<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: custompay/v1/card/card.proto

namespace Card_custompay;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>card_custompay.CardIdx</code>
 */
class CardIdx extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 idx = 1;</code>
     */
    protected $idx = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $idx
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Custompay\V1\Card\Card::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 idx = 1;</code>
     * @return int
     */
    public function getIdx()
    {
        return $this->idx;
    }

    /**
     * Generated from protobuf field <code>int32 idx = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setIdx($var)
    {
        GPBUtil::checkInt32($var);
        $this->idx = $var;

        return $this;
    }

}
