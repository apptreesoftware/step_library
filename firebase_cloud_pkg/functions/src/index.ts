import * as functions from 'firebase-functions';
import * as admin from 'firebase-admin';
import * as apptreeio from 'apptreeio'
import {GetDocumentInput, JsonObject, QueryAndQueueInput, QueryInput, QueryParameters, UpsertInput} from "./models";

admin.initializeApp();

export const getDocument = functions.https.onRequest(async (req, resp) => {
    const inputs = req.body as GetDocumentInput;
    apptreeio.validateInputs(req, "RecordPath");

    try {
        const docRef = admin.firestore().doc(inputs.RecordPath);
        const snapshot = await docRef.get();
        if (!snapshot.exists) {
            resp.send({Success: false, Message: "not found"});
        } else {
            resp.send({
                Success: true,
                Record: snapshot.data() as JsonObject,
            });
        }
    } catch (e) {
        resp.status(500).send(`error encountered: ${e}`);
    }
});

export const query = functions.https.onRequest(async (req, resp) => {
    const inputs = req.body as QueryInput;
    apptreeio.validateInputs(req, "CollectionPath");

    try {
        const records = await queryCollection(inputs.CollectionPath, inputs.QueryParams);
        resp.send({"Records": records});
    } catch (e) {
        resp.status(500).send(`error encountered: ${e}`);
    }
});

export const queryAndQueue = functions.https.onRequest(async (req, resp) => {
    const client = apptreeio.createClient(req);
    apptreeio.validateInputs(req, "CollectionPath");

    const inputs = req.body as QueryAndQueueInput;

    try {
        const records = await queryCollection(inputs.CollectionPath, inputs.QueryParams);
        for (const object of records) {
            await client.SpawnWorkflow(inputs.Workflow, object);
        }
        resp.send({"Records": records});
    } catch (e) {
        resp.status(500).send(`error encountered: ${e}`);
    }
});

export const upsert = functions.https.onRequest(async (req, resp) => {
    const inputs = req.body as UpsertInput;
    apptreeio.validateInputs(req, "Record", "RecordPath");
    if (!inputs.Merge) {
        inputs.Merge = false;
    }

    try {
        const docRef = admin.firestore().doc(inputs.RecordPath);
        await docRef.set(inputs.Record, {merge: inputs.Merge});

        const updatedRef = await docRef.get();
        const doc = updatedRef.data() as JsonObject;
        resp.send({"Record": doc});
    } catch (e) {
        resp.status(500).send(`error encountered: ${e}`);
    }
});

async function queryCollection(collectionPath: string, queryParams?: QueryParameters[]): Promise<JsonObject[]> {
    const collection = admin.firestore().collection(collectionPath);
    let response = undefined;
    if (!queryParams || queryParams.length === 0) {
        response = await collection.get();
    } else {
        let collQuery: FirebaseFirestore.Query | undefined = undefined;
        for (const queryParam of queryParams) {
            if (!collQuery) {
                collQuery = collection.where(queryParam.FieldName, queryParam.Operator, queryParam.FieldValue);
            } else {
                collQuery = collQuery.where(queryParam.FieldName, queryParam.Operator, queryParam.FieldValue);
            }
        }

        if (!collQuery) {
            throw Error("internal server error occurred creating document query");
        }
        response = await collQuery.get();
    }
    return  response.docs.map(d => d.data() as JsonObject);
}


