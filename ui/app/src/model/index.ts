import type {components} from './base-schema'

export type AuthenticationRequest = components['schemas']['AuthenticationRequest']
export type AuthenticationResponse = components['schemas']['AuthenticationResponse']

export type RenewTokenRequest = components['schemas']['RenewTokenRequest']
export type RenewTokenResponse = components['schemas']['RenewTokenResponse']

export type Token = AuthenticationResponse['token']
export type BooleanExpression = components['schemas']['BooleanExpression']
export type PairExpression = components['schemas']['PairExpression']

export type Status = components['schemas']['Status']
