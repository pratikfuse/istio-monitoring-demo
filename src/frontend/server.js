const { default: axios } = require('axios');
const express = require('express');
const handlebars = require('express-handlebars');

const app = express();

app.set('view engine', 'handlebars');

app.engine('handlebars', handlebars({
    layoutsDir: __dirname + '/views',
}));

app.use(express.static("public"));

app.get("/", async function (req, res) {
    try {
        const { data } = await axios.get("http://middleware-svc:9000");
        console.log(data);
        res.render('main', { time: data.time, client_ip: data.client_ip })
    } catch (error) {
        console.log(error);
        if (error.response) {
            // Request made and server responded
            res.render('main', { error: error.response.data })
        } else if (error.request) {
            // The request was made but no response was received
            res.render('main', { error: "No response from server" })
        } else {
            // Something happened in setting up the request that triggered an Error
            res.render('main', { error: "something went wrong" })
        }

    }
})

app.listen(3000, () => {
    console.log("running ")
});