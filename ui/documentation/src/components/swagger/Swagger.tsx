import React, { useContext } from "react"
import SwaggerUI from "swagger-ui-react"
import "swagger-ui-react/swagger-ui.css"

const KeepOriginal = React.createContext({})


// Create the layout component
class AugmentingLayout extends React.Component {
    render() {
        const {
            getComponent
        } = this.props

        const BaseLayout = getComponent("BaseLayout", true)
        const originalOperation = getComponent("operation", true)

        console.log(SwaggerUI)

        return (
            <div>
                <div className="myCustomHeader">
                    <h1>I have a custom header above Swagger-UI!</h1>
                </div>
                <KeepOriginal.Provider value={
                    {
                        originalOperation: originalOperation
                    }
                }>
                    <BaseLayout />
                </KeepOriginal.Provider>
            </div>
        )
    }
}

function OperationsNew(props: any) {
    const getComponent = props.getComponent

    return <>
        {/* <Operations {...props} /> */}
    </>
}

// Create the plugin that provides our layout component
const AugmentingLayoutPlugin = () => {
    return {
        components: {
            AugmentingLayout: AugmentingLayout,
        },
        wrapComponents: {
            operation: (Original, { React }) => props => {
                const { operation } = props
                // if (
                //     operation.get("path") === "/pet/findByStatus" &&
                //     operation.get("method") === "get" &&
                //     operation.get("op").size // i.e., resolved Operation has been provided
                // ) {
                const originalDescription = operation.getIn(["op", "description"])
                return React.createElement(Original, {
                    ...props,
                    operation: operation.setIn(
                        ["op", "description"],
                        originalDescription + "\n\n*Hello world!*"
                    )
                })
                // }
                return React.createElement(Original, props)
            }
        }
    }
}

// Add config to Request Snippets Configuration with an unique key like "node_native" 
const snippetConfig = {
    requestSnippetsEnabled: true,
    requestSnippets: {
        generators: {
            "node_native": {
                title: "NodeJs Native",
                syntax: "javascript"
            }
        }
    }
}

const SnippedGeneratorNodeJsPlugin = {
    fn: {
        // use `requestSnippetGenerator_` + key from config (node_native) for generator fn
        requestSnippetGenerator_node_native: (request: any) => {
            const url = new URL(request.get("url"))
            let isMultipartFormDataRequest = false
            const headers = request.get("headers")
            if (headers && headers.size) {
                request.get("headers").map((val: any, key: any) => {
                    isMultipartFormDataRequest = isMultipartFormDataRequest || /^content-type$/i.test(key) && /^multipart\/form-data$/i.test(val)
                })
            }
            const packageStr = url.protocol === "https:" ? "https" : "http"
            let reqBody = request.get("body")
            if (request.get("body")) {
                if (isMultipartFormDataRequest && ["POST", "PUT", "PATCH"].includes(request.get("method"))) {
                    return "throw new Error(\"Currently unsupported content-type: /^multipart\\/form-data$/i\");"
                } else {
                    // if (!Map.isMap(reqBody)) {
                    //     if (typeof reqBody !== "string") {
                    //         reqBody = JSON.stringify(reqBody)
                    //     }
                    // } else {
                    //     reqBody = getStringBodyOfMap(request)
                    // }

                    reqBody = '{"t": "Taleh"}'
                }
            } else if (!request.get("body") && request.get("method") === "POST") {
                reqBody = ""
            }

            const stringBody = "`" + (reqBody || "")
                .replace(/\\n/g, "\n")
                .replace(/`/g, "\\`")
                + "`"

            return `const http = require("${packageStr}");
const options = {
  "method": "${request.get("method")}",
  "hostname": "${url.host}",
  "port": ${url.port || "null"},
  "path": "${url.pathname}"${headers && headers.size ? `,
  "headers": {
    ${request.get("headers").map((val: any, key: any) => `"${key}": "${val}"`).valueSeq().join(",\n    ")}
  }` : ""}
};
const req = http.request(options, function (res) {
  const chunks = [];
  res.on("data", function (chunk) {
    chunks.push(chunk);
  });
  res.on("end", function () {
    const body = Buffer.concat(chunks);
    console.log(body.toString());
  });
});
${reqBody ? `\nreq.write(${stringBody});` : ""}
req.end();`
        }
    }
}




export function Swagger() {
    return <SwaggerUI url="http://localhost:9009/docs/openapi.json"
        presets={[AugmentingLayoutPlugin]}
        plugins={[SnippedGeneratorNodeJsPlugin]}
        layout='AugmentingLayout'
        {...snippetConfig}
    />
}