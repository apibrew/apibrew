import os
import yaml

from apibrew.config import Config


class ConfigLoader:
    @staticmethod
    def load():
        home = os.path.expanduser("~")
        apbr_config_file = home + '/.apbr/config'
        config_data = open(apbr_config_file, 'r').read()
        return Config(**yaml.load(config_data, yaml.FullLoader))

