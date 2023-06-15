import {Box, Table, TableBody, TableCell, TableHead, TableRow} from "@mui/material";

import {JSX} from "react";
import {NavigationTree} from "./NavigationTree.tsx";
import {NavigationItem, navigationItems} from "./navigation";
import Image from "../../proto/index.ts";

export function Sdk(): JSX.Element {

    console.log(Image.file.map(item => {
        console.log(item.name)
        console.log(item.bufExtension)
        console.log(item.options)
    }))

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
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>Parameter name</TableCell>
                        <TableCell>Type</TableCell>
                        <TableCell>Description</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    <TableRow>
                        <TableCell>username</TableCell>
                        <TableCell>string</TableCell>
                        <TableCell>Username of the user</TableCell>
                    </TableRow>
                    <TableRow>
                        <TableCell>password</TableCell>
                        <TableCell>string</TableCell>
                        <TableCell>Password of the user</TableCell>
                    </TableRow>
                    <TableRow>
                        <TableCell>term</TableCell>
                        <TableCell>enum</TableCell>
                        <TableCell>Term is to indicate how long the token will live. <br/> Values: short &gt; 1 min;
                            middle &gt; 2 hours</TableCell>
                    </TableRow>
                </TableBody>
            </Table>
            <h3>Response Parameters</h3>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>Parameter name</TableCell>
                        <TableCell>Type</TableCell>
                        <TableCell>Description</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    <TableRow>
                        <TableCell>username</TableCell>
                        <TableCell>string</TableCell>
                        <TableCell>Username of the user</TableCell>
                    </TableRow>
                    <TableRow>
                        <TableCell>password</TableCell>
                        <TableCell>string</TableCell>
                        <TableCell>Password of the user</TableCell>
                    </TableRow>
                    <TableRow>
                        <TableCell>term</TableCell>
                        <TableCell>enum</TableCell>
                        <TableCell>Term is to indicate how long the token will live. <br/> Values: short &gt; 1 min;
                            middle &gt; 2 hours</TableCell>
                    </TableRow>
                </TableBody>
            </Table>

            <h3>
                Examples
            </h3>

            {/*<h1>User APIs</h1>*/}
            {/*<h2>Country</h2>*/}
            {/*<h2>City</h2>*/}

            {/*END*/}
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
