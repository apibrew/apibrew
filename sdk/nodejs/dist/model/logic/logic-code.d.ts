export declare const LogicCodeResource: {
    resource: string;
    namespace: string;
};
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
export interface LogicCode {
    id: string;
    statements?: Statement[];
    version: number;
}
export declare const LogicCodeName = "LogicCode";
export declare const LogicCodeIdName = "Id";
export declare const LogicCodeStatementsName = "Statements";
export declare const LogicCodeVersionName = "Version";
