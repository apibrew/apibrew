import {type ModalTypeMap} from '@mui/material'
import React, {type ReactNode, useEffect} from 'react'

export interface BreadCramp {
    title: string
    link?: string
}

export interface LayoutOptions {
    showAlert: (alert: AlertOptions) => void
    showModal: (modal: ModalOptions) => ModalOperations
    setBreadCramps: (list: BreadCramp[]) => void
    breadCramps: BreadCramp[]
}

export interface ModalOptions {
    id?: string
    content: ReactNode
    props?: ModalTypeMap['props']
    onClose?: () => void
}

export interface ModalOperations {
    id?: string
    close: () => void
}

export interface AlertOptions {
    message: string
    severity: 'error' | 'warning' | 'info' | 'success'
}

export const LayoutContext = React.createContext<LayoutOptions>({} as any)

export function useBreadCramps(...list: BreadCramp[]) {
    const layoutContext = React.useContext(LayoutContext)

    useEffect(() => {
        layoutContext.setBreadCramps(list);
    }, [])
}
