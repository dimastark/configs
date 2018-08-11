import { brew } from './brew';
import { fish } from './fish';
import { git } from './git';
import { node } from './node';
import { python } from './python';

const command = process.argv[2];
const tools = { brew, fish, git, node, python };
const order = Object.keys(tools)
    .sort((a, b) => tools[a].order - tools[b].order);

for (const toolName of order) {
    const tool = tools[toolName];

    if (!(command in tool)) {
        continue;
    }

    console.log(`> ${toolName} ${command}`);

    const code = tool[command]();

    if (code !== 0) {
        process.exit(code);
    }
}
