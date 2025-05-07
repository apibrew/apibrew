const Person = resource('PersonTypescript')

Person.beforeCreate((person: any, event: any) => {
    if (!person.lastName) {
        person.lastName = 'Unknown'
    }

    return person
})

Person.afterCreate((person: any) => {
    console.log('Person created: ', person.firstName)
})

Person.afterCreate((person: any) => {
    if (person.firstName === 'PreventDelete') {
        console.log('Preventing delete')
        return false;
    }
})