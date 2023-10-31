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


class Backend(Enum):
    POSTGRESQL = "POSTGRESQL"
    MYSQL = "MYSQL"
    MONGODB = "MONGODB"
    REDIS = "REDIS"



class DataSource:
    id: str
    version: int
    auditData: AuditData
    name: str
    description: str
    backend: Backend
    options: dict[str, str]

    @staticmethod
    def entity_info():
        return EntityInfo("system", "DataSource", "system-datasource")




