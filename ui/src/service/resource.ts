import axios from 'axios'
import { BACKEND_URL } from '../config'
import { type Resource } from '../model'
import { TokenService } from './token'
import { handleError } from './error-handler'

export namespace ResourceService {
    interface ResourceListContainer {
        resources: Resource[]
    }

    export async function list(): Promise<Resource[]> {
        try {
            const result = await axios.get<ResourceListContainer>(`${BACKEND_URL}/system/resources`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data.resources
        } catch (e) {
            return await handleError(e)
        }
    }
}
