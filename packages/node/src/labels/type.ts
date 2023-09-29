export class Label {
    private name: string;
    private value: string;
    constructor(name: string, value: string) {
        this.name = name;
        this.value = value;
    }

    public getName(): string {
        return this.name;
    }

    public getValue(): string {
        return this.value;
    }

    static Name(name: string) : (value: string) => Label {
        return (value: string) => new Label(name, value);
    }
}