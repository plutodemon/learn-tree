// 类的使用

class Person {
    // 实例属性
    // private name: string;
    age: number;
    // 静态属性(类属性)
    static readonly type = 'Person';

    // 构造函数
    constructor(private name: string, age: number) {
        this.name = name;
        this.age = age;
    }

    say() {
        console.log(`name: ${this.name}, age: ${this.age}`);
    }

    set setName(name: string) {
        this.name = name;
    }

    get getName() {
        return this.name;
    }
}

const p = new Person('Tom', 188);
p.say();
p.setName = 'Jerry';
console.log(p.getName);

// 类的继承
class Student extends Person {
    // 子类的实例属性
    studentId: number;

    constructor(name: string, age: number, studentId: number) {
        // 调用父类的构造函数
        super(name, age);
        this.studentId = studentId;
    }

    // 方法的重写
    say() {
        console.log(`${this.getName} is studying...`);
    }
}

const s = new Student('Jerry', 20, 1001);
s.say();

// 抽象类 不能被实例化(不能new) 只能被继承
abstract class Animal {
    // 抽象方法 只能定义在抽象类中 不能有具体实现 必须在子类中实现
    abstract say(): void;
}

// 接口 用于定义对象的结构 限制类的结构
interface PersonInterface {
    name: string;
    age: number;

    say(): void;
}

interface PersonInterface {
    sex: string;
}

class MyPerson implements PersonInterface {
    name: string;
    age: number;

    say(): void {
        throw new Error("Method not implemented.");
    }

    sex: string;

    constructor(name: string, age: number, sex: string) {
        this.name = name;
        this.age = age;
        this.sex = sex;
    }
}

// 泛型
interface Inter {
    length: number;
}

function createArray<T extends Inter>(value: T): T[] {
    const arr = [];
    for (let i = 0; i < 10; i++) {
        arr[i] = value;
    }
    return arr;
}

createArray('a')
//createArray<number>(1)