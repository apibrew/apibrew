import { useEffect, useState } from 'react'
import { Urls } from '@apibrew/client/impl/client-impl'
import { Event } from '@apibrew/client/model/extension'
import { EntityInfo } from '@apibrew/client'
import { useClient } from '../context'

export function useWatcher(
  entityInfo: EntityInfo,
  filters?: { [key: string]: string }
): number {
  const client = useClient()

  const [counter, setCounter] = useState(0)

  const eiw = JSON.stringify(entityInfo)

  const triggerUpdate = () => {
    setCounter((counter) => counter + 1)
  }

  useEffect(() => {
    const controller = new AbortController()

    async function watchInternal() {
      while (!controller.signal.aborted) {
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
          throw new Error('No response body')
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

            if (event.id === 'heartbeat-message') {
              // console.log('Received heartbeat event')
            } else {
              triggerUpdate()
            }
          } catch (e) {
            console.error('Error reading data', e)
            break
          }
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
