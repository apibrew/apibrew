import { Client } from "../client"
import { RecordResourceInfo, Record } from "../model"
import { handleError } from "../api/error"

export function defineRecord<T extends Record<unknown>>(resourceInfo: RecordResourceInfo, record: T) {
    const client = Client.getDefaultClient()

    const repository = client.newRepository(resourceInfo)

    repository.create(record).then(resp => {
        console.log(resp)
    }, err => {
        console.error(handleError(err))
    })
}