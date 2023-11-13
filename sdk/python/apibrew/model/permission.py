from datetime import datetime
from typing import List, Optional
from enum import Enum
from apibrew.entity import Entity, EntityInfo
from __future__ import annotations

from apibrew.model.user import User
from apibrew.model.role import Role



class AuditData:
    createdBy: str
    updatedBy: str
    createdOn: datetime
    updatedOn: datetime

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


class Operation(Enum):
    READ = "READ"
    CREATE = "CREATE"
    UPDATE = "UPDATE"
    DELETE = "DELETE"
    FULL = "FULL"


class Permit(Enum):
    ALLOW = "ALLOW"
    REJECT = "REJECT"



class Permission:
    id: str
    version: int
    auditData: AuditData
    namespace: str
    resource: str
    recordSelector: BooleanExpression
    operation: Operation
    before: datetime
    after: datetime
    user: User
    role: Role
    permit: Permit
    localFlags: dict

    @staticmethod
    def entity_info():
        return EntityInfo("system", "Permission", "system-permission")




