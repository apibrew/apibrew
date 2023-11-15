import { useEffect, useState } from 'react'
import { useErrorHandler } from '../context/error-handler'
import { useRepository } from './repository'
import { Entity, EntityInfo, ListRecordParams } from '@apibrew/client'

export function useRecordByName<
  T extends Entity & {
    name: string
  }
>(entityInfo: EntityInfo, name: string, wi?: number): T | undefined {
  return useRecordBy<T>(entityInfo, { name: name } as any, wi)
}

export function useRecordBy<T extends Entity>(
  entityInfo: EntityInfo,
  identifier: Partial<T>,
  wi?: number
): T | undefined {
  const repository = useRepository<T>(entityInfo)

  const [record, setRecord] = useState<T>()
  const errorHandler = useErrorHandler()

  useEffect(() => {
    repository.load(identifier).then(setRecord, errorHandler)
  }, [entityInfo.restPath, JSON.stringify(identifier), wi])

  return record
}

export function useRecords<T extends Entity>(
  entityInfo: EntityInfo,
  params?: ListRecordParams,
  wi?: number
) {
  const repository = useRepository<T>(entityInfo)

  const [records, setRecords] = useState<T[]>()
  const errorHandler = useErrorHandler()

  useEffect(() => {
    repository.list(params).then((response) => {
      setRecords(response.content)
    }, errorHandler)
  }, [entityInfo.restPath, JSON.stringify(params), wi])

  return records
}
