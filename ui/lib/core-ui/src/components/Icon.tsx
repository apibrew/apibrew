import React from "react";

export interface IconProps {
    name: string
    group?: string
    size?: string
}

export function Icon(props: IconProps) {
    const group = props.group ?? 'material-icons'
    const size = props.size ?? 'md-18'

    return <span className={`${group} ${size}`}>{props.name}</span>
}
