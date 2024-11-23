import { Box, Button, CircularProgress, Grid2, Paper, Stack, TextField, Typography } from "@mui/material";
import { FormEvent, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
// import { toast } from "react-toastify";
// import api from "../../services/api";
// import { isAxiosError } from "axios";
import LockResetIcon from '@mui/icons-material/LockReset';

export default function ForgottenPassword() {
	const navigate = useNavigate()

	const [email, setEmail] = useState('')
	const [loading, setLoading] = useState(false)



	// useEffect(() => {
	// 	async function confirmEmail(slugId: string) {
	// 		setMessage("Email confirmation in progress...")
	// 		try {
	// 			await api.post('/confirm', {
	// 				slugId
	// 			})
	// 			toast('Account activated!', {
	// 				type: 'success',
	// 				draggable: false,
	// 			})
	// 			navigate("/login")
	// 		} catch (error: unknown) {
	// 			if (isAxiosError(error) && error.response) {
	// 				toast(error.response.data.message, {
	// 					type: 'error',
	// 					draggable: false,
	// 				})
	// 			}
	// 			else {
	// 				toast((error as Error).message, {
	// 					type: 'error',
	// 					draggable: false,
	// 				})
	// 			}
	// 			setMessage("Invalid email confirmation link")
	// 		}
	// 	}

	// 	if (!params.slugId) {
	// 		toast('Invalid email confirmation link', {
	// 			type: 'error',
	// 			draggable: false,
	// 		})
	// 		navigate("/")
	// 	} else {
	// 		console.log("params.slugId:", params.slugId)
	// 		confirmEmail(params.slugId)

	// 	}
	// }, [params.slugId, navigate])

	function handleSubmit(e: FormEvent<HTMLFormElement>) {
		console.log("teste", e)
		setLoading(false)
	}

	return (
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
							<Typography textAlign="center" component="h1" sx={{ typography: { xs: 'h2', md: 'h1' } }} color="primary.main"><b>Forgotten Password</b></Typography>
							<Grid2 sx={{ justifySelf: "center" }} container spacing={2} >
								<Grid2 size={{ xs: 12 }}>
									<TextField fullWidth required label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
								</Grid2>
								<Grid2 size={{ xs: 12 }}>
									<Stack spacing={2} direction={"row"} justifyContent={"end"}>
										<Button type="submit" variant="contained" startIcon={!loading ? <LockResetIcon /> : ""} sx={{ borderRadius: '32px' }}>{!loading ? <b>Password</b> : <CircularProgress color="inherit" size={21} sx={{ my: "1.75px", mx: "29.82px" }} />}</Button>
									</Stack>
								</Grid2>
							</Grid2>
							<Grid2 container spacing={2} ></Grid2>
						</Stack>
					</Grid2>
				</Grid2>
			</Paper>
		</Box>
	)
}
