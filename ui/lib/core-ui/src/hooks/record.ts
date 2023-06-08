import {useEffect, useState} from "react";
import {RecordService} from "../service";
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