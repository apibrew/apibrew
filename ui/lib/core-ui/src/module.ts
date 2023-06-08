import "./service/module.ts";
import * as ModuleService from "./service/module.ts";
import {ActionExecuteFunction} from "./components/logic/function/action-execute-function";
import {ResourceContextComponent} from "./components/context/ResourceContextComponent.tsx";
import {DashboardLayout} from "./layout";
import React, {ComponentType} from "react";
import {Error} from "./pages/error/Error.tsx";
import {SecurityConstraintsInput} from "./components/security/SecurityConstraintsInput.tsx";

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
        CrudSettingsFormConfig: lazyComponent(() => import("./components/custom-inputs/CrudSettingsFormConfig"), 'CrudSettingsFormConfig'),
        CrudSettingsGridConfig: lazyComponent(() => import("./components/custom-inputs/CrudSettingsGridConfig"), 'CrudSettingsGridConfig'),
        Test: lazyComponent(() => import("./test/test"), 'Test'),
        CrudPage: lazyComponent(() => import("./pages/crud-page/CrudPage"), 'CrudPage'),
        LoginPage: lazyComponent(() => import("./pages/login/login.tsx"), 'Login'),
        AppDesigner: lazyComponent(() => import("./pages/app-designer/index.tsx"), 'AppDesigner'),
        Error: lazyComponent(() => import("./pages/error/Error.tsx"), 'Error'),
        SecurityConstraintsInput: lazyComponent(() => import("./components/security/SecurityConstraintsInput.tsx"), 'SecurityConstraintsInput'),
        UserProfile: lazyComponent(() => import("./pages/user/UserProfile"), 'UserProfile'),
    },
    name: 'CoreUI',
})

ModuleService.registerLocalModule({
    exports: {
        Resource: ResourceContextComponent,
    },
    name: 'Context',
})
