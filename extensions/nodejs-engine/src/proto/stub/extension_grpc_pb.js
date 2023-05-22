// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_extension_pb = require('../stub/extension_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var model_query_pb = require('../model/query_pb.js');
var model_extension_pb = require('../model/extension_pb.js');

function serialize_stub_CreateExtensionRequest(arg) {
  if (!(arg instanceof stub_extension_pb.CreateExtensionRequest)) {
    throw new Error('Expected argument of type stub.CreateExtensionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateExtensionRequest(buffer_arg) {
  return stub_extension_pb.CreateExtensionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_CreateExtensionResponse(arg) {
  if (!(arg instanceof stub_extension_pb.CreateExtensionResponse)) {
    throw new Error('Expected argument of type stub.CreateExtensionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateExtensionResponse(buffer_arg) {
  return stub_extension_pb.CreateExtensionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteExtensionRequest(arg) {
  if (!(arg instanceof stub_extension_pb.DeleteExtensionRequest)) {
    throw new Error('Expected argument of type stub.DeleteExtensionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteExtensionRequest(buffer_arg) {
  return stub_extension_pb.DeleteExtensionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteExtensionResponse(arg) {
  if (!(arg instanceof stub_extension_pb.DeleteExtensionResponse)) {
    throw new Error('Expected argument of type stub.DeleteExtensionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteExtensionResponse(buffer_arg) {
  return stub_extension_pb.DeleteExtensionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetExtensionRequest(arg) {
  if (!(arg instanceof stub_extension_pb.GetExtensionRequest)) {
    throw new Error('Expected argument of type stub.GetExtensionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetExtensionRequest(buffer_arg) {
  return stub_extension_pb.GetExtensionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetExtensionResponse(arg) {
  if (!(arg instanceof stub_extension_pb.GetExtensionResponse)) {
    throw new Error('Expected argument of type stub.GetExtensionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetExtensionResponse(buffer_arg) {
  return stub_extension_pb.GetExtensionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListExtensionRequest(arg) {
  if (!(arg instanceof stub_extension_pb.ListExtensionRequest)) {
    throw new Error('Expected argument of type stub.ListExtensionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListExtensionRequest(buffer_arg) {
  return stub_extension_pb.ListExtensionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListExtensionResponse(arg) {
  if (!(arg instanceof stub_extension_pb.ListExtensionResponse)) {
    throw new Error('Expected argument of type stub.ListExtensionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListExtensionResponse(buffer_arg) {
  return stub_extension_pb.ListExtensionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateExtensionRequest(arg) {
  if (!(arg instanceof stub_extension_pb.UpdateExtensionRequest)) {
    throw new Error('Expected argument of type stub.UpdateExtensionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateExtensionRequest(buffer_arg) {
  return stub_extension_pb.UpdateExtensionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateExtensionResponse(arg) {
  if (!(arg instanceof stub_extension_pb.UpdateExtensionResponse)) {
    throw new Error('Expected argument of type stub.UpdateExtensionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateExtensionResponse(buffer_arg) {
  return stub_extension_pb.UpdateExtensionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Extension Service is for managing extensions
var ExtensionService = exports.ExtensionService = {
  list: {
    path: '/stub.Extension/List',
    requestStream: false,
    responseStream: false,
    requestType: stub_extension_pb.ListExtensionRequest,
    responseType: stub_extension_pb.ListExtensionResponse,
    requestSerialize: serialize_stub_ListExtensionRequest,
    requestDeserialize: deserialize_stub_ListExtensionRequest,
    responseSerialize: serialize_stub_ListExtensionResponse,
    responseDeserialize: deserialize_stub_ListExtensionResponse,
  },
  get: {
    path: '/stub.Extension/Get',
    requestStream: false,
    responseStream: false,
    requestType: stub_extension_pb.GetExtensionRequest,
    responseType: stub_extension_pb.GetExtensionResponse,
    requestSerialize: serialize_stub_GetExtensionRequest,
    requestDeserialize: deserialize_stub_GetExtensionRequest,
    responseSerialize: serialize_stub_GetExtensionResponse,
    responseDeserialize: deserialize_stub_GetExtensionResponse,
  },
  create: {
    path: '/stub.Extension/Create',
    requestStream: false,
    responseStream: false,
    requestType: stub_extension_pb.CreateExtensionRequest,
    responseType: stub_extension_pb.CreateExtensionResponse,
    requestSerialize: serialize_stub_CreateExtensionRequest,
    requestDeserialize: deserialize_stub_CreateExtensionRequest,
    responseSerialize: serialize_stub_CreateExtensionResponse,
    responseDeserialize: deserialize_stub_CreateExtensionResponse,
  },
  update: {
    path: '/stub.Extension/Update',
    requestStream: false,
    responseStream: false,
    requestType: stub_extension_pb.UpdateExtensionRequest,
    responseType: stub_extension_pb.UpdateExtensionResponse,
    requestSerialize: serialize_stub_UpdateExtensionRequest,
    requestDeserialize: deserialize_stub_UpdateExtensionRequest,
    responseSerialize: serialize_stub_UpdateExtensionResponse,
    responseDeserialize: deserialize_stub_UpdateExtensionResponse,
  },
  delete: {
    path: '/stub.Extension/Delete',
    requestStream: false,
    responseStream: false,
    requestType: stub_extension_pb.DeleteExtensionRequest,
    responseType: stub_extension_pb.DeleteExtensionResponse,
    requestSerialize: serialize_stub_DeleteExtensionRequest,
    requestDeserialize: deserialize_stub_DeleteExtensionRequest,
    responseSerialize: serialize_stub_DeleteExtensionResponse,
    responseDeserialize: deserialize_stub_DeleteExtensionResponse,
  },
};

exports.ExtensionClient = grpc.makeGenericClientConstructor(ExtensionService);
