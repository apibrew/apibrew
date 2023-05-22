// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_watch_pb = require('../stub/watch_pb.js');
var model_event_pb = require('../model/event_pb.js');
var model_query_pb = require('../model/query_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var gnostic_openapi_v3_annotations_pb = require('../gnostic/openapi/v3/annotations_pb.js');

function serialize_model_Event(arg) {
  if (!(arg instanceof model_event_pb.Event)) {
    throw new Error('Expected argument of type model.Event');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_Event(buffer_arg) {
  return model_event_pb.Event.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_WatchRequest(arg) {
  if (!(arg instanceof stub_watch_pb.WatchRequest)) {
    throw new Error('Expected argument of type stub.WatchRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_WatchRequest(buffer_arg) {
  return stub_watch_pb.WatchRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


// Watch service watching operations on records
var WatchService = exports.WatchService = {
  // Sends a greeting
watch: {
    path: '/stub.Watch/Watch',
    requestStream: false,
    responseStream: true,
    requestType: stub_watch_pb.WatchRequest,
    responseType: model_event_pb.Event,
    requestSerialize: serialize_stub_WatchRequest,
    requestDeserialize: deserialize_stub_WatchRequest,
    responseSerialize: serialize_model_Event,
    responseDeserialize: deserialize_model_Event,
  },
};

exports.WatchClient = grpc.makeGenericClientConstructor(WatchService);
