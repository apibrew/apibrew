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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.RepositoryExtensionImpl = exports.RepositoryImpl = exports.DhClient = void 0;
var axios_1 = __importDefault(require("axios"));
var DhClient = /** @class */ (function () {
    function DhClient(params) {
        this.params = params;
    }
    DhClient.prototype.authenticateWithUsernameAndPassword = function (username, password) {
        return __awaiter(this, void 0, void 0, function () {
            var authRequest, result;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        authRequest = {
                            username: username,
                            password: password,
                            term: "LONG"
                        };
                        return [4 /*yield*/, axios_1.default.post("http://".concat(this.params.Addr, "/authentication/token"), authRequest)];
                    case 1:
                        result = _a.sent();
                        this.params.token = result.data.token.content;
                        return [2 /*return*/];
                }
            });
        });
    };
    DhClient.prototype.newRepository = function (namespace, resource) {
        return new RepositoryImpl(this, {
            namespace: namespace,
            resource: resource,
            updateCheckVersion: false,
        });
    };
    DhClient.prototype.NewExtensionService = function (host, port) {
        return new ExtensionServiceImpl(host, port, host + ':' + port, this);
    };
    return DhClient;
}());
exports.DhClient = DhClient;
var ExtensionServiceImpl = /** @class */ (function () {
    function ExtensionServiceImpl(host, port, remoteHost, client) {
        this.host = host;
        this.port = port;
        this.remoteHost = remoteHost;
        this.client = client;
        this.functions = {};
    }
    ExtensionServiceImpl.prototype.getRemoteHost = function () {
        return this.remoteHost;
    };
    ExtensionServiceImpl.prototype.registerFunction = function (name, handler) {
        this.functions[name] = handler;
    };
    ExtensionServiceImpl.prototype.run = function () {
        return __awaiter(this, void 0, void 0, function () {
            var express, app;
            var _this = this;
            return __generator(this, function (_a) {
                express = require('express');
                app = express();
                app.use(express.json());
                app.get('/', function (req, res) {
                    res.send('ok');
                });
                app.post('/:name', function (req, res) { return __awaiter(_this, void 0, void 0, function () {
                    var name, request, response, e_1;
                    return __generator(this, function (_a) {
                        switch (_a.label) {
                            case 0:
                                name = req.params.name;
                                request = {
                                    name: name,
                                    request: req.body.content
                                };
                                _a.label = 1;
                            case 1:
                                _a.trys.push([1, 3, , 4]);
                                return [4 /*yield*/, this.functions[name](request.request)];
                            case 2:
                                response = _a.sent();
                                res.send({
                                    content: response
                                });
                                return [3 /*break*/, 4];
                            case 3:
                                e_1 = _a.sent();
                                console.log(e_1);
                                res.status(400).send({
                                    message: e_1.message
                                });
                                return [3 /*break*/, 4];
                            case 4: return [2 /*return*/];
                        }
                    });
                }); });
                console.log('starting extension service');
                app.listen(this.port, this.host, function () {
                    console.log("External service is listening on ".concat(_this.host));
                });
                return [2 /*return*/];
            });
        });
    };
    return ExtensionServiceImpl;
}());
var RepositoryImpl = /** @class */ (function () {
    function RepositoryImpl(client, params) {
        this.client = client;
        this.params = params;
    }
    RepositoryImpl.prototype.loadResources = function () {
        return __awaiter(this, void 0, void 0, function () {
            var result;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        if (this.resource) {
                            return [2 /*return*/];
                        }
                        return [4 /*yield*/, axios_1.default.get("http://".concat(this.client.params.Addr, "/system/resources/").concat(this.params.namespace, "/").concat(this.params.resource))];
                    case 1:
                        result = _a.sent();
                        this.resource = result.data.resource;
                        return [2 /*return*/];
                }
            });
        });
    };
    RepositoryImpl.prototype.create = function (entity) {
        return __awaiter(this, void 0, void 0, function () {
            var result;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, axios_1.default.post("http://".concat(this.client.params.Addr, "/records/").concat(this.params.namespace, "/").concat(this.params.resource), entity, {
                            headers: {
                                Authorization: "Bearer ".concat(this.client.params.token)
                            }
                        })];
                    case 1:
                        result = _a.sent();
                        return [2 /*return*/, result.data];
                }
            });
        });
    };
    RepositoryImpl.prototype.update = function (entity) {
        return __awaiter(this, void 0, void 0, function () {
            var result;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, axios_1.default.put("http://".concat(this.client.params.Addr, "/records/").concat(this.params.namespace, "/").concat(this.params.resource, "/").concat(entity.id), entity, {
                            headers: {
                                Authorization: "Bearer ".concat(this.client.params.token)
                            }
                        })];
                    case 1:
                        result = _a.sent();
                        return [2 /*return*/, result.data];
                }
            });
        });
    };
    RepositoryImpl.prototype.get = function (id) {
        return __awaiter(this, void 0, void 0, function () {
            var result;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, axios_1.default.get("http://".concat(this.client.params.Addr, "/records/").concat(this.params.namespace, "/").concat(this.params.resource, "/").concat(id), {
                            headers: {
                                Authorization: "Bearer ".concat(this.client.params.token)
                            }
                        })];
                    case 1:
                        result = _a.sent();
                        return [2 /*return*/, result.data];
                }
            });
        });
    };
    RepositoryImpl.prototype.load = function (entity) {
        return __awaiter(this, void 0, void 0, function () {
            var _i, _a, prop, val, result;
            return __generator(this, function (_b) {
                switch (_b.label) {
                    case 0:
                        if (!entity.id) return [3 /*break*/, 1];
                        return [2 /*return*/, this.get(entity.id)];
                    case 1: return [4 /*yield*/, this.loadResources()];
                    case 2:
                        _b.sent();
                        _i = 0, _a = this.resource.properties;
                        _b.label = 3;
                    case 3:
                        if (!(_i < _a.length)) return [3 /*break*/, 6];
                        prop = _a[_i];
                        if (!prop.unique) return [3 /*break*/, 5];
                        val = entity[prop.name];
                        return [4 /*yield*/, axios_1.default.get("http://".concat(this.client.params.Addr, "/").concat(this.params.namespace, "/").concat(this.params.resource, "?filters=").concat(prop.name, "&filters=").concat(val, "&limit=1"), {
                                headers: {
                                    Authorization: "Bearer ".concat(this.client.params.token)
                                }
                            })];
                    case 4:
                        result = _b.sent();
                        if (!result.data.total) {
                            return [3 /*break*/, 5];
                        }
                        return [2 /*return*/, result.data.content[0].properties];
                    case 5:
                        _i++;
                        return [3 /*break*/, 3];
                    case 6: throw new Error("Entity not found: ".concat(this.params.namespace, "/").concat(this.params.resource));
                }
            });
        });
    };
    RepositoryImpl.prototype.apply = function (entity) {
        return __awaiter(this, void 0, void 0, function () {
            var result;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, axios_1.default.patch("http://".concat(this.client.params.Addr, "/").concat(this.params.namespace, "/").concat(this.params.resource), entity, {
                            headers: {
                                Authorization: "Bearer ".concat(this.client.params.token)
                            }
                        })];
                    case 1:
                        result = _a.sent();
                        return [2 /*return*/, result.data];
                }
            });
        });
    };
    RepositoryImpl.prototype.find = function (params) {
        return __awaiter(this, void 0, void 0, function () {
            var result;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        if (!params.resolveReferences) {
                            params.resolveReferences = ["*"];
                        }
                        return [4 /*yield*/, axios_1.default.get("http://".concat(this.client.params.Addr, "/records/").concat(this.params.namespace, "/").concat(this.params.resource), {
                                headers: {
                                    Authorization: "Bearer ".concat(this.client.params.token)
                                }
                            })];
                    case 1:
                        result = _a.sent();
                        return [2 /*return*/, result.data];
                }
            });
        });
    };
    RepositoryImpl.prototype.extend = function (extensionService) {
        return new RepositoryExtensionImpl(this, extensionService, this.params.resource, this.params.namespace, this.client);
    };
    return RepositoryImpl;
}());
exports.RepositoryImpl = RepositoryImpl;
// ## repository extension
var RepositoryExtensionImpl = /** @class */ (function () {
    function RepositoryExtensionImpl(repository, extension, resourceName, namespace, client) {
        this.repository = repository;
        this.extension = extension;
        this.resourceName = resourceName;
        this.namespace = namespace;
        this.client = client;
        this.extensionRepository = new RepositoryImpl(client, {
            namespace: "system", resource: "extension", updateCheckVersion: false
        });
    }
    RepositoryExtensionImpl.prototype.onCreate = function (handler, finalize) {
        return __awaiter(this, void 0, void 0, function () {
            var extensionName, ext;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        extensionName = this.getExtensionName("OnCreate");
                        this.extension.registerFunction(extensionName, function (data) {
                            return __awaiter(this, void 0, void 0, function () {
                                var records, _i, _a, record, entity, response;
                                return __generator(this, function (_b) {
                                    switch (_b.label) {
                                        case 0:
                                            records = [];
                                            _i = 0, _a = data.request.records;
                                            _b.label = 1;
                                        case 1:
                                            if (!(_i < _a.length)) return [3 /*break*/, 4];
                                            record = _a[_i];
                                            return [4 /*yield*/, handler(record.properties)];
                                        case 2:
                                            entity = _b.sent();
                                            records.push({
                                                properties: entity
                                            });
                                            _b.label = 3;
                                        case 3:
                                            _i++;
                                            return [3 /*break*/, 1];
                                        case 4:
                                            response = {
                                                "response": {
                                                    '@type': 'type.googleapis.com/stub.CreateRecordResponse',
                                                    "records": records
                                                }
                                            };
                                            return [2 /*return*/, response];
                                    }
                                });
                            });
                        });
                        ext = {
                            name: extensionName,
                            namespace: this.namespace,
                            resource: this.resourceName,
                            instead: {
                                create: {
                                    kind: "httpCall",
                                    uri: "http://".concat(this.extension.getRemoteHost(), "/").concat(extensionName),
                                    method: 'POST',
                                },
                                finalize: finalize,
                            },
                        };
                        return [4 /*yield*/, this.extensionRepository.apply(ext)];
                    case 1:
                        _a.sent();
                        return [2 /*return*/];
                }
            });
        });
    };
    RepositoryExtensionImpl.prototype.onUpdate = function (handler, finalize) {
        return __awaiter(this, void 0, void 0, function () {
            var extensionName, ext;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        extensionName = this.getExtensionName("OnUpdate");
                        this.extension.registerFunction(extensionName, function (data) {
                            return __awaiter(this, void 0, void 0, function () {
                                var records, _i, _a, record, entity, response;
                                return __generator(this, function (_b) {
                                    switch (_b.label) {
                                        case 0:
                                            records = [];
                                            _i = 0, _a = data.request.records;
                                            _b.label = 1;
                                        case 1:
                                            if (!(_i < _a.length)) return [3 /*break*/, 4];
                                            record = _a[_i];
                                            return [4 /*yield*/, handler(record.properties)];
                                        case 2:
                                            entity = _b.sent();
                                            records.push({
                                                properties: entity
                                            });
                                            _b.label = 3;
                                        case 3:
                                            _i++;
                                            return [3 /*break*/, 1];
                                        case 4:
                                            response = {
                                                "response": {
                                                    '@type': 'type.googleapis.com/stub.UpdateRecordResponse',
                                                    "records": records
                                                }
                                            };
                                            return [2 /*return*/, response];
                                    }
                                });
                            });
                        });
                        ext = {
                            name: extensionName,
                            namespace: this.namespace,
                            resource: this.resourceName,
                            instead: {
                                update: {
                                    kind: "httpCall",
                                    uri: "http://".concat(this.extension.getRemoteHost(), "/").concat(extensionName),
                                    method: 'POST',
                                },
                                finalize: finalize,
                            },
                        };
                        return [4 /*yield*/, this.extensionRepository.apply(ext)];
                    case 1:
                        _a.sent();
                        return [2 /*return*/];
                }
            });
        });
    };
    RepositoryExtensionImpl.prototype.onDelete = function (handler) {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                //TODO implement me
                throw new Error("Method not implemented.");
            });
        });
    };
    RepositoryExtensionImpl.prototype.onGet = function (handler) {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                //TODO implement me
                throw new Error("Method not implemented.");
            });
        });
    };
    RepositoryExtensionImpl.prototype.onList = function (handler) {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                //TODO implement me
                throw new Error("Method not implemented.");
            });
        });
    };
    RepositoryExtensionImpl.prototype.getExtensionName = function (action) {
        return "".concat(this.namespace, "-").concat(this.resourceName, "-").concat(action);
    };
    return RepositoryExtensionImpl;
}());
exports.RepositoryExtensionImpl = RepositoryExtensionImpl;
