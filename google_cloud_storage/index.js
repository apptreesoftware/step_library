let apptree = require('apptree-workflow-sdk');
const fs = require('fs');
const path = require('path');
const {Storage} = require('@google-cloud/storage');

apptree.addStep('upload_file', '1.0', uploadFile);
apptree.addStep('download_file', '1.0', downloadFile);
apptree.run();


async function uploadFile(input) {
    apptree.validateInputs('Credential', 'FilePath', 'ProjectId', 'Bucket');
    let credential = input['Credential'];
    let filePath = input['FilePath'];
    const bucketName = input['Bucket'];
    const deleteOnUpload = input['DeleteOnUpload'];

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

    if (deleteOnUpload) {
        fs.unlink(filePath, (err) => {
            if (err) {
                console.error("file was uploaded but did not get deleted from the local filesystem");
            }
            console.info("file deleted");
        });
    }
    return {"Success" : true };
}

async function downloadFile(input) {
    apptree.validateInputs('Credential', 'FileName', 'ProjectId', 'Bucket', 'OutputDirectory');
    let credential = input['Credential'];
    let fileName = input['FileName'];
    const bucketName = input['Bucket'];
    const outputDirectory = input['OutputDirectory'];

    const storage = new Storage({
        projectId: input['ProjectId'],
        credentials: credential,
    });
    const bucket = storage.bucket(bucketName);
    const exists = await bucket.exists();
    if (!exists) {
        throw "Bucket does not exist";
    }

    await bucket.file(fileName).download({destination: `${outputDirectory}/${fileName}`});
    return {"Success": true}
}