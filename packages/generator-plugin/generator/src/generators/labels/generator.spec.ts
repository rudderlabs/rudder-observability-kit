import { createTreeWithEmptyWorkspace } from '@nx/devkit/testing';
import { Tree } from '@nx/devkit';

import { labelsGenerator, LabelFiles } from './generator';
import { LabelsGeneratorSchema } from './schema';

describe('Labels generator', () => {
  let tree: Tree;
  const options: LabelsGeneratorSchema = {};

  beforeEach(() => {
    tree = createTreeWithEmptyWorkspace();
  });

  it('should run successfully', async () => {
    await labelsGenerator(tree, options);

    const changes = tree.listChanges();
    const labelFileChanges = changes.filter(
      (c) => c.type === 'CREATE' && c.path.includes('labels')
    );
    LabelFiles.every((file) => {
      const constantFileChange = labelFileChanges.find((c) =>
        c.path.includes(file.root)
      );
      expect(constantFileChange).toBeTruthy();
    });
  });
});
