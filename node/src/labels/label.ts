export class Label {
    private name: string;
    private value: any;
    constructor(name: string, value: any) {
        this.name = name;
        this.value = value;
    }

    public getName(): string {
        return this.name;
    }

    public getValue(): any {
        return this.value;
    }

    static Name<T>(name: string) : (value: T) => Label {
        return (value: T) => new Label(name, value);
    }
}