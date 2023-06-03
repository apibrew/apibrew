import { Crud } from '@apibrew/core-ui'
import {useParams} from "react-router-dom";

export function CrudPage () {
  const params = useParams()

  return (
    <Crud namespace={params.namespace as string} resource={params.resource as string}/>
  )
}
