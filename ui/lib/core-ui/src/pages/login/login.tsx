import * as React from 'react'
import {useState, useEffect} from 'react'
import Avatar from '@mui/material/Avatar'
import Button from '@mui/material/Button'
import TextField from '@mui/material/TextField'
import Box from '@mui/material/Box'
import LockOutlinedIcon from '@mui/icons-material/LockOutlined'
import Typography from '@mui/material/Typography'
import Container from '@mui/material/Container'
import VisibilityIcon from '@mui/icons-material/Visibility'
import {useNavigate} from 'react-router-dom'
import * as TokenService from "../../service/token.ts";
import {authenticate} from "@apibrew/core-lib";

export function Login(): JSX.Element {
    const isLoggedIn = TokenService.isLoggedIn()
    useEffect(() => {
        console.log(isLoggedIn)
        if (isLoggedIn) {
            navigate('/dashboard/test')
        }
    }, [])

    const navigate = useNavigate()
    const [show, setShow] = useState('password')
    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        const data = new FormData(event.currentTarget)

        authenticate(data.get('username') as string, data.get('password') as string)
            .then(() => {
                navigate('/dashboard/test')
            }, (err: Error) => {
                console.error(err)

                alert('username or password is incorrect')
            })
    }

    return (
        <main
            style={{width: '450px', height: '500px', margin: '100px auto', borderRadius: '50px', textAlign: 'center'}}>
            <Container component="main" maxWidth="xs"
                       style={{borderRadius: '50px', height: '500px', width: '450px', background: 'none'}}>
                <Box
                    sx={{
                        marginTop: 2,
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center'
                    }}
                >
                    <Avatar sx={{m: 1, bgcolor: 'black'}}>
                        <LockOutlinedIcon/>
                    </Avatar>
                    <Typography component="h1" variant="h5" style={{color: 'white'}}>
                        Sign in
                    </Typography>
                    <Box component="form" onSubmit={handleSubmit} noValidate sx={{mt: 1}}>
                        <TextField
                            margin="normal"
                            required
                            fullWidth
                            id="username"
                            label="Username"
                            name="username"
                            autoComplete="username"
                            autoFocus
                        />
                        <button onClick={(e) => {
                            e.preventDefault()
                            e.stopPropagation()

                            if (show === 'password') {
                                setShow('text')
                            } else {
                                setShow('password')
                            }

                            return false
                        }}
                                style={{
                                    cursor: 'hand',
                                    position: 'relative',
                                    top: 60,
                                    marginLeft: 350,
                                    border: 0,
                                    borderRadius: 38,
                                    color: 'grey',
                                    width: '40px',
                                    height: '30px',
                                    zIndex: 1000,
                                    background: ' none'
                                }}>
                            <VisibilityIcon/>
                        </button>
                        <TextField
                            margin="normal"
                            required
                            fullWidth
                            name="password"
                            label="Password"
                            type={show}
                            id="password"
                            autoComplete="current-password"
                        />
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{mt: 3, mb: 2}}
                            style={{background: ' rgba(255, 255, 255, 0.15)'}}>
                            Sign In
                        </Button>
                    </Box>
                </Box>
            </Container>
        </main>
    )
}
