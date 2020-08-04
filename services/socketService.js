const uuid = require('uuid')
const configs = require('../configurations/pass')
const authController = require('../controller/authController')
const SocketService = {}

var connecteds = {}
var workers = {}

var callers = {}

SocketService.Register = (http) =>{
    if(configs.Socket.use){    
        console.log("[+] Socket Started")
        var io = require('socket.io')(http);
        
        io.on('connection', function(socket){
            console.log('[.] A worker appears');
            connecteds[socket.id] = {p: socket}

            socket.emit('welcome',{msg:"welcome dude"})
            
            socket.on("register",(data)=>{
                if (authController.CheckWorkerToken(data.token)){
                    console.log(`[+] Worker (${data.id}) connected -> ${socket.id}`);
                    workers[data.id] = {s:socket,id:data.id,timestamp: (new Date().getTime())}
                    connecteds[socket.id].w = workers[data.id]
                    socket.emit("registered",{sid:socket.id})
                }
                else 
                    socket.disconnect()
            })


            socket.on('response', (data)=>{
                if(connecteds[socket.id].w != undefined){
                    let c = callers[data.id]
                    c.lambda({success:data.success, m:c.metadata},data)

                    let timestamp = (new Date().getTime())
                    console.log[`[+] Predicted: callerID [${data.id}] worker [${connecteds[socket.id].w.id}] at [${timestamp}]`]
                }else
                    socket.disconnect()
            })

            socket.on('disconnect', function(){
                let id = "anonymus"
                if (connecteds[socket.id].w != undefined){
                    id = connecteds[socket.id].w.id
                    workers[id] = null
                    delete workers[id]
                }
                connecteds[socket.id] = null
                delete connecteds[socket.id]
                
                console.log(`[-] Worker (${id}) disconnected -> ${socket.id}`);
            });
        });
    }
}

SocketService.SendToWorker = (lambda,metadata, payload) =>{
    let timestamp = (new Date().getTime())
    let ws = Object.keys(workers)
    if(ws.length <= 0){
        console.log(`[x] No have workers online [${timestamp}]`)
        return lambda({success:false,m:metadata},{})
    }

    // TODO: make load balance here
    let w = workers[ws.pop()]

    let callerID = uuid.v4()
    callers[callerID] = {lambda,metadata,timestamp}
    w.s.emit('handle',{id:callerID,data:payload.data, path:payload.url, method:payload.method, timestamp})

    console.log(`[+] Predicting: callerID [${callerID}] worker [${w.id}] at [${timestamp}]`)
}

SocketService.ListWorkers= () =>{
    return Object.keys(workers)
}

module.exports = SocketService
