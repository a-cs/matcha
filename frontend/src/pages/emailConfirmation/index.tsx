import { Box, CircularProgress, Paper, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { toast } from "react-toastify";
import api from "../../services/api";
import { isAxiosError } from "axios";

export default function EmailConfirmation() {
	const params = useParams()
	const navigate = useNavigate()

	const [message, setMessage] = useState("Reading link information...")

	useEffect(() => {
		async function confirmEmail(slugId: string) {
			setMessage("Email confirmation in progress...")
			try {
				await api.post('/confirm', {
					slugId
				})
				toast('Account activated!', {
					type: 'success',
					draggable: false,
				})
				navigate("/login")
			} catch (error: unknown) {
				if(isAxiosError(error) && error.response){
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
				setMessage("Invalid email confirmation link")
			}
		}

		if (!params.slugId) {
			toast('Invalid email confirmation link', {
				type: 'error',
				draggable: false,
			})
			navigate("/")
		} else {
			console.log("params.slugId:", params.slugId)
			confirmEmail(params.slugId)

		}
	}, [params.slugId, navigate])

	return (
		<Box sx={{ m: { xs: 1, sm: 4 }, minHeight: "92vh", display: "flex", alignItems: "center", justifyContent: "center" }}>
			<Paper elevation={8} sx={{ padding: { xs: 2, sm: 4 }, borderRadius: '16px', minHeight: "80px", display: "flex", alignItems: "center", justifyContent: "center", backgroundColor: "primary.main", color: "secondary.contrastText" }}>
				<CircularProgress size="32px" sx={{ minWidth: "32px", minHeight: "32px" }} color="inherit" />
				<Typography sx={{ typography: { xs: "p", sm: "h5", md: "h4" }, ml: { xs: 2, sm: 4 } }}><b>{message}</b></Typography>
			</Paper>
		</Box>
	)
}
