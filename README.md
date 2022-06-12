# golib

Generic functions that can be used as a library

### contexts
used for copying contexts in case something needs to run asynchronously and not get canceled if parent ctx is canceled.

### dates
used for date related functions. 

### handler
holds only one function for now. GetRequestParam is used to get a request parameter value from map[string]string | map[string][]string | url.Values 

### hid
creates a unique human readable ID (not safe for scaling)

### logger
a wrap of logrus that I prefer

### pointers
is used to be able to write one-liners

### queue
a wrapper for rabbitmq. TODO: convert it to interface or plugin to support different queues

### routerwrapper
provides better way to create APIs with optional query parameters

### slices
holds common functions for slices

### stringtools
holds common functions for strings

----

<p style="text-align: center">
<a style="text-decoration: none" href="go.mod">
    <img src="https://img.shields.io/github/go-mod/go-version/mikarios/golib?style=plastic" alt="Go version">
</a>


<a href="https://codecov.io/gh/mikarios/golib" style="text-decoration: none">
    <img src="https://img.shields.io/codecov/c/github/mikarios/golib?label=codecov&style=plastic" alt="code coverage"/>
</a>

<a style="text-decoration: none" href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=plastic" alt="License: MIT">
</a>

<br />

<a style="text-decoration: none" href="https://github.com/mikarios/golib/stargazers">
    <img src="https://img.shields.io/github/stars/mikarios/golib.svg?style=plastic" alt="Stars">
</a>

<a style="text-decoration: none" href="https://github.com/mikarios/golib/fork">
    <img src="https://img.shields.io/github/forks/mikarios/golib.svg?style=plastic" alt="Forks">
</a>

<a style="text-decoration: none" href="https://github.com/mikarios/golib/issues">
    <img src="https://img.shields.io/github/issues/mikarios/golib.svg?style=plastic" alt="Issues">
</a>
</p>
