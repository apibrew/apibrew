import React from 'react'
import { BACKEND_URL } from '../config'

export function Test(): JSX.Element {
    return (
        <>
            Hello World {BACKEND_URL}
        </>
    )
}
