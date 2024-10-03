import { Box, Button, Container, Typography } from "@mui/material";

export default function SignUp() {
	return (
		<Container sx={{ background: "blue" }}>
			<Typography variant="h1">SignUp</Typography>
			<Box sx={{ my: 4 }}>
				<Button variant="contained" href="/" ><b>Home</b></Button>
				<Button variant="contained" color="secondary" href="/signup"><b>Signup</b></Button>
			</Box>
		</Container>
	)
}
