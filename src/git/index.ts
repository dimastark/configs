import * as path from 'path';

import { cp } from 'shelljs';

import { SystemTool } from '../SystemTool';

const gitConfigPath = path.join(__dirname, 'gitconfig');
const gitIgnorePath = path.join(__dirname, 'gitignore');

export const git: SystemTool = {
    order: 1,

    backup() {
        cp('~/.gitconfig', gitConfigPath);
        cp('~/.gitignore', gitIgnorePath);

        return 0;
    },

    restore() {
        cp(gitConfigPath, '~/.gitconfig');
        cp(gitIgnorePath, '~/.gitignore');

        return 0;
    }
};
