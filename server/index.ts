import express from "express";

const app = express();

app.get("/", function (_, res) {
  res.send("Hello World");
});

const PORT = 3001;
app.listen(PORT, () => {
  console.log(`⚡️[server]: Server is running at http://localhost:${PORT}`);
});
