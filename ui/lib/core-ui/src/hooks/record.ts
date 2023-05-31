import {useEffect, useState} from "react";
import {RecordService} from "../service";

export function useRecordByName<T>(resourceName: string, namespace: string, recordName: string): T {
    const [record, setRecord] = useState<T>()

    useEffect(() => {
        RecordService.findBy(namespace, resourceName, 'name', recordName).then(setRecord)
    }, [resourceName, namespace, recordName])

    return record
}