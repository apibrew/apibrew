from datetime import datetime
from typing import List, Optional
from enum import Enum
from apibrew.entity import Entity, EntityInfo
from __future__ import annotations

from apibrew.storage.model.storage_object import StorageObject




class Permission(Enum):
    DOWNLOAD = "DOWNLOAD"
    UPLOAD = "UPLOAD"



class Signature:
    id: str
    object: StorageObject
    permissions: list[Permission]
    expiration: datetime
    signature: str
    annotations: dict[str, str]
    version: int

    @staticmethod
    def entity_info():
        return EntityInfo("storage", "Signature", "storage-signature")




