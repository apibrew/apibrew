import {OpenAPIV3_1} from "openapi-types";
import axios from 'axios'

export function get(): Promise<OpenAPIV3_1.Document> {
    return axios.get<OpenAPIV3_1.Document>('http://localhost:9009/docs/openapi.json').then(resp => resp.data)
}