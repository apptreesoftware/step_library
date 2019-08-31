//Steps to access the EMS WebServices
// 8.30.2019
// abe.foster@apptreesoftware.com
let apptree = require('apptree-workflow-sdk');
let axios = require('axios');
apptree.addStep('get_templates', '1.0', getTemplates);

apptree.run();

async function getTemplates(inputs) {
    apptree.validateInputs('HostUrl', 'ClientId', 'Secret');
    const host = inputs['HostUrl'];
    const clientId = inputs['ClientId'];
    const secret = inputs['Secret'];
    const endpoint = "/platform/api/v1/webtemplates";
    const url = buildUrl(host, endpoint);
    let templates = [];

    let clientToken = await getClientToken(host, clientId, secret);

    let config = {
        headers: {'x-ems-api-token': clientToken}
    };
    await axios.get(url, config)
        .then(function(response){
            templates = response.data.results;
        })
        .catch(function (error) {
            throw(error);
        });
    return {'Templates' : templates};
}

async function getClientToken(host, clientId, secret) {
    const url = buildUrl(host, "/platform/api/v1/clientauthentication");
    const creds = {clientId: clientId, secret: secret};
    let clientToken = "";
    let config = {
        headers: {'Content-Type': 'application/json'}
    };
    await axios.post(url, JSON.stringify(creds), config)
        .then(function (response) {
            clientToken = response.data['clientToken'];
        })
        .catch(function (error) {
            throw(error);
        });
    return clientToken;
}

//Use this if the Call requires an Auth Token
async function getAuthToken(host, username, pw, token){
    const url = buildUrl(host, "/platform/api/v1/authentication");
    const emsToken = token;
    const creds = {username: username, password: pw};
    let authToken = "";

    let config = {
        headers: {'Content-Type': 'application/json','x-ems-api-token': emsToken}
    };
    await axios.post(url, JSON.stringify(creds), config)
        .then(function(response){
            authToken = JSON.parse(response.data['webToken']);
        })
        .catch(function (error) {
            throw(error);
        });
    return authToken;
}

function buildUrl(host, endpoint){
    let url = host + endpoint;
    return url;
}