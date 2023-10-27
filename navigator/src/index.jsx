import { render } from 'preact';
//import { LocationProvider, Router, Route } from 'preact-iso';
import Router from 'preact-router';
import Loader from './state/loader.js';
import Search from './state/search.js';
import Snackbar from './state/snackbar.js';
import { useEffect, useState } from "preact/hooks";
import Spinner from './components/Spinner.jsx';

import { Header } from './components/Header.jsx';
import { Home } from './pages/Home/index.jsx';
import { Downloads } from './pages/Downloads/index.jsx';
import { Login } from './pages/Login/index.jsx';
import { NotFound } from './pages/_404.jsx';
import './style.css';
import api from './models/api.js';
import { Snackbar as SnackbarComponent } from './components/Snackbar.jsx';

const PageLoader = () => {
	return (
		<div class="absolute w-screen flex justify-center items-center z-50" style={{ height: 'calc(100vh - 50px)' }}>
			<div class="absolute">
				<Spinner size={20}/>
			</div>
		</div>
	)
}
export function App() {
	const [loading, setLoading] = useState(false)
	const [searchQuery, setSearchQuery] = useState()
	const [searchResults, setSearchResults] = useState()
	const [snackbars, setSnackbars] = useState([])

	const addSnackbar = (message) => {
		setSnackbars([...snackbars, message])
		setTimeout(() => {
			setSnackbars(snackbars.slice(1))
		}, 5000)
	}

	useEffect(() => {
		const urlParts = new URL(window.location)
		const url = urlParts.pathname

		api.checkAuth().catch(() => {
			if (url !== '/login') {
				window.location.href = "/login"
			}
		})
	}, [])

	return (
		<Loader.Provider value={{ loading, set: setLoading }}>
			<Snackbar.Provider value={{ snackbars, add: addSnackbar }}>
				<Search.Provider value={{  query: searchQuery, results: searchResults, setQuery: setSearchQuery, setResults: setSearchResults }}>
					<Header />
					<SnackbarComponent />
					<div style={{ marginTop: '50px'}}>
						{loading && <PageLoader />}
						<main class="flex flex-co text-left" style={{ opacity: loading ? 0.5 : 1 }}>
							<Router>
								<Home path="/" />
								<Login path="/login" />
								<Downloads path="/downloads" />
								<NotFound default />
							</Router>
						</main>
					</div>
				</Search.Provider>
			</Snackbar.Provider>
		</Loader.Provider>
	);
}

render(<App />, document.getElementById('app'));
