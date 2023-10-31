from datetime import datetime
from typing import List, Optional
from enum import Enum
from apibrew.entity import Entity, EntityInfo
from __future__ import annotations

from apibrew.model.role import Role
from apibrew.model.permission import Permission



class AuditData:
    createdBy: str
    updatedBy: str
    createdOn: datetime
    updatedOn: datetime



class User:
    id: str
    version: int
    auditData: AuditData
    username: str
    password: str
    roles: list[Role]
    permissions: list[Permission]
    details: dict

    @staticmethod
    def entity_info():
        return EntityInfo("system", "User", "system-user")




