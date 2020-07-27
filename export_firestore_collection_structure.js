// Exports the structure of the specified collections. (Assumes all documents inside a given collection uses the same structure)
// Change out PROJECT_ID for your GCP/Firebase project.
// Make sure you are authenticated via the firebase CLI tool (`firebase login`)
// Run with `node export_firestore_collection_structure.js`
'use strict';
const fs = require('fs');
const admin = require('firebase-admin');
admin.initializeApp({projectId: 'PROJECT_ID'});
let db = admin.firestore();
const cols = ['collection1', 'collection2'];
const tables = {};

const extractDocStructure = (d, col) => {
  const metaObj = {};
  const dig = v => {
    if (typeof v !== 'object') {
      return typeof v;
    } else if (Object.keys(v)[0] === '0') {
      return [dig(v['0'])];
    } else if (Object.keys(v).length === 0) {
      return [];
    } else if (Object.keys(v)[0] === '_seconds') {
      return 'timestamp';
    } else {
      const obj = {};
      Object.keys(v).forEach(k => (obj[k] = dig(v[k])));
      return obj;
    }
  };
  Object.keys(d).forEach(k => (metaObj[k] = dig(d[k])));
  tables[col] = metaObj;
};

Promise.all(
  cols.map(col =>
    db
      .collection(col)
      .limit(1)
      .get()
      .then(snap => snap.docs[0])
      .then(d => ({id: d.id, ...d.data()}))
      .then(d => extractDocStructure(d, col))
      .catch(error => console.error(col, error))
  )
).then(_ => {
  const file = fs.createWriteStream('database.json');
  file.write(JSON.stringify(tables, null, 2));
});
