import {BooleanExpression} from "./model/extension";

export class BooleanExpressionBuilder {
    public static and(...expressions: BooleanExpression[]): BooleanExpression {
        return {
            and: expressions
        } as BooleanExpression
    }

    public static or(...expressions: BooleanExpression[]): BooleanExpression {
        return {
            or: expressions
        } as BooleanExpression
    }

    public static not(expression: BooleanExpression): BooleanExpression {
        return {
            not: expression
        } as BooleanExpression
    }

    public static equal(property: string, value: any): BooleanExpression {
        return {
            equal: {
                left: {
                    property: property
                },
                right: {
                    value: value
                }
            }
        } as BooleanExpression
    }

    public static eq(property: string, value: any): BooleanExpression {
        return this.equal(property, value)
    }

    public static notEqual(property: string, value: any): BooleanExpression {
        return this.not(this.equal(property, value))
    }

    public static greaterThan(property: string, value: any): BooleanExpression {
        return {
            greaterThan: {
                left: {
                    property: property
                },
                right: {
                    value: value
                }
            }
        } as BooleanExpression
    }

    public static gt(property: string, value: any): BooleanExpression {
        return this.greaterThan(property, value)
    }

    public static greaterThanOrEqual(property: string, value: any): BooleanExpression {
        return {
            greaterThanOrEqual: {
                left: {
                    property: property
                },
                right: {
                    value: value
                }
            }
        } as BooleanExpression
    }

    public static gte(property: string, value: any): BooleanExpression {
        return this.greaterThanOrEqual(property, value)
    }

    public static lessThan(property: string, value: any): BooleanExpression {
        return {
            lessThan: {
                left: {
                    property: property
                },
                right: {
                    value: value
                }
            }
        } as BooleanExpression
    }

    public static lt(property: string, value: any): BooleanExpression {
        return this.lessThan(property, value)
    }

    public static lessThanOrEqual(property: string, value: any): BooleanExpression {
        return {
            lessThanOrEqual: {
                left: {
                    property: property
                },
                right: {
                    value: value
                }
            }
        } as BooleanExpression
    }

    public static lte(property: string, value: any): BooleanExpression {
        return this.lessThanOrEqual(property, value)
    }

    public static in(property: string, values: any[]): BooleanExpression {
        return {
            in: {
                left: {
                    property: property
                },
                right: {
                    value: values
                }
            }
        } as BooleanExpression
    }

    public static notIn(property: string, values: any[]): BooleanExpression {
        return this.not(this.in(property, values))
    }

    public static isNull(property: string): BooleanExpression {
        return {
            isNull: {
                property: property
            }
        } as BooleanExpression
    }

    public static isNotNull(property: string): BooleanExpression {
        return this.not(this.isNull(property))
    }
}