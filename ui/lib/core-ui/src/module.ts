import {ModuleData} from "./model/module-data.ts";
import {ModuleService} from "./service/module.ts";
import {ActionExecuteFunction} from "./components/logic/function/action-execute-function.ts";

const moduleData: ModuleData = {
    exports: {
        ActionExecuteFunction: new ActionExecuteFunction(),
    },
    name: 'CoreUI',
    package: 'self',
}

ModuleService.registerLocalModule(moduleData)
