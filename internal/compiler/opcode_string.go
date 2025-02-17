// Code generated by "stringer -type=Opcode,AugOp,BuiltinOp"; DO NOT EDIT.

package compiler

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Nop-0]
	_ = x[Num-1]
	_ = x[Str-2]
	_ = x[Dupe-3]
	_ = x[Drop-4]
	_ = x[Swap-5]
	_ = x[Field-6]
	_ = x[FieldInt-7]
	_ = x[FieldByName-8]
	_ = x[FieldByNameStr-9]
	_ = x[Global-10]
	_ = x[Local-11]
	_ = x[Special-12]
	_ = x[ArrayGlobal-13]
	_ = x[ArrayLocal-14]
	_ = x[InGlobal-15]
	_ = x[InLocal-16]
	_ = x[AssignField-17]
	_ = x[AssignGlobal-18]
	_ = x[AssignLocal-19]
	_ = x[AssignSpecial-20]
	_ = x[AssignArrayGlobal-21]
	_ = x[AssignArrayLocal-22]
	_ = x[Delete-23]
	_ = x[DeleteAll-24]
	_ = x[IncrField-25]
	_ = x[IncrGlobal-26]
	_ = x[IncrLocal-27]
	_ = x[IncrSpecial-28]
	_ = x[IncrArrayGlobal-29]
	_ = x[IncrArrayLocal-30]
	_ = x[AugAssignField-31]
	_ = x[AugAssignGlobal-32]
	_ = x[AugAssignLocal-33]
	_ = x[AugAssignSpecial-34]
	_ = x[AugAssignArrayGlobal-35]
	_ = x[AugAssignArrayLocal-36]
	_ = x[Regex-37]
	_ = x[IndexMulti-38]
	_ = x[ConcatMulti-39]
	_ = x[Add-40]
	_ = x[Subtract-41]
	_ = x[Multiply-42]
	_ = x[Divide-43]
	_ = x[Power-44]
	_ = x[Modulo-45]
	_ = x[Equals-46]
	_ = x[NotEquals-47]
	_ = x[Less-48]
	_ = x[Greater-49]
	_ = x[LessOrEqual-50]
	_ = x[GreaterOrEqual-51]
	_ = x[Concat-52]
	_ = x[Match-53]
	_ = x[NotMatch-54]
	_ = x[Not-55]
	_ = x[UnaryMinus-56]
	_ = x[UnaryPlus-57]
	_ = x[Boolean-58]
	_ = x[Jump-59]
	_ = x[JumpFalse-60]
	_ = x[JumpTrue-61]
	_ = x[JumpEquals-62]
	_ = x[JumpNotEquals-63]
	_ = x[JumpLess-64]
	_ = x[JumpGreater-65]
	_ = x[JumpLessOrEqual-66]
	_ = x[JumpGreaterOrEqual-67]
	_ = x[Next-68]
	_ = x[Nextfile-69]
	_ = x[Exit-70]
	_ = x[ForIn-71]
	_ = x[BreakForIn-72]
	_ = x[CallBuiltin-73]
	_ = x[CallLengthArray-74]
	_ = x[CallSplit-75]
	_ = x[CallSplitSep-76]
	_ = x[CallSprintf-77]
	_ = x[CallUser-78]
	_ = x[CallNative-79]
	_ = x[Return-80]
	_ = x[ReturnNull-81]
	_ = x[Nulls-82]
	_ = x[Print-83]
	_ = x[Printf-84]
	_ = x[Getline-85]
	_ = x[GetlineField-86]
	_ = x[GetlineGlobal-87]
	_ = x[GetlineLocal-88]
	_ = x[GetlineSpecial-89]
	_ = x[GetlineArray-90]
	_ = x[EndOpcode-91]
}

const _Opcode_name = "NopNumStrDupeDropSwapFieldFieldIntFieldByNameFieldByNameStrGlobalLocalSpecialArrayGlobalArrayLocalInGlobalInLocalAssignFieldAssignGlobalAssignLocalAssignSpecialAssignArrayGlobalAssignArrayLocalDeleteDeleteAllIncrFieldIncrGlobalIncrLocalIncrSpecialIncrArrayGlobalIncrArrayLocalAugAssignFieldAugAssignGlobalAugAssignLocalAugAssignSpecialAugAssignArrayGlobalAugAssignArrayLocalRegexIndexMultiConcatMultiAddSubtractMultiplyDividePowerModuloEqualsNotEqualsLessGreaterLessOrEqualGreaterOrEqualConcatMatchNotMatchNotUnaryMinusUnaryPlusBooleanJumpJumpFalseJumpTrueJumpEqualsJumpNotEqualsJumpLessJumpGreaterJumpLessOrEqualJumpGreaterOrEqualNextNextfileExitForInBreakForInCallBuiltinCallLengthArrayCallSplitCallSplitSepCallSprintfCallUserCallNativeReturnReturnNullNullsPrintPrintfGetlineGetlineFieldGetlineGlobalGetlineLocalGetlineSpecialGetlineArrayEndOpcode"

var _Opcode_index = [...]uint16{0, 3, 6, 9, 13, 17, 21, 26, 34, 45, 59, 65, 70, 77, 88, 98, 106, 113, 124, 136, 147, 160, 177, 193, 199, 208, 217, 227, 236, 247, 262, 276, 290, 305, 319, 335, 355, 374, 379, 389, 400, 403, 411, 419, 425, 430, 436, 442, 451, 455, 462, 473, 487, 493, 498, 506, 509, 519, 528, 535, 539, 548, 556, 566, 579, 587, 598, 613, 631, 635, 643, 647, 652, 662, 673, 688, 697, 709, 720, 728, 738, 744, 754, 759, 764, 770, 777, 789, 802, 814, 828, 840, 849}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AugOpAdd-0]
	_ = x[AugOpSub-1]
	_ = x[AugOpMul-2]
	_ = x[AugOpDiv-3]
	_ = x[AugOpPow-4]
	_ = x[AugOpMod-5]
}

const _AugOp_name = "AugOpAddAugOpSubAugOpMulAugOpDivAugOpPowAugOpMod"

var _AugOp_index = [...]uint8{0, 8, 16, 24, 32, 40, 48}

func (i AugOp) String() string {
	if i < 0 || i >= AugOp(len(_AugOp_index)-1) {
		return "AugOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AugOp_name[_AugOp_index[i]:_AugOp_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BuiltinAtan2-0]
	_ = x[BuiltinClose-1]
	_ = x[BuiltinCos-2]
	_ = x[BuiltinExp-3]
	_ = x[BuiltinFflush-4]
	_ = x[BuiltinFflushAll-5]
	_ = x[BuiltinGsub-6]
	_ = x[BuiltinIndex-7]
	_ = x[BuiltinInt-8]
	_ = x[BuiltinLength-9]
	_ = x[BuiltinLengthArg-10]
	_ = x[BuiltinLog-11]
	_ = x[BuiltinMatch-12]
	_ = x[BuiltinRand-13]
	_ = x[BuiltinSin-14]
	_ = x[BuiltinSqrt-15]
	_ = x[BuiltinSrand-16]
	_ = x[BuiltinSrandSeed-17]
	_ = x[BuiltinSub-18]
	_ = x[BuiltinSubstr-19]
	_ = x[BuiltinSubstrLength-20]
	_ = x[BuiltinSystem-21]
	_ = x[BuiltinTolower-22]
	_ = x[BuiltinToupper-23]
}

const _BuiltinOp_name = "BuiltinAtan2BuiltinCloseBuiltinCosBuiltinExpBuiltinFflushBuiltinFflushAllBuiltinGsubBuiltinIndexBuiltinIntBuiltinLengthBuiltinLengthArgBuiltinLogBuiltinMatchBuiltinRandBuiltinSinBuiltinSqrtBuiltinSrandBuiltinSrandSeedBuiltinSubBuiltinSubstrBuiltinSubstrLengthBuiltinSystemBuiltinTolowerBuiltinToupper"

var _BuiltinOp_index = [...]uint16{0, 12, 24, 34, 44, 57, 73, 84, 96, 106, 119, 135, 145, 157, 168, 178, 189, 201, 217, 227, 240, 259, 272, 286, 300}

func (i BuiltinOp) String() string {
	if i < 0 || i >= BuiltinOp(len(_BuiltinOp_index)-1) {
		return "BuiltinOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BuiltinOp_name[_BuiltinOp_index[i]:_BuiltinOp_index[i+1]]
}
