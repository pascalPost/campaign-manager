// @generated by protoc-gen-es v1.8.0 with parameter "target=ts"
// @generated from file proto/cm/v1/cm.proto (package proto.cm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message proto.cm.v1.PingRequest
 */
export class PingRequest extends Message<PingRequest> {
  constructor(data?: PartialMessage<PingRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.cm.v1.PingRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PingRequest {
    return new PingRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PingRequest {
    return new PingRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PingRequest {
    return new PingRequest().fromJsonString(jsonString, options);
  }

  static equals(a: PingRequest | PlainMessage<PingRequest> | undefined, b: PingRequest | PlainMessage<PingRequest> | undefined): boolean {
    return proto3.util.equals(PingRequest, a, b);
  }
}

/**
 * @generated from message proto.cm.v1.PingResponse
 */
export class PingResponse extends Message<PingResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  constructor(data?: PartialMessage<PingResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.cm.v1.PingResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PingResponse {
    return new PingResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PingResponse {
    return new PingResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PingResponse {
    return new PingResponse().fromJsonString(jsonString, options);
  }

  static equals(a: PingResponse | PlainMessage<PingResponse> | undefined, b: PingResponse | PlainMessage<PingResponse> | undefined): boolean {
    return proto3.util.equals(PingResponse, a, b);
  }
}

/**
 * @generated from message proto.cm.v1.NewProjectRequest
 */
export class NewProjectRequest extends Message<NewProjectRequest> {
  /**
   * @generated from field: string project_name = 1;
   */
  projectName = "";

  /**
   * @generated from field: string csv_file_path = 2;
   */
  csvFilePath = "";

  constructor(data?: PartialMessage<NewProjectRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.cm.v1.NewProjectRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "csv_file_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NewProjectRequest {
    return new NewProjectRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NewProjectRequest {
    return new NewProjectRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NewProjectRequest {
    return new NewProjectRequest().fromJsonString(jsonString, options);
  }

  static equals(a: NewProjectRequest | PlainMessage<NewProjectRequest> | undefined, b: NewProjectRequest | PlainMessage<NewProjectRequest> | undefined): boolean {
    return proto3.util.equals(NewProjectRequest, a, b);
  }
}

/**
 * @generated from message proto.cm.v1.NewProjectResponse
 */
export class NewProjectResponse extends Message<NewProjectResponse> {
  /**
   * @generated from field: optional uint64 project_id = 1;
   */
  projectId?: bigint;

  constructor(data?: PartialMessage<NewProjectResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.cm.v1.NewProjectResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NewProjectResponse {
    return new NewProjectResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NewProjectResponse {
    return new NewProjectResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NewProjectResponse {
    return new NewProjectResponse().fromJsonString(jsonString, options);
  }

  static equals(a: NewProjectResponse | PlainMessage<NewProjectResponse> | undefined, b: NewProjectResponse | PlainMessage<NewProjectResponse> | undefined): boolean {
    return proto3.util.equals(NewProjectResponse, a, b);
  }
}

/**
 * @generated from message proto.cm.v1.GetSettingsRequest
 */
export class GetSettingsRequest extends Message<GetSettingsRequest> {
  constructor(data?: PartialMessage<GetSettingsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.cm.v1.GetSettingsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSettingsRequest {
    return new GetSettingsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSettingsRequest {
    return new GetSettingsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSettingsRequest {
    return new GetSettingsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetSettingsRequest | PlainMessage<GetSettingsRequest> | undefined, b: GetSettingsRequest | PlainMessage<GetSettingsRequest> | undefined): boolean {
    return proto3.util.equals(GetSettingsRequest, a, b);
  }
}

/**
 * @generated from message proto.cm.v1.GetSettingsResponse
 */
export class GetSettingsResponse extends Message<GetSettingsResponse> {
  /**
   * @generated from field: string working_dir = 1;
   */
  workingDir = "";

  /**
   * @generated from field: string lsf_username = 2;
   */
  lsfUsername = "";

  /**
   * @generated from field: string lsf_password = 3;
   */
  lsfPassword = "";

  constructor(data?: PartialMessage<GetSettingsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.cm.v1.GetSettingsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "working_dir", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "lsf_username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "lsf_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSettingsResponse {
    return new GetSettingsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSettingsResponse {
    return new GetSettingsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSettingsResponse {
    return new GetSettingsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetSettingsResponse | PlainMessage<GetSettingsResponse> | undefined, b: GetSettingsResponse | PlainMessage<GetSettingsResponse> | undefined): boolean {
    return proto3.util.equals(GetSettingsResponse, a, b);
  }
}

/**
 * @generated from message proto.cm.v1.SetSettingsRequest
 */
export class SetSettingsRequest extends Message<SetSettingsRequest> {
  /**
   * @generated from field: string working_dir = 1;
   */
  workingDir = "";

  /**
   * @generated from field: string lsf_username = 2;
   */
  lsfUsername = "";

  /**
   * @generated from field: string lsf_password = 3;
   */
  lsfPassword = "";

  constructor(data?: PartialMessage<SetSettingsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.cm.v1.SetSettingsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "working_dir", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "lsf_username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "lsf_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetSettingsRequest {
    return new SetSettingsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetSettingsRequest {
    return new SetSettingsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetSettingsRequest {
    return new SetSettingsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SetSettingsRequest | PlainMessage<SetSettingsRequest> | undefined, b: SetSettingsRequest | PlainMessage<SetSettingsRequest> | undefined): boolean {
    return proto3.util.equals(SetSettingsRequest, a, b);
  }
}

/**
 * @generated from message proto.cm.v1.SetSettingsResponse
 */
export class SetSettingsResponse extends Message<SetSettingsResponse> {
  /**
   * @generated from field: string working_dir = 1;
   */
  workingDir = "";

  /**
   * @generated from field: string lsf_username = 2;
   */
  lsfUsername = "";

  /**
   * @generated from field: string lsf_password = 3;
   */
  lsfPassword = "";

  constructor(data?: PartialMessage<SetSettingsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.cm.v1.SetSettingsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "working_dir", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "lsf_username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "lsf_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetSettingsResponse {
    return new SetSettingsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetSettingsResponse {
    return new SetSettingsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetSettingsResponse {
    return new SetSettingsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SetSettingsResponse | PlainMessage<SetSettingsResponse> | undefined, b: SetSettingsResponse | PlainMessage<SetSettingsResponse> | undefined): boolean {
    return proto3.util.equals(SetSettingsResponse, a, b);
  }
}

