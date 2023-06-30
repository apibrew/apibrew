import {useEffect, useState} from "react";
import {RecordService} from "@apibrew/ui-lib";
import {useErrorHandler} from "./error-handler.tsx";
import {useResourceByName} from "./resource.ts";

export function useRecordByName<T>(resourceName: string, namespace: string, recordName: string): T {
    return useRecordBy<T>(resourceName, namespace, 'name', recordName)
}

export function useRecordBy<T>(resourceName: string, namespace: string, key: string, value: string): T {
    const [record, setRecord] = useState<T>()
    const errorHandler = useErrorHandler()

    useEffect(() => {
        try {
            RecordService.findBy(namespace, resourceName, key, value).then(setRecord, errorHandler)
        } catch (e) {
            console.error(e)
            errorHandler(e)
        }
    }, [resourceName, namespace, key, value])

    return record
}

export function useRecords<T>(resourceName: string, namespace: string) {
    const [records, setRecords] = useState<T[]>([])
    const errorHandler = useErrorHandler()

    useEffect(() => {
        RecordService.list<T>(namespace, resourceName).then(setRecords, errorHandler)
    }, [resourceName, namespace])

    return records
}
