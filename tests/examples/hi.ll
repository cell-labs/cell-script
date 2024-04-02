target triple = "riscv-unknown"

%string = type { i64, i8* }

@str.0 = constant [1 x i8] c"\00"

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

declare i8* @realloc(i8* %0, i64 %1)

declare i8* @memcpy(i8* %dest, i8* %src, i64 %n)

declare i8* @strcat(i8* %0, i8* %1)

declare i8* @strcpy(i8* %0, i8* %1)

declare i8* @strncpy(i8* %0, i8* %1, i64 %2)

declare i8* @strndup(i8* %0, i64 %1)

declare void @exit(i32 %0)

define i64 @string_len(%string %input) {
entry:
	%0 = extractvalue %string %input, 0
	ret i64 %0
}

define void @global-init-0() {
block-1:
	ret void
}

define i32 @main() {
block-2:
	call void @global-init-0()
	ret i32 0
}
