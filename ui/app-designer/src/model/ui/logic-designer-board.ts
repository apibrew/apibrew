


// Sub Types

import { FunctionTrigger } from "../logic/function-trigger";
import { ResourceRule } from "../logic/resource-rule";
import { Function } from "../logic/function"

export interface Expression {
     type: boolean;

}

export interface Statement {
     type: boolean;
     var?: VarStatement;
     if?: IfStatement;
     for?: ForStatement;

}

export interface VarStatement {
     name: string;
     value: Expression;

}

export interface IfStatement {
     name: string;
     value: Expression;

}

export interface ForStatement {
     name: string;
     value: Expression;

}

export interface FunctionDefinition {
     function: Function;
     parameters?: string[];
     statements?: Statement[];

}

// Resource Type
export interface LogicDesignerBoard {
     id: string;
     triggerDefs?: FunctionTrigger[];
     resourceRules?: ResourceRule[];
     version: number;
     functionDefs?: FunctionDefinition[];

}
// Resource and Property Names
export const LogicDesignerBoardName = "LogicDesignerBoard";

export const LogicDesignerBoardIdName = "Id";

export const LogicDesignerBoardTriggerDefsName = "TriggerDefs";

export const LogicDesignerBoardResourceRulesName = "ResourceRules";

export const LogicDesignerBoardVersionName = "Version";

export const LogicDesignerBoardFunctionDefsName = "FunctionDefs";


