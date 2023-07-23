import type {components} from './base-schema'

export type AuthenticationRequest = components['schemas']['AuthenticationRequest']
export type AuthenticationResponse = components['schemas']['AuthenticationResponse']

export type RenewTokenRequest = components['schemas']['RenewTokenRequest']
export type RenewTokenResponse = components['schemas']['RenewTokenResponse']

export type Token = AuthenticationResponse['token']
export type BooleanExpression = components['schemas']['BooleanExpression']
export type PairExpression = components['schemas']['PairExpression']

export type User = components['schemas']['User']
export type Resource = components['schemas']['Resource']
export type ResourceProperty = components['schemas']['ResourceProperty']
export type Namespace = components['schemas']['Namespace']

export type Status = components['schemas']['Status']


export * from './annotations.ts'
export * from './component-interfaces.ts'
export * from './event.ts'
export * from './module-data.ts'
export * from './system/security-constraint.ts'
export * from './ui/module.ts'
