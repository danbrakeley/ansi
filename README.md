# ANSI

This package is a collection of constants and helpers to build ANSI/VT-100 escape sequences for console manipulation. The meat is in [ansi.go](./ansi.go).

Note that the SGR codes (for display attributes like color) can be combined using the `SGR()` function, which takes the sgr values from the `sgr` package. If you just want to use a single code, there are helpers in the main package for that, e.g. `ansi.SGR(sgr.Reset)` is the same as `ansi.Reset`. For example:

```go
fmt.Println(ansi.SGR(sgr.BgDarkRed, sgr.FgBlue), "Hello ", ansi.FsGreen, "World!", ansi.Reset)
```

On Windows, note that during init this package will make Windows API calls into GetConsoleMode and, if necessary, SetConsoleMode. This is to ensure the `ENABLE_VIRTUAL_TERMINAL_PROCESSING` flag (`0x04`) is set, which must be set for a Windows cmd prompt to actually process ANSI/VT-100 escape codes. The relevant code is in [init_windows.go](./init_windows.go).
