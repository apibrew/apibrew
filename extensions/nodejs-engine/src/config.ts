console.log(process.env)
export const APBR_ADDR = process.env.APBR_HOST || 'localhost:9009'
export const ENGINE_ADDR = process.env.ENGINE_HOST || 'localhost:23619'
export const ENGINE_REMOTE_ADDR = process.env.ENGINE_REMOTE_ADDR || `${ENGINE_ADDR}`
export const TOKEN = process.env.APBR_TOKE || ''
export const EXTENSION_NAME = 'nodejs-engine'

console.log(`APBR_ADDR: ${APBR_ADDR}; ENGINE_ADDR: ${ENGINE_ADDR}; ENGINE_REMOTE_ADDR: ${ENGINE_REMOTE_ADDR}; TOKEN: ${TOKEN}; EXTENSION_NAME: ${EXTENSION_NAME}`)
