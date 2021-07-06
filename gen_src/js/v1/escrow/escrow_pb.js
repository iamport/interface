// source: v1/escrow/escrow.proto
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
var global = Function('return this')();

var google_api_annotations_pb = require('../../google/api/annotations_pb.js');
goog.object.extend(proto, google_api_annotations_pb);
goog.exportSymbol('proto.escrow.EscrowRequest', null, global);
goog.exportSymbol('proto.escrow.EscrowResponse', null, global);
goog.exportSymbol('proto.escrow.Info', null, global);
goog.exportSymbol('proto.escrow.Logis', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.escrow.Info = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.escrow.Info, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.escrow.Info.displayName = 'proto.escrow.Info';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.escrow.Logis = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.escrow.Logis, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.escrow.Logis.displayName = 'proto.escrow.Logis';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.escrow.EscrowRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.escrow.EscrowRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.escrow.EscrowRequest.displayName = 'proto.escrow.EscrowRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.escrow.EscrowResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.escrow.EscrowResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.escrow.EscrowResponse.displayName = 'proto.escrow.EscrowResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.escrow.Info.prototype.toObject = function(opt_includeInstance) {
  return proto.escrow.Info.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.escrow.Info} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.escrow.Info.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    tel: jspb.Message.getFieldWithDefault(msg, 2, ""),
    addr: jspb.Message.getFieldWithDefault(msg, 3, ""),
    postcode: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.escrow.Info}
 */
proto.escrow.Info.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.escrow.Info;
  return proto.escrow.Info.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.escrow.Info} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.escrow.Info}
 */
proto.escrow.Info.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTel(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddr(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setPostcode(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.escrow.Info.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.escrow.Info.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.escrow.Info} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.escrow.Info.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getTel();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getAddr();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getPostcode();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.escrow.Info.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.escrow.Info} returns this
 */
proto.escrow.Info.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string tel = 2;
 * @return {string}
 */
proto.escrow.Info.prototype.getTel = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.escrow.Info} returns this
 */
proto.escrow.Info.prototype.setTel = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string addr = 3;
 * @return {string}
 */
proto.escrow.Info.prototype.getAddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.escrow.Info} returns this
 */
proto.escrow.Info.prototype.setAddr = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string postcode = 4;
 * @return {string}
 */
proto.escrow.Info.prototype.getPostcode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.escrow.Info} returns this
 */
proto.escrow.Info.prototype.setPostcode = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.escrow.Logis.prototype.toObject = function(opt_includeInstance) {
  return proto.escrow.Logis.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.escrow.Logis} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.escrow.Logis.toObject = function(includeInstance, msg) {
  var f, obj = {
    company: jspb.Message.getFieldWithDefault(msg, 1, ""),
    invoice: jspb.Message.getFieldWithDefault(msg, 2, ""),
    sentAt: jspb.Message.getFieldWithDefault(msg, 3, 0),
    appliedAt: jspb.Message.getFieldWithDefault(msg, 4, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.escrow.Logis}
 */
proto.escrow.Logis.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.escrow.Logis;
  return proto.escrow.Logis.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.escrow.Logis} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.escrow.Logis}
 */
proto.escrow.Logis.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setCompany(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setInvoice(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSentAt(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAppliedAt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.escrow.Logis.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.escrow.Logis.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.escrow.Logis} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.escrow.Logis.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCompany();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getInvoice();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getSentAt();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getAppliedAt();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
};


/**
 * optional string company = 1;
 * @return {string}
 */
proto.escrow.Logis.prototype.getCompany = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.escrow.Logis} returns this
 */
proto.escrow.Logis.prototype.setCompany = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string invoice = 2;
 * @return {string}
 */
proto.escrow.Logis.prototype.getInvoice = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.escrow.Logis} returns this
 */
proto.escrow.Logis.prototype.setInvoice = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int32 sent_at = 3;
 * @return {number}
 */
proto.escrow.Logis.prototype.getSentAt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.escrow.Logis} returns this
 */
proto.escrow.Logis.prototype.setSentAt = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 applied_at = 4;
 * @return {number}
 */
proto.escrow.Logis.prototype.getAppliedAt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.escrow.Logis} returns this
 */
proto.escrow.Logis.prototype.setAppliedAt = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.escrow.EscrowRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.escrow.EscrowRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.escrow.EscrowRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.escrow.EscrowRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    impUid: jspb.Message.getFieldWithDefault(msg, 1, ""),
    sender: (f = msg.getSender()) && proto.escrow.Info.toObject(includeInstance, f),
    receiver: (f = msg.getReceiver()) && proto.escrow.Info.toObject(includeInstance, f),
    logis: (f = msg.getLogis()) && proto.escrow.Logis.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.escrow.EscrowRequest}
 */
proto.escrow.EscrowRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.escrow.EscrowRequest;
  return proto.escrow.EscrowRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.escrow.EscrowRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.escrow.EscrowRequest}
 */
proto.escrow.EscrowRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setImpUid(value);
      break;
    case 2:
      var value = new proto.escrow.Info;
      reader.readMessage(value,proto.escrow.Info.deserializeBinaryFromReader);
      msg.setSender(value);
      break;
    case 3:
      var value = new proto.escrow.Info;
      reader.readMessage(value,proto.escrow.Info.deserializeBinaryFromReader);
      msg.setReceiver(value);
      break;
    case 4:
      var value = new proto.escrow.Logis;
      reader.readMessage(value,proto.escrow.Logis.deserializeBinaryFromReader);
      msg.setLogis(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.escrow.EscrowRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.escrow.EscrowRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.escrow.EscrowRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.escrow.EscrowRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getImpUid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSender();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.escrow.Info.serializeBinaryToWriter
    );
  }
  f = message.getReceiver();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.escrow.Info.serializeBinaryToWriter
    );
  }
  f = message.getLogis();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.escrow.Logis.serializeBinaryToWriter
    );
  }
};


/**
 * optional string imp_uid = 1;
 * @return {string}
 */
proto.escrow.EscrowRequest.prototype.getImpUid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.escrow.EscrowRequest} returns this
 */
proto.escrow.EscrowRequest.prototype.setImpUid = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional Info sender = 2;
 * @return {?proto.escrow.Info}
 */
proto.escrow.EscrowRequest.prototype.getSender = function() {
  return /** @type{?proto.escrow.Info} */ (
    jspb.Message.getWrapperField(this, proto.escrow.Info, 2));
};


/**
 * @param {?proto.escrow.Info|undefined} value
 * @return {!proto.escrow.EscrowRequest} returns this
*/
proto.escrow.EscrowRequest.prototype.setSender = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.escrow.EscrowRequest} returns this
 */
proto.escrow.EscrowRequest.prototype.clearSender = function() {
  return this.setSender(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.escrow.EscrowRequest.prototype.hasSender = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Info receiver = 3;
 * @return {?proto.escrow.Info}
 */
proto.escrow.EscrowRequest.prototype.getReceiver = function() {
  return /** @type{?proto.escrow.Info} */ (
    jspb.Message.getWrapperField(this, proto.escrow.Info, 3));
};


/**
 * @param {?proto.escrow.Info|undefined} value
 * @return {!proto.escrow.EscrowRequest} returns this
*/
proto.escrow.EscrowRequest.prototype.setReceiver = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.escrow.EscrowRequest} returns this
 */
proto.escrow.EscrowRequest.prototype.clearReceiver = function() {
  return this.setReceiver(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.escrow.EscrowRequest.prototype.hasReceiver = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional Logis logis = 4;
 * @return {?proto.escrow.Logis}
 */
proto.escrow.EscrowRequest.prototype.getLogis = function() {
  return /** @type{?proto.escrow.Logis} */ (
    jspb.Message.getWrapperField(this, proto.escrow.Logis, 4));
};


/**
 * @param {?proto.escrow.Logis|undefined} value
 * @return {!proto.escrow.EscrowRequest} returns this
*/
proto.escrow.EscrowRequest.prototype.setLogis = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.escrow.EscrowRequest} returns this
 */
proto.escrow.EscrowRequest.prototype.clearLogis = function() {
  return this.setLogis(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.escrow.EscrowRequest.prototype.hasLogis = function() {
  return jspb.Message.getField(this, 4) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.escrow.EscrowResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.escrow.EscrowResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.escrow.EscrowResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.escrow.EscrowResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    code: jspb.Message.getFieldWithDefault(msg, 1, 0),
    message: jspb.Message.getFieldWithDefault(msg, 2, ""),
    response: (f = msg.getResponse()) && proto.escrow.Logis.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.escrow.EscrowResponse}
 */
proto.escrow.EscrowResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.escrow.EscrowResponse;
  return proto.escrow.EscrowResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.escrow.EscrowResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.escrow.EscrowResponse}
 */
proto.escrow.EscrowResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCode(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setMessage(value);
      break;
    case 3:
      var value = new proto.escrow.Logis;
      reader.readMessage(value,proto.escrow.Logis.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.escrow.EscrowResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.escrow.EscrowResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.escrow.EscrowResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.escrow.EscrowResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCode();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getMessage();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.escrow.Logis.serializeBinaryToWriter
    );
  }
};


/**
 * optional int32 code = 1;
 * @return {number}
 */
proto.escrow.EscrowResponse.prototype.getCode = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.escrow.EscrowResponse} returns this
 */
proto.escrow.EscrowResponse.prototype.setCode = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string message = 2;
 * @return {string}
 */
proto.escrow.EscrowResponse.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.escrow.EscrowResponse} returns this
 */
proto.escrow.EscrowResponse.prototype.setMessage = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Logis response = 3;
 * @return {?proto.escrow.Logis}
 */
proto.escrow.EscrowResponse.prototype.getResponse = function() {
  return /** @type{?proto.escrow.Logis} */ (
    jspb.Message.getWrapperField(this, proto.escrow.Logis, 3));
};


/**
 * @param {?proto.escrow.Logis|undefined} value
 * @return {!proto.escrow.EscrowResponse} returns this
*/
proto.escrow.EscrowResponse.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.escrow.EscrowResponse} returns this
 */
proto.escrow.EscrowResponse.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.escrow.EscrowResponse.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 3) != null;
};


goog.object.extend(exports, proto.escrow);
