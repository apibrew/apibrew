const Person = resource('Person')
const Human = resource('Human')

const humanToPerson = item => {
    const parts = item.name.split(' ')
    return {
        firstName: parts[0],
        lastName: parts[1],
        id: item.id,
    }
}

const personToHuman = item => {
    return {
        name: `${item.firstName} ${item.lastName}`,
        id: item.id,
    }
}

Human.bindCreate(Person, personToHuman, humanToPerson)
Human.bindUpdate(Person, personToHuman, humanToPerson)
Human.bindGet(Person, personToHuman, humanToPerson)
Human.bindList(Person, personToHuman, humanToPerson)
Human.bindDelete(Person, personToHuman, humanToPerson)
