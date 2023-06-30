import {ActionExecuteFunction} from "./components/logic/function/action-execute-function.tsx";
import {ResourceContextComponent} from "./components/context/ResourceContextComponent.tsx";
import {DashboardLayout} from "./layout/index.ts";
import React, {ComponentType} from "react";
import {ModuleService} from "@apibrew/ui-lib";

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
        CrudSettingsFormConfig: lazyComponent(() => import("./components/custom-inputs/CrudSettingsFormConfig.tsx"), 'CrudSettingsFormConfig'),
        CrudSettingsGridConfig: lazyComponent(() => import("./components/custom-inputs/CrudSettingsGridConfig.tsx"), 'CrudSettingsGridConfig'),
        Test: lazyComponent(() => import("./test/test.tsx"), 'Test'),
        Crud: lazyComponent(() => import("./components/crud/Crud.tsx"), 'Crud'),
        LoginPage: lazyComponent(() => import("./pages/login/login.tsx"), 'Login'),
        Error: lazyComponent(() => import("./pages/error/Error.tsx"), 'Error'),
        SecurityConstraintsInput: lazyComponent(() => import("./components/security/SecurityConstraintsInput.tsx"), 'SecurityConstraintsInput'),
        UserProfile: lazyComponent(() => import("./pages/user/UserProfile.tsx"), 'UserProfile'),
    },
    name: 'CoreUI',
})

ModuleService.registerLocalModule({
    exports: {
        Resource: ResourceContextComponent,
    },
    name: 'Context',
})
