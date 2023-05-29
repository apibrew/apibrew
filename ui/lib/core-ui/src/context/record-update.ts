import React from "react";
import {Record} from "../service/record";

export const RecordUpdateContext = React.createContext<(record: Record) => void>(undefined)

export function useRecordUpdate(): (record: Record) => void {
    const RecordUpdate = React.useContext(RecordUpdateContext)
    if (!RecordUpdate) {
        throw new Error("useRecord must be used within a RecordProvider")
    }
    return RecordUpdate
}
