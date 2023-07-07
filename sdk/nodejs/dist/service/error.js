"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.handleError = void 0;
var axios_1 = require("axios");
function handleError(error) {
    var _a;
    if (error instanceof axios_1.AxiosError) {
        console.error('AxiosError:', error.request.url, error.request.body, error.code, (_a = error.response) === null || _a === void 0 ? void 0 : _a.data);
        return error;
    }
    else {
        return error;
    }
}
exports.handleError = handleError;
