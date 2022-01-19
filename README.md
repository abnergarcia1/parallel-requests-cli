# Parallel requests CLI

This tool help us to get the MD5 SHA from the response of a list of URLs
___
## Build
To build the app, we need to run the following command in the root of the project:
```
  make build
```
This will generate a binary called `parallel-requests-cli`
___
## Install
To build the go app, we need to run the following command in the root of the project:
```
  make install
```
This will build and install the binary called `parallel-requests-cli`
___
## Use

To use the app, you need to send as arguments the urls from where you want to get the SHA 
```
./parallel-requests-cli https://github.com/ http://www.google.com
```
and the response can be like:
```
parallel-requests-cli git:(main) ✗ ./parallel-requests-cli https://github.com/ http://www.google.com
http://www.google.com 30287771192cc5f242dfc36661e59fd1
https://github.com/ 7f6e213f7aab196fb01117cd9a6466a1
```
> **NOTE**: By default, the maximum parallel requests is 10, but you can change this parameter.
> Also, the urls must contain `http` or `https` in order to correctly make the request, otherwise, the tool will 
> use `http` protocol for the requests


### Parallel requests
Also we can set the maximum number of parallel requests using the `-parallel` parameter

example:
```
./parallel-requests-cli  -parallel 5 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
```

## Testing
In order to run the tests, just run:
```
make test
```