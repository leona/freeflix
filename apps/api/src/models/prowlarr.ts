import { prowlarr } from "../config.js";

const baseUrl = `http://prowlarr:${prowlarr.Port}/api/v1`;

const search = async ({ query }) => {
  console.log("searching for:", query);

  const response = await fetch(
    `${baseUrl}/search?apikey=${prowlarr.ApiKey}&query=${query}`
  );

  let data = await response.json();
  data = data.sort((a, b) => b.seeders - a.seeders);
  return data;
};

export default {
  search,
};
