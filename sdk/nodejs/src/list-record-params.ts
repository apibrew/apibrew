import {BooleanExpression} from "./model";

export interface ListRecordParams {
    query?: BooleanExpression
    filters?: { [key: string]: string }
    resolveReferences?: string[]
    limit?: number
    offset?: number
    useHistory?: boolean
    aggregation?: Aggregation
    sorting?: SortingItem[]
}

export interface Aggregation {
    items: AggregationItem[]
    grouping: GroupingItem[]
}

export interface AggregationItem {
    name: string
    algorithm: AggregationAlgorithm
    property: string
}

export interface GroupingItem {
    property: string
}

export interface SortingItem {
    property: string
    direction: Direction
}

export enum Direction {
    ASC = "ASC",
    DESC = "DESC"
}

export enum AggregationAlgorithm {
    SUM = "SUM",
    AVG = "AVG",
    MIN = "MIN",
    MAX = "MAX",
    COUNT = "COUNT"
}
