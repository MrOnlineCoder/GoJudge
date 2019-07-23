# GoJudge

GoJudge is a small programming contest management system written in Go.

## Note
Please note, this project is mostly educational. If you really want to use it in production environment (i.e. running local programming contests) - use it at your own risk!

### Goal of this project

Currently, there is no certain goal for it. I wanted to get experience using Go and at the same time was doing some contests on [Codeforces](https://codeforces.com) (at which I am no very good by the way). So, I decided to make a Go web application for judging.

### Getting Started

One of the things that I don't like in other contest systems is their installation. For example, [ejudge](https://ejudge.ru/) requires Linux and compilation from scratch. It is also quite complex to configure. Another one, [DOMJudge](https://www.domjudge.org), which is used at ACM ICPC, needs some dependencies, like MySQL server and also Linux OS.

GoJudge tries to minimize effort needed to install and setup a working instance of GoJudge.

**Pre-built binary**

You can download GoJudge prebuilt binary in GitHub Releases page for your system. It includes GoJudge executable, Sandbox runner and starting directory structure. 

Starting GoJudge is quite straightforward - just run `gojudge` main executable
```
cd gojudge
gojudge
```

**Compiling from source**

To compile and run from source, you will need Git and working Go installation.

```
git clone https://github.com/MrOnlineCoder/GoJudge.git
cd GoJudge
go get ./...
go build
gojudge
```

**Usage**

If GoJudge server has successfully started, open your favourite browser and browse to 

```
http://localhost:1337
```

GoJudge client interface must open just for you.

GoJudge will create default administrator user on first run. Use it to log in and setup your contest.

Default administrator username: **admin**

Default administrator password: **admin**


**Do not forget to change admin's password after successful login!**

### Features

Please note, not all features are well-tested and ready to use. 

* No dependencies or third-party software needed for running
* Uses SQLite for storing all data
* Easy to use client interface built in Vue and Bootstrap
* Supports different methods for checking participants' solutions
* Runs submissions in a safe cross-platform sandbox environment
* Cross-platform: can run both on Windows and Linux!

### TODO

* Saving contest to a file (or DB)
* Scoreboard
* Handle contest end and results
* Jury user interface
* Realtime submission status update via Socket.io
* Clarifications
* Token based checking
* Checking using external checker program
* Configuration files

### Architecture

* **Web server**: Golang [net/http](https://golang.org/pkg/net/http/) server with [gorilla/mux](https://github.com/gorilla/mux) router
* **Web client**: [Vue.js](https://vuejs.org/) SPA application
* **Web UI**: [Bootstrap](https://getbootstrap.com) via [Bootstrap-Vue](https://bootstrap-vue.js.org)
* **Judging system**: serveral worker [goroutines](https://tour.golang.org/concurrency/1), which receive incoming submission via channels

More on judging: GoJudge starts several worker goroutines (usually the number of workers equal the number of CPU cores on the machine). Each worker awaits for incoming submission and when one is avaliable, starts judging it. All commands are executed with [os.Exec](https://golang.org/pkg/os/exec/). Firstly, submission source code is written to a file in sandbox folder. Then, compiler is called to compile the submission. After that, compiled submission is ran through a set of tests, running via `sandbox_runner` each test. Final verdict is written to the database and source code and executable files for submission get deleted from sandbox folder.

### License
MIT (see LICENSE file)