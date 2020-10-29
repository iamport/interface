# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/subscribe/subscribe.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
require 'v1/payment/payment_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("v1/subscribe/subscribe.proto", :syntax => :proto3) do
    add_message "subscribe.OnetimePaymentRequest" do
      optional :merchant_uid, :string, 1
      optional :amount, :int32, 2
      optional :tax_free, :int32, 3
      optional :card_number, :string, 4
      optional :expiry, :string, 5
      optional :birth, :string, 6
      optional :pwd_2digit, :string, 7
      optional :customer_uid, :string, 8
      optional :pg, :string, 9
      optional :name, :string, 10
      optional :buyer_name, :string, 11
      optional :buyer_email, :string, 12
      optional :buyer_tel, :string, 13
      optional :buyer_addr, :string, 14
      optional :buyer_postcode, :string, 15
      optional :card_quota, :int32, 16
      optional :interest_free_by_merchant, :bool, 17
      optional :custom_data, :string, 18
      optional :notice_url, :string, 19
    end
    add_message "subscribe.OnetimePaymentResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "payment.Payment"
    end
    add_message "subscribe.AgainPaymentRequest" do
      optional :customer_uid, :string, 1
      optional :merchant_uid, :string, 2
      optional :amount, :int32, 3
      optional :tax_free, :int32, 4
      optional :name, :string, 5
      optional :buyer_name, :string, 6
      optional :buyer_email, :string, 7
      optional :buyer_tel, :string, 8
      optional :buyer_addr, :string, 9
      optional :buyer_postcode, :string, 10
      optional :card_quota, :int32, 11
      optional :interest_free_by_merchant, :bool, 12
      optional :custom_data, :string, 13
      optional :notice_url, :string, 14
    end
    add_message "subscribe.AgainPaymentResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "payment.Payment"
    end
    add_message "subscribe.PaymentScheduleParam" do
      optional :merchant_uid, :string, 1
      optional :schedule_at, :int32, 2
      optional :amount, :int32, 3
      optional :tax_free, :int32, 4
      optional :name, :string, 5
      optional :buyer_name, :string, 6
      optional :buyer_email, :string, 7
      optional :buyer_tel, :string, 8
      optional :buyer_addr, :string, 9
      optional :buyer_postcode, :string, 10
    end
    add_message "subscribe.UnitSchedulePaymentResponse" do
      optional :customer_uid, :string, 1
      optional :merchant_uid, :string, 2
      optional :imp_uid, :string, 3
      optional :schedule_at, :int32, 4
      optional :executed_at, :int32, 5
      optional :revoked_at, :int32, 6
      optional :amount, :int32, 7
      optional :name, :string, 8
      optional :buyer_name, :string, 9
      optional :buyer_email, :string, 10
      optional :buyer_tel, :string, 11
      optional :buyer_addr, :string, 12
      optional :buyer_postcode, :string, 13
      optional :custom_data, :string, 14
      optional :schedule_status, :string, 15
      optional :payment_status, :string, 16
      optional :fail_reason, :string, 17
    end
    add_message "subscribe.SchedulePayemntRequest" do
      optional :customer_uid, :string, 1
      optional :checking_amount, :int32, 2
      optional :card_number, :string, 3
      optional :expiry, :string, 4
      optional :birth, :string, 5
      optional :pwd_2digit, :string, 6
      optional :pg, :string, 7
      repeated :schedules, :message, 8, "subscribe.PaymentScheduleParam"
    end
    add_message "subscribe.SchedulePaymentResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      repeated :response, :message, 3, "subscribe.UnitSchedulePaymentResponse"
    end
    add_message "subscribe.UnschedulePaymentRequest" do
      optional :customer_uid, :string, 1
      repeated :merchant_uid, :string, 2
    end
    add_message "subscribe.UnschedulePaymentResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      repeated :response, :message, 3, "subscribe.UnitSchedulePaymentResponse"
    end
    add_message "subscribe.GetPaymentScheduleRequest" do
      optional :merchant_uid, :string, 1
    end
    add_message "subscribe.GetPaymentScheduleResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "subscribe.UnitSchedulePaymentResponse"
    end
    add_message "subscribe.GetPaymentScheduleByCustomerRequest" do
      optional :customer_uid, :string, 1
      optional :page, :int32, 2
      optional :from, :int32, 3
      optional :to, :int32, 4
      optional :schedule_status, :string, 5
    end
    add_message "subscribe.NestedGetPaymentScheduleByCustomerData" do
      optional :total, :int32, 1
      optional :previous, :int32, 2
      optional :next, :int32, 3
      repeated :list, :message, 4, "subscribe.UnitSchedulePaymentResponse"
    end
    add_message "subscribe.GetPaymentScheduleByCustomerResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "subscribe.NestedGetPaymentScheduleByCustomerData"
    end
  end
end

module Subscribe
  OnetimePaymentRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.OnetimePaymentRequest").msgclass
  OnetimePaymentResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.OnetimePaymentResponse").msgclass
  AgainPaymentRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.AgainPaymentRequest").msgclass
  AgainPaymentResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.AgainPaymentResponse").msgclass
  PaymentScheduleParam = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.PaymentScheduleParam").msgclass
  UnitSchedulePaymentResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.UnitSchedulePaymentResponse").msgclass
  SchedulePayemntRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.SchedulePayemntRequest").msgclass
  SchedulePaymentResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.SchedulePaymentResponse").msgclass
  UnschedulePaymentRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.UnschedulePaymentRequest").msgclass
  UnschedulePaymentResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.UnschedulePaymentResponse").msgclass
  GetPaymentScheduleRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.GetPaymentScheduleRequest").msgclass
  GetPaymentScheduleResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.GetPaymentScheduleResponse").msgclass
  GetPaymentScheduleByCustomerRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.GetPaymentScheduleByCustomerRequest").msgclass
  NestedGetPaymentScheduleByCustomerData = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.NestedGetPaymentScheduleByCustomerData").msgclass
  GetPaymentScheduleByCustomerResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe.GetPaymentScheduleByCustomerResponse").msgclass
end
