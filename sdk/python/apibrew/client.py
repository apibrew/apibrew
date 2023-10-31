import requests

from apibrew.apiexception import ApiException
from apibrew.configloader import ConfigLoader

class Urls:
    @staticmethod
    def events_url(url):
        return f"{url}/_events"

    @staticmethod
    def resource_url(url):
        return f"{url}/resources"

    @staticmethod
    def record_url(url, rest_path):
        return f"{url}/{rest_path}"

    @staticmethod
    def record_search_url(url, rest_path):
        return f"{url}/{rest_path}/_search"

    @staticmethod
    def record_watch_url(url, rest_path):
        return f"{url}/{rest_path}/_watch"

    @staticmethod
    def resource_by_name(url, namespace, name):
        return f"{Urls.resource_url(url)}/by-name/{namespace}/{name}"

    @staticmethod
    def resource_by_id(url, id):
        return f"{Urls.resource_url(url)}/{id}"

    @staticmethod
    def record_by_id_url(url, rest_path, id):
        return f"{url}/{rest_path}/{id}"

    @staticmethod
    def record_action_by_id_url(url, rest_path, id, action):
        return f"{url}/{rest_path}/{id}/_{action}"

    @staticmethod
    def authenticate(url):
        return f"{url}/authentication/token"


class Client:
    def __init__(self, url):
        self.url = url
        self.token = None
        self.bypass_extensions = False
        self.repository_class_map = {}
        self.repository_entity_map = {}

    @staticmethod
    def new_client():
        return Client.new_client_by_server_name(None)

    @staticmethod
    def new_client_by_server_name(server_name):
        config = ConfigLoader.load()

        print(config)

        if server_name is None:
            server_name = config.defaultServer

        print(config.servers[0])

        # server_config = next((item for item in config.servers if item.name == server_name), None)
        #
        # if server_config is None:
        #     raise Exception(f"Server not found: {server_name}")
        #
        # return Client.new_client_by_server_config(server_config)


    def authenticate_with_token(self, token):
        self.token = token

    def apply_resource(self, resource):
        resp = requests.get(Urls.resource_by_name(self.url, resource.namespace.name, resource.name),
                            headers=self.headers())
        exists_status = resp.status_code
        if exists_status == 200:
            return self.update_resource(resource)
        elif exists_status == 404:
            return self.create_resource(resource)
        else:
            return self.ensure_response_success(resp)

    def get_resource_by_name(self, namespace, name):
        return requests.get(Urls.resource_by_name(self.url, namespace, name), headers=self.headers()).json()

    def list_resources(self):
        result = requests.get(Urls.resource_url(self.url), headers=self.headers()).json()
        self.ensure_response_success(result)
        return result

    def create_resource(self, resource):
        result = requests.post(Urls.resource_url(self.url), json=resource, headers=self.headers()).json()
        self.ensure_response_success(result)
        return result

    def update_resource(self, resource):
        result = requests.post(Urls.resource_by_id(self.url, str(resource.id)), json=resource,
                               headers=self.headers()).json()
        self.ensure_response_success(result)
        return result

    def delete_resource(self, resource):
        result = requests.delete(Urls.resource_by_id(self.url, str(resource.id)), headers=self.headers()).json()
        self.ensure_response_success(result)
        return result

    def authenticate_with_username_and_password(self, username, password):
        body = {
            "username": username,
            "password": password,
            "term": "VERY_LONG"
        }
        result = requests.post(Urls.authenticate(self.url), json=body).json()
        if result.status_code == 200:
            self.token = result.token.content
        else:
            self.ensure_response_success(result)

    def new_client_authenticate_with_token(self, token):
        client = Client(self.url)
        client.bypass_extensions = self.bypass_extensions
        client.authenticate_with_token(token)
        return client

    def new_client_authenticate_with_username_and_password(self, username, password):
        client = Client(self.url)
        client.bypass_extensions = self.bypass_extensions

    def headers(self):
        if self.token is not None:
            return {"Authorization": f"Bearer {self.token}"}
        else:
            return {}

    @staticmethod
    def ensure_response_success(resp):
        if resp.status_code != 200:
            raise ApiException(resp.status_code, resp.content)
        return resp
