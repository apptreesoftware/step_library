//Steps to access the Microsoft Graph API
// 10.3.2019
// abe.foster@apptreesoftware.com
let apptree = require('apptree-workflow-sdk');
let axios = require('axios');
let qs = require('querystring');
apptree.addStep('get_auth_token', '1.0', getAuthToken);
apptree.addStep('get_users', '1.0', getUsers);

apptree.run();

async function getUsers(inputs) {
    apptree.validateInputs('AuthToken', 'Filter');
    const filter = inputs['Filter'];
    const authToken = inputs['AuthToken'];

    let users = await getAll(authToken,'users', filter);

    return {'Users' : users};
}

async function getAll(token, endPoint, filter){
    let objects = [];
    let hasNextLink = false;
    let nextLink = '';
    const url = buildUrl(endPoint, filter);
    let response = await axios.get(url, createConfig(false, token));
    if(response.data){
        objects.push(response.data.value);
    }
    if(response.data['@odata.nextLink']){
        hasNextLink = true;
        nextLink = response.data['@odata.nextLink'];
    }
    while(hasNextLink){
        await new Promise(done => setTimeout(done, 200));
        let nextResponse = await axios.get(nextLink, createConfig(false, token));
        objects.push(nextResponse.data.value);
        if(nextResponse.data['@odata.nextLink']){
            hasNextLink = true;
            nextLink = nextResponse.data['@odata.nextLink'];
        }else{
            hasNextLink = false;
        }
    }
    return objects;
}

async function getAuthToken(inputs){
    apptree.validateInputs('Tenant', 'Password', 'Username', 'ClientId', 'ClientSecret', 'Scope');
    const tenant = inputs['Tenant'];
    const pw = inputs['Password'];
    const username = inputs['Username'];
    const clientId = inputs['ClientId'];
    const clientSecret = inputs['ClientSecret'];
    const scope = inputs['Scope'];
    const url = `https://login.microsoftonline.com/${tenant}/oauth2/v2.0/token`;
    const login = true;
    const creds = {
        username: username,
        password: pw,
        client_id: clientId,
        client_secret: clientSecret,
        grant_type: 'Password',
        scope: scope
    };

    let response = await axios.post(url, qs.stringify(creds), createConfig(login));

    let authToken = response.data['access_token'];

    return {'AuthToken' : authToken};
}

function buildUrl(endPoint, filter){
    const host = 'https://graph.microsoft.com/v1.0/';
    let url = `${host}${endPoint}?$${filter}`;
    return url;
}

function createConfig(login, token){
    let config = {};
    if(login){
        config = {
            headers: {'Content-Type': 'application/x-www-form-urlencoded'}
        };
    }else{
        config = {
            headers: {'Content-Type': 'application/json', 'Authorization': 'Bearer ' + token}
        };
    }
    return config;
}