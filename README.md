# Invoice Tracker

This is a simple Web Invoice Tracker. In order to run this program, you have to install some things and knows how to use a CMD (Command Prompt) for Windows or Terminal for Linux.

## Introduction

You can track your receipt by simply inputting the details of each invoice. This helps you also visualize your expenses into organized collections of receipts. It also helps you keep a list of the ones who bought and where you bought it.

## Installation

First, you need to install go x.x.x on your machine and configured. You can download it on their page, which is [The Go Programming Language](https://golang.org/doc/install).

For **Windows**, you also have to install a [GCC Compiler](http://tdm-gcc.tdragon.net/download) and GIT. Also set correctly the Environment Variables.

For **Linux**:

```bash
$ nano .profile
```

And add this inside the .profile:

```bash
# set PATH so it includes user's private bin directories

PATH="$HOME/bin:$HOME/.local/bin:$PATH"
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Second, please do install [uAdmin](https://github.com/uadmin/uadmin), a Golang Web Framework. Please check their page in order to install it correctly.

## Run the program

**Windows**

For CMD, run it as Administrator.

```
C:\> cd Users\PC-Name\go\src\github.com\rmarasigan\receipt_tracking
C:\Users\PC-Name\go\src\github.com\rmarasigan\receipt_tracking> go build
C:\Users\PC-Name\go\src\github.com\rmarasigan\receipt_tracking> go run .\receipt_tracking
```

This will show you:
```bash
[   OK   ]   Initializing DB: [12/12]
[   OK   ]   Server Started: http://0.0.0.0:2563
         ___       __          _
  __  __/   | ____/ /___ ___  (_)___
 / / / / /| |/ __  / __ '__ \/ / __ \
/ /_/ / ___ / /_/ / / / / / / / / / /
\__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/
```

Open your browser and type: localhost:2563/receipt-tracker. It will show you the uAdmin Dashboard.

**Linux**

```bash
usename@pc-name:~$ cd go/src/github.com/rmarasigan/receipt_tracking
usename@pc-name:~/go/src/github.com/rmarasigan/receipt_tracking$ go build; ./receipt_tracking
```

Open your browser and type: 0.0.0.0:2563/receipt-tracker.
