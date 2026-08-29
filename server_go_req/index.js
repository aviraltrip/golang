const express = require("express");

const app = express();
const port = 3000;
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send("Welcome to my server!");
});

app.get("/get", (req, res) => {
  res.status(200).json({
    message: "Hello from my server",
  });
});

app.post("/post", (req, res) => {
  const myJson = req.body;

  res.status(200).json({
    message: "POST request received",
    data: myJson,
  });
});

app.post("/postform", (req, res) => {
  res.status(200).json({
    message: "Form data received",
    data: req.body,
  });
});

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});