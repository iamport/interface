// package: certification
// file: v1/certification/certification.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../../google/api/annotations_pb";

export class CertificationRequest extends jspb.Message {
  getImpUid(): string;
  setImpUid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CertificationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CertificationRequest): CertificationRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CertificationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CertificationRequest;
  static deserializeBinaryFromReader(message: CertificationRequest, reader: jspb.BinaryReader): CertificationRequest;
}

export namespace CertificationRequest {
  export type AsObject = {
    impUid: string,
  }
}

export class CertificationResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): Certification | undefined;
  setResponse(value?: Certification): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CertificationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CertificationResponse): CertificationResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CertificationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CertificationResponse;
  static deserializeBinaryFromReader(message: CertificationResponse, reader: jspb.BinaryReader): CertificationResponse;
}

export namespace CertificationResponse {
  export type AsObject = {
    code: number,
    message: string,
    response?: Certification.AsObject,
  }
}

export class Certification extends jspb.Message {
  getImpUid(): string;
  setImpUid(value: string): void;

  getMerchantUid(): string;
  setMerchantUid(value: string): void;

  getPgTid(): string;
  setPgTid(value: string): void;

  getPgProvider(): string;
  setPgProvider(value: string): void;

  getName(): string;
  setName(value: string): void;

  getGender(): string;
  setGender(value: string): void;

  getBirth(): number;
  setBirth(value: number): void;

  getBirthday(): string;
  setBirthday(value: string): void;

  getForeigner(): boolean;
  setForeigner(value: boolean): void;

  getPhone(): string;
  setPhone(value: string): void;

  getCarrier(): string;
  setCarrier(value: string): void;

  getCertified(): boolean;
  setCertified(value: boolean): void;

  getCertifiedAt(): number;
  setCertifiedAt(value: number): void;

  getUniqueKey(): string;
  setUniqueKey(value: string): void;

  getUniqueInSite(): string;
  setUniqueInSite(value: string): void;

  getOrigin(): string;
  setOrigin(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Certification.AsObject;
  static toObject(includeInstance: boolean, msg: Certification): Certification.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Certification, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Certification;
  static deserializeBinaryFromReader(message: Certification, reader: jspb.BinaryReader): Certification;
}

export namespace Certification {
  export type AsObject = {
    impUid: string,
    merchantUid: string,
    pgTid: string,
    pgProvider: string,
    name: string,
    gender: string,
    birth: number,
    birthday: string,
    foreigner: boolean,
    phone: string,
    carrier: string,
    certified: boolean,
    certifiedAt: number,
    uniqueKey: string,
    uniqueInSite: string,
    origin: string,
  }
}

