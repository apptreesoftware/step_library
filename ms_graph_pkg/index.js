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
    const url = buildUrl()

    let response = await axios.get(url, createConfig());

    let users = response.data.value;

    return {'Users' : users};
}

function formatDate(stringDate) {
    let date = new Date(stringDate);
    let hours = date.getHours();
    let minutes = date.getMinutes();
    let ampm = hours >= 12 ? 'pm' : 'am';
    hours = hours % 12;
    hours = hours ? hours : 12; // the hour '0' should be '12'
    minutes = minutes < 10 ? '0'+minutes : minutes;
    let strTime = hours + ':' + minutes + ' ' + ampm;
    return date.getMonth()+1 + "/" + date.getDate() + "/" + date.getFullYear() + "  " + strTime;
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

function buildUrl(endpoint, endPoint, filter){
    const host = 'https://graph.microsoft.com/v1.0/';
    let url = `${host}${endPoint}?$${filter}`;
    return url;
}

function createConfig(login){
    let config = {};
    if(login){
        config = {
            headers: {'Content-Type': 'application/x-www-form-urlencoded'}
        };
    }else{
        config = {
            headers: {'Content-Type': 'application/json'}
        };
    }
    return config;
}