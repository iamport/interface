<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/certification/certification.proto

namespace GPBMetadata\V1\Certification;

class Certification
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Api\Annotations::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
$v1/certification/certification.protocertification"\'
CertificationRequest
imp_uid (	"f
CertificationResponse
code (
message (	.
response (2.certification.Certification"�
Certification
imp_uid (	
merchant_uid (	
pg_tid (	
pg_provider (	
name (	
gender (	
birth (
birthday (	
	foreigner	 (
phone
 (	
carrier (	
	certified (
certified_at (

unique_key (	
unique_in_site (	
origin (	2�
CertificationService�
CertificationGetRPC#.certification.CertificationRequest$.certification.CertificationResponse"1���+)/api/payments/v1/certifications/{imp_uid}�
CertificationDeleteRPC#.certification.CertificationRequest$.certification.CertificationResponse"1���+*)/api/payments/v1/certifications/{imp_uid}BMZ8github.com/iamport/interface/gen_src/go/v1/certification�V1.Certificationbproto3'
        , true);

        static::$is_initialized = true;
    }
}

