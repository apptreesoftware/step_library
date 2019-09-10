let apptree = require('apptree-workflow-sdk');

apptree.addStep('new', '1.0', newObject);
apptree.addStep('update', '1.0', updateObject);

apptree.run();

function newObject(inputs) {
    let fields = inputs['Fields'];
    return { 'Record': fields };
}

function updateObject(inputs) {
    apptree.validateInputs('Record', 'Fields');

    let fields = inputs['Fields'];
    let object = inputs['Record'];
    if (object == null) {
        throw "You must provide the input `Record`";
    }

    let keys = Object.keys(fields);
    keys.forEach((k) => {
        if (apptree.debug) {
            console.log(`Setting ${k} = ${fields[k]}`);
        }
        object[k] = fields[k];
    });
    return { 'Record': object };
}