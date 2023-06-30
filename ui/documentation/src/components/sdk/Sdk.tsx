import { Box } from "@mui/material";

import { JSX, useEffect, useState } from "react";
import { NavigationTree } from "./NavigationTree.tsx";
import { NavigationItem, navigationItems } from "./navigation";
import { Loading } from "@apibrew/ui-lib";
import { OpenAPIV3_1 } from "openapi-types";
import { OpenApiService } from '../../service'
import { useErrorHandler } from "@apibrew/ui-lib";
import { SchemaElement } from "./SchemaElement.tsx";
import { relatedSchemas } from "../../util/openapi.ts";

export function Sdk(): JSX.Element {
    const [doc, setDoc] = useState<OpenAPIV3_1.Document>()
    const errorHandler = useErrorHandler()

    useEffect(() => {
        OpenApiService.get().then(setDoc, errorHandler)
    }, [])

    if (!doc) {
        return <Loading />
    }

    return <Box className='documentation' display={'flex'} flexDirection={'row'} height={'100%'} overflow='hidden'>
        <Box display={'flex'} flexDirection={'column'}
            sx={{ background: 'white', width: '300px', height: '100%', padding: '5px' }}>
            <NavigationTree items={navigationItems} />
        </Box>
        <Box display={'block'} m={5} overflow='scroll'>
            <a target="_blank" href="http://localhost:9009/docs/openapi.json">openapi.json</a>

            {navigationItems.map(item => <SdkNavigationElement key={item.name} item={item} />)}

            <h1>Authentication API</h1>
            <p>
                Authentication APIs are used to authenticate users and get access to the resources.
            </p>
            <p>
                For all endpoints, which needs you to be authenticated, you need to pass the access token in the header.
                The access token is obtained by calling the authenticate endpoint.
            </p>
            <br />
            <h2>Authenticate</h2>
            <div style={{
                paddingLeft: '15px'
            }}>
                <p>
                    This endpoint is used to authenticate the user and get the access token.
                </p>
                <p>
                    The access token is used to authenticate the user for all the endpoints which needs authentication.
                </p>

                <SchemaElement open={true} doc={doc} name='AuthenticationRequest' />
                <SchemaElement open={true} doc={doc} name='AuthenticationResponse' />
            </div>

            <h2>Refresh Token</h2>
            <div style={{
                paddingLeft: '15px'
            }}>
                <p>
                    This endpoint is used to authenticate the user and get the access token.
                </p>
                <p>
                    The access token is used to authenticate the user for all the endpoints which needs authentication.
                </p>

                <SchemaElement open={true} doc={doc} name='AuthenticationRequest' />
                <SchemaElement open={true} doc={doc} name='AuthenticationResponse' />
            </div>

            <h2>Related Schemas</h2>
            {relatedSchemas(doc, ['AuthenticationRequest', 'AuthenticationResponse']).map(item => {
                return <SchemaElement open={false} key={item} name={item} doc={doc} />
            })}
        </Box>
    </Box>
}

function SdkNavigationElement(props: { item: NavigationItem }) {
    if (props.item.component) {
        const Component = props.item.component

        return <Component />
    }

    return <>
        {props.item.children?.map(item => <SdkNavigationElement key={item.name} item={item} />)}
    </>
}
