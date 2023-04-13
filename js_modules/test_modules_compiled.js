// js_modules/test_modules_lib.js
var kafif = ":kafif:";
var test_modules_lib_default = kafif;

// js_modules/test_modules.js
function fNumber(value) {
  return parseInt(value);
}
function fString(value) {
  return test_modules_lib_default;
}
function fObject(value) {
  return {
    key: test_modules_lib_default
  };
}
function fFunction(value) {
  return function() {
    return test_modules_lib_default;
  };
}
