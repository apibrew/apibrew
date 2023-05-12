import { ModalTypeMap } from "@mui/material";
import React, { ReactNode } from "react";

export interface LayoutOptions {
    showAlert(alert: AlertOptions): void;
    showModal(modal: ModalOptions): ModalOperations;
}

export interface ModalOptions {
    id?: string
    content: ReactNode
    props?: ModalTypeMap['props'],
    onClose?: () => void
}

export interface ModalOperations {
    id?: string
    close(): void
}

export interface AlertOptions {
    message: string;
    severity: "error" | "warning" | "info" | "success";
}

export const LayoutContext = React.createContext<LayoutOptions>({} as any);