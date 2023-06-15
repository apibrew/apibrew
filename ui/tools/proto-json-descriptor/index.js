#! /usr/bin/env node

const protobufjs = require("protobufjs")
const walk = require("./walk").walk

const args = process.argv.slice(2)

if (args.length < 1) {
    console.log("Usage: node index.js <protofile | protofiledirectory>")
    process.exit(1)
}

const protofile = args[0]

function loadProtoFile(protofile) {
    return new Promise((resolve, reject) => {
        protobufjs.load(protofile, (err, root) => {
            if (err) {
                reject(err)
            } else {
                resolve(root)
            }
        })
    })
}

function convertProtoFileToDescriptor(protofile) {
    return loadProtoFile(protofile).then(root => {
        return JSON.stringify(root.toJSON(), null, 4)
    })
}

if (protofile.endsWith(".proto")) {
    convertProtoFileToDescriptor(protofile).then(descriptor => {
        console.log(descriptor)
    }).catch(err => {
        console.error(err)
        process.exit(1)
    })
} else {
    walk(protofile, function (err, results) {
        results = results.filter(item => item.endsWith('.proto'))
        convertProtoFileToDescriptor(results).then(descriptor => {
            console.log(descriptor)
        }).catch(err => {
            console.error(err)
            process.exit(1)
        })
    })
}
