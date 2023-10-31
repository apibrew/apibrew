from datetime import datetime
from typing import List, Optional
from enum import Enum
from apibrew.entity import Entity, EntityInfo
from __future__ import annotations






class Record:
    id: str
    properties: dict
    packedProperties: list[dict]

    @staticmethod
    def entity_info():
        return EntityInfo("system", "Record", "system-record")




