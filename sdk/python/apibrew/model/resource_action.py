from datetime import datetime
from typing import List, Optional
from enum import Enum
from apibrew.entity import Entity, EntityInfo
from __future__ import annotations

from apibrew.model.resource import Resource



class SubType:
    name: str
    title: str
    description: str
    properties: list[Property]

class Reference:
    resource: Resource
    cascade: bool
    backReference: str

class Property:
    name: str
    type: Type
    typeRef: str
    primary: bool
    required: bool
    unique: bool
    immutable: bool
    length: int
    item: Property
    reference: Reference
    defaultValue: dict
    enumValues: list[str]
    exampleValue: dict
    title: str
    description: str
    annotations: dict[str, str]

class AuditData:
    createdBy: str
    updatedBy: str
    createdOn: datetime
    updatedOn: datetime


class Type(Enum):
    BOOL = "BOOL"
    STRING = "STRING"
    FLOAT32 = "FLOAT32"
    FLOAT64 = "FLOAT64"
    INT32 = "INT32"
    INT64 = "INT64"
    BYTES = "BYTES"
    UUID = "UUID"
    DATE = "DATE"
    TIME = "TIME"
    TIMESTAMP = "TIMESTAMP"
    OBJECT = "OBJECT"
    MAP = "MAP"
    LIST = "LIST"
    REFERENCE = "REFERENCE"
    ENUM = "ENUM"
    STRUCT = "STRUCT"



class ResourceAction:
    id: str
    version: int
    auditData: AuditData
    resource: Resource
    name: str
    title: str
    description: str
    internal: bool
    types: list[SubType]
    input: list[Property]
    output: Property
    annotations: dict[str, str]

    @staticmethod
    def entity_info():
        return EntityInfo("system", "ResourceAction", "system-resourceaction")




