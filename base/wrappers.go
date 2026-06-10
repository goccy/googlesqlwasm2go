//go:build amd64 || arm64

package base

func callImport_0(m *Module, l0 int32) int32 {
	return m.Env.X_ZNK9googlesql12ASTIntoAlias13GetAsIdStringEv(m, l0)
}

func callImport_1(m *Module, l0 int32) int32 {
	return m.Env.X__cxa_allocate_exception(m, l0)
}

func callImport_2(m *Module, l0 int32, l1 int32, l2 int32) {
	m.Env.X__cxa_throw(m, l0, l1, l2)
}

func callImport_3(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	m.Env.X_ZN4absl12lts_2024072216raw_log_internal6RawLogENS0_11LogSeverityEPKciS4_z(m, l0, l1, l2, l3, l4)
}

func callImport_4(m *Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	return m.Wasmify.Callback_invoke(m, l0, l1, l2, l3)
}

func callImport_5(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Env.X__cxa_thread_atexit(m, l0, l1, l2)
}

func callImport_6(m *Module, l0 int32) int32 {
	return m.Env.U_isUWhiteSpace_76(m, l0)
}

func callImport_7(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Env.X_ZN4absl12lts_2024072213GetStackTraceEPPvii(m, l0, l1, l2)
}

func callImport_8(m *Module) int32 {
	return m.Env.X_ZN4absl12lts_2024072212log_internal8TimeZoneEv(m)
}

func callImport_9(m *Module) int32 {
	return m.Env.X_ZN4absl12lts_2024072212log_internal13IsInitializedEv(m)
}

func callImport_10(m *Module) int32 {
	return m.Env.X_ZN4absl12lts_2024072212log_internal24MaxFramesInLogStackTraceEv(m)
}

func callImport_11(m *Module) int32 {
	return m.Env.X_ZN4absl12lts_2024072212log_internal28ShouldSymbolizeLogStackTraceEv(m)
}

func callImport_12(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Env.X_ZN4absl12lts_202407229SymbolizeEPKvPci(m, l0, l1, l2)
}

func callImport_13(m *Module) int32 {
	return m.Env.X_ZN4absl12lts_2024072212log_internal12ExitOnDFatalEv(m)
}

func callImport_14(m *Module, l0 int32) int32 {
	return m.Env.X_ZN4absl12lts_2024072212log_internal24SetSuppressSigabortTraceEb(m, l0)
}

func callImport_15(m *Module, l0 int32, l1 int32) {
	m.Env.X_ZN4absl12lts_2024072212log_internal13WriteToStderrENSt3__217basic_string_viewIcNS2_11char_traitsIcEEEENS0_11LogSeverityE(m, l0, l1)
}

func callImport_16(m *Module, l0 int32) int32 {
	return m.Env.X_ZN6icu_769ErrorCodeD1Ev(m, l0)
}

func callImport_17(m *Module, l0 int32) int32 {
	return m.Env.X_ZNK6icu_769ErrorCode9errorNameEv(m, l0)
}

func callImport_18(m *Module, l0 int32) int32 {
	return m.Env.X_ZN6icu_768ByteSinkD2Ev(m, l0)
}

func callImport_19(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	return m.Env.Utf8_prevCharSafeBody_76(m, l0, l1, l2, l3, l4)
}

func callImport_20(m *Module, l0 int32) {
	m.Env.X_ZN6icu_767UMemorydlEPv(m, l0)
}

func callImport_21(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return m.Env.X_ZN6icu_768ByteSink15GetAppendBufferEiiPciPi(m, l0, l1, l2, l3, l4, l5)
}

func callImport_22(m *Module, l0 int32) {
	m.Env.X_ZN6icu_768ByteSink5FlushEv(m, l0)
}

func callImport_23(m *Module, l0 int32) int32 {
	return m.Env.X_ZN4absl12lts_2024072213time_internal4cctz12TimeZoneLibC4MakeERKNSt3__212basic_stringIcNS4_11char_traitsIcEENS4_9allocatorIcEEEE(m, l0)
}

func callImport_24(m *Module, l0 int32) int32 {
	return m.Env.X_ZN6icu_7611Normalizer223getNFKCCasefoldInstanceER10UErrorCode(m, l0)
}

func callImport_25(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Env.Utf8_back1SafeBody_76(m, l0, l1, l2)
}

func callImport_26(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Environ_get(m, l0, l1)
}

func callImport_27(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Environ_sizes_get(m, l0, l1)
}

func callImport_28(m *Module, l0 int32, l1 int64, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Clock_time_get(m, l0, l1, l2)
}

func callImport_29(m *Module, l0 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_close(m, l0)
}

func callImport_30(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_fdstat_get(m, l0, l1)
}

func callImport_31(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_fdstat_set_flags(m, l0, l1)
}

func callImport_32(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_prestat_get(m, l0, l1)
}

func callImport_33(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_prestat_dir_name(m, l0, l1, l2)
}

func callImport_34(m *Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_read(m, l0, l1, l2, l3)
}

func callImport_35(m *Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_seek(m, l0, l1, l2, l3)
}

func callImport_36(m *Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_write(m, l0, l1, l2, l3)
}

func callImport_37(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int64, l7 int32, l8 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_open(m, l0, l1, l2, l3, l4, l5, l6, l7, l8)
}

func callImport_38(m *Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Poll_oneoff(m, l0, l1, l2, l3)
}

func callImport_39(m *Module, l0 int32) {
	m.Wasi_snapshot_preview1.Proc_exit(m, l0)
}

func callImport_40(m *Module) int32 {
	return m.Wasi_snapshot_preview1.Sched_yield(m)
}

func loadGlobal_0(m *Module) int32     { return m.G0 }
func storeGlobal_0(m *Module, v int32) { m.G0 = v }

func callIndirect_type0(m *Module, idx int32, p0 int32, p1 int32) int64 {
	return m.T0[idx].(func(*Module, int32, int32) int64)(m, p0, p1)
}

func callIndirect_type1(m *Module, idx int32, p0 int32, p1 int32, p2 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32))(m, p0, p1, p2)
}

func callIndirect_type2(m *Module, idx int32, p0 int32, p1 int32) {
	m.T0[idx].(func(*Module, int32, int32))(m, p0, p1)
}

func callIndirect_type3(m *Module, idx int32, p0 int32) int32 {
	return m.T0[idx].(func(*Module, int32) int32)(m, p0)
}

func callIndirect_type4(m *Module, idx int32, p0 int32) {
	m.T0[idx].(func(*Module, int32))(m, p0)
}

func callIndirect_type5(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32))(m, p0, p1, p2, p3)
}

func callIndirect_type6(m *Module, idx int32, p0 int32, p1 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32) int32)(m, p0, p1)
}

func callIndirect_type7(m *Module, idx int32, p0 int32, p1 int32, p2 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32) int32)(m, p0, p1, p2)
}

func callIndirect_type8(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32, int32))(m, p0, p1, p2, p3, p4)
}

func callIndirect_type10(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3)
}

func callIndirect_type11(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32))(m, p0, p1, p2, p3, p4, p5)
}

func callIndirect_type12(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32, p6 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32, int32))(m, p0, p1, p2, p3, p4, p5, p6)
}

func callIndirect_type13(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3, p4)
}

func callIndirect_type14(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int64, p4 int32, p5 int64) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int64, int32, int64) int32)(m, p0, p1, p2, p3, p4, p5)
}

func callIndirect_type15(m *Module, idx int32) int32 {
	return m.T0[idx].(func(*Module) int32)(m)
}

func callIndirect_type16(m *Module, idx int32, p0 int32) int64 {
	return m.T0[idx].(func(*Module, int32) int64)(m, p0)
}

func callIndirect_type17(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3, p4, p5)
}

func callIndirect_type19(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32, p6 int32, p7 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, p0, p1, p2, p3, p4, p5, p6, p7)
}

func callIndirect_type20(m *Module, idx int32, p0 int32, p1 int64) {
	m.T0[idx].(func(*Module, int32, int64))(m, p0, p1)
}

func callIndirect_type21(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32, p6 int32, p7 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3, p4, p5, p6, p7)
}

func callIndirect_type22(m *Module, idx int32) {
	m.T0[idx].(func(*Module))(m)
}

func callIndirect_type23(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32, p6 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3, p4, p5, p6)
}

func callIndirect_type24(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32, p6 int32, p7 int32, p8 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32, int32, int32, int32))(m, p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

func callIndirect_type27(m *Module, idx int32, p0 int32, p1 int32, p2 int64) {
	m.T0[idx].(func(*Module, int32, int32, int64))(m, p0, p1, p2)
}

func callIndirect_type30(m *Module, idx int32, p0 int32, p1 int64, p2 int32) {
	m.T0[idx].(func(*Module, int32, int64, int32))(m, p0, p1, p2)
}

func callIndirect_type35(m *Module, idx int32, p0 int32, p1 int64, p2 int32) int64 {
	return m.T0[idx].(func(*Module, int32, int64, int32) int64)(m, p0, p1, p2)
}

func callIndirect_type48(m *Module, idx int32, p0 int32, p1 int32, p2 int64, p3 int32, p4 int32) {
	m.T0[idx].(func(*Module, int32, int32, int64, int32, int32))(m, p0, p1, p2, p3, p4)
}

func callIndirect_type50(m *Module, idx int32, p0 int32, p1 int32) float64 {
	return m.T0[idx].(func(*Module, int32, int32) float64)(m, p0, p1)
}

func callIndirect_type53(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int64) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int64) int32)(m, p0, p1, p2, p3, p4)
}

func callIndirect_type58(m *Module, idx int32, p0 int32, p1 int32) float32 {
	return m.T0[idx].(func(*Module, int32, int32) float32)(m, p0, p1)
}

func callIndirect_type68(m *Module, idx int32, p0 int32, p1 float64, p2 int32) {
	m.T0[idx].(func(*Module, int32, float64, int32))(m, p0, p1, p2)
}

func callIndirect_type75(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 float64) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, float64) int32)(m, p0, p1, p2, p3, p4)
}

func callIndirect_type91(m *Module, idx int32, p0 int32, p1 float32, p2 int32) {
	m.T0[idx].(func(*Module, int32, float32, int32))(m, p0, p1, p2)
}

func callIndirect_type96(m *Module, idx int32, p0 int64) {
	m.T0[idx].(func(*Module, int64))(m, p0)
}
