import {Permission} from "./model";

export interface TokenBody {
    exp: number
    uid: string
    username: string
    permissions: Permission[]
}