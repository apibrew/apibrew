"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isObjectModified = exports.isObject = void 0;
function isObject(item) {
    return (item && typeof item === 'object' && !Array.isArray(item));
}
exports.isObject = isObject;
function isObjectModified(source, updated) {
    if (JSON.stringify(source) === JSON.stringify(updated)) {
        return false;
    }
    if (isObject(source) && isObject(updated)) {
        for (var _i = 0, _a = Object.keys(updated); _i < _a.length; _i++) {
            var key = _a[_i];
            console.log('looking for diff: ', key);
            if (isObjectModified(source[key], updated[key])) {
                console.log('Diff found: ', source[key], updated[key]);
                return true;
            }
        }
        return false;
    }
    else if (Array.isArray(source) && Array.isArray(updated)) {
        if (source.length !== updated.length) {
            console.log('Diff found: ', source.length !== updated.length);
            return true;
        }
        for (var i = 0; i < source.length; i++) {
            if (isObjectModified(source[i], updated[i])) {
                console.log('Diff found: ', source[i], updated[i]);
                return true;
            }
        }
        return false;
    }
    else {
        console.log('Diff found{e}: ', source, updated);
        return true;
    }
}
exports.isObjectModified = isObjectModified;
