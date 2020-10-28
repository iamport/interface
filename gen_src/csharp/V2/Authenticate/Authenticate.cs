// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: v2/authenticate/authenticate.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace V2.Authenticate {

  /// <summary>Holder for reflection information generated from v2/authenticate/authenticate.proto</summary>
  public static partial class AuthenticateReflection {

    #region Descriptor
    /// <summary>File descriptor for v2/authenticate/authenticate.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static AuthenticateReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiJ2Mi9hdXRoZW50aWNhdGUvYXV0aGVudGljYXRlLnByb3RvEg9hdXRoZW50",
            "aWNhdGVfdjIaHGdvb2dsZS9hcGkvYW5ub3RhdGlvbnMucHJvdG8iMwoMVG9r",
            "ZW5SZXF1ZXN0Eg8KB2ltcF9rZXkYASABKAkSEgoKaW1wX3NlY3JldBgCIAEo",
            "CSJCCglUb2tlbkRhdGESFAoMYWNjZXNzX3Rva2VuGAEgASgJEhIKCmV4cGly",
            "ZWRfYXQYAiABKAUSCwoDbm93GAMgASgFIlwKDVRva2VuUmVzcG9uc2USDAoE",
            "Y29kZRgBIAEoBRIPCgdtZXNzYWdlGAIgASgJEiwKCHJlc3BvbnNlGAMgASgL",
            "MhouYXV0aGVudGljYXRlX3YyLlRva2VuRGF0YSJOChVQdWJLZXlSZWdpc3Rl",
            "clJlcXVlc3QSDwoHaW1wX2tleRgBIAEoCRISCgpwdWJsaWNfa2V5GAIgASgJ",
            "EhAKCHBhc3N3b3JkGAMgASgJIjcKFlB1YktleVJlZ2lzdGVyUmVzcG9uc2US",
            "DAoEY29kZRgBIAEoBRIPCgdtZXNzYWdlGAIgASgJMpMCChNBdXRoZW50aWNh",
            "dGVTZXJ2aWNlEm4KClRva2VuVjJSUEMSHS5hdXRoZW50aWNhdGVfdjIuVG9r",
            "ZW5SZXF1ZXN0Gh4uYXV0aGVudGljYXRlX3YyLlRva2VuUmVzcG9uc2UiIYLT",
            "5JMCGyIWL2FwaS92Mi91c2Vycy9nZXRUb2tlbjoBKhKLAQoTUmVnaXN0ZXJQ",
            "dWJLZXlWMlJQQxImLmF1dGhlbnRpY2F0ZV92Mi5QdWJLZXlSZWdpc3RlclJl",
            "cXVlc3QaJy5hdXRoZW50aWNhdGVfdjIuUHViS2V5UmVnaXN0ZXJSZXNwb25z",
            "ZSIjgtPkkwIdIhgvYXBpL3YyL3VzZXJzL3B1YmxpY19rZXk6ASpCS1o3Z2l0",
            "aHViLmNvbS9pYW1wb3J0L2ludGVyZmFjZS9nZW5fc3JjL2dvL3YyL2F1dGhl",
            "bnRpY2F0ZaoCD1YyLkF1dGhlbnRpY2F0ZWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Api.Annotations.AnnotationsReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::V2.Authenticate.TokenRequest), global::V2.Authenticate.TokenRequest.Parser, new[]{ "ImpKey", "ImpSecret" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::V2.Authenticate.TokenData), global::V2.Authenticate.TokenData.Parser, new[]{ "AccessToken", "ExpiredAt", "Now" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::V2.Authenticate.TokenResponse), global::V2.Authenticate.TokenResponse.Parser, new[]{ "Code", "Message", "Response" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::V2.Authenticate.PubKeyRegisterRequest), global::V2.Authenticate.PubKeyRegisterRequest.Parser, new[]{ "ImpKey", "PublicKey", "Password" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::V2.Authenticate.PubKeyRegisterResponse), global::V2.Authenticate.PubKeyRegisterResponse.Parser, new[]{ "Code", "Message" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class TokenRequest : pb::IMessage<TokenRequest>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<TokenRequest> _parser = new pb::MessageParser<TokenRequest>(() => new TokenRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<TokenRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::V2.Authenticate.AuthenticateReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public TokenRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public TokenRequest(TokenRequest other) : this() {
      impKey_ = other.impKey_;
      impSecret_ = other.impSecret_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public TokenRequest Clone() {
      return new TokenRequest(this);
    }

    /// <summary>Field number for the "imp_key" field.</summary>
    public const int ImpKeyFieldNumber = 1;
    private string impKey_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string ImpKey {
      get { return impKey_; }
      set {
        impKey_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "imp_secret" field.</summary>
    public const int ImpSecretFieldNumber = 2;
    private string impSecret_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string ImpSecret {
      get { return impSecret_; }
      set {
        impSecret_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as TokenRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(TokenRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (ImpKey != other.ImpKey) return false;
      if (ImpSecret != other.ImpSecret) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (ImpKey.Length != 0) hash ^= ImpKey.GetHashCode();
      if (ImpSecret.Length != 0) hash ^= ImpSecret.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (ImpKey.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(ImpKey);
      }
      if (ImpSecret.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(ImpSecret);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (ImpKey.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(ImpKey);
      }
      if (ImpSecret.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(ImpSecret);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (ImpKey.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(ImpKey);
      }
      if (ImpSecret.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(ImpSecret);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(TokenRequest other) {
      if (other == null) {
        return;
      }
      if (other.ImpKey.Length != 0) {
        ImpKey = other.ImpKey;
      }
      if (other.ImpSecret.Length != 0) {
        ImpSecret = other.ImpSecret;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            ImpKey = input.ReadString();
            break;
          }
          case 18: {
            ImpSecret = input.ReadString();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            ImpKey = input.ReadString();
            break;
          }
          case 18: {
            ImpSecret = input.ReadString();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class TokenData : pb::IMessage<TokenData>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<TokenData> _parser = new pb::MessageParser<TokenData>(() => new TokenData());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<TokenData> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::V2.Authenticate.AuthenticateReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public TokenData() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public TokenData(TokenData other) : this() {
      accessToken_ = other.accessToken_;
      expiredAt_ = other.expiredAt_;
      now_ = other.now_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public TokenData Clone() {
      return new TokenData(this);
    }

    /// <summary>Field number for the "access_token" field.</summary>
    public const int AccessTokenFieldNumber = 1;
    private string accessToken_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string AccessToken {
      get { return accessToken_; }
      set {
        accessToken_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "expired_at" field.</summary>
    public const int ExpiredAtFieldNumber = 2;
    private int expiredAt_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int ExpiredAt {
      get { return expiredAt_; }
      set {
        expiredAt_ = value;
      }
    }

    /// <summary>Field number for the "now" field.</summary>
    public const int NowFieldNumber = 3;
    private int now_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Now {
      get { return now_; }
      set {
        now_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as TokenData);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(TokenData other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (AccessToken != other.AccessToken) return false;
      if (ExpiredAt != other.ExpiredAt) return false;
      if (Now != other.Now) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (AccessToken.Length != 0) hash ^= AccessToken.GetHashCode();
      if (ExpiredAt != 0) hash ^= ExpiredAt.GetHashCode();
      if (Now != 0) hash ^= Now.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (AccessToken.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(AccessToken);
      }
      if (ExpiredAt != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(ExpiredAt);
      }
      if (Now != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(Now);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (AccessToken.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(AccessToken);
      }
      if (ExpiredAt != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(ExpiredAt);
      }
      if (Now != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(Now);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (AccessToken.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(AccessToken);
      }
      if (ExpiredAt != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(ExpiredAt);
      }
      if (Now != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Now);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(TokenData other) {
      if (other == null) {
        return;
      }
      if (other.AccessToken.Length != 0) {
        AccessToken = other.AccessToken;
      }
      if (other.ExpiredAt != 0) {
        ExpiredAt = other.ExpiredAt;
      }
      if (other.Now != 0) {
        Now = other.Now;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            AccessToken = input.ReadString();
            break;
          }
          case 16: {
            ExpiredAt = input.ReadInt32();
            break;
          }
          case 24: {
            Now = input.ReadInt32();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            AccessToken = input.ReadString();
            break;
          }
          case 16: {
            ExpiredAt = input.ReadInt32();
            break;
          }
          case 24: {
            Now = input.ReadInt32();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class TokenResponse : pb::IMessage<TokenResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<TokenResponse> _parser = new pb::MessageParser<TokenResponse>(() => new TokenResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<TokenResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::V2.Authenticate.AuthenticateReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public TokenResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public TokenResponse(TokenResponse other) : this() {
      code_ = other.code_;
      message_ = other.message_;
      response_ = other.response_ != null ? other.response_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public TokenResponse Clone() {
      return new TokenResponse(this);
    }

    /// <summary>Field number for the "code" field.</summary>
    public const int CodeFieldNumber = 1;
    private int code_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Code {
      get { return code_; }
      set {
        code_ = value;
      }
    }

    /// <summary>Field number for the "message" field.</summary>
    public const int MessageFieldNumber = 2;
    private string message_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Message {
      get { return message_; }
      set {
        message_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "response" field.</summary>
    public const int ResponseFieldNumber = 3;
    private global::V2.Authenticate.TokenData response_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::V2.Authenticate.TokenData Response {
      get { return response_; }
      set {
        response_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as TokenResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(TokenResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Code != other.Code) return false;
      if (Message != other.Message) return false;
      if (!object.Equals(Response, other.Response)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Code != 0) hash ^= Code.GetHashCode();
      if (Message.Length != 0) hash ^= Message.GetHashCode();
      if (response_ != null) hash ^= Response.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (Code != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(Code);
      }
      if (Message.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Message);
      }
      if (response_ != null) {
        output.WriteRawTag(26);
        output.WriteMessage(Response);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (Code != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(Code);
      }
      if (Message.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Message);
      }
      if (response_ != null) {
        output.WriteRawTag(26);
        output.WriteMessage(Response);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Code != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Code);
      }
      if (Message.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Message);
      }
      if (response_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Response);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(TokenResponse other) {
      if (other == null) {
        return;
      }
      if (other.Code != 0) {
        Code = other.Code;
      }
      if (other.Message.Length != 0) {
        Message = other.Message;
      }
      if (other.response_ != null) {
        if (response_ == null) {
          Response = new global::V2.Authenticate.TokenData();
        }
        Response.MergeFrom(other.Response);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Code = input.ReadInt32();
            break;
          }
          case 18: {
            Message = input.ReadString();
            break;
          }
          case 26: {
            if (response_ == null) {
              Response = new global::V2.Authenticate.TokenData();
            }
            input.ReadMessage(Response);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 8: {
            Code = input.ReadInt32();
            break;
          }
          case 18: {
            Message = input.ReadString();
            break;
          }
          case 26: {
            if (response_ == null) {
              Response = new global::V2.Authenticate.TokenData();
            }
            input.ReadMessage(Response);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class PubKeyRegisterRequest : pb::IMessage<PubKeyRegisterRequest>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<PubKeyRegisterRequest> _parser = new pb::MessageParser<PubKeyRegisterRequest>(() => new PubKeyRegisterRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<PubKeyRegisterRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::V2.Authenticate.AuthenticateReflection.Descriptor.MessageTypes[3]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PubKeyRegisterRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PubKeyRegisterRequest(PubKeyRegisterRequest other) : this() {
      impKey_ = other.impKey_;
      publicKey_ = other.publicKey_;
      password_ = other.password_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PubKeyRegisterRequest Clone() {
      return new PubKeyRegisterRequest(this);
    }

    /// <summary>Field number for the "imp_key" field.</summary>
    public const int ImpKeyFieldNumber = 1;
    private string impKey_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string ImpKey {
      get { return impKey_; }
      set {
        impKey_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "public_key" field.</summary>
    public const int PublicKeyFieldNumber = 2;
    private string publicKey_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string PublicKey {
      get { return publicKey_; }
      set {
        publicKey_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "password" field.</summary>
    public const int PasswordFieldNumber = 3;
    private string password_ = "";
    /// <summary>
    /// want to make it work without session dependency
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Password {
      get { return password_; }
      set {
        password_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as PubKeyRegisterRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(PubKeyRegisterRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (ImpKey != other.ImpKey) return false;
      if (PublicKey != other.PublicKey) return false;
      if (Password != other.Password) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (ImpKey.Length != 0) hash ^= ImpKey.GetHashCode();
      if (PublicKey.Length != 0) hash ^= PublicKey.GetHashCode();
      if (Password.Length != 0) hash ^= Password.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (ImpKey.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(ImpKey);
      }
      if (PublicKey.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(PublicKey);
      }
      if (Password.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Password);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (ImpKey.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(ImpKey);
      }
      if (PublicKey.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(PublicKey);
      }
      if (Password.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Password);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (ImpKey.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(ImpKey);
      }
      if (PublicKey.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(PublicKey);
      }
      if (Password.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Password);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(PubKeyRegisterRequest other) {
      if (other == null) {
        return;
      }
      if (other.ImpKey.Length != 0) {
        ImpKey = other.ImpKey;
      }
      if (other.PublicKey.Length != 0) {
        PublicKey = other.PublicKey;
      }
      if (other.Password.Length != 0) {
        Password = other.Password;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            ImpKey = input.ReadString();
            break;
          }
          case 18: {
            PublicKey = input.ReadString();
            break;
          }
          case 26: {
            Password = input.ReadString();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            ImpKey = input.ReadString();
            break;
          }
          case 18: {
            PublicKey = input.ReadString();
            break;
          }
          case 26: {
            Password = input.ReadString();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class PubKeyRegisterResponse : pb::IMessage<PubKeyRegisterResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<PubKeyRegisterResponse> _parser = new pb::MessageParser<PubKeyRegisterResponse>(() => new PubKeyRegisterResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<PubKeyRegisterResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::V2.Authenticate.AuthenticateReflection.Descriptor.MessageTypes[4]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PubKeyRegisterResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PubKeyRegisterResponse(PubKeyRegisterResponse other) : this() {
      code_ = other.code_;
      message_ = other.message_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PubKeyRegisterResponse Clone() {
      return new PubKeyRegisterResponse(this);
    }

    /// <summary>Field number for the "code" field.</summary>
    public const int CodeFieldNumber = 1;
    private int code_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Code {
      get { return code_; }
      set {
        code_ = value;
      }
    }

    /// <summary>Field number for the "message" field.</summary>
    public const int MessageFieldNumber = 2;
    private string message_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Message {
      get { return message_; }
      set {
        message_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as PubKeyRegisterResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(PubKeyRegisterResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Code != other.Code) return false;
      if (Message != other.Message) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Code != 0) hash ^= Code.GetHashCode();
      if (Message.Length != 0) hash ^= Message.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (Code != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(Code);
      }
      if (Message.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Message);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (Code != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(Code);
      }
      if (Message.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Message);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Code != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Code);
      }
      if (Message.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Message);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(PubKeyRegisterResponse other) {
      if (other == null) {
        return;
      }
      if (other.Code != 0) {
        Code = other.Code;
      }
      if (other.Message.Length != 0) {
        Message = other.Message;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Code = input.ReadInt32();
            break;
          }
          case 18: {
            Message = input.ReadString();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 8: {
            Code = input.ReadInt32();
            break;
          }
          case 18: {
            Message = input.ReadString();
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code