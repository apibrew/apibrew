from datetime import datetime
from typing import List, Optional
from enum import Enum
from apibrew.entity import Entity, EntityInfo
from __future__ import annotations

from apibrew.model.namespace import Namespace
from apibrew.model.data_source import DataSource



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

class SubType:
    name: str
    title: str
    description: str
    properties: list[Property]

class AuditData:
    createdBy: str
    updatedBy: str
    createdOn: datetime
    updatedOn: datetime

class IndexProperty:
    name: str
    order: Order

class Index:
    properties: list[IndexProperty]
    indexType: IndexType
    unique: bool
    annotations: dict[str, str]

class Reference:
    resource: Resource
    cascade: bool
    backReference: str


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


class Order(Enum):
    UNKNOWN = "UNKNOWN"
    ASC = "ASC"
    DESC = "DESC"


class IndexType(Enum):
    BTREE = "BTREE"
    HASH = "HASH"



class Resource:
    id: str
    version: int
    auditData: AuditData
    name: str
    namespace: Namespace
    virtual: bool
    properties: list[Property]
    indexes: list[Index]
    types: list[SubType]
    immutable: bool
    abstract: bool
    checkReferences: bool
    dataSource: DataSource
    entity: str
    catalog: str
    title: str
    description: str
    annotations: dict[str, str]

    @staticmethod
    def entity_info():
        return EntityInfo("system", "Resource", "resources")




