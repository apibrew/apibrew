const Human = resource("Human")

Human.beforeUpdate(human => {
    if (human.age > 100) {
        throw new Error("age must be less than 100")
    }
})