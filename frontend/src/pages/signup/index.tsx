import { Box, Button, CircularProgress, Grid2, IconButton, InputAdornment, Paper, Stack, TextField, Typography } from "@mui/material";
import PersonAddIcon from '@mui/icons-material/PersonAdd';
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import ArrowBackIosNewIcon from '@mui/icons-material/ArrowBackIosNew';
import { FormEvent, useState } from "react";
import api from "../../services/api";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";

export default function SignUp() {
	const navigate = useNavigate()

	const [showPassword, setShowPassword] = useState(false)
	const [email, setEmail] = useState('')
	const [username, setUsername] = useState('')
	const [password, setPassword] = useState('')
	const [firstName, setFirstName] = useState('')
	const [lastName, setLastName] = useState('')
	const [loading, setLoading] = useState(false)

	function togglePasswordVisibility() {
		setShowPassword(currentState => !currentState)
	}

	async function handleSubmit(e: FormEvent<HTMLFormElement>) {
		e.preventDefault()
		setLoading(true)
		try {
			await api.post('/signup', {
				email,
				username,
				password,
				firstName,
				lastName
			})
			toast("Account Created. An email link was sent to confirm your account.", {
				type: 'success',
				draggable: false,
			})
			navigate("/login")
			setLoading(false)
		} catch (error: unknown) {
			console.log("error:", error)
			toast((error as Error).message, {
				type: 'error',
				draggable: false,
			})
			setLoading(false)
		}
	}

	return (
		<>
			<Box m={4}>
				<Paper elevation={8} sx={{ padding: 2, borderRadius: '16px', minHeight: "92vh" }}>
					<Grid2 container spacing={2}>
						<Grid2 size={{ xs: 12, sm: 6 }} sx={{ position: "relative" }}>
							<IconButton size="large" href="/login" sx={{ position: "absolute", top: { xs: -20, md: 0 }, left: { xs: -20, md: 0 } }}>
								<ArrowBackIosNewIcon fontSize="inherit" />
							</IconButton>
							<Stack alignItems={"center"} justifyContent={"space-around"} minHeight="100%" >
								<Box component="img" sx={{ width: "100%", height: "100%", maxWidth: "700px", maxHeight: "700px" }} src="./MatchaLogoWithText.png" />
							</Stack>
						</Grid2>
						<Grid2 size={{ xs: 12, sm: 6 }}>
							<Stack component="form" onSubmit={(e) => handleSubmit(e)} paddingTop={4} spacing={2} justifyContent={"space-between"} minHeight={"88vh"}>
								<Typography textAlign="center" component="h1" sx={{ typography: { xs: 'h2', md: 'h1' } }} color="primary.main"><b>Sign Up</b></Typography>
								<Grid2 container spacing={2} >
									<Grid2 size={{ xs: 12 }}>
										<TextField fullWidth required label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
									</Grid2>
									<Grid2 size={{ xs: 12, md: 6 }}>
										<TextField fullWidth required label="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
									</Grid2>
									<Grid2 size={{ xs: 12, md: 6 }}>
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
									<Grid2 size={{ xs: 12, md: 6 }}>
										<TextField fullWidth required label="First Name" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
									</Grid2>
									<Grid2 size={{ xs: 12, md: 6 }}>
										<TextField fullWidth required label="Last Name" value={lastName} onChange={(e) => setLastName(e.target.value)} />
									</Grid2>
									<Grid2 size={{ xs: 12 }} textAlign={"end"}>
										<Button type="submit" variant="contained" startIcon={!loading ? <PersonAddIcon /> : ""} sx={{ borderRadius: '32px' }}>{!loading ? <b>Sign up</b> : <CircularProgress color="inherit" size={21} sx={{ my: "1.75px", mx: "29.82px" }} />}</Button>
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
