import { walk } from "jsr:@std/fs@1.0.16/walk";
import { join } from "jsr:@std/path@1.0.8";
import { stringify } from "jsr:@std/yaml@1.0.5";

const items = [];

for await (
  const entry of walk(".", {
    exts: [".ts"],
    skip: [/^node_modules$/, /^\..+/],
  })
) {
  if (!entry.isFile) continue;

  const module = await import(join(Deno.cwd(), entry.path));

  if (module && module.default) {
    if (Array.isArray(module.default)) {
      items.push(...module.default);
    } else {
      items.push(module.default);
    }
  }
}

await Deno.stdout.write(
  new TextEncoder().encode(
    stringify({
      apiVersion: "v1",
      kind: "List",
      items: items,
    }),
  ),
);
