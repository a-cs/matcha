import { createTheme } from '@mui/material/styles';

export const theme = createTheme({
	palette: {
		mode: 'light',
		primary: {
			main: '#820ad1',
		},
		secondary: {
			main: '#20a4f3',
			contrastText: 'rgba(255,255,255,0.87)',
		},
		background: {
			default: '#f6f7f8',
			paper: '#f6f7f8',
		},
		error: {
			main: '#ff3c38',
		},
		success: {
			main: '#0ead69',
			contrastText: 'rgba(255,255,255,0.87)',
		},
		info: {
			main: '#20a4f3',
		},
		warning: {
			main: '#fd7506',
		},
	},
});