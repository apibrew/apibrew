import { Client } from "../client"
import { Function, FunctionResource, RecordResourceInfo, Resource } from "../model"
import { Lambda, LambdaResource } from "../model/logic/lambda"
import { ResourceApi } from "../api"
import { handleError } from "../api/error"
import { getModule, registerModuleChild } from "./module-def"

export function defineResource(resource: Resource) {
    const client = Client.getDefaultClient()

    ResourceApi.apply(client.provider()(), resource).then(resp => {
        console.log(resp)
    }, err => {
        console.error(handleError(err))
    })
}