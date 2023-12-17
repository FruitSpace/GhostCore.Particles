const viable = async ()=> {
    console.log("Hi")
    let r = await fetch('https://ya.ru').then(e=>e.headers.get("Location"))
    console.log(r)
}

viable()