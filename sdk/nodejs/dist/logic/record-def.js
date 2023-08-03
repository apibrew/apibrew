"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineRecord = void 0;
var client_1 = require("../client");
var error_1 = require("../api/error");
function defineRecord(resourceInfo, record) {
    var client = client_1.Client.getDefaultClient();
    var repository = client.newRepository(resourceInfo);
    repository.create(record).then(function (resp) {
        console.log(resp);
    }, function (err) {
        console.error((0, error_1.handleError)(err));
    });
}
exports.defineRecord = defineRecord;
