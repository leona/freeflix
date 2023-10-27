import { createContext } from "preact";

export default createContext({
  query: null,
  results: null,
  setQuery: (query) => null,
  setResults: (results) => null,
});
