# ts-manifest

> [!WARNING]
>
> Since the limitation of `deno`'s dynamic import, this project uses child
> process to evaluate `.ts` files.

## Usage

Everything `.ts` file in the current directory will be evaluated with `deno`,
and the `default` export will be combined into a Kubernetes yaml manifest.

## Credits

GUO YANKE, MIT License
