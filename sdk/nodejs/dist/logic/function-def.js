"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineFunction = void 0;
var client_1 = require("../client");
var function_1 = require("../model/logic/function");
var module_def_1 = require("./module-def");
function defineFunction(name, args, fn) {
    var client = client_1.Client.getDefaultClient();
    var functionRepository = client.newRepository(function_1.FunctionResource);
    var module = (0, module_def_1.getModule)();
    functionRepository.apply({
        package: module.package,
        name: name,
        args: args.map(function (arg) {
            return {
                name: arg
            };
        }),
        module: {
            id: module.id,
        },
        engine: {
            name: 'nodejs-engine'
        }
    }).then(function (resp) {
        console.log(resp);
    }, function (err) {
        console.error(err);
    });
    (0, module_def_1.registerModuleChild)(name, fn);
}
exports.defineFunction = defineFunction;
