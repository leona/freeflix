import { createContext } from "preact";

export default createContext({
  loading: false,
  set: (state) => null,
});
