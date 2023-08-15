interface AdvancedUser {
    name: string
    age: number
    friends: AdvancedUser[]
}

interface AdvancedResult {
    people: { name: string, age: number}[]
    ages: number[]
}
