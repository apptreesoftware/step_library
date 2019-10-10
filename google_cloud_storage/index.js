let apptree = require('apptree-workflow-sdk');
const fs = require('fs');
const path = require('path');
const {Storage} = require('@google-cloud/storage');

apptree.addStep('upload_file', '1.0', uploadFile);
apptree.addStep('download_file', '1.0', downloadFile);
apptree.addStep('move_file', '1.0', moveFile);
apptree.run();


async function uploadFile(input) {
    apptree.validateInputs('Credentials', 'FilePath', 'ProjectId', 'Bucket');
    let credential = input['Credentials'];
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
    apptree.validateInputs('Credentials', 'FileName', 'ProjectId', 'Bucket', 'OutputDirectory');
    let credential = input['Credentials'];
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
    return {"Success": true};
}

async function moveFile(input) {
    apptree.validateInputs('Credentials', 'FileName', 'SourceBucket', 'DestinationBucket', 'ProjectId');
    let credential = input['Credentials'];
    let fileName = input['FileName'];
    const scBucketName = input['SourceBucket'];
    const destBucketName = input['DestinationBucket'];

    const storage = new Storage({
        projectId: input['ProjectId'],
        credentials: credential,
    });

    const sourceBucket = storage.bucket(scBucketName);
    let exists = await sourceBucket.exists();
    if (!exists) {
        throw "source bucket does not exist";
    }

    const destBucket = storage.bucket(destBucketName);
    exists = await destBucket.exists();
    if (!exists) {
        throw "destination bucket does not exist";
    }

    await sourceBucket.file(fileName).move(destBucket);
    return {"Success": true};
}