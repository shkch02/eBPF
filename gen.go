package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cflags "-I./include" counter src/counter.c
