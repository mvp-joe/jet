/**
 * @param {AdvancedUser} input
 * @returns {AdvancedResult}
 */
function main(input) {
    const people = []
    const ages = []
    walkFriend(input, friend => {
        people.push({name: friend.name, age: friend.age})
        ages.push(friend.age)
    })
    return {
        people,
        ages,
    }
}

/**
 * @param {AdvancedUser} friend
 * @param {function(AdvancedUser)} callback
 */
function walkFriend(friend, callback) {
    callback(friend)
    friend.friends?.forEach(friend => {
        walkFriend(friend, callback)
    });
}
