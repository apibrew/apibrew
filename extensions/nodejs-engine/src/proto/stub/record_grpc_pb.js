// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_record_pb = require('../stub/record_pb.js');
var model_record_pb = require('../model/record_pb.js');
var model_query_pb = require('../model/query_pb.js');
var model_error_pb = require('../model/error_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');

function serialize_model_Record(arg) {
  if (!(arg instanceof model_record_pb.Record)) {
    throw new Error('Expected argument of type model.Record');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_Record(buffer_arg) {
  return model_record_pb.Record.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ApplyRecordRequest(arg) {
  if (!(arg instanceof stub_record_pb.ApplyRecordRequest)) {
    throw new Error('Expected argument of type stub.ApplyRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ApplyRecordRequest(buffer_arg) {
  return stub_record_pb.ApplyRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ApplyRecordResponse(arg) {
  if (!(arg instanceof stub_record_pb.ApplyRecordResponse)) {
    throw new Error('Expected argument of type stub.ApplyRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ApplyRecordResponse(buffer_arg) {
  return stub_record_pb.ApplyRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_CreateRecordRequest(arg) {
  if (!(arg instanceof stub_record_pb.CreateRecordRequest)) {
    throw new Error('Expected argument of type stub.CreateRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateRecordRequest(buffer_arg) {
  return stub_record_pb.CreateRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_CreateRecordResponse(arg) {
  if (!(arg instanceof stub_record_pb.CreateRecordResponse)) {
    throw new Error('Expected argument of type stub.CreateRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateRecordResponse(buffer_arg) {
  return stub_record_pb.CreateRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteRecordRequest(arg) {
  if (!(arg instanceof stub_record_pb.DeleteRecordRequest)) {
    throw new Error('Expected argument of type stub.DeleteRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteRecordRequest(buffer_arg) {
  return stub_record_pb.DeleteRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteRecordResponse(arg) {
  if (!(arg instanceof stub_record_pb.DeleteRecordResponse)) {
    throw new Error('Expected argument of type stub.DeleteRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteRecordResponse(buffer_arg) {
  return stub_record_pb.DeleteRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetRecordRequest(arg) {
  if (!(arg instanceof stub_record_pb.GetRecordRequest)) {
    throw new Error('Expected argument of type stub.GetRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetRecordRequest(buffer_arg) {
  return stub_record_pb.GetRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetRecordResponse(arg) {
  if (!(arg instanceof stub_record_pb.GetRecordResponse)) {
    throw new Error('Expected argument of type stub.GetRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetRecordResponse(buffer_arg) {
  return stub_record_pb.GetRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListRecordRequest(arg) {
  if (!(arg instanceof stub_record_pb.ListRecordRequest)) {
    throw new Error('Expected argument of type stub.ListRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListRecordRequest(buffer_arg) {
  return stub_record_pb.ListRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListRecordResponse(arg) {
  if (!(arg instanceof stub_record_pb.ListRecordResponse)) {
    throw new Error('Expected argument of type stub.ListRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListRecordResponse(buffer_arg) {
  return stub_record_pb.ListRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ReadStreamRequest(arg) {
  if (!(arg instanceof stub_record_pb.ReadStreamRequest)) {
    throw new Error('Expected argument of type stub.ReadStreamRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ReadStreamRequest(buffer_arg) {
  return stub_record_pb.ReadStreamRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_SearchRecordRequest(arg) {
  if (!(arg instanceof stub_record_pb.SearchRecordRequest)) {
    throw new Error('Expected argument of type stub.SearchRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_SearchRecordRequest(buffer_arg) {
  return stub_record_pb.SearchRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_SearchRecordResponse(arg) {
  if (!(arg instanceof stub_record_pb.SearchRecordResponse)) {
    throw new Error('Expected argument of type stub.SearchRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_SearchRecordResponse(buffer_arg) {
  return stub_record_pb.SearchRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateMultiRecordRequest(arg) {
  if (!(arg instanceof stub_record_pb.UpdateMultiRecordRequest)) {
    throw new Error('Expected argument of type stub.UpdateMultiRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateMultiRecordRequest(buffer_arg) {
  return stub_record_pb.UpdateMultiRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateMultiRecordResponse(arg) {
  if (!(arg instanceof stub_record_pb.UpdateMultiRecordResponse)) {
    throw new Error('Expected argument of type stub.UpdateMultiRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateMultiRecordResponse(buffer_arg) {
  return stub_record_pb.UpdateMultiRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateRecordRequest(arg) {
  if (!(arg instanceof stub_record_pb.UpdateRecordRequest)) {
    throw new Error('Expected argument of type stub.UpdateRecordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateRecordRequest(buffer_arg) {
  return stub_record_pb.UpdateRecordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateRecordResponse(arg) {
  if (!(arg instanceof stub_record_pb.UpdateRecordResponse)) {
    throw new Error('Expected argument of type stub.UpdateRecordResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateRecordResponse(buffer_arg) {
  return stub_record_pb.UpdateRecordResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_WriteStreamResponse(arg) {
  if (!(arg instanceof stub_record_pb.WriteStreamResponse)) {
    throw new Error('Expected argument of type stub.WriteStreamResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_WriteStreamResponse(buffer_arg) {
  return stub_record_pb.WriteStreamResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Record service is an abstract service for records of all resources. You can do CRUD like operations with Record service
var RecordService = exports.RecordService = {
  create: {
    path: '/stub.Record/Create',
    requestStream: false,
    responseStream: false,
    requestType: stub_record_pb.CreateRecordRequest,
    responseType: stub_record_pb.CreateRecordResponse,
    requestSerialize: serialize_stub_CreateRecordRequest,
    requestDeserialize: deserialize_stub_CreateRecordRequest,
    responseSerialize: serialize_stub_CreateRecordResponse,
    responseDeserialize: deserialize_stub_CreateRecordResponse,
  },
  update: {
    path: '/stub.Record/Update',
    requestStream: false,
    responseStream: false,
    requestType: stub_record_pb.UpdateRecordRequest,
    responseType: stub_record_pb.UpdateRecordResponse,
    requestSerialize: serialize_stub_UpdateRecordRequest,
    requestDeserialize: deserialize_stub_UpdateRecordRequest,
    responseSerialize: serialize_stub_UpdateRecordResponse,
    responseDeserialize: deserialize_stub_UpdateRecordResponse,
  },
  apply: {
    path: '/stub.Record/Apply',
    requestStream: false,
    responseStream: false,
    requestType: stub_record_pb.ApplyRecordRequest,
    responseType: stub_record_pb.ApplyRecordResponse,
    requestSerialize: serialize_stub_ApplyRecordRequest,
    requestDeserialize: deserialize_stub_ApplyRecordRequest,
    responseSerialize: serialize_stub_ApplyRecordResponse,
    responseDeserialize: deserialize_stub_ApplyRecordResponse,
  },
  //
// Not implemented yet
updateMulti: {
    path: '/stub.Record/UpdateMulti',
    requestStream: false,
    responseStream: false,
    requestType: stub_record_pb.UpdateMultiRecordRequest,
    responseType: stub_record_pb.UpdateMultiRecordResponse,
    requestSerialize: serialize_stub_UpdateMultiRecordRequest,
    requestDeserialize: deserialize_stub_UpdateMultiRecordRequest,
    responseSerialize: serialize_stub_UpdateMultiRecordResponse,
    responseDeserialize: deserialize_stub_UpdateMultiRecordResponse,
  },
  delete: {
    path: '/stub.Record/Delete',
    requestStream: false,
    responseStream: false,
    requestType: stub_record_pb.DeleteRecordRequest,
    responseType: stub_record_pb.DeleteRecordResponse,
    requestSerialize: serialize_stub_DeleteRecordRequest,
    requestDeserialize: deserialize_stub_DeleteRecordRequest,
    responseSerialize: serialize_stub_DeleteRecordResponse,
    responseDeserialize: deserialize_stub_DeleteRecordResponse,
  },
  list: {
    path: '/stub.Record/List',
    requestStream: false,
    responseStream: false,
    requestType: stub_record_pb.ListRecordRequest,
    responseType: stub_record_pb.ListRecordResponse,
    requestSerialize: serialize_stub_ListRecordRequest,
    requestDeserialize: deserialize_stub_ListRecordRequest,
    responseSerialize: serialize_stub_ListRecordResponse,
    responseDeserialize: deserialize_stub_ListRecordResponse,
  },
  search: {
    path: '/stub.Record/Search',
    requestStream: false,
    responseStream: false,
    requestType: stub_record_pb.SearchRecordRequest,
    responseType: stub_record_pb.SearchRecordResponse,
    requestSerialize: serialize_stub_SearchRecordRequest,
    requestDeserialize: deserialize_stub_SearchRecordRequest,
    responseSerialize: serialize_stub_SearchRecordResponse,
    responseDeserialize: deserialize_stub_SearchRecordResponse,
  },
  readStream: {
    path: '/stub.Record/ReadStream',
    requestStream: false,
    responseStream: true,
    requestType: stub_record_pb.ReadStreamRequest,
    responseType: model_record_pb.Record,
    requestSerialize: serialize_stub_ReadStreamRequest,
    requestDeserialize: deserialize_stub_ReadStreamRequest,
    responseSerialize: serialize_model_Record,
    responseDeserialize: deserialize_model_Record,
  },
  //
// Not implemented yet
writeStream: {
    path: '/stub.Record/WriteStream',
    requestStream: true,
    responseStream: false,
    requestType: model_record_pb.Record,
    responseType: stub_record_pb.WriteStreamResponse,
    requestSerialize: serialize_model_Record,
    requestDeserialize: deserialize_model_Record,
    responseSerialize: serialize_stub_WriteStreamResponse,
    responseDeserialize: deserialize_stub_WriteStreamResponse,
  },
  get: {
    path: '/stub.Record/Get',
    requestStream: false,
    responseStream: false,
    requestType: stub_record_pb.GetRecordRequest,
    responseType: stub_record_pb.GetRecordResponse,
    requestSerialize: serialize_stub_GetRecordRequest,
    requestDeserialize: deserialize_stub_GetRecordRequest,
    responseSerialize: serialize_stub_GetRecordResponse,
    responseDeserialize: deserialize_stub_GetRecordResponse,
  },
};

exports.RecordClient = grpc.makeGenericClientConstructor(RecordService);
