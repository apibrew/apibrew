import { useEffect, useState } from 'react'
import { Urls } from '@apibrew/client/impl/client-impl'
import { Event } from '@apibrew/client/model/extension'
import { ApiException, EntityInfo } from '@apibrew/client'
import { useClient, useErrorHandler } from '../context'

export function useWatcher(
  entityInfo: EntityInfo,
  filters?: { [key: string]: string }
): number {
  const client = useClient()
  const errorHandler = useErrorHandler()

  const [counter, setCounter] = useState(0)

  const eiw = JSON.stringify(entityInfo)

  const triggerUpdate = () => {
    setCounter((counter) => counter + 1)
  }

  useEffect(() => {
    const controller = new AbortController()

    async function watchInternal() {
      while (!controller.signal.aborted) {
        try {
          console.log('Starting watch')
          const response = await fetch(
            Urls.recordWatchUrl(client.getUrl(), entityInfo.restPath, filters),
            {
              method: 'GET',
              headers: {
                'Content-Type': 'text/event-stream',
                ...client.headers()
              },
              signal: controller.signal
            }
          )

          if (!response.body) {
            if (controller.signal.aborted) {
              console.log('request aborted')
              break
            }

            console.log('watcher no content')
            break
          }

          const reader = response.body
            .pipeThrough(new TextDecoderStream())
            .getReader()
          while (true) {
            try {
              const { value, done } = await reader.read()
              if (done) break
              if (!value) {
                break
              }
              const event = JSON.parse(value) as Event

              if (response.status !== 200) {
                throw ApiException.fromError(event)
              }

              if (event.id === 'heartbeat-message') {
                // console.log('Received heartbeat event')
              } else {
                triggerUpdate()
              }
            } catch (e) {
              console.error(e)
              break
            }
          }
        } catch (e) {
          console.error(e)

          if (e instanceof DOMException) {
            return
          }

          if (errorHandler) {
            errorHandler(e)
          }
        } finally {
          await new Promise((resolve) => setTimeout(resolve, 1000))
        }
      }
    }

    watchInternal().then(() => {
      console.log('Watch completed')
    })

    return () => {
      console.log('Aborting watch')
      controller.abort()
    }
  }, [eiw])

  return counter
}
