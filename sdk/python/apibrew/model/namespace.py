from datetime import datetime
from typing import List, Optional
from enum import Enum
from apibrew.entity import Entity, EntityInfo
from __future__ import annotations




class AuditData:
    createdBy: str
    updatedBy: str
    createdOn: datetime
    updatedOn: datetime



class Namespace:
    id: str
    version: int
    auditData: AuditData
    name: str
    description: str
    details: dict

    @staticmethod
    def entity_info():
        return EntityInfo("system", "Namespace", "system-namespace")




