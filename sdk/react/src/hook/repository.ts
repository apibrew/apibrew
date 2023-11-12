import { useClient } from '../context'
import { Entity, EntityInfo, Repository } from '@apibrew/client'

export function useRepository<T extends Entity>(
  entityInfo: EntityInfo
): Repository<T> {
  const client = useClient()

  return client.repository<T>(entityInfo)
}
