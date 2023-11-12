import { useClient } from '../context'
import { TokenBody } from '@apibrew/client/token-body'

export function useTokenBody(): TokenBody {
  const client = useClient()

  const token = client.getTokenBody()

  if (!token) {
    throw new Error('Token body not found')
  }

  return token
}
