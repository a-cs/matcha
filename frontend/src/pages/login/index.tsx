import { Box, Button, CircularProgress, Grid2, IconButton, InputAdornment, Paper, Stack, TextField, Typography } from "@mui/material";
import PersonAddIcon from '@mui/icons-material/PersonAdd';
import LoginIcon from '@mui/icons-material/Login';
import LockResetIcon from '@mui/icons-material/LockReset';
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import { FormEvent, useState } from "react";
import api from "../../services/api";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import { isAxiosError } from "axios";

export default function Login() {
	const navigate = useNavigate()

	const [showPassword, setShowPassword] = useState(false)
	const [email, setEmail] = useState('')
	const [password, setPassword] = useState('')
	const [loading, setLoading] = useState(false)

	function togglePasswordVisibility() {
		setShowPassword(currentState => !currentState)
	}

	async function handleSubmit(e: FormEvent<HTMLFormElement>) {
		e.preventDefault()
		setLoading(true)
		try {
			await api.post('/login', {
				email,
				password,
			})
			navigate("/")
			setLoading(false)
		} catch (error: unknown) {
			if(isAxiosError(error) && error.response && error.response.data.message){
				toast(error.response.data.message, {
					type: 'error',
					draggable: false,
				})
			}
			else {
				toast((error as Error).message, {
					type: 'error',
					draggable: false,
				})
			}
			setLoading(false)
		}
	}

	return (
		<>
			<Box m={4}>
				<Paper elevation={8} sx={{ padding: 2, borderRadius: '16px', minHeight: "92vh" }}>
					<Grid2 container spacing={2}>
						<Grid2 size={{ xs: 12, sm: 6 }} sx={{ position: "relative" }}>
							<Stack alignItems={"center"} justifyContent={"space-around"} minHeight="100%" >
								<Box component="img" sx={{ width: "100%", height: "100%", maxWidth: "700px", maxHeight: "700px" }} src="./MatchaLogoWithText.png" />
							</Stack>
						</Grid2>
						<Grid2 size={{ xs: 12, sm: 6 }}>
							<Stack component="form" onSubmit={(e) => handleSubmit(e)} paddingTop={4} justifyContent={"space-between"} spacing={2} minHeight={"88vh"}>
								<Typography textAlign="center" component="h1" sx={{ typography: { xs: 'h2', md: 'h1' } }} color="primary.main"><b>Login</b></Typography>
								<Grid2 sx={{ justifySelf: "center" }} container spacing={2} >
									<Grid2 size={{ xs: 12, lg: 6 }}>
										<TextField fullWidth required label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
									</Grid2>
									<Grid2 size={{ xs: 12, lg: 6 }}>
										<TextField fullWidth required label="Password" value={password} onChange={(e) => setPassword(e.target.value)} type={showPassword ? "text" : "password"} slotProps={{
											input: {
												endAdornment: (
													<InputAdornment position="end">
														<IconButton tabIndex={-1} size="large" onClick={togglePasswordVisibility}>

															{showPassword ?
																<VisibilityIcon fontSize="inherit" /> : <VisibilityOffIcon fontSize="inherit" />
															}
														</IconButton>
													</InputAdornment>
												)
											}
										}} />
									</Grid2>
									<Grid2 size={{ xs: 12 }}>
										<Stack spacing={2} direction={"row"} justifyContent={"space-between"}>
											<Button type="button" variant="outlined" startIcon={<PersonAddIcon />} sx={{ borderWidth: "3px", borderRadius: '32px' }} href="/signup"><b>SignUp</b></Button>
											<Button type="button" variant="outlined" startIcon={<LockResetIcon />} sx={{ borderWidth: "3px", borderRadius: '32px' }} href="/forgottenPassword"><b>Password</b></Button>
											<Button type="submit" variant="contained" startIcon={!loading ? <LoginIcon /> : ""} sx={{ borderRadius: '32px' }}>{!loading ? <b>Login</b> : <CircularProgress color="inherit" size={21} sx={{ my: "1.75px", mx: "29.82px" }} />}</Button>
										</Stack>
									</Grid2>
								</Grid2>
								<Grid2 container spacing={2} ></Grid2>
							</Stack>
						</Grid2>
					</Grid2>
				</Paper>
			</Box>
		</>
	)
}
