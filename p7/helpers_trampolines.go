//go:build amd64 || arm64

package p7

import _ "unsafe"

import base "github.com/goccy/googlesqlwasm2go/base"

//go:linkname _xF32_eq github.com/goccy/googlesqlwasm2go/base.F32_eq
func _xF32_eq(x0 float32, x1 float32) int32
func f32_eq(x0 float32, x1 float32) int32 { return _xF32_eq(x0, x1) }

//go:linkname _xF64_abs github.com/goccy/googlesqlwasm2go/base.F64_abs
func _xF64_abs(x0 float64) float64
func f64_abs(x0 float64) float64 { return _xF64_abs(x0) }

//go:linkname _xF64_eq github.com/goccy/googlesqlwasm2go/base.F64_eq
func _xF64_eq(x0 float64, x1 float64) int32
func f64_eq(x0 float64, x1 float64) int32 { return _xF64_eq(x0, x1) }

//go:linkname _xF64_gt github.com/goccy/googlesqlwasm2go/base.F64_gt
func _xF64_gt(x0 float64, x1 float64) int32
func f64_gt(x0 float64, x1 float64) int32 { return _xF64_gt(x0, x1) }

//go:linkname _xF64_le github.com/goccy/googlesqlwasm2go/base.F64_le
func _xF64_le(x0 float64, x1 float64) int32
func f64_le(x0 float64, x1 float64) int32 { return _xF64_le(x0, x1) }

//go:linkname _xF64_lt github.com/goccy/googlesqlwasm2go/base.F64_lt
func _xF64_lt(x0 float64, x1 float64) int32
func f64_lt(x0 float64, x1 float64) int32 { return _xF64_lt(x0, x1) }

//go:linkname _xF64_ne github.com/goccy/googlesqlwasm2go/base.F64_ne
func _xF64_ne(x0 float64, x1 float64) int32
func f64_ne(x0 float64, x1 float64) int32 { return _xF64_ne(x0, x1) }

//go:linkname _xF64_promote_f32 github.com/goccy/googlesqlwasm2go/base.F64_promote_f32
func _xF64_promote_f32(x0 float32) float64
func f64_promote_f32(x0 float32) float64 { return _xF64_promote_f32(x0) }

//go:linkname _xI32_clz github.com/goccy/googlesqlwasm2go/base.I32_clz
func _xI32_clz(x0 int32) int32
func i32_clz(x0 int32) int32 { return _xI32_clz(x0) }

//go:linkname _xI32_div_s github.com/goccy/googlesqlwasm2go/base.I32_div_s
func _xI32_div_s(x0 int32, x1 int32) int32
func i32_div_s(x0 int32, x1 int32) int32 { return _xI32_div_s(x0, x1) }

//go:linkname _xI32_div_u_s github.com/goccy/googlesqlwasm2go/base.I32_div_u_s
func _xI32_div_u_s(x0 int32, x1 int32) int32
func i32_div_u_s(x0 int32, x1 int32) int32 { return _xI32_div_u_s(x0, x1) }

//go:linkname _xI32_eqz github.com/goccy/googlesqlwasm2go/base.I32_eqz
func _xI32_eqz(x0 int32) int32
func i32_eqz(x0 int32) int32 { return _xI32_eqz(x0) }

//go:linkname _xI32_extend8_s github.com/goccy/googlesqlwasm2go/base.I32_extend8_s
func _xI32_extend8_s(x0 int32) int32
func i32_extend8_s(x0 int32) int32 { return _xI32_extend8_s(x0) }

//go:linkname _xI32_rem_u_s github.com/goccy/googlesqlwasm2go/base.I32_rem_u_s
func _xI32_rem_u_s(x0 int32, x1 int32) int32
func i32_rem_u_s(x0 int32, x1 int32) int32 { return _xI32_rem_u_s(x0, x1) }

//go:linkname _xI32_wrap_i64 github.com/goccy/googlesqlwasm2go/base.I32_wrap_i64
func _xI32_wrap_i64(x0 int64) int32
func i32_wrap_i64(x0 int64) int32 { return _xI32_wrap_i64(x0) }

//go:linkname _xI64_clz github.com/goccy/googlesqlwasm2go/base.I64_clz
func _xI64_clz(x0 int64) int64
func i64_clz(x0 int64) int64 { return _xI64_clz(x0) }

//go:linkname _xI64_ctz github.com/goccy/googlesqlwasm2go/base.I64_ctz
func _xI64_ctz(x0 int64) int64
func i64_ctz(x0 int64) int64 { return _xI64_ctz(x0) }

//go:linkname _xI64_div_s github.com/goccy/googlesqlwasm2go/base.I64_div_s
func _xI64_div_s(x0 int64, x1 int64) int64
func i64_div_s(x0 int64, x1 int64) int64 { return _xI64_div_s(x0, x1) }

//go:linkname _xI64_div_u_s github.com/goccy/googlesqlwasm2go/base.I64_div_u_s
func _xI64_div_u_s(x0 int64, x1 int64) int64
func i64_div_u_s(x0 int64, x1 int64) int64 { return _xI64_div_u_s(x0, x1) }

//go:linkname _xI64_eqz github.com/goccy/googlesqlwasm2go/base.I64_eqz
func _xI64_eqz(x0 int64) int32
func i64_eqz(x0 int64) int32 { return _xI64_eqz(x0) }

//go:linkname _xI64_extend16_s github.com/goccy/googlesqlwasm2go/base.I64_extend16_s
func _xI64_extend16_s(x0 int64) int64
func i64_extend16_s(x0 int64) int64 { return _xI64_extend16_s(x0) }

//go:linkname _xI64_extend_i32_s github.com/goccy/googlesqlwasm2go/base.I64_extend_i32_s
func _xI64_extend_i32_s(x0 int32) int64
func i64_extend_i32_s(x0 int32) int64 { return _xI64_extend_i32_s(x0) }

//go:linkname _xI64_extend_i32_u github.com/goccy/googlesqlwasm2go/base.I64_extend_i32_u
func _xI64_extend_i32_u(x0 int32) int64
func i64_extend_i32_u(x0 int32) int64 { return _xI64_extend_i32_u(x0) }

//go:linkname _xI64_rem_s github.com/goccy/googlesqlwasm2go/base.I64_rem_s
func _xI64_rem_s(x0 int64, x1 int64) int64
func i64_rem_s(x0 int64, x1 int64) int64 { return _xI64_rem_s(x0, x1) }

//go:linkname _xMemoryCopy github.com/goccy/googlesqlwasm2go/base.MemoryCopy
func _xMemoryCopy(m *base.Module, dst int32, src int32, n int32)
func memoryCopy(m *base.Module, dst int32, src int32, n int32) { _xMemoryCopy(m, dst, src, n) }

//go:linkname _xMemoryFill github.com/goccy/googlesqlwasm2go/base.MemoryFill
func _xMemoryFill(m *base.Module, dst int32, val int32, n int32)
func memoryFill(m *base.Module, dst int32, val int32, n int32) { _xMemoryFill(m, dst, val, n) }

//go:linkname _xWasm_trap_div_zero github.com/goccy/googlesqlwasm2go/base.Wasm_trap_div_zero
func _xWasm_trap_div_zero()
func wasm_trap_div_zero() { _xWasm_trap_div_zero() }
