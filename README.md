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
