"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SpecialProperty = exports.isAnnotationEnabled = void 0;
function isAnnotationEnabled(annotations, annotationName) {
    return Boolean(annotations && annotations[annotationName] === 'true');
}
exports.isAnnotationEnabled = isAnnotationEnabled;
exports.SpecialProperty = 'SpecialProperty';
