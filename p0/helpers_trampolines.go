//go:build amd64 || arm64

package p0

import _ "unsafe"

//go:linkname _xI32_div_s github.com/goccy/googlesqlwasm2go/base.I32_div_s
func _xI32_div_s(x0 int32, x1 int32) int32
func i32_div_s(x0 int32, x1 int32) int32 { return _xI32_div_s(x0, x1) }

//go:linkname _xI32_eqz github.com/goccy/googlesqlwasm2go/base.I32_eqz
func _xI32_eqz(x0 int32) int32
func i32_eqz(x0 int32) int32 { return _xI32_eqz(x0) }

//go:linkname _xWasm_trap_div_zero github.com/goccy/googlesqlwasm2go/base.Wasm_trap_div_zero
func _xWasm_trap_div_zero()
func wasm_trap_div_zero() { _xWasm_trap_div_zero() }
