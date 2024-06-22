# ANSI

This package is a collection of constants and helpers to build ANSI/VT-100 escape sequences for console manipulation. The meat is in [ansi.go](./ansi.go).

Note that the codes for `Reset` through to `BgWhite` are just the numbers, and not the full escape code. You need to wrap those codes in a call to SGR in order to build a working code. For example:

```go
fmt.Println(ansi.SGR(ansi.BgDarkRed, ansi.FgBlue), "Hello ", ansi.SGR(ansi.FsGreen), "World!", ansi.SGR(ansi.Reset))
```

On Windows, note that during init this package will make Windows API calls into GetConsoleMode and, if necessary, SetConsoleMode. This is to ensure the `ENABLE_VIRTUAL_TERMINAL_PROCESSING` flag (`0x04`) is set, which must be set for a Windows cmd prompt to actually process ANSI/VT-100 escape codes. The relevant code is in [init_windows.go](./init_windows.go).

## TODO

I needed to break this out from Frog so other projects could use it without adding unnecessary dependencies, but there's some usability stuff I may come back and do, like making the use of SGR codes simpler. Maybe move just those constants to an SGR package, so you'd do this: `ansi.SGR(sgr.BgDarkRed, sgr.FgBlue)`? And then I could add `ansi.FgBlue` as a constant that has the full SGR escape code built-in.
