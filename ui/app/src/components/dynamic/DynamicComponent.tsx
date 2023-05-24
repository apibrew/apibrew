import React, {ReactNode} from "react";
import * as babel from "@babel/standalone";
import Button from "@mui/material/Button";

export interface DynamicComponentProps {
    componentCode: string
    children?: ReactNode
}

export function DynamicComponent(props: DynamicComponentProps) {
    let code = babel.transform(props.componentCode, {presets: ['react', 'es2015']}).code!;

    code = code.replace(/"use strict";\n+/g, '');

    const params: { [key: string]: any } = {
        exports: {},
        React: React,
        require: (name: string) => {
            switch (name) {
                case '@mui/material/Button':
                    return Button;
            }

            throw new Error(`Cannot find module '${name}'`);
        }
    }

    const paramKeys = Object.keys(params)
    const paramValues = paramKeys.map(key => params[key])

    const func = new Function(...paramKeys, `return (function(){${code}})()`);

    func(...paramValues);

    const Component = params.exports.default;

    return <Component/>
}