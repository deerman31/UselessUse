// dev.ts
import dev from "$fresh/dev.ts";

// const port = parseInt(Deno.env.get("PORT") || "8000");
const port = parseInt(Deno.env.get("FRONTEND_PORT") || "8000");

await dev(import.meta.url, "./main.ts", { port });
