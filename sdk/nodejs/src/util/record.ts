export function label(record: any): string {
    const nameProps = ['name', 'username', 'title', 'label', 'id'];

    if (!record) {
        return ''
    }

    for (const nameProp of nameProps) {
        if (record[nameProp]) {
            return record[nameProp];
        }
    }

    if (typeof record === 'object') {
        return '- object -'
    }

    return record + '';
}