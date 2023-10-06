

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
annotations?: object;
createdBy?: string;
updatedBy?: string;
createdOn?: string;
updatedOn?: string;
version: number;

}
// Resource and Property Names
export const ResourceRuleName = "ResourceRule";

export const ResourceRuleIdName = "Id";

export const ResourceRuleNameName = "Name";

export const ResourceRuleResourceName = "Resource";

export const ResourceRuleNamespaceName = "Namespace";

export const ResourceRuleConditionFunctionName = "ConditionFunction";

export const ResourceRuleAnnotationsName = "Annotations";

export const ResourceRuleCreatedByName = "CreatedBy";

export const ResourceRuleUpdatedByName = "UpdatedBy";

export const ResourceRuleCreatedOnName = "CreatedOn";

export const ResourceRuleUpdatedOnName = "UpdatedOn";

export const ResourceRuleVersionName = "Version";


