import type { PlopTypes } from "@turbo/gen";
import data from "./labels.json";

type Label = {
  name: string;
  value: string;
};
// Learn more about Turborepo Generators at https://turbo.build/repo/docs/core-concepts/monorepos/code-generation

function toSnakeUpperCase(label: string): Label {
  const newName = label
    .match(/[A-Z]{2,}(?=[A-Z][a-z]+[0-9]*|\b)|[A-Z]?[a-z]+[0-9]*|[A-Z]|[0-9]+/g)
    ?.map((x) => x.toUpperCase())
    .join("_");

  if (!newName) {
    throw new Error(`Unable to convert label ${label} to snake case`);
  }
  return {
    name: newName,
    value: label,
  };
}

// eslint-disable-next-line import/no-default-export -- Turbo generators require default export
export default function generator(plop: PlopTypes.NodePlopAPI): void {
  // A simple generator to add a new React component to the internal UI library
  const labels = { labels: data.labels.map(toSnakeUpperCase) };
  plop.setGenerator("labels", {
    description: "Generate common labels for go, typescript, and python",
    prompts: [],
    actions: [
      {
        type: "add",
        data: labels,
        path: "go/labels/common.go",
        templateFile: "templates/go.hbs",
        force: true,
      },
      {
        type: "add",
        data: labels,
        path: "node/src/labels/common.ts",
        templateFile: "templates/node.hbs",
        force: true,
      },
      {
        type: "add",
        data: labels,
        path: "python/labels/common.py",
        templateFile: "templates/python.hbs",
        force: true,
      },
    ],
  });
}
