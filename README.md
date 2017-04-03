# Description

`jstpl` is an utility to quickly format a json input to the user desired format using golang `text/template` package.

# Usage

Using a file as input and output

**input.json**
```json
{
    "name": "toto"
}
```

**output.tpl**
```
Hello {{.name}}!
```

**bash**
```bash
$ jstpl -in input.json -f output.tpl
Hello toto!
```

Using standard input and command line format

**bash**
```bash
$ echo '{"name":"toto"}' | jstpl 'Hello {{.name}}!"
Hello toto!
```

# Plugins

`jstpl` makes use of plugins introduced in go 1.8. They must be located on the same directory as `jstpl`, and be prefixed by `jstpl-plugin-` to be loaded.

See https://github.com/Inozuma/jstpl-plugin-strings for an example.
