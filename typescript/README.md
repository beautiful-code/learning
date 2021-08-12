# Typescript Exercises

* Open `learning/typescript` directory in VSCode.
* Run `npm install` to install the dependencies.
* Typescript exercises are present in `learning/typescript/src` folder.
* Understand the `Objectives` present at the top of the each file and follow the `Instructions` to complete the exercises.

## Steps for checking Typescript Errors and Running the Exercises
* Go to `learning/typescript/tsconfig.json`.
* Before starting an exercise uncomment it from the tsconfig's `include` array.
* Run `npm run -s tsc-watch`. This will watch typescript files and re-compile them whenever they are modified.
  If there any typescript errors it will show them in the terminal.
* Associated JS file will be emitted in `learning/typescript/dist` directory. Run it using the below command.
  ```
  node ./dist/<file_name>.js

  Example:
  node ./dist/index.js
  ```