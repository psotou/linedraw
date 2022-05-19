# Usage

![Go](https://github.com/psotou/linedraw/actions/workflows/go.yml/badge.svg)

Import the package in order to use the line drawing of your choice:

```go
import (
    "github.com/psotou/linedraw"
)
```

And then use one of the four line drawing available by calling the type of line as follows:

```go
// we need to pass the length of the desired line to the factory function
line := NewLine(12)
line.Horizontal.Bold()
fmt.Println("hola")
line.Horizontal.Light()
```

This will output the following:

```shell
━━━━━━━━━━━━━━━━━━━━━━━━━
hola
─────────────────────────
```

The other options available are `Vertical`, `HorizontalDotted` and `VerticalDotted`, which will produce the following output:

```shell
┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃
hola
│││││││││││││││││││││││││
```

```shell
┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅
hola
┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄
```

```shell
┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇┇
hola
┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆┆
```
