class Entity:
    def __init__(self):
        self.id = None

    def get_id(self):
        return self.id

    def set_id(self, id):
        self.id = id

    def get_property(self, property):
        return getattr(self, property)

    def set_property(self, property, value):
        setattr(self, property, value)

    def get_property_map(self):
        return self.__dict__

    def __str__(self):
        return str(self.__dict__)

    def __repr__(self):
        return str(self.__dict__)


class EntityInfo:
    namespace: str
    resource: str
    rest_path: str

    def __init__(self, namespace, resource, rest_path):
        self.namespace = namespace
        self.resource = resource
        self.rest_path = rest_path

    def __str__(self):
        return f"EntityInfo: {self.namespace} {self.resource} {self.rest_path}"

    @staticmethod
    def from_resource(resource):
        return EntityInfo(resource.namespace.name, resource.name, EntityInfo.get_rest_path(resource))

    @staticmethod
    def get_rest_path(resource):
        if resource.annotations and resource.annotations.get("OpenApiRestPath"):
            return resource.annotations.get("OpenApiRestPath")
        elif resource.namespace.name == "default":
            return EntityInfo.slug(resource.name)
        else:
            return EntityInfo.slug(resource.namespace.name + "/" + resource.name)

    @staticmethod
    def slug(name):
        return name.lower().replace("[^a-z0-9]+", "-")

    @staticmethod
    def from_entity_class(entity_class):
        return entity_class.entity_info
