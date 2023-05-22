// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_rest_record_pb = require('../../stub/rest/record_pb.js');
var model_record_pb = require('../../model/record_pb.js');
var model_query_pb = require('../../model/query_pb.js');
var model_error_pb = require('../../model/error_pb.js');
var google_api_annotations_pb = require('../../google/api/annotations_pb.js');
var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');

function serialize_rest_ApplyRecordRequest(arg) {
  if (!(arg instanceof stub_rest_record_pb.ApplyRecordRequest)) {
    throw new Error('Expected argument of type rest.ApplyRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rest_ApplyRecordRequest(buffer_arg) {
  return stub_rest_record_pb.ApplyRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rest_ApplyRecordResponse(arg) {
  if (!(arg instanceof stub_rest_record_pb.ApplyRecordResponse)) {
    throw new Error('Expected argument of type rest.ApplyRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rest_ApplyRecordResponse(buffer_arg) {
  return stub_rest_record_pb.ApplyRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rest_CreateRecordRequest(arg) {
  if (!(arg instanceof stub_rest_record_pb.CreateRecordRequest)) {
    throw new Error('Expected argument of type rest.CreateRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rest_CreateRecordRequest(buffer_arg) {
  return stub_rest_record_pb.CreateRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rest_CreateRecordResponse(arg) {
  if (!(arg instanceof stub_rest_record_pb.CreateRecordResponse)) {
    throw new Error('Expected argument of type rest.CreateRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rest_CreateRecordResponse(buffer_arg) {
  return stub_rest_record_pb.CreateRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rest_DeleteRecordRequest(arg) {
  if (!(arg instanceof stub_rest_record_pb.DeleteRecordRequest)) {
    throw new Error('Expected argument of type rest.DeleteRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rest_DeleteRecordRequest(buffer_arg) {
  return stub_rest_record_pb.DeleteRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rest_DeleteRecordResponse(arg) {
  if (!(arg instanceof stub_rest_record_pb.DeleteRecordResponse)) {
    throw new Error('Expected argument of type rest.DeleteRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rest_DeleteRecordResponse(buffer_arg) {
  return stub_rest_record_pb.DeleteRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rest_UpdateRecordRequest(arg) {
  if (!(arg instanceof stub_rest_record_pb.UpdateRecordRequest)) {
    throw new Error('Expected argument of type rest.UpdateRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rest_UpdateRecordRequest(buffer_arg) {
  return stub_rest_record_pb.UpdateRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rest_UpdateRecordResponse(arg) {
  if (!(arg instanceof stub_rest_record_pb.UpdateRecordResponse)) {
    throw new Error('Expected argument of type rest.UpdateRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rest_UpdateRecordResponse(buffer_arg) {
  return stub_rest_record_pb.UpdateRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Record service is an abstract service for records of all resources. You can do CRUD like operations with Record service
var RecordService = exports.RecordService = {
  create: {
    path: '/rest.Record/Create',
    requestStream: false,
    responseStream: false,
    requestType: stub_rest_record_pb.CreateRecordRequest,
    responseType: stub_rest_record_pb.CreateRecordResponse,
    requestSerialize: serialize_rest_CreateRecordRequest,
    requestDeserialize: deserialize_rest_CreateRecordRequest,
    responseSerialize: serialize_rest_CreateRecordResponse,
    responseDeserialize: deserialize_rest_CreateRecordResponse,
  },
  update: {
    path: '/rest.Record/Update',
    requestStream: false,
    responseStream: false,
    requestType: stub_rest_record_pb.UpdateRecordRequest,
    responseType: stub_rest_record_pb.UpdateRecordResponse,
    requestSerialize: serialize_rest_UpdateRecordRequest,
    requestDeserialize: deserialize_rest_UpdateRecordRequest,
    responseSerialize: serialize_rest_UpdateRecordResponse,
    responseDeserialize: deserialize_rest_UpdateRecordResponse,
  },
  apply: {
    path: '/rest.Record/Apply',
    requestStream: false,
    responseStream: false,
    requestType: stub_rest_record_pb.ApplyRecordRequest,
    responseType: stub_rest_record_pb.ApplyRecordResponse,
    requestSerialize: serialize_rest_ApplyRecordRequest,
    requestDeserialize: deserialize_rest_ApplyRecordRequest,
    responseSerialize: serialize_rest_ApplyRecordResponse,
    responseDeserialize: deserialize_rest_ApplyRecordResponse,
  },
  delete: {
    path: '/rest.Record/Delete',
    requestStream: false,
    responseStream: false,
    requestType: stub_rest_record_pb.DeleteRecordRequest,
    responseType: stub_rest_record_pb.DeleteRecordResponse,
    requestSerialize: serialize_rest_DeleteRecordRequest,
    requestDeserialize: deserialize_rest_DeleteRecordRequest,
    responseSerialize: serialize_rest_DeleteRecordResponse,
    responseDeserialize: deserialize_rest_DeleteRecordResponse,
  },
};

exports.RecordClient = grpc.makeGenericClientConstructor(RecordService);
