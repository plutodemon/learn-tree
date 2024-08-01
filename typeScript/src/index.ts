// 类的使用

class Person {
    // 实例属性
    name: string;
    age: number;
    // 静态属性(类属性)
    static readonly type = 'Person';

    // 构造函数
    constructor(name: string, age: number) {
        this.name = name;
        this.age = age;
    }

    say() {
        console.log(`name: ${this.name}, age: ${this.age}`);
    }
}

const p = new Person('Tom', 188);
p.say();

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
        console.log(`${this.name} is studying...`);
    }
}

const s = new Student('Jerry', 20, 1001);
s.say();