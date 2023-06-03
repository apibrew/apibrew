// @generated by protoc-gen-es v1.2.1 with parameter "target=ts"
// @generated from file stub/watch.proto (package stub, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { EventSelector } from "../model/event_pb";

/**
 * @generated from message stub.WatchRequest
 */
export class WatchRequest extends Message<WatchRequest> {
  /**
   * @generated from field: string token = 1;
   */
  token = "";

  /**
   * @generated from field: model.EventSelector selector = 2;
   */
  selector?: EventSelector;

  constructor(data?: PartialMessage<WatchRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stub.WatchRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "selector", kind: "message", T: EventSelector },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WatchRequest {
    return new WatchRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WatchRequest {
    return new WatchRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WatchRequest {
    return new WatchRequest().fromJsonString(jsonString, options);
  }

  static equals(a: WatchRequest | PlainMessage<WatchRequest> | undefined, b: WatchRequest | PlainMessage<WatchRequest> | undefined): boolean {
    return proto3.util.equals(WatchRequest, a, b);
  }
}

