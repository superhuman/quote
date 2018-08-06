A small utility to quote and unquote strings.

It assumes JSON escaping, but that's mostly compatible with what other programming
languages do (go, swift, javascript...).

### Installation

go install github.com/superhuman/quote

### Usage

```
echo -n 'h"i' | quote
# "h\"i"
```

```
echo -n '"h\"i"' | quote -d
# h"i
```

