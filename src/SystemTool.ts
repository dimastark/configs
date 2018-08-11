export interface SystemTool {
    order: number;

    backup(): number;
    restore(): number;
    cleanup?(): number;
}
