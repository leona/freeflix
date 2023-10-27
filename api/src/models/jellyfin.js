const baseUrl = "http://jellyfin:8096";

const search = async ({ query }) => {
  console.log("searching for:", query);
  const response = await fetch(`${baseUrl}/Search/Hints?searchTerm=${query}`, {
    headers: {
      Authorization: `MediaBrowser Token="${process.env.JELLYFIN_API_KEY}"`,
    },
  });

  console.log("search response code:", response.statusText);
  const data = await response.json();
  console.log("got data:", data);
  return data;
};

const authenticate = async ({ username, password }) => {
  console.log("logging in with", username);

  const response = await fetch(`${baseUrl}/Users/authenticatebyname`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-Emby-Authorization": `MediaBrowser Client="Jellyfin Web", Device="Chrome", DeviceId="123", Version="10.8.11"`,
    },
    body: JSON.stringify({
      Username: username,
      Pw: password,
    }),
  });

  console.log("auth response code:", response.statusText);
  const data = await response.json();

  if (!data?.User?.Name) {
    throw new Error("Login failed");
  }

  return data;
};

export default {
  search,
  baseUrl,
  authenticate,
};
