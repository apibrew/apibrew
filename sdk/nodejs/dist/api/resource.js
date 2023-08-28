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
exports.apply = exports.save = exports.getByName = exports.get = exports.remove = exports.update = exports.create = exports.list = void 0;
var axios_1 = __importDefault(require("axios"));
var util_1 = require("../util");
function list(config) {
    return __awaiter(this, void 0, void 0, function () {
        var result;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, axios_1.default.get("".concat(config.backendUrl, "/resources"), {
                        headers: {
                            Authorization: "Bearer ".concat(config.token)
                        }
                    })];
                case 1:
                    result = _a.sent();
                    return [2 /*return*/, result.data];
            }
        });
    });
}
exports.list = list;
function create(config, resource) {
    return __awaiter(this, void 0, void 0, function () {
        var result;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, axios_1.default.post("".concat(config.backendUrl, "/resources"), resource, {
                        headers: {
                            Authorization: "Bearer ".concat(config.token)
                        }
                    })];
                case 1:
                    result = _a.sent();
                    return [2 /*return*/, result.data];
            }
        });
    });
}
exports.create = create;
function update(config, resource) {
    return __awaiter(this, void 0, void 0, function () {
        var result;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, axios_1.default.put("".concat(config.backendUrl, "/resources/").concat(resource.id), resource, {
                        headers: {
                            Authorization: "Bearer ".concat(config.token)
                        }
                    })];
                case 1:
                    result = _a.sent();
                    return [2 /*return*/, result.data];
            }
        });
    });
}
exports.update = update;
function remove(config, resource) {
    return __awaiter(this, void 0, void 0, function () {
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, axios_1.default.delete("".concat(config.backendUrl, "/resources/").concat(resource.id), {
                        headers: {
                            Authorization: "Bearer ".concat(config.token)
                        }
                    })];
                case 1:
                    _a.sent();
                    return [2 /*return*/];
            }
        });
    });
}
exports.remove = remove;
function get(config, resourceId) {
    return __awaiter(this, void 0, void 0, function () {
        var result;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, axios_1.default.get("".concat(config.backendUrl, "/resources/").concat(resourceId), {
                        headers: {
                            Authorization: "Bearer ".concat(config.token)
                        }
                    })];
                case 1:
                    result = _a.sent();
                    return [2 /*return*/, result.data];
            }
        });
    });
}
exports.get = get;
function getByName(config, resourceName, namespace) {
    return __awaiter(this, void 0, void 0, function () {
        var result;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!namespace) {
                        namespace = 'default';
                    }
                    return [4 /*yield*/, axios_1.default.get("".concat(config.backendUrl, "/resources/by-name/").concat(namespace, "/").concat(resourceName), {
                            headers: {
                                Authorization: "Bearer ".concat(config.token)
                            }
                        })];
                case 1:
                    result = _a.sent();
                    return [2 /*return*/, result.data];
            }
        });
    });
}
exports.getByName = getByName;
function save(config, resource) {
    return __awaiter(this, void 0, void 0, function () {
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!resource.id) return [3 /*break*/, 2];
                    return [4 /*yield*/, update(config, resource)];
                case 1: return [2 /*return*/, _a.sent()];
                case 2: return [4 /*yield*/, create(config, resource)];
                case 3: return [2 /*return*/, _a.sent()];
            }
        });
    });
}
exports.save = save;
function apply(config, resource) {
    return __awaiter(this, void 0, void 0, function () {
        var existingResource, e_1;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    _a.trys.push([0, 3, , 5]);
                    return [4 /*yield*/, getByName(config, resource.name, resource.namespace.name)];
                case 1:
                    existingResource = _a.sent();
                    resource.id = existingResource.id;
                    if (!(0, util_1.isObjectModified)(existingResource, resource)) {
                        return [2 /*return*/, existingResource];
                    }
                    console.log('Updating resource', resource.name);
                    return [4 /*yield*/, update(config, resource)];
                case 2: return [2 /*return*/, _a.sent()];
                case 3:
                    e_1 = _a.sent();
                    console.log('Creating resource', resource.name);
                    return [4 /*yield*/, create(config, resource)];
                case 4: return [2 /*return*/, _a.sent()];
                case 5: return [2 /*return*/];
            }
        });
    });
}
exports.apply = apply;
