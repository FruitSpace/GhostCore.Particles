# GCModules Runtime

Communication between JS runtimes and core itself is done through `GCRuntime` which introduces following functions to VM:
> `gcruntime.on` event listener is not really a javascript Event/CustomEvent, but rather a hook that will register
callback inside GhostCore and will be triggered as soon as matching event is fired
> 
> `callback` can return modified data (useful for preprocessors) or nothing
```ts
gcruntime.on = (
    module: string, //relay event only for specified module
    event: string,  //relay event only for specified type
    callback: (data: object) => (object | void)
): void => {/*@__PURE__*/}


//Example that recieves event:
gcruntime.on("pong", "ping", (e) => console.log(e))
//Example that modifies event:
gcruntime.on("pong", "ping", (e) => ({...e, value: 17}))
```

> `gcruntime.emit` sends event to all runtimes and core itself
> 
> `emit` cannot return any values or have callbacks
```ts
gcruntime.emit = (
    event: string,
    data: object
): void => {/*@__PURE__*/}


// Example that sends event (module "pong"):
gcruntime.emit("ping", {ping: "pong", value: 10})
```

### Why only one-way events?

Imagine a situation when there is one `ping` emitter and 10 receivers, 8 of which modifies data. Which value to return?

A much better way of implementing two-way communication is through emitters and listeners on both sides:
```js
//pong.js
gcruntime.emit("ping", "Hi!") //pong module

gcruntime.on("*", "pong", (e)=>console.log("It responded!"))
```
```js
//responder.js
gcruntime.on("pong", "ping", (e)=>{
    gcruntime.emit("pong", e+" - and I say 'Hello!'")
})
```
Let's see what just happened:
1) `pong` module sent `ping` event with `Hi!` as it's body
2) `responder` module got this event because
   - it listens only for `pong` module events only
   - it listens only for `ping` events
3) and `responder` sends `pong` event with `Hi! - and I say 'Hello!'` as it's body
4) `pong` module got this event because
    - it uses `*` wildcard and doesn't care who sent event
    - it listens only for `pong` event
5) `pong` module logs this historical event to console

### Omg but I can just listen on all-wildcards and capture all events
Yes. I mean these are your scripts, you should know what they are doing. Also you can use `gcruntime.on('*','*',()=>{})`
for debugging purposes

### Wait, but why callbacks can return modified data then?
Well, while runtimes are isolated and opaque for GhostCore, the core itself may emit some `:before` events, which data 
is expected to be modified, like `levelName`.

Let's make a filter for bad words as an example
```js
gcruntime.on('gc','onLevelUpload:before', (data)=>{
    let name = data.name
    name = name.replace(/(.+)(bad)(.+)/, "$1***$3")
    return {...data, name: name}
})
```
> Note: only `gc` module events support data modification. Also, you can't name your module `gc`
 


---
## Obsolete


### LevelPacks

- [ ] `onGauntletNew` - invoked when new gauntlet is created
- [ ] `onMapPackNew` - invoked when new map pack is created

### Communication

- To Be Done

### Essential

```go
package modules

// Plugin describes Plugin type; PluginCore is also Plugin
type Plugin interface {
	PreInit(*PluginCore, ...interface{})
	Unload(...interface{})
}

// PreInit Invoked to load anything
func PreInit(pch *PluginCore, args ...interface{})

// Unload Unloads everything
func Unload(args ...interface{})
```

### Player

```go
package modules

// OnPlayerNew Invoked when player is registered, but not yet activated account
func OnPlayerNew(uid int, uname string, email string)

// OnPlayerActivate Invoked when player first activated account
func OnPlayerActivate(uid int, uname string)

// OnPlayerLogin invoked when player commits login (regular, not gjp)
func OnPlayerLogin(uid int, uname string)

// OnPlayerBackup invoked when player uploads their backup
func OnPlayerBackup(uid int, decryptedBackup string)

//onPlayerSync is forbidden

// OnPlayerScoreUpdate invoked when player updates their score
func OnPlayerScoreUpdate(uid int, uname string, stats map[string]int)

var stats = map[string]int{
	"stars": 0,
	"diamonds": 0,
	"coins": 0,
	"ucoins": 0,
	"demons": 0,
	"cpoints": 0,
	"orbs": 0,
	"moons": 0,
	"special": 0,
	"lvlsCompleted": 0,
}
```

### Level

```go
package modules

import "strconv"

// OnLevelUpload invoked when level was uploaded
func OnLevelUpload(id int, name string, builder string, desc string)

// OnLevelUpdate invoked when level was updated
func OnLevelUpdate(id int, name string, builder string, desc string)

// OnLevelDelete invoked when level was deleted
func OnLevelDelete(id int, name string, builder string)

// OnLevelRate invoked when level was rated/rerated
func OnLevelRate(id int, name string, builder string, stars int, likes int, downloads int, length int, demonDiff int, isEpic bool, isFeatured bool, ratedBy map[string]string)

// OnLevelReport invoked when level was reported
func OnLevelReport(id int, name string, builder string, player string)

// OnLevelScore invoked when player published their score in level scoreboard
func OnLevelScore(id int, name string, player string, percent int, coins int)

var ratedBy = map[string]string{
	"uid":   strconv.Itoa(1),
	"uname": "Username",
}
```

### Communtication

```go
package modules

// Not implemented yet
```