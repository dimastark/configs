import * as path from 'path';

import { cp } from 'shelljs';

import { SystemTool } from '../SystemTool';

const fishConfigPath = path.join(__dirname, 'config.fish');
const fishFunctionsPath = path.join(__dirname, 'functions');

export const fish: SystemTool = {
    order: 1,

    backup() {
        cp('~/.config/fish/config.fish', fishConfigPath);
        cp('-R', '~/.config/fish/functions', __dirname);

        return 0;
    },

    restore() {
        cp(fishConfigPath, '~/.config/fish/config.fish');
        cp('-R', fishFunctionsPath, '~/.config/fish/functions');

        return 0;
    }
};
