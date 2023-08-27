import { NodeVM } from "vm2"
import { moduleInitTemplate } from "./function-template"
import { Module } from "@apibrew/client"
import * as fs from 'fs'
import path from "path"
import { FN_DIR } from "./config"
import { apbrClient } from "./client"

export async function initModule(module: Module) {
    const moduleInitContent = moduleInitTemplate(module)

    fs.writeFileSync(path.join(FN_DIR + '/', module.id + '/module_init.js'), moduleInitContent)

    const vm = new NodeVM({
        sandbox: {
            module: module,
            exports: exports,
            apbrClient: apbrClient
        },
        console: 'inherit',
        require: {
            external: true,
            builtin: ['*'],
        },
        allowAsync: true,
        timeout: 1000,
    })

    vm.runFile(FN_DIR + `/${module.id}/module_init.js`)
}
