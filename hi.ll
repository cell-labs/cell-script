target triple = "riscv-unknown"

%string = type { i64, i8* }
%error = type { i1, %string }

@str.0 = constant [1 x i8] c"\00"
@str.1 = constant [4 x i8] c"%d\0A\00"
@str.2 = constant [4 x i8] c"%s\0A\00"
@str.3 = constant [2 x i8] c"1\00"

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %size)

declare i8* @realloc(i8* %ptr, i64 %size)

declare i8* @memcpy(i8* %dest, i8* %src, i64 %n)

declare i8* @memset(i8* %dest, i32 %c, i32 %n)

declare i8* @strcat(i8* %0, i8* %1)

declare i8* @strcpy(i8* %dst, i8* %src)

declare i8* @strncpy(i8* %dst, i8* %src, i64 %n)

declare i8* @strndup(i8* %str, i64 %n)

declare i64 @strcmp(i8* %lhs, i8* %rhs)

declare void @syscall_exit(i8 %exit_code)

define i64 @string_len(%string %input) {
entry:
	%0 = extractvalue %string %input, 0
	ret i64 %0
}

define void @global-init-0() {
block-1:
	ret void
}

define i64 @main() {
block-2:
	call void @global-init-0()
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	store i64 1, i64* %1
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	%3 = getelementptr [2 x i8], [2 x i8]* @str.3, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = load %string, %string* %0
	call void @debug_Printlnstring(%string %4)
	call void @debug_Printlnuint64(i64 1)
	call void @debug_Printlnuint8(i8 1)
	call void @debug_Printlnuint16(i16 1)
	call void @debug_Printlnuint32(i32 1)
	call void @debug_Printlnuint64(i64 1)
	call void @debug_Printlnuint128(i128 1)
	call void @debug_Printlnuint256(i256 1)
	ret i64 0
}

define %string @global_method_error_Errorpointererror(%error* %p-3) {
block-4:
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	store i64 0, i64* %1
	%3 = getelementptr [1 x i8], [1 x i8]* @str.0, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = getelementptr %error, %error* %p-3, i32 0, i32 1
	%5 = load %string, %string* %4
	ret %string %5
}

define %string @global_method_error_Errorpointererror_jump(i8* %unsafe-ptr) {
block-5:
	%0 = bitcast i8* %unsafe-ptr to %error*
	%1 = call %string @global_method_error_Errorpointererror(%error* %0)
	ret %string %1
}

define i1 @global_method_error_IsNonepointererror(%error* %p-6) {
block-7:
	%0 = alloca i1
	store i1 false, i1* %0
	%1 = getelementptr %error, %error* %p-6, i32 0, i32 0
	%2 = load i1, i1* %1
	ret i1 %2
}

define i1 @global_method_error_IsNonepointererror_jump(i8* %unsafe-ptr) {
block-8:
	%0 = bitcast i8* %unsafe-ptr to %error*
	%1 = call i1 @global_method_error_IsNonepointererror(%error* %0)
	ret i1 %1
}

define i1 @global_method_error_NotNonepointererror(%error* %p-9) {
block-10:
	%0 = alloca i1
	store i1 false, i1* %0
	%1 = getelementptr %error, %error* %p-9, i32 0, i32 0
	%2 = load i1, i1* %1
	%3 = xor i1 %2, true
	ret i1 %3
}

define i1 @global_method_error_NotNonepointererror_jump(i8* %unsafe-ptr) {
block-11:
	%0 = bitcast i8* %unsafe-ptr to %error*
	%1 = call i1 @global_method_error_NotNonepointererror(%error* %0)
	ret i1 %1
}

define %string @global_method_uint256_toStringuint256(i256 %p-12) {
block-13:
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	store i64 0, i64* %1
	%3 = getelementptr [1 x i8], [1 x i8]* @str.0, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = alloca { i32, i32, i32, i8* }
	%len-15 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 0
	%cap-16 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 1
	%offset-17 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 2
	%backing-18 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 3
	%5 = trunc i64 0 to i32
	store i32 %5, i32* %len-15
	store i32 2, i32* %cap-16
	store i32 0, i32* %offset-17
	%6 = mul i32 2, 1
	%7 = sext i32 %6 to i64
	%slicezero-19 = call i8* @malloc(i64 %7)
	%8 = bitcast i8* %slicezero-19 to i8*
	store i8* %8, i8** %backing-18
	%9 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 3
	%loadedbackingarrayptr-20 = load i8*, i8** %9
	%len-21 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 0
	%10 = trunc i64 0 to i32
	store i32 %10, i32* %len-21
	%cap-22 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 1
	store i32 2, i32* %cap-22
	%11 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4
	%digits-23 = alloca { i32, i32, i32, i8* }
	store { i32, i32, i32, i8* } %11, { i32, i32, i32, i8* }* %digits-23
	%i-24 = alloca i256
	store i256 %p-12, i256* %i-24
	br label %block-25-cond

block-25-cond:
	%12 = load i256, i256* %i-24
	%13 = icmp ugt i256 %12, 0
	br i1 %13, label %block-26-body, label %block-28-after

block-26-body:
	%14 = load i256, i256* %i-24
	%15 = urem i256 %14, 10
	%16 = alloca i8
	%17 = trunc i256 %15 to i8
	store i8 %17, i8* %16
	%18 = load i8, i8* %16
	%mod-29 = alloca i8
	store i8 %18, i8* %mod-29
	%sliceToAppendTo-33 = alloca { i32, i32, i32, i8* }
	%len-34 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 0
	%cap-35 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 1
	%19 = load i32, i32* %len-34
	%20 = load i32, i32* %cap-35
	%21 = icmp ult i32 %19, %20
	br i1 %21, label %block-32-append-existing-block, label %block-30-copy-slice

block-27-after-body:
	%22 = load i256, i256* %i-24
	%23 = udiv i256 %22, 10
	store i256 %23, i256* %i-24
	br label %block-25-cond

block-28-after:
	%24 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23
	%25 = extractvalue { i32, i32, i32, i8* } %24, 0
	%26 = alloca i64
	%27 = zext i32 %25 to i64
	store i64 %27, i64* %26
	%28 = load i64, i64* %26
	%length-51 = alloca i64
	store i64 %28, i64* %length-51
	%sliceToAppendTo-55 = alloca { i32, i32, i32, i8* }
	%len-56 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 0
	%cap-57 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 1
	%29 = load i32, i32* %len-56
	%30 = load i32, i32* %cap-57
	%31 = icmp ult i32 %29, %30
	br i1 %31, label %block-54-append-existing-block, label %block-52-copy-slice

block-30-copy-slice:
	%copy-to-new-slice-36 = alloca { i32, i32, i32, i8* }
	%len-37 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-36, i32 0, i32 0
	%cap-38 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-36, i32 0, i32 1
	%offset-39 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-36, i32 0, i32 2
	%backing-40 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-36, i32 0, i32 3
	%prev-len-41 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 0
	%prev-cap-42 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 1
	%32 = load i32, i32* %prev-len-41
	%33 = load i32, i32* %prev-cap-42
	store i32 %32, i32* %len-37
	store i32 0, i32* %offset-39
	%34 = mul i32 %33, 2
	%35 = zext i32 %34 to i64
	%36 = mul i64 %35, 20
	%slice-grow-43 = call i8* @malloc(i64 %36)
	store i32 %34, i32* %cap-38
	%37 = bitcast i8* %slice-grow-43 to i8*
	store i8* %37, i8** %backing-40
	%prev-backarr-44 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 3
	%38 = load i8*, i8** %prev-backarr-44
	%prev-backarray-casted-45 = bitcast i8* %38 to i8*
	%39 = alloca i32
	store i32 0, i32* %39
	%40 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-36
	store { i32, i32, i32, i8* } %40, { i32, i32, i32, i8* }* %sliceToAppendTo-33
	br label %block-46-copy-slice-bytes

block-31-add-to-slice:
	%sliceLenPtr-47 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %sliceToAppendTo-33, i32 0, i32 0
	%slicelen-48 = load i32, i32* %sliceLenPtr-47
	%backingarrayptr-49 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %sliceToAppendTo-33, i32 0, i32 3
	%41 = load i8*, i8** %backingarrayptr-49
	%store-ptr-50 = getelementptr i8, i8* %41, i32 %slicelen-48
	%42 = load i8, i8* %mod-29
	%43 = add i8 %42, 48
	store i8 %43, i8* %store-ptr-50
	%44 = add i32 %slicelen-48, 1
	store i32 %44, i32* %sliceLenPtr-47
	%45 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %sliceToAppendTo-33
	store { i32, i32, i32, i8* } %45, { i32, i32, i32, i8* }* %digits-23
	br label %block-27-after-body

block-32-append-existing-block:
	%46 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23
	store { i32, i32, i32, i8* } %46, { i32, i32, i32, i8* }* %sliceToAppendTo-33
	br label %block-31-add-to-slice

block-46-copy-slice-bytes:
	%47 = load i32, i32* %39
	%48 = getelementptr i8, i8* %38, i32 %47
	%49 = load i32, i32* %39
	%50 = getelementptr i8, i8* %37, i32 %49
	%51 = load i8, i8* %48
	store i8 %51, i8* %50
	%52 = load i32, i32* %39
	%53 = add i32 1, %52
	store i32 %53, i32* %39
	%54 = icmp ult i32 %53, %32
	br i1 %54, label %block-46-copy-slice-bytes, label %block-31-add-to-slice

block-52-copy-slice:
	%copy-to-new-slice-58 = alloca { i32, i32, i32, i8* }
	%len-59 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-58, i32 0, i32 0
	%cap-60 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-58, i32 0, i32 1
	%offset-61 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-58, i32 0, i32 2
	%backing-62 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-58, i32 0, i32 3
	%prev-len-63 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 0
	%prev-cap-64 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 1
	%55 = load i32, i32* %prev-len-63
	%56 = load i32, i32* %prev-cap-64
	store i32 %55, i32* %len-59
	store i32 0, i32* %offset-61
	%57 = mul i32 %56, 2
	%58 = zext i32 %57 to i64
	%59 = mul i64 %58, 20
	%slice-grow-65 = call i8* @malloc(i64 %59)
	store i32 %57, i32* %cap-60
	%60 = bitcast i8* %slice-grow-65 to i8*
	store i8* %60, i8** %backing-62
	%prev-backarr-66 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23, i32 0, i32 3
	%61 = load i8*, i8** %prev-backarr-66
	%prev-backarray-casted-67 = bitcast i8* %61 to i8*
	%62 = alloca i32
	store i32 0, i32* %62
	%63 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-58
	store { i32, i32, i32, i8* } %63, { i32, i32, i32, i8* }* %sliceToAppendTo-55
	br label %block-68-copy-slice-bytes

block-53-add-to-slice:
	%sliceLenPtr-69 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %sliceToAppendTo-55, i32 0, i32 0
	%slicelen-70 = load i32, i32* %sliceLenPtr-69
	%backingarrayptr-71 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %sliceToAppendTo-55, i32 0, i32 3
	%64 = load i8*, i8** %backingarrayptr-71
	%store-ptr-72 = getelementptr i8, i8* %64, i32 %slicelen-70
	%65 = alloca i8
	%66 = trunc i64 0 to i8
	store i8 %66, i8* %65
	%67 = load i8, i8* %65
	store i8 %67, i8* %store-ptr-72
	%68 = add i32 %slicelen-70, 1
	store i32 %68, i32* %sliceLenPtr-69
	%69 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %sliceToAppendTo-55
	store { i32, i32, i32, i8* } %69, { i32, i32, i32, i8* }* %digits-23
	%70 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23
	%71 = alloca { i32, i32, i32, i8* }
	%72 = extractvalue { i32, i32, i32, i8* } %70, 2
	%73 = zext i32 %72 to i64
	%74 = add i64 0, %73
	%75 = load i64, i64* %length-51
	%76 = sub i64 %75, %74
	%77 = trunc i64 %76 to i32
	%78 = trunc i64 %74 to i32
	%79 = extractvalue { i32, i32, i32, i8* } %70, 1
	%len-73 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %71, i32 0, i32 0
	store i32 %77, i32* %len-73
	%cap-74 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %71, i32 0, i32 1
	store i32 %79, i32* %cap-74
	%offset-75 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %71, i32 0, i32 2
	store i32 %78, i32* %offset-75
	%backing-76 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %71, i32 0, i32 3
	%80 = extractvalue { i32, i32, i32, i8* } %70, 3
	store i8* %80, i8** %backing-76
	%81 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %71
	%82 = extractvalue { i32, i32, i32, i8* } %81, 0
	%83 = extractvalue { i32, i32, i32, i8* } %81, 2
	%84 = extractvalue { i32, i32, i32, i8* } %81, 3
	%85 = getelementptr i8, i8* %84, i32 %83
	%86 = zext i32 %82 to i64
	%87 = alloca %string
	%len-77 = getelementptr %string, %string* %87, i32 0, i32 0
	store i64 %86, i64* %len-77
	%str-78 = getelementptr %string, %string* %87, i32 0, i32 1
	store i8* %85, i8** %str-78
	%88 = load %string, %string* %87
	%str-79 = alloca %string
	store %string %88, %string* %str-79
	%i-80 = alloca i64
	store i64 0, i64* %i-80
	br label %block-81-cond

block-54-append-existing-block:
	%89 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-23
	store { i32, i32, i32, i8* } %89, { i32, i32, i32, i8* }* %sliceToAppendTo-55
	br label %block-53-add-to-slice

block-68-copy-slice-bytes:
	%90 = load i32, i32* %62
	%91 = getelementptr i8, i8* %61, i32 %90
	%92 = load i32, i32* %62
	%93 = getelementptr i8, i8* %60, i32 %92
	%94 = load i8, i8* %91
	store i8 %94, i8* %93
	%95 = load i32, i32* %62
	%96 = add i32 1, %95
	store i32 %96, i32* %62
	%97 = icmp ult i32 %96, %55
	br i1 %97, label %block-68-copy-slice-bytes, label %block-53-add-to-slice

block-81-cond:
	%98 = load i64, i64* %length-51
	%99 = udiv i64 %98, 2
	%100 = load i64, i64* %length-51
	%101 = sub i64 %100, %99
	%102 = load i64, i64* %i-80
	%103 = icmp ult i64 %102, %101
	br i1 %103, label %block-82-body, label %block-84-after

block-82-body:
	%104 = load i64, i64* %i-80
	%105 = load %string, %string* %str-79
	%106 = extractvalue %string %105, 0
	%107 = extractvalue %string %105, 1
	%108 = getelementptr i8, i8* %107, i64 %104
	%109 = load i8, i8* %108
	%v-85 = alloca i8
	store i8 %109, i8* %v-85
	%110 = load i64, i64* %i-80
	%111 = load %string, %string* %str-79
	%112 = extractvalue %string %111, 0
	%113 = extractvalue %string %111, 1
	%114 = getelementptr i8, i8* %113, i64 %110
	%115 = load i64, i64* %length-51
	%116 = sub i64 %115, 1
	%117 = load i64, i64* %i-80
	%118 = sub i64 %116, %117
	%119 = load %string, %string* %str-79
	%120 = extractvalue %string %119, 0
	%121 = extractvalue %string %119, 1
	%122 = getelementptr i8, i8* %121, i64 %118
	%123 = load i8, i8* %122
	store i8 %123, i8* %114
	%124 = load i64, i64* %length-51
	%125 = sub i64 %124, 1
	%126 = load i64, i64* %i-80
	%127 = sub i64 %125, %126
	%128 = load %string, %string* %str-79
	%129 = extractvalue %string %128, 0
	%130 = extractvalue %string %128, 1
	%131 = getelementptr i8, i8* %130, i64 %127
	%132 = load i8, i8* %v-85
	store i8 %132, i8* %131
	br label %block-83-after-body

block-83-after-body:
	%133 = load i64, i64* %i-80
	%134 = add i64 %133, 1
	store i64 %134, i64* %i-80
	br label %block-81-cond

block-84-after:
	%135 = load %string, %string* %str-79
	ret %string %135
}

define %string @global_method_uint256_toStringuint256_jump(i8* %unsafe-ptr) {
block-14:
	%0 = bitcast i8* %unsafe-ptr to i256*
	%1 = load i256, i256* %0
	%2 = call %string @global_method_uint256_toStringuint256(i256 %1)
	ret %string %2
}

define %string @global_method_uint128_toStringuint128(i128 %p-86) {
block-87:
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	store i64 0, i64* %1
	%3 = getelementptr [1 x i8], [1 x i8]* @str.0, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = alloca { i32, i32, i32, i8* }
	%len-89 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 0
	%cap-90 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 1
	%offset-91 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 2
	%backing-92 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 3
	%5 = trunc i64 0 to i32
	store i32 %5, i32* %len-89
	store i32 2, i32* %cap-90
	store i32 0, i32* %offset-91
	%6 = mul i32 2, 1
	%7 = sext i32 %6 to i64
	%slicezero-93 = call i8* @malloc(i64 %7)
	%8 = bitcast i8* %slicezero-93 to i8*
	store i8* %8, i8** %backing-92
	%9 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 3
	%loadedbackingarrayptr-94 = load i8*, i8** %9
	%len-95 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 0
	%10 = trunc i64 0 to i32
	store i32 %10, i32* %len-95
	%cap-96 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4, i32 0, i32 1
	store i32 2, i32* %cap-96
	%11 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %4
	%digits-97 = alloca { i32, i32, i32, i8* }
	store { i32, i32, i32, i8* } %11, { i32, i32, i32, i8* }* %digits-97
	%i-98 = alloca i128
	store i128 %p-86, i128* %i-98
	br label %block-99-cond

block-99-cond:
	%12 = load i128, i128* %i-98
	%13 = icmp ugt i128 %12, 0
	br i1 %13, label %block-100-body, label %block-102-after

block-100-body:
	%14 = load i128, i128* %i-98
	%15 = urem i128 %14, 10
	%16 = alloca i8
	%17 = trunc i128 %15 to i8
	store i8 %17, i8* %16
	%18 = load i8, i8* %16
	%mod-103 = alloca i8
	store i8 %18, i8* %mod-103
	%sliceToAppendTo-107 = alloca { i32, i32, i32, i8* }
	%len-108 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-97, i32 0, i32 0
	%cap-109 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-97, i32 0, i32 1
	%19 = load i32, i32* %len-108
	%20 = load i32, i32* %cap-109
	%21 = icmp ult i32 %19, %20
	br i1 %21, label %block-106-append-existing-block, label %block-104-copy-slice

block-101-after-body:
	%22 = load i128, i128* %i-98
	%23 = udiv i128 %22, 10
	store i128 %23, i128* %i-98
	br label %block-99-cond

block-102-after:
	%24 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-97
	%25 = extractvalue { i32, i32, i32, i8* } %24, 0
	%26 = extractvalue { i32, i32, i32, i8* } %24, 2
	%27 = extractvalue { i32, i32, i32, i8* } %24, 3
	%28 = getelementptr i8, i8* %27, i32 %26
	%29 = zext i32 %25 to i64
	%30 = alloca %string
	%len-125 = getelementptr %string, %string* %30, i32 0, i32 0
	store i64 %29, i64* %len-125
	%str-126 = getelementptr %string, %string* %30, i32 0, i32 1
	store i8* %28, i8** %str-126
	%31 = load %string, %string* %30
	%str-127 = alloca %string
	store %string %31, %string* %str-127
	%32 = load %string, %string* %str-127
	%33 = call i64 @string_len(%string %32)
	%length-128 = alloca i64
	store i64 %33, i64* %length-128
	%i-129 = alloca i64
	store i64 0, i64* %i-129
	br label %block-130-cond

block-104-copy-slice:
	%copy-to-new-slice-110 = alloca { i32, i32, i32, i8* }
	%len-111 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-110, i32 0, i32 0
	%cap-112 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-110, i32 0, i32 1
	%offset-113 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-110, i32 0, i32 2
	%backing-114 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-110, i32 0, i32 3
	%prev-len-115 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-97, i32 0, i32 0
	%prev-cap-116 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-97, i32 0, i32 1
	%34 = load i32, i32* %prev-len-115
	%35 = load i32, i32* %prev-cap-116
	store i32 %34, i32* %len-111
	store i32 0, i32* %offset-113
	%36 = mul i32 %35, 2
	%37 = zext i32 %36 to i64
	%38 = mul i64 %37, 20
	%slice-grow-117 = call i8* @malloc(i64 %38)
	store i32 %36, i32* %cap-112
	%39 = bitcast i8* %slice-grow-117 to i8*
	store i8* %39, i8** %backing-114
	%prev-backarr-118 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-97, i32 0, i32 3
	%40 = load i8*, i8** %prev-backarr-118
	%prev-backarray-casted-119 = bitcast i8* %40 to i8*
	%41 = alloca i32
	store i32 0, i32* %41
	%42 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %copy-to-new-slice-110
	store { i32, i32, i32, i8* } %42, { i32, i32, i32, i8* }* %sliceToAppendTo-107
	br label %block-120-copy-slice-bytes

block-105-add-to-slice:
	%sliceLenPtr-121 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %sliceToAppendTo-107, i32 0, i32 0
	%slicelen-122 = load i32, i32* %sliceLenPtr-121
	%backingarrayptr-123 = getelementptr { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %sliceToAppendTo-107, i32 0, i32 3
	%43 = load i8*, i8** %backingarrayptr-123
	%store-ptr-124 = getelementptr i8, i8* %43, i32 %slicelen-122
	%44 = load i8, i8* %mod-103
	%45 = add i8 %44, 48
	store i8 %45, i8* %store-ptr-124
	%46 = add i32 %slicelen-122, 1
	store i32 %46, i32* %sliceLenPtr-121
	%47 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %sliceToAppendTo-107
	store { i32, i32, i32, i8* } %47, { i32, i32, i32, i8* }* %digits-97
	br label %block-101-after-body

block-106-append-existing-block:
	%48 = load { i32, i32, i32, i8* }, { i32, i32, i32, i8* }* %digits-97
	store { i32, i32, i32, i8* } %48, { i32, i32, i32, i8* }* %sliceToAppendTo-107
	br label %block-105-add-to-slice

block-120-copy-slice-bytes:
	%49 = load i32, i32* %41
	%50 = getelementptr i8, i8* %40, i32 %49
	%51 = load i32, i32* %41
	%52 = getelementptr i8, i8* %39, i32 %51
	%53 = load i8, i8* %50
	store i8 %53, i8* %52
	%54 = load i32, i32* %41
	%55 = add i32 1, %54
	store i32 %55, i32* %41
	%56 = icmp ult i32 %55, %34
	br i1 %56, label %block-120-copy-slice-bytes, label %block-105-add-to-slice

block-130-cond:
	%57 = load i64, i64* %length-128
	%58 = udiv i64 %57, 2
	%59 = load i64, i64* %length-128
	%60 = sub i64 %59, %58
	%61 = load i64, i64* %i-129
	%62 = icmp ult i64 %61, %60
	br i1 %62, label %block-131-body, label %block-133-after

block-131-body:
	%63 = load i64, i64* %i-129
	%64 = load %string, %string* %str-127
	%65 = extractvalue %string %64, 0
	%66 = extractvalue %string %64, 1
	%67 = getelementptr i8, i8* %66, i64 %63
	%68 = load i8, i8* %67
	%v-134 = alloca i8
	store i8 %68, i8* %v-134
	%69 = load i64, i64* %i-129
	%70 = load %string, %string* %str-127
	%71 = extractvalue %string %70, 0
	%72 = extractvalue %string %70, 1
	%73 = getelementptr i8, i8* %72, i64 %69
	%74 = load i64, i64* %length-128
	%75 = sub i64 %74, 1
	%76 = load i64, i64* %i-129
	%77 = sub i64 %75, %76
	%78 = load %string, %string* %str-127
	%79 = extractvalue %string %78, 0
	%80 = extractvalue %string %78, 1
	%81 = getelementptr i8, i8* %80, i64 %77
	%82 = load i8, i8* %81
	store i8 %82, i8* %73
	%83 = load i64, i64* %length-128
	%84 = sub i64 %83, 1
	%85 = load i64, i64* %i-129
	%86 = sub i64 %84, %85
	%87 = load %string, %string* %str-127
	%88 = extractvalue %string %87, 0
	%89 = extractvalue %string %87, 1
	%90 = getelementptr i8, i8* %89, i64 %86
	%91 = load i8, i8* %v-134
	store i8 %91, i8* %90
	br label %block-132-after-body

block-132-after-body:
	%92 = load i64, i64* %i-129
	%93 = add i64 %92, 1
	store i64 %93, i64* %i-129
	br label %block-130-cond

block-133-after:
	%94 = load %string, %string* %str-127
	ret %string %94
}

define %string @global_method_uint128_toStringuint128_jump(i8* %unsafe-ptr) {
block-88:
	%0 = bitcast i8* %unsafe-ptr to i128*
	%1 = load i128, i128* %0
	%2 = call %string @global_method_uint128_toStringuint128(i128 %1)
	ret %string %2
}

define void @debug_Printlnuint8(i8 %p-135) {
block-136:
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	store i64 3, i64* %1
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	%3 = getelementptr [4 x i8], [4 x i8]* @str.1, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = load %string, %string* %0
	%5 = extractvalue %string %4, 1
	%6 = call i32 (i8*, ...) @printf(i8* %5, i8 %p-135)
	ret void
}

define void @debug_Printlnuint16(i16 %p-137) {
block-138:
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	store i64 3, i64* %1
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	%3 = getelementptr [4 x i8], [4 x i8]* @str.1, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = load %string, %string* %0
	%5 = extractvalue %string %4, 1
	%6 = call i32 (i8*, ...) @printf(i8* %5, i16 %p-137)
	ret void
}

define void @debug_Printlnuint32(i32 %p-139) {
block-140:
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	store i64 3, i64* %1
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	%3 = getelementptr [4 x i8], [4 x i8]* @str.1, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = load %string, %string* %0
	%5 = extractvalue %string %4, 1
	%6 = call i32 (i8*, ...) @printf(i8* %5, i32 %p-139)
	ret void
}

define void @debug_Printlnuint64(i64 %p-141) {
block-142:
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	store i64 3, i64* %1
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	%3 = getelementptr [4 x i8], [4 x i8]* @str.1, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = load %string, %string* %0
	%5 = extractvalue %string %4, 1
	%6 = call i32 (i8*, ...) @printf(i8* %5, i64 %p-141)
	ret void
}

define void @debug_Printlnuint128(i128 %p-143) {
block-144:
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	store i64 3, i64* %1
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	%3 = getelementptr [4 x i8], [4 x i8]* @str.2, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = load %string, %string* %0
	%5 = alloca i128
	store i128 %p-143, i128* %5
	%6 = load i128, i128* %5
	%7 = call %string @global_method_uint128_toStringuint128(i128 %6)
	%8 = extractvalue %string %4, 1
	%9 = extractvalue %string %7, 1
	%10 = call i32 (i8*, ...) @printf(i8* %8, i8* %9)
	ret void
}

define void @debug_Printlnuint256(i256 %p-145) {
block-146:
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	store i64 3, i64* %1
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	%3 = getelementptr [4 x i8], [4 x i8]* @str.2, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = load %string, %string* %0
	%5 = alloca i256
	store i256 %p-145, i256* %5
	%6 = load i256, i256* %5
	%7 = call %string @global_method_uint256_toStringuint256(i256 %6)
	%8 = extractvalue %string %4, 1
	%9 = extractvalue %string %7, 1
	%10 = call i32 (i8*, ...) @printf(i8* %8, i8* %9)
	ret void
}

define void @debug_Printlnstring(%string %p-147) {
block-148:
	%paramPtr-149 = alloca %string
	store %string %p-147, %string* %paramPtr-149
	%0 = alloca %string
	%1 = getelementptr %string, %string* %0, i32 0, i32 0
	store i64 3, i64* %1
	%2 = getelementptr %string, %string* %0, i32 0, i32 1
	%3 = getelementptr [4 x i8], [4 x i8]* @str.2, i32 0, i32 0
	store i8* %3, i8** %2
	%4 = load %string, %string* %0
	%5 = extractvalue %string %4, 1
	%6 = load %string, %string* %paramPtr-149
	%7 = extractvalue %string %6, 1
	%8 = call i32 (i8*, ...) @printf(i8* %5, i8* %7)
	ret void
}

