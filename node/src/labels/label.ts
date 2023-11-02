export class Label<T> {
    private name: string;
    private value: T;
    constructor(name: string, value: T) {
        this.name = name;
        this.value = value;
    }

    public getName(): string {
        return this.name;
    }

    public getValue(): T {
        return this.value;
    }

    static Name<T>(name: string) : (value: T) => Label<T> {
        return (value: T) => new Label<T>(name, value);
    }
}