import {Resource} from "../../model"
import {Record} from "../../service/record"
import {FormConfig as CrudFormConfig} from "../../model/ui/crud.ts";
import {ResourceContext} from "../../context/resource";
import {ValueContext} from "../../context/value";
import {RecordContext} from "../../context/record";
import React from "react";
import {StructElement} from "./form-elements/StructElement.tsx";

export interface FormProps {
    resource: Resource
    record: Record
    readOnly?: boolean
    setRecord: (record: Record) => void
    formConfig: CrudFormConfig
}

export function Form(props: FormProps) {
    return (
        <ResourceContext.Provider value={props.resource}>
            <RecordContext.Provider value={props.record}>
                <ValueContext.Provider value={{
                    value: props.record,
                    onChange: props.setRecord,
                    readOnly: props.readOnly,
                }}>
                    {props.formConfig.children &&
                        <StructElement properties={props.resource.properties} config={props.formConfig}/>}
                </ValueContext.Provider>
            </RecordContext.Provider>
        </ResourceContext.Provider>
    );
}
