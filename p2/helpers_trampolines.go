//go:build amd64 || arm64

package p2

import _ "unsafe"

import base "github.com/goccy/googlesqlwasm2go/base"

//go:linkname _xI32_div_s github.com/goccy/googlesqlwasm2go/base.I32_div_s
func _xI32_div_s(x0 int32, x1 int32) int32
func i32_div_s(x0 int32, x1 int32) int32 { return _xI32_div_s(x0, x1) }

//go:linkname _xI32_eqz github.com/goccy/googlesqlwasm2go/base.I32_eqz
func _xI32_eqz(x0 int32) int32
func i32_eqz(x0 int32) int32 { return _xI32_eqz(x0) }

//go:linkname _xI32_extend8_s github.com/goccy/googlesqlwasm2go/base.I32_extend8_s
func _xI32_extend8_s(x0 int32) int32
func i32_extend8_s(x0 int32) int32 { return _xI32_extend8_s(x0) }

//go:linkname _xI32_wrap_i64 github.com/goccy/googlesqlwasm2go/base.I32_wrap_i64
func _xI32_wrap_i64(x0 int64) int32
func i32_wrap_i64(x0 int64) int32 { return _xI32_wrap_i64(x0) }

//go:linkname _xI64_ctz github.com/goccy/googlesqlwasm2go/base.I64_ctz
func _xI64_ctz(x0 int64) int64
func i64_ctz(x0 int64) int64 { return _xI64_ctz(x0) }

//go:linkname _xI64_eqz github.com/goccy/googlesqlwasm2go/base.I64_eqz
func _xI64_eqz(x0 int64) int32
func i64_eqz(x0 int64) int32 { return _xI64_eqz(x0) }

//go:linkname _xI64_extend_i32_u github.com/goccy/googlesqlwasm2go/base.I64_extend_i32_u
func _xI64_extend_i32_u(x0 int32) int64
func i64_extend_i32_u(x0 int32) int64 { return _xI64_extend_i32_u(x0) }

//go:linkname _xMemoryFill github.com/goccy/googlesqlwasm2go/base.MemoryFill
func _xMemoryFill(m *base.Module, dst int32, val int32, n int32)
func memoryFill(m *base.Module, dst int32, val int32, n int32) { _xMemoryFill(m, dst, val, n) }

//go:linkname _xWasm_trap_div_zero github.com/goccy/googlesqlwasm2go/base.Wasm_trap_div_zero
func _xWasm_trap_div_zero()
func wasm_trap_div_zero() { _xWasm_trap_div_zero() }
