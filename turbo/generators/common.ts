export type Label = {
  name: string;
  type: string;
};

export type LabelWithVariable = Label & {
  variable: string;
};
type StringMapFunction = (string) => string;

export const idFunction = (str: string) => str;

const camelCaseRegex = /\b[a-zA-Z][a-zA-Z0-9]*\b/;

export function mapLabels(
  labels: Label[],
  varFn: StringMapFunction,
  typeFn: StringMapFunction
): { labels: LabelWithVariable[] } {
  return {
    labels: labels.map((label) => {
      return {
        variable: varFn(label.name),
        name: label.name,
        type: typeFn(label.type),
      };
    }),
  };
}

function splitCamelCase(name: string): string[] {
  return captializeFirstLetter(name).split(/(?=[A-Z])/);
}

export function getSnakeCaseIdentifierName(labelName: string): string {
    const words = splitCamelCase(labelName);
    return words.join("_").toUpperCase();
}

export function getCamelCaseIdentifierName(labelName: string): string {
  const words = splitCamelCase(labelName);

  if (words[words.length - 1] === "Id") {
    words[words.length - 1] = "ID";
  }
  return words.join("");
}

function captializeFirstLetter(str: string): string {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

export function validateLabels(allLabels: Record<string, Label[]>): void {
  Object.keys(allLabels).forEach((serivce) => {
    const labels = allLabels[serivce];
    const labelNames = labels.map((label) => label.name);
    const invalidLabels = labelNames.filter(
      (labelName) => !camelCaseRegex.test(labelName)
    );
    if (invalidLabels.length > 0) {
      throw new Error(`Labels:${invalidLabels} are invalid`);
    }
    const hasDuplicates = new Set(labelNames).size !== labelNames.length;
    if (hasDuplicates) {
      throw new Error("Labels contain duplicate names");
    }
  });
}

export function mapToNodeType(type: string): string {
  switch (type) {
    case "int":
      return "number";
    case "int64":
      return "number";
    case "float32":
      return "number";
    case "float64":
      return "number";
    case "bool":
      return "boolean";
    case "Time":
        return "Date";
    default:
      return type;
  }
}
