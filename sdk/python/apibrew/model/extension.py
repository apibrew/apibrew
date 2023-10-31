from datetime import datetime
from typing import List, Optional
from enum import Enum
from apibrew.entity import Entity, EntityInfo
from __future__ import annotations

from apibrew.model.resource import Resource
from apibrew.model.record import Record



class BooleanExpression:
    and_: list[BooleanExpression]
    or_: list[BooleanExpression]
    not_: BooleanExpression
    equal: PairExpression
    lessThan: PairExpression
    greaterThan: PairExpression
    lessThanOrEqual: PairExpression
    greaterThanOrEqual: PairExpression
    in_: PairExpression
    isNull: Expression
    regexMatch: RegexMatchExpression

class PairExpression:
    left: Expression
    right: Expression

class RegexMatchExpression:
    pattern: str
    expression: Expression

class Expression:
    property: str
    value: dict

class AuditData:
    createdBy: str
    updatedBy: str
    createdOn: datetime
    updatedOn: datetime

class FunctionCall:
    host: str
    functionName: str

class HttpCall:
    uri: str
    method: str

class ChannelCall:
    channelKey: str

class ExternalCall:
    functionCall: FunctionCall
    httpCall: HttpCall
    channelCall: ChannelCall

class EventSelector:
    actions: list[Action]
    recordSelector: BooleanExpression
    namespaces: list[str]
    resources: list[str]
    ids: list[str]
    annotations: dict[str, str]

class RecordSearchParams:
    query: BooleanExpression
    limit: int
    offset: int
    resolveReferences: list[str]

class Event:
    id: str
    action: Action
    recordSearchParams: RecordSearchParams
    actionSummary: str
    actionDescription: str
    resource: Resource
    records: list[Record]
    finalizes: bool
    sync: bool
    time: datetime
    total: int
    actionName: str
    input: dict
    output: dict
    annotations: dict[str, str]
    error: Error

class ErrorField:
    recordId: str
    property: str
    message: str
    value: dict

class Error:
    code: Code
    message: str
    fields: list[ErrorField]


class Action(Enum):
    CREATE = "CREATE"
    UPDATE = "UPDATE"
    DELETE = "DELETE"
    GET = "GET"
    LIST = "LIST"
    OPERATE = "OPERATE"


class Code(Enum):
    UNKNOWN_ERROR = "UNKNOWN_ERROR"
    RECORD_NOT_FOUND = "RECORD_NOT_FOUND"
    UNABLE_TO_LOCATE_PRIMARY_KEY = "UNABLE_TO_LOCATE_PRIMARY_KEY"
    INTERNAL_ERROR = "INTERNAL_ERROR"
    PROPERTY_NOT_FOUND = "PROPERTY_NOT_FOUND"
    RECORD_VALIDATION_ERROR = "RECORD_VALIDATION_ERROR"
    RESOURCE_VALIDATION_ERROR = "RESOURCE_VALIDATION_ERROR"
    AUTHENTICATION_FAILED = "AUTHENTICATION_FAILED"
    ALREADY_EXISTS = "ALREADY_EXISTS"
    ACCESS_DENIED = "ACCESS_DENIED"
    BACKEND_ERROR = "BACKEND_ERROR"
    UNIQUE_VIOLATION = "UNIQUE_VIOLATION"
    REFERENCE_VIOLATION = "REFERENCE_VIOLATION"
    RESOURCE_NOT_FOUND = "RESOURCE_NOT_FOUND"
    UNSUPPORTED_OPERATION = "UNSUPPORTED_OPERATION"
    EXTERNAL_BACKEND_COMMUNICATION_ERROR = "EXTERNAL_BACKEND_COMMUNICATION_ERROR"
    EXTERNAL_BACKEND_ERROR = "EXTERNAL_BACKEND_ERROR"
    RATE_LIMIT_ERROR = "RATE_LIMIT_ERROR"



class Extension:
    id: str
    version: int
    auditData: AuditData
    name: str
    description: str
    selector: EventSelector
    order: int
    finalizes: bool
    sync: bool
    responds: bool
    call: ExternalCall
    annotations: dict[str, str]

    @staticmethod
    def entity_info():
        return EntityInfo("system", "Extension", "system-extension")




