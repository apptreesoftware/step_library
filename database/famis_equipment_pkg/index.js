let apptree = require('apptree-workflow-sdk');
const {Client} = require('pg');

apptree.addStep('equipment_lookup', '1.0', equipmentLookup);

apptree.run();

async function equipmentLookup(input) {
    apptree.validateInputs("ConnectionString");
    let connectionStr = input["ConnectionString"];
    console.log(`Connecting to DB ${connectionStr}`);
    const client = new Client({
        connectionString: connectionStr
    });
    await client.connect();
    console.log(`Connecting Established`);
    let equipmentType = input['EquipmentType'];
    let equipmentSubType = input['EquipmentSubType'];
    let equipmentSubType2 = input['EquipmentSubType2'];
    if (equipmentType == null || equipmentType.length === 0) {
        console.log(`Querying for equipment type`);
        const result = await client.query(`
            SELECT distinct xref.va_equip_type,
                            case
                                when va_equip_subtype1 = '%'
                                    then workflow_id
                                else null
                                end workflow_id
            from at_equip_ts_xref xref
            order by 1;
        `);
        let options = result.rows.map((r) => r['va_equip_type'].trim());
        return {Options: options, Workflow: null};
    } else if (equipmentSubType == null || equipmentSubType.length === 0) {
        let result = await client.query(`SELECT workflow_id FROM  at_equip_ts_xref WHERE va_equip_type = '${equipmentType}'`);
        if (result.rowCount === 1) {
            const workflow = result.rows[0]['workflow_id'];
            return {Workflow: workflow, Options: []}
        }
        result = await client.query(`
            SELECT distinct xref.va_equip_type,
                            xref.va_equip_subtype1,
                            case
                                when va_equip_subtype2 = '%'
                                    then workflow_id
                                else null
                                end workflow_id
            from at_equip_ts_xref xref
            where xref.va_equip_subtype1 != '%'
              and xref.va_equip_type = '${equipmentType}'
            order by 1, 2;`);
        const options = result.rows.map((r) => r['va_equip_subtype1']);
        return {Options: options, Workflow: null};
    } else if (equipmentSubType2 == null || equipmentSubType2.length === 0) {
        let result = await client.query(`
        SELECT workflow_id FROM  at_equip_ts_xref 
        WHERE va_equip_type = '${equipmentType}' AND va_equip_subtype1 = '${equipmentSubType}'`);
        if (result.rowCount === 1) {
            const workflow = result.rows[0]['workflow_id'];
            return {Workflow: workflow, Options: []}
        }
        result = await client.query(`
            SELECT distinct xref.va_equip_type,
                            xref.va_equip_subtype1,
                            xref.va_equip_subtype2,
                            xref.workflow_id
            from at_equip_ts_xref xref
            where xref.va_equip_subtype2 != '%'
              and xref.va_equip_type = '${equipmentType}'
              and xref.va_equip_subtype1 = '${equipmentSubType}'
            order by 1, 2, 3;
        `);
        const options = result.rows.map((r) => r['va_equip_subtype2']);
        return {Options: options, Workflow: null};
    } else {
        const result = await client.query(`
            SELECT distinct xref.va_equip_type,
                    xref.va_equip_subtype1,
                    xref.va_equip_subtype2,
                    xref.workflow_id
            from at_equip_ts_xref xref
            where xref.va_equip_subtype2 = '${equipmentSubType2}'
                and xref.va_equip_type = '${equipmentType}'
                and xref.va_equip_subtype1 = '${equipmentSubType}'
            order by 1, 2, 3;`);
        if (result.rowCount === 1) {
            const workflow = result.rows[0]['workflow_id'];
            return {Workflow: workflow, Options: []}
        } else {
            throw `Could not find a workflow for equipment ${equipmentType} -> ${equipmentSubType} -> ${equipmentSubType2}`;
        }
    }
}
