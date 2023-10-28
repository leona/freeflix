const baseUrl = `http://client`;

const downloads = async () => {
  const response = await fetch(`${baseUrl}/downloads`, {
    method: "GET",
  });

  if (!response.ok) {
    console.log(
      "Failed to get downloads:",
      response.statusText,
      await response.text()
    );
    throw new Error("Failed to get downloads");
  }

  return await response.json();
};

const add = async (magnet) => {
  console.log("adding torrent:", magnet);

  const response = await fetch(`${baseUrl}/add`, {
    method: "POST",
    body: JSON.stringify({
      magnet,
    }),
  });

  const text = await response.text();

  console.log(
    "added new torrent response:",
    text,
    "status:",
    response.statusText
  );

  if (!response.ok) {
    throw new Error("Failed to add torrent");
  }
};

const remove = async (hash) => {
  console.log("deleting torrent:", hash);

  const response = await fetch(`${baseUrl}/remove/${hash}`, {
    method: "DELETE",
  });

  const text = await response.text();

  console.log(
    "deleted torrent response:",
    text,
    "status:",
    response.statusText
  );
};

const removeByTitle = async (title) => {
  console.log("deleting torrent:", title);

  const response = await fetch(`${baseUrl}/remove`, {
    method: "DELETE",
    body: JSON.stringify({
      title,
    }),
  });

  const text = await response.text();

  console.log(
    "deleted torrent response:",
    text,
    "status:",
    response.statusText
  );
};

export default {
  downloads,
  add,
  remove,
  removeByTitle,
};
