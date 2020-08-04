const socketService = require("../services/socketService")

var Predict = {}


Predict.Process = (req,res) =>{
    let body = req.body
   
    socketService.SendToWorker((r,data)=>{
        if(!r.success){
            r.m.send(r.success)
        }else
            r.m.send(data.data)
    },res,{data:body, url:req.url, method:req.method})
}

Predict.ListWorkers = (req,res) =>{
    res.send(socketService.ListWorkers())
}

module.exports = Predict