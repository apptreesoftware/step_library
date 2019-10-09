let apptree = require('apptree-workflow-sdk');

apptree.addStep('length', '1.0', getLength);
apptree.addStep('trimleft', '1.0', trimLeftString);
apptree.addStep('trimright', '1.0', trimRightString);
apptree.addStep('trim', '1.0', trimString);
apptree.addStep('substring', '1.0', substringString);
apptree.addStep('concat', '1.0', concatString);
apptree.addStep('indexof', '1.0', indexOfString);
apptree.run();

function getLength(inputs) {
    apptree.validateInputs('String');
    const string = inputs['String'];

    return { 'Count': string.length };

}

function trimLeftString(inputs) {
    apptree.validateInputs('String');
    var string = inputs['String'];

    return { 'String': string.trimLeft() };

}

function trimRightString(inputs) {
    apptree.validateInputs('String');
    var string = inputs['String'];

    return { 'String': string.trimRight() };

}

function trimString(inputs) {
    apptree.validateInputs('String');
    var string = inputs['String'];

    return { 'String': string.trim() };

}

function substringString(inputs) {
    apptree.validateInputs('String');
    const string = inputs['String'];
    let startIndex = inputs['StartIndex'];
    let endIndex = inputs['EndIndex'];

    if (startIndex == null) {
        startIndex = 0;
    }
    if (endIndex == null) {
        endIndex = string.length;
    }
    return { 'String': string.substring(startIndex, endIndex) };

}

function concatString(inputs) {
    apptree.validateInputs('String1', 'String2');
    var string1 = inputs['String1'];
    var string2 = inputs['String2'];

        return { 'String': string1.concat(string2) };

}

function indexOfString(inputs) {
    apptree.validateInputs('String', 'SearchString');
    const string = inputs['String'];
    const searchString = inputs['SearchString'];
    let startPosition = inputs['StartPosition'];

    if (startPosition == null) {
        startPosition = 0;
    }

    return { 'StringPosition': string.indexOf(searchString, startPosition) };

}