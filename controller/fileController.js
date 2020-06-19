const fs = require('fs')

const constants = require('../configurations/constants')

const FileController = {}

FileController.GetCertFile =  (req, res) => {
    res.setHeader('Content-Type','application/pdf')
    
    let file = req.params.id
    file = file.toString().split(".")
    fs.createReadStream( constants.Paths.Upload+"/"+file[0]).pipe(res)
    //res.download(constants.Paths.Upload+"/"+file[0],req.params.id)
}

FileController.GetCertFile2 =  (req, res) => {
    res.setHeader('Content-Type','application/pdf')
    
    let file = req.params.id
    file = file.toString().split(".")
    //fs.createReadStream( constants.Paths.Upload+"/"+file[0]).pipe(res)
    res.download(constants.Paths.Upload+"/"+file[0],req.params.id)
}

module.exports = FileController