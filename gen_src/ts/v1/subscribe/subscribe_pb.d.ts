// package: subscribe
// file: v1/subscribe/subscribe.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../../google/api/annotations_pb";
import * as v1_payment_payment_pb from "../../v1/payment/payment_pb";

export class OnetimePaymentRequest extends jspb.Message {
  getMerchantUid(): string;
  setMerchantUid(value: string): void;

  getAmount(): number;
  setAmount(value: number): void;

  getTaxFree(): number;
  setTaxFree(value: number): void;

  getCardNumber(): string;
  setCardNumber(value: string): void;

  getExpiry(): string;
  setExpiry(value: string): void;

  getBirth(): string;
  setBirth(value: string): void;

  getPwd2digit(): string;
  setPwd2digit(value: string): void;

  getCustomerUid(): string;
  setCustomerUid(value: string): void;

  getPg(): string;
  setPg(value: string): void;

  getName(): string;
  setName(value: string): void;

  getBuyerName(): string;
  setBuyerName(value: string): void;

  getBuyerEmail(): string;
  setBuyerEmail(value: string): void;

  getBuyerTel(): string;
  setBuyerTel(value: string): void;

  getBuyerAddr(): string;
  setBuyerAddr(value: string): void;

  getBuyerPostcode(): string;
  setBuyerPostcode(value: string): void;

  getCardQuota(): number;
  setCardQuota(value: number): void;

  getInterestFreeByMerchant(): boolean;
  setInterestFreeByMerchant(value: boolean): void;

  getCustomData(): string;
  setCustomData(value: string): void;

  getNoticeUrl(): string;
  setNoticeUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OnetimePaymentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: OnetimePaymentRequest): OnetimePaymentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: OnetimePaymentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OnetimePaymentRequest;
  static deserializeBinaryFromReader(message: OnetimePaymentRequest, reader: jspb.BinaryReader): OnetimePaymentRequest;
}

export namespace OnetimePaymentRequest {
  export type AsObject = {
    merchantUid: string,
    amount: number,
    taxFree: number,
    cardNumber: string,
    expiry: string,
    birth: string,
    pwd2digit: string,
    customerUid: string,
    pg: string,
    name: string,
    buyerName: string,
    buyerEmail: string,
    buyerTel: string,
    buyerAddr: string,
    buyerPostcode: string,
    cardQuota: number,
    interestFreeByMerchant: boolean,
    customData: string,
    noticeUrl: string,
  }
}

export class OnetimePaymentResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): v1_payment_payment_pb.Payment | undefined;
  setResponse(value?: v1_payment_payment_pb.Payment): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OnetimePaymentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: OnetimePaymentResponse): OnetimePaymentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: OnetimePaymentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OnetimePaymentResponse;
  static deserializeBinaryFromReader(message: OnetimePaymentResponse, reader: jspb.BinaryReader): OnetimePaymentResponse;
}

export namespace OnetimePaymentResponse {
  export type AsObject = {
    code: number,
    message: string,
    response?: v1_payment_payment_pb.Payment.AsObject,
  }
}

export class AgainPaymentRequest extends jspb.Message {
  getCustomerUid(): string;
  setCustomerUid(value: string): void;

  getMerchantUid(): string;
  setMerchantUid(value: string): void;

  getAmount(): number;
  setAmount(value: number): void;

  getTaxFree(): number;
  setTaxFree(value: number): void;

  getName(): string;
  setName(value: string): void;

  getBuyerName(): string;
  setBuyerName(value: string): void;

  getBuyerEmail(): string;
  setBuyerEmail(value: string): void;

  getBuyerTel(): string;
  setBuyerTel(value: string): void;

  getBuyerAddr(): string;
  setBuyerAddr(value: string): void;

  getBuyerPostcode(): string;
  setBuyerPostcode(value: string): void;

  getCardQuota(): number;
  setCardQuota(value: number): void;

  getInterestFreeByMerchant(): boolean;
  setInterestFreeByMerchant(value: boolean): void;

  getCustomData(): string;
  setCustomData(value: string): void;

  getNoticeUrl(): string;
  setNoticeUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AgainPaymentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AgainPaymentRequest): AgainPaymentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AgainPaymentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AgainPaymentRequest;
  static deserializeBinaryFromReader(message: AgainPaymentRequest, reader: jspb.BinaryReader): AgainPaymentRequest;
}

export namespace AgainPaymentRequest {
  export type AsObject = {
    customerUid: string,
    merchantUid: string,
    amount: number,
    taxFree: number,
    name: string,
    buyerName: string,
    buyerEmail: string,
    buyerTel: string,
    buyerAddr: string,
    buyerPostcode: string,
    cardQuota: number,
    interestFreeByMerchant: boolean,
    customData: string,
    noticeUrl: string,
  }
}

export class AgainPaymentResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): v1_payment_payment_pb.Payment | undefined;
  setResponse(value?: v1_payment_payment_pb.Payment): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AgainPaymentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AgainPaymentResponse): AgainPaymentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AgainPaymentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AgainPaymentResponse;
  static deserializeBinaryFromReader(message: AgainPaymentResponse, reader: jspb.BinaryReader): AgainPaymentResponse;
}

export namespace AgainPaymentResponse {
  export type AsObject = {
    code: number,
    message: string,
    response?: v1_payment_payment_pb.Payment.AsObject,
  }
}

export class PaymentScheduleParam extends jspb.Message {
  getMerchantUid(): string;
  setMerchantUid(value: string): void;

  getScheduleAt(): number;
  setScheduleAt(value: number): void;

  getAmount(): number;
  setAmount(value: number): void;

  getTaxFree(): number;
  setTaxFree(value: number): void;

  getName(): string;
  setName(value: string): void;

  getBuyerName(): string;
  setBuyerName(value: string): void;

  getBuyerEmail(): string;
  setBuyerEmail(value: string): void;

  getBuyerTel(): string;
  setBuyerTel(value: string): void;

  getBuyerAddr(): string;
  setBuyerAddr(value: string): void;

  getBuyerPostcode(): string;
  setBuyerPostcode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PaymentScheduleParam.AsObject;
  static toObject(includeInstance: boolean, msg: PaymentScheduleParam): PaymentScheduleParam.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PaymentScheduleParam, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PaymentScheduleParam;
  static deserializeBinaryFromReader(message: PaymentScheduleParam, reader: jspb.BinaryReader): PaymentScheduleParam;
}

export namespace PaymentScheduleParam {
  export type AsObject = {
    merchantUid: string,
    scheduleAt: number,
    amount: number,
    taxFree: number,
    name: string,
    buyerName: string,
    buyerEmail: string,
    buyerTel: string,
    buyerAddr: string,
    buyerPostcode: string,
  }
}

export class UnitSchedulePaymentResponse extends jspb.Message {
  getCustomerUid(): string;
  setCustomerUid(value: string): void;

  getMerchantUid(): string;
  setMerchantUid(value: string): void;

  getImpUid(): string;
  setImpUid(value: string): void;

  getScheduleAt(): number;
  setScheduleAt(value: number): void;

  getExecutedAt(): number;
  setExecutedAt(value: number): void;

  getRevokedAt(): number;
  setRevokedAt(value: number): void;

  getAmount(): number;
  setAmount(value: number): void;

  getName(): string;
  setName(value: string): void;

  getBuyerName(): string;
  setBuyerName(value: string): void;

  getBuyerEmail(): string;
  setBuyerEmail(value: string): void;

  getBuyerTel(): string;
  setBuyerTel(value: string): void;

  getBuyerAddr(): string;
  setBuyerAddr(value: string): void;

  getBuyerPostcode(): string;
  setBuyerPostcode(value: string): void;

  getCustomData(): string;
  setCustomData(value: string): void;

  getScheduleStatus(): string;
  setScheduleStatus(value: string): void;

  getPaymentStatus(): string;
  setPaymentStatus(value: string): void;

  getFailReason(): string;
  setFailReason(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnitSchedulePaymentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UnitSchedulePaymentResponse): UnitSchedulePaymentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnitSchedulePaymentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnitSchedulePaymentResponse;
  static deserializeBinaryFromReader(message: UnitSchedulePaymentResponse, reader: jspb.BinaryReader): UnitSchedulePaymentResponse;
}

export namespace UnitSchedulePaymentResponse {
  export type AsObject = {
    customerUid: string,
    merchantUid: string,
    impUid: string,
    scheduleAt: number,
    executedAt: number,
    revokedAt: number,
    amount: number,
    name: string,
    buyerName: string,
    buyerEmail: string,
    buyerTel: string,
    buyerAddr: string,
    buyerPostcode: string,
    customData: string,
    scheduleStatus: string,
    paymentStatus: string,
    failReason: string,
  }
}

export class SchedulePayemntRequest extends jspb.Message {
  getCustomerUid(): string;
  setCustomerUid(value: string): void;

  getCheckingAmount(): number;
  setCheckingAmount(value: number): void;

  getCardNumber(): string;
  setCardNumber(value: string): void;

  getExpiry(): string;
  setExpiry(value: string): void;

  getBirth(): string;
  setBirth(value: string): void;

  getPwd2digit(): string;
  setPwd2digit(value: string): void;

  getPg(): string;
  setPg(value: string): void;

  clearSchedulesList(): void;
  getSchedulesList(): Array<PaymentScheduleParam>;
  setSchedulesList(value: Array<PaymentScheduleParam>): void;
  addSchedules(value?: PaymentScheduleParam, index?: number): PaymentScheduleParam;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SchedulePayemntRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SchedulePayemntRequest): SchedulePayemntRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SchedulePayemntRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SchedulePayemntRequest;
  static deserializeBinaryFromReader(message: SchedulePayemntRequest, reader: jspb.BinaryReader): SchedulePayemntRequest;
}

export namespace SchedulePayemntRequest {
  export type AsObject = {
    customerUid: string,
    checkingAmount: number,
    cardNumber: string,
    expiry: string,
    birth: string,
    pwd2digit: string,
    pg: string,
    schedulesList: Array<PaymentScheduleParam.AsObject>,
  }
}

export class SchedulePaymentResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  clearResponseList(): void;
  getResponseList(): Array<UnitSchedulePaymentResponse>;
  setResponseList(value: Array<UnitSchedulePaymentResponse>): void;
  addResponse(value?: UnitSchedulePaymentResponse, index?: number): UnitSchedulePaymentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SchedulePaymentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SchedulePaymentResponse): SchedulePaymentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SchedulePaymentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SchedulePaymentResponse;
  static deserializeBinaryFromReader(message: SchedulePaymentResponse, reader: jspb.BinaryReader): SchedulePaymentResponse;
}

export namespace SchedulePaymentResponse {
  export type AsObject = {
    code: number,
    message: string,
    responseList: Array<UnitSchedulePaymentResponse.AsObject>,
  }
}

export class UnschedulePaymentRequest extends jspb.Message {
  getCustomerUid(): string;
  setCustomerUid(value: string): void;

  clearMerchantUidList(): void;
  getMerchantUidList(): Array<string>;
  setMerchantUidList(value: Array<string>): void;
  addMerchantUid(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnschedulePaymentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnschedulePaymentRequest): UnschedulePaymentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnschedulePaymentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnschedulePaymentRequest;
  static deserializeBinaryFromReader(message: UnschedulePaymentRequest, reader: jspb.BinaryReader): UnschedulePaymentRequest;
}

export namespace UnschedulePaymentRequest {
  export type AsObject = {
    customerUid: string,
    merchantUidList: Array<string>,
  }
}

export class UnschedulePaymentResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  clearResponseList(): void;
  getResponseList(): Array<UnitSchedulePaymentResponse>;
  setResponseList(value: Array<UnitSchedulePaymentResponse>): void;
  addResponse(value?: UnitSchedulePaymentResponse, index?: number): UnitSchedulePaymentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnschedulePaymentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UnschedulePaymentResponse): UnschedulePaymentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnschedulePaymentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnschedulePaymentResponse;
  static deserializeBinaryFromReader(message: UnschedulePaymentResponse, reader: jspb.BinaryReader): UnschedulePaymentResponse;
}

export namespace UnschedulePaymentResponse {
  export type AsObject = {
    code: number,
    message: string,
    responseList: Array<UnitSchedulePaymentResponse.AsObject>,
  }
}

export class GetPaymentScheduleRequest extends jspb.Message {
  getMerchantUid(): string;
  setMerchantUid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPaymentScheduleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPaymentScheduleRequest): GetPaymentScheduleRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPaymentScheduleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPaymentScheduleRequest;
  static deserializeBinaryFromReader(message: GetPaymentScheduleRequest, reader: jspb.BinaryReader): GetPaymentScheduleRequest;
}

export namespace GetPaymentScheduleRequest {
  export type AsObject = {
    merchantUid: string,
  }
}

export class GetPaymentScheduleResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): UnitSchedulePaymentResponse | undefined;
  setResponse(value?: UnitSchedulePaymentResponse): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPaymentScheduleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetPaymentScheduleResponse): GetPaymentScheduleResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPaymentScheduleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPaymentScheduleResponse;
  static deserializeBinaryFromReader(message: GetPaymentScheduleResponse, reader: jspb.BinaryReader): GetPaymentScheduleResponse;
}

export namespace GetPaymentScheduleResponse {
  export type AsObject = {
    code: number,
    message: string,
    response?: UnitSchedulePaymentResponse.AsObject,
  }
}

export class GetPaymentScheduleByCustomerRequest extends jspb.Message {
  getCustomerUid(): string;
  setCustomerUid(value: string): void;

  getPage(): number;
  setPage(value: number): void;

  getFrom(): number;
  setFrom(value: number): void;

  getTo(): number;
  setTo(value: number): void;

  getScheduleStatus(): string;
  setScheduleStatus(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPaymentScheduleByCustomerRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPaymentScheduleByCustomerRequest): GetPaymentScheduleByCustomerRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPaymentScheduleByCustomerRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPaymentScheduleByCustomerRequest;
  static deserializeBinaryFromReader(message: GetPaymentScheduleByCustomerRequest, reader: jspb.BinaryReader): GetPaymentScheduleByCustomerRequest;
}

export namespace GetPaymentScheduleByCustomerRequest {
  export type AsObject = {
    customerUid: string,
    page: number,
    from: number,
    to: number,
    scheduleStatus: string,
  }
}

export class NestedGetPaymentScheduleByCustomerData extends jspb.Message {
  getTotal(): number;
  setTotal(value: number): void;

  getPrevious(): number;
  setPrevious(value: number): void;

  getNext(): number;
  setNext(value: number): void;

  clearListList(): void;
  getListList(): Array<UnitSchedulePaymentResponse>;
  setListList(value: Array<UnitSchedulePaymentResponse>): void;
  addList(value?: UnitSchedulePaymentResponse, index?: number): UnitSchedulePaymentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NestedGetPaymentScheduleByCustomerData.AsObject;
  static toObject(includeInstance: boolean, msg: NestedGetPaymentScheduleByCustomerData): NestedGetPaymentScheduleByCustomerData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NestedGetPaymentScheduleByCustomerData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NestedGetPaymentScheduleByCustomerData;
  static deserializeBinaryFromReader(message: NestedGetPaymentScheduleByCustomerData, reader: jspb.BinaryReader): NestedGetPaymentScheduleByCustomerData;
}

export namespace NestedGetPaymentScheduleByCustomerData {
  export type AsObject = {
    total: number,
    previous: number,
    next: number,
    listList: Array<UnitSchedulePaymentResponse.AsObject>,
  }
}

export class GetPaymentScheduleByCustomerResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): NestedGetPaymentScheduleByCustomerData | undefined;
  setResponse(value?: NestedGetPaymentScheduleByCustomerData): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPaymentScheduleByCustomerResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetPaymentScheduleByCustomerResponse): GetPaymentScheduleByCustomerResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPaymentScheduleByCustomerResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPaymentScheduleByCustomerResponse;
  static deserializeBinaryFromReader(message: GetPaymentScheduleByCustomerResponse, reader: jspb.BinaryReader): GetPaymentScheduleByCustomerResponse;
}

export namespace GetPaymentScheduleByCustomerResponse {
  export type AsObject = {
    code: number,
    message: string,
    response?: NestedGetPaymentScheduleByCustomerData.AsObject,
  }
}

