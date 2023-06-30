import axios from "axios";
import { test1, test2 } from "./functions.js";

export function Test1(params) {
    return test1(params)
}

export function Test2(params) {
    return test2(params)
}
