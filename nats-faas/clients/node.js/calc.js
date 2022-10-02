import { connect, StringCodec } from "nats";

console.log("🤖 Connecting to the server...")
const nc = await connect({ servers: ["localhost:4222"] });

// create a codec
const sc = StringCodec();

// subject/topic
const subject = "faas";

const addition = {
  operation: "+",
  operand1: 30,
  operand2: 12
}

const subtraction = {
  operation: "-",
  operand1: 50,
  operand2: 8
}

const multiplication = {
  operation: "*",
  operand1: 21,
  operand2: 2
}

const division = {
  operation: "/",
  operand1: 84,
  operand2: 2
}

async function callCalcOperation(operation) {

  await nc.request(subject, sc.encode(JSON.stringify(operation)), { timeout: 800 })
    .then((message) => {
      console.log(`🟢 [${operation.operation}] Result from the wasm function: ${sc.decode(message.data)}`);
    })
    .catch((err) => {
      console.log(`🔴 Error with request: ${err.message}`);
    });

}

for (let step = 0; step < 100; step++) {
  console.log("🟠 Iteration number:", step)
  await callCalcOperation(addition)
  await callCalcOperation(subtraction)
  await callCalcOperation(multiplication)
  await callCalcOperation(division)
}

await nc.close();

