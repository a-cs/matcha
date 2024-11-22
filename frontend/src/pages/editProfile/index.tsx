import { Autocomplete, Box, Button, Chip, CircularProgress, Grid2, MenuItem, Paper, Stack, TextField, Typography } from "@mui/material";
import { FormEvent, SyntheticEvent, useEffect, useState } from "react";
import EditIcon from '@mui/icons-material/Edit';
import AccountBoxIcon from '@mui/icons-material/AccountBox';
import PhotoIcon from '@mui/icons-material/Photo';
import AddIcon from '@mui/icons-material/Add';

export default function EditProfile() {
	const [firstName, setFirstName] = useState('')
	const [lastName, setLastName] = useState('')
	const [gender, setGender] = useState('')
	const [biography, setBiography] = useState('')
	const [sexualPreferences, SetSexualPreferences] = useState<string[] | null>([])
	const [tags, SetTags] = useState<string[] | null>([])
	const [loading, setLoading] = useState(false)

	const [genderOptions, setGenderOptions] = useState<string[]>([])
	const [sexualPreferencesOptions, setSexualPreferencesOptions] = useState<string[]>([])
	const [tagsOptions, setTagsOptions] = useState<string[]>([])
	const [profileSrc, setpProfileSrc] = useState("")
	const [picturesSrc, setPicturesSrc] = useState<string[]>([])


	const profile = "https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
	const pictureSrc1 = "https://images.unsplash.com/photo-1476254592636-2a8e23b256fe?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
	const pictureSrc2 = "https://images.unsplash.com/photo-1474291102916-622af5ff18bb?q=80&w=2669&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
	const pictureSrc3 = "https://images.unsplash.com/photo-1525715843408-5c6ec44503b1?q=80&w=2670&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
	const pictureSrc4 = "https://images.unsplash.com/photo-1478144113946-d55adda4e24e?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"

	const maximumNumberOfOptions = 2

	const defautlPaperSxValues = {
		borderRadius: '8px',
		padding: "8px",
		display: "flex",
		alignItems: "center",
		justifyContent: "center",
		border: "2px solid",
		borderColor: "#ccc",
		color: "#ccc",
		position: "relative",
		// borderBottomRightRadius: 0,
		transition: "0.2s",
		"&:hover": {
			boxShadow: 12,
			borderColor: "primary.main",
			color: "primary.main",
		}
	}

	const profilePaperSxValues = {
		minWidth: "132px",
		width: "132px",
		minHeight: "132px",
		height: "132px",
		...defautlPaperSxValues
	}

	const picturesPaperSxValues = {
		minWidth: "100px",
		width: "100px",
		minHeight: "100px",
		height: "100px",
		...defautlPaperSxValues
	}

	let formData = {
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
		"profile_picture": "profileSrc"
		// "view_counter": 2,
		// "is_online": true,
		// "last_online_at": "2024-11-11T23:59:59Z",
		// "account_status": "active"
	}


	async function handleSubmit(e: FormEvent<HTMLFormElement>) {
		e.preventDefault()
		setLoading(true)
		formData = {
			"id": 123,
			"user_id": 123,
			"first_name": firstName,
			"last_name": lastName,
			// "location": "any" To DO,
			// "likes_counter": 2,
			gender,
			"tags_list": tags ? tags : [""],
			biography,
			"sexual_preference": sexualPreferences ? sexualPreferences : [""],
			"pictures": picturesSrc,
			"profile_picture": profileSrc
			// "view_counter": 2,
			// "is_online": true,
			// "last_online_at": "2024-11-11T23:59:59Z",
			// "account_status": "active"
		}
		console.log("teste", formData)
		setLoading(false)
	}

	function handleSexualPrefereceChange(_event: SyntheticEvent<Element, Event>, newValue: string[] | null) {
		SetSexualPreferences(newValue)
	}

	function handleTagsChange(_event: SyntheticEvent<Element, Event>, newValue: string[] | null) {
		if (tags && tags?.length < maximumNumberOfOptions || newValue && newValue.length < maximumNumberOfOptions) {
			SetTags(newValue)
		}
	}

	function handleChangeProfilePicture() {
		//todo
		console.log("change profile pic")
	}

	function handleChangePictureByIndex(index: number) {
		//todo
		console.log("change profile pic index: ", index)
	}

	function handleAddPicture() {
		//todo
		console.log("add profile pic")
	}

	useEffect(() => {
		setGenderOptions(["male", "female", "non-binary"])
		setPicturesSrc([
			pictureSrc1,
			pictureSrc2,
			pictureSrc3,
			pictureSrc4
		])
		setpProfileSrc(profile)
		setSexualPreferencesOptions(["asexual", "bisexual", "gay", "lesbian", "straight"])
		setTagsOptions(["vegan", "geek", "piercing"])
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
									<Grid2 container columnGap={{ xs: "32px", md: "64px" }} rowGap={"24px"} size={{ xs: 12 }} direction={"row"} my={6} justifyContent={"center"}>
										<Grid2 size={{ xs: 12 }} display={"flex"} justifyContent={"center"}>
											<Paper component={Button} onClick={handleChangeProfilePicture} elevation={4} sx={profilePaperSxValues} title="Profile">
												<Box sx={{ backgroundColor: "transparent", position: "absolute", bottom: -5, right: -4, display: "flex", justifyContent: "center", alignItems: "center", borderRadius: "4px" }}>
													<AccountBoxIcon sx={{
														fontSize: "24px", color: "inherit", borderBottomRightRadius: 16
													}} />
												</Box>
												{profileSrc ?
													(<>
														<Box component={"img"} sx={{ borderRadius: '8px', maxWidth: "110px", height: "110px", border: "1px solid", borderColor: "#ccc" }} src={profileSrc}>
														</Box>
													</>
													) : (
														<>
															<Box sx={{ display: "flex", justifyContent: "center", alignItems: "center", borderRadius: "4px", color: "inherit" }}>
																<AddIcon sx={{
																	fontSize: "64px", color: "inherit",
																}} />
															</Box>
														</>
													)

												}
											</Paper>
										</Grid2>
										{
											picturesSrc?.map((src: string, index) => (
												<Grid2 key={index} size={{ xs: 5, md: 1 }} display={"flex"} justifyContent={"center"}>
													<Paper component={Button} onClick={() => handleChangePictureByIndex(index)} elevation={4} sx={picturesPaperSxValues}>
														<Box sx={{ backgroundColor: "transparent", position: "absolute", bottom: -5, right: -4, display: "flex", justifyContent: "center", alignItems: "center", borderRadius: "4px" }}>
															<PhotoIcon sx={{
																fontSize: "24px", color: "inherit", borderBottomRightRadius: 16
															}} />
														</Box>
														<Box component={"img"} sx={{ borderRadius: '8px', maxWidth: "84px", height: "84px", border: "1px solid", borderColor: "#ccc" }} src={src}>
														</Box>
													</Paper>
												</Grid2>
											))
										}
										{
											picturesSrc.length < 4 &&
											(
												<Grid2 size={{ xs: 5, md: 1 }} display={"flex"} justifyContent={"center"}>
													<Paper component={Button} onClick={handleAddPicture} elevation={4} sx={picturesPaperSxValues}>
														<Box sx={{ backgroundColor: "transparent", position: "absolute", bottom: -5, right: -4, display: "flex", justifyContent: "center", alignItems: "center", borderRadius: "4px" }}>
															<PhotoIcon sx={{
																fontSize: "24px", color: "inherit", borderBottomRightRadius: 16
															}} />
														</Box>
														<Box sx={{ display: "flex", justifyContent: "center", alignItems: "center", borderRadius: "4px", color: "inherit" }}>
															<AddIcon sx={{
																fontSize: "64px", color: "inherit",
															}} />
														</Box>
													</Paper>
												</Grid2>
											)
										}
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
									<Grid2 size={{ xs: 12, md: 6 }}>
										<Autocomplete
											disablePortal
											multiple
											options={sexualPreferencesOptions}
											renderInput={(params) => (
												<TextField label="Sexual Preference" {...params} />
											)}
											renderTags={(tagValue, getTagProps) =>
												tagValue.map((option, index) => {
													const { key, ...tagProps } = getTagProps({ index });
													return <Chip key={key} variant="outlined" color="primary" label={option} {...tagProps} />;
												})
											}
											value={sexualPreferences ? sexualPreferences : [""]}
											onChange={handleSexualPrefereceChange}
										/>
									</Grid2>
									<Grid2 size={{ xs: 12, md: 6 }}>
										<Autocomplete
											disablePortal
											multiple
											options={tagsOptions}
											renderInput={(params) => (
												<TextField label="Interests" helperText={
													tags && tags.length >= 2 ? `You can select a maximum of ${maximumNumberOfOptions} options.` : ""
												} {...params} />
											)}
											renderTags={(tagValue, getTagProps) =>
												tagValue.map((option, index) => {
													const { key, ...tagProps } = getTagProps({ index });
													return <Chip key={key} variant="outlined" color="primary" label={option} {...tagProps} />;
												})
											}
											value={tags ? tags : [""]}
											onChange={handleTagsChange}
										/>
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
