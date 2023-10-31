from typing import List


class Authentication:
    username: str
    password: str
    token: str

    def __init__(self, **entries):
        self.__dict__.update(entries)


class Server:
    name: str
    host: str
    port: int
    httpPort: int
    insecure: bool
    authentication: Authentication

    def __init__(self, **entries):
        self.__dict__.update(entries)


class Config:
    type: str
    defaultServer: str
    servers: List[Server]

    def __init__(self, **entries):
        self.__dict__.update(entries)

    def __str__(self):
        return f"Config: {self.type} {self.defaultServer} {self.servers}"
