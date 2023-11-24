import { useEffect, useState } from 'react'
import { Code, Resource } from '@apibrew/client/model'
import { useErrorHandler } from '../context/error-handler'
import { useClient } from '../context'
import { ApiException } from '@apibrew/client'

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

export function useResourceById(id: string): Resource | undefined {
  const errorHandler = useErrorHandler()
  const [resource, setResource] = useState<Resource>()

  const client = useClient()

  useEffect(() => {
    client.listResources().then((response) => {
      const resource = response.find((resource) => resource.id === id)
      if (resource) {
        setResource(resource)
      } else {
        throw new ApiException(
          Code.RESOURCE_NOT_FOUND,
          `Resource with id ${id} not found`
        )
      }
    }, errorHandler)
  }, [id])

  return resource
}
