"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineResource = void 0;
var client_1 = require("../client");
var service_1 = require("../service");
var error_1 = require("../service/error");
function defineResource(resource) {
    var client = client_1.Client.getDefaultClient();
    service_1.ResourceService.apply(client.provider()(), resource).then(function (resp) {
        console.log(resp);
    }, function (err) {
        console.error((0, error_1.handleError)(err));
    });
}
exports.defineResource = defineResource;
