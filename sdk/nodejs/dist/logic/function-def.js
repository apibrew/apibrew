"use strict";
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineFunction = void 0;
var client_1 = require("../client");
var function_1 = require("../model/logic/function");
var module_def_1 = require("./module-def");
function defineFunction(funcProps, fn) {
    var client = client_1.Client.getDefaultClient();
    var functionRepository = client.newRepository(function_1.FunctionResource);
    functionRepository.apply(__assign(__assign({}, funcProps), { module: {
            id: (0, module_def_1.getModuleId)()
        }, engine: {
            name: 'nodejs-engine'
        } })).then(function (resp) {
        console.log(resp);
    }, function (err) {
        console.error(err);
    });
    (0, module_def_1.registerModuleChild)(funcProps.name, fn);
}
exports.defineFunction = defineFunction;
