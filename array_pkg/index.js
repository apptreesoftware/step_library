let apptree = require('apptree-workflow-sdk');

apptree.addStep('find', '1.0', findObject);
apptree.addStep('insert_object', '1.0', insertObject);
apptree.addStep('remove_object', '1.0', removeObject);

apptree.run();

function findObject(inputs) {
    apptree.validateInputs('Records', 'MatchField','MatchValue');
    const array = inputs['Records'];
    const field = inputs['MatchField'];
    const value = inputs['MatchValue'];
    let object = {};

    array.forEach(function(item){
        const objectVal = item[field];
        if (apptree.debug) {
            console.log(`Comparing ${objectVal} == ${value}`);
        }
        if(objectVal === value){
            object = item;
        }
    });

    return { 'Record': object };
}

function insertObject(inputs) {
    apptree.validateInputs('Records', 'Object');
    const array = inputs['Records'];
    const object = inputs['Object'];

    array.push(object);

    return { 'Records': array };
}

function removeObject(inputs) {
    apptree.validateInputs('Records');
    const array = inputs['Records'];
    const field = inputs['MatchField'];
    const value = inputs['MatchValue'];
    const index = inputs['Index'];
    let result = [];

    if(index == null){
        if(field == null){
            throw("If not removing by Index, MatchField must be provided");
        }
        if(value == null){
            throw("If not removing by Index, MatchValue must be provided");
        }
        array.forEach(function(item){
            if(item[field] === value){
                result = arrayRemove(array, item);
            }
        });
    }else {
        result = array;
        result.splice(index, 1);
    }

    return { 'Records': result };
}

function arrayRemove(arr, value) {

    return arr.filter(function(ele){
        return ele !== value;
    });
}