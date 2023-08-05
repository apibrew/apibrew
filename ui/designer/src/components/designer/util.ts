import { Resource } from '@apibrew/client';
import { AppDesignerBoard } from '../../model/app-designer-board';


export function checkResourceAllowedOnBoard(board: AppDesignerBoard, resource: Resource): boolean {
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

            if (selector === resource.namespace.name) {
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