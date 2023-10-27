import './style.scss';
import { route } from 'preact-router';

import { useEffect, useState, useContext } from "preact/hooks";
import api from '../../models/api'
import { bytes } from '../../utilities/normalise'
import { SearchBtn, SecondaryBtn } from '../../components/Button'
import Loader from '../../state/loader.js';
import Search from '../../state/search.js';
import Snackbar from '../../state/snackbar.js';

const SearchInput = ({ value }) => {
	const search = useContext(Search)

	const onSubmit = (e) => {
		console.log("submit", search)
		e.preventDefault()
		route(`/?query=${search.query}`)
	}

	useEffect(() => {
		if (value) {
			search.setQuery(value)
			return
		}
	}, [value])

	return (
		<form class="w-full py-5" onSubmit={onSubmit}>   
			<label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Search</label>
			<div class="relative">
				<div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
					<svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
						<path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
					</svg>
				</div>
				<input type="search" ref={input => input && input.focus()} autoFocus value={search.query} onInput={(e) => search.setQuery(e.target.value)} id="default-search" class="block w-full p-4 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white border-transparent focus:border-transparent focus:outline-none" placeholder="Search tv/film e.g. The bear s01e02" required />
				<SearchBtn type="submit">Search</SearchBtn>
			</div>
	</form>
	)
}

const ItemAction = ({ result , downloads, onDownload}) => {
	const loader = useContext(Loader)
	const snackbar = useContext(Snackbar)

  const onClickDownload = async () => {
		loader.set(true)

		try {
			await api.queue(result.magnetUrl || result.downloadUrl)
			await new Promise(r => setTimeout(r, 500));

			snackbar.add({
				message: "Added to download queue",
				type: 'success'
			})
		} catch (err) {
			snackbar.add({
				message: err.message,
				type: 'error'
			})
		}

		loader.set(false)
		await new Promise(r => setTimeout(r, 3000));
		onDownload()
  }

	const combinedDownloads = [...(downloads.active || []), ...(downloads.complete || [])]
	const exists = combinedDownloads.find((download) => download.name.replace(/-/g, ' ') == result.title.replace(/-/g, ' '))

  if (exists) {
    return (
      <p>Already exists</p>
    )
  } else {
    return (
			<div class="flex">
				<SecondaryBtn onClick={onClickDownload}>Download</SecondaryBtn>
			</div>
    )
  }
}

const Item = ({ result, downloads, onDownload }) => {
  return (
    <div class="flex flex-row justify-between">
			<div class="flex flex-col">
				<a href={result.guid} target="_blank">
					<p class="text-sm">{result.title}</p>
				</a>
				<div class="flex space-x-2">
					<p class="text-xs">Size: <strong>{bytes(result.size, ["KB", "MB", "GB"])}</strong></p>		
					<p class="text-xs">Age: <strong>{result.age} days</strong></p>
					<p class="text-xs">Source: <strong>{result.indexer}</strong></p>
					<p class="text-xs">Seeders: <strong>{result.seeders}</strong></p>
				</div>
			</div>
      <ItemAction downloads={downloads} onDownload={onDownload} result={result}/>
    </div>
  )
}

const Results = ({ query }) => {
	const search= useContext(Search)
	const loader = useContext(Loader)
	const [downloads, setDownloads] = useState([])
	const refreshDownloads = () => api.downloads().then((data) => setDownloads(data))
	const snackbar = useContext(Snackbar)

	if (query?.length) {
		useEffect(() => {
			console.log("Loading results")
			loader.set(true)
			refreshDownloads()

			api.search(query).then((data) => {
				search.setResults(data)
				loader.set(false)
			}).catch((err) => {
				snackbar.add({
					message: err.message,
					type: 'error'
				})

				loader.set(false)
			})
		}, [query])
	}

	return (
		<div class="flex flex-col space-y-3 pb-4">
			{search.results?.length === 0 && (
        <p class="text-xl py-5 text-white">No results found</p>
			)}
			{search.results?.map((result, index) => (
				<Item result={result} downloads={downloads} onDownload={refreshDownloads} key={result.guid} />
			))}
		</div>
	)
}

export function Home() {
	const { query } = this.props

	return (
		<div class="home flex-col w-full">
			<SearchInput value={query} />
			<Results query={query} />
		</div>
	);
}
