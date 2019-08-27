let apptree = require('apptree-workflow-sdk');

apptree.addStep('new', '1.0', newObject);
apptree.addStep('update', '1.0', updateObject);

apptree.run();

function newObject(inputs) {
    var fields = inputs['Fields'];
    return { 'Record': fields };
}

function updateObject(inputs) {
    var fields = inputs['Fields'];
    var object = inputs['Record'];

    if (object == null) {
        process.stderr.write("You must provide the input `Record`");
        return;
    }

    var keys = Object.entries(fields);
    for (var key in keys) {
        object.key = fields[key];
    }
    return { 'Record': object };
}