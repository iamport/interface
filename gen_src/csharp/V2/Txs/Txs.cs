// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: v2/txs/txs.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace V2.Txs {

  /// <summary>Holder for reflection information generated from v2/txs/txs.proto</summary>
  public static partial class TxsReflection {

    #region Descriptor
    /// <summary>File descriptor for v2/txs/txs.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static TxsReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChB2Mi90eHMvdHhzLnByb3RvEgZ0eHNfdjIaHGdvb2dsZS9hcGkvYW5ub3Rh",
            "dGlvbnMucHJvdG8aFHYyL2Jhc2lzL2Jhc2lzLnByb3RvIrUCChRHZXRUeHNS",
            "ZWNvcmRzUmVxdWVzdBIPCgdpbXBfdWlkGAEgASgJEhQKDG1lcmNoYW50X3Vp",
            "ZBgCIAEoCRIPCgdwYWlkX2J5GAMgASgJEg4KBnN0YXR1cxgEIAMoCRIMCgRw",
            "YWdlGAUgASgFEg8KB3NvcnRpbmcYBiABKAkSDQoFbGltaXQYByABKAUSEwoL",
            "cGdfcHJvdmlkZXIYCCADKAkSEgoKcGF5X21ldGhvZBgJIAMoCRISCgpidXll",
            "cl9uYW1lGAogASgJEhEKCWJ1eWVyX3RlbBgLIAEoCRITCgtidXllcl9lbWFp",
            "bBgMIAEoCRINCgVzdGFydBgNIAEoCRILCgNlbmQYDiABKAkSDwoHc2FuZGJv",
            "eBgPIAEoCBIVCg1vdXRwdXRfZm9ybWF0GBAgAygJIlkKFEdldFR4c1JlY29y",
            "ZHNSZXBvbnNlEgwKBGNvZGUYASABKAUSDwoHbWVzc2FnZRgCIAEoCRIiCghy",
            "ZXNwb25zZRgDIAMoCzIQLmJhc2lzX3YyLlVuaXRUeDJ9CgpUeHNTZXJ2aWNl",
            "Em8KE0dldFR4c1JlY29yZFNlcnZpY2USHC50eHNfdjIuR2V0VHhzUmVjb3Jk",
            "c1JlcXVlc3QaHC50eHNfdjIuR2V0VHhzUmVjb3Jkc1JlcG9uc2UiHILT5JMC",
            "FhIUL2FwaS9wYXltZW50cy92Mi90eHNCOVouZ2l0aHViLmNvbS9pYW1wb3J0",
            "L2ludGVyZmFjZS9nZW5fc3JjL2dvL3YyL3R4c6oCBlYyLlR4c2IGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Api.Annotations.AnnotationsReflection.Descriptor, global::V2.Basis.BasisReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::V2.Txs.GetTxsRecordsRequest), global::V2.Txs.GetTxsRecordsRequest.Parser, new[]{ "ImpUid", "MerchantUid", "PaidBy", "Status", "Page", "Sorting", "Limit", "PgProvider", "PayMethod", "BuyerName", "BuyerTel", "BuyerEmail", "Start", "End", "Sandbox", "OutputFormat" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::V2.Txs.GetTxsRecordsReponse), global::V2.Txs.GetTxsRecordsReponse.Parser, new[]{ "Code", "Message", "Response" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class GetTxsRecordsRequest : pb::IMessage<GetTxsRecordsRequest>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<GetTxsRecordsRequest> _parser = new pb::MessageParser<GetTxsRecordsRequest>(() => new GetTxsRecordsRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<GetTxsRecordsRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::V2.Txs.TxsReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetTxsRecordsRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetTxsRecordsRequest(GetTxsRecordsRequest other) : this() {
      impUid_ = other.impUid_;
      merchantUid_ = other.merchantUid_;
      paidBy_ = other.paidBy_;
      status_ = other.status_.Clone();
      page_ = other.page_;
      sorting_ = other.sorting_;
      limit_ = other.limit_;
      pgProvider_ = other.pgProvider_.Clone();
      payMethod_ = other.payMethod_.Clone();
      buyerName_ = other.buyerName_;
      buyerTel_ = other.buyerTel_;
      buyerEmail_ = other.buyerEmail_;
      start_ = other.start_;
      end_ = other.end_;
      sandbox_ = other.sandbox_;
      outputFormat_ = other.outputFormat_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetTxsRecordsRequest Clone() {
      return new GetTxsRecordsRequest(this);
    }

    /// <summary>Field number for the "imp_uid" field.</summary>
    public const int ImpUidFieldNumber = 1;
    private string impUid_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string ImpUid {
      get { return impUid_; }
      set {
        impUid_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "merchant_uid" field.</summary>
    public const int MerchantUidFieldNumber = 2;
    private string merchantUid_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string MerchantUid {
      get { return merchantUid_; }
      set {
        merchantUid_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "paid_by" field.</summary>
    public const int PaidByFieldNumber = 3;
    private string paidBy_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string PaidBy {
      get { return paidBy_; }
      set {
        paidBy_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "status" field.</summary>
    public const int StatusFieldNumber = 4;
    private static readonly pb::FieldCodec<string> _repeated_status_codec
        = pb::FieldCodec.ForString(34);
    private readonly pbc::RepeatedField<string> status_ = new pbc::RepeatedField<string>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<string> Status {
      get { return status_; }
    }

    /// <summary>Field number for the "page" field.</summary>
    public const int PageFieldNumber = 5;
    private int page_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Page {
      get { return page_; }
      set {
        page_ = value;
      }
    }

    /// <summary>Field number for the "sorting" field.</summary>
    public const int SortingFieldNumber = 6;
    private string sorting_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Sorting {
      get { return sorting_; }
      set {
        sorting_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "limit" field.</summary>
    public const int LimitFieldNumber = 7;
    private int limit_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Limit {
      get { return limit_; }
      set {
        limit_ = value;
      }
    }

    /// <summary>Field number for the "pg_provider" field.</summary>
    public const int PgProviderFieldNumber = 8;
    private static readonly pb::FieldCodec<string> _repeated_pgProvider_codec
        = pb::FieldCodec.ForString(66);
    private readonly pbc::RepeatedField<string> pgProvider_ = new pbc::RepeatedField<string>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<string> PgProvider {
      get { return pgProvider_; }
    }

    /// <summary>Field number for the "pay_method" field.</summary>
    public const int PayMethodFieldNumber = 9;
    private static readonly pb::FieldCodec<string> _repeated_payMethod_codec
        = pb::FieldCodec.ForString(74);
    private readonly pbc::RepeatedField<string> payMethod_ = new pbc::RepeatedField<string>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<string> PayMethod {
      get { return payMethod_; }
    }

    /// <summary>Field number for the "buyer_name" field.</summary>
    public const int BuyerNameFieldNumber = 10;
    private string buyerName_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string BuyerName {
      get { return buyerName_; }
      set {
        buyerName_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "buyer_tel" field.</summary>
    public const int BuyerTelFieldNumber = 11;
    private string buyerTel_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string BuyerTel {
      get { return buyerTel_; }
      set {
        buyerTel_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "buyer_email" field.</summary>
    public const int BuyerEmailFieldNumber = 12;
    private string buyerEmail_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string BuyerEmail {
      get { return buyerEmail_; }
      set {
        buyerEmail_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "start" field.</summary>
    public const int StartFieldNumber = 13;
    private string start_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Start {
      get { return start_; }
      set {
        start_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "end" field.</summary>
    public const int EndFieldNumber = 14;
    private string end_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string End {
      get { return end_; }
      set {
        end_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "sandbox" field.</summary>
    public const int SandboxFieldNumber = 15;
    private bool sandbox_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Sandbox {
      get { return sandbox_; }
      set {
        sandbox_ = value;
      }
    }

    /// <summary>Field number for the "output_format" field.</summary>
    public const int OutputFormatFieldNumber = 16;
    private static readonly pb::FieldCodec<string> _repeated_outputFormat_codec
        = pb::FieldCodec.ForString(130);
    private readonly pbc::RepeatedField<string> outputFormat_ = new pbc::RepeatedField<string>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<string> OutputFormat {
      get { return outputFormat_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as GetTxsRecordsRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(GetTxsRecordsRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (ImpUid != other.ImpUid) return false;
      if (MerchantUid != other.MerchantUid) return false;
      if (PaidBy != other.PaidBy) return false;
      if(!status_.Equals(other.status_)) return false;
      if (Page != other.Page) return false;
      if (Sorting != other.Sorting) return false;
      if (Limit != other.Limit) return false;
      if(!pgProvider_.Equals(other.pgProvider_)) return false;
      if(!payMethod_.Equals(other.payMethod_)) return false;
      if (BuyerName != other.BuyerName) return false;
      if (BuyerTel != other.BuyerTel) return false;
      if (BuyerEmail != other.BuyerEmail) return false;
      if (Start != other.Start) return false;
      if (End != other.End) return false;
      if (Sandbox != other.Sandbox) return false;
      if(!outputFormat_.Equals(other.outputFormat_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (ImpUid.Length != 0) hash ^= ImpUid.GetHashCode();
      if (MerchantUid.Length != 0) hash ^= MerchantUid.GetHashCode();
      if (PaidBy.Length != 0) hash ^= PaidBy.GetHashCode();
      hash ^= status_.GetHashCode();
      if (Page != 0) hash ^= Page.GetHashCode();
      if (Sorting.Length != 0) hash ^= Sorting.GetHashCode();
      if (Limit != 0) hash ^= Limit.GetHashCode();
      hash ^= pgProvider_.GetHashCode();
      hash ^= payMethod_.GetHashCode();
      if (BuyerName.Length != 0) hash ^= BuyerName.GetHashCode();
      if (BuyerTel.Length != 0) hash ^= BuyerTel.GetHashCode();
      if (BuyerEmail.Length != 0) hash ^= BuyerEmail.GetHashCode();
      if (Start.Length != 0) hash ^= Start.GetHashCode();
      if (End.Length != 0) hash ^= End.GetHashCode();
      if (Sandbox != false) hash ^= Sandbox.GetHashCode();
      hash ^= outputFormat_.GetHashCode();
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
      if (ImpUid.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(ImpUid);
      }
      if (MerchantUid.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(MerchantUid);
      }
      if (PaidBy.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(PaidBy);
      }
      status_.WriteTo(output, _repeated_status_codec);
      if (Page != 0) {
        output.WriteRawTag(40);
        output.WriteInt32(Page);
      }
      if (Sorting.Length != 0) {
        output.WriteRawTag(50);
        output.WriteString(Sorting);
      }
      if (Limit != 0) {
        output.WriteRawTag(56);
        output.WriteInt32(Limit);
      }
      pgProvider_.WriteTo(output, _repeated_pgProvider_codec);
      payMethod_.WriteTo(output, _repeated_payMethod_codec);
      if (BuyerName.Length != 0) {
        output.WriteRawTag(82);
        output.WriteString(BuyerName);
      }
      if (BuyerTel.Length != 0) {
        output.WriteRawTag(90);
        output.WriteString(BuyerTel);
      }
      if (BuyerEmail.Length != 0) {
        output.WriteRawTag(98);
        output.WriteString(BuyerEmail);
      }
      if (Start.Length != 0) {
        output.WriteRawTag(106);
        output.WriteString(Start);
      }
      if (End.Length != 0) {
        output.WriteRawTag(114);
        output.WriteString(End);
      }
      if (Sandbox != false) {
        output.WriteRawTag(120);
        output.WriteBool(Sandbox);
      }
      outputFormat_.WriteTo(output, _repeated_outputFormat_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (ImpUid.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(ImpUid);
      }
      if (MerchantUid.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(MerchantUid);
      }
      if (PaidBy.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(PaidBy);
      }
      status_.WriteTo(ref output, _repeated_status_codec);
      if (Page != 0) {
        output.WriteRawTag(40);
        output.WriteInt32(Page);
      }
      if (Sorting.Length != 0) {
        output.WriteRawTag(50);
        output.WriteString(Sorting);
      }
      if (Limit != 0) {
        output.WriteRawTag(56);
        output.WriteInt32(Limit);
      }
      pgProvider_.WriteTo(ref output, _repeated_pgProvider_codec);
      payMethod_.WriteTo(ref output, _repeated_payMethod_codec);
      if (BuyerName.Length != 0) {
        output.WriteRawTag(82);
        output.WriteString(BuyerName);
      }
      if (BuyerTel.Length != 0) {
        output.WriteRawTag(90);
        output.WriteString(BuyerTel);
      }
      if (BuyerEmail.Length != 0) {
        output.WriteRawTag(98);
        output.WriteString(BuyerEmail);
      }
      if (Start.Length != 0) {
        output.WriteRawTag(106);
        output.WriteString(Start);
      }
      if (End.Length != 0) {
        output.WriteRawTag(114);
        output.WriteString(End);
      }
      if (Sandbox != false) {
        output.WriteRawTag(120);
        output.WriteBool(Sandbox);
      }
      outputFormat_.WriteTo(ref output, _repeated_outputFormat_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (ImpUid.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(ImpUid);
      }
      if (MerchantUid.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(MerchantUid);
      }
      if (PaidBy.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(PaidBy);
      }
      size += status_.CalculateSize(_repeated_status_codec);
      if (Page != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Page);
      }
      if (Sorting.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Sorting);
      }
      if (Limit != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Limit);
      }
      size += pgProvider_.CalculateSize(_repeated_pgProvider_codec);
      size += payMethod_.CalculateSize(_repeated_payMethod_codec);
      if (BuyerName.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(BuyerName);
      }
      if (BuyerTel.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(BuyerTel);
      }
      if (BuyerEmail.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(BuyerEmail);
      }
      if (Start.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Start);
      }
      if (End.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(End);
      }
      if (Sandbox != false) {
        size += 1 + 1;
      }
      size += outputFormat_.CalculateSize(_repeated_outputFormat_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(GetTxsRecordsRequest other) {
      if (other == null) {
        return;
      }
      if (other.ImpUid.Length != 0) {
        ImpUid = other.ImpUid;
      }
      if (other.MerchantUid.Length != 0) {
        MerchantUid = other.MerchantUid;
      }
      if (other.PaidBy.Length != 0) {
        PaidBy = other.PaidBy;
      }
      status_.Add(other.status_);
      if (other.Page != 0) {
        Page = other.Page;
      }
      if (other.Sorting.Length != 0) {
        Sorting = other.Sorting;
      }
      if (other.Limit != 0) {
        Limit = other.Limit;
      }
      pgProvider_.Add(other.pgProvider_);
      payMethod_.Add(other.payMethod_);
      if (other.BuyerName.Length != 0) {
        BuyerName = other.BuyerName;
      }
      if (other.BuyerTel.Length != 0) {
        BuyerTel = other.BuyerTel;
      }
      if (other.BuyerEmail.Length != 0) {
        BuyerEmail = other.BuyerEmail;
      }
      if (other.Start.Length != 0) {
        Start = other.Start;
      }
      if (other.End.Length != 0) {
        End = other.End;
      }
      if (other.Sandbox != false) {
        Sandbox = other.Sandbox;
      }
      outputFormat_.Add(other.outputFormat_);
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
            ImpUid = input.ReadString();
            break;
          }
          case 18: {
            MerchantUid = input.ReadString();
            break;
          }
          case 26: {
            PaidBy = input.ReadString();
            break;
          }
          case 34: {
            status_.AddEntriesFrom(input, _repeated_status_codec);
            break;
          }
          case 40: {
            Page = input.ReadInt32();
            break;
          }
          case 50: {
            Sorting = input.ReadString();
            break;
          }
          case 56: {
            Limit = input.ReadInt32();
            break;
          }
          case 66: {
            pgProvider_.AddEntriesFrom(input, _repeated_pgProvider_codec);
            break;
          }
          case 74: {
            payMethod_.AddEntriesFrom(input, _repeated_payMethod_codec);
            break;
          }
          case 82: {
            BuyerName = input.ReadString();
            break;
          }
          case 90: {
            BuyerTel = input.ReadString();
            break;
          }
          case 98: {
            BuyerEmail = input.ReadString();
            break;
          }
          case 106: {
            Start = input.ReadString();
            break;
          }
          case 114: {
            End = input.ReadString();
            break;
          }
          case 120: {
            Sandbox = input.ReadBool();
            break;
          }
          case 130: {
            outputFormat_.AddEntriesFrom(input, _repeated_outputFormat_codec);
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
            ImpUid = input.ReadString();
            break;
          }
          case 18: {
            MerchantUid = input.ReadString();
            break;
          }
          case 26: {
            PaidBy = input.ReadString();
            break;
          }
          case 34: {
            status_.AddEntriesFrom(ref input, _repeated_status_codec);
            break;
          }
          case 40: {
            Page = input.ReadInt32();
            break;
          }
          case 50: {
            Sorting = input.ReadString();
            break;
          }
          case 56: {
            Limit = input.ReadInt32();
            break;
          }
          case 66: {
            pgProvider_.AddEntriesFrom(ref input, _repeated_pgProvider_codec);
            break;
          }
          case 74: {
            payMethod_.AddEntriesFrom(ref input, _repeated_payMethod_codec);
            break;
          }
          case 82: {
            BuyerName = input.ReadString();
            break;
          }
          case 90: {
            BuyerTel = input.ReadString();
            break;
          }
          case 98: {
            BuyerEmail = input.ReadString();
            break;
          }
          case 106: {
            Start = input.ReadString();
            break;
          }
          case 114: {
            End = input.ReadString();
            break;
          }
          case 120: {
            Sandbox = input.ReadBool();
            break;
          }
          case 130: {
            outputFormat_.AddEntriesFrom(ref input, _repeated_outputFormat_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class GetTxsRecordsReponse : pb::IMessage<GetTxsRecordsReponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<GetTxsRecordsReponse> _parser = new pb::MessageParser<GetTxsRecordsReponse>(() => new GetTxsRecordsReponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<GetTxsRecordsReponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::V2.Txs.TxsReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetTxsRecordsReponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetTxsRecordsReponse(GetTxsRecordsReponse other) : this() {
      code_ = other.code_;
      message_ = other.message_;
      response_ = other.response_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetTxsRecordsReponse Clone() {
      return new GetTxsRecordsReponse(this);
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
    private static readonly pb::FieldCodec<global::V2.Basis.UnitTx> _repeated_response_codec
        = pb::FieldCodec.ForMessage(26, global::V2.Basis.UnitTx.Parser);
    private readonly pbc::RepeatedField<global::V2.Basis.UnitTx> response_ = new pbc::RepeatedField<global::V2.Basis.UnitTx>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::V2.Basis.UnitTx> Response {
      get { return response_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as GetTxsRecordsReponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(GetTxsRecordsReponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Code != other.Code) return false;
      if (Message != other.Message) return false;
      if(!response_.Equals(other.response_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Code != 0) hash ^= Code.GetHashCode();
      if (Message.Length != 0) hash ^= Message.GetHashCode();
      hash ^= response_.GetHashCode();
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
      response_.WriteTo(output, _repeated_response_codec);
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
      response_.WriteTo(ref output, _repeated_response_codec);
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
      size += response_.CalculateSize(_repeated_response_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(GetTxsRecordsReponse other) {
      if (other == null) {
        return;
      }
      if (other.Code != 0) {
        Code = other.Code;
      }
      if (other.Message.Length != 0) {
        Message = other.Message;
      }
      response_.Add(other.response_);
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
            response_.AddEntriesFrom(input, _repeated_response_codec);
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
            response_.AddEntriesFrom(ref input, _repeated_response_codec);
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
