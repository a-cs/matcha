import { Box, Button, Container, Typography } from "@mui/material";

export default function Home() {
	return (
		<Container sx={{ background: "red" }}>
			<Typography variant="h1">Home</Typography>
			<Box sx={{ my: 4 }}>
				<Button variant="contained" href="/" ><b>Home</b></Button>
				<Button variant="contained" color="secondary" href="/signup"><b>Signup</b></Button>
			</Box>
		</Container>
	)
}
