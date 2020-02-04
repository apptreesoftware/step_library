import * as functions from 'firebase-functions';
import * as admin from 'firebase-admin';
import * as apptreeio from 'apptreeio'
import {GetDocumentInput, JsonObject, QueryAndQueueInput, QueryInput, QueryParameters, UpsertInput} from "./models";

admin.initializeApp();

export const getDocument = functions.https.onRequest(async (req, resp) => {
    const inputs = req.body as GetDocumentInput;
    if (!inputs.RecordPath || inputs.RecordPath == "") {
        resp.status(500).send('record path is a required input');
        return;
    }

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
    if (!inputs.CollectionPath || inputs.CollectionPath === "") {
        resp.status(500).send(`collection path is a required input`);
        return;
    }

    try {
        const records = await queryCollection(inputs.CollectionPath, inputs.QueryParams);
        resp.send({"Records": records});
    } catch (e) {
        resp.status(500).send(`error encountered: ${e}`);
    }
});

export const queryAndQueue = functions.https.onRequest(async (req, resp) => {
    const client = apptreeio.createClient(req);

    const inputs = req.body as QueryAndQueueInput;
    if (!inputs.CollectionPath || inputs.CollectionPath === "") {
        resp.status(500).send(`collection path is a required input`);
        return;
    }

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
    if (!inputs.Record) {
        resp.status(500).send(`record to upsert is a required input`);
        return;
    }
    if (!inputs.RecordPath) {
        resp.status(500).send(`record path is a required input`);
        return;
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


