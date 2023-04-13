// Provided context this

const gcruntime = {}


gcruntime.on = (module, event, callback) => {
    // Imported Go function
    // Usage example:
    gcruntime.on("calculator", "userStarsCalculated", (data)=>{
        fetch("<webhook>", {method: "POST", body: JSON.stringify({
                message: `${data.user} reached ${data.stars}. Congratulations!`
            })}).catch()
    })
}

gcruntime.emit = (event, data) => {
    // Imported Go function, will call all gcruntime.on functions in all runtimes
    // Usage example:
    // ... some math stuff maybe
    gcruntime.emit("userStarsCalculated", {
        user: "GhostCat",
        stars: 500
    })
}


export default gcruntime