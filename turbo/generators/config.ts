import type { PlopTypes } from "@turbo/gen";
import { type ActionType } from "node-plop";
import { labels } from "./labels.json";
import {
  Label,
  getCamelCaseIdentifierName,
  getSnakeCaseIdentifierName,
  idFunction,
  mapLabels,
  mapToNodeType,
  validateLabels,
} from "./common";

function generateGoActions(labels: Record<string, Label[]>): ActionType[] {
  return Object.keys(labels).map((serivce) => {
    return {
      type: "add",
      data: mapLabels(labels[serivce], getCamelCaseIdentifierName, idFunction),
      path: `go/labels/${serivce}.go`,
      templateFile: "templates/go.hbs",
      force: true,
    };
  });
}

export function generateNodeActions(
  labels: Record<string, Label[]>
): ActionType[] {
  return Object.keys(labels).map((serivce) => {
    return {
      type: "add",
      data: mapLabels(labels[serivce], getCamelCaseIdentifierName, mapToNodeType),
      path: `node/src/labels/${serivce}.ts`,
      templateFile: "templates/node.hbs",
      force: true,
    };
  });
}

function generatePythonActions(labels: Record<string, Label[]>): ActionType[] {
  return Object.keys(labels).map((serivce) => {
    return {
      type: "add",
      data: mapLabels(labels[serivce], getSnakeCaseIdentifierName, idFunction),
      path: `python/labels/${serivce}.py`,
      templateFile: "templates/python.hbs",
      force: true,
    };
  });
}

// Learn more about Turborepo Generators at https://turbo.build/repo/docs/core-concepts/monorepos/code-generation
export default function generator(plop: PlopTypes.NodePlopAPI): void {
  validateLabels(labels);
  const actions = [...generateGoActions(labels), ...generateNodeActions(labels), ...generatePythonActions(labels)];
  plop.setGenerator("labels", {
    description: "Generate labels",
    prompts: [],
    actions,
  });
}
