// src/index.js
import express, { Express, Request, Response} from "express";
import dotenv from "dotenv";
import cors from "cors";
import bodyParser, { text } from "body-parser";
import morgan from "morgan";

import fetch from './fetch.js';
import Cache from './cache.js';
import {stripMensa, stripCampus} from "./stripper.js";

dotenv.config();

const app: Express = express().use(cors({ origin: '*' })).use(bodyParser.json());
app.use(morgan('combined'))

const port = process.env.PORT || 3000;

const baseUrl = "https://www.imensa.de/";

const cache = new Cache();

app.get("/", (req: Request, res: Response) => {
  res.send("Mensa API");
});

app.get("/api/:Ort/:Mensa?", (req: Request, res: Response) => {
    if (req.params.Ort === null) {
        return res.send("Invalid request");
    }
    let cachedData = cache.get();
    if (cachedData !== null) {
        return res.send(cachedData);
    }else {
        let url = baseUrl + req.params.Ort.toLowerCase(); 
        fetch(url).then((data) => {
            let stripedData = null;
            if (req.params.Mensa !== null) {
                let stripedData = stripMensa(data);
            }else {
                let stripedData = stripCampus(data);
            }

            cache.set(req.params.Ort, stripedData);
            res.send(stripedData);
        });
    }

});

app.get("/api", (req: Request, res: Response) => {
    res.send("/api/:Ort/:Mensa?");
});

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});
