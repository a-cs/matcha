import { Route, Routes } from 'react-router-dom'

import Home from './pages/home'
import SignUp from './pages/signup'

export function Router() {
	return (
		<Routes>
			<Route index element={<Home />} />
			<Route path="/signup" element={<SignUp />} />
		</Routes>
	)
}