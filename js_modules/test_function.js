

function fNumber(value) {
    return parseInt(value)
}

function fString(value) {
    return value.toString()
}

function fObject(value) {
    return {
        key: value
    }
}

function fFunction(value) {
    return function () {
        return value
    }
}