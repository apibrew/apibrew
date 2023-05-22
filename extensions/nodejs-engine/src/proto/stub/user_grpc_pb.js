// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_user_pb = require('../stub/user_pb.js');
var model_error_pb = require('../model/error_pb.js');
var model_user_pb = require('../model/user_pb.js');
var model_query_pb = require('../model/query_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var gnostic_openapi_v3_annotations_pb = require('../gnostic/openapi/v3/annotations_pb.js');

function serialize_stub_CreateUserRequest(arg) {
  if (!(arg instanceof stub_user_pb.CreateUserRequest)) {
    throw new Error('Expected argument of type stub.CreateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateUserRequest(buffer_arg) {
  return stub_user_pb.CreateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_CreateUserResponse(arg) {
  if (!(arg instanceof stub_user_pb.CreateUserResponse)) {
    throw new Error('Expected argument of type stub.CreateUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateUserResponse(buffer_arg) {
  return stub_user_pb.CreateUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteUserRequest(arg) {
  if (!(arg instanceof stub_user_pb.DeleteUserRequest)) {
    throw new Error('Expected argument of type stub.DeleteUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteUserRequest(buffer_arg) {
  return stub_user_pb.DeleteUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteUserResponse(arg) {
  if (!(arg instanceof stub_user_pb.DeleteUserResponse)) {
    throw new Error('Expected argument of type stub.DeleteUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteUserResponse(buffer_arg) {
  return stub_user_pb.DeleteUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetUserRequest(arg) {
  if (!(arg instanceof stub_user_pb.GetUserRequest)) {
    throw new Error('Expected argument of type stub.GetUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetUserRequest(buffer_arg) {
  return stub_user_pb.GetUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetUserResponse(arg) {
  if (!(arg instanceof stub_user_pb.GetUserResponse)) {
    throw new Error('Expected argument of type stub.GetUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetUserResponse(buffer_arg) {
  return stub_user_pb.GetUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListUserRequest(arg) {
  if (!(arg instanceof stub_user_pb.ListUserRequest)) {
    throw new Error('Expected argument of type stub.ListUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListUserRequest(buffer_arg) {
  return stub_user_pb.ListUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListUserResponse(arg) {
  if (!(arg instanceof stub_user_pb.ListUserResponse)) {
    throw new Error('Expected argument of type stub.ListUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListUserResponse(buffer_arg) {
  return stub_user_pb.ListUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateUserRequest(arg) {
  if (!(arg instanceof stub_user_pb.UpdateUserRequest)) {
    throw new Error('Expected argument of type stub.UpdateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateUserRequest(buffer_arg) {
  return stub_user_pb.UpdateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateUserResponse(arg) {
  if (!(arg instanceof stub_user_pb.UpdateUserResponse)) {
    throw new Error('Expected argument of type stub.UpdateUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateUserResponse(buffer_arg) {
  return stub_user_pb.UpdateUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// User service is for managing users
var UserService = exports.UserService = {
  create: {
    path: '/stub.User/Create',
    requestStream: false,
    responseStream: false,
    requestType: stub_user_pb.CreateUserRequest,
    responseType: stub_user_pb.CreateUserResponse,
    requestSerialize: serialize_stub_CreateUserRequest,
    requestDeserialize: deserialize_stub_CreateUserRequest,
    responseSerialize: serialize_stub_CreateUserResponse,
    responseDeserialize: deserialize_stub_CreateUserResponse,
  },
  update: {
    path: '/stub.User/Update',
    requestStream: false,
    responseStream: false,
    requestType: stub_user_pb.UpdateUserRequest,
    responseType: stub_user_pb.UpdateUserResponse,
    requestSerialize: serialize_stub_UpdateUserRequest,
    requestDeserialize: deserialize_stub_UpdateUserRequest,
    responseSerialize: serialize_stub_UpdateUserResponse,
    responseDeserialize: deserialize_stub_UpdateUserResponse,
  },
  delete: {
    path: '/stub.User/Delete',
    requestStream: false,
    responseStream: false,
    requestType: stub_user_pb.DeleteUserRequest,
    responseType: stub_user_pb.DeleteUserResponse,
    requestSerialize: serialize_stub_DeleteUserRequest,
    requestDeserialize: deserialize_stub_DeleteUserRequest,
    responseSerialize: serialize_stub_DeleteUserResponse,
    responseDeserialize: deserialize_stub_DeleteUserResponse,
  },
  list: {
    path: '/stub.User/List',
    requestStream: false,
    responseStream: false,
    requestType: stub_user_pb.ListUserRequest,
    responseType: stub_user_pb.ListUserResponse,
    requestSerialize: serialize_stub_ListUserRequest,
    requestDeserialize: deserialize_stub_ListUserRequest,
    responseSerialize: serialize_stub_ListUserResponse,
    responseDeserialize: deserialize_stub_ListUserResponse,
  },
  get: {
    path: '/stub.User/Get',
    requestStream: false,
    responseStream: false,
    requestType: stub_user_pb.GetUserRequest,
    responseType: stub_user_pb.GetUserResponse,
    requestSerialize: serialize_stub_GetUserRequest,
    requestDeserialize: deserialize_stub_GetUserRequest,
    responseSerialize: serialize_stub_GetUserResponse,
    responseDeserialize: deserialize_stub_GetUserResponse,
  },
};

exports.UserClient = grpc.makeGenericClientConstructor(UserService);
