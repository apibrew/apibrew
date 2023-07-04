import { Function } from "./function";
export declare const ResourceRuleResource: {
    resource: string;
    namespace: string;
};
export interface ResourceRule {
    id: string;
    name: string;
    resource: string;
    namespace: string;
    conditionFunction: Function;
    version: number;
}
export declare const ResourceRuleName = "ResourceRule";
export declare const ResourceRuleIdName = "Id";
export declare const ResourceRuleNameName = "Name";
export declare const ResourceRuleResourceName = "Resource";
export declare const ResourceRuleNamespaceName = "Namespace";
export declare const ResourceRuleConditionFunctionName = "ConditionFunction";
export declare const ResourceRuleVersionName = "Version";
