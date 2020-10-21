<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/authenticate/token.proto

namespace Authenticate;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>authenticate.Token</code>
 */
class Token extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string access_token = 1;</code>
     */
    protected $access_token = '';
    /**
     * Generated from protobuf field <code>int32 now = 2;</code>
     */
    protected $now = 0;
    /**
     * Generated from protobuf field <code>int32 expired_at = 3;</code>
     */
    protected $expired_at = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $access_token
     *     @type int $now
     *     @type int $expired_at
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\V1\Authenticate\Token::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string access_token = 1;</code>
     * @return string
     */
    public function getAccessToken()
    {
        return $this->access_token;
    }

    /**
     * Generated from protobuf field <code>string access_token = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setAccessToken($var)
    {
        GPBUtil::checkString($var, True);
        $this->access_token = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 now = 2;</code>
     * @return int
     */
    public function getNow()
    {
        return $this->now;
    }

    /**
     * Generated from protobuf field <code>int32 now = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setNow($var)
    {
        GPBUtil::checkInt32($var);
        $this->now = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 expired_at = 3;</code>
     * @return int
     */
    public function getExpiredAt()
    {
        return $this->expired_at;
    }

    /**
     * Generated from protobuf field <code>int32 expired_at = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setExpiredAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->expired_at = $var;

        return $this;
    }

}

