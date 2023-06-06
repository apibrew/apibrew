import {ExportOptions} from "./model/module-data.ts";
import {ModuleService} from "./service/module.ts";
import {ActionExecuteFunction} from "./components/logic/function/action-execute-function";
import {RecordService} from "./service";
import {Layout} from "./model/ui/layout.ts";
import {prepareLayoutComponent} from "./components/dynamic/Layout.tsx";
import {ResourceContextComponent} from "./components/context/ResourceContextComponent.tsx";
import {DashboardLayout} from "./layout";
import React, {ComponentType} from "react";
import {Error} from "./pages/error/Error.tsx";

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

export function lazyComponent<T, C extends ComponentType<any>>(func: () => Promise<T>, componentName: keyof T) {
    const lazyCom: () => Promise<{ default: C }> = () => func().then(imp => {
        return {
            default: imp[componentName] as C
        }
    })

    return React.lazy<C>(lazyCom)
}

ModuleService.registerLocalModule({
    exports: {
        DashboardLayout: DashboardLayout,
        ActionExecuteFunction: new ActionExecuteFunction(),
        FunctionScriptInput: lazyComponent(() => import("./components/custom-inputs/FunctionScriptInput"), 'FunctionScriptInput'),
        CrudSettingsFormConfig: lazyComponent(() => import("./components/custom-inputs/CrudSettingsFormConfig"), 'CrudSettingsFormConfig'),
        CrudSettingsGridConfig: lazyComponent(() => import("./components/custom-inputs/CrudSettingsGridConfig"), 'CrudSettingsGridConfig'),
        Test: lazyComponent(() => import("./test/test"), 'Test'),
        CrudPage: lazyComponent(() => import("./pages/crud-page/CrudPage"), 'CrudPage'),
        AppDesigner: lazyComponent(() => import("./pages/app-designer/index.tsx"), 'AppDesigner'),
        Error: lazyComponent(() => import("./pages/error/Error.tsx"), 'Error'),
    },
    name: 'CoreUI',
})

ModuleService.registerLocalModule({
    exports: {
        Resource: ResourceContextComponent,
    },
    name: 'Context',
})
