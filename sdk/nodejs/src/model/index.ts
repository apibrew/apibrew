import type { components } from './base-schema'

export type AuthenticationRequest = components['schemas']['AuthenticationRequest']
export type AuthenticationResponse = components['schemas']['AuthenticationResponse']

export type RenewTokenRequest = components['schemas']['RenewTokenRequest']
export type RenewTokenResponse = components['schemas']['RenewTokenResponse']

export interface Token {
    /**
     * Format: enum
     * @enum {string}
     */
    term: TokenTerm;
    content: string;
    /** Format: date-time */
    expiration: string;
}

export type BooleanExpression = components['schemas']['BooleanExpression']
export type PairExpression = components['schemas']['PairExpression']

export type User = components['schemas']['User']
export type Resource = components['schemas']['Resource']
export type ResourceProperty = components['schemas']['ResourceProperty']
export type Namespace = components['schemas']['Namespace']

export type Status = components['schemas']['Status']

export type TokenTerm = "VERY_SHORT" | "SHORT" | "MIDDLE" | "LONG" | "VERY_LONG"

export type Extension = components["schemas"]["Extension"]

export type Event = components["schemas"]["Event"]

export * from './annotations'
export * from './security-constraint'
export * from './record'

export * from './logic/function'
export * from './logic/function-execution'
export * from './logic/function-trigger'
export * from './logic/logic-code'
export * from './logic/module'
export * from './logic/resource-rule'
export * from './logic/schedule'
