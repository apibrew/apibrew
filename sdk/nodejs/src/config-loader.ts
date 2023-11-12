import {Config} from "./config";
import * as yaml from "js-yaml";

let moduleNames = [];

export class ConfigLoader {
    public static load(): Config {
        const os = require('os')
        const fs = require('fs')
        const home = os.homedir();

        const apbrConfigFile = home + '/.apbr/config';

        const configData = fs.readFileSync(apbrConfigFile, 'utf8');

        return yaml.load(configData) as any;
    }

    public static loadServerConfig(serverName: string) {
        const config = ConfigLoader.load();

        if (!serverName) {
            serverName = config.defaultServer;
        }

        const serverConfig = config.servers.find((item: any) => item.name == serverName);

        if (!serverConfig) {
            throw new Error("Server not found: " + serverName);
        }

        return serverConfig
    }
}
