"use strict";

const WASM_URL = "wasm.wasm";
let wasm;

const initWasm = () => {
  const go = new Go();

  if ("instantiateStreaming" in WebAssembly) {
    WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(
      function (obj) {
        wasm = obj.instance;
        go.run(wasm);
      }
    );
  }
};

initWasm();
