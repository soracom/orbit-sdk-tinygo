# orbit-sdk-tinygo [![PkgGoDev](https://pkg.go.dev/badge/github.com/soracom/orbit-sdk-tinygo)](https://pkg.go.dev/github.com/soracom/orbit-sdk-tinygo) [![Go Report Card](https://goreportcard.com/badge/github.com/soracom/orbit-sdk-tinygo)](https://goreportcard.com/report/github.com/soracom/orbit-sdk-tinygo)

A SORACOM Orbit SDK is for [TinyGo](https://tinygo.org/) to generate wasm modules.

## Notice

Please use the latest **tagged** release that supports the features currently available on the Soracom platform. The `main` branch may contain unreleased features and could break your code.

## Synopsis

### Hello world

```go
import sdk github.com/soracom/orbit-sdk-tinygo

// For processing uplink (UE to SORACOM)
//export uplink
func uplink() sdk.ErrorCode {
	sdk.Log("hello world, uplink")
	return 0
}

// For processing downlink (SORACOM to UE)
//export downlink
func downlink() sdk.ErrorCode {
	sdk.Log("hello world, downlink")
	return 0
}
```

And more pragmatic examples are in [examples/](./examples/).

## Description

Please see [godoc](https://pkg.go.dev/github.com/soracom/orbit-sdk-tinygo) for detailed instructions on how to use this SDK.

## Known issues

### Running wasm on the SORACOM Orbit "sandbox" runtime

There are some heavily restricted rules to execute your wasm module on the SORACOM Orbit runtime. For example:

- It doesn't allow to execute a wasm module that depends on `syscall/js.*`
  - Basically, wasm that is generated by TinyGo depends on `js` runtime (see also: https://tinygo.org/webassembly/webassembly/).
  - However, the Orbit runtime doesn't support the `js` runtime. So __it is necessary to avoid such dependencies carefully__.
    - For example, `fmt` package is not allowed because the package requires `syscall/js.*` modules.
- `json` package is not supported by TinyGo.
  - see also: https://tinygo.org/lang-support/stdlib/
  - This SDK uses the following libraries for workaround.
    - [moznion/go-json-ice](https://github.com/moznion/go-json-ice)
    - [moznion/jsonparser#without_js_runtime_tinygo](https://github.com/moznion/jsonparser/tree/without_js_runtime_tinygo)

## License

```
Copyright 2020 SORACOM, Inc. https://soracom.io/

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

```

