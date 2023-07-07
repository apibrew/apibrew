import { Client, Extension, Function, FunctionResource, FunctionTrigger, FunctionTriggerResource, Module, ModuleResource, ResourceRuleResource } from '@apibrew/client'
import { APBR_ADDR, TOKEN } from './config'
import { Lambda, LambdaResource } from './model/lambda'
import { FunctionExecutionEngine, FunctionExecutionEngineResource } from './model/function-execution-engine'


export const apbrClient = new Client(`http://${APBR_ADDR}`)
apbrClient.authenticateToken(TOKEN)

export const extensionRepository = apbrClient.newRepository<Extension>({ namespace: 'system', resource: 'extensions' })
export const functionExecutionEngineRepository = apbrClient.newRepository<FunctionExecutionEngine>(FunctionExecutionEngineResource)
export const functionRepository = apbrClient.newRepository<Function>(FunctionResource)
export const moduleRepository = apbrClient.newRepository<Module>(ModuleResource)
export const triggerRepository = apbrClient.newRepository<FunctionTrigger>(FunctionTriggerResource)
export const lambdaRepository = apbrClient.newRepository<Lambda>(LambdaResource)
export const functionTriggerRepository = apbrClient.newRepository<FunctionTrigger>(FunctionTriggerResource)
export const resourceRuleRpository = apbrClient.newRepository<Function>(ResourceRuleResource)
