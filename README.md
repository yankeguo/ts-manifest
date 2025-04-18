# ts-manifest

> [!WARNING]
>
> Due to the limitation of `deno`'s dynamic import, this project uses child process to evaluate `.ts` files.

## Usage

- `ts-manifest -cache` to cache all `deno` dependencies

- `ts-manifest` to evaluate every `.ts` files in the current directory with `deno`, and all `default` exports will be combined into a Kubernetes `List` json.

## Credits

GUO YANKE, MIT License
