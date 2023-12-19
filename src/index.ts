// src/index.js
import express, { Express, Request, Response} from "express";
import dotenv from "dotenv";
import cors from "cors";
import bodyParser from "body-parser";
import morgan from "morgan";

import fetch from './fetch.js';

dotenv.config();

const app: Express = express().use(cors({ origin: '*' })).use(bodyParser.json());
app.use(morgan('combined'))

const port = process.env.PORT || 3000;

const baseUrl = "https://www.imensa.de/";

app.get("/", (req: Request, res: Response) => {
  res.send("Mensa API");
});

app.get("/api/:Ort/:Mensa", (req: Request, res: Response) => {
    if (req.params.Ort === null) {
        return res.send("Invalid request");
    }
    let url = baseUrl + req.params.Ort.toLowerCase(); 
    fetch(url).then((data) => {
        res.send(data);
    });

});

app.get("/api", (req: Request, res: Response) => {
    res.send("/api/:Ort/:Mensa");
});

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});
