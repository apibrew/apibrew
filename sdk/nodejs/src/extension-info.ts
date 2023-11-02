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
        finalizes: Boolean(extensionInfo.finalizer),
        sync: Boolean(extensionInfo.sync),
        responds: Boolean(extensionInfo.responds),
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

export function extensionInfoToString(extensionInfo?: ExtensionInfo) {
    if (!extensionInfo) {
        return '';
    }
    return extensionInfo.namespaces + "/" + extensionInfo.resources + "/" + extensionInfo.action + "/" + extensionInfo.order + "/" + extensionInfo.sync + "/" + extensionInfo.responds + "/" + extensionInfo.finalizer;
}