/*
* Log Dao
*/
const baseDAO = require('../dao.js');
const pass = require('../../configurations/pass')

var dao = {}
dao = Object.create(baseDAO);

dao.Save = log => new Promise((resolve,reject)=>{
    resolve(true)
})


module.exports = dao;