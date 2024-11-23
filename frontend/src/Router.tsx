import { Route, Routes } from 'react-router-dom'

import Home from './pages/home'
import SignUp from './pages/signup'
import EmailConfirmation from './pages/emailConfirmation'
import Login from './pages/login'
import EditProfile from './pages/editProfile'
import ForgottenPassword from './pages/forgottenPassword'

export function Router() {
	return (
		<Routes>
			<Route index element={<Home />} />
			<Route path="/signup" element={<SignUp />} />
			<Route path="/login" element={<Login />} />
			<Route path="/confirm/:slugId?" element={<EmailConfirmation />} />
			<Route path="/edit" element={<EditProfile />} />
			<Route path='/forgottenPassword' element={<ForgottenPassword/>} />
		</Routes>
	)
}