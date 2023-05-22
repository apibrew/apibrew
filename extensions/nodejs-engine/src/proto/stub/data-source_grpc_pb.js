// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var stub_data$source_pb = require('../stub/data-source_pb.js');
var model_data$source_pb = require('../model/data-source_pb.js');
var model_error_pb = require('../model/error_pb.js');
var model_resource_pb = require('../model/resource_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var gnostic_openapi_v3_annotations_pb = require('../gnostic/openapi/v3/annotations_pb.js');

function serialize_stub_CreateDataSourceRequest(arg) {
  if (!(arg instanceof stub_data$source_pb.CreateDataSourceRequest)) {
    throw new Error('Expected argument of type stub.CreateDataSourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateDataSourceRequest(buffer_arg) {
  return stub_data$source_pb.CreateDataSourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_CreateDataSourceResponse(arg) {
  if (!(arg instanceof stub_data$source_pb.CreateDataSourceResponse)) {
    throw new Error('Expected argument of type stub.CreateDataSourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_CreateDataSourceResponse(buffer_arg) {
  return stub_data$source_pb.CreateDataSourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteDataSourceRequest(arg) {
  if (!(arg instanceof stub_data$source_pb.DeleteDataSourceRequest)) {
    throw new Error('Expected argument of type stub.DeleteDataSourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteDataSourceRequest(buffer_arg) {
  return stub_data$source_pb.DeleteDataSourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_DeleteDataSourceResponse(arg) {
  if (!(arg instanceof stub_data$source_pb.DeleteDataSourceResponse)) {
    throw new Error('Expected argument of type stub.DeleteDataSourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_DeleteDataSourceResponse(buffer_arg) {
  return stub_data$source_pb.DeleteDataSourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetDataSourceRequest(arg) {
  if (!(arg instanceof stub_data$source_pb.GetDataSourceRequest)) {
    throw new Error('Expected argument of type stub.GetDataSourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetDataSourceRequest(buffer_arg) {
  return stub_data$source_pb.GetDataSourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_GetDataSourceResponse(arg) {
  if (!(arg instanceof stub_data$source_pb.GetDataSourceResponse)) {
    throw new Error('Expected argument of type stub.GetDataSourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_GetDataSourceResponse(buffer_arg) {
  return stub_data$source_pb.GetDataSourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListDataSourceRequest(arg) {
  if (!(arg instanceof stub_data$source_pb.ListDataSourceRequest)) {
    throw new Error('Expected argument of type stub.ListDataSourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListDataSourceRequest(buffer_arg) {
  return stub_data$source_pb.ListDataSourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListDataSourceResponse(arg) {
  if (!(arg instanceof stub_data$source_pb.ListDataSourceResponse)) {
    throw new Error('Expected argument of type stub.ListDataSourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListDataSourceResponse(buffer_arg) {
  return stub_data$source_pb.ListDataSourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListEntitiesRequest(arg) {
  if (!(arg instanceof stub_data$source_pb.ListEntitiesRequest)) {
    throw new Error('Expected argument of type stub.ListEntitiesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListEntitiesRequest(buffer_arg) {
  return stub_data$source_pb.ListEntitiesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_ListEntitiesResponse(arg) {
  if (!(arg instanceof stub_data$source_pb.ListEntitiesResponse)) {
    throw new Error('Expected argument of type stub.ListEntitiesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_ListEntitiesResponse(buffer_arg) {
  return stub_data$source_pb.ListEntitiesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_PrepareResourceFromEntityRequest(arg) {
  if (!(arg instanceof stub_data$source_pb.PrepareResourceFromEntityRequest)) {
    throw new Error('Expected argument of type stub.PrepareResourceFromEntityRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_PrepareResourceFromEntityRequest(buffer_arg) {
  return stub_data$source_pb.PrepareResourceFromEntityRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_PrepareResourceFromEntityResponse(arg) {
  if (!(arg instanceof stub_data$source_pb.PrepareResourceFromEntityResponse)) {
    throw new Error('Expected argument of type stub.PrepareResourceFromEntityResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_PrepareResourceFromEntityResponse(buffer_arg) {
  return stub_data$source_pb.PrepareResourceFromEntityResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_StatusRequest(arg) {
  if (!(arg instanceof stub_data$source_pb.StatusRequest)) {
    throw new Error('Expected argument of type stub.StatusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_StatusRequest(buffer_arg) {
  return stub_data$source_pb.StatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_StatusResponse(arg) {
  if (!(arg instanceof stub_data$source_pb.StatusResponse)) {
    throw new Error('Expected argument of type stub.StatusResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_StatusResponse(buffer_arg) {
  return stub_data$source_pb.StatusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateDataSourceRequest(arg) {
  if (!(arg instanceof stub_data$source_pb.UpdateDataSourceRequest)) {
    throw new Error('Expected argument of type stub.UpdateDataSourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateDataSourceRequest(buffer_arg) {
  return stub_data$source_pb.UpdateDataSourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stub_UpdateDataSourceResponse(arg) {
  if (!(arg instanceof stub_data$source_pb.UpdateDataSourceResponse)) {
    throw new Error('Expected argument of type stub.UpdateDataSourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stub_UpdateDataSourceResponse(buffer_arg) {
  return stub_data$source_pb.UpdateDataSourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// DataSource Service is for managing data sources
var DataSourceService = exports.DataSourceService = {
  create: {
    path: '/stub.DataSource/Create',
    requestStream: false,
    responseStream: false,
    requestType: stub_data$source_pb.CreateDataSourceRequest,
    responseType: stub_data$source_pb.CreateDataSourceResponse,
    requestSerialize: serialize_stub_CreateDataSourceRequest,
    requestDeserialize: deserialize_stub_CreateDataSourceRequest,
    responseSerialize: serialize_stub_CreateDataSourceResponse,
    responseDeserialize: deserialize_stub_CreateDataSourceResponse,
  },
  list: {
    path: '/stub.DataSource/List',
    requestStream: false,
    responseStream: false,
    requestType: stub_data$source_pb.ListDataSourceRequest,
    responseType: stub_data$source_pb.ListDataSourceResponse,
    requestSerialize: serialize_stub_ListDataSourceRequest,
    requestDeserialize: deserialize_stub_ListDataSourceRequest,
    responseSerialize: serialize_stub_ListDataSourceResponse,
    responseDeserialize: deserialize_stub_ListDataSourceResponse,
  },
  update: {
    path: '/stub.DataSource/Update',
    requestStream: false,
    responseStream: false,
    requestType: stub_data$source_pb.UpdateDataSourceRequest,
    responseType: stub_data$source_pb.UpdateDataSourceResponse,
    requestSerialize: serialize_stub_UpdateDataSourceRequest,
    requestDeserialize: deserialize_stub_UpdateDataSourceRequest,
    responseSerialize: serialize_stub_UpdateDataSourceResponse,
    responseDeserialize: deserialize_stub_UpdateDataSourceResponse,
  },
  delete: {
    path: '/stub.DataSource/Delete',
    requestStream: false,
    responseStream: false,
    requestType: stub_data$source_pb.DeleteDataSourceRequest,
    responseType: stub_data$source_pb.DeleteDataSourceResponse,
    requestSerialize: serialize_stub_DeleteDataSourceRequest,
    requestDeserialize: deserialize_stub_DeleteDataSourceRequest,
    responseSerialize: serialize_stub_DeleteDataSourceResponse,
    responseDeserialize: deserialize_stub_DeleteDataSourceResponse,
  },
  get: {
    path: '/stub.DataSource/Get',
    requestStream: false,
    responseStream: false,
    requestType: stub_data$source_pb.GetDataSourceRequest,
    responseType: stub_data$source_pb.GetDataSourceResponse,
    requestSerialize: serialize_stub_GetDataSourceRequest,
    requestDeserialize: deserialize_stub_GetDataSourceRequest,
    responseSerialize: serialize_stub_GetDataSourceResponse,
    responseDeserialize: deserialize_stub_GetDataSourceResponse,
  },
  //
// Status will return connection status of data source
status: {
    path: '/stub.DataSource/Status',
    requestStream: false,
    responseStream: false,
    requestType: stub_data$source_pb.StatusRequest,
    responseType: stub_data$source_pb.StatusResponse,
    requestSerialize: serialize_stub_StatusRequest,
    requestDeserialize: deserialize_stub_StatusRequest,
    responseSerialize: serialize_stub_StatusResponse,
    responseDeserialize: deserialize_stub_StatusResponse,
  },
  //
// List entities will return all entities from data source
listEntities: {
    path: '/stub.DataSource/ListEntities',
    requestStream: false,
    responseStream: false,
    requestType: stub_data$source_pb.ListEntitiesRequest,
    responseType: stub_data$source_pb.ListEntitiesResponse,
    requestSerialize: serialize_stub_ListEntitiesRequest,
    requestDeserialize: deserialize_stub_ListEntitiesRequest,
    responseSerialize: serialize_stub_ListEntitiesResponse,
    responseDeserialize: deserialize_stub_ListEntitiesResponse,
  },
  //
// PrepareResourceFromEntity will return resource from data source based on entity.
// It is for database first approach. If you already have an entity/table on data source and your want to create resource based on it, you can call this endpoint to do it.
prepareResourceFromEntity: {
    path: '/stub.DataSource/PrepareResourceFromEntity',
    requestStream: false,
    responseStream: false,
    requestType: stub_data$source_pb.PrepareResourceFromEntityRequest,
    responseType: stub_data$source_pb.PrepareResourceFromEntityResponse,
    requestSerialize: serialize_stub_PrepareResourceFromEntityRequest,
    requestDeserialize: deserialize_stub_PrepareResourceFromEntityRequest,
    responseSerialize: serialize_stub_PrepareResourceFromEntityResponse,
    responseDeserialize: deserialize_stub_PrepareResourceFromEntityResponse,
  },
};

exports.DataSourceClient = grpc.makeGenericClientConstructor(DataSourceService);
