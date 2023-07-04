


export const LogicCodeResource = {
    resource: "LogicCode",
    namespace: "logic",
};

// Sub Types

export interface Expression {
     type: boolean;

}

export interface Statement {
     type: boolean;
     var?: VarStatement;

}

export interface VarStatement {
     name: string;
     value: Expression;

}

// Resource Type
export interface LogicCode {
    id: string;
statements?: Statement[];
version: number;

}
// Resource and Property Names
export const LogicCodeName = "LogicCode";

export const LogicCodeIdName = "Id";

export const LogicCodeStatementsName = "Statements";

export const LogicCodeVersionName = "Version";


