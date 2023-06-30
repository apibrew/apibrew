import { Resource } from '@apibrew/ui-lib';
import { AppDesignerBoard } from "../../../model/ui/designer-board.ts";

export function checkResourceAllowedOnBoard(board: AppDesignerBoard, resource: Resource): boolean {
    console.log(board.resourceSelector, board.name, resource.name, resource.namespace)
    if (board.resourceSelector) {
        let found = false
        for (const selector of board.resourceSelector) {
            if (selector == '*') {
                found = true
                break
            }

            if (selector === resource.name) {
                found = true
                break
            }
        }

        if (!found) {
            return false
        }
    }

    if (board.namespaceSelector) {
        let found = false
        for (const selector of board.namespaceSelector) {
            if (selector == '*') {
                found = true
                break
            }

            if (selector === resource.namespace) {
                found = true
                break
            }
        }

        if (!found) {
            return false
        }
    }

    return true
}