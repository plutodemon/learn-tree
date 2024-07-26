// ts语法
let str: string = 'hello world';
let num: number = 1;

function rand(min: number, max: number): number {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

let flag: boolean | number;
// unknown 为类型安全的any
let un: unknown;
// 断言
str = un as string;
str = <string>un;

// object
let userObj: { name: string, age?: number };
userObj = {name: 'a'};

// 任意属性
let userObj2: { name: string, [prop: string]: any };
userObj2 = {name: 'a', sex: 'man'}

let userObj3: { name: string } & { age: number };
userObj3 = {name: 'a', age: 1};

// 函数类型声明
let fn: (a: number, b: number) => number;
fn = (a, b) => a + b;

// 数组 元组 枚举
let arr: number[] = [1, 2, 3];
let arr2: Array<number> = [1, 2, 3];
let arr3: [string, number] = ['a', 1];

enum Color {Red, Green, Blue}

// 类型别名
type t1 = 1 | 2 | 3;
let t: t1;
t = 1;