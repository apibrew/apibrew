import { Crud } from "../../components/crud/Crud"

export interface CrudPageProps {
    namespace: string
    resource: string
}

export function CrudPage(props: CrudPageProps) {
    return (
        <Crud namespace={props.namespace} resource={props.resource} />
    )
}