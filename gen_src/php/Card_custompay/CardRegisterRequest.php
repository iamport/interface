<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: custompay/v1/card/card.proto

namespace Card_custompay;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>card_custompay.CardRegisterRequest</code>
 */
class CardRegisterRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string card_idx = 1;</code>
     */
    protected $card_idx = '';
    /**
     * Generated from protobuf field <code>string alias = 2;</code>
     */
    protected $alias = '';
    /**
     * Generated from protobuf field <code>string customer_uid = 3;</code>
     */
    protected $customer_uid = '';
    /**
     * Generated from protobuf field <code>string pg = 4;</code>
     */
    protected $pg = '';
    /**
     * Generated from protobuf field <code>string expiry = 5;</code>
     */
    protected $expiry = '';
    /**
     * Generated from protobuf field <code>string birth = 6;</code>
     */
    protected $birth = '';
    /**
     * Generated from protobuf field <code>string pwd_2digit = 7;</code>
     */
    protected $pwd_2digit = '';
    /**
     * Generated from protobuf field <code>string customer_name = 8;</code>
     */
    protected $customer_name = '';
    /**
     * Generated from protobuf field <code>string customer_tel = 9;</code>
     */
    protected $customer_tel = '';
    /**
     * Generated from protobuf field <code>string customer_email = 10;</code>
     */
    protected $customer_email = '';
    /**
     * Generated from protobuf field <code>string customer_addr = 11;</code>
     */
    protected $customer_addr = '';
    /**
     * Generated from protobuf field <code>string customer_postcode = 12;</code>
     */
    protected $customer_postcode = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $card_idx
     *     @type string $alias
     *     @type string $customer_uid
     *     @type string $pg
     *     @type string $expiry
     *     @type string $birth
     *     @type string $pwd_2digit
     *     @type string $customer_name
     *     @type string $customer_tel
     *     @type string $customer_email
     *     @type string $customer_addr
     *     @type string $customer_postcode
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Custompay\V1\Card\Card::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string card_idx = 1;</code>
     * @return string
     */
    public function getCardIdx()
    {
        return $this->card_idx;
    }

    /**
     * Generated from protobuf field <code>string card_idx = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setCardIdx($var)
    {
        GPBUtil::checkString($var, True);
        $this->card_idx = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string alias = 2;</code>
     * @return string
     */
    public function getAlias()
    {
        return $this->alias;
    }

    /**
     * Generated from protobuf field <code>string alias = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setAlias($var)
    {
        GPBUtil::checkString($var, True);
        $this->alias = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string customer_uid = 3;</code>
     * @return string
     */
    public function getCustomerUid()
    {
        return $this->customer_uid;
    }

    /**
     * Generated from protobuf field <code>string customer_uid = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomerUid($var)
    {
        GPBUtil::checkString($var, True);
        $this->customer_uid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string pg = 4;</code>
     * @return string
     */
    public function getPg()
    {
        return $this->pg;
    }

    /**
     * Generated from protobuf field <code>string pg = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setPg($var)
    {
        GPBUtil::checkString($var, True);
        $this->pg = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string expiry = 5;</code>
     * @return string
     */
    public function getExpiry()
    {
        return $this->expiry;
    }

    /**
     * Generated from protobuf field <code>string expiry = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setExpiry($var)
    {
        GPBUtil::checkString($var, True);
        $this->expiry = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string birth = 6;</code>
     * @return string
     */
    public function getBirth()
    {
        return $this->birth;
    }

    /**
     * Generated from protobuf field <code>string birth = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setBirth($var)
    {
        GPBUtil::checkString($var, True);
        $this->birth = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string pwd_2digit = 7;</code>
     * @return string
     */
    public function getPwd2Digit()
    {
        return $this->pwd_2digit;
    }

    /**
     * Generated from protobuf field <code>string pwd_2digit = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setPwd2Digit($var)
    {
        GPBUtil::checkString($var, True);
        $this->pwd_2digit = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string customer_name = 8;</code>
     * @return string
     */
    public function getCustomerName()
    {
        return $this->customer_name;
    }

    /**
     * Generated from protobuf field <code>string customer_name = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomerName($var)
    {
        GPBUtil::checkString($var, True);
        $this->customer_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string customer_tel = 9;</code>
     * @return string
     */
    public function getCustomerTel()
    {
        return $this->customer_tel;
    }

    /**
     * Generated from protobuf field <code>string customer_tel = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomerTel($var)
    {
        GPBUtil::checkString($var, True);
        $this->customer_tel = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string customer_email = 10;</code>
     * @return string
     */
    public function getCustomerEmail()
    {
        return $this->customer_email;
    }

    /**
     * Generated from protobuf field <code>string customer_email = 10;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomerEmail($var)
    {
        GPBUtil::checkString($var, True);
        $this->customer_email = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string customer_addr = 11;</code>
     * @return string
     */
    public function getCustomerAddr()
    {
        return $this->customer_addr;
    }

    /**
     * Generated from protobuf field <code>string customer_addr = 11;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomerAddr($var)
    {
        GPBUtil::checkString($var, True);
        $this->customer_addr = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string customer_postcode = 12;</code>
     * @return string
     */
    public function getCustomerPostcode()
    {
        return $this->customer_postcode;
    }

    /**
     * Generated from protobuf field <code>string customer_postcode = 12;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomerPostcode($var)
    {
        GPBUtil::checkString($var, True);
        $this->customer_postcode = $var;

        return $this;
    }

}

