import { Client } from '@apibrew/client'
import { APBR_ADDR, TOKEN } from './config'


export const apbrClient = new Client(`http://${APBR_ADDR}`)
apbrClient.authenticateToken(TOKEN)
