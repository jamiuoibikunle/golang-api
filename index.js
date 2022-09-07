const app = require("express")();
const fs = require("fs");

const PORT = process.env.PORT || "8080";


app.use("/", (req, res) => {
    fs.readFile("./json/states.json", "utf8", (err, data) => {
        res.send({ message: data })
    })
})

app.listen(PORT, () => {
    console.log("Listening on port " + PORT);
});
