# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v2/authenticate/authenticate.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("v2/authenticate/authenticate.proto", :syntax => :proto3) do
    add_message "authenticate_v2.TokenRequest" do
      optional :imp_key, :string, 1
      optional :imp_secret, :string, 2
    end
    add_message "authenticate_v2.TokenData" do
      optional :access_token, :string, 1
      optional :expired_at, :int32, 2
      optional :now, :int32, 3
    end
    add_message "authenticate_v2.TokenResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "authenticate_v2.TokenData"
    end
    add_message "authenticate_v2.PubKeyRegisterRequest" do
      optional :imp_key, :string, 1
      optional :public_key, :string, 2
      optional :password, :string, 3
    end
    add_message "authenticate_v2.PubKeyRegisterResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
    end
  end
end

module AuthenticateV2
  TokenRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("authenticate_v2.TokenRequest").msgclass
  TokenData = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("authenticate_v2.TokenData").msgclass
  TokenResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("authenticate_v2.TokenResponse").msgclass
  PubKeyRegisterRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("authenticate_v2.PubKeyRegisterRequest").msgclass
  PubKeyRegisterResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("authenticate_v2.PubKeyRegisterResponse").msgclass
end
