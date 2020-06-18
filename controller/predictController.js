const socketService = require("../services/socketService")

var Predict = {}


Predict.Process = (req,res) =>{

    let body = req.body

    socketService.SendToWorker((r,data)=>{
        if(!r.success){
            r.m.send(r.success)
        }else
            r.send(data)
    },res,body)
}

module.exports = Predict