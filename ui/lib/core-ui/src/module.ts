import {ExportOptions} from "./model/module-data.ts";
import {ModuleService} from "./service/module.ts";
import {ActionExecuteFunction} from "./components/logic/function/action-execute-function";
import {RecordService} from "./service";
import {Layout} from "./model/ui/layout.ts";
import {prepareLayoutComponent} from "./components/dynamic/Layout.tsx";
import {Form} from "./components/form/Form.tsx";
import {ResourceContextComponent} from "./components/context/ResourceContextComponent.tsx";
import {RouterComponent} from "./components/basic/RouterComponent.tsx";
import {FunctionScriptInput} from "./components/custom-inputs/FunctionScriptInput.tsx";
import {CrudSettingsFormConfig} from "./components/custom-inputs/CrudSettingsFormConfig.tsx";
import {CrudSettingsGridConfig} from "./components/custom-inputs/CrudSettingsGridConfig.tsx";

ModuleService.registerLocalModuleAwait(RecordService.list<Layout>('ui', 'Layout').then(layouts => {
    const exports: ExportOptions = {}

    layouts.forEach(layout => {
        exports[layout.name] = prepareLayoutComponent(layout)
    })

    return {
        exports: exports,
        name: 'Layout',
    }
}))

ModuleService.registerLocalModule({
    exports: {
        ActionExecuteFunction: new ActionExecuteFunction(),
        FunctionScriptInput: FunctionScriptInput,
        Form: Form,
        Router: RouterComponent,
        CrudSettingsFormConfig: CrudSettingsFormConfig,
        CrudSettingsGridConfig: CrudSettingsGridConfig,
    },
    name: 'CoreUI',
})

ModuleService.registerLocalModule({
    exports: {
        Resource: ResourceContextComponent,
    },
    name: 'Context',
})
