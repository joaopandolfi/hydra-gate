const uuid = require('uuid')
const configs = require('../configurations/pass')
const authController = require('../controller/authController')
const SockerService = {}

var connecteds = {}
var workers = {}

var callers = {}

SockerService.Register = (http) =>{
    if(configs.Socket.use){    
        console.log("[+] Socket Started")
        var io = require('socket.io')(http);
        
        io.on('connection', function(socket){
            console.log('[.] A worker appears');
            connecteds[socket.id] = {p: socket}

            socket.emit("ping",{msg:"welcome dude"})

            socket.on('pong',(data)=>{
                if (authController.CheckWorkerToken(data.token)){
                    console.log(`[+] Worker (${data.id}) connected -> ${socket.id}`);
                    workers[data.id] = {s:socket,id:data.id,timestamp: (new Date().getTime())}
                    connecteds[socket.id].w = workers[data.id]
                }
                else 
                    socket.disconnect()
            })


            spcket.on('predicted', (data)=>{
                if(connecteds[socket.id].w != undefined){
                    let c = callers[callerID]
                    c.lambda({success:true, m:metadata},data)

                    let timestamp = (new Date().getTime())
                    console.log[`[+] Predicted: callerID [${data.callerID}] worker [${connecteds[socket.id].w.id}] at [${timestamp}]`]
                }else
                    socket.disconnect()
            })

            socket.on('disconnect', function(){
                let id = connecteds[socket.id].w.id
                connecteds[socket.id] = null
                workers[id] = null
                delete connecteds[socket.id]
                delete workers[id]
                console.log(`[-] Worker (${id}) disconnected -> ${socket.id}`);
            });
        });
    }
}

SockerService.SendToWorker = (lambda,metadata, data) =>{
    let timestamp = (new Date().getTime())
    let ws = Object.keys(workers)
    if(ws.length <= 0){
        console.log(`[x] No have workers online [${timestamp}]`)
        return lambda({success:false,m:metadata},{})
    }

    let w = workers[ws.pop()]

    let callerID = uuid.v4()
    callers[callerID] = {lambda,metadata,timestamp}
    w.s.emit('predict',{data, timestamp})

    console.log(`[+] Predicting: callerID [${callerID}] worker [${w.id}] at [${timestamp}]`)
}

module.exports = SockerService
