// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_generic_pb = require('../stub/generic_pb.js');
var model_query_pb = require('../model/query_pb.js');
var model_error_pb = require('../model/error_pb.js');
var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');
var google_protobuf_any_pb = require('google-protobuf/google/protobuf/any_pb.js');

function serialize_stub_CreateRequest(arg) {
  if (!(arg instanceof stub_generic_pb.CreateRequest)) {
    throw new Error('Expected argument of type stub.CreateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateRequest(buffer_arg) {
  return stub_generic_pb.CreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_CreateResponse(arg) {
  if (!(arg instanceof stub_generic_pb.CreateResponse)) {
    throw new Error('Expected argument of type stub.CreateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateResponse(buffer_arg) {
  return stub_generic_pb.CreateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteRequest(arg) {
  if (!(arg instanceof stub_generic_pb.DeleteRequest)) {
    throw new Error('Expected argument of type stub.DeleteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteRequest(buffer_arg) {
  return stub_generic_pb.DeleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteResponse(arg) {
  if (!(arg instanceof stub_generic_pb.DeleteResponse)) {
    throw new Error('Expected argument of type stub.DeleteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteResponse(buffer_arg) {
  return stub_generic_pb.DeleteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetRequest(arg) {
  if (!(arg instanceof stub_generic_pb.GetRequest)) {
    throw new Error('Expected argument of type stub.GetRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetRequest(buffer_arg) {
  return stub_generic_pb.GetRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetResponse(arg) {
  if (!(arg instanceof stub_generic_pb.GetResponse)) {
    throw new Error('Expected argument of type stub.GetResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetResponse(buffer_arg) {
  return stub_generic_pb.GetResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListRequest(arg) {
  if (!(arg instanceof stub_generic_pb.ListRequest)) {
    throw new Error('Expected argument of type stub.ListRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListRequest(buffer_arg) {
  return stub_generic_pb.ListRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListResponse(arg) {
  if (!(arg instanceof stub_generic_pb.ListResponse)) {
    throw new Error('Expected argument of type stub.ListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListResponse(buffer_arg) {
  return stub_generic_pb.ListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_SearchRequest(arg) {
  if (!(arg instanceof stub_generic_pb.SearchRequest)) {
    throw new Error('Expected argument of type stub.SearchRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_SearchRequest(buffer_arg) {
  return stub_generic_pb.SearchRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_SearchResponse(arg) {
  if (!(arg instanceof stub_generic_pb.SearchResponse)) {
    throw new Error('Expected argument of type stub.SearchResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_SearchResponse(buffer_arg) {
  return stub_generic_pb.SearchResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateMultiRequest(arg) {
  if (!(arg instanceof stub_generic_pb.UpdateMultiRequest)) {
    throw new Error('Expected argument of type stub.UpdateMultiRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateMultiRequest(buffer_arg) {
  return stub_generic_pb.UpdateMultiRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateMultiResponse(arg) {
  if (!(arg instanceof stub_generic_pb.UpdateMultiResponse)) {
    throw new Error('Expected argument of type stub.UpdateMultiResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateMultiResponse(buffer_arg) {
  return stub_generic_pb.UpdateMultiResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateRequest(arg) {
  if (!(arg instanceof stub_generic_pb.UpdateRequest)) {
    throw new Error('Expected argument of type stub.UpdateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateRequest(buffer_arg) {
  return stub_generic_pb.UpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateResponse(arg) {
  if (!(arg instanceof stub_generic_pb.UpdateResponse)) {
    throw new Error('Expected argument of type stub.UpdateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateResponse(buffer_arg) {
  return stub_generic_pb.UpdateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


//
// Not implemented yet
var GenericService = exports.GenericService = {
  create: {
    path: '/stub.Generic/Create',
    requestStream: false,
    responseStream: false,
    requestType: stub_generic_pb.CreateRequest,
    responseType: stub_generic_pb.CreateResponse,
    requestSerialize: serialize_stub_CreateRequest,
    requestDeserialize: deserialize_stub_CreateRequest,
    responseSerialize: serialize_stub_CreateResponse,
    responseDeserialize: deserialize_stub_CreateResponse,
  },
  update: {
    path: '/stub.Generic/Update',
    requestStream: false,
    responseStream: false,
    requestType: stub_generic_pb.UpdateRequest,
    responseType: stub_generic_pb.UpdateResponse,
    requestSerialize: serialize_stub_UpdateRequest,
    requestDeserialize: deserialize_stub_UpdateRequest,
    responseSerialize: serialize_stub_UpdateResponse,
    responseDeserialize: deserialize_stub_UpdateResponse,
  },
  updateMulti: {
    path: '/stub.Generic/UpdateMulti',
    requestStream: false,
    responseStream: false,
    requestType: stub_generic_pb.UpdateMultiRequest,
    responseType: stub_generic_pb.UpdateMultiResponse,
    requestSerialize: serialize_stub_UpdateMultiRequest,
    requestDeserialize: deserialize_stub_UpdateMultiRequest,
    responseSerialize: serialize_stub_UpdateMultiResponse,
    responseDeserialize: deserialize_stub_UpdateMultiResponse,
  },
  delete: {
    path: '/stub.Generic/Delete',
    requestStream: false,
    responseStream: false,
    requestType: stub_generic_pb.DeleteRequest,
    responseType: stub_generic_pb.DeleteResponse,
    requestSerialize: serialize_stub_DeleteRequest,
    requestDeserialize: deserialize_stub_DeleteRequest,
    responseSerialize: serialize_stub_DeleteResponse,
    responseDeserialize: deserialize_stub_DeleteResponse,
  },
  list: {
    path: '/stub.Generic/List',
    requestStream: false,
    responseStream: false,
    requestType: stub_generic_pb.ListRequest,
    responseType: stub_generic_pb.ListResponse,
    requestSerialize: serialize_stub_ListRequest,
    requestDeserialize: deserialize_stub_ListRequest,
    responseSerialize: serialize_stub_ListResponse,
    responseDeserialize: deserialize_stub_ListResponse,
  },
  search: {
    path: '/stub.Generic/Search',
    requestStream: false,
    responseStream: false,
    requestType: stub_generic_pb.SearchRequest,
    responseType: stub_generic_pb.SearchResponse,
    requestSerialize: serialize_stub_SearchRequest,
    requestDeserialize: deserialize_stub_SearchRequest,
    responseSerialize: serialize_stub_SearchResponse,
    responseDeserialize: deserialize_stub_SearchResponse,
  },
  get: {
    path: '/stub.Generic/Get',
    requestStream: false,
    responseStream: false,
    requestType: stub_generic_pb.GetRequest,
    responseType: stub_generic_pb.GetResponse,
    requestSerialize: serialize_stub_GetRequest,
    requestDeserialize: deserialize_stub_GetRequest,
    responseSerialize: serialize_stub_GetResponse,
    responseDeserialize: deserialize_stub_GetResponse,
  },
};

exports.GenericClient = grpc.makeGenericClientConstructor(GenericService);
