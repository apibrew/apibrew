package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var ResourcePropertyProperties = []*model.ResourceProperty{
	{
		Name:  "name",
		Title: util.Pointer("Name"),
		Description: util.Pointer(`The name of the property. 
Name is the main parameter of property, it is used to identify the property. It is also used to name record properties. 
For example {"title": "Lord of the Rings"} And there "title" is a property, and it is defined by name "title", in its Resource 
		`),
		Type:         model.ResourceProperty_STRING,
		Length:       256,
		ExampleValue: structpb.NewStringValue("title"),
		Required:     false, // can be optional for item types
	},
	{
		Name:        "type",
		Title:       util.Pointer("Type"),
		Description: util.Pointer(`The type of the property. Property Data Types can be one of it. Types can be written with all capital letters.`),
		Type:        model.ResourceProperty_ENUM,
		EnumValues: []string{
			"BOOL",
			"STRING",
			"FLOAT32",
			"FLOAT64",
			"INT32",
			"INT64",
			"BYTES",
			"UUID",
			"DATE",
			"TIME",
			"TIMESTAMP",
			"OBJECT",
			"MAP",
			"LIST",
			"REFERENCE",
			"ENUM",
			"STRUCT",
		},
		ExampleValue: structpb.NewStringValue("STRING"),
		Required:     true,
	},
	{
		Name:  "typeRef",
		Title: util.Pointer("Type Reference"),
		Description: util.Pointer(`The type reference of the property. It is only used for STRUCT type. 
When you used STRUCT type, you need to define your type inside types of resource and then you can use its name as typeRef.`),
		Type:         model.ResourceProperty_STRING,
		Length:       256,
		ExampleValue: structpb.NewStringValue("BookPublishingDetails"),
		Required:     false,
	},
	{
		Name:  "primary",
		Title: util.Pointer("Primary"),
		Description: util.Pointer(`The primary property of the resource. It is used to identify the resource. When it is not supplied, an id property is automatically created.
Normally primary property should not be provided. It is only used for special cases. If provided, it can break some functionalities of system. 
If Primary is provided, it should be a single property. It can not be a list or map.
If Primary is provided, internal id property will not be created.
`),
		Type:     model.ResourceProperty_BOOL,
		Required: true,

		DefaultValue: structpb.NewBoolValue(false),
	},
	{
		Name:  "required",
		Title: util.Pointer("Required"),
		Description: util.Pointer(`This property indicates that whether or not given property is required.
When creating/updating records, if required property is not and defaultValue is given in property, the system will allow request but will use default value instead.
(In all cases if default value is provided it will be used in case of property absence)
`),
		Type:     model.ResourceProperty_BOOL,
		Required: true,

		DefaultValue: structpb.NewBoolValue(false),
	},
	{
		Name:  "unique",
		Title: util.Pointer("Unique"),
		Description: util.Pointer(`This property indicates that whether or not given property is unique.
Unique property is only working for single property, for combination of properties to become unique, you can use indexes with unique flag 
`),
		Type:     model.ResourceProperty_BOOL,
		Required: true,

		DefaultValue: structpb.NewBoolValue(false),
	},
	{
		Name:        "immutable",
		Title:       util.Pointer("Immutable"),
		Description: util.Pointer("This property indicates that whether or not given property is immutable. Immutable properties can not be updated."),
		Type:        model.ResourceProperty_BOOL,
		Required:    true,

		DefaultValue: structpb.NewBoolValue(false),
	},
	{
		Name:         "length",
		Title:        util.Pointer("Length"),
		Description:  util.Pointer("This property indicates the length of the property. It is only used for STRING type."),
		Type:         model.ResourceProperty_INT32,
		Required:     true,
		ExampleValue: structpb.NewNumberValue(256),
		DefaultValue: structpb.NewNumberValue(256),
	},
	{
		Name:        "item",
		Title:       util.Pointer("Item"),
		Description: util.Pointer(`This property indicates the item type of the property. It is only used for LIST and MAP types.`),
		Type:        model.ResourceProperty_STRUCT,
		Required:    false,
		ExampleValue: structpb.NewStructValue(&structpb.Struct{
			Fields: map[string]*structpb.Value{
				"type": structpb.NewStringValue("STRING"),
			},
		}),
		TypeRef: util.Pointer("Property"),
	},
	{
		Name:  "reference",
		Title: util.Pointer("Reference"),
		Description: util.Pointer(`This property indicates the reference type of the property. It is only used for REFERENCE type.
When you use REFERENCE type, you need to provide reference details.
Reference details is used to locate referenced resource
When providing reference details, you need to provide namespace and resource name of the referenced resource.
If you don't provide namespace, it will be assumed as the same namespace with the resource.
`),
		Type: model.ResourceProperty_STRUCT,
		ExampleValue: structpb.NewStructValue(&structpb.Struct{
			Fields: map[string]*structpb.Value{
				"resource": structpb.NewStructValue(&structpb.Struct{
					Fields: map[string]*structpb.Value{
						"namespace": structpb.NewStringValue("default"),
						"resource":  structpb.NewStringValue("Book"),
					},
				}),
			},
		}),
		TypeRef: util.Pointer("Reference"),
	},
	{
		Name:  "defaultValue",
		Title: util.Pointer("Default Value"),
		Description: util.Pointer(`This property indicates the default value of the property. 
It is used when creating/updating records and property is not provided.
`),
		Type:         model.ResourceProperty_OBJECT,
		ExampleValue: structpb.NewStringValue("Lord of the Rings"),
		Required:     false,
	},
	{
		Name:        "enumValues",
		Title:       util.Pointer("Enum Values"),
		Description: util.Pointer(`This property is only used with ENUM type. This property indicates the enum values of the property.`),
		Type:        model.ResourceProperty_LIST,
		Item: &model.ResourceProperty{
			Type: model.ResourceProperty_STRING,
		},
		ExampleValue: structpb.NewListValue(&structpb.ListValue{
			Values: []*structpb.Value{
				structpb.NewStringValue("UNKNOWN"),
				structpb.NewStringValue("ASC"),
				structpb.NewStringValue("DESC"),
			},
		}),
		Required: false,
	},
	{
		Name:         "exampleValue",
		Title:        util.Pointer("Example Value"),
		Description:  util.Pointer(`This property indicates the example value of the property.`),
		Type:         model.ResourceProperty_OBJECT,
		ExampleValue: structpb.NewStringValue(`no-book-name`),
		Required:     false,
	},
	{
		Name:         "title",
		Title:        util.Pointer("Title"),
		Description:  util.Pointer(`This property indicates the title of the property. It is used to have meaningful names for the properties.`),
		Type:         model.ResourceProperty_STRING,
		Length:       256,
		ExampleValue: structpb.NewStringValue(`Book Title`),
		Required:     false,
	},
	{
		Name:         "description",
		Title:        util.Pointer("Description"),
		Description:  util.Pointer(`This property indicates the description of the property. It is used to have meaningful description for the properties.`),
		Type:         model.ResourceProperty_STRING,
		Length:       256,
		ExampleValue: structpb.NewStringValue(`Book Title is a property of Book Resource. It represents the title of the book.`),
		Required:     false,
	},
	special.AnnotationsProperty,
}

var PropertyType = &model.ResourceSubType{
	Name:        "Property",
	Title:       "Property",
	Description: "Property is a type that represents a property of a resource. It is like an API properties or properties of class in a programming language",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: ResourcePropertyProperties,
}

var SubTypeType = &model.ResourceSubType{
	Name:        "SubType",
	Title:       "Sub Type",
	Description: "Sub Type is a type that represents a sub type of a resource. It is mostly used by STRUCT type to define the properties of the struct. ",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: []*model.ResourceProperty{
		{
			Name:         "name",
			Title:        util.Pointer("Name"),
			Description:  util.Pointer("The name of the sub type. "),
			Type:         model.ResourceProperty_STRING,
			ExampleValue: structpb.NewStringValue("Book"),
			Required:     true,
		},
		{
			Name:         "title",
			Title:        util.Pointer("Title"),
			Description:  util.Pointer("The title of the sub type. It is used to have meaningful names for the sub types."),
			Type:         model.ResourceProperty_STRING,
			Length:       256,
			ExampleValue: structpb.NewStringValue("Book"),
			Required:     false,
		},
		{
			Name:         "description",
			Title:        util.Pointer("Description"),
			Description:  util.Pointer("The description of the sub type. It is used to have meaningful description for the sub types. "),
			Type:         model.ResourceProperty_STRING,
			Length:       256,
			ExampleValue: structpb.NewStringValue("Book is a sub type of Resource. It represents a book in the system. "),
			Required:     false,
		},
		{
			Name:        "properties",
			Title:       util.Pointer("Properties"),
			Description: util.Pointer("The properties of the sub type. It is used to define the properties of the sub type. "),
			Type:        model.ResourceProperty_LIST,
			Required:    true,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Property"),
			},
			ExampleValue: structpb.NewListValue(&structpb.ListValue{
				Values: []*structpb.Value{
					structpb.NewStructValue(&structpb.Struct{
						Fields: map[string]*structpb.Value{
							"name": structpb.NewStringValue("title"),
							"type": structpb.NewStringValue("STRING"),
						},
					}),
				},
			}),
		},
	},
}

var ReferenceType = &model.ResourceSubType{
	Name:        "Reference",
	Title:       "Reference",
	Description: "Reference is a type that represents a reference to another resource. It is used to define the reference to another resource. ",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: []*model.ResourceProperty{
		{
			Name:  "resource",
			Title: util.Pointer("Resource"),
			Description: util.Pointer(`This property indicates the resource of the reference.
When providing resource, you need to provide namespace and resource name of the referenced resource.
If you don't provide namespace, it will be assumed as the same namespace with the resource.
`),
			Type: model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: "system",
				Resource:  "Resource",
			},
			ExampleValue: structpb.NewStructValue(&structpb.Struct{
				Fields: map[string]*structpb.Value{
					"namespace": structpb.NewStructValue(&structpb.Struct{
						Fields: map[string]*structpb.Value{
							"name": structpb.NewStringValue("test-namespace"),
						},
					}),
					"name": structpb.NewStringValue("Book"),
				},
			}),
		},
		{
			Name:  "cascade",
			Title: util.Pointer("Cascade"),
			Description: util.Pointer(`This property indicates that whether or not given reference is cascade.
If it is true, when referenced resource record is deleted, all the records that are referencing to that resource will be deleted.
`),
			Type: model.ResourceProperty_BOOL,
		},
		{
			Name:  "backReference",
			Title: util.Pointer("Back Reference"),
			Description: util.Pointer(`This property indicates that whether or not given reference is back reference.
Back reference is reverse of reference, If resource A has reference to resource B, in that case resource B can have back reference to resource A.
Back reference is used only as List.
Backreference should be the name of property in the referenced resource. (like author inside book)
For example:
	Book -> Author.
	Book will have reference to Author. And Author can have back reference to the list of books

`),
			Type:         model.ResourceProperty_STRING,
			ExampleValue: structpb.NewStringValue("author"),
		},
	},
}

var ResourceResource = &model.Resource{
	Name:        "Resource",
	Namespace:   "system",
	Title:       util.Pointer("Resource"),
	Description: util.Pointer("Resource is a top level resource that represents a model and API in the system"),
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource",
	},
	Types: []*model.ResourceSubType{
		PropertyType,
		SubTypeType,
		special.AuditDataSubType,
		{
			Name: "IndexProperty",
			Properties: []*model.ResourceProperty{
				{
					Name:     "name",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				{
					Name:     "order",
					Type:     model.ResourceProperty_ENUM,
					Required: false,
					EnumValues: []string{
						"UNKNOWN", "ASC", "DESC",
					},
					DefaultValue: structpb.NewStringValue("ASC"),
				},
			},
		},
		{
			Name: "Index",
			Properties: []*model.ResourceProperty{
				{
					Name: "properties",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("IndexProperty"),
					},
				},
				{
					Name:     "indexType",
					Type:     model.ResourceProperty_ENUM,
					Required: false,
					EnumValues: []string{
						"BTREE", "HASH",
					},
					DefaultValue: structpb.NewStringValue("BTREE"),
				},
				{
					Name:     "unique",
					Type:     model.ResourceProperty_BOOL,
					Required: false,
				},
				special.AnnotationsProperty,
			},
		},
		ReferenceType,
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperty,
		{
			Name:         "name",
			Title:        util.Pointer("Name"),
			Description:  util.Pointer("The name of the resource. Name is the main parameter of resource, it is used to identify the resource. It is also used to name API endpoints."),
			Type:         model.ResourceProperty_STRING,
			Length:       256,
			Required:     true,
			Unique:       false,
			ExampleValue: structpb.NewStringValue("Book"),
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:        "namespace",
			Title:       util.Pointer("Namespace"),
			Description: util.Pointer("The namespace of the resource. Namespace is used to group resources. It is also used to name API endpoints together with Resource. "),
			Type:        model.ResourceProperty_REFERENCE,
			Required:    true,
			Reference: &model.Reference{
				Resource:  NamespaceResource.Name,
				Namespace: NamespaceResource.Namespace,
				Cascade:   false,
			},
			ExampleValue: structpb.NewStructValue(&structpb.Struct{
				Fields: map[string]*structpb.Value{
					"name": structpb.NewStringValue("system"),
				},
			}),
		},
		{
			Name:  "virtual",
			Title: util.Pointer("Virtual"),
			Description: util.Pointer(`This property indicates that whether or not given resource is virtual. 
Virtual resources are not stored in database. They are created on the fly.
Virtual resources are used to prepare bind them to extensions or nano codes, etc. without touching to database.
`),
			Type:     model.ResourceProperty_BOOL,
			Required: true,

			DefaultValue: structpb.NewBoolValue(false),
		},
		{
			Name:        "properties",
			Title:       util.Pointer("Properties"),
			Description: util.Pointer(`This property indicates the properties of the resource.`),
			Type:        model.ResourceProperty_LIST,
			Required:    true,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Property"),
			},
			ExampleValue: structpb.NewListValue(&structpb.ListValue{
				Values: []*structpb.Value{
					structpb.NewStructValue(&structpb.Struct{
						Fields: map[string]*structpb.Value{
							"name": structpb.NewStringValue("title"),
							"type": structpb.NewStringValue("name"),
						},
					}),
					structpb.NewStructValue(&structpb.Struct{
						Fields: map[string]*structpb.Value{
							"name": structpb.NewStringValue("type"),
							"type": structpb.NewStringValue("STRING"),
						},
					}),
				},
			}),
		},
		{
			Name:  "indexes",
			Title: util.Pointer("Indexes"),
			Description: util.Pointer(`This property indicates the indexes of the resource.
Indexes are used to speed up the queries. Indexes are used to define complex unique constraints.
`),
			Type:     model.ResourceProperty_LIST,
			Required: false,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Index"),
			},
		},
		{
			Name:  "types",
			Title: util.Pointer("Types"),
			Description: util.Pointer(`This property indicates the types of the resource.
This is used to hav sub types, which will be used by other properties which has type STRUCT.
`),
			Type:     model.ResourceProperty_LIST,
			Required: false,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("SubType"),
			},
		},
		{
			Name:         "immutable",
			Title:        util.Pointer("Immutable"),
			Description:  util.Pointer("This property indicates that whether or not given resource is immutable. Immutable resources can not be updated or deleted."),
			Type:         model.ResourceProperty_BOOL,
			Required:     true,
			DefaultValue: structpb.NewBoolValue(false),
		},
		{
			Name:  "abstract",
			Title: util.Pointer("Abstract"),
			Description: util.Pointer(`This property indicates that whether or not given resource is abstract.
Abstract resources are not stored in database. No record related operation is allowed in abstract resources.
Abstract resources are mostly used for code generation (for abstract types, etc.)
`),
			Type:     model.ResourceProperty_BOOL,
			Required: true,

			DefaultValue: structpb.NewBoolValue(false),
		},
		{
			Name:        "checkReferences",
			Title:       util.Pointer("Check References"),
			Description: util.Pointer(`This property indicates that whether or not check references is enabled. Check references resources are used to check references to other resources. It is acting if enabled only in create/update operations`),
			Type:        model.ResourceProperty_BOOL,
			Required:    true,

			DefaultValue: structpb.NewBoolValue(false),
		},
		{
			Name:  "dataSource",
			Title: util.Pointer("Data Source"),
			Description: util.Pointer(`This property indicates the data source of the resource.
Data source is used to store the records of the resource.
Each resource can have only one data source. But data source can be different from resource to resource.
Updating data source of a resource is not migrating any data.
DataSource property is used for non-virtual resources.
If DataSource is not provided, default DataSource will be used
`),
			Type:     model.ResourceProperty_REFERENCE,
			Required: false,
			Reference: &model.Reference{
				Resource:  DataSourceResource.Name,
				Namespace: DataSourceResource.Namespace,
				Cascade:   false,
			},
		},
		{
			Name:         "entity",
			Title:        util.Pointer("Entity"),
			Description:  util.Pointer(`This property indicates the entity of the resource. By entity, table name is considered for relational databases`),
			Type:         model.ResourceProperty_STRING,
			Length:       256,
			Required:     false,
			ExampleValue: structpb.NewStringValue("book"),
		},
		{
			Name:         "catalog",
			Title:        util.Pointer("Catalog"),
			Description:  util.Pointer(`This property indicates the catalog of the resource. By catalog, schema name is considered for relational databases.`),
			Type:         model.ResourceProperty_STRING,
			Length:       256,
			Required:     false,
			ExampleValue: structpb.NewStringValue("public"),
		},
		{
			Name:         "title",
			Title:        util.Pointer("Title"),
			Description:  util.Pointer(`This property indicates the title of the resource. It is used to have meaningful names for the resources.`),
			Type:         model.ResourceProperty_STRING,
			Length:       256,
			Required:     false,
			ExampleValue: structpb.NewStringValue(`Book`),
		},
		{
			Name:         "description",
			Title:        util.Pointer("Description"),
			Description:  util.Pointer(`This property indicates the description of the resource. It is used to have meaningful description for the resources.`),
			Type:         model.ResourceProperty_STRING,
			Length:       256,
			Required:     false,
			ExampleValue: structpb.NewStringValue(`Book is a resource in the system. It represents a book in the system.`),
		},
		special.AnnotationsProperty,
	},
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name: "namespace",
				},
				{
					Name: "name",
				},
			},
			Unique: true,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:     annotations.Enabled,
		annotations.RestApiDisabled: annotations.Enabled,
		annotations.OpenApiGroup:    OpenApiMeta,
		annotations.OpenApiRestPath: "resources",
	},
}
