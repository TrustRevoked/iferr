# iferr

Given a Go file and a byte position, it will generate an appropriate `if err != nil { ... }` statement

## Installation

```
go get github.com/TrustRevoked/iferr
```

## Usage

```
% iferr /path/to/file.go 123
if err != nil {
return "", nil, err
}
```

For functions which don't return an error type, it will use `log.Fatal(err)` instead.

## vim integration

```
function GoIfErr()
  let path = '/tmp/goiferr.txt'
  let position = line2byte(line('.'))+col('.')
  execute 'write!' path
  let output = system('iferr '.shellescape(path).' '.shellescape(position))
  let chomped = substitute(output, '\n\+$', '', '')
  return chomped
endfunction

iabbrev <expr> iferr GoIfErr()
```

Then you can type "iferr" anywhere and it'll be expanded.

## Limitations

Currently it only looks for top-level declarations. So it won't produce the correct output for inline functions.

## License

MIT
