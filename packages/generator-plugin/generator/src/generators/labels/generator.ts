import { formatFiles, generateFiles, Tree } from '@nx/devkit';
import {join} from 'path';
import { LabelsGeneratorSchema } from './schema';
import { Label, labels } from './labels';

type LabelMetadata = {
  lang: string;
  root: string;
  path: string;
};

export const LabelFiles: LabelMetadata[] = [
  {
    lang: 'go',
    root: 'go',
    path: 'labels',
  },
  {
    lang: 'node',
    root: 'node',
    path: 'src/labels',
  },
  {
    lang: 'python',
    root: 'python',
    path: 'labels',
  },
];

export async function labelsGenerator(
  tree: Tree,
  _options: LabelsGeneratorSchema
) {
  LabelFiles.forEach((file) => {
    const projectRoot = join('packages', file.root, file.path);
    generateFiles(tree, join(__dirname, 'files', file.lang), projectRoot, {
      labels: labels.map(toSnakeUpperCase),
    });
  });
  await formatFiles(tree);
}

function toSnakeUpperCase(label: Label): Label {
  const newName = label.name
    .match(/[A-Z]{2,}(?=[A-Z][a-z]+[0-9]*|\b)|[A-Z]?[a-z]+[0-9]*|[A-Z]|[0-9]+/g)
    ?.map((x) => x.toUpperCase())
    .join('_');

  return {
    name: newName || label.name,
    value: label.value,
  };
}

export default labelsGenerator;
