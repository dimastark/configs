import * as path from 'path';

import { exec } from 'shelljs';

import { SystemTool } from '../SystemTool';

const bundlePath = path.join(__dirname, 'Brewfile');

export const brew: SystemTool = {
    order: 0,

    backup() {
        return exec(`brew bundle dump --file="${bundlePath}"`).code;
    },

    restore() {
        return exec(`brew bundle --file="${bundlePath}"`).code;
    },

    cleanup() {
        return (
            exec(`brew bundle cleanup --file="${bundlePath}"`).code ||
            exec('brew upgrade').code ||
            exec('brew cask upgrade').code ||
            exec('brew cleanup -s').code ||
            exec('brew cask cleanup').code
        );
    }
};
