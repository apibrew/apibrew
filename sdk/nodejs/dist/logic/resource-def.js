"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineResource = void 0;
var client_1 = require("../client");
var api_1 = require("../api");
var error_1 = require("../api/error");
function defineResource(resource) {
    var client = client_1.Client.getDefaultClient();
    api_1.ResourceApi.apply(client.provider()(), resource).then(function (resp) {
        console.log(resp);
    }, function (err) {
        console.error((0, error_1.handleError)(err));
    });
}
exports.defineResource = defineResource;
