import {BooleanExpression} from "./model/extension";

export interface ListRecordParams {
    query?: BooleanExpression
    filters?: { [key: string]: string }
    resolveReferences?: string[]
    limit?: number
    offset?: number
    useHistory?: boolean
}