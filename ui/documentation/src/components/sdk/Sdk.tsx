import {Box} from "@mui/material";

import {JSX} from "react";
import {NavigationTree} from "./NavigationTree.tsx";
import {NavigationItem, navigationItems} from "./navigation";
import {getMessageTypeDescriptor, getProtoDescriptor, GetProtoImage} from "../../proto";
import {ProtoMessageElement} from "./ProtoMessageElement.tsx";
import {useProtoImageLoading} from "../../proto/loaded-hook.ts";
import {Loading} from "@apibrew/core-lib";

export function Sdk(): JSX.Element {
    const loading = useProtoImageLoading()

    if (!loading) {
        return <Loading/>
    }

    const authenticationStubProto = getProtoDescriptor('stub', 'authentication');
    const AuthenticationRequest = getMessageTypeDescriptor(authenticationStubProto, 'AuthenticationRequest');
    const AuthenticationResponse = getMessageTypeDescriptor(authenticationStubProto, 'AuthenticationResponse');

    const ProtoImage = GetProtoImage()

    return <Box className='documentation' display={'flex'} flexDirection={'row'} height={'100%'} overflow='hidden'>
        <Box display={'flex'} flexDirection={'column'}
             sx={{background: 'white', width: '300px', height: '100%', padding: '5px'}}>
            <NavigationTree items={navigationItems}/>
        </Box>
        <Box display={'block'} m={5} overflow='scroll'>
            {navigationItems.map(item => <SdkNavigationElement item={item}/>)}

            {/*Begin*/}
            <h1>Authentication API</h1>
            <p>
                Authentication APIs are used to authenticate users and get access to the resources.
            </p>
            <p>
                For all endpoints, which needs you to be authenticated, you need to pass the access token in the header.
                The access token is obtained by calling the authenticate endpoint.
            </p>
            <br/>
            <h2>Authenticate</h2>
            <p>
                This endpoint is used to authenticate the user and get the access token.
            </p>
            <p>
                The access token is used to authenticate the user for all the endpoints which needs authentication.
            </p>
            <h3>Request Parameters</h3>
            <ProtoMessageElement open={true} messageType={AuthenticationRequest} protoFile={authenticationStubProto}/>
            <h3>Response Parameters</h3>
            <ProtoMessageElement open={true} messageType={AuthenticationResponse} protoFile={authenticationStubProto}/>

            <h3>
                Examples
            </h3>

            {/*<h1>User APIs</h1>*/}
            {/*<h2>Country</h2>*/}
            {/*<h2>City</h2>*/}

            {/*END*/}


            <h2>Schema</h2>
            <h3>Stub</h3>
            {ProtoImage.file.filter(item => item.package === 'stub').map(file => <>
                {(file.messageType ?? []).map(messageType => <>
                    <ProtoMessageElement messageType={messageType} protoFile={file}/>
                </>)}
            </>)}
            <h3>Model</h3>
            {ProtoImage.file.filter(item => item.package === 'model').map(file => <>
                {(file.messageType ?? []).map(messageType => <>
                    <ProtoMessageElement messageType={messageType} protoFile={file}/>
                </>)}
            </>)}
        </Box>
    </Box>
}

function SdkNavigationElement(props: { item: NavigationItem }) {
    if (props.item.component) {
        const Component = props.item.component

        return <Component/>
    }

    return <>
        {props.item.children?.map(item => <SdkNavigationElement item={item}/>)}
    </>
}
