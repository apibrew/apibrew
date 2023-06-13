import {ModuleService} from "@apibrew/core-lib";
import {Sdk} from "./components/sdk/Sdk.tsx";

export function registerModule() {
    ModuleService.registerLocalModule({
        exports: {
            Sdk: Sdk,
        },
        name: 'Documentation',
    })
}