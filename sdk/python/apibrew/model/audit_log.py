from datetime import datetime
from typing import List, Optional
from enum import Enum
from apibrew.entity import Entity, EntityInfo
from __future__ import annotations





class Operation(Enum):
    CREATE = "CREATE"
    UPDATE = "UPDATE"
    DELETE = "DELETE"



class AuditLog:
    id: str
    version: int
    namespace: str
    resource: str
    recordId: str
    time: datetime
    username: str
    operation: Operation
    properties: dict
    annotations: dict[str, str]

    @staticmethod
    def entity_info():
        return EntityInfo("system", "AuditLog", "system-auditlog")




