#!/usr/bin/env bash
go-bindata -pkg=template  -ignore=.go -nomemcopy  tmpl/...