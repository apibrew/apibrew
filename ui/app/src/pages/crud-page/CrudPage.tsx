import {useParams} from "react-router-dom";
import {Crud} from "../../components";

export function CrudPage () {
  const params = useParams()

  return (
    <Crud namespace={params.namespace as string} resource={params.resource as string}/>
  )
}
