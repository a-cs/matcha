import { Box, Button, Grid2, IconButton, InputAdornment, Paper, Stack, TextField, Typography } from "@mui/material";
import PersonAddIcon from '@mui/icons-material/PersonAdd';
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import { useState } from "react";
// import logo from "./logoMatcha.png"

export default function SignUp() {
	const [showPassword, setShowPassword] = useState(false)

	function togglePasswordVisibility() {
		setShowPassword(currentState => !currentState)
	}

	return (
		<Box m={4}>
			<Paper elevation={8} sx={{ padding: 2, borderRadius: '16px', minHeight: "92vh" }}>
				<Grid2 container spacing={2}>
					<Grid2 size={{ xs: 12, sm: 6 }}>
						<Stack alignItems={"center"} justifyContent={"space-around"} minHeight="100%" >
							<Box component="img" sx={{ width: "100%", height: "100%", maxWidth: "700px", maxHeight: "700px" }} src="./logoMatcha.png" />
						</Stack>
					</Grid2>
					<Grid2 size={{ xs: 12, sm: 6 }}>

						<Stack paddingTop={4} spacing={2} justifyContent={"space-between"} minHeight={"88vh"}>
							<Typography textAlign="center" component="h1" sx={{ typography: { xs: 'h2', md: 'h1' } }} color="primary.main"><b>Sign Up</b></Typography>
							<Grid2 container spacing={2}>
								<Grid2 size={{ xs: 12 }}>
									<TextField fullWidth required label="Email" />
								</Grid2>
								<Grid2 size={{ xs: 12, md: 6 }}>
									<TextField fullWidth required label="Username" />
								</Grid2>
								<Grid2 size={{ xs: 12, md: 6 }}>
									<TextField fullWidth required label="Password" type={showPassword ? "text" : "password"} slotProps={{
										input: {
											endAdornment: (
												<InputAdornment position="end">
													<IconButton size="large" onClick={togglePasswordVisibility}>

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
									<TextField fullWidth required label="First Name" />
								</Grid2>
								<Grid2 size={{ xs: 12, md: 6 }}>

									<TextField fullWidth required label="Last Name" />
								</Grid2>
							</Grid2>
							<Stack paddingTop={4} spacing={2} direction={"row"} justifyContent={"end"}>

								<Button variant="contained" color="secondary" href="/" ><b>Home</b></Button>
								<Button variant="contained" href="/signup" startIcon={<PersonAddIcon />} sx={{ borderRadius: '32px' }}><b>Sign up</b></Button>
							</Stack>
						</Stack>
					</Grid2>
				</Grid2>
			</Paper>
		</Box>
	)
}
