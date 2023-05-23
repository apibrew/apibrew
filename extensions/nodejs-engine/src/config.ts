export const APBR_HOST = process.env.APBR_HOST || 'localhost'
export const APBR_PORT = process.env.APBR_PORT || 9009
export const ENGINE_HOST = process.env.ENGINE_HOST || 'localhost'
export const ENGINE_PORT = process.env.ENGINE_PORT || 23619
export const ENGINE_REMOTE_ADDR = process.env.ENGINE_REMOTE_ADDR || `${ENGINE_HOST}:${ENGINE_PORT}`
export const TOKEN = process.env.APBR_TOKE || ''
export const EXTENSION_NAME = 'nodejs-engine'

console.log(`APBR_HOST: ${APBR_HOST}; APBR_PORT: ${APBR_PORT}; ENGINE_HOST: ${ENGINE_HOST}; ENGINE_PORT: ${ENGINE_PORT}; TOKEN: ${TOKEN}; EXTENSION_NAME: ${EXTENSION_NAME}; ENGINE_REMOTE_ADDR: ${ENGINE_REMOTE_ADDR}`)
