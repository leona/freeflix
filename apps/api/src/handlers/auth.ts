import jellyfin from "../models/jellyfin.js";
import { sign } from "hono/jwt";

export const authCheck = async (ctx) => {
  return ctx.json({
    message: "Authenticated",
  });
};

export const authenticate = async (ctx) => {
  const data = await ctx.req.json();
  const { username, password } = data;

  try {
    await jellyfin.authenticate({ username, password });
    const token = await sign({ username }, process.env.JWT_SECRET);

    return ctx.json({
      jwt: token,
    });
  } catch (e) {
    console.log("Failed to login", e);
    ctx.status(401);

    return ctx.json({
      message: "Invalid Credentials",
    });
  }
};
