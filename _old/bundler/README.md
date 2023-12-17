## GhostCore.Modules Bundler

While GhostCore Module can be a single file fully compatible with ECMAScript 5.1 standards and therefore GC.M JS Engine,
it's assumed that typical project would use modern javascript (or even typescript) and will rely on many npm packages.

Many packages use different ES standards and usually have the following features:
- ES2016+ imports: `import module from '@author/module'`
- Async event loops and promises
- Callbacks and handlers
- and so on

That introduces necessity of transpiling and bundling code to be compatible with CommonJS. For that purposes `babel` 
and `esbuild` are used.


### Initial project requirements

GC.M JS Engine ships with custom runtime and connectors, for that bundler is needed to verify compatibility
and bundle runtime with your project.

- Firstly initialize your project by
```shell
$ npm init
```
- Install bundler and add target to your package.json file
```shell
$ # something like npm i ghostcore-bundler
gcmbundle init
```

That will create GC.M config file
`gcm-config.json`
```json
{
  "script": "Script Name",
  "version": "0.1",
  "author": "Author Name",
  
  "entrypoint": "src/index.js",
  "npm": true
  
}
```
Where `entrypoint` is obviously your main file that will run and `npm` is a parameter that tells if bundler needs
to fetch npm packages listed in package.json before bundling (Usually that is used when building by CI or FruitSpace
pipelines).

`script`, `version`, `author` are self-explanatory

Also, it will add target
```json
  "build": "gcmbundle bundle"
```

Default output is at `dist/gcmodule.js`