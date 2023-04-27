import type { components } from './schema'

export type AuthenticationRequest = components['schemas']['AuthenticationRequest']
export type AuthenticationResponse = components['schemas']['AuthenticationResponse']
export type Token = AuthenticationResponse['token']
