# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/authenticate/token.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("v1/authenticate/token.proto", :syntax => :proto3) do
    add_message "authenticate.Token" do
      optional :access_token, :string, 1
      optional :now, :int32, 2
      optional :expired_at, :int32, 3
    end
    add_message "authenticate.TokenRequest" do
      optional :imp_key, :string, 1
      optional :imp_secret, :string, 2
    end
    add_message "authenticate.TokenResponse" do
      optional :code, :int32, 1
      optional :message, :string, 2
      optional :response, :message, 3, "authenticate.Token"
    end
  end
end

module Authenticate
  Token = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("authenticate.Token").msgclass
  TokenRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("authenticate.TokenRequest").msgclass
  TokenResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("authenticate.TokenResponse").msgclass
end
