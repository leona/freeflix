import { useEffect, useState, useContext } from "preact/hooks";
import api from '@/models/api'
import Loader from '@/state/loader'
import Snackbar from '@/state/snackbar'
import { PrimaryBtn, DangerBtn } from '@/components/Button'

const ItemAction = ({ result }) => {
  const loader = useContext(Loader)
  const snackbar = useContext(Snackbar)

  const remove = async () => {
    loader.set(true)
    await api.remove(result.hash)
    await new Promise(r => setTimeout(r, 500));
    loader.set(false)
    
    snackbar.add({
      message: "Removed from download queue",
      type: 'info'
    })
  }

  const removeByTitle = async () => {
    loader.set(true)
    await api.removeByTitle(result.name)
    await new Promise(r => setTimeout(r, 500));
    loader.set(false)
    
    snackbar.add({
      message: "Files deleted",
      type: 'info'
    })
  }

  const watch = async () => {
    loader.set(true)

    try {
      const watch = await api.watch(result.name)
      window.open(watch.url, '_blank')
    } catch (err) {
      snackbar.add({
        message: err.message,
        type: 'error'
      })
    }

    loader.set(false)
  }

  if (result?.progress < 100) {
    return (
      <div class="flex flex-row">
        <DangerBtn onClick={remove}>Cancel</DangerBtn>
      </div>
    )
  } else {
    return (
      <div class="flex flex-row space-x-3">
        <PrimaryBtn onClick={watch}>Watch</PrimaryBtn>
        <DangerBtn onClick={removeByTitle}>Delete</DangerBtn>
      </div>
    )
  }
}

const Item = ({ result }) => {
  return (
    <div class="flex flex-row justify-between">
			<div class="flex flex-col">
				<p class="text-sm">{result.name}</p>
        {result.progress > 0 && (
          <div class="flex space-x-2">
            {result.progress && (<p class="text-xs">Progress: <strong>{result.progress}% - {result.downloaded} / {result.size}</strong></p>)}	
            {result.speed && (<p class="text-xs">Speed: <strong>{result.speed}/s</strong></p>)}
          </div>
        )}
			</div>
      <ItemAction result={result} />
    </div>
  )
}

export function Downloads() {
  const { query } = this.props
  const [data, setData] = useState([])

  const refresh = () => api.downloads().then((data) => setData(data))
  
  useEffect(() => {
    refresh()
    const interval = setInterval(refresh, 2000)

    return () => {
      clearInterval(interval)
    }
  }, [query])
  
  const combinedData = [...(data.active || []), ...(data.complete || [])]

  if (!combinedData?.length) {
    return (
      <div class="flex flex-col w-full">
        <p class="text-xl py-5 text-white">No downloads</p>
      </div>
    )
  }
  
	return (
		<div class="downloads flex-col text-left h-full w-full space-y-5 py-4">
        {combinedData.map((result) => (
          <Item result={result} />
        ))}
		</div>
	);
}
