// @generated by protoc-gen-connect-es v0.9.1 with parameter "target=ts,import_extension="
// @generated from file stub/record.proto (package stub, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ApplyRecordRequest, ApplyRecordResponse, CreateRecordRequest, CreateRecordResponse, DeleteRecordRequest, DeleteRecordResponse, GetRecordRequest, GetRecordResponse, ListRecordRequest, ListRecordResponse, ReadStreamRequest, SearchRecordRequest, SearchRecordResponse, UpdateMultiRecordRequest, UpdateMultiRecordResponse, UpdateRecordRequest, UpdateRecordResponse, WriteStreamResponse } from "./record_pb";
import { MethodKind } from "@bufbuild/protobuf";
import { Record } from "../model/record_pb";

/**
 * Record service is an abstract service for records of all resources. You can do CRUD like operations with Record service
 *
 * @generated from service stub.Record
 */
export const Record = {
  typeName: "stub.Record",
  methods: {
    /**
     * @generated from rpc stub.Record.Create
     */
    create: {
      name: "Create",
      I: CreateRecordRequest,
      O: CreateRecordResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Record.Update
     */
    update: {
      name: "Update",
      I: UpdateRecordRequest,
      O: UpdateRecordResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Record.Apply
     */
    apply: {
      name: "Apply",
      I: ApplyRecordRequest,
      O: ApplyRecordResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * Not implemented yet
     *
     * @generated from rpc stub.Record.UpdateMulti
     */
    updateMulti: {
      name: "UpdateMulti",
      I: UpdateMultiRecordRequest,
      O: UpdateMultiRecordResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Record.Delete
     */
    delete: {
      name: "Delete",
      I: DeleteRecordRequest,
      O: DeleteRecordResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Record.List
     */
    list: {
      name: "List",
      I: ListRecordRequest,
      O: ListRecordResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Record.Search
     */
    search: {
      name: "Search",
      I: SearchRecordRequest,
      O: SearchRecordResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Record.ReadStream
     */
    readStream: {
      name: "ReadStream",
      I: ReadStreamRequest,
      O: Record,
      kind: MethodKind.ServerStreaming,
    },
    /**
     *
     * Not implemented yet
     *
     * @generated from rpc stub.Record.WriteStream
     */
    writeStream: {
      name: "WriteStream",
      I: Record,
      O: WriteStreamResponse,
      kind: MethodKind.ClientStreaming,
    },
    /**
     * @generated from rpc stub.Record.Get
     */
    get: {
      name: "Get",
      I: GetRecordRequest,
      O: GetRecordResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
