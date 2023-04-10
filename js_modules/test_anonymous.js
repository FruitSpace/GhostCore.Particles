

const fNumber = (value)=>{
    return parseInt(value)
}

const fString = (value)=>{
    return value.toString()
}

const fObject = (value)=>{
    return {
        key: value
    }
}

const fFunction = (value)=>{
    return ()=>value
}