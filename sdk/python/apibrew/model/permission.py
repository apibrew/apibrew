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


class PropertyMode(Enum):
    PROPERTY_MATCH_ONLY = "PROPERTY_MATCH_ONLY"
    PROPERTY_MATCH_ANY = "PROPERTY_MATCH_ANY"


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
    property: str
    propertyValue: str
    propertyMode: PropertyMode
    operation: Operation
    recordIds: list[str]
    before: datetime
    after: datetime
    user: User
    role: Role
    permit: Permit
    localFlags: dict

    @staticmethod
    def entity_info():
        return EntityInfo("system", "Permission", "system-permission")




