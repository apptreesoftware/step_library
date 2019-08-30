let apptree = require('apptree-workflow-sdk');

apptree.addStep('get_templates', '1.0', getTemplates());

apptree.run();

let tokens = {};
tokens.clientToken = null;
tokens.apiToken = null;

function getTemplates(inputs) {
    apptree.validateInputs('HostUrl', 'ClientId', 'Secret', 'Username', 'Password');
    const host = inputs['HostUrl'];
    const clientId = inputs['ClientId'];
    const secret = inputs['Secret'];
    const username = inputs['Username'];
    const password = inputs['Password'];
    const endpoint = "platform/api/v1/webtemplates";
    const request = new XMLHttpRequest();
    const url = buildUrl(host, endpoint);
    let response = "";

     tokens = authorize(host, clientId, secret, username, password);

    request.open("GET", url);
    request.send();

    request.onreadystatechange=function(){
        if(this.readyState === XMLHttpRequest.DONE && this.status ===200){
            response = request.responseText;
            return{'Status': request.status, "Records": request.body};
        }else{
            return{"Status": request.status, "Message": request.responseText};
        }
    }
}

function authorize(host, clientId, secret, username, password){
    let clientResponse = getClientToken(host, clientId, secret);
    if(clientResponse['Status'] === "Success"){
        getAuthToken(host, username, password, tokens.clientToken);
    }
}

function getClientToken(host, clientId, secret){
    const url = buildUrl(host, "/platform/api/v1/clientauthentication");
    const creds = {clientId: clientId, secret: secret};
    const clientRequest = new XMLHttpRequest();

    clientRequest.open("POST", url);
    clientRequest.setRequestHeader("Content-Type", "application/json");

    clientRequest.onreadystatechange=function() {
        if (this.readyState === XMLHttpRequest.DONE && this.status === 200) {
            var json = JSON.parse(clientRequest.responseText);
            tokens.clientToken = json.clientToken;
            return{"Status": "Success"};
        } else {
            return {"Status": clientRequest.status, "Message": clientRequest.responseText};
        }
    };
    clientRequest.send(JSON.stringify(creds));
}

function getAuthToken(host, username, pw, token){
    const url = buildUrl(host, "/platform/api/v1/authentication");
    const emsToken = token;
    const authRequest = new XMLHttpRequest();
    const creds = {username: username, password: pw};

    authRequest.open("POST", url);
    authRequest.setRequestHeader("Content-Type", "application/json");
    authRequest.setRequestHeader("x-ems-api-token", emsToken);

    authRequest.onreadystatechange=function() {
        if (this.readyState === XMLHttpRequest.DONE && this.status === 200) {
            var json = JSON.parse(authRequest.responseText);
            tokens.authToken = json['x-ems-api-token'];
            return{"Status": "Success"}
        } else {
            return {"Status": clientRequest.status, "Message": clientRequest.responseText};
        }
    };

    authRequest.send(JSON.stringify(creds));
}

function buildUrl(host, endpoint){
    let url = host + endpoint;
    return url;
}