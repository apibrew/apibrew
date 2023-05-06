import axios from 'axios'
import {BACKEND_URL} from '../config'
import {Resource} from "../model";
import {TokenService} from "./token";

export namespace ResourceService {
    interface ResourceListContainer {
        resources: Resource[]
    }

    export async function list(): Promise<Resource[]> {
        const result = await axios.get<ResourceListContainer>(`${BACKEND_URL}/system/resources`, {
            headers: {
                Authorization: `Bearer ${await TokenService.get()}`
            }
        })

        return result.data.resources
    }
}
