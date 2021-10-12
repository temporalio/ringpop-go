ringpop-go [![Build Status](https://travis-ci.org/uber/ringpop-go.svg?branch=master)](https://travis-ci.org/uber/ringpop-go) [![Coverage Status](https://coveralls.io/repos/uber/ringpop-go/badge.svg?branch=master&service=github)](https://coveralls.io/github/uber/ringpop-go?branch=master)
==========

**(This project is no longer under active development. Temporal will eventually deprecate usage of Ringpop.)**

Ringpop is a library that brings cooperation and coordination to distributed
applications ([see Uber announcement blogpost](https://eng.uber.com/ringpop-open-source-nodejs-library/)). It maintains a consistent hash ring on top of a membership
protocol and provides request forwarding as a routing convenience. It can be
used to shard your application in a way that's scalable and fault tolerant.

Getting started
---------------

To install ringpop-go:

```
go get github.com/temporalio/ringpop-go
```

Developing
----------

First make certain that `thrift` v0.9.3
(OSX: `brew install https://gist.githubusercontent.com/chrislusf/8b4e7c19551ba220232f037b43c0eaf3/raw/01465b867b8ef9af7c7c3fa830c83666c825122d/thrift.rb`) and `glide` are
in your path (above). Then,

```
make setup
```

to install remaining golang dependencies and install the pre-commit hook.

Finally, run the tests by doing (note ensure you have enough file descriptors using `ulimit -n` - atleast 8192 reccomended.):

```
make test
```

Documentation
--------------

Interested in where to go from here? Read the docs at
[ringpop.readthedocs.org](https://ringpop.readthedocs.org)
