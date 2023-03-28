// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package nethttp

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfHttpMethodInvocation struct {
	StartMonotimeNs uint64
	Regs            struct {
		R15     uint64
		R14     uint64
		R13     uint64
		R12     uint64
		Rbp     uint64
		Rbx     uint64
		R11     uint64
		R10     uint64
		R9      uint64
		R8      uint64
		Rax     uint64
		Rcx     uint64
		Rdx     uint64
		Rsi     uint64
		Rdi     uint64
		OrigRax uint64
		Rip     uint64
		Cs      uint64
		Eflags  uint64
		Rsp     uint64
		Ss      uint64
	}
}

type bpfHttpRequestTrace struct {
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Method          [6]uint8
	Path            [100]uint8
	Status          uint16
	RemoteAddr      [50]uint8
	Host            [256]uint8
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	UprobeServeHTTP       *ebpf.ProgramSpec `ebpf:"uprobe_ServeHTTP"`
	UprobeServeHttpReturn *ebpf.ProgramSpec `ebpf:"uprobe_ServeHttp_return"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	Events              *ebpf.MapSpec `ebpf:"events"`
	OngoingHttpRequests *ebpf.MapSpec `ebpf:"ongoing_http_requests"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	Events              *ebpf.Map `ebpf:"events"`
	OngoingHttpRequests *ebpf.Map `ebpf:"ongoing_http_requests"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.Events,
		m.OngoingHttpRequests,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	UprobeServeHTTP       *ebpf.Program `ebpf:"uprobe_ServeHTTP"`
	UprobeServeHttpReturn *ebpf.Program `ebpf:"uprobe_ServeHttp_return"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.UprobeServeHTTP,
		p.UprobeServeHttpReturn,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_x86.o
var _BpfBytes []byte
