# go-tendo

[![Go Report Card](https://goreportcard.com/badge/github.com/andrewlader/go-tendo)](https://goreportcard.com/report/github.com/andrewlader/go-tendo)
[![Build Status](https://travis-ci.org/AndrewLader/go-tendo.svg?branch=master)](https://travis-ci.org/AndrewLader/go-tendo)
[![Coverage Status](https://coveralls.io/repos/github/AndrewLader/go-tendo/badge.svg)](https://coveralls.io/github/AndrewLader/go-tendo)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/AndrewLader/go-tendo/blob/master/LICENSE)

Go application used to walk a Go project and report on the number of structs, methods and functions in each package.

```
 ██████╗  ██████╗    ████████╗███████╗███╗   ██╗██████╗  ██████╗ 
██╔════╝ ██╔═══██╗   ╚══██╔══╝██╔════╝████╗  ██║██╔══██╗██╔═══██╗
██║  ███╗██║   ██║█████╗██║   █████╗  ██╔██╗ ██║██║  ██║██║   ██║
██║   ██║██║   ██║╚════╝██║   ██╔══╝  ██║╚██╗██║██║  ██║██║   ██║
╚██████╔╝╚██████╔╝      ██║   ███████╗██║ ╚████║██████╔╝╚██████╔╝
 ╚═════╝  ╚═════╝       ╚═╝   ╚══════╝╚═╝  ╚═══╝╚═════╝  ╚═════╝ 
```

### Usage

After building and installing `go-tendo` into a known path, navigate to the desired Go project and use the following command:

```
go-tendo [--log={logLevel}] {targetPath}
```

*example*

```
go-tendo --log=info ./
```

This command sets the logging level to output *info* related elements only, and inspect the source code at the target path of `./`

<img alt="Sample Output" src="https://github.com/AndrewLader/go-tendo/blob/master/images/go-tendo%20output.png" width="515px" />

**Log Levels**

* *LogAll* - All output is displayed
* *LogTrace* - Trace and above (e.g., trace, info, warnings and errors) output is displayed
* *LogInfo* - Info related output and above is displayed
* *LogWarnings* - Default output is displayed along with any warnings or errors
* *LogErrors* - Default output is displayed along with any errors

### License

[MIT](https://github.com/AndrewLader/go-tendo/blob/master/LICENSE)

