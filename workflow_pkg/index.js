let apptree = require('apptree-workflow-sdk');

apptree.addStep('exit', '1.0', exitWorkflowSuccess);
apptree.addStep('fail', '1.0', exitWorkflowFail);

apptree.run();

function exitWorkflowSuccess() {
    if (apptree.debug) {
        console.log(`Exiting with code ${apptree.EXIT_WORKFLOW_SUCCESS_CODE}`);
    }
    process.exit(apptree.EXIT_WORKFLOW_SUCCESS_CODE);
}

function exitWorkflowFail() {
    if (apptree.debug) {
        console.log(`Exiting with code ${apptree.EXIT_WORKFLOW_FAIL_CODE}`);
    }
    process.exit(apptree.EXIT_WORKFLOW_FAIL_CODE);
}