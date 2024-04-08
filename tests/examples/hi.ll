target triple = "riscv-unknown"

%string = type { i64, i8* }

@str.0 = constant [1 x i8] c"\00"
@str.1 = constant [3 x i8] c"hi\00"

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

declare i8* @realloc(i8* %0, i64 %1)

declare i8* @memcpy(i8* %dest, i8* %src, i64 %n)

declare i8* @strcat(i8* %0, i8* %1)

declare i8* @strcpy(i8* %0, i8* %1)

declare i8* @strncpy(i8* %0, i8* %1, i64 %2)

declare i8* @strndup(i8* %0, i64 %1)

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
        store i64 2, i64* %1
        %2 = getelementptr %string, %string* %0, i32 0, i32 1
        %3 = getelementptr [3 x i8], [3 x i8]* @str.1, i32 0, i32 0
        store i8* %3, i8** %2
        %4 = load %string, %string* %0
        %5 = extractvalue %string %4, 1
        %6 = call i32 (i8*, ...) @printf(i8* %5)
        ret i64 0
}