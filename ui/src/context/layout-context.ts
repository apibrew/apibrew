import React from "react";

export interface LayoutOptions {
    showAlert(alert: AlertOptions): void;
}

export interface AlertOptions {
    message: string;
    severity: "error" | "warning" | "info" | "success";
}

export const LayoutContext = React.createContext<LayoutOptions>({} as any);