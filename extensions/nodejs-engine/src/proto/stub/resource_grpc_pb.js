// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_resource_pb = require('../stub/resource_pb.js');
var model_error_pb = require('../model/error_pb.js');
var model_resource_pb = require('../model/resource_pb.js');
var model_resource$migration_pb = require('../model/resource-migration_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var gnostic_openapi_v3_annotations_pb = require('../gnostic/openapi/v3/annotations_pb.js');

function serialize_stub_CreateResourceRequest(arg) {
  if (!(arg instanceof stub_resource_pb.CreateResourceRequest)) {
    throw new Error('Expected argument of type stub.CreateResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateResourceRequest(buffer_arg) {
  return stub_resource_pb.CreateResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_CreateResourceResponse(arg) {
  if (!(arg instanceof stub_resource_pb.CreateResourceResponse)) {
    throw new Error('Expected argument of type stub.CreateResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateResourceResponse(buffer_arg) {
  return stub_resource_pb.CreateResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteResourceRequest(arg) {
  if (!(arg instanceof stub_resource_pb.DeleteResourceRequest)) {
    throw new Error('Expected argument of type stub.DeleteResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteResourceRequest(buffer_arg) {
  return stub_resource_pb.DeleteResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteResourceResponse(arg) {
  if (!(arg instanceof stub_resource_pb.DeleteResourceResponse)) {
    throw new Error('Expected argument of type stub.DeleteResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteResourceResponse(buffer_arg) {
  return stub_resource_pb.DeleteResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetResourceByNameRequest(arg) {
  if (!(arg instanceof stub_resource_pb.GetResourceByNameRequest)) {
    throw new Error('Expected argument of type stub.GetResourceByNameRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetResourceByNameRequest(buffer_arg) {
  return stub_resource_pb.GetResourceByNameRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetResourceByNameResponse(arg) {
  if (!(arg instanceof stub_resource_pb.GetResourceByNameResponse)) {
    throw new Error('Expected argument of type stub.GetResourceByNameResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetResourceByNameResponse(buffer_arg) {
  return stub_resource_pb.GetResourceByNameResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetResourceRequest(arg) {
  if (!(arg instanceof stub_resource_pb.GetResourceRequest)) {
    throw new Error('Expected argument of type stub.GetResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetResourceRequest(buffer_arg) {
  return stub_resource_pb.GetResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetResourceResponse(arg) {
  if (!(arg instanceof stub_resource_pb.GetResourceResponse)) {
    throw new Error('Expected argument of type stub.GetResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetResourceResponse(buffer_arg) {
  return stub_resource_pb.GetResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetSystemResourceRequest(arg) {
  if (!(arg instanceof stub_resource_pb.GetSystemResourceRequest)) {
    throw new Error('Expected argument of type stub.GetSystemResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetSystemResourceRequest(buffer_arg) {
  return stub_resource_pb.GetSystemResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetSystemResourceResponse(arg) {
  if (!(arg instanceof stub_resource_pb.GetSystemResourceResponse)) {
    throw new Error('Expected argument of type stub.GetSystemResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetSystemResourceResponse(buffer_arg) {
  return stub_resource_pb.GetSystemResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListResourceRequest(arg) {
  if (!(arg instanceof stub_resource_pb.ListResourceRequest)) {
    throw new Error('Expected argument of type stub.ListResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListResourceRequest(buffer_arg) {
  return stub_resource_pb.ListResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListResourceResponse(arg) {
  if (!(arg instanceof stub_resource_pb.ListResourceResponse)) {
    throw new Error('Expected argument of type stub.ListResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListResourceResponse(buffer_arg) {
  return stub_resource_pb.ListResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_PrepareResourceMigrationPlanRequest(arg) {
  if (!(arg instanceof stub_resource_pb.PrepareResourceMigrationPlanRequest)) {
    throw new Error('Expected argument of type stub.PrepareResourceMigrationPlanRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_PrepareResourceMigrationPlanRequest(buffer_arg) {
  return stub_resource_pb.PrepareResourceMigrationPlanRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_PrepareResourceMigrationPlanResponse(arg) {
  if (!(arg instanceof stub_resource_pb.PrepareResourceMigrationPlanResponse)) {
    throw new Error('Expected argument of type stub.PrepareResourceMigrationPlanResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_PrepareResourceMigrationPlanResponse(buffer_arg) {
  return stub_resource_pb.PrepareResourceMigrationPlanResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateResourceRequest(arg) {
  if (!(arg instanceof stub_resource_pb.UpdateResourceRequest)) {
    throw new Error('Expected argument of type stub.UpdateResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateResourceRequest(buffer_arg) {
  return stub_resource_pb.UpdateResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateResourceResponse(arg) {
  if (!(arg instanceof stub_resource_pb.UpdateResourceResponse)) {
    throw new Error('Expected argument of type stub.UpdateResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateResourceResponse(buffer_arg) {
  return stub_resource_pb.UpdateResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Resource service is for managing resources
var ResourceService = exports.ResourceService = {
  create: {
    path: '/stub.Resource/Create',
    requestStream: false,
    responseStream: false,
    requestType: stub_resource_pb.CreateResourceRequest,
    responseType: stub_resource_pb.CreateResourceResponse,
    requestSerialize: serialize_stub_CreateResourceRequest,
    requestDeserialize: deserialize_stub_CreateResourceRequest,
    responseSerialize: serialize_stub_CreateResourceResponse,
    responseDeserialize: deserialize_stub_CreateResourceResponse,
  },
  update: {
    path: '/stub.Resource/Update',
    requestStream: false,
    responseStream: false,
    requestType: stub_resource_pb.UpdateResourceRequest,
    responseType: stub_resource_pb.UpdateResourceResponse,
    requestSerialize: serialize_stub_UpdateResourceRequest,
    requestDeserialize: deserialize_stub_UpdateResourceRequest,
    responseSerialize: serialize_stub_UpdateResourceResponse,
    responseDeserialize: deserialize_stub_UpdateResourceResponse,
  },
  delete: {
    path: '/stub.Resource/Delete',
    requestStream: false,
    responseStream: false,
    requestType: stub_resource_pb.DeleteResourceRequest,
    responseType: stub_resource_pb.DeleteResourceResponse,
    requestSerialize: serialize_stub_DeleteResourceRequest,
    requestDeserialize: deserialize_stub_DeleteResourceRequest,
    responseSerialize: serialize_stub_DeleteResourceResponse,
    responseDeserialize: deserialize_stub_DeleteResourceResponse,
  },
  list: {
    path: '/stub.Resource/List',
    requestStream: false,
    responseStream: false,
    requestType: stub_resource_pb.ListResourceRequest,
    responseType: stub_resource_pb.ListResourceResponse,
    requestSerialize: serialize_stub_ListResourceRequest,
    requestDeserialize: deserialize_stub_ListResourceRequest,
    responseSerialize: serialize_stub_ListResourceResponse,
    responseDeserialize: deserialize_stub_ListResourceResponse,
  },
  //
// PrepareResourceMigrationPlan will prepare the migration plan for the resources, it will not do any migration. It will just return the plan for the migration.
prepareResourceMigrationPlan: {
    path: '/stub.Resource/PrepareResourceMigrationPlan',
    requestStream: false,
    responseStream: false,
    requestType: stub_resource_pb.PrepareResourceMigrationPlanRequest,
    responseType: stub_resource_pb.PrepareResourceMigrationPlanResponse,
    requestSerialize: serialize_stub_PrepareResourceMigrationPlanRequest,
    requestDeserialize: deserialize_stub_PrepareResourceMigrationPlanRequest,
    responseSerialize: serialize_stub_PrepareResourceMigrationPlanResponse,
    responseDeserialize: deserialize_stub_PrepareResourceMigrationPlanResponse,
  },
  get: {
    path: '/stub.Resource/Get',
    requestStream: false,
    responseStream: false,
    requestType: stub_resource_pb.GetResourceRequest,
    responseType: stub_resource_pb.GetResourceResponse,
    requestSerialize: serialize_stub_GetResourceRequest,
    requestDeserialize: deserialize_stub_GetResourceRequest,
    responseSerialize: serialize_stub_GetResourceResponse,
    responseDeserialize: deserialize_stub_GetResourceResponse,
  },
  getByName: {
    path: '/stub.Resource/GetByName',
    requestStream: false,
    responseStream: false,
    requestType: stub_resource_pb.GetResourceByNameRequest,
    responseType: stub_resource_pb.GetResourceByNameResponse,
    requestSerialize: serialize_stub_GetResourceByNameRequest,
    requestDeserialize: deserialize_stub_GetResourceByNameRequest,
    responseSerialize: serialize_stub_GetResourceByNameResponse,
    responseDeserialize: deserialize_stub_GetResourceByNameResponse,
  },
  getSystemResource: {
    path: '/stub.Resource/GetSystemResource',
    requestStream: false,
    responseStream: false,
    requestType: stub_resource_pb.GetSystemResourceRequest,
    responseType: stub_resource_pb.GetSystemResourceResponse,
    requestSerialize: serialize_stub_GetSystemResourceRequest,
    requestDeserialize: deserialize_stub_GetSystemResourceRequest,
    responseSerialize: serialize_stub_GetSystemResourceResponse,
    responseDeserialize: deserialize_stub_GetSystemResourceResponse,
  },
};

exports.ResourceClient = grpc.makeGenericClientConstructor(ResourceService);
