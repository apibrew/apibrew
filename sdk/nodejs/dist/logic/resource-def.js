"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineResource = void 0;
var client_1 = require("../client");
var service_1 = require("../service");
function defineResource(resource) {
    var client = client_1.Client.getDefaultClient();
    service_1.ResourceService.create(client.provider()(), resource).then(function (resp) {
        console.log(resp);
    }, function (err) {
        console.error(err);
    });
}
exports.defineResource = defineResource;
