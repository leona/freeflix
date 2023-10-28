import { createClient } from "@apps/api/src/client";

const copyClient = () => {
  return createClient({
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  });
};

export let client = copyClient();

const authenticate = async ({ username, password }) => {
  const response = await client.auth.$post({
    json: {
      username,
      password,
    },
  });

  const data = await response.json();

  if (!data.jwt) {
    throw new Error("Login failed");
  }

  localStorage.setItem("token", data.jwt);
  client = copyClient();
};

const queue = async (url) => {
  console.log("queueing:", url);

  const response = await client.queue.$post({
    json: {
      url,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to add torrent");
  }
};

const search = async (query) => {
  console.log("searching for:", query);

  const response = await client.search.$get({
    query: {
      query,
    },
  });

  return await response.json();
};

const suggest = async (query) => {
  const response = await client.suggest.$get({
    query: {
      query,
    },
  });

  return await response.json();
};

const downloads = async () => {
  const response = await client.downloads.$get();
  return await response.json();
};

const remove = async (hash) => {
  const response = await client.remove.$delete({
    json: {
      hash,
    },
  });

  return await response.json();
};

const removeByTitle = async (title) => {
  const response = await client["remove-title"].$delete({
    json: {
      title,
    },
  });

  return await response.json();
};

const watch = async (name) => {
  const response = await client.watch.$get({
    query: {
      name,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to get watch URL");
  }
  return await response.json();
};

const checkAuth = async () => {
  const response = await client.auth.$get();
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
  suggest,
};
