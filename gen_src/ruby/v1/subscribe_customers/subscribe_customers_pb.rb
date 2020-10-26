# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/subscribe_customers/subscribe_customers.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
require 'v1/payment/payment_pb'
require 'v1/subscribe/subscribe_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("v1/subscribe_customers/subscribe_customers.proto", :syntax => :proto3) do
    add_message "subscribe_customers.CustomerBillingKey" do
      optional :card_code, :string, 1
      optional :card_name, :string, 2
      optional :card_number, :string, 3
      optional :card_type, :int32, 4
      optional :customer_addr, :string, 5
      optional :customer_email, :string, 6
      optional :customer_name, :string, 7
      optional :customer_postcode, :string, 8
      optional :customer_tel, :string, 9
      optional :customer_uid, :string, 10
      optional :inserted, :int32, 11
      optional :pg_id, :string, 12
      optional :pg_provider, :string, 13
      optional :updated, :int32, 14
    end
    add_message "subscribe_customers.GetMultipleCustomerBillingKeyRequest" do
      repeated :customer_uid, :string, 1
    end
    add_message "subscribe_customers.GetMultipleCustomerBillingKeyResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      repeated :response, :message, 3, "subscribe_customers.CustomerBillingKey"
    end
    add_message "subscribe_customers.DeleteCustomerBillingKeyRequest" do
      optional :customer_uid, :string, 1
      optional :reason, :string, 2
      optional :requester, :string, 3
    end
    add_message "subscribe_customers.DeleteCustomerBillingKeyResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "subscribe_customers.CustomerBillingKey"
    end
    add_message "subscribe_customers.GetCustomerBillingKeyRequest" do
      optional :customer_uid, :string, 1
    end
    add_message "subscribe_customers.GetCustomerBillingKeyResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "subscribe_customers.CustomerBillingKey"
    end
    add_message "subscribe_customers.InsertCustomerBillingKeyRequest" do
      optional :customer_uid, :string, 1
      optional :pg, :string, 2
      optional :card_number, :string, 3
      optional :expiry, :string, 4
      optional :birth, :string, 5
      optional :pwd_2digit, :string, 6
      optional :customer_name, :string, 7
      optional :customer_tel, :string, 8
      optional :customer_email, :string, 9
      optional :customer_addr, :string, 10
      optional :customer_postcode, :string, 11
    end
    add_message "subscribe_customers.InsertCustomerBillingKeyResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "subscribe_customers.CustomerBillingKey"
    end
    add_message "subscribe_customers.GetPaidByBillingKeyListRequest" do
      optional :customer_uid, :string, 1
      optional :page, :int32, 2
    end
    add_message "subscribe_customers.NestedGetPaidByBillingKeyListData" do
      optional :total, :int32, 1
      optional :previous, :int32, 2
      optional :next, :int32, 3
      repeated :list, :message, 4, "payment.Payment"
    end
    add_message "subscribe_customers.GetPaidByBillingKeyListResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "subscribe_customers.NestedGetPaidByBillingKeyListData"
    end
  end
end

module SubscribeCustomers
  CustomerBillingKey = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.CustomerBillingKey").msgclass
  GetMultipleCustomerBillingKeyRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.GetMultipleCustomerBillingKeyRequest").msgclass
  GetMultipleCustomerBillingKeyResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.GetMultipleCustomerBillingKeyResponse").msgclass
  DeleteCustomerBillingKeyRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.DeleteCustomerBillingKeyRequest").msgclass
  DeleteCustomerBillingKeyResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.DeleteCustomerBillingKeyResponse").msgclass
  GetCustomerBillingKeyRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.GetCustomerBillingKeyRequest").msgclass
  GetCustomerBillingKeyResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.GetCustomerBillingKeyResponse").msgclass
  InsertCustomerBillingKeyRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.InsertCustomerBillingKeyRequest").msgclass
  InsertCustomerBillingKeyResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.InsertCustomerBillingKeyResponse").msgclass
  GetPaidByBillingKeyListRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.GetPaidByBillingKeyListRequest").msgclass
  NestedGetPaidByBillingKeyListData = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.NestedGetPaidByBillingKeyListData").msgclass
  GetPaidByBillingKeyListResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("subscribe_customers.GetPaidByBillingKeyListResponse").msgclass
end
