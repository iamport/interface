<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v2/txs/txs.proto

namespace Txs_v2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>txs_v2.GetTxsRecordsRequest</code>
 */
class GetTxsRecordsRequest extends \Google\Protobuf\Internal\Message
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
     * Generated from protobuf field <code>string paid_by = 3;</code>
     */
    protected $paid_by = '';
    /**
     * Generated from protobuf field <code>string status = 4;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>int32 page = 5;</code>
     */
    protected $page = 0;
    /**
     * Generated from protobuf field <code>string sorting = 6;</code>
     */
    protected $sorting = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $imp_uid
     *     @type string $merchant_uid
     *     @type string $paid_by
     *     @type string $status
     *     @type int $page
     *     @type string $sorting
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\V2\Txs\Txs::initOnce();
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
     * Generated from protobuf field <code>string paid_by = 3;</code>
     * @return string
     */
    public function getPaidBy()
    {
        return $this->paid_by;
    }

    /**
     * Generated from protobuf field <code>string paid_by = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setPaidBy($var)
    {
        GPBUtil::checkString($var, True);
        $this->paid_by = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 4;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 page = 5;</code>
     * @return int
     */
    public function getPage()
    {
        return $this->page;
    }

    /**
     * Generated from protobuf field <code>int32 page = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setPage($var)
    {
        GPBUtil::checkInt32($var);
        $this->page = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string sorting = 6;</code>
     * @return string
     */
    public function getSorting()
    {
        return $this->sorting;
    }

    /**
     * Generated from protobuf field <code>string sorting = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setSorting($var)
    {
        GPBUtil::checkString($var, True);
        $this->sorting = $var;

        return $this;
    }

}

