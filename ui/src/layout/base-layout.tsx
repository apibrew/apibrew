import { createTheme, ThemeProvider } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'

export interface BaseLayoutProps {
    children: JSX.Element | JSX.Element[]
}

export function BaseLayout(props: BaseLayoutProps): JSX.Element {
    const theme = createTheme()

    return <>
        <ThemeProvider theme={theme}>
            <CssBaseline />
            {props.children}
        </ThemeProvider>
    </>
}
