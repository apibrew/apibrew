import React from "react";
import {Record} from "../service/record";

export const RecordContext = React.createContext<Record | undefined>(undefined)

export function useRecord<T>() {
    const Record = React.useContext(RecordContext)
    if (!Record) {
        throw new Error("useRecord must be used within a RecordProvider")
    }
    return Record as T
}