import { ModuleService } from "@apibrew/ui-lib";

export function registerModule() {
    ModuleService.registerLocalModule({
        exports: {
            ResourceDesigner: () => import('./components/resource-designer/ResourceDesigner'),
            LogicDesigner: () => import('./components/logic-designer/LogicDesigner'),
        },
        name: 'Documentation',
    })
}