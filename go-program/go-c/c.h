
// 由Go语言实现的接口 export printStr1
extern void printStr1(char* str);

// 申明 C 接口 无需转换参数: string <=> _GoString_
extern void printStr2(_GoString_ s); // go ^1.10
