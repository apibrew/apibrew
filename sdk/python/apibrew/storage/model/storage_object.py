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



class StorageObject:
    id: str
    name: str
    annotations: dict[str, str]
    contentType: str
    size: int
    allowDownloadPublicly: bool
    allowUploadPublicly: bool
    version: int
    auditData: AuditData

    @staticmethod
    def entity_info():
        return EntityInfo("storage", "StorageObject", "storage-storageobject")




