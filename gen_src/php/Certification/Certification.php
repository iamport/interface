<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/certification/certification.proto

namespace Certification;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>certification.Certification</code>
 */
class Certification extends \Google\Protobuf\Internal\Message
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
     * Generated from protobuf field <code>string pg_tid = 3;</code>
     */
    protected $pg_tid = '';
    /**
     * Generated from protobuf field <code>string pg_provider = 4;</code>
     */
    protected $pg_provider = '';
    /**
     * Generated from protobuf field <code>string name = 5;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string gender = 6;</code>
     */
    protected $gender = '';
    /**
     * Generated from protobuf field <code>int32 birth = 7;</code>
     */
    protected $birth = 0;
    /**
     * Generated from protobuf field <code>string birthday = 8;</code>
     */
    protected $birthday = '';
    /**
     * Generated from protobuf field <code>bool foreigner = 9;</code>
     */
    protected $foreigner = false;
    /**
     * Generated from protobuf field <code>string phone = 10;</code>
     */
    protected $phone = '';
    /**
     * Generated from protobuf field <code>string carrier = 11;</code>
     */
    protected $carrier = '';
    /**
     * Generated from protobuf field <code>bool certified = 12;</code>
     */
    protected $certified = false;
    /**
     * Generated from protobuf field <code>int32 certified_at = 13;</code>
     */
    protected $certified_at = 0;
    /**
     * Generated from protobuf field <code>string unique_key = 14;</code>
     */
    protected $unique_key = '';
    /**
     * Generated from protobuf field <code>string unique_in_site = 15;</code>
     */
    protected $unique_in_site = '';
    /**
     * Generated from protobuf field <code>string origin = 16;</code>
     */
    protected $origin = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $imp_uid
     *     @type string $merchant_uid
     *     @type string $pg_tid
     *     @type string $pg_provider
     *     @type string $name
     *     @type string $gender
     *     @type int $birth
     *     @type string $birthday
     *     @type bool $foreigner
     *     @type string $phone
     *     @type string $carrier
     *     @type bool $certified
     *     @type int $certified_at
     *     @type string $unique_key
     *     @type string $unique_in_site
     *     @type string $origin
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\V1\Certification\Certification::initOnce();
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
     * Generated from protobuf field <code>string pg_tid = 3;</code>
     * @return string
     */
    public function getPgTid()
    {
        return $this->pg_tid;
    }

    /**
     * Generated from protobuf field <code>string pg_tid = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setPgTid($var)
    {
        GPBUtil::checkString($var, True);
        $this->pg_tid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string pg_provider = 4;</code>
     * @return string
     */
    public function getPgProvider()
    {
        return $this->pg_provider;
    }

    /**
     * Generated from protobuf field <code>string pg_provider = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setPgProvider($var)
    {
        GPBUtil::checkString($var, True);
        $this->pg_provider = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string name = 5;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Generated from protobuf field <code>string name = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string gender = 6;</code>
     * @return string
     */
    public function getGender()
    {
        return $this->gender;
    }

    /**
     * Generated from protobuf field <code>string gender = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setGender($var)
    {
        GPBUtil::checkString($var, True);
        $this->gender = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 birth = 7;</code>
     * @return int
     */
    public function getBirth()
    {
        return $this->birth;
    }

    /**
     * Generated from protobuf field <code>int32 birth = 7;</code>
     * @param int $var
     * @return $this
     */
    public function setBirth($var)
    {
        GPBUtil::checkInt32($var);
        $this->birth = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string birthday = 8;</code>
     * @return string
     */
    public function getBirthday()
    {
        return $this->birthday;
    }

    /**
     * Generated from protobuf field <code>string birthday = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setBirthday($var)
    {
        GPBUtil::checkString($var, True);
        $this->birthday = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool foreigner = 9;</code>
     * @return bool
     */
    public function getForeigner()
    {
        return $this->foreigner;
    }

    /**
     * Generated from protobuf field <code>bool foreigner = 9;</code>
     * @param bool $var
     * @return $this
     */
    public function setForeigner($var)
    {
        GPBUtil::checkBool($var);
        $this->foreigner = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string phone = 10;</code>
     * @return string
     */
    public function getPhone()
    {
        return $this->phone;
    }

    /**
     * Generated from protobuf field <code>string phone = 10;</code>
     * @param string $var
     * @return $this
     */
    public function setPhone($var)
    {
        GPBUtil::checkString($var, True);
        $this->phone = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string carrier = 11;</code>
     * @return string
     */
    public function getCarrier()
    {
        return $this->carrier;
    }

    /**
     * Generated from protobuf field <code>string carrier = 11;</code>
     * @param string $var
     * @return $this
     */
    public function setCarrier($var)
    {
        GPBUtil::checkString($var, True);
        $this->carrier = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool certified = 12;</code>
     * @return bool
     */
    public function getCertified()
    {
        return $this->certified;
    }

    /**
     * Generated from protobuf field <code>bool certified = 12;</code>
     * @param bool $var
     * @return $this
     */
    public function setCertified($var)
    {
        GPBUtil::checkBool($var);
        $this->certified = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 certified_at = 13;</code>
     * @return int
     */
    public function getCertifiedAt()
    {
        return $this->certified_at;
    }

    /**
     * Generated from protobuf field <code>int32 certified_at = 13;</code>
     * @param int $var
     * @return $this
     */
    public function setCertifiedAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->certified_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string unique_key = 14;</code>
     * @return string
     */
    public function getUniqueKey()
    {
        return $this->unique_key;
    }

    /**
     * Generated from protobuf field <code>string unique_key = 14;</code>
     * @param string $var
     * @return $this
     */
    public function setUniqueKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->unique_key = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string unique_in_site = 15;</code>
     * @return string
     */
    public function getUniqueInSite()
    {
        return $this->unique_in_site;
    }

    /**
     * Generated from protobuf field <code>string unique_in_site = 15;</code>
     * @param string $var
     * @return $this
     */
    public function setUniqueInSite($var)
    {
        GPBUtil::checkString($var, True);
        $this->unique_in_site = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string origin = 16;</code>
     * @return string
     */
    public function getOrigin()
    {
        return $this->origin;
    }

    /**
     * Generated from protobuf field <code>string origin = 16;</code>
     * @param string $var
     * @return $this
     */
    public function setOrigin($var)
    {
        GPBUtil::checkString($var, True);
        $this->origin = $var;

        return $this;
    }

}

