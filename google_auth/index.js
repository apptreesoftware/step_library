let apptree = require('apptree-workflow-sdk');
const { auth } = require("google-auth-library");
apptree.addStep('get_token', '1.0', getToken);
apptree.run();


async function getToken(input) {
    apptree.validateInputs('Credential', 'Scopes');
    let credential = input['Credential'];
    let scopes = input['Scopes'];
    const client = auth.fromJSON(credential);
    client.scopes= scopes;
    const token = await client.getAccessToken();
    return {"Token" : token.token};
}
