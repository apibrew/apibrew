// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ext_function_pb = require('../ext/function_pb.js');
var model_record_pb = require('../model/record_pb.js');
var model_resource_pb = require('../model/resource_pb.js');
var model_query_pb = require('../model/query_pb.js');
var model_error_pb = require('../model/error_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var google_protobuf_any_pb = require('google-protobuf/google/protobuf/any_pb.js');
var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');
var model_event_pb = require('../model/event_pb.js');

function serialize_ext_FunctionCallRequest(arg) {
  if (!(arg instanceof ext_function_pb.FunctionCallRequest)) {
    throw new Error('Expected argument of type ext.FunctionCallRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ext_FunctionCallRequest(buffer_arg) {
  return ext_function_pb.FunctionCallRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ext_FunctionCallResponse(arg) {
  if (!(arg instanceof ext_function_pb.FunctionCallResponse)) {
    throw new Error('Expected argument of type ext.FunctionCallResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ext_FunctionCallResponse(buffer_arg) {
  return ext_function_pb.FunctionCallResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var FunctionService = exports.FunctionService = {
  functionCall: {
    path: '/ext.Function/FunctionCall',
    requestStream: false,
    responseStream: false,
    requestType: ext_function_pb.FunctionCallRequest,
    responseType: ext_function_pb.FunctionCallResponse,
    requestSerialize: serialize_ext_FunctionCallRequest,
    requestDeserialize: deserialize_ext_FunctionCallRequest,
    responseSerialize: serialize_ext_FunctionCallResponse,
    responseDeserialize: deserialize_ext_FunctionCallResponse,
  },
};

exports.FunctionClient = grpc.makeGenericClientConstructor(FunctionService);
