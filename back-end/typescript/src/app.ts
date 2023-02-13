import express from "express";
import { authRouter } from "./auth.controllers";
import { userRouter } from "./user.controllers";
var app = express();

const port = 3000;

app.get("/", function (req, res) {
  res.send("hello world");
});

app.use("/auth", authRouter);

app.use("/user", userRouter);

app.listen(port, () => {
  console.log(`Running at http://localhost:${port}/`);
});
