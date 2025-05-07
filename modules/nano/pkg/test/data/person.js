const Person = resource('Person')

Person.beforeCreate((person, event) => {
    if (!person.lastName) {
        person.lastName = 'Unknown'
    }

    return person
})

Person.afterCreate(person => {
    console.log('Person created: ', person.firstName)
})

Person.afterCreate(person => {
    if (person.firstName === 'PreventDelete') {
        console.log('Preventing delete')
        return false;
    }
})