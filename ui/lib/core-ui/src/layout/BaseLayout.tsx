import {createTheme, ThemeProvider} from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import React, {Fragment} from "react";
import {AlertOptions, LayoutContext, LayoutOptions, ModalOperations, ModalOptions} from "../context/layout-context.ts";
import {Alert, Modal, Snackbar} from "@mui/material";

export interface BaseLayoutProps {
    children: JSX.Element | JSX.Element[]
}

export function BaseLayout(props: BaseLayoutProps): JSX.Element {
    const [snackBarOpen, setSnackBarOpen] = React.useState(false)
    const [alert, setAlert] = React.useState<AlertOptions>()
    const [modals, setModals] = React.useState<ModalOptions[]>([])

    const layoutOptions: LayoutOptions = {
        showAlert(alert: AlertOptions): void {
            setAlert(alert)
            setSnackBarOpen(true)
        },
        showModal(modal: ModalOptions): ModalOperations {
            modal.id = Math.random().toString(36).substr(2, 9)

            setModals([...modals, modal])

            return {
                id: modal.id,
                close(): void {
                    setModals(modals.filter(m => m.id !== modal.id))
                    if (modal.onClose) {
                        modal.onClose()
                    }
                }
            }
        }
    }

    const modalContainer = (
        <React.Fragment>
            {modals.map(modal => <Fragment key={modal.id}>
                <Modal {...modal.props}
                       open={true}
                       onClose={() => {
                           setModals(modals.filter(m => m.id !== modal.id))
                           if (modal.onClose) {
                               modal.onClose()
                           }
                       }}>
                    <React.Fragment>
                        {modal.content}
                    </React.Fragment>
                </Modal>
            </Fragment>)}
        </React.Fragment>
    )
    const snackBar = <Snackbar open={snackBarOpen}
                               autoHideDuration={6000}
                               onClose={() => {
                                   setSnackBarOpen(false)
                               }}>
        <Alert onClose={() => {
            setSnackBarOpen(false)
        }} severity={alert?.severity} sx={{width: '100%'}}>{alert?.message}</Alert>
    </Snackbar>

    const theme = createTheme({
        palette: {
            mode: 'light',
            background: {
                default: '#F7F9FC'
            }
        },
        components: {
            MuiButton: {
                defaultProps: {
                    size: 'small',
                    sx: {
                        margin: 0.5
                    }
                },
            },
        }
    })
    return <React.Fragment>
        <ThemeProvider theme={theme}>
            <CssBaseline/>
            <LayoutContext.Provider value={layoutOptions}>
                {props.children}
                {modalContainer}
                {snackBar}
            </LayoutContext.Provider>
        </ThemeProvider>
    </React.Fragment>
}
