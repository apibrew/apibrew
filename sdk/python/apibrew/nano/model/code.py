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


class Language(Enum):
    PYTHON = "PYTHON"
    JAVASCRIPT = "JAVASCRIPT"


class ContentFormat(Enum):
    TEXT = "TEXT"
    TAR = "TAR"
    TAR_GZ = "TAR_GZ"



class Code:
    id: str
    name: str
    language: Language
    content: str
    contentFormat: ContentFormat
    annotations: dict[str, str]
    version: int
    auditData: AuditData

    @staticmethod
    def entity_info():
        return EntityInfo("nano", "Code", "nano-code")




