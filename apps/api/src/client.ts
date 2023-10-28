import { hc } from "hono/client";
import app from ".";

const host = "/api";

export const createClient = ({ headers }) => {
  return hc<typeof app>(host, {
    headers,
  });
};
