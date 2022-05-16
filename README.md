# Usage

Import the package in order to use the line drawing of your choice:

```go
import (
    "github.com/psotou/linedraw"
)
```

And then use one of the four line drawing available by calling the type of line as follows:

```go
// we need to pass the length of the desired line to the factory function
hline := linedraw.NewHLine(12)
hline.Bold()
fmt.Println("hola")
hline.Light()
```

This will output the following:

```shell
━━━━━━━━━━━━━━━━━━━━━━━━━
hola
─────────────────────────
```

The other options available are vertical lines, dotted horizontal lines and dotted vertical lines, which will produce the following output:

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
