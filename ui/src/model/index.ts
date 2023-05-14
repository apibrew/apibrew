import type { components } from './schema'

export type AuthenticationRequest = components['schemas']['AuthenticationRequest']
export type AuthenticationResponse = components['schemas']['AuthenticationResponse']
export type Token = AuthenticationResponse['token']

export type User = components['schemas']['User']
export type Resource = components['schemas']['Resource']
export type ResourceProperty = components['schemas']['ResourceProperty']
export type Namespace = components['schemas']['Namespace']
