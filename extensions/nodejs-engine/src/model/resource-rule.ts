

import { Function } from "./function";


export const ResourceRuleResource = {
    resource: "ResourceRule",
    namespace: "logic",
};

// Sub Types

// Resource Type
export interface ResourceRule {
    id: string;
name: string;
resource: string;
namespace: string;
conditionFunction: Function;
version: number;

}
// Resource and Property Names
export const ResourceRuleName = "ResourceRule";

export const ResourceRuleIdName = "Id";

export const ResourceRuleNameName = "Name";

export const ResourceRuleResourceName = "Resource";

export const ResourceRuleNamespaceName = "Namespace";

export const ResourceRuleConditionFunctionName = "ConditionFunction";

export const ResourceRuleVersionName = "Version";


