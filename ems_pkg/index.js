//Steps to access the EMS WebServices
// 8.30.2019
// abe.foster@apptreesoftware.com
let apptree = require('apptree-workflow-sdk');
let axios = require('axios');
apptree.addStep('get_templates', '1.0', getTemplates);
apptree.addStep('get_client_token', '1.0', getClientToken);
apptree.addStep('get_auth_token', '1.0', getAuthToken);
apptree.addStep('get_setup_types', '1.0', getSetupTypes);
apptree.addStep('get_timezones', '1.0', getTimezones);
apptree.addStep('get_room_types', '1.0', getRoomTypes);
apptree.addStep('get_buildings', '1.0', getBuildings);
apptree.addStep('search_buildings', '1.0', searchBuildings);
apptree.addStep('get_floors_by_buildingid', '1.0', getFloorsByBuildingId);
apptree.addStep('get_rooms_by_buildingid', '1.0', getRoomsByBuildingId);
apptree.addStep('get_event_types', '1.0', getEventTypes);
apptree.addStep('get_groups', '1.0', getGroups);
apptree.addStep('get_groups_by_webuserid', '1.0', getGroupsByWebUserId);
apptree.addStep('search_bookings_by_roomid', '1.0', searchBookingsByRoomId);
apptree.addStep('search_room_availability', '1.0', searchRoomAvailability);
apptree.addStep('create_reservation', '1.0', createReservation);
apptree.addStep('create_availability_markdown', '1.0', createAvailabilityMarkdown);

apptree.run();

async function createReservation(inputs) {
    apptree.validateInputs('AuthToken', 'HostUrl', 'EmailAddress', 'EventName',
        'GroupId', 'EventTypeId', 'ProcessTemplateId', 'RoomRecordType', 'Phone', 'BillingReference',
        'StartTime', 'EndTime', 'RoomId');
    const host = inputs['HostUrl'];
    const authToken = inputs['AuthToken'];
    const contactId = inputs['ContactId'];
    const comment = inputs['Comment'];
    const emailAddress = inputs['EmailAddress'];
    const eventName = inputs['EventName'];
    const roomId = inputs['RoomId'];
    const startTime = inputs['StartTime'];
    const endTime = inputs['EndTime']
    const groupId = inputs['GroupId'];
    const eventTypeId = inputs['EventTypeId'];
    const processTemplateId = inputs['ProcessTemplateId'];
    const roomRecordType = inputs['RoomRecordType'];
    const phone = inputs['Phone'];
    const billingReference = inputs['BillingReference'];
    const endpoint = "/platform/api/v1/reservations/actions/create";
    const url = buildUrl(host, endpoint);
    let bookings = [{endTime: endTime, startTime: startTime, roomId: roomId, roomRecordType: roomRecordType, eventTypeId: eventTypeId}];

    let data = {contactId: contactId, comment: comment, emailAddress: emailAddress, eventName: eventName, bookings: bookings,
        groupId: groupId, eventTypeId: eventTypeId, processTemplateId: processTemplateId, roomRecordType: roomRecordType,
        phone: phone, billingReference: billingReference};

    let response = await axios.post(url, data, createConfig(authToken));

    let bookingIds = response.data.id;

    return {'Bookings' : bookingIds};
}

async function searchRoomAvailability(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl', 'BuildingId', 'Attendance', 'RoomTypeIds',
        'FloorIds', 'Dates', 'EventStartTime', 'EventEndTime', 'ExcludeUnavailable', 'WebTemplateId', 'TimeZoneId');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const buildingId = inputs['BuildingId'];
    const attendance = inputs['Attendance'];
    const floorIds = inputs['FloorIds'];
    const roomTypeIds = inputs['RoomTypeIds'];
    const dates = inputs['Dates'];
    const eventStartTime = inputs['EventStartTime'];
    const eventEndTime = inputs['EventEndTime'];
    const excludeUnavailable = inputs['ExcludeUnavailable'];
    const webTemplateId = inputs['WebTemplateId'];
    const timeZoneId = inputs['TimeZoneId'];
    const endpoint = "/platform/api/v1/bookings/actions/search";
    const url = buildUrl(host, endpoint);

    let data = {buildingId: buildingId, attendance: attendance, roomTypeIds: roomTypeIds, floorIds: floorIds, dates: dates,
        eventStartTime: eventStartTime, eventEndTime: eventEndTime, excludeUnavailable: excludeUnavailable,
        webTemplateId: webTemplateId, timeZoneId: timeZoneId};

    let response = await axios.post(url, data, createConfig(clientToken));

    let rooms = response.data.results;

    return {'Rooms' : rooms};
}

async function searchBookingsByRoomId(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl', 'RoomId');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const roomIds = inputs['RoomId'];
    const endpoint = "/platform/api/v1/bookings/actions/search";
    const url = buildUrl(host, endpoint);
    let array = [roomIds]
    let d = new Date();
    d.setDate(d.getDate() - 2);
    let date = d.toISOString();
    let data = {roomIds: array, includeCancelled: false, minDateChanged: date};
    let bookings = [];
    try{
        let response = await axios.post(url, data, createConfig(clientToken));
        bookings = response.data.results;
        return {'Bookings' : bookings};
    }catch(err){
        return {'Bookings' : null};
    }
}

async function getFloorsByBuildingId(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl', 'BuildingId');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const buildingId = inputs['BuildingId'];
    const endpoint = `/platform/api/v1/groups?facilityId=${buildingId}`;
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let floors = response.data.results;

    return {'Floors' : floors};
}

async function getRoomsByBuildingId(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl', 'BuildingId');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const buildingIds = inputs['BuildingId'];
    const endpoint = `/platform/api/v1/rooms/actions/search`;
    const url = buildUrl(host, endpoint);
    let data = {buildingIds: [buildingIds]};

    let response = await axios.post(url, data, createConfig(clientToken));

    let rooms = response.data.results;

    return {'Rooms' : rooms};
}

async function getTemplates(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const endpoint = "/platform/api/v1/webtemplates";
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let templates = response.data.results;

    return {'Templates' : templates};
}

async function getGroupsByWebUserId(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl', 'WebUserId');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const webUserId = inputs['WebUserId'];
    const endpoint = `/platform/api/v1/groups?webUserId=${webUserId}`;
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let groups = response.data.results;

    return {'Groups' : groups};
}

async function getGroups(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const endpoint = "/platform/api/v1/groups";
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let groups = response.data.results;

    return {'Groups' : groups};
}

async function getEventTypes(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const endpoint = "/platform/api/v1/eventtypes";
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let eventTypes = response.data.results;

    return {'EventTypes' : eventTypes};
}

async function getBuildings(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const endpoint = "/platform/api/v1/buildings";
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let buildings = response.data.results;

    return {'Buildings' : buildings};
}

async function searchBuildings(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl', 'SearchText');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const searchText = inputs['SearchText'];
    const endpoint = `/platform/api/v1/buildings?searchText=${searchText}`;
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let buildings = response.data.results;

    return {'Buildings' : buildings};
}

async function getRoomTypes(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const endpoint = "/platform/api/v1/roomtypes?webEnabled=true";
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let roomTypes = response.data.results;

    return {'RoomTypes' : roomTypes};
}

async function getSetupTypes(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const endpoint = "/platform/api/v1/setuptypes";
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let types = response.data.results;

    return {'SetupTypes' : types};
}

async function getTimezones(inputs) {
    apptree.validateInputs('ClientToken', 'HostUrl');
    const host = inputs['HostUrl'];
    const clientToken = inputs['ClientToken'];
    const endpoint = "/platform/api/v1/timezones";
    const url = buildUrl(host, endpoint);

    let response = await axios.get(url, createConfig(clientToken));

    let timezones = response.data;

    return {'Timezones' : timezones};
}

async function createAvailabilityMarkdown(inputs) {
    apptree.validateInputs('Bookings');
    const bookings = inputs['Bookings'];
    let roomDes = "";
    if(bookings[0] !== undefined){
        let book = bookings[0];
        roomDes = book.room.description;
        bookings.forEach(function (book) {
            book.eventStartTime = formatDate(book.eventStartTime);
            book.eventEndTime = formatDate(book.eventEndTime);
        });
    }else{
        return null;
    }
    
    const markdown = `# ${roomDes}\n${bookings.map(b => `- ${b.eventStartTime} to ${b.eventEndTime}\n`).join('')}`;

    return {'Markdown' : markdown};
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

async function getClientToken(inputs) {
    apptree.validateInputs('HostUrl', 'ClientId', 'Secret');
    const host = inputs['HostUrl'];
    const clientId = inputs['ClientId'];
    const secret = inputs['Secret'];
    const url = buildUrl(host, "/platform/api/v1/clientauthentication");
    const creds = {clientId: clientId, secret: secret};

    let response = await axios.post(url, JSON.stringify(creds), createConfig());

    let clientToken = response.data['clientToken'];

    return {'ClientToken' : clientToken};
}

//Use this if the Call requires an Auth Token
async function getAuthToken(inputs){
    apptree.validateInputs('HostUrl', 'Password', 'Username', 'ClientToken');
    const host = inputs['HostUrl'];
    const pw = inputs['Password'];
    const username = inputs['Username'];
    const clientToken = inputs['ClientToken'];
    const url = buildUrl(host, "/platform/api/v1/authentication");
    const creds = {username: username, password: pw};

    let response = await axios.post(url, JSON.stringify(creds), createConfig(clientToken));

    let authToken = response.data['webToken'];

    return {'AuthToken' : authToken};
}

function buildUrl(host, endpoint){
    let url = host + endpoint;
    return url;
}

function createConfig(clientToken){
    let config = {};
    if(clientToken){
        config = {
            headers: {'Content-Type': 'application/json','x-ems-api-token': clientToken}
        };
    }else{
        config = {
            headers: {'Content-Type': 'application/json'}
        };
    }
    return config;
}