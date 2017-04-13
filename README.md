# gitlab-reporting
Reporting Tool for GitLab projects, written in Go

## Install go
If not already done, install go to get started

Follow the instructions at https://golang.org/doc/install to install and setup go

## Getting started

get the code by executing the following command at $GOPATH/src

```
go get github.com/adorsys/gitlab-reporting
```

or by cloning the repository into $GOPATH/src/github.com/adorsys

```
git clone https://github.com/adorsys/gitlab-reporting.git
```
if cloned execute following command from $GOPATH/src
```
go install github.com/adorsys/gitlab-reporting
```

## Usage

go to directory $GOPATH/src/github.com/adorsys/gitlab-reporting and execute

```
./../../../../bin/gitlab-reporting
```
The tool is now running at http://127.0.0.1:9090

To use the tool you´ll need:

- gitlab url where the project is
- gitlab private token
- project namespace/name

## Change address

To change the default ip adress execute tool with ip flag and the chosen adress
```
./../../../../bin/gitlab-reporting -ip {ip adress}
```

To change the default port execute tool with port flag and the chosen portnumber
```
./../../../../bin/gitlab-reporting -port {portnumber}
```
