import { connect, StringCodec } from "nats";

const nc = await connect({ servers: ["localhost:4222"] });

// create a codec
const sc = StringCodec();

// subject/topic
const subject = "faas";

const addition = {
  operation: "+",
  operand1: 30,
  operand2: 12
};

const subtraction = {
  operation: "-",
  operand1: 50,
  operand2: 8
};

const multiplication = {
  operation: "*",
  operand1: 21,
  operand2: 2
};

const division = {
  operation: "/",
  operand1: 84,
  operand2: 2
};

async function wasmCalc(operation) {
  let jsonString = JSON.stringify(operation)
  
  await nc.request(subject, sc.encode(jsonString), { timeout: 1000 })
    .then((message) => {
      let result = sc.decode(message.data)
      console.log("🟢 Result:", result);
    })
    .catch((err) => {
      console.log("🔴 Error:", err.message);
    });

}

  // the client makes a request and receives a promise for a message
  // by default the request times out after 1s (1000 millis) and has
  // no payload.

await wasmCalc(addition)
await wasmCalc(subtraction)
await wasmCalc(multiplication)
await wasmCalc(division)

await nc.close();

