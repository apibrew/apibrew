import { cwd } from "process"

export const APBR_ADDR = process.env.APBR_HOST || 'localhost:9009'
export const ENGINE_ADDR = process.env.ENGINE_HOST || 'http://localhost:23619'
export const ENGINE_REMOTE_ADDR = process.env.ENGINE_REMOTE_ADDR || `${ENGINE_ADDR}`
export const TOKEN = process.env.APBR_TOKEN || 'eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tL2FwaWJyZXcvYXBpYnJldyIsInN1YiI6ImFkbWluIiwiYXVkIjpbImdpdGh1Yi5jb20vYXBpYnJldy9hcGlicmV3Il0sImV4cCI6MTc1MDkyMDc4OCwibmJmIjoxNjg3ODQ4Nzg4LCJpYXQiOjE2ODc4NDg3ODgsImp0aSI6ImEyNzhkM2QwLTUyOWYtNGI5Mi1iMzFjLWEwZTA3ZGRjN2FiNyIsInNlY3VyaXR5Q29uc3RyYWludHMiOlt7Im5hbWVzcGFjZSI6IioiLCJyZXNvdXJjZSI6IioiLCJwcm9wZXJ0eSI6IioiLCJvcGVyYXRpb24iOiJGVUxMIiwicm9sZSI6InJvb3QifV0sInVzZXJuYW1lIjoiYWRtaW4iLCJyb2xlcyI6WyJyb290Il0sInVpZCI6IjI5MDA1NTY2LWZlYjktMTFlZC1hNGM5LWM2YWFjNjRmMTliMiJ9.LZr5a0OJHixRMbFungSddyX9AsdLzPjzAiRdoPgusu_PqplO8CJJikWn8RGoVJ8nHchVywTD5qVv9BAXxhpQHsaYcmZ8vY3-rAKv2fhEwmlh4NuXD4jUAnCeO-7K1cP1kNQKvw78-EmpJ48Xfzfe7s-IUPrFEPnnbxCIRBJ8pOY'
export const EXTENSION_NAME = 'nodejs-engine'
export const FN_DIR = process.env.FN_DIR || `${cwd()}/fn`

console.log(`APBR_ADDR: ${APBR_ADDR}; ENGINE_ADDR: ${ENGINE_ADDR}; ENGINE_REMOTE_ADDR: ${ENGINE_REMOTE_ADDR}; TOKEN: ${TOKEN}; EXTENSION_NAME: ${EXTENSION_NAME}`)