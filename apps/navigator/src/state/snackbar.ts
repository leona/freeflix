import { createContext } from "preact";

export default createContext({
  snackbars: [],
  add: (snackbar) => null,
});
