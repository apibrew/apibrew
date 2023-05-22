// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_authentication_pb = require('../stub/authentication_pb.js');
var model_error_pb = require('../model/error_pb.js');
var model_token_pb = require('../model/token_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var gnostic_openapi_v3_annotations_pb = require('../gnostic/openapi/v3/annotations_pb.js');

function serialize_stub_AuthenticationRequest(arg) {
  if (!(arg instanceof stub_authentication_pb.AuthenticationRequest)) {
    throw new Error('Expected argument of type stub.AuthenticationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_AuthenticationRequest(buffer_arg) {
  return stub_authentication_pb.AuthenticationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_AuthenticationResponse(arg) {
  if (!(arg instanceof stub_authentication_pb.AuthenticationResponse)) {
    throw new Error('Expected argument of type stub.AuthenticationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_AuthenticationResponse(buffer_arg) {
  return stub_authentication_pb.AuthenticationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_RenewTokenRequest(arg) {
  if (!(arg instanceof stub_authentication_pb.RenewTokenRequest)) {
    throw new Error('Expected argument of type stub.RenewTokenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_RenewTokenRequest(buffer_arg) {
  return stub_authentication_pb.RenewTokenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_RenewTokenResponse(arg) {
  if (!(arg instanceof stub_authentication_pb.RenewTokenResponse)) {
    throw new Error('Expected argument of type stub.RenewTokenResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_RenewTokenResponse(buffer_arg) {
  return stub_authentication_pb.RenewTokenResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Authentication Service is for authentication related operations
var AuthenticationService = exports.AuthenticationService = {
  //
// Authentication with username/password and create new token
// Later on, you need to use this token to access other services, for grpc, you need to set the token on request. For Rest, you need to set the token on Authorization header with Bearer prefix
authenticate: {
    path: '/stub.Authentication/Authenticate',
    requestStream: false,
    responseStream: false,
    requestType: stub_authentication_pb.AuthenticationRequest,
    responseType: stub_authentication_pb.AuthenticationResponse,
    requestSerialize: serialize_stub_AuthenticationRequest,
    requestDeserialize: deserialize_stub_AuthenticationRequest,
    responseSerialize: serialize_stub_AuthenticationResponse,
    responseDeserialize: deserialize_stub_AuthenticationResponse,
  },
  //
// Renew token with existing token
renewToken: {
    path: '/stub.Authentication/RenewToken',
    requestStream: false,
    responseStream: false,
    requestType: stub_authentication_pb.RenewTokenRequest,
    responseType: stub_authentication_pb.RenewTokenResponse,
    requestSerialize: serialize_stub_RenewTokenRequest,
    requestDeserialize: deserialize_stub_RenewTokenRequest,
    responseSerialize: serialize_stub_RenewTokenResponse,
    responseDeserialize: deserialize_stub_RenewTokenResponse,
  },
};

exports.AuthenticationClient = grpc.makeGenericClientConstructor(AuthenticationService);
