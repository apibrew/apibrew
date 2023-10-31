class ApiException(Exception):
    def __init__(self, code, message):
        self.code = code
        self.message = message

    def to_dict(self):
        return {'code': self.code, 'message': self.message}

    def __str__(self):
        return f"ApiException: {self.code} {self.message}"
