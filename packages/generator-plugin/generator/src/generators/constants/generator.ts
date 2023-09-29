import { formatFiles, generateFiles, Tree } from '@nx/devkit';
import {join} from 'path';
import { ConstantsGeneratorSchema } from './schema';
import { Constant, constants } from './constants';

type ConstantFileMetadata = {
  lang: string;
  root: string;
  path: string;
};

export const ConstantFiles: ConstantFileMetadata[] = [
  {
    lang: 'go',
    root: 'go-observability-kit',
    path: 'constants',
  },
  {
    lang: 'typescript',
    root: 'node-observability-kit',
    path: 'src',
  },
  {
    lang: 'python',
    root: 'python-observability-kit',
    path: '',
  },
];

export async function constantsGenerator(
  tree: Tree,
  _options: ConstantsGeneratorSchema
) {
  ConstantFiles.forEach((file) => {
    const projectRoot = join('packages', file.root, file.path);
    generateFiles(tree, join(__dirname, 'files', file.lang), projectRoot, {
      constants: constants.map(toSnakeUpperCase),
    });
  });
  await formatFiles(tree);
}

function toSnakeUpperCase(constant: Constant): Constant {
  const newName = constant.name
    .match(/[A-Z]{2,}(?=[A-Z][a-z]+[0-9]*|\b)|[A-Z]?[a-z]+[0-9]*|[A-Z]|[0-9]+/g)
    ?.map((x) => x.toUpperCase())
    .join('_');

  return {
    name: newName || constant.name,
    value: constant.value,
  };
}

export default constantsGenerator;
