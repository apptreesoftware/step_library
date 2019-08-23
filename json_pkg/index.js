#!/usr/bin/env node
'use strict';

let apptree = require('apptree-workflow-sdk');

apptree.addStep('parse', '1.0', parse);
apptree.addStep('parse_array', '1.0', parseArray);
apptree.addStep('set_all', '1.0', setAll);
apptree.addStep('set', '1.0', set);
apptree.run();


function parse(input) {
  var out = JSON.parse(input["text"]);
  return { "record": out };
}

function parseArray(input) {
  var out = JSON.parse(input["text"]);
  return { "records": out };
}

function setAll(input) {
  const records = input['records'] || [];
  const fieldName = input['key'];
  const value = input['value'];

  records.forEach((r) => r[fieldName] = value);
  return { "records": records };
}

function set(input) {
  var record = input['record'] || {};
  var fieldName = input['key'];
  var value = input['value'];

  record[fieldName] = value;
  return { "record": record };
}
