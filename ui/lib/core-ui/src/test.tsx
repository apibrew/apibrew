const code = `
Object.defineProperty(exports, Symbol.toStringTag, { value: "Module" });
const jsxRuntime = require("react/jsx-runtime");
const coreUi = require("core-ui");
function Component1() {
  return /* @__PURE__ */ jsxRuntime.jsx(jsxRuntime.Fragment, { children: /* @__PURE__ */ jsxRuntime.jsx(coreUi.PageLayout, { children: /* @__PURE__ */ jsxRuntime.jsx("h1", { children: "Hello World from Component1" }) }) });
}
exports.Component1 = Component1;
//# sourceMappingURL=skeleton.cjs.js.map

`

import * as jsxRuntime from 'react/jsx-runtime'
import * as CoreUI from './index.ts'

const func = new Function('exports', 'require', code)

const exports: {[key: string]: any} = {}

func(exports, (path: string) => {
    switch (path) {
        case 'react/jsx-runtime':
            return jsxRuntime
        case 'core-ui':
            return CoreUI
    }
})

const Component1 = exports.Component1

console.log(exports)

export function Test() {
    return <div>
        <Component1/>
    </div>
}
