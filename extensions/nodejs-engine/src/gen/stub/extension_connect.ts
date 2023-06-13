// @generated by protoc-gen-connect-es v0.9.1 with parameter "target=ts,import_extension="
// @generated from file stub/extension.proto (package stub, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateExtensionRequest, CreateExtensionResponse, DeleteExtensionRequest, DeleteExtensionResponse, GetExtensionRequest, GetExtensionResponse, ListExtensionRequest, ListExtensionResponse, UpdateExtensionRequest, UpdateExtensionResponse } from "./extension_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * Extension Service is for managing extensions
 *
 * @generated from service stub.Extension
 */
export const Extension = {
  typeName: "stub.Extension",
  methods: {
    /**
     * @generated from rpc stub.Extension.List
     */
    list: {
      name: "List",
      I: ListExtensionRequest,
      O: ListExtensionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Extension.Get
     */
    get: {
      name: "Get",
      I: GetExtensionRequest,
      O: GetExtensionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Extension.Create
     */
    create: {
      name: "Create",
      I: CreateExtensionRequest,
      O: CreateExtensionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Extension.Update
     */
    update: {
      name: "Update",
      I: UpdateExtensionRequest,
      O: UpdateExtensionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stub.Extension.Delete
     */
    delete: {
      name: "Delete",
      I: DeleteExtensionRequest,
      O: DeleteExtensionResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
