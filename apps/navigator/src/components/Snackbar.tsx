import { useContext } from "preact/hooks";
import SnackbarState from '@/state/snackbar'

export function Snackbar() {
  const snackbars = useContext(SnackbarState)

  const colors = {
    'info': 'bg-blue-500',
    'success': 'bg-green-500',
    'error': 'bg-red-500',
  }

  return (
    <>
      {snackbars.snackbars.map((snackbar) => (
        <div class="fixed bottom-0 right-0 m-4 z-50">
          <div class={`${colors[snackbar.type] || colors.info} rounded-lg shadow-xl py-3 px-4 m-5 flex flex-row`}>
            <div class="flex flex-row space-x-2">
              <p class="text-sm font-semibold">{snackbar.type.toUpperCase()}</p>
              <p class="text-sm">{snackbar.message}</p>
            </div>
          </div>
        </div>
      ))}
    </>
  )
}