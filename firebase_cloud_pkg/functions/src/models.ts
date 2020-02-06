

export interface QueryInput {
    CollectionPath: string;
    QueryParams: QueryParameters[];
}

export interface QueryAndQueueInput {
    CollectionPath: string;
    QueryParams: QueryParameters[];
    Workflow: string;
}

export interface UpsertInput {
    Record: JsonObject;
    RecordPath: string;
    Merge?: boolean;
}

export interface GetDocumentInput {
    RecordPath: string;
}

export interface QueryParameters {
    FieldName: string;
    Operator: FirebaseFirestore.WhereFilterOp;
    FieldValue: string;
}

export type JsonObject = { [prop: string]: any };