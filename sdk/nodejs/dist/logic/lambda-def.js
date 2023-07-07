"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.fireLambda = exports.defineLambda = void 0;
var client_1 = require("../client");
var model_1 = require("../model");
var lambda_1 = require("../model/logic/lambda");
var module_def_1 = require("./module-def");
var error_1 = require("../service/error");
function parseLambdaEventSelectorPattern(eventSelectorPattern) {
    var parts = eventSelectorPattern.split(':');
    var resourceFullName = parts[0];
    var action = parts[1];
    var subParts = resourceFullName.split('/');
    var resourceName;
    var resourceNamespace;
    if (subParts.length === 1) {
        resourceName = subParts[0];
        resourceNamespace = 'default';
    }
    else if (subParts.length === 2) {
        resourceNamespace = subParts[0];
        resourceName = subParts[1];
    }
    else {
        throw new Error('Invalid resource name: ' + resourceFullName);
    }
    return {
        resourceFullName: resourceFullName,
        resourceName: resourceName,
        resourceNamespace: resourceNamespace,
        action: action
    };
}
function defineLambda(name, eventSelectorPattern, fn) {
    var client = client_1.Client.getDefaultClient();
    var functionRepository = client.newRepository(model_1.FunctionResource);
    var lambdaRepository = client.newRepository(lambda_1.LambdaResource);
    var module = (0, module_def_1.getModule)();
    function createLambda() {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        console.log('before create lambda');
                        return [4 /*yield*/, functionRepository.apply({
                                package: module.package,
                                name: 'Lambda_' + name,
                                args: [{
                                        name: 'element'
                                    }],
                                module: {
                                    id: module.id,
                                },
                                engine: {
                                    name: 'nodejs-engine'
                                }
                            }).then(function (resp) {
                                console.log(resp);
                            }, function (err) {
                                console.error((0, error_1.handleError)(err));
                            })];
                    case 1:
                        _a.sent();
                        console.log('after create lambda');
                        return [4 /*yield*/, lambdaRepository.apply({
                                package: module.package,
                                name: 'Lambda_' + name,
                                eventSelectorPattern: eventSelectorPattern,
                                function: {
                                    package: module.package,
                                    name: 'Lambda_' + name,
                                },
                            }).then(function (resp) {
                                console.log(resp);
                            }, function (err) {
                                console.error((0, error_1.handleError)(err));
                            })];
                    case 2:
                        _a.sent();
                        console.log('after create lambda 2');
                        return [2 /*return*/];
                }
            });
        });
    }
    createLambda();
    (0, module_def_1.registerModuleChild)('Lambda_' + name, fn);
}
exports.defineLambda = defineLambda;
function fireLambda(trigger, element) {
    var module = (0, module_def_1.getModule)();
    var client = client_1.Client.getDefaultClient();
    var parsed = parseLambdaEventSelectorPattern(trigger);
    var repository = client.newRepository({
        namespace: parsed.resourceNamespace,
        resource: parsed.resourceName,
    });
    element.action = parsed.action;
    repository.create(element).then(function (resp) {
        console.log('Lambda ' + trigger + ' fired');
    }, function (err) {
        console.error((0, error_1.handleError)(err));
    });
}
exports.fireLambda = fireLambda;
