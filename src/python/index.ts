import * as path from 'path';

import { exec } from 'shelljs';

import { SystemTool } from '../SystemTool';

const requirementsPath = path.join(__dirname, 'requirements.txt');

export const python: SystemTool = {
    order: 2,

    backup() {
        return exec(`pip3 freeze > ${requirementsPath}`).code;
    },

    restore() {
        return (
            exec(`pip3 install -r ${requirementsPath}`).code ||
            this.cleanup()
        );
    },

    cleanup() {
        return (
            exec(`pip3 freeze | grep -v -f ${requirementsPath} - | xargs pip3 uninstall -y`).code ||
            exec('pip3 freeze | grep -v "^\\-e" | cut -d = -f 1  | xargs -n1 pip3 install -U').code
        );
    }
};
