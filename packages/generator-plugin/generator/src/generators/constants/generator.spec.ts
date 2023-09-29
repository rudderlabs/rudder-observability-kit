import { createTreeWithEmptyWorkspace } from '@nx/devkit/testing';
import { Tree } from '@nx/devkit';

import { constantsGenerator, ConstantFiles } from './generator';
import { ConstantsGeneratorSchema } from './schema';

describe('constants generator', () => {
  let tree: Tree;
  const options: ConstantsGeneratorSchema = {};

  beforeEach(() => {
    tree = createTreeWithEmptyWorkspace();
  });

  it('should run successfully', async () => {
    await constantsGenerator(tree, options);

    const changes = tree.listChanges();
    const constantFileChanges = changes.filter(
      (c) => c.type === 'CREATE' && c.path.includes('constants')
    );
    ConstantFiles.every((file) => {
      const constantFileChange = constantFileChanges.find((c) =>
        c.path.includes(file.root)
      );
      expect(constantFileChange).toBeTruthy();
    });
  });
});
