import pkg from 'sqlite3';
const { Database } = pkg;

export const db = new Database('../db.sqlite3');
