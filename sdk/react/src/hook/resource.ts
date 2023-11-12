import { useEffect, useState } from 'react'
import { Resource } from '@apibrew/client/model'
import { useErrorHandler } from '../context/error-handler'
import { useClient } from '../context'

export function useResourceByName(
  resourceName: string,
  namespace = 'default'
): Resource | undefined {
  const errorHandler = useErrorHandler()
  const [resource, setResource] = useState<Resource>()

  const client = useClient()

  useEffect(() => {
    client
      .getResourceByName(namespace, resourceName)
      .then(setResource, errorHandler)
  }, [resourceName, namespace])

  return resource
}
