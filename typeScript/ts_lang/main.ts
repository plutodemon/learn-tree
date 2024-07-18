function greeter(person: string): string {
    if (person === "Jane User") {
        return "Hello, Jane User";
    }
    return "Hello, " + person;
}

let user = "Jane";

console.log(greeter(user));