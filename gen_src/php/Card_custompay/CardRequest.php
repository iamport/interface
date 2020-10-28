<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: custompay/v1/card/card.proto

namespace Card_custompay;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>card_custompay.CardRequest</code>
 */
class CardRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 card_idx = 1;</code>
     */
    protected $card_idx = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $card_idx
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Custompay\V1\Card\Card::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 card_idx = 1;</code>
     * @return int
     */
    public function getCardIdx()
    {
        return $this->card_idx;
    }

    /**
     * Generated from protobuf field <code>int32 card_idx = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setCardIdx($var)
    {
        GPBUtil::checkInt32($var);
        $this->card_idx = $var;

        return $this;
    }

}

