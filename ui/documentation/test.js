const packets = [
    {
        "name": "n1",
        "type": "output",
    },
    {
        "name": "n1",
        "type": "output",
    },
    {
        "name": "n1",
        "type": "output",
    },
    {
        "name": "n1",
        "type": "output",
    },
    {
        "name": "n1",
        "type": "output",
    },
    {
        "name": "n11",
        "type": "input",
    },
    {
        "name": "n122",
        "type": "input",
    },
    {
        "name": "n13",
        "type": "input",
    },
    {
        "name": "n14",
        "type": "input",
    },
    {
        "name": "n15",
        "type": "output",
    },
    {
        "name": "n16",
        "type": "input",
    },
    {
        "name": "n17",
        "type": "input",
    },
    {
        "name": "n18",
        "type": "input",
    },
    {
        "name": "n19",
        "type": "input",
    },
    {
        "name": "n111",
        "type": "input",
    },
    {
        "name": "n112",
        "type": "input",
    }
]



let firstInput;
let lastInput;

for (const packet of packets) {
    // packetsin icindeki obyekti tipe beraberleshd

   // console.log(packet, packet.type, 'input', packet.type == 'input')
    if (packet.type == 'input') {
        // neyi packet
        // neye firstInput
        lastInput = packet

        console.log(firstInput)

        // firstInput.push(packet.type)
        // packet.type.push(firstInput)
        console.log(packet)
    }
}

console.log('NETICE::')

console.log(firstInput)
console.log(lastInput)
