import fs from 'fs';

export const readFileContent = async (filePath: string): Promise<string | ''> => {
  let result = '';
  if ((await fs.promises.access(filePath), fs.constants.F_OK)) {
    result = await fs.promises.readFile(filePath, 'utf8');
  }
  return result;
};
