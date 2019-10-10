let apptree = require('apptree-workflow-sdk');
let path = require('path');

const {Storage} = require('@google-cloud/storage');
apptree.addStep('upload_file', '1.0', uploadFile);
apptree.run();


async function uploadFile(input) {
    apptree.validateInputs('Credential', 'FilePath', 'ProjectId', 'Bucket');
    let credential = input['Credential'];
    let filePath = input['FilePath'];
    const bucketName = input['Bucket'];
    const storage = new Storage({
        projectId: input['ProjectId'],
        credentials: credential,
    });
    const bucket = storage.bucket(bucketName);
    const exists = await bucket.exists();
    if (!exists) {
        await bucket.create();
    }
    console.log(`Uploading ${filePath}`);
    await bucket.upload(filePath, {

    });
    return {"Success" : true };
}