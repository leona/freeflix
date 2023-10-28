import { Hono } from "hono";
import * as handlers from "./handlers";
import { cors } from "hono/cors";
import { jwt } from "hono/jwt";

const app = new Hono();

app.use(
  "/*",
  cors({
    origin: ["*"],
  })
);

if (!process.env.JWT_SECRET?.length) {
  throw new Error("JWT_SECRET is not set");
}

app.post("/auth", handlers.authenticate);

app.use(
  "/*",
  jwt({
    secret: process.env.JWT_SECRET,
  })
);

app.get("/search", handlers.search);
app.get("/downloads", handlers.downloads);
app.post("/queue", handlers.queue);
app.delete("/remove", handlers.remove);
app.delete("/remove-title", handlers.removeTitle);
app.get("/watch", handlers.watch);
app.get("/auth", handlers.authCheck);
app.get("/suggest", handlers.suggest);

console.log("Listening on port 3000");
export default app;
