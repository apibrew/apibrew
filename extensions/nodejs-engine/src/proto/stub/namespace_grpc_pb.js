// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_namespace_pb = require('../stub/namespace_pb.js');
var model_namespace_pb = require('../model/namespace_pb.js');
var model_error_pb = require('../model/error_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var gnostic_openapi_v3_annotations_pb = require('../gnostic/openapi/v3/annotations_pb.js');

function serialize_stub_CreateNamespaceRequest(arg) {
  if (!(arg instanceof stub_namespace_pb.CreateNamespaceRequest)) {
    throw new Error('Expected argument of type stub.CreateNamespaceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateNamespaceRequest(buffer_arg) {
  return stub_namespace_pb.CreateNamespaceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_CreateNamespaceResponse(arg) {
  if (!(arg instanceof stub_namespace_pb.CreateNamespaceResponse)) {
    throw new Error('Expected argument of type stub.CreateNamespaceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateNamespaceResponse(buffer_arg) {
  return stub_namespace_pb.CreateNamespaceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteNamespaceRequest(arg) {
  if (!(arg instanceof stub_namespace_pb.DeleteNamespaceRequest)) {
    throw new Error('Expected argument of type stub.DeleteNamespaceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteNamespaceRequest(buffer_arg) {
  return stub_namespace_pb.DeleteNamespaceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteNamespaceResponse(arg) {
  if (!(arg instanceof stub_namespace_pb.DeleteNamespaceResponse)) {
    throw new Error('Expected argument of type stub.DeleteNamespaceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteNamespaceResponse(buffer_arg) {
  return stub_namespace_pb.DeleteNamespaceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetNamespaceRequest(arg) {
  if (!(arg instanceof stub_namespace_pb.GetNamespaceRequest)) {
    throw new Error('Expected argument of type stub.GetNamespaceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetNamespaceRequest(buffer_arg) {
  return stub_namespace_pb.GetNamespaceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetNamespaceResponse(arg) {
  if (!(arg instanceof stub_namespace_pb.GetNamespaceResponse)) {
    throw new Error('Expected argument of type stub.GetNamespaceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetNamespaceResponse(buffer_arg) {
  return stub_namespace_pb.GetNamespaceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListNamespaceRequest(arg) {
  if (!(arg instanceof stub_namespace_pb.ListNamespaceRequest)) {
    throw new Error('Expected argument of type stub.ListNamespaceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListNamespaceRequest(buffer_arg) {
  return stub_namespace_pb.ListNamespaceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListNamespaceResponse(arg) {
  if (!(arg instanceof stub_namespace_pb.ListNamespaceResponse)) {
    throw new Error('Expected argument of type stub.ListNamespaceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListNamespaceResponse(buffer_arg) {
  return stub_namespace_pb.ListNamespaceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateNamespaceRequest(arg) {
  if (!(arg instanceof stub_namespace_pb.UpdateNamespaceRequest)) {
    throw new Error('Expected argument of type stub.UpdateNamespaceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateNamespaceRequest(buffer_arg) {
  return stub_namespace_pb.UpdateNamespaceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateNamespaceResponse(arg) {
  if (!(arg instanceof stub_namespace_pb.UpdateNamespaceResponse)) {
    throw new Error('Expected argument of type stub.UpdateNamespaceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateNamespaceResponse(buffer_arg) {
  return stub_namespace_pb.UpdateNamespaceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Namespace Service is for managing namespaces
var NamespaceService = exports.NamespaceService = {
  create: {
    path: '/stub.Namespace/Create',
    requestStream: false,
    responseStream: false,
    requestType: stub_namespace_pb.CreateNamespaceRequest,
    responseType: stub_namespace_pb.CreateNamespaceResponse,
    requestSerialize: serialize_stub_CreateNamespaceRequest,
    requestDeserialize: deserialize_stub_CreateNamespaceRequest,
    responseSerialize: serialize_stub_CreateNamespaceResponse,
    responseDeserialize: deserialize_stub_CreateNamespaceResponse,
  },
  list: {
    path: '/stub.Namespace/List',
    requestStream: false,
    responseStream: false,
    requestType: stub_namespace_pb.ListNamespaceRequest,
    responseType: stub_namespace_pb.ListNamespaceResponse,
    requestSerialize: serialize_stub_ListNamespaceRequest,
    requestDeserialize: deserialize_stub_ListNamespaceRequest,
    responseSerialize: serialize_stub_ListNamespaceResponse,
    responseDeserialize: deserialize_stub_ListNamespaceResponse,
  },
  update: {
    path: '/stub.Namespace/Update',
    requestStream: false,
    responseStream: false,
    requestType: stub_namespace_pb.UpdateNamespaceRequest,
    responseType: stub_namespace_pb.UpdateNamespaceResponse,
    requestSerialize: serialize_stub_UpdateNamespaceRequest,
    requestDeserialize: deserialize_stub_UpdateNamespaceRequest,
    responseSerialize: serialize_stub_UpdateNamespaceResponse,
    responseDeserialize: deserialize_stub_UpdateNamespaceResponse,
  },
  delete: {
    path: '/stub.Namespace/Delete',
    requestStream: false,
    responseStream: false,
    requestType: stub_namespace_pb.DeleteNamespaceRequest,
    responseType: stub_namespace_pb.DeleteNamespaceResponse,
    requestSerialize: serialize_stub_DeleteNamespaceRequest,
    requestDeserialize: deserialize_stub_DeleteNamespaceRequest,
    responseSerialize: serialize_stub_DeleteNamespaceResponse,
    responseDeserialize: deserialize_stub_DeleteNamespaceResponse,
  },
  get: {
    path: '/stub.Namespace/Get',
    requestStream: false,
    responseStream: false,
    requestType: stub_namespace_pb.GetNamespaceRequest,
    responseType: stub_namespace_pb.GetNamespaceResponse,
    requestSerialize: serialize_stub_GetNamespaceRequest,
    requestDeserialize: deserialize_stub_GetNamespaceRequest,
    responseSerialize: serialize_stub_GetNamespaceResponse,
    responseDeserialize: deserialize_stub_GetNamespaceResponse,
  },
};

exports.NamespaceClient = grpc.makeGenericClientConstructor(NamespaceService);
