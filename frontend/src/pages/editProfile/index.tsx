import { Box, Button, CircularProgress, Grid2, MenuItem, Paper, Stack, TextField, Typography } from "@mui/material";
import { FormEvent, useEffect, useState } from "react";
import EditIcon from '@mui/icons-material/Edit';
import AddPhotoAlternateIcon from '@mui/icons-material/AddPhotoAlternate';
import AccountBoxIcon from '@mui/icons-material/AccountBox';
import PhotoIcon from '@mui/icons-material/Photo';

export default function EditProfile() {
	const [firstName, setFirstName] = useState('')
	const [lastName, setLastName] = useState('')
	const [gender, setGender] = useState('')
	const [biography, setBiography] = useState('')
	const [loading, setLoading] = useState(false)

	const [genderOptions, setGenderOptions] = useState<string[]>([])
	// const pictureSrc = "https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"

	// const pictures = [
	// 	pictureSrc, pictureSrc, pictureSrc, pictureSrc
	// ]

	const formData = {
		"id": 123,
		"user_id": 123,
		"first_name": "any",
		"last_name": "any",
		// "location": "any" To DO,
		// "likes_counter": 2,
		"gender": "any",
		"tags_list": [
			"any1",
			"any2"
		],
		"biography": "any",
		"sexual_preference": [
			"any1",
			"any2"
		],
		"pictures": [
			"any1",
			"any2"
		],
		// "view_counter": 2,
		// "is_online": true,
		// "last_online_at": "2024-11-11T23:59:59Z",
		// "account_status": "active"
	}


	async function handleSubmit(e: FormEvent<HTMLFormElement>) {
		e.preventDefault()
		console.log("teste", formData)
		setLoading(true)
	}

	useEffect(() => {
		setGenderOptions(["male", "female", "non-binary"])
	}, [])


	return (
		<>
			<Box m={4}>
				<Paper elevation={8} sx={{ padding: 2, borderRadius: '16px', minHeight: "92vh" }}>
					<Grid2 container spacing={2}>
						<Grid2 size={{ xs: 12 }}>
							<Stack component="form" onSubmit={(e) => handleSubmit(e)} paddingTop={4} justifyContent={"space-between"} minHeight={"88vh"}>
								<Stack mb={4} spacing={{ xs: 1, md: 2 }}>
									<Typography textAlign="center" component="h1" sx={{ typography: { xs: 'h2', md: 'h1' } }} color="primary.main"><b>Profile</b></Typography>
									<Typography textAlign="center" component="h3" sx={{ typography: { xs: 'h4', md: 'h3' } }} color="#bdbdbd">Complete your profile information.</Typography>
								</Stack>
								<Grid2 container spacing={2} >
									<Grid2 size={{ xs: 12 }} textAlign={"center"}>
										<Button type="button" variant="contained" sx={{ fontSize: '32px', borderRadius: '8px', width: "132px", height: "132px", flexDirection: "column", justifyContent: "center", alignItems: "center" }}>{!loading ? <> <AddPhotoAlternateIcon sx={{ fontSize: '64px' }} /> <b>Photo</b> </> : <CircularProgress color="inherit" size={21} sx={{ my: "1.75px", mx: "29.82px" }} />}</Button>
									</Grid2>
									<Grid2 container columnGap={{ xs: "32px", md: "64px" }} rowGap={"24px"} size={{ xs: 12 }} direction={"row"} my={6} justifyContent={"center"}>
										<Grid2 size={{ xs: 12 }} display={"flex"} justifyContent={"center"}>
											<Paper elevation={8} sx={{ borderRadius: '16px', padding: "8px", width: "132px", height: "132px", display: "flex", alignItems: "center", justifyContent: "center", border: "2px solid", borderColor: "primary.main", position: "relative", borderBottomRightRadius: 0 }} title="Profile">
												<AccountBoxIcon sx={{
													fontSize: "22px", color: "primary.main", position: "absolute", bottom: -4, right: -4
												}} />
												<Box component={"img"} sx={{ borderRadius: '16px', height: "120px", border: "2px solid", borderColor: "#fff" }} src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D">
												</Box>
											</Paper>
										</Grid2>
										<Grid2 size={{ xs: 5, md: 1 }} display={"flex"} justifyContent={"center"}>
											<Paper elevation={4} sx={{ borderRadius: '12px', padding: "8px", width: "132px", height: "132px", display: "flex", alignItems: "center", justifyContent: "center", border:"2px solid", borderColor: "secondary.main", position: "relative", borderBottomRightRadius: 0 }}>
												<PhotoIcon sx={{
													fontSize: "16px", color: "secondary.main", position: "absolute", bottom: -2, right: -3
												}} />
												<Box component={"img"} sx={{ borderRadius: '16px', height: "120px", border: "2px solid", borderColor: "#ccc" }} src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D">
												</Box>
											</Paper>
										</Grid2>
										<Grid2 size={{ xs: 5, md: 1 }} display={"flex"} justifyContent={"center"}>
											<Paper elevation={4} sx={{ borderRadius: '12px', padding: "8px", width: "132px", height: "132px", display: "flex", alignItems: "center", justifyContent: "center", border:"2px solid", borderColor: "secondary.main", position: "relative", borderBottomRightRadius: 0 }}>
												<PhotoIcon sx={{
													fontSize: "16px", color: "secondary.main", position: "absolute", bottom: -2, right: -3
												}} />
												<Box component={"img"} sx={{ borderRadius: '16px', height: "120px", border: "2px solid", borderColor: "#ccc" }} src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D">
												</Box>
											</Paper>
										</Grid2>
										<Grid2 size={{ xs: 5, md: 1 }} display={"flex"} justifyContent={"center"}>
											<Paper elevation={4} sx={{ borderRadius: '12px', padding: "8px", width: "132px", height: "132px", display: "flex", alignItems: "center", justifyContent: "center", border:"2px solid", borderColor: "secondary.main", position: "relative", borderBottomRightRadius: 0 }}>
												<PhotoIcon sx={{
													fontSize: "16px", color: "secondary.main", position: "absolute", bottom: -2, right: -3
												}} />
												<Box component={"img"} sx={{ borderRadius: '16px', height: "120px", border: "2px solid", borderColor: "#ccc" }} src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D">
												</Box>
											</Paper>
										</Grid2>
										<Grid2 size={{ xs: 5, md: 1 }} display={"flex"} justifyContent={"center"}>
											<Paper elevation={4} sx={{ borderRadius: '12px', padding: "8px", width: "132px", height: "132px", display: "flex", alignItems: "center", justifyContent: "center", border:"2px solid", borderColor: "secondary.main", position: "relative", borderBottomRightRadius: 0 }}>
												<PhotoIcon sx={{
													fontSize: "16px", color: "secondary.main", position: "absolute", bottom: -2, right: -3
												}} />
												<Box component={"img"} sx={{ borderRadius: '16px', height: "120px", border: "2px solid", borderColor: "#ccc" }} src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D">
												</Box>
											</Paper>
										</Grid2>
									</Grid2>
								</Grid2>
								<Grid2 container spacing={2} >
									<Grid2 size={{ xs: 12, md: 4 }}>
										<TextField fullWidth required label="First Name" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
									</Grid2>
									<Grid2 size={{ xs: 12, md: 4 }}>
										<TextField fullWidth required label="Last Name" value={lastName} onChange={(e) => setLastName(e.target.value)} />
									</Grid2>
									<Grid2 size={{ xs: 12, md: 4 }}>
										<TextField fullWidth required select label="Gender" value={gender} onChange={(e) => setGender(e.target.value)}>
											{
												genderOptions?.map((option) => (
													<MenuItem key={option} value={option}>
														{option}
													</MenuItem>
												))
											}
										</TextField>
									</Grid2>
									<Grid2 size={{ xs: 12 }}>
										<TextField fullWidth required multiline
											rows={2} label="Biography" value={biography} onChange={(e) => setBiography(e.target.value)} />
									</Grid2>
									<Grid2 size={{ xs: 12 }} textAlign={"end"}>
										<Button type="submit" variant="contained" startIcon={!loading ? <EditIcon /> : ""} sx={{ borderRadius: '32px' }}>{!loading ? <b>Profile</b> : <CircularProgress color="inherit" size={21} sx={{ my: "1.75px", mx: "29.82px" }} />}</Button>
									</Grid2>
								</Grid2>
							</Stack>
						</Grid2>
					</Grid2>
				</Paper >
			</Box >
		</>
	)
}
