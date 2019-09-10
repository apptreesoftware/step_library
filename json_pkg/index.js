#!/usr/bin/env node
'use strict';

let apptree = require('apptree-workflow-sdk');

apptree.addStep('parse', '1.0', parse);
apptree.addStep('parse_array', '1.0', parseArray);
apptree.run();


function parse(input) {
    let out = JSON.parse(input["String"]);
    return {"Record": out};
}

function parseArray(input) {
    var out = JSON.parse(input["String"]);
    return {"Records": out};
}
