import {Action, EventSelector, Extension} from "./model/extension";

export interface ExtensionInfo {
    namespaces?: string[]
    resources?: string[]
    action: Action
    sync: boolean
    responds: boolean
    finalizer: boolean
    order: number
    sealResource?: boolean
}

export function toExtension(extensionInfo: ExtensionInfo): Extension {
    return {
        name: extensionInfo.namespaces + "/" + extensionInfo.resources + "/" + extensionInfo.action + "/" + extensionInfo.order + "/" + extensionInfo.sync + "/" + extensionInfo.responds + "/" + extensionInfo.finalizer,
        finalizes: extensionInfo.finalizer,
        sync: extensionInfo.sync,
        responds: extensionInfo.responds,
        order: extensionInfo.order,
        selector: prepareEventSelector(extensionInfo),
    } as Extension;
}

export function prepareEventSelector(extensionInfo: ExtensionInfo): EventSelector {
    return {
        namespaces: extensionInfo.namespaces,
        resources: extensionInfo.resources,
        actions: [extensionInfo.action],
    } as EventSelector;
}

//     }
//
//     public Extension toExtension() {
//         Extension extension = new Extension();
//
//         extension.setName(namespaces + "/" + resources + "/" + action + "/" + order + "/" + sync + "/" + responds + "/" + finalizer);
//
//         extension.setFinalizes(finalizer);
//         extension.setSync(sync);
//         extension.setResponds(responds);
//         extension.setOrder(order);
//         extension.setSelector(prepareEventSelector());
//
//
//         return extension;
//     }
//
//     private Extension.EventSelector prepareEventSelector() {
//         Extension.EventSelector eventSelector = new Extension.EventSelector();
//
//         if (namespaces != null) {
//             eventSelector.setNamespaces(namespaces);
//         }
//
//         if (resources != null) {
//             eventSelector.setResources(resources);
//         }
//
//         if (action != null) {
//             eventSelector.setActions(List.of(action));
//         }
//
//         return eventSelector;
//     }
//
//     @Override
//     public String toString() {
//         return "ExtensionInfo{" +
//                 "namespaces=" + namespaces +
//                 ", resources=" + resources +
//                 ", action=" + action +
//                 ", sync=" + sync +
//                 ", responds=" + responds +
//                 ", finalizer=" + finalizer +
//                 ", order=" + order +
//                 '}';
//     }
// }