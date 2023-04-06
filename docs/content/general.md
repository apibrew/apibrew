# Table of contents

* Core elements - Inside Data Handler there are 6 main elements
    * [Resource](#resource)
    * [Record](#record)
    * [Data Source](#data-source)
    * [Namespace](#namespace)
    * [User](#user)
    * [Extension](#extension)
* [Property Types](#property-types)
* [Security Context](#security-context)
* [Resource Reference](#resource-reference)
* [Annotations](#annotations)

# Core elements

![](/dh_elements.png)

## Resource

### Overview

Resource is main element among all elements. Resource is for defining schema of your structure. After that you can do
crud operation inside resource.
When you crete new resource, it is defining same structure on its backend.

Depending on data source backend, it can store and manage data differently.

But for all data sources and for different backends (sql, mongo, redis, etc.). Everything is working as same.
**It means that, if you use postgresql and moved your resources from postgresql to mongodb everything will work as is.
**. It is internal logic of data-handler how it is handling operation on which backend. At high level, end user expects
that all data source backends are the same.

### Resource definition

Proto file: [resource.proto](https://github.com/tislib/data-handler/blob/master/proto/model/resource.proto)
Resource has the following properties:

* **id** - unique resource id
* **name** - unique resource name
* **namespace** - each resource is kept inside a namespace. One namespace can have multiple resources
* **sourceConfig** - source config is to configure resource and bind it to data-source and an entity inside data source.
  An entity is like a table on sql databases or collection on mongodb etc.
    * **dataSource** - data source name: where resource structure and its data will be physically exists. Data source
      name is required if resource is not virtual
    * **catalog** - catalog is like a folder/schema/database. It is changing from backend to backend. Basically it is
      for grouping entities
    * **entity** - entity name an item on datasource backend where resource will be bound. For sql databases it is table
      name, for mongo it is collection name, etc.
* **properties** - list of properties of resource. This properties will be used by records of resource. Properties is
  columns on sql databases. For schemaless data structures properties is only managed by
    * **name** - property name.
    * **type** - type of property - see [property-types](#property-types) section
    * **mapping** - mapping is like a column name, it is binding name to entity. For abstraction purposes property name
      is not used while communicating to resource backend. Instead mapping is used as a key of property
    * **required** - this is to mark property as required
    * **primary** - this is to mark property as primary. Primary properties is like a part of primary key. Primary
      property(s) is used in to identify record.
    * **length** - length property is only valid and required for String typed properties
    * **immutable** - immutable is to mark property as immutable. If marked, updates on this field on records will be
      discarded
    * **securityContext** - security context is to apply ACL to resource property -
      see [security context](#security-context)
    * **defaultValue** - defaultValue is default value.
    * **exampleValue** - exampleValue is example value. It is an informative column
    * **enumValues** - enumValues is used if property type is an enum
    * **reference** - reference property is only valid and required for Reference types.
      See [Resource Reference](#resource-reference)
        * **referencedResource** - referenced resource name
        * **cascade** - if cascade is true, delete/update operations will be cascaded to back referenced resources
    * **subType** - subType is used for complex types(list, map). For list, subType is element type. For map, it is
      value type(key type is always string)
    * **title** - property type. It is an informative column
    * **description** - property type. It is an informative column
    * **annotations** - property annotations - see [annotations](#annotations)
* **indexes** - list of resource indexes. Its implementation is depending on data source backend and may not be
  supported by
  some backends.
    * **properties** - list of properties inside single index. Normally you will need only single property. Multi
      property will be needed for multi property indexes(for complex indexes)
        * **name** - property name
        * **order** - for ordered indexes(like BTREE), order of data
    * **indexType** - Index type(BTREE, HASH)
    * **annotations** - Index annotations - see [annotations](#annotations)
* **securityContext** - security context is to apply ACL to resource property -
  see [security context](#security-context)
* **virtual** - If virtual is true. Operations will not phisically affect datasource/backend. Virtual resources is for
  extension purposes. Their behaviors can be extended and altered. It can also be used to integrate 3rd party systems.
* **immutable** - If immutable is true. Update and Delete operation will not be allowed on records of that resource
* **annotations** - resource annotations - see [annotations](#annotations)

### Special properties

When new resource is created or updated, *data-handler* appends some special properties to the resource.
These are:

* id - id field is primary key and record identifier. id special property will be added if resource does not any primary
  key. You can also prevent this happening by annotating resource with `DoPrimaryKeyLookup` annotation
* audit - audit special properties are for audit purposes and consist of 4 properties. By default, audit special
  properties will not be enabled. And it can be enabled by annotating resource with `EnableAudit` annotation.
    * created_on - if audit is enabled, this property will hold information about username who created record
    * created_by - if audit is enabled, this property will hold information when record created
    * updated_on - if audit is enabled, this property will hold information about username who updated record last time
    * updated_by - if audit is enabled, this property will hold information when record updated last time
* version - version property is added if you don't have such property and resource *is not* annotated with
  DisableVersion annotation

### Examples

#### City, Country

country.yml

```yaml
type: resource
name: country
sourceConfig:
  dataSource: default
  entity: country
properties:
  - name: name
    type: STRING
    length: 255
    required: true
    unique: true
  - name: description
    type: STRING
    length: 255
```

Now let's create country resource

```
dhctl apply -f country.yml
```

city.yml

```yaml
type: resource
name: city
sourceConfig:
  dataSource: default
  entity: city
properties:
  - name: name
    type: STRING
    length: 255
    required: true
    unique: true
  - name: description
    type: STRING
    length: 255
  - name: country
    type: REFERENCE
    length: 255
    reference:
      referencedResource: country
      cascade: true
```

Now let's create city resource

```
dhctl apply -f city.yml
```

So, by this way, you will create both country and city resources

## Record

### Overview

Record is a peace of data of resource. In basic words, If Resource is a table, Record is a row; If Resource is a
collection, Record is an item, etc.

### Record definition

Proto file: [resource.proto](https://github.com/tislib/data-handler/blob/master/proto/model/record.proto)

Record has the following properties:

* **id** - unique record id, it is for identifying record. Record id comes from its resource property which is primary.
  If Resource has multiple primary properties, system will join them with dash. If no primary, id will be empty string
* **properties** - Properties is a map. Where key is property name and value is its value according to record. Mostly
  Properties is like a record body. It can be even considered record itself.
* **propertiesPacked** - This property is only available to GRPC and will be used instead of properties if pack mode is
  enabled. If pack mode enabled, properties will not be sent, instead propertiesPacked will be sent. It is for saving
  space and cpu for transferring many accounts.

### Examples

#### City, Country

data.yml

```yaml
type: record
resource: country
properties:
  name: Azerbaijan
  description: Land of fire
---
type: record
resource: city
properties:
  name: Baku
  country:
    name: Azerbaijan # This is for matching country by name
```

Now let's create country resource

```
dhctl apply -f country.yml
```

## Data source

### Overview

Data source is for connecting our Resources to Databases/Data stores. Data source is main part of resource. And without
Datasource, Resource cannot physically store any data without datasource.

Data source is also an abstraction point for various databases.

### Data source definition

Proto file: [resource.proto](https://github.com/tislib/data-handler/blob/master/proto/model/data-source.proto)

Data source has the following properties:

* **id** - unique data source id
* **name** - unique data source name
* **description** - Datasource description
* **backend** - Data source backend is an enum. And you define which database you will use by setting backend. Backend
  has following values
    * POSTGRESQL - postgresql database
    * VIRTUAL - virtual backend is not a real backend. It is not storing any data inside it. It is just for extension purposes. Which resource will be extended, they can be configured to virtual backend.
    * MYSQL
    * ORACLE
    * MONGODB
    * REDIS
* **params** - params is to configure database backend connection configuration params. For each backend type we have
  different available params.
    * **postgresqlParams** - Postgresql connection params
        * **username** - username
        * **password** - password
        * **host** - host
        * **port** - port
        * **dbName** - Database name
        * **defaultSchema** - defaultSchema
  * **mysqlParams** - Mysql connection params
      * **username** - username
      * **password** - password
      * **host** - host
      * **port** - port
      * **dbName** - Database name
      * **defaultSchema** - defaultSchema
  * **mongoParams** - Mongo connection params
      * **uri** - Mongodb connection string
      * **dbName** - Database name

## Namespace
### Overview
Namespace is for grouping resources.

### Namespace definition
* **id** - unique namespace id
* **name** - unique namespace name
* **description** - Datasource description

### Special namespaces:
1. Default namespace (name=default): Default namespace is auto create on initial setup. When new resource is created, if you have not defined namespace upon resource definition, default namespace
2. System namespace (name=system): System namespace is for holding all system resources.

## User
### Overview
User is for authentication purposes.

### User definition
* **id** - unique id
* **username** - unique username
* **password** - password for authentication
* **securityContext** - security context is to apply ACL to resource property -
  see [security context](#security-context)

## Extension
### Overview
Extensions is one of the main features of Data Handler. Extensions is for extending capabilities of Data handler.
So you can define custom resources and you can define how it will work. 

Extensions can be developed technically in any language which supports grpc protocol. But currently we have built in support for golang.
So you can define a resource, an extension and you can develop how your resource will work.

### Extension definition
* **id** - unique namespace id
* **name** - unique namespace name
* **description** - Datasource description
* **namespace** - namespace of resource you want to extend
* **resource** - name of resource you want to extend
