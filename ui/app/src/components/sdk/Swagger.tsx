import SwaggerUI from "swagger-ui-react"
import "swagger-ui-react/swagger-ui.css"
import { Resource } from "../../model"
import { SdkComponentProps } from "./Sdk"
import React, { useEffect } from "react"
import PropTypes from "prop-types"
import { BACKEND_URL } from "../../config"
import axios from "axios"
import { Box } from "@mui/material"
import { TokenService } from "../../service/token"

export default class BaseLayout extends React.Component<any, any> {

    static propTypes = {
        errSelectors: PropTypes.object.isRequired,
        errActions: PropTypes.object.isRequired,
        specSelectors: PropTypes.object.isRequired,
        oas3Selectors: PropTypes.object.isRequired,
        oas3Actions: PropTypes.object.isRequired,
        getComponent: PropTypes.func.isRequired
    }

    render() {
        let { errSelectors, specSelectors, getComponent } = this.props
        let VersionPragmaFilter = getComponent("VersionPragmaFilter")
        let Operations = getComponent("operations", true)
        let Models = getComponent("Models", true)
        let Row = getComponent("Row")
        let Col = getComponent("Col")
        let Errors = getComponent("errors", true)
        let isSwagger2 = specSelectors.isSwagger2()
        let isOAS3 = specSelectors.isOAS3()

        const isSpecEmpty = !specSelectors.specStr()

        const loadingStatus = specSelectors.loadingStatus()

        let loadingMessage = null

        if (loadingStatus === "loading") {
            loadingMessage = <div className="info">
                <div className="loading-container">
                    <div className="loading"></div>
                </div>
            </div>
        }

        if (loadingStatus === "failed") {
            loadingMessage = <div className="info">
                <div className="loading-container">
                    <h4 className="title">Failed to load API definition.</h4>
                    <Errors />
                </div>
            </div>
        }

        if (loadingStatus === "failedConfig") {
            const lastErr = errSelectors.lastError()
            const lastErrMsg = lastErr ? lastErr.get("message") : ""
            loadingMessage = <div className="info failed-config">
                <div className="loading-container">
                    <h4 className="title">Failed to load remote configuration.</h4>
                    <p>{lastErrMsg}</p>
                </div>
            </div>
        }

        if (!loadingMessage && isSpecEmpty) {
            loadingMessage = <h4>No API definition provided.</h4>
        }

        if (loadingMessage) {
            return <div className="swagger-ui">
                <div className="loading-container">
                    {loadingMessage}
                </div>
            </div>
        }

        const servers = specSelectors.servers()
        const schemes = specSelectors.schemes()

        return (
            <div style={{width: '100%'}} className='swagger-ui'>
                <VersionPragmaFilter isSwagger2={isSwagger2} isOAS3={isOAS3} alsoShow={<Errors />}>
                    <Row>
                        <Col mobile={12} desktop={12} >
                            <Operations />
                        </Col>
                    </Row>
                    <Row>
                        <Col mobile={12} desktop={12} >
                            <Models />
                        </Col>
                    </Row>
                </VersionPragmaFilter>
            </div>
        )
    }
}

const LayoutPlugin = () => {
    return {
        components: {
            BaseLayout: BaseLayout
        }
    }
}

export function Swagger(props: SdkComponentProps): JSX.Element {
    const [spec, setSpec] = React.useState<any>(null)

    useEffect(() => {
        if (props.resource) {
            axios.get(`${BACKEND_URL}/docs/api.json`)
                .then(res => res.data)
                .then(resp => {
                    for (let pathItem of Object.keys(resp.paths)) {
                        if (!pathItem.startsWith('/'+props.resource!.name)) {
                            delete resp.paths[pathItem]
                        }
                    }

                    delete resp.tags

                    for (let componentItem of Object.keys(resp.components.schemas)) {
                        if (componentItem == `item-${props.resource!.name}`) {
                            continue
                        }

                        if (componentItem == `list-${props.resource!.name}`) {
                            continue
                        }

                        delete resp.components.schemas[componentItem]
                    }

                    setSpec(resp)

                    resp.servers = [
                        {
                            url: `${BACKEND_URL}`
                        }
                    ]
                })
        }
    }, [props.resource])

    if (!spec) {
        return <div>Loading...</div>
    }
    
    return <SwaggerUI
        filter={true}
        plugins={[LayoutPlugin]}
        requestInterceptor={async (req: any) => {
            req.headers['Authorization'] = `Bearer ${await TokenService.get()}`
            return req
        }}
        layout="BaseLayout"
        spec={spec}/>
}