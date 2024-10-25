import { BrowserRouter } from "react-router-dom"
import { Router } from "./Router"
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from "react-toastify";
import "./App.css"

function App() {

	return (
		<BrowserRouter>
			<Router />
			<ToastContainer theme="colored" />
		</BrowserRouter>
	)
}

export default App
