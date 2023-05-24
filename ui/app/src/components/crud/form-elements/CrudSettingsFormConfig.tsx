import {CrudFormConfig} from "../../../model/schema";
import {Card, CardContent, CardHeader, Grid, Tab, Tabs} from "@mui/material";
import Box from "@mui/material/Box";
import React from "react";
import Button from "@mui/material/Button";
import {CameraRear, ContentCopy, Delete, Group, TableRows, TextFields, Title} from "@mui/icons-material";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";

export interface CrudSettingsFormConfigProps {
    config: CrudFormConfig
}

export function CrudSettingsFormConfig(props: CrudSettingsFormConfigProps) {
    return (
        <>
            <Box sx={{flexGrow: 1}} m={2}>
                <Grid container spacing={2}>
                    <Grid xs={2}>
                        <Box m={1}>
                            <Card sx={{height: '600px'}}>
                                <CardHeader title={<>
                                    <center>Toolbox</center>
                                </>}/>
                               <CardContent>
                                   <List disablePadding={true}>
                                       <ListItem disablePadding={true}>
                                           <Button startIcon={<Title/>}>Title</Button>
                                       </ListItem>
                                       <ListItem disablePadding={true}>
                                           <Button startIcon={<ContentCopy/>}>Text</Button>
                                       </ListItem>
                                       <ListItem disablePadding={true}>
                                           <Button startIcon={<TextFields/>}>Text Field</Button>
                                       </ListItem>
                                       <ListItem disablePadding={true}>
                                           <Button startIcon={<TableRows/>}>Tab</Button>
                                       </ListItem>
                                       <ListItem disablePadding={true}>
                                           <Button startIcon={<Group/>}>Group</Button>
                                       </ListItem>
                                       <ListItem disablePadding={true}>
                                           <Button startIcon={<CameraRear/>}>Section</Button>
                                       </ListItem>
                                   </List>
                               </CardContent>
                            </Card>
                        </Box>
                    </Grid>
                    <Grid xs={10}>
                        <Box m={1}>
                            <Card sx={{height: '600px'}}>
                                <CardHeader title={<>
                                    <center>Form</center>
                                </>}/>
                                <CardContent>
                                    <Box sx={{borderBottom: 1, borderColor: 'divider'}}>
                                        <Tabs>
                                        </Tabs>
                                    </Box>
                                </CardContent>
                            </Card>
                        </Box>
                    </Grid>
                </Grid>
            </Box>
        </>
    )
}