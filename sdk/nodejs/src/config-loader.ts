import {Config} from "./config";
import * as os from "os";
import * as fs from "fs";
import * as yaml from "js-yaml";

export class ConfigLoader {
    static load(): Config {
        const home = os.homedir();

        const apbrConfigFile = home + '/.apbr/config';

        const configData = fs.readFileSync(apbrConfigFile, 'utf8');

        return yaml.load(configData) as any;
    }
}
