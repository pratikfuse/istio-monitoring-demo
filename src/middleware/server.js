const http = require('http');
const { default: axios } = require('axios');

function handleError(error) {
    if (error.response) {
        // Request made and server responded
        console.log(error.response);
    } else if (error.request) {
        // The request was made but no response was received
        console.log(error.request);
    } else {
        // Something happened in setting up the request that triggered an Error
        console.log("fuck you")
    }
}

//create a server object:
http.createServer(function (req, res) {
    axios.get("http://backend-svc:4000").then(response => {
        console.log(response.data);
        res.setHeader("Content-Type", "application/json");
        res.write(`
            {
                "time": "${response.data.time}",
                "client_ip": "${response.data.client_ip}"
            }
        `);
        res.end();
    }).catch(err => {
        handleError(err)
        res.setHeader("Content-Type", "application/json");
        res.write(`
            {
                "error": "something went wrong"
            }
        `)
        res.end();
    })

}).listen(9000); //the server object listens on port 8080