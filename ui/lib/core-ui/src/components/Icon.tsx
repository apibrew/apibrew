import React from "react";

export interface IconProps {
    name: string
    group?: string
    size?: number
    color?: string
    weight?: number
}

export function Icon(props: IconProps) {
    const group = props.group ?? 'material-icons'

    return <span className={`${group}`} style={{
        color: props.color ?? 'inherit',
        fontSize: props.size ?? '24px',
        fontWeight: props.weight ?? 400

    }}>{props.name}</span>
}
