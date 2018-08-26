import * as path from 'path';

import { exec, rm } from 'shelljs';

import { SystemTool } from '../SystemTool';

const bundlePath = path.join(__dirname, 'Brewfile');

export const brew: SystemTool = {
    order: 0,

    backup() {
        rm(bundlePath);

        return exec(`brew bundle dump --file="${bundlePath}"`).code;
    },

    restore() {
        return exec(`brew bundle --file="${bundlePath}"`).code;
    },

    cleanup() {
        return (
            exec('brew upgrade').code ||
            exec('brew cask upgrade').code ||
            exec('brew cleanup -s').code
        );
    }
};
