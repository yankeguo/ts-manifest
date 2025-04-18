package main

import (
	"bytes"
	"flag"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/yankeguo/rg"
)

var evalScript = template.Must(template.New("").Parse(`
{{range $i, $file := .Files}}
import * as file{{$i}} from {{$file}};
{{end}}
const items = [];
{{range $i, $file := .Files}}
if (file{{$i}}.default) {
  if (Array.isArray(file{{$i}}.default)) {
    items.push(...file{{$i}}.default);
  } else if (typeof file{{$i}}.default === "object") {
    items.push(file{{$i}}.default);
  }
}
{{end}}
await Deno.stdout.write(new TextEncoder().encode(JSON.stringify({ apiVersion: "v1", kind: "List", items }, null, 2)));
`))

func createEvalScript(files []string) string {
	for i, file := range files {
		files[i] = strconv.Quote("." + string(filepath.Separator) + file)
	}
	buf := &bytes.Buffer{}
	rg.Must0(evalScript.Execute(buf, map[string]any{
		"Files": files,
	}))
	return buf.String()
}

func main() {
	var err error
	defer func() {
		if err == nil {
			return
		}
		log.Printf("exited with error: %s", err.Error())
		os.Exit(1)
	}()
	defer rg.Guard(&err)

	log.SetOutput(os.Stderr)

	var (
		optCache bool
	)

	fset := flag.NewFlagSet("ts-manifest", flag.ExitOnError)
	fset.SetOutput(os.Stderr)
	fset.BoolVar(&optCache, "cache", false, "cache dependencies")
	fset.Parse(os.Args[1:])

	var files []string

	fs.WalkDir(os.DirFS("."), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == "." {
				return nil
			}
			if d.Name() == "node_modules" || strings.HasPrefix(d.Name(), ".") {
				return fs.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".ts") {
			return nil
		}
		files = append(files, path)
		return nil
	})

	if optCache {
		cmd := exec.Command("deno", append([]string{"cache"}, files...)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		rg.Must0(cmd.Run())
		return
	}

	cmd := exec.Command("deno", "run", "-A", "-")
	cmd.Stdin = bytes.NewReader([]byte(createEvalScript(files)))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	rg.Must0(cmd.Run())
}
