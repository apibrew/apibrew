// source: model/annotations.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

var google_protobuf_descriptor_pb = require('google-protobuf/google/protobuf/descriptor_pb.js');
goog.object.extend(proto, google_protobuf_descriptor_pb);
goog.exportSymbol('proto.model.hclBlock', null, global);
goog.exportSymbol('proto.model.hclIgnore', null, global);
goog.exportSymbol('proto.model.hclLabel', null, global);

/**
 * A tuple of {field number, class constructor} for the extension
 * field named `hclBlock`.
 * @type {!jspb.ExtensionFieldInfo<string>}
 */
proto.model.hclBlock = new jspb.ExtensionFieldInfo(
    1144,
    {hclBlock: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[1144] = new jspb.ExtensionFieldBinaryInfo(
    proto.model.hclBlock,
    jspb.BinaryReader.prototype.readString,
    jspb.BinaryWriter.prototype.writeString,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[1144] = proto.model.hclBlock;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `hclLabel`.
 * @type {!jspb.ExtensionFieldInfo<string>}
 */
proto.model.hclLabel = new jspb.ExtensionFieldInfo(
    1145,
    {hclLabel: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[1145] = new jspb.ExtensionFieldBinaryInfo(
    proto.model.hclLabel,
    jspb.BinaryReader.prototype.readString,
    jspb.BinaryWriter.prototype.writeString,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[1145] = proto.model.hclLabel;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `hclIgnore`.
 * @type {!jspb.ExtensionFieldInfo<boolean>}
 */
proto.model.hclIgnore = new jspb.ExtensionFieldInfo(
    1146,
    {hclIgnore: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[1146] = new jspb.ExtensionFieldBinaryInfo(
    proto.model.hclIgnore,
    jspb.BinaryReader.prototype.readBool,
    jspb.BinaryWriter.prototype.writeBool,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[1146] = proto.model.hclIgnore;

goog.object.extend(exports, proto.model);
