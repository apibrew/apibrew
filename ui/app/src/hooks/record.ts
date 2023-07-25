import { useEffect, useState } from "react";
import { RecordService } from "@apibrew/ui-lib";
import { useErrorHandler } from "./error-handler.tsx";

export function useRecordByName<T>(resourceName: string, namespace: string, recordName: string): T | undefined {
    return useRecordBy<T>(resourceName, namespace, 'name', recordName)
}

export function useRecordBy<T>(resourceName: string, namespace: string, key: string, value: string): T | undefined {
    const [record, setRecord] = useState<T>()
    const errorHandler = useErrorHandler()

    useEffect(() => {
        try {
            RecordService.findBy<T>(namespace, resourceName, key, value).then(setRecord, errorHandler)
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
