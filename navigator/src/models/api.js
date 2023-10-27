const baseUrl = "/api";
let loginToken = localStorage.getItem("token");

const queue = async (url) => {
  console.log("queueing:", url);

  const response = await fetch(`${baseUrl}/queue`, {
    method: "POST",
    body: JSON.stringify({
      url,
    }),
    headers: {
      Authorization: `Bearer ${loginToken}`,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to add torrent");
  }
};

const search = async (query) => {
  console.log("searching for:", query);
  const response = await fetch(`${baseUrl}/search?query=${query}`, {
    headers: {
      Authorization: `Bearer ${loginToken}`,
    },
  });
  return await response.json();
};

const downloads = async () => {
  const response = await fetch(`${baseUrl}/downloads`, {
    headers: {
      Authorization: `Bearer ${loginToken}`,
    },
  });

  return await response.json();
};

const remove = async (hash) => {
  const response = await fetch(`${baseUrl}/remove`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${loginToken}`,
    },
    body: JSON.stringify({
      hash,
    }),
  });
  return await response.json();
};

const removeByTitle = async (title) => {
  const response = await fetch(`${baseUrl}/remove-title`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${loginToken}`,
    },
    body: JSON.stringify({
      title,
    }),
  });
  return await response.json();
};

const watch = async (name) => {
  const response = await fetch(`${baseUrl}/watch?query=${name}`, {
    headers: {
      Authorization: `Bearer ${loginToken}`,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to get watch URL");
  }
  return await response.json();
};

const authenticate = async ({ username, password }) => {
  const response = await fetch(`${baseUrl}/auth`, {
    method: "POST",
    body: JSON.stringify({
      username,
      password,
    }),
  });

  const data = await response.json();

  if (!data.jwt) {
    throw new Error("Login failed");
  }

  loginToken = data.jwt;
  localStorage.setItem("token", loginToken);
};

const checkAuth = async () => {
  const response = await fetch(`${baseUrl}/auth`, {
    headers: {
      Authorization: `Bearer ${loginToken}`,
    },
  });

  const data = await response.json();

  if (!data.message || !response.ok) {
    throw new Error("Login failed");
  }
};

export default {
  queue,
  search,
  downloads,
  remove,
  removeByTitle,
  watch,
  authenticate,
  checkAuth,
};
