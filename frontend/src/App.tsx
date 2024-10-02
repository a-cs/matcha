import { Box, Button, Container } from "@mui/material"

function App() {

	return (
		<Container maxWidth="sm">
			<Box sx={{ my: 4 }}>
				<Button variant="contained" ><b>Home</b></Button>
				<Button variant="contained" color="secondary" ><b>Home</b></Button>
				<Button variant="contained" color="error" ><b>Home</b></Button>
				<Button variant="contained" color="success" ><b>Home</b></Button>
			</Box>
		</Container>
	)
}

export default App
