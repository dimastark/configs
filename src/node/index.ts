import { exec } from 'shelljs';

import { SystemTool } from '../SystemTool';

export const node: SystemTool = {
    order: 2,

    backup() {
        return 0;
    },

    restore() {
        return (
            exec('npm i -g npm').code
        );
    },

    cleanup() {
        return exec(
            'npm ls -gp --depth=0 | ' +
            'awk -F/ "/node_modules/ && !/\\/npm$/ {print $NF}" | ' +
            'xargs npm -g rm'
        ).code;
    }
};
