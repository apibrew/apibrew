import React from "react";

export interface ValueOptions<T> {
    value: T
    onChange: (value: T) => void
    readOnly?: boolean
}

export const ValueContext = React.createContext<ValueOptions<any> | undefined>(undefined)

export function useValue() {
    const Value = React.useContext(ValueContext)
    if (!Value) {
        throw new Error("useValue must be used within a ValueProvider")
    }
    return Value
}