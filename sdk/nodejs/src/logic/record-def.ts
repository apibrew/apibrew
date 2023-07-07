import { Client } from "../client"
import { Function, FunctionResource, RecordResourceInfo, Record } from "../model"
import { Lambda, LambdaResource } from "../model/logic/lambda"
import { ResourceService } from "../service"
import { handleError } from "../service/error"

export function defineRecord<T extends Record<unknown>>(resourceInfo: RecordResourceInfo, record: T) {
    const client = Client.getDefaultClient()

    const repository = client.newRepository(resourceInfo)

    repository.create(record).then(resp => {
        console.log(resp)
    }, err => {
        console.error(handleError(err))
    })
}